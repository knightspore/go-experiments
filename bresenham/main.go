package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

func main() {

	data := ReadObjFile("./models/rock.obj")

	var points []Point

	for _, val := range data {
		points = append(points, val.ToPoint())
	}

	render(func() {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		poly := NewPolygon(points...)
		poly.Print(0)
		time.Sleep(5 * time.Second)
	})
}

func render(f func()) {

	err := termbox.Init()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

	f()

}
