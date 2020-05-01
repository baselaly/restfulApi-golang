//+build !wireinject
// Code generated by Wire. DO NOT EDIT.

//go:generate wire

package main

import (
	"go_restful/user"

	"github.com/jinzhu/gorm"
)

// Injectors from wire.go:

func InitUserApi(db *gorm.DB) user.UserApi {
	userRepository := user.ProvideUserRepository(db)
	userService := user.ProvideUserService(userRepository)
	userApi := user.ProvideUserAPI(userService)
	return userApi
}
