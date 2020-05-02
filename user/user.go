package user

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"not null;unique"`
	Password string `json:"password"`
}

// TransformedUser format user for response
type TransformedUser struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}
