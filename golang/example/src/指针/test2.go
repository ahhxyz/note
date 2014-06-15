package main

import (
	"fmt"
)

type longtou map[string]string

func main() {
	va := &longtou{}
	(*va)["baidu"] = "bbbb"
	(*va)["google"] = "cccc"

	for k, v := range *va {
		fmt.Println(k, v)
	}

}
