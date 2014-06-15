package main

import "fmt"

type Db struct{
	Name string
}
type Model struct{
	*Db
	Data string		
}

type User struct{
	*Model
	Id int
	
}

func (u *User) Put(){
	
	u.Name="公共的Name"

	
}
func (u *User) Get(){
	fmt.Println("%v\n",u.Name)
}


func main() {
	u1:=&User{&Model{new(Db),"d1"},12}
	u2:=&User{&Model{new(Db),"d2"},13}
	u1.Put()
	u2.Get()//不会输出："公共的Name"，虽然这个字段的类型是以指针类型嵌入的，但u1和u2中，指针指向了类型的不同实例
	
}