package main

import "fmt"

type S struct {
	i string
}

func (p S) Get() string {
	return p.i
}
func (p S) Put(v string) {
	p.i = v
}

type I interface {
	Get() string
	Put(v string)
}

func main() {
	p := new(S) //这里不能定义成指针类型或定义成接口类型；
	var i I
	p.i = "Hello World!"
	i = p
	fmt.Println(i.Get())
}
