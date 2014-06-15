package main

import (
	"fmt"
	"reflect"
)

func main() {

	Pre := [][]interface{}{
		[]interface{}{"uname", "yhm", "unique", "用户名已经存在！", "_", "_"}, //{数据库字段名，表单字段名，内置验证条件，验证失败的错误提示信息，附加验证方法，自动完成}
		[]interface{}{"nick", "nich", "require", "昵称不能为空！", "_", "_"},
		[]interface{}{"remark", "beizhu", "_", "_", "_", ""},
		[]interface{}{"level", "_", "_", "_", "_", 1},
		[]interface{}{"create_time", "_", "_", "_", "_", ""},
		//[]interface{}{"_", "isSubmited", "_", "_", "_", Lib.SetTime},
	}
	for _, val := range Pre {
		v := reflect.ValueOf(val[0])
		fmt.Printf("%v\n", v)

	}
}
