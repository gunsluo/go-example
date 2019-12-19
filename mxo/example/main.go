package main

import (
	"fmt"
	"time"

	"github.com/gunsluo/go-example/mxo/storage"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
)

const (
	driverPostgres = "postgres"
	dsnPostgres    = "postgres://postgres:password@localhost:5432/xo?sslmode=disable"

	driverMssql = "mssql"
	dsnMssql    = "sqlserver://SA:Tes9ting@localhost:1433/instance?database=xo&encrypt=disable"
)

func main() {
	//testStoage(driverPostgres, dsnPostgres)
	testStoage(driverMssql, dsnMssql)
}

func testStoage(driver, dsn string) {
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		panic(err)
	}

	s, err := storage.New(driver, storage.Config{Logger: logrus.New()})
	if err != nil {
		panic(err)
	}

	testStoageAndDB(s, db)

	db.Close()
}

func testStoageAndDB(s storage.Storage, db *sqlx.DB) {
	user := &storage.User{
		Subject:     "luoji",
		CreatedDate: time.Now(),
		ChangedDate: time.Now(),
		DeletedDate: time.Now(),
	}

	err := s.InsertUser(db, user)
	if err != nil {
		panic(err)
	}

	u, err := s.UserByID(db, user.ID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("query user: %v\n", u)
}
