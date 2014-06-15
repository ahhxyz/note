package main

import "fmt"
import "reflect"

var (
	which  string
	whichs []string
	joins  []string
	join   string
)

type User struct {
	View map[string]interface{}
}

func test(t interface{}) {
}
func main() {

	m := new(User) //这个实例必须也是对应控制器用要使用的实例

	m.View = map[string]interface{}{
		"User": []string{"id", "name"},
		"Department": map[string]string{
			"did":   "did",
			"dname": "dname",
			"_on":   "Department.id=User.id",
		},
	}

	for _, fields := range m.View {
		//rf := reflect.ValueOf(fields)
		switch reflect.TypeOf(fields).Kind() {

		case reflect.Map:
			fmt.Printf("%s\n", fields.(map[string]string))
		}

	}
}
