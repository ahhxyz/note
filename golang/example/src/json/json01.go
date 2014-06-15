package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// json encode
	j1 := make(map[string]interface{})
	j1["name"] = "豆蔻"
	j1["url"] = "http://www.dotcoo.com/"

	js1, err := json.Marshal(j1)
	if err != nil {
		panic(err)
	}

	println(string(js1))

	// json decode
	j2 := make(map[string]interface{})
	err = json.Unmarshal(js1, &j2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", j2)
}
