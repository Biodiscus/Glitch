package main

import (
	"glitter.jemoeders.website/biodiscus/glitch/pkg/stl"
	"log"
)

func main() {
	p, err := stl.NewParser("data/cup/10_ml.stl")
	if err != nil {
		log.Fatal(err)
	}

	solid := p.Parse()
	log.Println("Got Solid:", solid)
}
