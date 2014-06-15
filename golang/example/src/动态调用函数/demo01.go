package main

import (
	"动态调用函数/test"
)

func main() {
	funcs := map[string]func(in ...interface{}){"demo": test.Demo}
	funcs["demo"]("动态调用其它包中的函数")

}
