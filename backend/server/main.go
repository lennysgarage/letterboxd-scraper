package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	fetchlist "github.com/lennysgarage/letterboxd-scraper/fetchlist"
)

func intersectWatchlists(watchlist []fetchlist.Movie, numLinks int) []fetchlist.Movie {
	intersection := make([]fetchlist.Movie, 0)
	hash := make(map[fetchlist.Movie]int)

	for _, movie := range watchlist {
		hash[movie] += 1
	}

	for movie, count := range hash {
		if count == numLinks {
			intersection = append(intersection, movie)
		}
	}

	return intersection
}

func unionWatchlists(watchlist []fetchlist.Movie) []fetchlist.Movie {
	union := make([]fetchlist.Movie, 0)
	hash := make(map[fetchlist.Movie]bool)

	for _, movie := range watchlist {
		hash[movie] = true
	}

	for movie := range hash {
		union = append(union, movie)
	}

	return union
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	var wg sync.WaitGroup

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/api", func(c *gin.Context) {
		links := c.QueryArray("src")
		intersection := c.Query("i")

		var movieList []fetchlist.Movie
		// Fetch all user's watchlists
		for _, link := range links {
			wg.Add(1)

			go func(link string) {
				defer wg.Done()
				movies := fetchlist.FetchWatchlist(link)
				movieList = append(movieList, movies...)
			}(link)
		}
		wg.Wait()

		if len(links) > 1 && intersection == "true" {
			// Create intersected watchlist
			movieList = intersectWatchlists(movieList, len(links))
		} else { // Union		
			movieList = unionWatchlists(movieList)
		}

		if len(movieList) != 0 {
			c.JSON(http.StatusOK, movieList)
		} else {
			c.Status(http.StatusNotFound)
		}
	})

	err := router.Run(":" + port)
	if err != nil {
		log.Println("server crashed", err)
	}
}
