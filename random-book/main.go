package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	books := GetBooks()

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(books), func(i, j int) { books[i], books[j] = books[j], books[i] })

	fmt.Printf("Option 1: %q\n", books[0])
	fmt.Printf("Option 2: %q\n", books[1])
	fmt.Printf("Option 3: %q\n", books[2])
}

func GetBooks() []string {

	var books []string

	readFile, err := os.Open("/home/c/Documents/Planner/📔Reading List.md")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		if strings.HasPrefix(text, "- [ ]") {
			book := strings.Trim(strings.Split(text, "] ")[1], " ")
			books = append(books, book)
		}
	}

	readFile.Close()

	return books
}
