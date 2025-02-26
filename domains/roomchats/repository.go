package roomchats

import (
	"chatroom-api/database"
	"chatroom-api/entities"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

var db = database.SetupDatabaseConnection()
var elasticClient = database.SetupElasticConnection()

type RoomchatRepository interface {
	CreateRoomchat(roomchat *entities.Roomchat) error
	JoinRoomchat(userId uint64, roomId uint64) (entities.RoomchatUser, error)
	GetRoomchatByUserId(userIds []int64) (entities.Roomchat, int, error)
	GetRoomchatUsers(userId int64) ([]entities.Roomchat, error)
	GetChatHistories(roomId string) ([]entities.ChatHistory, error)
}

type roomchatRepository struct{}

func NewRoomchatRepository() RoomchatRepository {
	return &roomchatRepository{}
}

func (r *roomchatRepository) CreateRoomchat(roomchat *entities.Roomchat) error {
	// Save to DB using gorm
	err := db.Create(&roomchat)
	if err.Error != nil {
		return err.Error
	}

	return nil
}

func (r *roomchatRepository) JoinRoomchat(userId uint64, roomId uint64) (entities.RoomchatUser, error) {
	// Save to DB using gorm
	roomchatUser := entities.RoomchatUser{
		RoomchatID: roomId,
		UserID:     userId,
		CreatedAt:  time.Now().UTC(), // Use UTC time
		UpdatedAt:  time.Now().UTC(), // Use UTC time
	}

	err := db.Create(&roomchatUser)
	if err.Error != nil {
		return roomchatUser, err.Error
	}

	return roomchatUser, nil
}

func (r *roomchatRepository) GetRoomchatByUserId(userIds []int64) (entities.Roomchat, int, error) {
	var roomchat entities.Roomchat
	var roomchatIDs []int64

	// Ambil roomchat_id yang memiliki user sesuai userIds
	err := db.Table("roomchat_users").
		Select("roomchat_id").
		Where("user_id IN (?)", userIds).
		Group("roomchat_id").
		Having("COUNT(user_id) = ?", len(userIds)).
		Pluck("roomchat_id", &roomchatIDs).Error

	if err != nil {
		return roomchat, http.StatusInternalServerError, err
	}

	if len(roomchatIDs) == 0 {
		return roomchat, http.StatusNotFound, nil // Tidak ada room yang sesuai
	}

	// Pastikan hanya mengambil room yang berisi tepat userIds yang diminta
	var validRoomchatID int64
	for _, roomID := range roomchatIDs {
		var userCount int64
		err := db.Table("roomchat_users").
			Where("roomchat_id = ?", roomID).
			Count(&userCount).Error

		//logging user count
		logrus.Info("User count: ", userCount)
		if err != nil {
			return roomchat, http.StatusInternalServerError, err
		}

		if userCount == int64(len(userIds)) {
			validRoomchatID = roomID
			break
		}
	}

	if validRoomchatID == 0 {
		return roomchat, http.StatusNotFound, nil // Tidak ada room yang sesuai
	}

	// Ambil data roomchat
	err = db.Where("id = ?", validRoomchatID).First(&roomchat).Error
	if err != nil {
		return roomchat, http.StatusInternalServerError, err
	}

	return roomchat, http.StatusOK, nil
}

func (r *roomchatRepository) GetRoomchatUsers(userId int64) ([]entities.Roomchat, error) {
	var roomchats []entities.Roomchat
	var roomchatIds []int64
	// Ambil roomchat_id yang memiliki user sesuai userIds
	err := db.Table("roomchat_users").
		Select("roomchat_id").
		Where("user_id = ?", userId).
		Pluck("roomchat_id", &roomchatIds).Error
	if err != nil {
		return roomchats, err
	}

	if len(roomchatIds) == 0 {
		return []entities.Roomchat{}, nil // Tidak ada room yang sesuai
	}

	// Ambil data roomchat berdasarkan roomchatIds
	err = db.Where("id IN (?)", roomchatIds).Find(&roomchats).Error
	if err != nil {
		return nil, err
	}

	return roomchats, nil
}

func (r *roomchatRepository) GetChatHistories(roomId string) ([]entities.ChatHistory, error) {

	chatHistoriesIndex := os.Getenv("ELASTIC_CHAT_HISTORIES_INDEX")
	// Search chat histories from elastic search
	chatHistories := []entities.ChatHistory{}
	ctx := context.Background()
	query := elastic.NewMatchQuery("room_id", roomId)
	searchResult, err := elasticClient.Search().
		Index(chatHistoriesIndex).
		Query(query).
		Sort("created_at", false).
		Do(ctx)
	if err != nil {
		return chatHistories, err
	}

	for _, hit := range searchResult.Hits.Hits {
		var chatHistory entities.ChatHistory
		err := json.Unmarshal(hit.Source, &chatHistory)
		if err != nil {
			return chatHistories, err
		}
		chatHistories = append(chatHistories, chatHistory)
	}

	return chatHistories, nil
}
