package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	vertexShaderSource = `
    #version 460
    in vec3 vp;
    void main() {
        gl_Position = vec4(vp, 1);
    }` + "\x00"

	fragmentShaderSource = `
	#version 460
	out vec4 frag_colour;
	vec3 color = vec3(1, 0.42, 0.5);
	void main() {
    frag_colour = vec4(color,1);
	}
` + "\x00"
)

type Renderer struct {
	W, H    int
	Window  *glfw.Window
	Program *uint32
}

// NewRenderer returns a Renderer with go-glfw and go-gl initialized
func NewRenderer(w, h int, title string) (*Renderer, error) {
	window, err := initGlfw(w, h, title)
	if err != nil {
		return nil, err
	}
	program, err := initOpenGL()
	if err != nil {
		return nil, err
	}
	r := &Renderer{w, h, window, &program}
	return r, nil
}

func (r *Renderer) Render(fps int, f func()) {
	step := time.Second / time.Duration(fps)
	for !r.Window.ShouldClose() {
		t := time.Now()
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.UseProgram(*r.Program)
		f()
		time.Sleep(step - time.Since(t))
		glfw.PollEvents()
		r.Window.SwapBuffers()
	}
}

func initGlfw(w, h int, title string) (*glfw.Window, error) {

	if err := glfw.Init(); err != nil {
		return nil, err
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(w, h, title, nil, nil)
	if err != nil {
		return nil, err
	}

	window.MakeContextCurrent()

	window.SetSizeCallback(func(w *glfw.Window, width, height int) {
		gl.Viewport(0, 0, int32(width), int32(height))
	})

	return window, nil

}

func initOpenGL() (uint32, error) {
	if err := gl.Init(); err != nil {
		return 0, err
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	prog := gl.CreateProgram()

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}
	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)

	gl.LinkProgram(prog)

	gl.ClearColor(0.1, 0.1, 0.2, 1.0)

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return prog, nil
}

func AttachShader(prog *uint32, vertexShaderSrc, fragmentShaderSrc string) error {

	return nil
}

// Shader Compilation
func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
