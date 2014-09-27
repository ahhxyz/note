package main

import (
	//"encoding/json"
	"fmt"
	"reflect"
)

var test int

type Person struct {
	Username string `json:"name"`                        // 显式指定名称
	Age      int    `json:"age,string"`                  // 显式指定 JSON 数据类型
	Address  string `json:"address,omitempty" orm:"orm"` // 如果值为空，则忽略该字段。
	Data     []byte `json`                               // 显式忽略的字段
}

func main() {
	p := &Person{"Tom", 15, "", []byte{1, 2}}
	v := reflect.ValueOf(p).Elem().Type()
	//var cdts string
	for i := 0; i < v.NumField(); i++ {
		get := v.Field(i).Tag
		//json.Unmarshal([]byte(v.Field(i).Tag.Get("json")), &cdts)
		fmt.Printf("%v\n", get == "")
		//fmt.Println("%v:%v\n", v.Field(i).Name, v.Field(i).Tag) //将tag输出出来
	}
}
