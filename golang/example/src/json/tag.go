package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Username string `json:"name"`              // 显式指定名称
	Age      int    `json:"age,string"`        // 显式指定 JSON 数据类型
	Address  string `json:"address,omitempty"` // 如果值为空，则忽略该字段。
	Data     []byte `json:"-"`                 // 显式忽略的字段
}

func main() {
	p := &Person{"Tom", 15, "", []byte{1, 2}}
	b, _ := json.MarshalIndent(p, "", "\t")
	fmt.Println(string(b))
	//p2 := new(Person)
	//json.Unmarshal(b, p2)
}
