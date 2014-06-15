package main

import (
	"fmt"
	"time"
)

func timer() {
	t := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-t.C: //通道阻塞，直到一秒后通道的信息发送出去
			fmt.Println("test timer")
		}
	}
}

func main() {
	go timer() //前面有go关键字的这条语句和后面的语句是同时执行的，这就是并发.并发的作用是执行此条语句的同时可以执行后面的语句，尤其是某个语句要不停地执行时。
	fmt.Println("test")

	timer()
}
