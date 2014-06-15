package main

import (
	"fmt"
	//"reflect"
)

type User struct {
	name string
}
type App struct {
	A *User
}

func (u *User) Get() {
	fmt.Println(u.name)
}
func (a *App) Get() {
	fmt.Println(a.A.name)

}
func main() {
	a := new(App)
	a.A = new(User)
	u := new(User)
	a.A.name = "霍去病"
	a.Get()
	u.Get()

}
