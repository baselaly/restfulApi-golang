package user

import (
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
