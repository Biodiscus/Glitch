package main

import (
	"glitter.jemoeders.website/biodiscus/glitch/pkg/mvc/controller"
	"glitter.jemoeders.website/biodiscus/glitch/pkg/mvc/driver"
	"glitter.jemoeders.website/biodiscus/glitch/pkg/window"
	"log"
	"runtime"
)

const Width = 500
const Height = 500

func init() {
	runtime.LockOSThread()
}

var mvcDriver driver.MvcDriver

func main() {
	w, err := window.NewWindow(Width, Height, "Testing")
	if err != nil {
		log.Fatal("Error when creating a new window: ", err)
	}

	setupMVC()

	w.SetRunStep(step)
	w.Run()
}

func setupMVC() {
	mvcDriver.SetController(controller.HomeController{})
}

func step(delta float64) {

	mvcDriver.Update(delta)
	mvcDriver.Render(delta)

}