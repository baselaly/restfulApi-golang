package main

import (
	"go_restful/user"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func InitUserApi(db *gorm.DB) user.UserApi {
	wire.Build(user.ProvideUserRepository, user.ProvideUserService, user.ProvideUserApi)
	return user.UserApi{}
}
