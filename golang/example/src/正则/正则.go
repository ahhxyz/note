package main

import (
	"fmt"
	"regexp"
)

func main() {
	rlt := regexp.MatchString("/api/([0-9]+)", "/api/123")
	fmt.Printf(rlt)
}
