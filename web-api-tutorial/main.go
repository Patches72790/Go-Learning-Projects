package main

import (
	"fmt"
	"net/http"
	"reflect"
	"tutorial/web-api/routes"

	"github.com/gin-gonic/gin"
)

var albums = []routes.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()

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
