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
	account := &storage.Account{
		Subject:     "luoji",
		CreatedDate: storage.NullTime{Time: time.Now(), Valid: true},
		ChangedDate: storage.NullTime{Time: time.Now(), Valid: true},
		//DeletedDate: time.Now(),
	}

	err := s.InsertAccount(db, account)
	if err != nil {
		panic(err)
	}
	fmt.Printf("insert account: id=%d, err=%v\n", account.ID, err)

	account.Email = "luoji@gmail.com"
	err = s.UpdateAccount(db, account)
	if err != nil {
		panic(err)
	}

	account.Subject = "luoji1"
	err = s.UpsertAccount(db, account)
	if err != nil {
		panic(err)
	}

	user := &storage.User{
		Subject:     account.Subject,
		CreatedDate: storage.NullTime{Time: time.Now(), Valid: true},
		ChangedDate: storage.NullTime{Time: time.Now(), Valid: true},
	}

	err = s.InsertUser(db, user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("insert user: id=%d, err=%v\n", user.ID, err)

	user.Name = "luoji"
	err = s.UpdateUser(db, user)
	if err != nil {
		panic(err)
	}

	user.Name = "luoji2"
	err = s.UpsertUser(db, user)
	if err != nil {
		panic(err)
	}

	err = s.DeleteUser(db, user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("delete user: %v\n", err)

	err = s.DeleteAccount(db, account)
	if err != nil {
		panic(err)
	}
	fmt.Printf("delete account: %v\n", err)

	/*
		u, err := s.AccountByID(db, account.ID)
		if err != nil {
			panic(err)
		}

		fmt.Printf("query account: %v\n", u)
	*/

}
