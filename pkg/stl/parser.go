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

func (p *Parser) Parse() {

}