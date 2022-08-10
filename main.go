package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID 		string 	`json:id`
	NAME 	string 	`json:name`
	ARTIST 	string 	`json:artist`
	PRICE 	float64 `json:price`
}

var albums = []album{
	{ID:"1", NAME:"album1", ARTIST:"artist1", PRICE:211},
	{ID:"2", NAME:"album2", ARTIST:"artist2", PRICE:212},
	{ID:"3", NAME:"album3", ARTIST:"artist2", PRICE:213},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", addAlbum)

	router.Run("localhost:8000")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func addAlbum(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	for _, album := range albums {
		if album.ID == newAlbum.ID {
			c.IndentedJSON(http.StatusConflict, gin.H{"message": "album with same id already exists"})
			return
		}
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, album  := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}