package routes

import (
	"context"
	"fmt"
	"time"

	"net/http"

	"tutorial/web-api/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var albumCollection = config.GetDBCollection(config.DB, "AlbumInfo")

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var albums []Album
		results, err := albumCollection.Find(ctx, bson.M{})

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Albums Not Found"})
		}

		defer results.Close(ctx)
		for results.TryNext(ctx) {
			var album Album
			if err = results.Decode(&album); err != nil {
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

		newAlbum := Album{
			ID:     "4",
			Title:  "Electric LadyLand",
			Artist: "Jimi Hendrix",
			Price:  13.99,
		}

		result, err := albumCollection.InsertOne(ctx, newAlbum)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error posting album"})
			return
		}

		c.IndentedJSON(http.StatusCreated, result)
	}
}

func GetAlbumById() gin.HandlerFunc {
	return func(c *gin.Context) {
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
}
