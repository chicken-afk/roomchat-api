package roomchats

type RoomchatRequest struct {
	RoomName  string `json:"room_name"`
	CreatedBy uint64 `json:"created_by"`
}
