package roomchats

import (
	"chatroom-api/commons"
	"chatroom-api/domains/users"
	"chatroom-api/entities"
)

type RoomchatService interface {
	CreateRoomchat(roomchat RoomchatRequest) (entities.Roomchat, error)
	JoinRoomchat(email string, roomId uint64) (entities.RoomchatUser, error)
	GetRoomchatByUserId(userIds []int64) (entities.Roomchat, int, error)
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
	roomchatEntity.RoomId = commons.RandomRoomchatCode()
	err := roomchatRepo.CreateRoomchat(&roomchatEntity)
	if err != nil {
		return roomchatEntity, err
	}
	return roomchatEntity, nil
}

func (r *roomchatService) JoinRoomchat(email string, roomId uint64) (entities.RoomchatUser, error) {
	//Call Repo
	roomchatRepo := NewRoomchatRepository()
	userRepo := users.NewUserRepository()
	var userEntity entities.User
	err := userRepo.GetUserByEmail(email, &userEntity)
	if err != nil {
		return entities.RoomchatUser{}, err
	}
	roomchatUser, err := roomchatRepo.JoinRoomchat(userEntity.ID, roomId)
	if err != nil {
		return roomchatUser, err
	}
	return roomchatUser, nil
}

func (r *roomchatService) GetRoomchatByUserId(userIds []int64) (entities.Roomchat, int, error) {
	//Call Repo
	roomchatRepo := NewRoomchatRepository()
	roomchat, httpStatus, err := roomchatRepo.GetRoomchatByUserId(userIds)
	if err != nil {
		return roomchat, httpStatus, err
	}
	return roomchat, httpStatus, nil
}
