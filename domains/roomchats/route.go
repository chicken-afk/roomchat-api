package roomchats

import (
	"chatroom-api/middleware"

	"github.com/gin-gonic/gin"
)

var roomController = NewRoomchatController()
var routeMiddleware = middleware.NewMiddleware()

func Router(r *gin.RouterGroup) *gin.RouterGroup {
	/* Router */
	r.POST("/create-roomchat", routeMiddleware.Auth(), roomController.StartRoomchat)

	return r
}
