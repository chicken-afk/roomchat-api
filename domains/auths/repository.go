package auths

import (
	"goboilerplate/commons"
	"goboilerplate/database"
	"goboilerplate/entities"
)

var db = database.SetupDatabaseConnection()

type AuthRepository interface {
	CreateUser(request RegisterRequest) (interface{}, error)
	Login(request LoginRequest) (interface{}, error)
}

type authRepository struct{}

func NewAuthRepository() AuthRepository {
	return &authRepository{}
}

func (a *authRepository) CreateUser(request RegisterRequest) (interface{}, error) {
	newUser := new(entities.User)
	newUser.Email = request.Email
	//Hash encrypt password
	newUser.Password, _ = commons.EncryptPassword(request.Password)
	db.Create(&newUser)
	return newUser, nil
}

func (a *authRepository) Login(request LoginRequest) (interface{}, error) {
	return nil, nil
}
