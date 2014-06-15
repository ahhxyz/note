package main

import "fmt"

func main() {
	row := []string{}
	res := [][]string{}
	data := make([]string, 3)
	data[0] = "预期值01"
	data[1] = "预期值02"
	row = append(row, data...)//slice是引用类型，只有用data...这种方式时，函数append会为原来的切片row 重新分配空间,所以row不会引用原来的数组或被追加的数组。否则会引用数组data，那么后面data更改时，row也会跟着变化而出现非预期的结果。
	res = append(res, row)
	row = []string{}
	data[0] = "新预期值01"
	data[1] = "新预期值02"
	row = append(row, data...)
	res = append(res, row)
	data[0] = "测试值01"
	data[2] = "测试值02"

	fmt.Printf("%s\n", res)
}
