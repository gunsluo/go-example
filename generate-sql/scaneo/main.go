package main

import (
	"log"

	"github.com/gunsluo/go-example/generate-sql/scaneo/models"
	"github.com/jmoiron/sqlx"
)

const (
	dsn = "dbname=luoji host=localhost port=26257 sslmode=disable user=root"
)

func main() {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	rows, err := db.Query("select * from post")
	if err != nil {
		log.Println(err)
		return
	}

	posts, err := models.ScanPosts(rows) // ScanPosts was auto-generated!
	if err != nil {
		log.Println(err)
	}

	log.Println(posts)
}
