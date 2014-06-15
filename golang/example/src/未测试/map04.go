package main

import "fmt"

type person struct {
	name, position string
}

var m map[string]person

func main() {
	m = make(map[string]person)
	m["info01"] = person{
		"陈汤", "左金吾大将军",
	}
	m["info02"] = person{
		"甘延寿", "右武卫大将军",
	}
	fmt.Println(m["info01"])
	fmt.Println(m)
}
