package roomchats

type RoomchatRequest struct {
	RoomName   string `json:"room_name"`
	RoomchatId string `json:"roomchat_id"`
	CreatedBy  uint64 `json:"created_by"`
}

type StartRoomchatRequest struct {
	Email string `json:"email"`
}

type SendMessageRequest struct {
	RoomchatId string `json:"roomchat_id"`
	Message    string `json:"message"`
}

type RoomlistResponse struct {
	RoomId    string         `json:"room_id"`
	RoomName  string         `json:"room_name"`
	CreatedAt string         `json:"created_at"`
	Users     []UserResponse `json:"users"`
}

type UserResponse struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
}
