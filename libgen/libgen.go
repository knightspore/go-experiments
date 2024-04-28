package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

const (
	SELECTOR         = "body > table.c > tbody > tr"
	BASE_URL         = "https://libgen.is/search.php"
	FSTR_BASE_PARAMS = "?req=%s&open=0&res=10&view=simple&phrase=1&column=def"
)

type LibgenClient struct {
	Results map[string][]Book
}

type Book struct {
	ID        int      `json:"id"`
	Author    string   `json:"author"`
	Title     string   `json:"title"`
	Publisher string   `json:"publisher"`
	Year      int      `json:"year"`
	Pages     int      `json:"pages"`
	Language  string   `json:"language"`
	Size      string   `json:"size"`
	Extension string   `json:"extension"`
	Mirrors   []string `json:"mirrors"`
}

func (b Book) ToString() string {
	j, err := json.Marshal(b)
	if err != nil {
		return ""
	}
	return string(j)
}

func NewLibgenClient() *LibgenClient {
	return &LibgenClient{
		Results: make(map[string][]Book),
	}
}

func (l *LibgenClient) SearchHash(query string) string {
	hash := sha256.New()
	hash.Write([]byte(query))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (l *LibgenClient) CreateURL(query string) string {
	formattedQuery := strings.ReplaceAll(query, " ", "%20")
	return BASE_URL + fmt.Sprintf(FSTR_BASE_PARAMS, formattedQuery)
}

func (l *LibgenClient) ParseRow(e *colly.HTMLElement) Book {
	var values = make([]string, 9)
	var mirrors []string

	e.ForEach("td", func(i int, e *colly.HTMLElement) {
		if i < 9 {
			values[i] = e.Text
		}
		if strings.HasPrefix(e.Text, "[") && (e.Text != "[edit]") {
			mirrors = append(mirrors, e.ChildAttr("a", "href"))
		}
	})

	id, _ := strconv.Atoi(values[0])
	year, _ := strconv.Atoi(values[4])
	pages, _ := strconv.Atoi(values[5])

	return Book{
		ID:        id,
		Author:    values[1],
		Title:     values[2],
		Publisher: values[3],
		Year:      year,
		Pages:     pages,
		Language:  values[6],
		Size:      values[7],
		Extension: values[8],
		Mirrors:   mirrors,
	}

}

func (l *LibgenClient) Scrape(url string, hash string) {
	c := colly.NewCollector()
	c.OnHTML(SELECTOR, func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, "[") {
			book := l.ParseRow(e)
			l.Results[hash] = append(l.Results[hash], book)
		}
	})
	c.Visit(url)
}

func (l *LibgenClient) FormatResults(hash string) string {
	books := ""
	for _, book := range l.Results[hash] {
		fmt.Println(book.ToString())
		books += book.ToString() + ","
	}
	if len(books) > 1 {
		books = books[:len(books)-1] // remove trailing comma
	}
	return fmt.Sprintf("[%s]", books)
}

func (l *LibgenClient) Search(query string) {
	searchHash := l.SearchHash(query)
	l.Results[searchHash] = []Book{}
	url := l.CreateURL(query)
	l.Scrape(url, searchHash)
}
