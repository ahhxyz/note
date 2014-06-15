package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
}
