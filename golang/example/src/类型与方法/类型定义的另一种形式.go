package main

import (
	"fmt"
)

type A struct {
	B struct {
		C string
	}
}

func main() {
	u := A{B: struct{ C string }{"甘延寿"}}
	fmt.Println(u.B.C)
}
