package routes

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"net/http"

	"tutorial/web-api/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var albumCollection = config.GetDBCollection(config.DB, "AlbumInfo")

var albums = []Album{
	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var albums []Album
		results, err := albumCollection.Find(ctx, bson.M{})

		if err != nil {
			fmt.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Albums Not Found"})
			return
		}

		defer results.Close(ctx)
		for results.TryNext(ctx) {
			var album Album
			if err = results.Decode(&album); err != nil {
				fmt.Println(err)
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving album"})
			}

			albums = append(albums, album)
		}

		c.IndentedJSON(http.StatusOK, albums)
	}
}

func PostAlbums() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		raw, err := c.GetRawData()
		fmt.Println(string(raw))
		var newAlbum Album
		err = c.ShouldBindJSON(&newAlbum)

		if err != nil {
			msg := "Error finding title or artist from post request"
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": msg})
			return
		}

		result, err := albumCollection.InsertOne(ctx, newAlbum)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error posting album"})
			return
		}

		msg := fmt.Sprint("Inserted album with new id: %v", result.InsertedID)
		c.IndentedJSON(http.StatusCreated, gin.H{"message": msg})
	}
}

func GetAlbumById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": fmt.Errorf("Error parsing int id: %v", err)})
		}

		result := albumCollection.FindOne(ctx, bson.M{"ID": id})
		var album Album

		err = result.Err()
		if err != nil {
			msg := fmt.Errorf("Error finding album with id %v", id)
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": msg})
			return
		}
		err = result.Decode(&album)
		if err != nil {
			msg := fmt.Errorf("Error decoding album with id %v", id)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": msg})
			return
		}

		c.IndentedJSON(http.StatusOK, album)
	}
}
