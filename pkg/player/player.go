package player

import (
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

type Player struct {
	orientation mgl64.Vec3
	position mgl64.Vec3
}

func NewPlayer() *Player {
	p := new(Player)
	return p
}

func (p *Player) GetPosition() mgl64.Vec3 {
	return p.position
}

func (p *Player) GetOrientation() mgl64.Vec3 {
	return p.orientation
}

func (p *Player) Translate(x, y, z float64) {
	vec := mgl64.Vec3{x, y, z}
	p.position = p.position.Add(vec)
}

func (p *Player) Rotate(x, y float64) {
	angleX := math.Mod(p.orientation.X() + x, 360.0)
	angleY := math.Mod(p.orientation.Y() + y, 360.0)

	p.orientation = mgl64.Vec3{angleX, angleY}
}
func (p *Player) Forward(val float64) {

}
func (p *Player) Backward(val float64) {

}
func (p *Player) Left(val float64) {

}
func (p *Player) Right(val float64) {

}
