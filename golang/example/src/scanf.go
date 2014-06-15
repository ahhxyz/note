package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Printf("请输入：\n")
	fmt.Scanf("%s", &s) //如果使用了fmt.Scan()，那么，程序不会等待用户的输入而跳过此语句执行完其它语句后就会退出。
	fmt.Printf("输入的是：%s\n", s)
}
