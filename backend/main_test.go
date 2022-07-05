package main

import (
	"os"
	"testing"
)

func BenchmarkMain(b *testing.B) {

	// go test -bench=. -benchmem
	os.Args = append(os.Args, "https://letterboxd.com/lennysgarage/watchlist/")
	os.Args = append(os.Args, "https://letterboxd.com/jojosweetermanz/watchlist/")
	os.Args = append(os.Args, "https://letterboxd.com/kennychiwa/watchlist/")

	main()
}
