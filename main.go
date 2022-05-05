package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"github.com/gocolly/colly"
)

type Movie struct {
	Title string
	Link  string
}

func fetchWatchlist(username string) [][]string {
	var movies [][]string
	c := colly.NewCollector(
		colly.AllowedDomains("letterboxd.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	c.OnHTML(".pagination", func(e *colly.HTMLElement) {
		nextPage := e.ChildAttr(".paginate-nextprev a.next", "href")
		c.Visit(e.Request.AbsoluteURL(nextPage))
	})

	// Find all movies in watchlist
	c.OnHTML(".poster-list li", func(e *colly.HTMLElement) {
		film := e.ChildAttr("div", "data-target-link")

		movie := Movie{}
		movie.Title = film[6 : len(film)-1]
		movie.Link = "https://letterboxd.com" + film

		row := []string{movie.Title, movie.Link}
		movies = append(movies, row)
	})

	c.Visit(fmt.Sprintf("https://letterboxd.com/%s/watchlist/page/1/", username))

	return movies
}

func writeWatchlist(username string, movies [][]string) {
	file, err := os.Create(fmt.Sprintf("watchlist-%s.csv", username))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	headers := []string{"Title", "Link"}
	writer.Write(headers)

	for _, movie := range movies {
		writer.Write(movie)
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Username required")
		os.Exit(1)
	}
	username := os.Args[1]

	movies := fetchWatchlist(username)
	writeWatchlist(username, movies)
}
