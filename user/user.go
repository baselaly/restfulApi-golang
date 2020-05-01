package user

import (
	"time"
)

// User model
type User struct {
	ID        int16     `json:"id" gorm:"primary_key"`
	Email     string    `json:"email" gorm:"not null;unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
