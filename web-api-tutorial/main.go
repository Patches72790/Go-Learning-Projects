package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	//"os"
	"reflect"
	"time"
	"tutorial/web-api/routes"

	"context"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var albums = []routes.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func connect() (context.Context, *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//DB_URI := fmt.Sprintf("mongodb://%s:%s@%s:%s/?maxPoolSize=20&w=majority",
	//	os.Getenv("MONGO_DB_USERNAME"),
	//	os.Getenv("MONGO_DB_PASS"),
	//	os.Getenv("MONGO_HOST"),
	//	os.Getenv("MONGO_PORT"),
	//)

	//DB_URI := "mongodb://127.0.0.1:27017"

	DB_URI := fmt.Sprintf("%s", os.Getenv("MONGO_ATLAS_URI"))

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(DB_URI))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
		log.Fatal("Error pinging db client")
	}

	fmt.Println("Successfully connected and pinged client")

	return ctx, client
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	connect()

	albums := routes.ReadAlbumsFromDisk()
	router.Use(routes.WriteAlbumsToDisk(albums))

	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	albums, albums_exist := c.Get("albums")
	if !albums_exist {
		return
	}
	c.IndentedJSON(http.StatusOK, albums)

	c.Next()
}

func postAlbums(c *gin.Context) {
	var albums, _ = c.Get("albums")

	if albums == nil {
		return
	}

	fmt.Println("albums", albums)
	fmt.Println(reflect.TypeOf(albums))
	var newAlbum routes.Album

	var cast_albums = albums.([]routes.Album)
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	var new_albums = append(cast_albums, newAlbum)
	c.Set("album", new_albums)
	c.IndentedJSON(http.StatusCreated, newAlbum)

	c.Next()
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	var errorMessage = fmt.Sprintf("album not found with id %s", id)
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": errorMessage})
}
