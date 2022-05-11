package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type Movie struct {
	Title string
	Link  string
}

func fetchWatchlist(link string) [][]string {
	var movies [][]string
	c := colly.NewCollector(
		colly.AllowedDomains("letterboxd.com"),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 4})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	c.OnHTML(".pagination", func(e *colly.HTMLElement) {
		nextPage := e.ChildAttr(".paginate-nextprev a.next", "href")
		c.Visit(e.Request.AbsoluteURL(nextPage))
	})

	// Find all movies in list
	c.OnHTML(".poster-list li", func(e *colly.HTMLElement) {
		film := e.ChildAttr("div", "data-target-link")

		movie := Movie{}
		movie.Title = strings.Replace(film[6:len(film)-1], "-", " ", -1)
		movie.Link = "https://letterboxd.com" + film

		row := []string{movie.Title, movie.Link}
		movies = append(movies, row)
	})

	c.Visit(link)

	c.Wait()
	return movies
}

func writeWatchlist(link string, movies [][]string) {
	link = strings.Replace(link, "https://letterboxd.com/", "", 1)
	file, err := os.Create(fmt.Sprintf("%s.csv", strings.Replace(link, "/", ":", -1)))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	headers := []string{"Title", "LetterboxdURI"}
	writer.Write(headers)

	for _, movie := range movies {
		writer.Write(movie)
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("link to letterboxd watchlist or other public list required")
		os.Exit(1)
	}

	link := os.Args[1]
	movies := fetchWatchlist(link)
	writeWatchlist(link, movies)

}
