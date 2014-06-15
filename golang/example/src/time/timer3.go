package main

import (
	"fmt"
	"time"
)

func testTimer() {
	go func() {
		fmt.Println("test timer")
	}()

}

func timer() {
	timer := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timer.C:
			testTimer()
		}
	}
}

func main() {
	go timer()
	timer()
}
