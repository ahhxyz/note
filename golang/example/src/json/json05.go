package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Page   int
	Fruits []string
}

func main() {
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
	fmt.Println(res.Page)
}
