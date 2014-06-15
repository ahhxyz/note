package main

import (
	"fmt"
)

type User struct {
	Id   int
	Name string
}

func (u *User) Hi(s string) {
	fmt.Println("方法接收者其实是函数的第一个参数，", u.Name, s)
}
func main() {
	u := User{Name: "类型的Name字段的值，"}
	(*User).Hi(&u, "传进来的参数")
}
