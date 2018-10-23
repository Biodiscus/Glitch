package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"log"
	"runtime"
)

const Width = 500
const Height = 500

var (
	texture   uint32
	rotationX float32
	rotationY float32
)

var Triangle = []float32{
	4.68435 - 7.79994 - 8.98095,
	-4.74887 - 7.62968 - 8.8107,
	4.59622 - 7.81663 - 8.99765,
	4.74887 - 7.75432 - 8.93533,
	-4.74887 - 7.62968 - 8.8107,
	4.68435 - 7.79994 - 8.98095,
}

var previousTime float64
var angle float64

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	window, err := glfw.CreateWindow(800, 600, "Test", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	setupScene()

	previousTime = glfw.GetTime()

	for !window.ShouldClose() {
		drawScene()
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func setupScene() {
	gl.Enable(gl.DEPTH_TEST)

	gl.ClearColor(0, 0, 0, 0)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Frustum(-4.0, 4.0, -4.0, 4.0, 0.1, 10.0)


	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(0.0, 0.0, -0.11)

}

func drawScene() {
	time := glfw.GetTime()
	diff := time - previousTime
	previousTime = time

	angle += diff / 100

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(0.0, 0.0, -0.11)
	//gl.Translated(0, 0, -0.11 - angle)
	gl.Rotated(angle * 100, 0.0, 1.0, 0.0)

	gl.PointSize(3.0)
	gl.Begin(gl.POINTS)
	for x := float32(-4.0); x <= 4.0; x += 0.5 {
		for y := float32(-4.0); y <= 4.0; y += 0.5 {
			for z := float32(0); z <= 2.0; z += 0.1 {
				gl.Color3f(1, 0, 0)
				gl.Vertex3f(x, y, z)
			}
		}
	}
	gl.End()
}
