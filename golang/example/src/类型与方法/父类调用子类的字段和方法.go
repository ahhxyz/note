package main

import (
	"fmt"
	"reflect"
)

type Model struct {
	View string
	Self interface{}
}
type In interface {
	Get()
}

func (m *Model) Get() {
	v := reflect.ValueOf(m.Self)
	self := v.Elem().FieldByName("Model")
	fmt.Printf("%s\n%s\n%s\n", v.Elem().FieldByName("Id").String(), self.FieldByName("View").String(), m.View)
	method := v.MethodByName("G")
	in := []reflect.Value{}
	method.Call(in)
}

type User struct {
	Model
	Id string
}

func (u *User) G() {
	println("子类的方法")
}
func main() {

	u := new(User)
	u.Self = u
	u.Id = "子类的字段值"
	u.View = "父类的字段值"
	u.Get()

}
