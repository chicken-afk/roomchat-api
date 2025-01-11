package commons

import (
	"os"

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

func DialogError(message string, httpStatus int, ctx *gin.Context) {
	//ctx.JSON(httpStatus, res)
	ctx.AbortWithStatusJSON(httpStatus, gin.H{
		"success": false,
		"message": message,
	})
}
