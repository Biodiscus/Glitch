package view

import "log"

type HomeView struct{}

func (h HomeView) Render(delta float64) {
	log.Println("HomeView render:", delta)
}
