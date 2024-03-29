package main

import (
	"gin_blogs/controllers"
	"gin_blogs/models"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := gin.Default()
	r.Use(gin.Logger())

	r.LoadHTMLGlob("templates/**/*")

	models.ConnectDatabase()
	models.DBMigrate()

	r.GET("/blogs", controllers.BlogsIndex)
	r.GET("/blogs/:id", controllers.BlogsShow)

	log.Println("Server started!")
	r.Run() // Default Port 8080
}
