package window

import (
	"errors"
	"fmt"
	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Window struct {
	width int
	height int

	title string

	step RunCallback

	previousTime float64

	glfwWindow *glfw.Window
}

type RunCallback func(delta float64)

func NewWindow(width, height int, title string) (*Window, error) {
	w := new(Window)
	w.width = width
	w.height = height
	w.title = title

	err := w.setupGLFW()
	if err != nil {
		return nil, err
	}

	err = w.setupScene()
	if err != nil {
		return nil, err
	}


	return w, nil
}

func (w *Window) setupGLFW() error {
	if err := glfw.Init(); err != nil {
		return errors.New(fmt.Sprint("failed to initialize glfw:", err))
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

	window, err := glfw.CreateWindow(w.width, w.height, w.title, nil, nil)
	if err != nil {
		return errors.New(fmt.Sprint("failed to create glfw window:", err))
	}

	window.SetInputMode(glfw.CursorMode, glfw.CursorHidden)
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		return errors.New(fmt.Sprint("failed to initialize OpenGL: ", err))
	}

	w.glfwWindow = window

	return nil
}

func (w *Window) setupScene() error {
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
	//gl.Frustum(-20.0, 20.0, -20.0, 20.0, 0.1, 10.0)
	gl.Ortho(-40.0, 40.0, -40.0, 40.0, 1, 80.0)

	return nil
}

func (w *Window) SetRunStep(call RunCallback) {
	w.step = call
}

func (w *Window) Run() {
	w.previousTime = glfw.GetTime()

	for !w.glfwWindow.ShouldClose() { //|| escape
		// Calculate the delta time to give the step callback
		time := glfw.GetTime()
		delta := time - w.previousTime
		w.previousTime = time


		if w.step != nil {
			w.step(delta)
		}

		w.glfwWindow.SwapBuffers()
		glfw.PollEvents()
	}
}

func (w *Window) Destroy() {
	glfw.Terminate()
}
