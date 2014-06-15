package main
import(
   "fmt"
)
/*
@以下代码的语法不保证正确性，熟悉语法用的
*/
type human struct{
   name string
   age int
}
type woman struct{
	human //匿名字段用来实现继承
	isBeauty bool
}
type man struct{
   human //匿名字段用来实现继承
   isStrong bool

}
type all interface{
	getName() string
}
func (h *human) getName()string{ //human类实现了all接口
 return h.name
}

func main() {
 var w woman 
 w.name="刘亦菲"
 w.age=26
 w.isBeauty=true
 fmt.Println(w)
 fmt.Println(w.getName())//这个方法来自于继承。
 /*
 @输出结果：
 {{刘亦菲 26} true}
 刘亦菲
 */
}