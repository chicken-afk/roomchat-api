package entities

import "time"

type Roomchat struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"` // Primary key
	RoomName  string    `gorm:"not null" json:"room_name"`          // Unique and not null
	RoomId    string    `gorm:"not null" json:"room_id"`
	CreatedBy uint64    `gorm:"null" json:"created_by"`           // Not null
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"` // Auto-generated timestamp
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"` // Auto-generated timestamp
	Users     []User    `gorm:"many2many:roomchat_users;foreignKey:ID;joinForeignKey:RoomchatID;References:ID;joinReferences:UserID" json:"users"`
}
