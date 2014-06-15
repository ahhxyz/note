package main

import (
	"fmt"
)

func InsertSort(a []int) []int {
	var i, j int
	for i = 1; i < len(a); i++ {
		key := a[i]
		for j = i - 1; j >= 0 && a[j] > key; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = key
	}
	return a
}

func main() {
	var num int
	var a []int
	fmt.Printf("本次排序需要输入5个数字：\n")
	for i := 0; i < 5; i++ {
		fmt.Printf("请输入一个数字并回车：\n")
		fmt.Scanf("%d", &num)
		a = append(a, num)
	}
	fmt.Printf("未排序：%v\n", a)
	fmt.Printf("已排序：%v\n", InsertSort(a))
}
