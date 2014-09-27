package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

func main() {
	time1 := time.Now().UnixNano()

	db, _ := sql.Open("postgres", "user=postgres password=123456 dbname=cims sslmode=disable")

	time2 := time.Now().UnixNano()

	_sql := "SELECT id FROM cims_test"
	stmt, _ := db.Prepare(_sql)

	time3 := time.Now().UnixNano()
	res, _ := stmt.Exec()
	num, _ := res.RowsAffected()
	//= stmt.RowsAffected()
	//_, _ = stmt.Query()

	time4 := time.Now().UnixNano()
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n", time1, time2, time3, time4, num)
	db.Close()
}
