package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	str := `{"num":6.13,"strs":{"a":"b"}}`

	// json decode
	j := make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &j)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", j, j["strs"]["a"])
}
