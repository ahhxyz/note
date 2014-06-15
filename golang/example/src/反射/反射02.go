package main

import (
	"fmt"
	"reflect"
)

type User struct {
	id   int
	Name string
	Age  int
}
type Human struct {
	User
	title string
}

func (u *User) Hello() {
	fmt.Println("动态调用方法")
}
func main() {
	var u Human
	t := reflect.TypeOf(u)
	fmt.Println(t.FieldByIndex([]int{0, 2}).Name)
}
