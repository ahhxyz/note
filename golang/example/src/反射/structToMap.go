package main

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

type Person struct {
	Name string
	Age  uint
}

func main() {
	betty := Person{
		Name: "Betty",
		Age:  82,
	}

	urlValues := structToMap(reflect.ValueOf(&betty))
	fmt.Printf("%#v", urlValues["Age"][0])

}

func structToMap(v reflect.Value) url.Values {
	values := url.Values{}
	iVal := v.Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)

		var v string
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(f.Bytes())
		case string:
			v = f.String()
		}
		values.Set(typ.Field(i).Name, v)
	}
	/*
		structFields := make(map[string]interface{}, 0)
		for fieldName, valArr := range values {
			structFields[fieldName] = valArr[0]
		}
	*/
	return values
}
