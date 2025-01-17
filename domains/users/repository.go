package users

import (
	"chatroom-api/database"
	"chatroom-api/entities"
)

var db = database.SetupDatabaseConnection()

type UserRepository interface {
	CreateUser(user *entities.User) error
	GetUserByEmail(email string, user *entities.User) error
	GetUserById(id uint64, user *entities.User) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (u *userRepository) CreateUser(user *entities.User) error {
	err := db.Create(&user)
	if err.Error != nil {
		return err.Error
	}

	return nil
}

func (u *userRepository) GetUserByEmail(email string, user *entities.User) error {
	err := db.Where("email = ?", email).First(&user)
	if err.Error != nil {
		return err.Error
	}

	return nil
}

func (u *userRepository) GetUserById(id uint64, user *entities.User) error {
	err := db.Where("id = ?", id).First(&user)
	if err.Error != nil {
		return err.Error
	}

	return nil
}
