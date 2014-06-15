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
		rf := reflect.ValueOf(fields)
		switch reflect.TypeOf(fields).Kind() {
		case reflect.Slice:

			for i := 0; i < rf.Len(); i++ {
				fmt.Println("%s\n", rf.Index(i))
			}

		case reflect.Map:
			//allKeys := rf.MapKeys()
			for index, val := range fields.(map[string]string) {
				fmt.Println("%s\n%s\n%s\n", index, val, fields.(map[string]string))
			}
		}

	}
}
