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

func (u User) Hello(name string) {
	fmt.Println("长水校尉：", name, "，射声校尉", u.Name)
}
func main() {
	u := User{12, "甘延寿", 30}
	v := reflect.ValueOf(u)
	mv := v.MethodByName("ello")
	fmt.Printf("%v\n", mv.IsValid())
	//args := []reflect.Value{reflect.ValueOf("陈汤")} //参数必须是一个slice
	//mv.Call(args)
}
