package main

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
	"github.com/xo/dburl"
)

func main() {
	//dsn := `oracle://ac/password@127.0.0.1:1521/ORCLPDB1`
	dsn := `oracle://ac:password@127.0.0.1:1521/ORCLPDB1`
	u, err := dburl.Parse(dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.Driver, u.DSN)

	db, err := sql.Open(u.Driver, u.DSN)
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
