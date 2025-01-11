package commons

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/twinj/uuid"
)

type JwtService interface {
	GenerateToken(userId uint64, identifierId uint64) (*TokenDetail, error)
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
