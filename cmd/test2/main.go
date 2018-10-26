package main

import (
	"glitter.jemoeders.website/biodiscus/glitch/pkg/window"
	"log"
	"runtime"
)

const Width = 500
const Height = 500

func init() {
	runtime.LockOSThread()
}

func main() {
	w, err := window.NewWindow(Width, Height, "Testing")
	if err != nil {
		log.Fatal("Error when creating a new window: ", err)
	}

	w.SetRunStep(step)
	w.Run()
}

func step(delta float64) {
}