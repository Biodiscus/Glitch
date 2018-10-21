package data

import "fmt"

// solid <name>, name can be omitted, but if that's the case a space has to be there
type BeginSolid struct {
	Name string

	Facets []Facet

	EndFile EndSolid
}

func (b BeginSolid) String() string {
	return fmt.Sprint("BeginSolid {", "Name:", b.Name, ", Facets: ", b.Facets, ", EndSolid: ", b.EndFile, "}")
}

// endsolid <name>
type EndSolid struct {
	Name string
}

func (e EndSolid) String() string {
	return fmt.Sprint("EndSolid {", "Name:", e.Name, "}")
}