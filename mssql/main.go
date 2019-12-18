package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

const (
	driver = "sqlserver"
	dsn    = "sqlserver://SA:Tes9ting@localhost:1433/instance?database=xo"
)

func main() {
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		panic(err)
	}

	test(db)

	db.Close()
}

func test(db *sqlx.DB) {
	stmt, err := db.Prepare("SELECT [id], [subject] FROM [USER]")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	var id int64
	var subject string
	err = row.Scan(&id, &subject)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows")
			return
		}
		panic(err)
	}

	fmt.Printf("user id:%d subject:%s\n", id, subject)

}
