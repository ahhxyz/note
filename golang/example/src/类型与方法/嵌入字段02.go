package main

import (
	"fmt"
)

type User struct {
	name string
}
type App struct {
	*User
}

func (u *User) Get() {
	fmt.Println(u.name)
}
func (a *App) Get() {
	fmt.Println(a.name)

}
func main() {
	
	a=&App{new(User)}//或a=&App{&User}} //嵌入了指针类型，就必须这样实例化

	a.name = "霍去病"
	a.Get()
	

}
