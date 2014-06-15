package main

import (
	"fmt"

	"flag"
)

func main() {

	flag.Parse()

	fmt.Println(flag.Args())

	//$ go run flag.go 1 2
	//output:[1 2]

}
