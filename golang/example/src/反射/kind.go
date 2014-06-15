package main

import (
	"fmt"
	"reflect"
)

func Hello(s ...string) string {
	return "test"
}
func ref(i interface{}) {
	fmt.Printf("%v\n", reflect.ValueOf(i).Kind())
}
func main() {
	modelFn := map[string]func(in ...string) string{
		"SetName": Hello,
	}

	ref(modelFn)
}
