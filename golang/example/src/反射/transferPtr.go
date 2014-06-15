/*
 *本例子测试将指针传入函数后，指针内容改变后，是否能反映到其它地方。
 */

package main

import (
	"fmt"
	"reflect"
)

type Orm struct {
	Ctx interface{}
}

func NewOrm(m interface{}) *Orm {
	o := new(Orm)
	o.Ctx = m
	return o
}
func (o *Orm) Get() {
	v := reflect.ValueOf(o.Ctx).Elem().FieldByName("title").String()
	fmt.Printf("%v\n", v)

}

type User struct {
	title string
}

func main() {
	u := new(User)
	u.title = "原标题"
	o := NewOrm(u)
	o.Get()
	u.title = "新标题"
	o.Get()

}
