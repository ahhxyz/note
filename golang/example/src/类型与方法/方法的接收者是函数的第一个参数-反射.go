package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
}
type Member struct {
	User
	Age int
}

func (u *User) Hi() {
	v := reflect.ValueOf(u)
	fmt.Println("%v\n", v) //putput:<*main.User Value>
}
func main() {
	m := Member{Age: 1000}
	(*Member).Hi(&m)
}
