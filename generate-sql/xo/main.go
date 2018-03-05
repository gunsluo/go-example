package main

import (
	"fmt"

	"github.com/gunsluo/go-example/xo/models"
	"github.com/jmoiron/sqlx"
	"github.com/knq/dburl"
	_ "github.com/lib/pq"
)

const (
	dsn = "pgsql://root@localhost:26257/luoji?sslmode=disable"
)

// xo -v "pgsql://root@localhost:26257/?sslmode=disable" -s luoji -o models --template-path templates
func main() {

	url, err := dburl.Parse(dsn)
	if err != nil {
		panic(err)
	}

	db, err := sqlx.Open(url.Driver, url.DSN)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	account := &models.Account{}
	account.ID = 7
	account.Balance = models.Float64(100.01)

	if err := account.Save(db); err != nil {
		panic(err)
	}

	fmt.Println("->", account.Exists())
}
