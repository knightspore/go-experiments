package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

type Lightswitch struct {
	Name string
	On   bool
}

type PageData struct {
	Lights []Lightswitch
}

func main() {

	// Setup Database

	db, err := sql.Open("sqlite3", "./lights.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Setup Server

	fs := http.FileServer(http.Dir("./www"))
	tmpl := template.Must(template.ParseFiles("./templates/lightswitch.html"))

	// Register Routes

	http.Handle("/", fs)

	http.HandleFunc("/lightswitches", func(w http.ResponseWriter, _ *http.Request) {
		var switches []Lightswitch
		rows, err := db.Query("select name, on_status from lights")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var name string
			var on_status bool
			err = rows.Scan(&name, &on_status)
			if err != nil {
				log.Fatal(err)
			}
			switches = append(switches, Lightswitch{name, on_status})
		}
		data := PageData{Lights: switches}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {

		// Read body

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("<p>Error: Network error.</p>"))
			log.Println(err)
		}
		r.Body.Close()

		// Get Name

		data, err := url.ParseQuery(string(body))
		if err != nil {
			w.Write([]byte("<p>Error: Could not read name.</p>"))
			log.Println(err)
		}

		name := data["name"][0]

		// Create new Light

		if len(name) == 0 {
			w.Write([]byte("<p>Error: Could not read name.</p>"))
		} else {
			_, err := db.Exec("INSERT INTO lights (name, on_status) VALUES (?,?);", name, false)
			if err != nil {
				w.Write([]byte("<p>Error: Error saving light.</p>"))
				log.Printf("Error adding %q\n", name)
				log.Println(err)
			} else {
				log.Printf("Added %q\n", name)
				w.Header().Set("HX-Trigger", "additem")
				w.Write([]byte(fmt.Sprintf("<p>Added %q</p>", name)))
			}
		}
	})

	http.HandleFunc("/toggle", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		on := r.URL.Query().Get("on")
		_, err := db.Exec("UPDATE lights SET on_status = ? where name = ?;", on, name)
		if err != nil {
			log.Printf("Error toggling %s\n", name)
			log.Println(err)
		} else {
			log.Printf("Toggling %q\n", name)
			w.Header().Set("HX-Trigger", "onoff")
		}
	})

	// Start Server
	log.Println("Listening on :3000...")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
