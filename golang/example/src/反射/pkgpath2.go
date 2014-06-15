package main

import (
	//"fmt"
	//"reflect"
	p "test/path"
)

type Two struct {
	p.One
	t string
}

func main() {
	t := new(Two)
	t.Demo()

}
