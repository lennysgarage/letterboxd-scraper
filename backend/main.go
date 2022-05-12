package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	
	fetchlist "github.com/lennysgarage/letterboxd-scraper/fetchlist"
)

func writeList(link string, movies [][]string) {
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
		writer.Write(movie)
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("link to letterboxd watchlist or other public list required")
		os.Exit(1)
	}

	link := os.Args[1]
	movies := fetchlist.FetchWatchlist(link)
	writeList(link, movies)

}
