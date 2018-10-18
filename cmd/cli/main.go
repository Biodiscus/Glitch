package main

import (
	"glitter.jemoeders.website/biodiscus/glitch/pkg/stl"
	"log"
)

func main() {
	p, err := stl.NewParser("data/cup/100_ml.stl")
	if err != nil {
		log.Fatal(err)
	}

	p.Parse()
}
