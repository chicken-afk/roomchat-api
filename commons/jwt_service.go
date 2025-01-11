package commons

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/twinj/uuid"
)

type JwtService interface {
	GenerateToken(userId uint64, identifierId uint64) (*TokenDetail, error)
	ValidateJwtToken(r *http.Request) (*UserValidateDTO, error)
	VerifyToken(r *http.Request) (*jwt.Token, error)
	ExtractToken(r *http.Request) (bool, string)
	ValidateRefreshToken(tokenString string) (*RefreshTokenValidateDTO, error)
	VerifyRefreshToken(tokenString string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewJwtService() JwtService {
	return &jwtService{}
}

type TokenDetail struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	AccessUuid   string `json:"accessUuid"`
	RefreshUuid  string `json:"refreshUuid"`
	AtExpires    int64  `json:"atExpires"`
	RtExpires    int64  `json:"rtExpires"`
}

type UserValidateDTO struct {
	Authorized   string `json:"authorized"`
	AccessUuid   string `json:"accessUuid"`
	UserId       uint64 `json:"userId"`
	IdentifierId uint64 `json:"identifierId"`
	Exp          string `json:"exp"`
}

type RefreshTokenValidateDTO struct {
	RefreshUuid string `json:"refreshUuid"`
	UserId      uint64 `json:"userId"`
	Exp         string `json:"exp"`
}

func (j jwtService) GenerateToken(userId uint64, identifierId uint64) (*TokenDetail, error) {
	tokenDetail := &TokenDetail{}
	duration, err := strconv.Atoi(os.Getenv("TOKEN_DURATION"))
	if err != nil {
		duration = 30
	}

	tokenDetail.AtExpires = time.Now().Add(time.Minute * time.Duration(duration)).Unix()
	tokenDetail.AccessUuid = uuid.NewV4().String()

	tokenDetail.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	tokenDetail.RefreshUuid = uuid.NewV4().String()

	/*Creating Access Token*/
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["accessUuid"] = tokenDetail.AccessUuid
	atClaims["userId"] = userId
	atClaims["identifierId"] = identifierId
	atClaims["exp"] = tokenDetail.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tokenDetail.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	/*Creating Refresh Token*/
	rtClaims := jwt.MapClaims{}
	rtClaims["refreshUuid"] = tokenDetail.RefreshUuid
	rtClaims["userId"] = userId
	rtClaims["exp"] = tokenDetail.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	tokenDetail.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return tokenDetail, nil
}

func (j jwtService) ValidateJwtToken(r *http.Request) (*UserValidateDTO, error) {
	token, err := j.VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {

		accessUuid := claims["accessUuid"].(string)
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["userId"]), 10, 64)
		if err != nil {
			return nil, err
		}
		identifierId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["identifierId"]), 10, 64)
		if err != nil {
			return nil, err
		}

		exp := claims["exp"].(float64)
		expInt := int64(exp)
		tm := time.Unix(expInt, 0)
		timeStr := tm.String()

		return &UserValidateDTO{
			Authorized:   "true",
			AccessUuid:   accessUuid,
			UserId:       userId,
			IdentifierId: identifierId,
			Exp:          timeStr,
		}, nil
	}
	return nil, err
}

func (j jwtService) VerifyToken(r *http.Request) (*jwt.Token, error) {
	_, tokenString := j.ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (j jwtService) ExtractToken(r *http.Request) (bool, string) {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return true, strArr[1]
	}
	return false, bearToken
}

// refresh token
func (j jwtService) ValidateRefreshToken(tokenString string) (*RefreshTokenValidateDTO, error) {
	token, err := j.VerifyRefreshToken(tokenString)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUuid := claims["refreshUuid"].(string)
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["userId"]), 10, 64)
		if err != nil {
			return nil, err
		}

		exp := claims["exp"].(float64)
		expInt := int64(exp)
		tm := time.Unix(expInt, 0)
		timeStr := tm.String()

		return &RefreshTokenValidateDTO{
			RefreshUuid: refreshUuid,
			UserId:      userId,
			Exp:         timeStr,
		}, nil
	}
	return nil, err
}

func (j jwtService) VerifyRefreshToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
