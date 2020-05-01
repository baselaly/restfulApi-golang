package main

import (
	"go_restful/user"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func initDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root@/gorest")
	if err != nil {
		log.Fatalln("err")
	}
	db.AutoMigrate(&user.User{})
	return db
}

func main() {
	db := initDB()
	defer db.Close()

	userAPI := InitUserApi(db)

	route := gin.Default()

	route.GET("/users", userAPI.FindAll)
	route.GET("/user/:id", userAPI.FindByID)
	route.POST("/user/create", userAPI.Create)

	err := route.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
