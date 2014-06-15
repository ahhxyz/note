package main

import (
	"fmt"

	"os"
)

func main() {

	fmt.Println(os.Args)
	//$ go run osArgs.go 1 2
	//output:[/tmp/go-build424024630/command-line-arguments/_obj/exe/osArgs 1 2]

}
