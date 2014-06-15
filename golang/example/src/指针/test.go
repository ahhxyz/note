package main

import (
	"fmt"
)

type longtou map[int]string

func main() {
	va := &longtou{}
	for i := 0; i < 100; i++ {
		(*va)[i] = "aaaa"
	}
	for k, v := range *va {
		fmt.Println(k, v)
	}

}
