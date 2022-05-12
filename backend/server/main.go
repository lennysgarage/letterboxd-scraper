package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	fetchlist "github.com/lennysgarage/letterboxd-scraper/fetchlist"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/api", func(c *gin.Context) {
		link := c.Query("src")
		if movies := fetchlist.FetchWatchlist(link); len(movies) != 0 {
			c.JSON(http.StatusOK, movies)
		} else {
			c.Status(http.StatusNotFound)
		}
	})

	router.Run(":" + port)
}
