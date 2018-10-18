package pkg

import "log"

type Test struct {

}

func (t *Test) Hello() {
	log.Println("World!")
}
