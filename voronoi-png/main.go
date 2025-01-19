package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand/v2"
	"os"
)

const (
	NUM_POINTS = 20
	WIDTH      = 800
	HEIGHT     = 800
)

type Vec2 struct {
	x, y int
}

type CenterPoint struct {
	pos   Vec2
	color color.RGBA
}

func main() {
	points := make([]CenterPoint, NUM_POINTS)
	for i := 0; i < NUM_POINTS; i++ {
		points[i] = CenterPoint{
			pos:   Vec2{rand.IntN(WIDTH), rand.IntN(HEIGHT)},
			color: color.RGBA{uint8(rand.Int()), uint8(rand.Int()), uint8(rand.Int()), 255},
		}
	}

	upLeft := image.Point{0, 0}
	lowRight := image.Point{int(WIDTH), int(HEIGHT)}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			shortest := points[0]
			shortestDist := distance(Vec2{x, y}, points[0].pos)

			for _, point := range points {
				dist := distance(Vec2{x, y}, point.pos)
				if dist < shortestDist {
					shortest = point
					shortestDist = dist
				}

			}

			img.Set(x, y, shortest.color)
		}
	}

	for _, point := range points {
		img.Set(
			point.pos.x,
			point.pos.y,
			color.RGBA{0, 0, 0, 255},
		)
	}

	f, _ := os.Create("image.png")
	png.Encode(f, img)
}

func distance(v1, v2 Vec2) float64 {
	distX := math.Pow(float64(v2.x-v1.x), 2.0)
	distY := math.Pow(float64(v2.y-v1.y), 2.0)
	return math.Sqrt(distX + distY)
}
