package roomchats

import (
	"chatroom-api/database"
	"chatroom-api/entities"
)

var db = database.SetupDatabaseConnection()

type RoomchatRepository interface {
	CreateRoomchat(roomchat entities.Roomchat) error
	JoinRoomchat(userId uint64, roomId uint64) error
}

type roomchatRepository struct{}

func NewRoomchatRepository() RoomchatRepository {
	return &roomchatRepository{}
}

func (r *roomchatRepository) CreateRoomchat(roomchat entities.Roomchat) error {
	// Save to DB using gorm
	err := db.Create(&roomchat)
	if err.Error != nil {
		return err.Error
	}

	return nil
}

func (r *roomchatRepository) JoinRoomchat(userId uint64, roomId uint64) error {
	// Save to DB using gorm
	roomchatUser := entities.RoomchatUser{
		RoomchatID: roomId,
		UserID:     userId,
	}

	err := db.Create(&roomchatUser)
	if err.Error != nil {
		return err.Error
	}

	return nil
}
