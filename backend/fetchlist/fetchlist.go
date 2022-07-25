package fetchlist

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type Movie struct {
	Title string
	Link  string
}

func FetchWatchlist(link string) []Movie {
	var movies []Movie
	// var movies [][]string
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

		// row := []string{movie.Title, movie.Link}
		movies = append(movies, movie)
	})

	c.Visit(link)

	c.Wait()
	return movies
}
