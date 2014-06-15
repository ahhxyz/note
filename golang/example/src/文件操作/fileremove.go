package main

import (
	//"fmt"
	"os"
)

func main() {
	wd, _ := os.Getwd()
	//fmt.Printf(wd)
	os.Remove(wd + "/del/del")
}
