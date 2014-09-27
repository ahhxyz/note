package main

import (
	"sync"
)

var l sync.Mutex //线程锁
var a string     //共享变量

func f() {
	a = "hello, world" //在线程中修改共享变量的值
	l.Unlock()         //释放线程锁
}

func main() {
	l.Lock() //开启线程锁，作用时在后面的线程为处理完共享变量之前，其它线程不能修改它的值
	go f()   //开启一个线程
	l.Lock()
	print(a)
}
