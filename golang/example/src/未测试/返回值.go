package main

import "fmt"

func split(sum int) (x, y int) {//返回的不仅可以是函数的参数，也可以是任何的结果变量；
	x = sum * 4/9
	y = sum - x
	return  //如果return没有参数，那么就会将返回值列表定义的变量返回；
}

func main() {
	fmt.Println(split(17))
    //输出： 7 10
}