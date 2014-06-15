package main

import (
	"fmt"
)

func foo(s interface{}) {
	fmt.Printf("%v\n", s)
}
func main() {
	funcs := map[string]func(in interface{}){"f": foo}
	funcs["f"]("hi")

}
