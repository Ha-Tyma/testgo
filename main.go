package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Harvest", Artist: "Neil Young", Price: 20.50},
	{ID: "2", Title: "In Rainbows", Artist: "Radiohead", Price: 22.50},
	{ID: "3", Title: "Mutter", Artist: "Rammstein", Price: 23.50},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run(":8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
