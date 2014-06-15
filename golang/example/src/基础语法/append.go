package main

import "fmt"

func main() {
	var res [][]string
	data := make([]string, 2)
	data[0] = "预期值"
	res = append(res, data)
	data[0] = "新值"
	data[1] = "新增值"
	fmt.Printf("%s\n", res)
}
