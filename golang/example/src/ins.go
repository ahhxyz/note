package main

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	id   int
	name string
}

func main() {
	fmt.Printf("%v\n", strings.Split("fsdfs", ";"))

	var u0 User
	u0.name = "u0"
	u1 := &User{}
	u1.name = "u1"
	u2 := new(User)
	u2.name = "u2"
	fmt.Println(reflect.TypeOf(u0), reflect.TypeOf(u1), reflect.TypeOf(u2), u0.name, u1.name, u2.name)

}
