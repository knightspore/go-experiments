package main

type Point struct {
	X, Y int
}

type Coordinate struct {
	X, Y, Z int
}

type Line struct {
	Points []Point
}

// ToPoint returns a 3d Co-Ordinate as a 2d Point
func (c *Coordinate) ToPoint() Point {

	// TODO
	// This does not cover camera or perspective.
	// It's simple, and passes in that we get a simple
	// preview of a .obj file in the terminal.

	var scale int

	switch {
	case c.X < 5:
		scale = 10
	case c.X < 10:
		scale = 5
	default:
		scale = 1
	}

	x := (c.X * c.Z) * scale
	y := (c.Y * c.Z) * scale

	return Point{x, y}
}

// Print iterates through all points on a line accepting a Print Function and a Refresh Function.
func (l *Line) Print(printFn func(x, y int)) {
	for _, pt := range l.Points {
		go printFn(pt.X, pt.Y)
	}
}

// Ends returns the start and end points of a line
func (l *Line) Ends() (Point, Point) {
	return l.Points[0], l.Points[len(l.Points)-1]
}

// NewLine returns a line of calculated points using Bresenham's line algorithm.
// https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
func NewLine(x0, y0, x1, y1 int) Line {

	l := Line{}

	if Abs(y1-y0) < Abs(x1-x0) {
		if x0 > x1 {
			l.plotLow(x1, y1, x0, y0)
		} else {
			l.plotLow(x0, y0, x1, y1)
		}
	} else {
		if y0 > y1 {
			l.plotHigh(x1, y1, x0, y0)
		} else {
			l.plotHigh(x0, y0, x1, y1)
		}
	}

	return l

}

// plotLow covers extended slopes of Bresenham's line algorithm.
func (l *Line) plotLow(x0, y0, x1, y1 int) {

	dx := x1 - x0
	dy := y1 - y0
	yi := 1

	if dy < 0 {
		yi = -1
		dy = -dy
	}

	D := (2 * dy) - dx
	y := y0

	for x := x0; x <= x1; x++ {
		l.Points = append(l.Points, Point{x, y})

		if D > 0 {
			y = y + yi
			D = D + (2 * (dy - dx))
		} else {
			D = D + 2*dy
		}
	}
}

// plotHigh covers extended slopes of Bresenham's line algorithm.
func (l *Line) plotHigh(x0, y0, x1, y1 int) {

	dx := x1 - x0
	dy := y1 - y0
	xi := 1

	if dx < 0 {
		xi = -1
		dx = -dx
	}

	D := (2 * dx) - dy
	x := x0

	for y := y0; y <= y1; y++ {
		l.Points = append(l.Points, Point{x, y})

		if D > 0 {
			x = x + xi
			D = D + (2 * (dx - dy))
		} else {
			D = D + 2*dx
		}
	}

}
