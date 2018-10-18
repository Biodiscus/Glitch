package data

// solid <name>, name can be omitted, but if that's the case a space has to be there
type BeginFile struct {
	Name string

	Facets []Facet

	EndFile EndFile
}

// endsolid <name>
type EndFile struct {
	Name string
}
