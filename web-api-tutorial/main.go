package main

import (
	"log"

	"tutorial/web-api/config"
	"tutorial/web-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()

	router := gin.Default()

	router.GET("/albums", routes.GetAlbums())
	router.POST("/albums", routes.PostAlbums())
	router.GET("/albums/:id", routes.GetAlbumById())
	router.Run("localhost:8080")
}
