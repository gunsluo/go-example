package main

import (
	"time"

	"github.com/gunsluo/go-example/mxo/storage/mssql"
	"github.com/jmoiron/sqlx"

	_ "github.com/denisenkom/go-mssqldb"
)

const (
	//driver = "postgres"
	//dsn    = "postgres://postgres:password@localhost:5432/xo?sslmode=disable"

	driver = "mssql"
	dsn    = "sqlserver://SA:Tes9ting@localhost:1433/instance?database=xo&encrypt=disable"
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
	user := &mssql.User{
		Subject:     "luoji",
		CreatedDate: time.Now(),
		ChangedDate: time.Now(),
		DeletedDate: time.Now(),
	}

	err := user.Insert(db)
	if err != nil {
		panic(err)
	}
}
