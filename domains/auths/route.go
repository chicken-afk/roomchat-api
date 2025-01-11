package auths

import (
	"github.com/gin-gonic/gin"
)

var controller = NewAuthController()

func Router(r *gin.RouterGroup) *gin.RouterGroup {
	/* Router */
	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)

	return r
}
