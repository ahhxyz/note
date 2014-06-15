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

func (u *User) Hello(name string) {
	fmt.Println("长水校尉：", name, "，射声校尉：", u.Name)
}
func main() {
	u := new(User)
	u.Name = "甘延寿"
	v := reflect.ValueOf(u)
	fmt.Printf("%s", reflect.TypeOf(v))
	mv := v.MethodByName("Hello")
	args := []reflect.Value{reflect.ValueOf("陈汤")} //参数必须是一个slice
	mv.Call(args)
}
