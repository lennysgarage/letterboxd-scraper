package main

import (
	"os"
	"testing"
)

func BenchmarkMain(b *testing.B) {

	// go test -bench=. -benchmem
	os.Args = append(os.Args, "https://letterboxd.com/lennysgarage/watchlist/")
	os.Args = append(os.Args, "jojosweetermanz")
	os.Args = append(os.Args, "https://letterboxd.com/kennychiwa/watchlist/")

	main()

	// Remove benchmark files.
	os.Remove("-test.bench=..csv")
	os.Remove("-test.benchmem=true.csv")
	os.Remove("-test.paniconexit0.csv")
	os.Remove("-test.timeout=10m0s.csv")
}
