package main

import (
	"io/ioutil"
)

func main() {
	contents, _ := ioutil.ReadFile("demo.txt")
	println(string(contents))
	ioutil.WriteFile("filename", contents, 0x777)
}
