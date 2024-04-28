package main

import (
	"strings"
	"testing"
)

func TestBook(t *testing.T) {
	t.Run("book can be converted to json", func(t *testing.T) {
		b := Book{
			ID:        1,
			Author:    "test",
			Title:     "test",
			Publisher: "test",
			Year:      2000,
			Pages:     100,
			Language:  "test",
			Size:      "test",
			Extension: "test",
			Mirrors:   []string{"test"},
		}
		json := b.ToString()
		if json != "{\"id\":1,\"author\":\"test\",\"title\":\"test\",\"publisher\":\"test\",\"year\":2000,\"pages\":100,\"language\":\"test\",\"size\":\"test\",\"extension\":\"test\",\"mirrors\":[\"test\"]}" {
			t.Errorf("got %v want %v", json, "{\"id\":1,\"author\":\"test\",\"title\":\"test\",\"publisher\":\"test\",\"year\":2000,\"pages\":100,\"language\":\"test\",\"size\":\"test\",\"extension\":\"test\",\"mirrors\":[\"test\"]}")
		}
	})
}

func TestLibgenClient(t *testing.T) {

	libgen := NewLibgenClient()

	var hash string

	t.Run("client can create search hash", func(t *testing.T) {
		hash = libgen.SearchHash("test")
		libgen.Results[hash] = []Book{}
		printed := libgen.FormatResults(hash)
		if printed != "[]" {
			t.Errorf("got %v want %v", printed, "[]")
		}
	})

	var url string

	t.Run("client can create search url", func(t *testing.T) {
		url = libgen.CreateURL("test")
		want := "https://libgen.is/search.php?req=test&open=0&res=10&view=simple&phrase=1&column=def"
		if url != want {
			t.Errorf("got %v want %v", url, want)
		}
	})

	t.Run("client can scrape live site", func(t *testing.T) {
		libgen.Scrape(url, hash)

		got := libgen.Results[libgen.SearchHash("test")][0].ID
		want := Book{ID: 643}

		if got != want.ID {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("client can format results as json", func(t *testing.T) {
		printed := libgen.FormatResults(hash)
		if strings.Contains(printed, "643") == false {
			t.Errorf("got %v want %v", printed, "id")
		}

		start := printed[:2]
		end := printed[len(printed)-2:]

		if start != "[{" || end != "}]" {
			t.Errorf("got %v want %v", start, "[{")
			t.Errorf("got %v want %v", end, "}]")
		}
	})

}
