package middleware

import (
	"goboilerplate/commons"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Middleware interface {
	HeaderAuth() gin.HandlerFunc
}

type middleware struct{}

func NewMiddleware() Middleware {
	return &middleware{}
}

func (m *middleware) HeaderAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logrus.Info("HeaderAuth")
		requestHeaderToken := ctx.GetHeader("x-api-key") // Get header token
		logrus.Info("Header token:", requestHeaderToken)
		success := commons.ValidateHeaderToken(requestHeaderToken)
		if !success {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized",
			})
			return
		}
		ctx.Next()
	}
}
