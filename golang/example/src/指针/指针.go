package main

import (
//"fmt"
//"reflect"
)

type User struct {
	name string
}

func (u *User) Get() {
	println(u.name)
}
/*
 *如果这里的方法接收者是User类型，那么下面将不会输出任何内容
 *我们知道，方法接收者本质是函数的一个参数，而参数都是值类型的传递，
 *1.如果方法接收者不是指针类型，那么这里的u其实是实例u1的一个副本，
 *所以这时实际设置的是实例u1副本的字段值，实例u1本身自然就无法获取了任何内容了。
 *2.如果方法接收者是指针类型，那么这里的u同样是实例u2的一个副本,
 *但这个副本和u2指向同一个地址，所以这时实际设置的是同一个实例的字段值。
 */
func (u *User) Set(name string) { 
	u.name = name
}
func main() {
	var u1 User
	u2 := new(User)

	u1.Set("青春")
	u1.Get()

	u2.Set("青春2")
	u2.Get()

}
