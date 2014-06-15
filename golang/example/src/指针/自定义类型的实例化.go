package main

import (
	"fmt"
)

type User struct {
	id   int
	Name string
}

func main() {
	var _U *User
	/*
	   _U.Name="汉朝" //这个操作会报错，因为计算机尚未给_U分配内存
	*/
	fmt.Println("\n", _U) //输出：<nil>

	U := new(User) //创建类型User的一个实例，分配内存并以零值填充，最后返回新分配的内存空间的地址值
	/*
	   或
	   var U *User
	   U=new(User)
	*/

	fmt.Println("\n", U) //输出： &{0 }
	var u User           //此时计算机已经为u分配了内存空间并以零值填充
	fmt.Println("\n", u) //输出：{0 }

}
