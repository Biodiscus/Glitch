package controller

import (
	"glitter.jemoeders.website/biodiscus/glitch/pkg/mvc/driver"
	"glitter.jemoeders.website/biodiscus/glitch/pkg/mvc/view"
	"log"
)

type HomeController struct{
	driver.Controller

	view view.HomeView
}

func (h HomeController) Initialize() {
}

func (h HomeController) GetView() driver.View {
	return h.view
}


func (h HomeController) Update(delta float64) {
	log.Println("Updating")
}
