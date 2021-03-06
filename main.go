package main

import (
	"github.com/danielmoledo/hello-world/controllers"
	"github.com/danielmoledo/hello-world/model"
	"github.com/danielmoledo/hello-world/repository"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitializeDB() *gorm.DB{
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&model.HelloWorld{})

	return database
}

func main() {
	repository.DB = InitializeDB()
	r := gin.Default()
	r.POST("/hello-world", controllers.HelloWorld)
	r.Run()
}
