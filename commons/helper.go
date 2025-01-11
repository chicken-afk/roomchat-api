package commons

import (
	"os"

	"github.com/sirupsen/logrus"
)

func ValidateHeaderToken(token string) bool {
	logrus.Info("ValidateHeaderToken")
	logrus.Info("Token from env:", os.Getenv("API_KEY"))
	return token == os.Getenv("API_KEY")
}
