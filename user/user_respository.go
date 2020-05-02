package user

import (
	"errors"

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

func (u *UserRepository) FindByID(id int) (User, error) {
	var user User
	record := u.DB.First(&user, id).RecordNotFound()
	if record {
		return User{}, errors.New("no records found")
	}
	return user, nil
}

func (u *UserRepository) Create(user User) (User, error) {
	err := u.DB.Create(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (u *UserRepository) Delete(id int) bool {
	user, _ := u.FindByID(id)
	u.DB.Delete(&user)
	return true
}
