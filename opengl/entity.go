package main

import (
	"math"
	"math/rand"

	"github.com/go-gl/gl/v4.6-core/gl"
)

type Entity struct {
	Points     []float32
	vao        *uint32
	vbo        *uint32
	DirectionX Direction
	DirectionY Direction
	Speed      float32
}

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

func (d Direction) String() string {
	return [...]string{"Up", "Left", "Down", "Right"}[d]
}

func (e *Entity) MakeVao() {
	gl.GenBuffers(1, e.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, *e.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(e.Points), gl.Ptr(e.Points), gl.STATIC_DRAW) // 4 x len(points) accounts for the 4-bytes that make a up a float32 (32-bits)

	gl.GenVertexArrays(1, e.vao)
	gl.BindVertexArray(*e.vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, *e.vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
}

func NewPlayer(points []float32) *Entity {
	var vao, vbo uint32
	e := &Entity{points, &vao, &vbo, Left, Up, 0.05}
	e.MakeVao()
	return e
}

func newEntity(points []float32) *Entity {
	var vao, vbo uint32
	e := &Entity{points, &vao, &vbo, Direction(rand.Intn(2)), Direction(rand.Intn(2)) + 2, 0.02}
	e.MakeVao()
	return e
}

func NewCircle(cx, cy, rad float32, w, h int) *Entity {
	var floats []float32
	deg := w / h
	for t := float64(0); t < 2*math.Pi; t += 0.01 {
		var x, y, z float32
		x = cx + rad*float32(math.Cos(t))
		y = cy + rad*float32(deg)*float32(math.Sin(t))
		z = rand.Float32()
		floats = append(floats, []float32{x, y, z}...)
	}
	return newEntity(floats)
}

func NewSquare(start, size float32) *Entity {
	points := []float32{
		start, start + size, 0, // Top Left
		start + size, start + size, 0, // Top Right
		start, start, 0, // Bottom Left
		start + size, start + size, 0, // Top Right
		start + size, start, 0, // Bottom Right
		start, start, 0, // Bottom Left
	}
	return newEntity(points)
}

// Draw Functions

func (e *Entity) Draw(style uint32) {
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(e.Points), gl.Ptr(e.Points), gl.DYNAMIC_DRAW) // 4 x len(points) accounts for the 4-bytes that make a up a float32 (32-bits)
	gl.BindBuffer(gl.ARRAY_BUFFER, *e.vbo)
	gl.BindVertexArray(*e.vao)
	gl.DrawArrays(style, 0, int32(len(e.Points)/2)) // TODO: Split Points, Lines, Tris
}

func (e *Entity) Bounds() (bool, bool, bool, bool) {

	var up, left, down, right bool

	for i := 1; i < len(e.Points); i += 3 {
		y := e.Points[i]
		if y >= 0.9 {
			up = true
			break
		}
	}

	for i := 0; i < len(e.Points); i += 3 {
		x := e.Points[i]
		if x <= -0.9 {
			left = true
			break
		}
	}

	for i := 1; i < len(e.Points); i += 3 {
		y := e.Points[i]
		if y <= -0.9 {
			down = true
			break
		}
	}

	for i := 0; i < len(e.Points); i += 3 {
		x := e.Points[i]
		if x >= 0.9 {
			right = true
			break
		}
	}

	return up, left, down, right

}

func (e *Entity) Up() {
	e.AnimateY(func(i int) {
		e.Points[i] += e.Speed
	})
}

func (e *Entity) Left() {
	e.AnimateX(func(i int) {
		e.Points[i] -= e.Speed
	})
}

func (e *Entity) Down() {
	e.AnimateY(func(i int) {
		e.Points[i] -= e.Speed
	})
}

func (e *Entity) Right() {
	e.AnimateX(func(i int) {
		e.Points[i] += e.Speed
	})
}

func randSpeed() float32 {
	min, max := float32(0.001), float32(0.005)
	return min + rand.Float32()*(max-min)
}

func (e *Entity) Animate(f func(i int)) {
	for i := 0; i < len(e.Points); i++ {
		f(i)
	}
}

func (e *Entity) AnimateY(f func(i int)) {
	for i := 1; i < len(e.Points); i += 3 {
		f(i)
	}
}

func (e *Entity) AnimateX(f func(i int)) {

	for i := 0; i < len(e.Points); i += 3 {
		f(i)
	}
}

func (e *Entity) ScreenSaver() {
bounds:
	for i := 0; i < len(e.Points); i += 3 {
		x, y := e.Points[i], e.Points[i+1]

		switch {
		case x >= 1.0:
			e.DirectionX = Left
		case x <= -1.0:
			e.DirectionX = Right
		}

		switch {
		case y >= 1.0:
			e.DirectionY = Down
		case y <= -1.0:
			e.DirectionY = Up
		}

		break bounds
	}

	switch e.DirectionX {
	case Left:
		e.Left()
	case Right:
		e.Right()
	}

	switch e.DirectionY {
	case Up:
		e.Up()
	case Down:
		e.Down()
	}
}
