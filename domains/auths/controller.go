package auths

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var service = NewAuthService()

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type authController struct {
}

func NewAuthController() AuthController {
	return &authController{}
}

func (a *authController) Login(c *gin.Context) {
	/* Get params body */
	logrus.Info("Login")

	/** Get login request**/
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	/** Call service **/
	data, httpStatus, err := service.Login(request)
	if err != nil {
		c.JSON(httpStatus, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Login",
		"data":    data,
	})
}

func (a *authController) Register(c *gin.Context) {
	logrus.Info("Register")

	/** Get register request**/
	var request RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	/** Call service **/
	data, headerStatus, err := service.Register(request)
	if err != nil {
		c.JSON(headerStatus, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Register Success",
		"data":    data,
	})
}
