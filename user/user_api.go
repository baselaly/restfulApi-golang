package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	userService UserService
}

func ProvideUserApi(u UserService) UserApi {
	return UserApi{userService: u}
}

func (userApi *UserApi) FindAll(c *gin.Context) {
	users := userApi.userService.GetAll()
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (userApi *UserApi) FindById(c *gin.Context) string {
	return c.Param("id")
}
