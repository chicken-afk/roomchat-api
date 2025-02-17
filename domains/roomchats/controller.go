package roomchats

import (
	"chatroom-api/commons"
	"chatroom-api/domains/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoomchatController interface {
	StartRoomchat(c *gin.Context)
	SendMessage(c *gin.Context)
	GetRoomchats(c *gin.Context)
}

type roomchatController struct{}

func NewRoomchatController() RoomchatController {
	return &roomchatController{}
}

func (r *roomchatController) StartRoomchat(c *gin.Context) {
	tokenData, err := commons.GetTokenFromMiddleware(c)
	if err != nil {
		c.JSON(401, gin.H{
			"success": false,
			"message": "Unauthorized",
		})
		return
	}
	//Get User From tokenid
	userServ := users.NewUserService()
	userLogin, _, _ := userServ.GetUserById(tokenData.UserId)

	//bind request
	var request StartRoomchatRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	//Check if the email already registered
	userRequest, httpStatus, err := userServ.GetUserByEmail(request.Email)
	if err != nil {
		if httpStatus == 404 {
			c.JSON(404, gin.H{
				"success": false,
				"message": "User not found",
			})
			return
		}
		c.JSON(httpStatus, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	//Check if the email is the same as the login user
	if userLogin.Email == request.Email {
		c.JSON(400, gin.H{
			"success": false,
			"message": "You can't chat with yourself",
		})
		return
	}

	roomchatServ := NewRoomchatService()
	//Check if the roomchat already exist
	var userIds []int64
	userIds = append(userIds, int64(userLogin.ID))
	userIds = append(userIds, int64(userRequest.ID))
	roomchat, httpStatus, err := roomchatServ.GetRoomchatByUserId(userIds)
	if httpStatus == http.StatusInternalServerError {
		c.JSON(httpStatus, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	if httpStatus == http.StatusOK {
		c.JSON(httpStatus, gin.H{
			"success": true,
			"message": "Success Create or Get Roomchat",
			"data":    roomchat,
		})
		return
	}
	//Call service
	var roomName = commons.RandomRoomchatCode()

	var roomchatData RoomchatRequest
	roomchatData.RoomName = roomName
	roomchatData.CreatedBy = tokenData.UserId
	res, err := roomchatServ.CreateRoomchat(roomchatData)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	//Join user email to roomchat
	_, err = roomchatServ.JoinRoomchat(request.Email, res.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return

	}
	//Join user login to roomchat
	_, err = roomchatServ.JoinRoomchat(userLogin.Email, res.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	/* Get user from context */
	c.JSON(200, gin.H{
		"success": true,
		"message": "Success Create or Get Roomchat",
		"data":    res,
	})
	return
}

func (r *roomchatController) SendMessage(c *gin.Context) {
	tokenData, err := commons.GetTokenFromMiddleware(c)
	if err != nil {
		c.JSON(401, gin.H{
			"success": false,
			"message": "Unauthorized",
		})
		return
	}
	//Get User From tokenid
	userServ := users.NewUserService()
	userLogin, _, _ := userServ.GetUserById(tokenData.UserId)

	/* Get user from context */
	c.JSON(200, gin.H{
		"success": true,
		"message": "Success Send Message",
		"data":    userLogin,
	})
	return
}

func (r *roomchatController) GetRoomchats(c *gin.Context) {
	tokenData, err := commons.GetTokenFromMiddleware(c)
	if err != nil {
		c.JSON(401, gin.H{
			"success": false,
			"message": "Unauthorized",
		})
		return
	}
	//Get User From tokenid
	userServ := users.NewUserService()
	userLogin, _, _ := userServ.GetUserById(tokenData.UserId)

	roomchatServ := NewRoomchatService()
	//Get roomchat by user id
	roomchat, err := roomchatServ.GetRoomchatUsers(int64(userLogin.ID))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	/* Get user from context */
	c.JSON(200, gin.H{
		"success": true,
		"message": "Success Get Roomchat",
		"data":    roomchat,
	})
	return
}
