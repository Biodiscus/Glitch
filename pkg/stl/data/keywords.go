package data

import "regexp"

var SolidRegex = regexp.MustCompile(`^solid ([a-zA-Z_])+$`)
var EndSolidRegex = regexp.MustCompile(`^endsolid ([a-zA-Z_])+&`)

var FacetRegex = regexp.MustCompile(`facet normal .*`)
var FloatRegex = regexp.MustCompile(`([-+]?[0-9]*\.?[0-9]+)`)

var EndFacetRegex = regexp.MustCompile(`endfacet`)
var OuterLoopRegex = regexp.MustCompile(`outer loop`)
var EndLoopRegex = regexp.MustCompile(`endloop`)

var VertexRegex = regexp.MustCompile(`vertex .*`)
