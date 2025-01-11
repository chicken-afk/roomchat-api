package middleware

import (
	"goboilerplate/commons"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Middleware interface {
	HeaderAuth() gin.HandlerFunc
	Auth() gin.HandlerFunc
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

func (m *middleware) Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtServ := commons.NewJwtService()
		token, err := jwtServ.ValidateJwtToken(ctx.Request)
		if err != nil {
			logrus.Error(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized, token invalid",
			})
			return
		}

		ctx.Set("token", token)
		ctx.Next()
	}
}
