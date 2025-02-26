package entities

type ChatHistory struct {
	RoomID    string `json:"room_id"`
	Email     string `json:"email"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}
