package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if fi, err := os.Stat("pwd.go"); err == nil {
		//Atim := reflect.ValueOf(fi.Sys()).Elem().FieldByName("").Field(0).Int()
		//fmt.Printf("%v\n", fi.ModTime().Unix())

		if fi.ModTime().Unix()+3600 < time.Now().Unix() {
			fmt.Printf("%v\n", fi.ModTime().Unix())
		}

	}
}
