package tests

import (
	"chatroom-api/domains/roomchats"
	"testing"

	"github.com/joho/godotenv"
	"gopkg.in/stretchr/testify.v1/require"
)

func TestCreateRoomchat(t *testing.T) {
	godotenv.Load()

	var roomchatRequest roomchats.RoomchatRequest
	roomchatRequest.RoomName = "Room 1"
	roomchatRequest.CreatedBy = 1
	//Call Service
	roomchatService := roomchats.NewRoomchatService()
	createdRoomchat, err := roomchatService.CreateRoomchat(roomchatRequest)
	if err != nil {
		t.Errorf("Error while creating roomchat: %v", err)
	}
	require.NotEmpty(t, createdRoomchat)
	require.EqualValues(t, roomchatRequest.RoomName, createdRoomchat.RoomName)
}
