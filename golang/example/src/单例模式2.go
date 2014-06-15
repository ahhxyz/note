package main

import "fmt"

var _Db *Db

type Db struct {
	Name string
}

func (p *Db) Set(Name string) {
	p.Name = Name
}

func (p *Db) Say() {
	fmt.Println(p.Name)
}
func main() {
	if _Db == nil {
		_Db = new(Db)
	}
	_Db.Set("甘延寿")
	_Db.Say()
}
