package user

import (
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	Name     string
	Age      int
	Email    string
	Password string
}
