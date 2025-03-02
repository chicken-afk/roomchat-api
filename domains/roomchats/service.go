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
	GetRoomchatUsers(userId int64) ([]RoomlistResponse, error)
	GetChatHistories(roomId string) ([]entities.ChatHistory, error)
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

func (r *roomchatService) GetRoomchatUsers(userId int64) ([]RoomlistResponse, error) {
	//Call Repo
	var roomlist []RoomlistResponse
	roomchatRepo := NewRoomchatRepository()
	roomchat, err := roomchatRepo.GetRoomchatUsers(userId)
	//Mapping roomchat to roomlist
	for _, room := range roomchat {
		var userlist []UserResponse
		for _, user := range room.Users {
			userlist = append(userlist, UserResponse{
				ID:    user.ID,
				Email: user.Email,
			})
		}
		roomlist = append(roomlist, RoomlistResponse{
			RoomId:    room.RoomId,
			RoomName:  room.RoomName,
			CreatedAt: room.CreatedAt.Format("2006-01-02 15:04:05"),
			Users:     userlist,
		})
	}
	if err != nil {
		return roomlist, err
	}
	return roomlist, nil
}

func (r *roomchatService) GetChatHistories(roomId string) ([]entities.ChatHistory, error) {
	//Call Repo
	roomchatRepo := NewRoomchatRepository()
	chatHistories, err := roomchatRepo.GetChatHistories(roomId)
	if err != nil {
		return chatHistories, err
	}
	return chatHistories, nil
}
