package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	//"time"
)

func main() {
	_, _ = ioutil.ReadFile("demo.txt")
	//println(string(contents))
	//ioutil.WriteFile("demo.txt", contents, 0x777)

	fInfo, err := os.Stat("demo.txt") //该函数不会更改文件的访问时间
	if err == nil {
		info := fInfo.Sys()
		v := reflect.ValueOf(info).Elem()
		Atim := v.FieldByName("Atim").Field(0).Int()
		Mtim := v.FieldByName("Mtim").Field(0).Int()
		Ctim := v.FieldByName("Ctim").Field(0).Int()
		fmt.Printf("最后访问时间：%s\n修改时间：%s\n创建时间：%s\n", Atim, Mtim, Ctim)
		/*		if Atim.CanSet() {
					Atim.Field(0).SetInt(time.Now().Unix())
				}

				fmt.Printf("最后访问时间：%s\n", Atim.Field(0).Int())
		*/
	}
}
