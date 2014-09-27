package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	SESSION := make(map[string]interface{}, 0)
	SESSION["name"] = "名字"
	app := make(map[string]string)
	app["name"] = "应用名称"
	SESSION["APP"] = app
	delete(SESSION, "APP")
	str, _ := json.Marshal(SESSION)
	_ = json.Unmarshal([]byte(string(str)), &SESSION)

	fmt.Printf("%v\n", SESSION)
}
