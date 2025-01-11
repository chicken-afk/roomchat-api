package auths

import (
	"errors"
	"goboilerplate/commons"
	"goboilerplate/domains/users"
	"goboilerplate/entities"
	"net/http"

	"github.com/sirupsen/logrus"
)

var userRepo = users.NewUserRepository()

type AuthService interface {
	Login(request LoginRequest) (interface{}, error)
	Register(request RegisterRequest) (interface{}, int, error)
}

type aService struct {
}

func NewAuthService() AuthService {
	return &aService{}
}

func (a *aService) Login(request LoginRequest) (interface{}, error) {
	/**Call Repository**/

	var response LoginResponse
	response.Token = "ewafaofjds920r3jsda=29"
	response.Uuid = "24802-4534-klasd-3j329"
	response.Email = request.Email

	return response, nil
}

func (a *aService) Register(request RegisterRequest) (interface{}, int, error) {
	/**Call Repository**/
	var response RegisterResponse
	var user entities.User
	user.Email = request.Email
	hashPassword, err := commons.EncryptPassword(request.Password)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	user.Password = hashPassword

	//Check if user already exist
	var existUser entities.User
	err = userRepo.GetUserByEmail(user.Email, &existUser)
	logrus.Info("Exist User")
	logrus.Info(existUser)
	logrus.Info(err)

	if existUser.ID != 0 {
		logrus.Info("User already exist")
		err := errors.New("user already exist")
		return nil, http.StatusUnprocessableEntity, err
	}
	logrus.Info("creating new user")
	errRepo := userRepo.CreateUser(&user)
	if errRepo != nil {
		return nil, http.StatusInternalServerError, errRepo
	}

	response.Email = request.Email
	response.Status = "Active"
	response.Token = "ewafaofjds920r3jsda=29"
	response.ID = user.ID

	return response, http.StatusOK, nil
}
