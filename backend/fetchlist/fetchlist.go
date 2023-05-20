package fetchlist

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

type Movie struct {
	Title string
	Link  string
}

func FetchWatchlist(link string) []Movie {
	var movies []Movie
	c := colly.NewCollector(
		colly.AllowedDomains("letterboxd.com"),
		colly.Async(true),
	)

	c.OnError(func(e *colly.Response, err error) {
		fmt.Println("Something went wrong: ", err)
	})

	// Determines if a link to a list or a username.
	link = formatInput(link)

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 4})

	extensions.RandomUserAgent(c)
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting: ", r.URL.String())
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

		movies = append(movies, movie)
	})

	c.Visit(link)

	c.Wait()
	return movies
}

func formatInput(s string) string {
	if strings.HasPrefix(s, "http") {
		return s
	}

	return fmt.Sprintf("https://letterboxd.com/%s/watchlist/page/1/", s)
}
