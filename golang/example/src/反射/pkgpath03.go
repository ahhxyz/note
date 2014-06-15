package main

import (
	//"fmt"
	//"reflect"
	p "pkg"
)

type Two struct {
	p.One
	t string
}

func main() {
	t := new(Two)
	t.Demo()

}
