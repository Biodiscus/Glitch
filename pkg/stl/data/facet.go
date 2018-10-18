package data

// facet normal ni nj nk
// 		outer loop
//			...
//		endloop
// endfacet
type Facet struct {
	I float32
	J float32
	K float32

	vertices []Vertex
}