package stl

import (
	"errors"
	"io/ioutil"
	"os"
)

type Parser struct {
	content string
}

func NewParser(path string) (p *Parser, err error) {
	p = new(Parser)

	file, err := os.Open(path)
	if err != nil {
		err = errors.New("error opening file")
		return
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		err = errors.New("error reading file")
		return
	}

	p.content = string(bytes)

	return
}

// A STL file starts with solid: <name>

//facet normal ni nj nk
//    outer loop
//        vertex v1x v1y v1z
//        vertex v2x v2y v2z
//        vertex v3x v3y v3z
//    endloop
//endfacet

// It will end with: endsolid <name>
func (p *Parser) Parse() {

}