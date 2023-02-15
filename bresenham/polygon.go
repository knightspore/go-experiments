package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nsf/termbox-go"
)

type Polygon struct {
	Vertices []Point
	Edges    []Line
}

// NewPolygon returns a polygon with the number of points provided.
func NewPolygon(verts ...Point) Polygon {
	p := Polygon{
		Vertices: verts,
	}
	p.CalculateEdges()
	return p
}

// NewPolygonFromObj returns a Polygon loaded from a .obj file
func NewPolygonFromObj(p string) Polygon {
	points := LoadObj(p)
	polygon := NewPolygon(points...)
	return polygon
}

// OldCalculateEdges computeseges for a polygon with populated vertices.
// It doesn't work so well - it simply runs through a slice of Points
// and draws lines between them, linking the last again to the first.
func (p *Polygon) OldCalculateEdges() {
	close := NewLine(p.Vertices[len(p.Vertices)-1].X, p.Vertices[len(p.Vertices)-1].Y, p.Vertices[0].X, p.Vertices[0].Y)
	p.Edges = append(p.Edges, close)

	for i := 1; i < len(p.Vertices); i++ {
		v1, v2 := p.Vertices[i-1], p.Vertices[i]
		edge := NewLine(v1.X, v1.Y, v2.X, v2.Y)
		p.Edges = append(p.Edges, edge)
	}
}

// CalculateEdges computes all eges for a polygon with populated vertices.
// It draws a line from each Point to each other Point in the polygon.
// For simple viewing of .obj files, I preferred this method.
// It's almost like a cheap wireframe
func (p *Polygon) CalculateEdges() {

	var lines []Line

	lineChannel := make(chan Line)

	for _, v := range p.Vertices {
		for _, v2 := range p.Vertices {
			go func(v, v2 Point) {
				l := NewLine(v.X, v.Y, v2.X, v2.Y)
				lineChannel <- l
			}(v, v2)
		}
	}

	for i := 0; i < len(p.Vertices)*len(p.Vertices); i++ {
		l := <-lineChannel
		lines = append(lines, l)
	}

	p.Edges = lines

}

// Print iterates through all edges of the polygon and prints them.
func (p *Polygon) Print(clr int) {
	w, h := termbox.Size()
	cX, cY := w/2, h/2

	for _, edge := range p.Edges {
		edge.Print(
			func(x, y int) {
				termbox.SetCell((x/2)+cX, (y-2)+cY, '*', termbox.ColorWhite, termbox.ColorDefault)
			},
		)
	}

	termbox.Flush()
}

// LoadObj loads a slice of Points from a .obj file
func LoadObj(p string) []Point {
	var points []Point

	pointChannel := make(chan Point)

	objData := ReadObjFile(p)
	for _, val := range objData {
		go func(val Coordinate) {
			p := val.ToPoint()
			pointChannel <- p
		}(val)
	}

	for i := 0; i < len(objData); i++ {
		p := <-pointChannel
		points = append(points, p)
	}

	return points
}

// ReadObjFile returns a slice of 3d Coordinates from a .obj file
func ReadObjFile(path string) []Coordinate {

	var coords []Coordinate

	readFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		if strings.HasPrefix(text, "v ") {
			current := make([]int, 3)
			values := strings.Split(text, " ")[1:]
			for i, v := range values {
				c := StoFtoI(v)
				current[i] = c
			}
			coords = append(coords, Coordinate{current[0], current[1], current[2]})
		}
	}

	readFile.Close()

	return coords

}
