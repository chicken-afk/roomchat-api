package users

import (
	"goboilerplate/entities"
	"net/http"
)

type UserService interface {
	GetUserById(id uint64) (entities.User, int, error)
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (u *userService) GetUserById(id uint64) (entities.User, int, error) {
	// Call Repository
	userRepo := NewUserRepository()

	var user entities.User
	err := userRepo.GetUserById(id, &user)
	if err != nil {
		return user, http.StatusInternalServerError, err
	}

	return user, http.StatusOK, nil
}
