package test

import (
	"fmt"
)

func Demo(in ...interface{}) {
	fmt.Printf("%v\n%v\n", in, in[0])
}
