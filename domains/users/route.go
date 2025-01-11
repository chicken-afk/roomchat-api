package users

import (
	"goboilerplate/middleware"

	"github.com/gin-gonic/gin"
)

var controller = NewUserController()
var routeMiddleware = middleware.NewMiddleware()

func Router(r *gin.RouterGroup) *gin.RouterGroup {
	/* Router */
	r.GET("/profile", routeMiddleware.Auth(), controller.Profile)

	return r
}
