package tests

import (
	"chatroom-api/database"
	"chatroom-api/domains/roomchats"
	"io/ioutil"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gopkg.in/stretchr/testify.v1/require"
)

func TestCreateRoomchat(t *testing.T) {
	// Suppress logrus output by setting the log level to PanicLevel
	logrus.SetOutput(ioutil.Discard)
	// Load environment variables
	err := godotenv.Load()
	require.NoError(t, err, "Failed to load environment variables")

	// Set up database connection
	db := database.SetupDatabaseConnection()
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// Run GORM AutoMigrate
	// err = db.AutoMigrate(
	// 	&entities.User{},
	// 	&entities.Roomchat{},
	// 	&entities.RoomchatUser{},
	// )

	require.NoError(t, err, "Failed to run migrations")
	// Create a RoomchatRequest
	var roomchatRequest roomchats.RoomchatRequest
	roomchatRequest.RoomName = "Room 1"
	roomchatRequest.CreatedBy = 1

	// Create RoomchatService
	roomchatService := roomchats.NewRoomchatService()

	// Call CreateRoomchat
	createdRoomchat, err := roomchatService.CreateRoomchat(roomchatRequest)
	require.NoError(t, err, "Failed to create roomchat")
	require.NotEmpty(t, createdRoomchat, "Created roomchat should not be empty")
	require.NotEqual(t, createdRoomchat.ID, 0, "Created roomchat ID should not be 0")
	require.Equal(t, roomchatRequest.CreatedBy, createdRoomchat.CreatedBy, "CreatedBy should match the request")

	// Call JoinRoomchat
	logrus.Info("Joining roomchat" + strconv.FormatUint(createdRoomchat.ID, 10))
	roomchatUser, err := roomchatService.JoinRoomchat("andikayuda628@gmail.com", createdRoomchat.ID)
	require.NoError(t, err, "Failed to join roomchat")
	require.NotEmpty(t, roomchatUser, "Joined roomchat user should not be empty")
	require.NotEqual(t, roomchatUser.ID, 0, "RoomchatUser ID should not be 0")
	require.Equal(t, roomchatUser.RoomchatID, createdRoomchat.ID, "RoomchatUser.RoomchatID should match the created roomchat ID")
}

func TestJoinRoomchat(t *testing.T) {
	// Suppress logrus output by setting the log level to PanicLevel
	logrus.SetLevel(logrus.PanicLevel)
	// Load environment variables
	err := godotenv.Load()
	require.NoError(t, err, "Failed to load environment variables")

	// Set up database connection
	db := database.SetupDatabaseConnection()
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// Create RoomchatService
	roomchatService := roomchats.NewRoomchatService()

	// Call JoinRoomchat
	roomchatUser, err := roomchatService.JoinRoomchat(
		"andikayuda628@gmail.com",
		2,
	)
	require.NoError(t, err, "Failed to join roomchat")
	require.NotEmpty(t, roomchatUser, "Joined roomchat user should not be empty")
	require.NotEqual(t, roomchatUser.ID, 0, "RoomchatUser ID should not be 0")
	require.EqualValues(t, roomchatUser.RoomchatID, 2, "RoomchatUser.RoomchatID should match the created roomchat ID")
}
