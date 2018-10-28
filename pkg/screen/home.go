package screen

import (
	"github.com/go-gl/gl/v2.1/gl"
	"glitter.jemoeders.website/biodiscus/glitch/pkg/player"
	"glitter.jemoeders.website/biodiscus/glitch/pkg/stl"
	"glitter.jemoeders.website/biodiscus/glitch/pkg/stl/data"
	"glitter.jemoeders.website/biodiscus/glitch/pkg/window"
	"log"
)

type Home struct {
	window.Screen

	solid data.BeginSolid
	angle float64

	player player.Player
}

func (h *Home) Initialize() {
	p, err := stl.NewParser("data/cup/10_ml.stl")
	if err != nil {
		log.Fatal("Error parsing STL file:", err)
	}
	h.solid = p.Parse()


	h.player.Translate(0, 0, -40.0)
}

func (h *Home) Update(delta float64) {
	h.angle += delta * 10.0
	pos := h.player.GetPosition()
	orientation := h.player.GetOrientation()

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	gl.Translated(pos.X(), pos.Y(), pos.Z())
	gl.Rotated(orientation.X(), 1.0, 0.0, 0.0)
	gl.Rotated(orientation.Y(), 0.0, 1.0, 0.0)
}

func (h * Home) Render(delta float64) {
	gl.Color3f(1,0 ,0 )

	gl.PushMatrix()
	gl.Rotated(h.angle, 1.0, 0.0, 0.0)

	gl.Begin(gl.TRIANGLES)
	for _, facet := range h.solid.Facets {
		gl.Color3f(1, 0, 0)
		gl.Normal3f(facet.I, facet.J, facet.K)

		for _, vertex := range facet.Vertices {
			gl.Vertex3f(vertex.X, vertex.Y, vertex.Z)
		}
	}

	gl.End()
	gl.PopMatrix()
}



