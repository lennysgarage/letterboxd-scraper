package fetchlist

import (
	"fmt"
	"strings"
	"log"

	"github.com/gocolly/colly"
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

	err := c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 4})
	if err != nil {
		log.Println("Failed to setup colly limit ", err)
	}

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	c.OnHTML(".pagination", func(e *colly.HTMLElement) {
		nextPage := e.ChildAttr(".paginate-nextprev a.next", "href")
		err = c.Visit(e.Request.AbsoluteURL(nextPage))
		if err != nil {
			log.Println("Failed to visit absoluteURL", err)
		}
	})

	// Find all movies in list
	c.OnHTML(".poster-list li", func(e *colly.HTMLElement) {
		film := e.ChildAttr("div", "data-target-link")

		movie := Movie{}
		movie.Title = strings.Replace(film[6:len(film)-1], "-", " ", -1)
		movie.Link = "https://letterboxd.com" + film

		movies = append(movies, movie)
	})

	err = c.Visit(link)
	if err != nil {
		log.Println("Failed to visit webpage", err)
	}

	c.Wait()
	return movies
}
