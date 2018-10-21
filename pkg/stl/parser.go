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
func (p *Parser) Parse() data.BeginSolid {
	var solid data.BeginSolid
	var err error

	lines := strings.Split(p.content, NewLine)

	facets := make([]data.Facet, 0)

	// For now assume that the files are valid and can only contain 1 solid
	for index, line := range lines {
		if data.FacetRegex.MatchString(line) {
			//index := p.parseFacet(index, lines)
			facet, err := p.parseFacet(index, lines)
			if err != nil {
				log.Fatal("Error parsing STL file: ", err)
			}

			facets = append(facets, facet)
		} else if data.SolidRegex.MatchString(line) {
			solid.Name, err = p.parseName(line)
			if err != nil {
				log.Fatal("Error parsing solid: ", err)
			}
		} else if data.EndSolidRegex.MatchString(line) {
			solid.EndFile.Name, err = p.parseName(line)
			if err != nil {
				log.Fatal("Error parsing end solid: ", err)
			}

			// Break for now
			break;
		}
	}

	solid.Facets = facets

	return solid
}

func (p *Parser) parseName(line string) (string, error) {
	arr := strings.Split(line, " ")
	if len(arr) != 2 {
		return "", errors.New("expected the second argument to be a name, got: "+line)
	}

	return arr[1], nil
}

func (p *Parser) parseFacet(index int, lines []string) (data.Facet, error) {
	var err error

	facet := data.Facet{}

	// After a facet declaration, there should be a outer loop
	outer := lines[index+1]
	if !data.OuterLoopRegex.MatchString(outer) {
		return facet, errors.New("expected a `outer loop` after a `facet`, found: "+outer)
	}

	vertices := make([]data.Vertex, 0)
	// After the outer come the vertices
	for _, line := range lines[index+2:] {
		if data.EndFacetRegex.MatchString(line) {
			break;
		} else  if data.VertexRegex.MatchString(line){
			vertex, err := p.parseVertex(line)
			if err != nil {
				log.Fatal(err)
			}
			vertices = append(vertices, vertex)
		}
	}

	declaration := lines[index]
	facet, err = p.parseFacetDeclaration(declaration)
	facet.Vertices = vertices

	if err != nil {
		return facet, errors.New("expected a I J K value in the facet declaration, got: "+declaration)
	}


	return facet, nil
}

func (p *Parser) parseFacetDeclaration(declaration string) (data.Facet, error){
	floats := data.FloatRegex.FindAllString(declaration, -1)

	var err error

	facet := data.Facet{}

	facet.I, err = p.parseFloat(floats[0])
	if err != nil {
		return facet, errors.New("error parsing I of facet")
	}

	facet.J, err = p.parseFloat(floats[1])
	if err != nil {
		return facet, errors.New("error parsing J of facet")
	}

	facet.K, err = p.parseFloat(floats[2])
	if err != nil {
		return facet, errors.New("error parsing K of facet")
	}


	return facet, nil
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