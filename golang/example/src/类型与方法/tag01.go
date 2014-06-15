package main

import (
	"fmt"
	"reflect"
)

var test int

type Person struct {
	Username string `json:"name"`              // 显式指定名称
	Age      int    `json:"age,string"`        // 显式指定 JSON 数据类型
	Address  string `json:"address,omitempty"` // 如果值为空，则忽略该字段。
	Data     []byte `json:"-" test`            // 显式忽略的字段
}

func main() {
	p := &Person{"Tom", 15, "", []byte{1, 2}}
	v := reflect.ValueOf(p).Elem().Type()
	for i := 0; i < v.NumField(); i++ {
		fmt.Println("%v:%v\n", v.Field(i).Name, v.Field(i).Tag) //将tag输出出来
	}
}
