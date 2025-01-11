package entities

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"` // Primary key
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`  // Unique and not null
	Password  string    `gorm:"not null" json:"password"`           // Not null
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`   // Auto-generated timestamp
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`   // Auto-generated timestamp
}
