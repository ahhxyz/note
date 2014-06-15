package main

import (
	"fmt"
)

var _Db *Db

type Db struct {
	Name string
	ctx  *Db
}

func _getDb() *Db {
	if _Db == nil {
		_Db = new(Db)
	}
	return _Db
}

func (p *Db) Set(Name string) {
	p.Name = Name
}

func (p *Db) Say() {
	fmt.Println(p.Name)
}
func main() {
	_getDb()
	_Db.Set("甘延寿")
	_Db.Say()
}
