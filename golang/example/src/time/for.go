package main

import (
	"fmt"
	"time"
)

func timer() {
	var t time.Duration = 2
	timer := time.NewTicker(t * time.Second)
	for {
		select {
		case <-timer.C:
			/* go */ func() {
				fmt.Println("test timer")
			}()
		}
	}
}

func main() {
	timer() //该函数是死循环，如果同时还有其它操作的话，前面就必须加上go关键字，否则其他操作不会被执行。
}
