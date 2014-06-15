package main

import (
	"fmt"
	"reflect"
)

type One string
type Two struct {
	One
	Name string
}
type I interface {
	Say()
}

func (o *One) Say() {
	//fmt.Printf("%v\n", p.Name)
	t := reflect.TypeOf(o)  //	*main.One
	v := reflect.ValueOf(o) //	<*main.One Value>
	fmt.Printf(":%v\n", t)
	fmt.Printf(":%v\n", v)

}
func main() {
	var i I
	i = new(Two)
	i.Say()
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	fmt.Printf(":%v\n", t) //*main.Two
	fmt.Printf(":%v\n", v) //<*main.Two Value>
}
