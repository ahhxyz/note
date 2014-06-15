package main

import (
	"fmt"
	"reflect"
)

var Ins map[string]interface{}

type User struct {
	name string
}

func (u *User) Get() {
	println("demo")
}
func main() {
	u2 := new(User)
	Ins = map[string]interface{}{"User": u2}
	//Ins["User"].Get()
	v := reflect.ValueOf(Ins["User"])
	fmt.Printf("%s\n", v)
	method := v.MethodByName("Get")
	in := make([]reflect.Value, 0)
	method.Call(in)
}
