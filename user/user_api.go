package user

import (
	"strconv"
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
)
// UserApi Struct to Connect service with route
type UserApi struct {
	userService UserService
}

// ProvideUserAPI for provide user api on wire.go
func ProvideUserAPI(u UserService) UserApi {
	return UserApi{userService: u}
}
// FindAll to get all users
func (userApi *UserApi) FindAll(c *gin.Context) {
	users := userApi.userService.GetAll()
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// FindByID to get Single User
func (userApi *UserApi) FindByID(c *gin.Context) {
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		log.Fatalln(err)
	}
	user:=userApi.userService.GetByID(id)
	c.JSON(http.StatusOK, gin.H{"user": user})
}
