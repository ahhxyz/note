package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	time1 := time.Now().UnixNano()
	var id int64
	db, _ := sql.Open("mysql", "root:123456@tcp(localhost:3307)/cimser?charset=utf8")

	time2 := time.Now().UnixNano()

	_sql := "INSERT INTO cims_user(uname,pword,remark,level,status,create_time,nick,allow_db_counts,update_time) VALUES('吴鑫','e10adc3949ba59abbe56e057f20f883e','备注','1','1','1405151495','','5000','1405151495')"
	stmt, err := db.Prepare(_sql)

	time3 := time.Now().UnixNano()
	res, err := stmt.Exec()
	if err != nil {
		panic(err)
	}
	id, err = res.LastInsertId()
	fmt.Printf("%v\n", id, err)
	//num, _ := res.RowsAffected()
	//= stmt.RowsAffected()
	//_, _ = stmt.Query()

	time4 := time.Now().UnixNano()
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n", time1, time2, time3, time4)
	db.Close()
}
