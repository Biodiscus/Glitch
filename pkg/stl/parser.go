package stl

import (
	"errors"
	"glitter.jemoeders.website/biodiscus/glitch/pkg/stl/data"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const NewLine = "\n"

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
	lines := strings.Split(p.content, NewLine)

	// For now assume that the files are valid and can only contain 1 solid
	for index, line := range lines {
		if data.FacetRegex.MatchString(line) {
			p.parseFacet(index, lines)
		}
	}
}

func (p *Parser) parseFacet(index int, lines []string) (int, error) {
	newIndex := index + 4

	// After a facet decleration, there should be a outer loop
	outer := lines[index+1]
	if !data.OuterLoopRegex.MatchString(outer) {
		return -1, errors.New("expected a `outer loop` after a `facet`, found: "+outer)
	}

	vertexes := make([]data.Vertex, 0)
	// After the outer come the vertices
	for i, line := range lines[index+2:] {
		if data.EndLoopRegex.MatchString(line) {
			newIndex = i
			break;
		} else {
			vertex, err := p.parseVertex(line)
			if err != nil {
				log.Fatal(err)
			}
			vertexes = append(vertexes, vertex)
		}
	}


	log.Println(vertexes)
	return newIndex, nil
}

func (p *Parser) parseVertex(line string) (data.Vertex, error) {
	floats := data.FloatRegex.FindAllString(line, -1)

	var err error

	vertex := data.Vertex{}

	vertex.X, err = p.parseFloat(floats[0])
	if err != nil {
		return vertex, errors.New("error parsing X of vertex")
	}

	vertex.Y, err = p.parseFloat(floats[1])
	if err != nil {
		return vertex, errors.New("error parsing Y of vertex")
	}

	vertex.Z, err = p.parseFloat(floats[2])
	if err != nil {
		return vertex, errors.New("error parsing Z of vertex")
	}


	return vertex, nil
}

// strconv.ParseFloat actually returns a float64 (double), even if we specify the bitSize to be 32
// So this will cast it as well
func (p *Parser) parseFloat(str string) (float32, error){
	val, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0.0, err
	}
	return float32(val), err
}