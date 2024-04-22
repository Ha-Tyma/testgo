package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Artist struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist Artist  `json:"artist"`
	Price  float64 `json:"price"`
}

var artists = []Artist{
	{ID: "1", Name: "Rammstein", Country: "DE"},
	{ID: "2", Name: "Neil Young", Country: "US"},
	{ID: "3", Name: "Radiohead", Country: "UK"},
}

var albums = []Album{
	{ID: "1", Title: "Harvest", Artist: artists[1], Price: 20.50},
	{ID: "2", Title: "In Rainbows", Artist: artists[2], Price: 22.50},
	{ID: "3", Title: "Mutter", Artist: artists[0], Price: 23.50},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbum)
	router.GET("/artists", getArtists)
	router.GET("/artists/:id", getArtist)

	router.Run(":8080")
}

func getArtist(c *gin.Context) {
	id := c.Param("id")
	for _, a := range artists {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
		}
	}
}

func getArtists(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, artists)
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbum(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
		}
	}
}
