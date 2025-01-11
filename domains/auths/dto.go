package auths

import "goboilerplate/commons"

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Uuid  string `json:"uuid"`
	Email string `json:"email"`
}

type RegisterResponse struct {
	ID     int                 `json:"uuid"`
	Email  string              `json:"email"`
	Status string              `json:"status"`
	Token  commons.TokenDetail `json:"token"`
}
