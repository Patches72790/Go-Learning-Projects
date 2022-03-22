package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func WriteAlbumsToDisk(albums []Album) gin.HandlerFunc {

	fmt.Println("writing albums to disk")
	return func(c *gin.Context) {
		var albums, _ = c.Get("albums")
		if albums == nil {
			return
		}

		var cast_albums = albums.([]Album)

		j, _ := json.Marshal(cast_albums)

		error := os.WriteFile("./albums2.json", j, 777)
		if error != nil {
			log.Fatal(error)
		}

		fmt.Println("writing albums to disk")
		c.Next()
	}
}

func ReadAlbumsFromDisk() []Album {
	jsonFile, err := os.Open("./albums.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	bytes, _ := ioutil.ReadAll(jsonFile)

	var result []Album

	json.Unmarshal([]byte(bytes), &result)

	return result
}
