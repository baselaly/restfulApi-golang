package user

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"not null;unique"`
	Password string `json:"password"`
}
