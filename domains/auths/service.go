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
	Login(request LoginRequest) (interface{}, int, error)
	Register(request RegisterRequest) (interface{}, int, error)
}

type aService struct {
}

func NewAuthService() AuthService {
	return &aService{}
}

func (a *aService) Login(request LoginRequest) (interface{}, int, error) {
	/**Call Repository**/
	userRepo := users.NewUserRepository()

	/**Get user by email**/
	var user entities.User
	err := userRepo.GetUserByEmail(request.Email, &user)
	if err != nil && err.Error() != "record not found" {
		return nil, http.StatusInternalServerError, err
	}
	if user.ID == 0 {
		return nil, http.StatusUnprocessableEntity, errors.New("user not found")
	}
	/**validate password**/
	isCorrectPassword := commons.VerifyPassword(user.Password, request.Password)
	if !isCorrectPassword {
		return nil, http.StatusUnprocessableEntity, errors.New("password not match")
	}
	/**Generate token**/
	jwtServ := commons.NewJwtService()
	token, err := jwtServ.GenerateToken(uint64(user.ID), uint64(user.ID))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var response LoginResponse
	response.Token = *token
	response.Email = request.Email
	response.ID = user.ID
	response.Status = "Active"

	return response, http.StatusOK, nil
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
	response.ID = user.ID
	//Generate token JWT
	jwtServ := commons.NewJwtService()
	token, err := jwtServ.GenerateToken(uint64(user.ID), uint64(user.ID))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	response.Token = *token

	return response, http.StatusOK, nil
}
