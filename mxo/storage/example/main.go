package main

import (
	"github.com/gunsluo/go-example/mxo/storage/postgres"
	"github.com/jmoiron/sqlx"
)

const (
	driver = "postgres"
	dsn    = "postgres://postgres:password@localhost:5432/xo?sslmode=disable"
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
	user := &postgres.User{
		Subject: "luoji",
	}

	err := user.Insert(db)
	if err != nil {
		panic(err)
	}
}
