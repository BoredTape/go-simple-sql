package main

import (
	"github.com/BoredTape/go-simple-sql"
	"fmt"
)

func main() {
	var db go_simple_sql.CONN
	err := db.InitDB("127.0.0.1", "3306", "root", "root", "database", "utf8")
	if err != nil {
		fmt.Println(err)
	}
	sql := "SELECT * FROM table WHERE id=1"
	result, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
