package main

import "time"

func t(c <-chan time.Time) {
	<-c
	println("Timer expired")

}
func main() {

	go t(time.NewTimer(time.Second * 2).C)

}
