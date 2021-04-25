package logic

import (
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

func newEngine(mockfns ...func(sqlmock.Sqlmock)) *xorm.Engine {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	engine, _ := xorm.NewEngine("postgres", "postgres://postgres:password@127.0.0.1:5432/xorm?sslmode=disable")
	engine.DB().DB = db
	engine.ShowSQL(true)
	for _, fn := range mockfns {
		fn(mock)
	}

	return engine
}

func newSession(mockfns ...func(sqlmock.Sqlmock)) *xorm.Session {
	engine := newEngine(mockfns...)
	return engine.NewSession()
}
