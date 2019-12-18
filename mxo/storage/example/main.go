package main

import (
	"github.com/gunsluo/go-example/mxo/storage/sqlserver"
	"github.com/jmoiron/sqlx"

	_ "github.com/denisenkom/go-mssqldb"
)

const (
	//driver = "postgres"
	//dsn    = "postgres://postgres:password@localhost:5432/xo?sslmode=disable"

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
	user := &sqlserver.User{
		Subject: "luoji",
	}

	err := user.Insert(db)
	if err != nil {
		panic(err)
	}
}
