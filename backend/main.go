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

func writeList(link string, movies []fetchlist.Movie) error {
	link = strings.Replace(link, "https://letterboxd.com/", "", 1) // remove url in name of file

	file, err := os.Create(fmt.Sprintf("%s.csv", strings.Replace(link, "/", ":", -1)))
	if err != nil {
		return err
	}

	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"Title", "LetterboxdURI"}
	err = writer.Write(headers)
	if err != nil {
		return err
	}

	for _, movie := range movies {
		err = writer.Write([]string{movie.Title, movie.Link})
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("link to letterboxd watchlist or other public list required")
		os.Exit(1)
	}

	var wg sync.WaitGroup

	for _, link := range os.Args[1:] {
		wg.Add(1)

		go func(link string) {
			defer wg.Done()
			movies := fetchlist.FetchWatchlist(link)

			err := writeList(link, movies)
			if err != nil {
				log.Fatal("Error writing list, %w", err)
			}
		}(link)
	}
	wg.Wait()

	log.Println("Finished writing list!")
}
