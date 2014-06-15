package main

import (
	"fmt"
	"reflect"
)

type User struct {
	id   int
	name string
}

func (u User) Hello() {
	fmt.Println("动态调用方法")
}
func main() {
	var u User
	t := reflect.TypeOf(u)
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Type)
	}
}
