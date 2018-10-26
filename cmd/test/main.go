package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	player2 "glitter.jemoeders.website/biodiscus/glitch/pkg/player"
	"glitter.jemoeders.website/biodiscus/glitch/pkg/stl"
	"glitter.jemoeders.website/biodiscus/glitch/pkg/stl/data"
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
var facets []data.Facet
var up, down, left, right, escape bool

func init() {
	runtime.LockOSThread()
}

func parse() data.BeginSolid{
	p, err := stl.NewParser("data/cup/10_ml.stl")
	if err != nil {
		log.Fatal(err)
	}

	solid := p.Parse()
	return solid
}

var player player2.Player

func main() {
	facets = parse().Facets
	player.Translate(0, 0, -40.0)

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
	window.SetInputMode(glfw.CursorMode, glfw.CursorHidden)
	window.MakeContextCurrent()


	if err := gl.Init(); err != nil {
		panic(err)
	}


	setupScene()

	previousTime = glfw.GetTime()
	window.SetKeyCallback(keyCallback)
	window.SetCursorPosCallback(cursorCallback)

	for !window.ShouldClose() || escape {
		drawScene()
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func keyCallback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	state := action == glfw.Press || action == glfw.Repeat
	if key == glfw.KeyW {
		up = state
	} else if key == glfw.KeyS {
		down = state
	} else if key == glfw.KeyA {
		left = state
	} else if key == glfw.KeyD {
		right = state
	}
}

func cursorCallback (w *glfw.Window, xpos float64, ypos float64){
	centerX, centerY := getWindowCenter(w)

	angleY := (centerX - xpos) / 10.0
	angleX := (ypos - centerY) / 10.0

	player.Rotate(angleX, angleY)

	centerMouse(w)
}

func getWindowCenter(w *glfw.Window) (x, y float64){
	windowPosX, windowPosY := w.GetPos()

	centerX := float64(windowPosX) + (float64(Width) / 2.0)
	centerY := float64(windowPosY) + (float64(Height) / 2.0)
	return centerX, centerY
}

func centerMouse(w *glfw.Window) {
	centerX, centerY := getWindowCenter(w)
	w.SetCursorPos(centerX, centerY)
}

func setupScene() {
	gl.Enable(gl.DEPTH_TEST)

	gl.Enable(gl.LIGHTING)
	ambient := []float32{0.5, 0.5, 0.5, 1}
	diffuse := []float32{1, 1, 1, 1}
	lightPosition := []float32{-5, 5, 10, 0}
	gl.Lightfv(gl.LIGHT0, gl.AMBIENT, &ambient[0])
	gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, &diffuse[0])
	gl.Lightfv(gl.LIGHT0, gl.POSITION, &lightPosition[0])
	gl.Enable(gl.LIGHT0)

	gl.ClearColor(0, 0, 0, 0)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()

	//projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(Width) / Height, 0.1, 10.0)
	//projectionUniform := gl.GetUniformLocation(program, gl.Str("projection\x00"))
	//gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	//gl.Frustum(-20.0, 20.0, -20.0, 20.0, 0.1, 10.0)
	gl.Ortho(-40.0, 40.0, -40.0, 40.0, 1, 80.0)
}

func drawScene() {
	time := glfw.GetTime()
	diff := time - previousTime
	previousTime = time
	angle += diff * 10.0

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	if up || down || left || right {
		if up {
			player.Translate(0, 0, 0.1)
		} else if down {
			player.Translate(0, 0, -0.1)
		}
		//if up {
		//	player.Forward(0.01)
		//} else if down {
		//	player.Backward(-0.01)
		//}

		if left {
			//player.Left(0.01)
			player.Translate(0.01, 0, 0)
		} else if right {
			//player.Right(-0.01)
			player.Translate(-0.01, 0, 0)
		}
	}

	pos := player.GetPosition()
	orientation := player.GetOrientation()

	gl.Translated(pos.X(), pos.Y(), pos.Z())
	gl.Rotated(orientation.X(), 1.0, 0.0, 0.0)
	gl.Rotated(orientation.Y(), 0.0, 1.0, 0.0)

	gl.Color3f(1,0 ,0 )

	gl.PushMatrix()
	gl.Rotated(angle, 1.0, 0.0, 0.0)

	gl.Begin(gl.TRIANGLES)
	for _, facet := range facets {
		gl.Color3f(1, 0, 0)
		gl.Normal3f(facet.I, facet.J, facet.K)

		for _, vertex := range facet.Vertices {
			gl.Vertex3f(vertex.X, vertex.Y, vertex.Z)
		}
	}

	gl.End()
	gl.PopMatrix()

}
