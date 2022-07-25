package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	fetchlist "github.com/lennysgarage/letterboxd-scraper/fetchlist"
)

func writeList(link string, movies []fetchlist.Movie) {
	link = strings.Replace(link, "https://letterboxd.com/", "", 1) // remove url in name of file
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
		writer.Write([]string{movie.Title, movie.Link})
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("link to letterboxd watchlist or other public list required")
		os.Exit(1)
	}

	var wg sync.WaitGroup

	for _, link := range os.Args[1:] {
		wg.Add(1)

		go func(link string) {
			defer wg.Done()
			movies := fetchlist.FetchWatchlist(link)
			writeList(link, movies)
		}(link)
	}
	wg.Wait()
}
