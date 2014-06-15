package main

import (
	//"fmt"
	"os"
	"reflect"
)

func main() {
	dir, err := os.Open("dir")

	list, err := dir.Readdir(-1)
	dir.Close()
	println(len(list))

	for _, fi := range list {
		info := fi.Sys()
		v := reflect.ValueOf(info).Elem()
		Atim := v.FieldByName("Atim").Field(0).Int()
		println("\n", Atim)
	}

	//println(list)
	println(err)
}
