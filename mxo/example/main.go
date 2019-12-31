package main

import (
	"database/sql"
	"fmt"
	"os"
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
	var mod string
	if len(os.Args) <= 1 {
		mod = "postgres"
	} else {
		mod = os.Args[1]
	}

	fmt.Println("run mod:", mod)
	switch mod {
	case "postgres":
		testStoage(driverPostgres, dsnPostgres)
	case "mssql":
		testStoage(driverMssql, dsnMssql)
	default:
		fmt.Println("invalid parameter, it should be 'postgres' & 'mssql'")
	}
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

	user.Name = sql.NullString{Valid: true, String: "luoji"}
	err = s.UpdateUser(db, user)
	if err != nil {
		panic(err)
	}

	user.Name = sql.NullString{Valid: true, String: "luoji2"}
	err = s.UpsertUser(db, user)
	if err != nil {
		panic(err)
	}

	// query
	a, err := s.AccountByID(db, account.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("query account: %v\n", a)

	a, err = s.AccountBySubject(db, account.Subject)
	if err != nil {
		panic(err)
	}
	fmt.Printf("query account by subject: %v\n", a)

	a, err = s.AccountInUser(db, user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("query account in user: %v\n", a)

	as, err := s.GetMostRecentAccount(db, 10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("query account in user: %d %v\n", len(as), as)

	as, err = s.GetMostRecentChangedAccount(db, 10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("query account in user: %d %v\n", len(as), as)

	as, err = s.GetAllAccount(db, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("query all account in user: %d %v\n", len(as), as)

	count, err := s.CountAllAccount(db, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("count all account in user: %d\n", count)

	u, err := s.UserByID(db, user.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("query user: %v\n", u)

	us, err := s.UsersBySubjectFK(db, account.Subject, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("query user by subject fk: %s %v\n", account.Subject, len(us))

	total, err := s.CountUsersBySubjectFK(db, account.Subject, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("total user by subject fk: %s %v\n", account.Subject, total)

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
}
