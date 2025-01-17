package commons

import (
	"errors"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func ValidateHeaderToken(token string) bool {
	logrus.Info("ValidateHeaderToken")
	logrus.Info("Token from env:", os.Getenv("API_KEY"))
	return token == os.Getenv("API_KEY")
}

func EncryptPassword(password string) (string, error) {
	logrus.Info("EncryptPassword")

	// Generate a bcrypt hash of the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Error("Error encrypting password: ", err)
		return "", err
	}

	return string(hashedPassword), nil
}

// VerifyPassword checks if the provided plain text password matches the hashed password
func VerifyPassword(hashedPassword, password string) bool {
	logrus.Info("VerifyPassword")

	// Compare the hashed password with the plain text password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		logrus.Warn("Password verification failed: ", err)
		return false
	}

	return true
}

func DialogError(message string, httpStatus int, ctx *gin.Context) {
	//ctx.JSON(httpStatus, res)
	ctx.AbortWithStatusJSON(httpStatus, gin.H{
		"success": false,
		"message": message,
	})
}

func GetTokenFromMiddleware(ctx *gin.Context) (*UserValidateDTO, error) {
	token, exists := ctx.Get("token")
	if !exists {
		return nil, errors.New("token not exist from middleware")
	}

	tokenData, ok := token.(*UserValidateDTO)
	if !ok {
		return nil, errors.New("err when binding token")
	}

	return tokenData, nil
}

func RandomRoomchatCode() string {
	logrus.Info("RandomRoomchatCode")
	// Generate a random room code

	return "ROOM" + RandomString(6)
}

// RandomString generates a random string of the specified length.
func RandomString(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomStr := make([]byte, length)
	for i := range randomStr {
		randomStr[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(randomStr)
}
