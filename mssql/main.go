package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

const (
	driver = "mssql"
	dsn    = "sqlserver://SA:Tes9ting@localhost:1433/instance?database=xo&encrypt=disable"
)

func main() {
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		panic(err)
	}

	test(db)
	test2(db)

	db.Close()
}

func test(db *sqlx.DB) {
	//OUTPUT INSERTED.id

	const sqlstr = `INSERT INTO dbo.[user] (` +
		`subject` +
		`) OUTPUT INSERTED.id VALUES (` +
		`$1` +
		`)`

	var id int64
	err := db.QueryRow(sqlstr, "luoji").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("-->", id)
}

func test2(db *sqlx.DB) {
	stmt, err := db.Prepare("SELECT [id], [subject] FROM [USER]")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var id int64
	var subject string
	for rows.Next() {
		err = rows.Scan(&id, &subject)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("no rows")
				return
			}
			panic(err)
		}

		fmt.Printf("user id:%d subject:%s\n", id, subject)
	}
}
