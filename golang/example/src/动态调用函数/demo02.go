package main

import (
	"reflect"
	"动态调用函数/test"
)

func main() {
	funcs := map[string]interface{}{"demo": test.Demo}
	/*
		     *funcs["demo"]定义时的类型为interface{}，而接收到的类型为func(...interface {})，
			 *这里要反射接收到的类型，故要使用reflect.ValueOf()而不是reflect.TypeOf()
	*/
	v := reflect.ValueOf(funcs["demo"])
	in := []reflect.Value{reflect.ValueOf("动态调用其它包中的函数")}
	v.Call(in)

}
