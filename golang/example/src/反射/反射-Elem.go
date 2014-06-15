package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("类型错误或属性不可写！")
	} else {
		v = v.Elem()
	}
	if f := v.FieldByName("Name"); f.Kind() == reflect.String {
		f.Set(reflect.ValueOf("更改值"))
	}
	fmt.Printf("%s\n", v)
}
func main() {
	u := User{1, "奥巴马", 100}
	Set(&u)
	fmt.Printf("%s\n", u.Name)
}
