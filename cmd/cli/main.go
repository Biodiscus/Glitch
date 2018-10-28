package main

import (
	screen2 "glitter.jemoeders.website/biodiscus/glitch/pkg/screen"
	"glitter.jemoeders.website/biodiscus/glitch/pkg/window"
	"log"
	"runtime"
)

const Width = 800
const Height = 600

func init() {
	runtime.LockOSThread()
}

var screen *screen2.Home

func main() {
	w, err := window.NewWindow(Width, Height, "Testing")
	if err != nil {
		log.Fatal("Error when creating a new window: ", err)
	}

	screen = new(screen2.Home)
	screen.Initialize()

	w.SetRunStep(step)
	w.Run()
}

func step(delta float64) {
	screen.Update(delta)
	screen.Render(delta)
}