package main

import (
	"fmt"
	"reflect"
	p "test/path"
)

func main() {
	var o p.One

	fmt.Printf(":%v\n", reflect.TypeOf(o).PkgPath()) //<*main.Two Value>
}
