package data

import "fmt"

type Vertex struct {
	X float32
	Y float32
	Z float32
}

func (v Vertex) String() string {
	return fmt.Sprint("Vertex { X: ", v.X, ", Y: ", v.Y, ", Z: ", v.Z, "}")
}