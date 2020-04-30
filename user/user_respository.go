package user

import (
	"github.com/jinzhu/gorm"
)

// UserRepository for user
type UserRepository struct {
	DB *gorm.DB
}

func ProvideUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

func (u *UserRepository) FindAll() []User {
	var users []User
	u.DB.Find(&users)
	return users
}

func (u *UserRepository) FindByID(id int) User {
	var user User
	u.DB.First(&user, id)
	return user
}

func (u *UserRepository) Create(user User) User {
	u.DB.Create(&user)
	return user
}

func (u *UserRepository) Delete(id int) bool {
	user := u.FindByID(id)
	u.DB.Delete(&user)
	return true
}
