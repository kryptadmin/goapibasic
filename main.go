package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record
type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

// albums slice to seed record albums
var albums = []album{
	{ID: "1", Title: "Waltair veeraiah", Artist: "dsp", Price: 10.00},
	{ID: "2", Title: "Pushpaa", Artist: "dsp", Price: 18.00},
	{ID: "3", Title: "Jalse", Artist: "dsp", Price: 20.00},
}

// getAlbums responds with list of albums in JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbum adds an album to the albums with the json received
func postAlbum(c *gin.Context) {
	var newAlbum album

	//call BindJSON to attach received json to the newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns the album as a response
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbum)

	router.Run("localhost:8000")
}