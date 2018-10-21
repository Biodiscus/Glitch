package data

import "fmt"

// facet normal ni nj nk
// 		outer loop
//			...
//		endloop
// endfacet
type Facet struct {
	I float32
	J float32
	K float32

	Vertices []Vertex
}

func (f Facet) String() string {
	return fmt.Sprint("Facet { I: ", f.I, ", J: ", f.J, ", K: ", f.K, ", Vertices: ", f.Vertices,  "}")
}