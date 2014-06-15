package main

import "fmt"

func main() {
	m := map[string]interface{}{
		"person01": []string{"id", "name"},
		"person02": "甘延寿",
		"info": map[string]string{
			"朝代": "大汉王朝",
			"民族": "汉族",
		},
	}
	for i, v := range m {
		fmt.Println("index:%s\nvalue:%s\n", i, v)
	}
}
