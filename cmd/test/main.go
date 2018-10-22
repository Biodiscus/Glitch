package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"log"
	"runtime"
)

const Width = 500
const Height = 500

var Triangle = []float32{
	0, 0.5, 0, // top
	-0.5, -0.5, 0, // left
	0.5, -0.5, 0, // right
}

func main() {
	runtime.LockOSThread()

	window := initGlfw()
	defer window.Destroy()

	program := initOpenGl()

	vao := makeVao(Triangle)


	for !window.ShouldClose() {
		draw(vao, window, program)
	}
}

func makeVao(points []float32) uint32 {
	var vbo uint32
	var vao uint32

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4 * len(points), gl.Ptr(points), gl.STATIC_DRAW)

	gl.GenVertexArrays(	1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

func draw(vao uint32, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(Triangle)))

	glfw.PollEvents()
	window.SwapBuffers()
}

func initOpenGl() uint32 {
	if err := gl.Init(); err != nil {
		log.Fatal("Error initializing OpenGl: ", err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL Version:", version)

	program := gl.CreateProgram()
	gl.LinkProgram(program)

	return program
}

func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		log.Fatal("Error initializing window: ", err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(Width, Height, "Testing Window", nil, nil)
	if err != nil {
		log.Fatal("Error creating window: ", err)
	}
	window.MakeContextCurrent()

	return window
}