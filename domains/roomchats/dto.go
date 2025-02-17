package roomchats

type RoomchatRequest struct {
	RoomName   string `json:"room_name"`
	RoomchatId string `json:"roomchat_id"`
	CreatedBy  uint64 `json:"created_by"`
}

type StartRoomchatRequest struct {
	Email string `json:"email"`
}
