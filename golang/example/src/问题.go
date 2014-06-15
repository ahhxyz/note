package main

import (
	"fmt"
)

type User struct {
	Id   int
	Name string
}

func (u *User) Hi(s string) {
	fmt.Println("Hi", u.Name, s)
}
func main() {
	u := User{Name: "world"}
	(*User).Hi(&u, "hah")
	(*User).Name //报错
}
