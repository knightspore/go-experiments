package main

import (
	"fmt"
	"os"
	"strings"
)

func parseArgs() string {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Please use a search term: libgen-search <search term>")
		os.Exit(0)
	}
	return strings.Join(args, " ")
}

func main() {
	query := parseArgs()
	client := NewLibgenClient()
	client.Search(query)
	fmt.Println(client.FormatResults(client.SearchHash(query)))
}
