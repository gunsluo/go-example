package main

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type TZDemo struct {
	ID          int64       `json:"id" db:"id"`                     // id
	Tz          time.Time   `json:"tz" db:"tz"`                     // tz
	CreatedDate pq.NullTime `json:"created_date" db:"created_date"` // created_date
	ChangedDate pq.NullTime `json:"changed_date" db:"changed_date"` // changed_date
	DeletedDate pq.NullTime `json:"deleted_date" db:"deleted_date"` // deleted_date
}

const (
	driver = "postgres"
	dsn    = "postgres://postgres:password@postgres:5432/postgres?sslmode=disable"
)

func main() {
	now := time.Now()
	fmt.Println("now:", now)

	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		// Bad DSN, we quit immediately
		panic(err)
	}

	d := TZDemo{Tz: now}
	var sqlstr = `INSERT INTO tzdemo(tz) VALUES($1)`
	_, err = db.Exec(sqlstr, d.Tz)
	if err != nil {
		panic(err)
	}

	sqlstr = "select id, tz, created_date, changed_date, deleted_date from tzdemo"
	q, err := db.Query(sqlstr)
	if err != nil {
		panic(err)
	}
	defer q.Close()

	for q.Next() {
		dd := TZDemo{}

		// scan
		err = q.Scan(&dd.ID, &dd.Tz, &dd.CreatedDate, &dd.ChangedDate, &dd.DeletedDate)
		if err != nil {
			panic(err)
		}

		fmt.Println("--->", dd)
	}

	fmt.Println("Ok")
}

/*
CREATE TABLE IF NOT EXISTS tzdemo (
    id bigserial NOT NULL,

    tz timestamp NOT NULL,
    created_date timestamp DEFAULT now(),
    changed_date timestamp DEFAULT now(),
    deleted_date timestamp,

    CONSTRAINT tzdemo_pk PRIMARY KEY (id)
);
*/
