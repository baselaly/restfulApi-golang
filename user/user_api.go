package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}
	user := userApi.userService.GetByID(id)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Create to create new user
func (userApi *UserApi) Create(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	hashBytes, _ := bcrypt.GenerateFromPassword([]byte(password), 15)
	user := User{Email: email, Password: string(hashBytes)}
	newUser, err := userApi.userService.Create(user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": newUser})
}
