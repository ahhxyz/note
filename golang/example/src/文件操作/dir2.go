package main

import (
	//"flag"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	//"time"
)

func getFilelist(path string) {
	///var paths []string
	err := filepath.Walk(path, func(path string, fInfo os.FileInfo, err error) error {
		if fInfo == nil {
			return err
		}
		if !fInfo.IsDir() {
			//fInfo, _ := os.Stat(S.file) //该函数不会更改文件的访问时间
			info := fInfo.Sys()
			v := reflect.ValueOf(info).Elem()
			Atim := v.FieldByName("Atim").Field(0).Int()
			fmt.Printf("%s\n", Atim)

			//return fInfo
		}
		//println(fInfo.Sys())
		//paths = append(paths, path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	//return paths
}

func main() {
	getFilelist("dir")
}
