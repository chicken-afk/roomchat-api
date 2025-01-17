package roomchats

import (
	"chatroom-api/domains/users"
	"chatroom-api/entities"
)

type RoomchatService interface {
	CreateRoomchat(roomchat RoomchatRequest) (entities.Roomchat, error)
	JoinRoomchat(email string, roomId uint64) error
}

type roomchatService struct{}

func NewRoomchatService() RoomchatService {
	return &roomchatService{}
}

func (r *roomchatService) CreateRoomchat(roomchat RoomchatRequest) (entities.Roomchat, error) {
	//Call Repo
	roomchatRepo := NewRoomchatRepository()
	var roomchatEntity entities.Roomchat
	roomchatEntity.RoomName = roomchat.RoomName
	roomchatEntity.CreatedBy = roomchat.CreatedBy
	err := roomchatRepo.CreateRoomchat(roomchatEntity)
	if err != nil {
		return roomchatEntity, err
	}
	return roomchatEntity, nil
}

func (r *roomchatService) JoinRoomchat(email string, roomId uint64) error {
	//Call Repo
	roomchatRepo := NewRoomchatRepository()
	userRepo := users.NewUserRepository()
	var userEntity entities.User
	err := userRepo.GetUserByEmail(email, &userEntity)
	if err != nil {
		return err
	}
	err = roomchatRepo.JoinRoomchat(userEntity.ID, roomId)
	if err != nil {
		return err
	}
	return nil
}
