package main

import (
	"math/rand"
	"runtime"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	w            = 1080
	h            = 1080
	fps          = 75
	circlesCount = 10
)

func main() {

	// Init
	runtime.LockOSThread()

	r, err := NewRenderer(w, h, "Test")
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	// Generate Circles
	var circles []*Entity

	for i := 0; i < circlesCount; i++ {
		min, max := float32(-0.9), float32(0.9)
		cx := min + rand.Float32()*(max-min)
		cy := min + rand.Float32()*(max-min)
		rad := 0.0025 + rand.Float32()*(0.025-0.0025)
		c := NewCircle(cx, cy, rad, w, h)
		circles = append(circles, c)
	}

	/*
		p := NewPlayer([]float32{
			-0.9, -0.8, 0,
			-0.9, -0.9, 0,
			-0.8, -0.9, 0,
		})
	*/

	// Bind Keys
	r.Window.SetKeyCallback(func(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		// u, l, d, r := p.Bounds()
		press := action == glfw.Press || action == glfw.Repeat
		switch {
		case key == glfw.KeyEscape && press:
			window.SetShouldClose(true)
			/*
				case key == glfw.KeyA && press:
					if !l {
						p.Left()
					}
				case key == glfw.KeyD && press:
					if !r {
						p.Right()
					}
				case key == glfw.KeyW && press:
					if !u {
						p.Up()
					}
				case key == glfw.KeyS && press:
					if !d {
						p.Down()
					}
			*/
		}
	})

	// Render
	r.Render(fps, func() {
		// p.Draw(gl.LINES)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.UseProgram(*r.Program)
		for _, cir := range circles {
			cir.ScreenSaver()
			cir.Draw(gl.POINTS)
		}
	})

}
