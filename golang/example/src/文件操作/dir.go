package main

import (
	//"flag"
	"fmt"
	"os"
	"path/filepath"
)

func getFilelist(path string) []string {
	var paths []string
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		//println(path)
		paths = append(paths, path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return paths
}

func main() {
	fmt.Printf("%s\n", getFilelist("dir"))
}
