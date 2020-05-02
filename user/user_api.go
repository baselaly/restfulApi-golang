package user

import (
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
	var _users []TransformedUser

	for _, user := range users {
		_users = append(_users, TransformedUser{ID: user.ID, Email: user.Email})
	}
	c.JSON(http.StatusOK, gin.H{"users": _users})
}

// FindByID to get Single User
func (userApi *UserApi) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	user, err := userApi.userService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "no record found"})
		return
	}
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

// Delete to delete user by id
func (userApi *UserApi) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	state, err := userApi.userService.Delete(id)

	if err != nil && state == false {
		c.JSON(http.StatusNotFound, gin.H{"message": "cant found record to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Deleted Successfully"})
}
