package entities

type RoomchatUser struct {
	ID         uint64 `gorm:"primary_key:auto_increment" json:"id"`
	RoomchatID uint64 `gorm:"not null" json:"roomchat_id"`
	UserID     uint64 `gorm:"not null" json:"user_id"`
	CreatedAt  string `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  string `gorm:"autoUpdateTime" json:"updated_at"`
}
