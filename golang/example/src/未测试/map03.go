package main

import "fmt"

var m map[string]string

func main() {
    m = make(map[string]string)
    m["name01"] = ( "陈汤")
    m["name02"] = ( "甘延寿")
    fmt.Println(m["name01"])
    fmt.Println(m)
}