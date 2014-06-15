package main

import (
	"fmt"
	"time"
)

func main() {
	//const layout = "2,1, 2006"
	//t := time.Date(2009, time.November, 10, 15, 0, 0, 0, time.Local)
	fmt.Println(time.Now().Format("2006-01-02") + "-" + fmt.Sprintf("%d", time.Now().Hour()))
	//fmt.Println(t.UTC().Format(layout))
}
