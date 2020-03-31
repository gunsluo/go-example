package main

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

func main() {
	db, err := sql.Open("godror", `c##admin/password@127.0.0.1:1521/ORCLCDB`)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	var user string
	err = db.QueryRow("select user from dual").Scan(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Successful 'as sysdba' connection. Current user is: %v\n", user)
}
