package users

import (
	"goboilerplate/commons"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Profile(c *gin.Context)
}

type userController struct {
}

func NewUserController() UserController {
	return &userController{}
}

func (u *userController) Profile(c *gin.Context) {
	/**Get jwt data**/
	tokenData, err := commons.GetTokenFromMiddleware(c)
	if err != nil {
		c.JSON(401, gin.H{
			"success": false,
			"message": "Unauthorized",
		})
		return
	}

	//Call service
	userServ := NewUserService()
	user, httpStatus, err := userServ.GetUserById(tokenData.UserId)
	if err != nil {
		c.JSON(httpStatus, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	//Response
	res := new(ProfileResponse)
	res.ID = user.ID
	res.Email = user.Email
	res.CreatedAt = user.CreatedAt
	res.UpdatedAt = user.UpdatedAt
	res.Status = "Active"

	/* Get user from context */
	c.JSON(200, gin.H{
		"success": true,
		"message": "Success Get profile",
		"data":    res,
	})
	return
}
