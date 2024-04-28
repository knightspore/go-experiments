package main

import (
	"testing"

	"github.com/gocolly/colly/v2"
)

func TestBook(t *testing.T) {
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
}

func TestLibgenClient(t *testing.T) {

	libgen := NewLibgenClient()

	t.Run("client can create search hash", func(t *testing.T) {
		hash := libgen.SearchHash("test")
		libgen.Results[hash] = []Book{}
		printed := libgen.FormatResults(hash)
		if printed != "[]" {
			t.Errorf("got %v want %v", printed, "[]")
		}
	})

	t.Run("client can create search url", func(t *testing.T) {
		url := libgen.CreateURL("test")
		want := "https://libgen.is/search.php?req=test&open=0&res=10&view=simple&phrase=1&column=def"
		if url != want {
			t.Errorf("got %v want %v", url, want)
		}
	})

	t.Run("client can parse row", func(t *testing.T) {
		el := &colly.HTMLElement{
			Text: `<tr valign="top" bgcolor="#C6DEFF"><td>643</td>
				<td><a href="search.php?req=Larry J. Crockett&amp;column[]=author">Larry J. Crockett</a></td>
				<td width="500"><a href="search.php?req=Ablex+Series+in+Artificial+Intelligence&amp;column=series"><font face="Times" color="green"><i>Ablex Series in Artificial Intelligence</i></font></a><br><a href="book/index.php?md5=2F2DBA2A621B693BB95601C16ED680F8" title="" id="643">The Turing Test and the Frame Problem: AI's Mistaken Understanding of Intelligence<br> <font face="Times" color="green"><i>9780893919269, 0893919268</i></font></a></td>
				<td>Ablex Publishing Corporation</td>
				<td nowrap="">1994</td>
				<td>216</td>
				<td>English</td>
				<td nowrap="">517 Kb</td>
				<td nowrap="">gz</td>
				<td><a href="http://library.lol/main/2F2DBA2A621B693BB95601C16ED680F8" title="Libgen &amp; IPFS &amp; Tor">[1]</a></td><td><a href="http://libgen.li/ads.php?md5=2F2DBA2A621B693BB95601C16ED680F8" title="Libgen.li">[2]</a></td>
				<td><a href="https://library.bz/main/edit/2F2DBA2A621B693BB95601C16ED680F8" title="Libgen Librarian">[edit]</a></td>
				</tr>`,
		}

		row := libgen.ParseRow(el)

		want := Book{
			ID:        643,
			Author:    "Larry J. Crockett",
			Title:     "The Turing Test and the Frame Problem: AI's Mistaken Understanding of Intelligence",
			Publisher: "Ablex Publishing Corporation",
			Year:      1994,
			Pages:     216,
			Language:  "English",
			Size:      "517 Kb",
			Extension: "gz",
			Mirrors:   []string{"http://library.lol/main/2F2DBA2A621B693BB95601C16ED680F8", "http://libgen.li/ads.php?md5=2F2DBA2A621B693BB95601C16ED680F8", "https://library.bz/main/edit/2F2DBA2A621B693BB95601C16ED680F8"},
		}

		if row.ID != want.ID {
			t.Errorf("got %v want %v", row.ID, want.ID)
		}
		if row.Author != want.Author {
			t.Errorf("got %v want %v", row.Author, want.Author)
		}
	})

	t.Run("live search is working", func(t *testing.T) {
		want := Book{
			ID: 643,
		}

		libgen.Search("test")

		got := libgen.Results[libgen.SearchHash("test")][0].ID

		if got != want.ID {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
