package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:luoji@tcp(192.168.0.7:3306)/cloudzone?charset=utf8")
	checkErr(err)
	//db.SetMaxIdleConns(1000)
	//db.SetMaxOpenConns(1000)

	name := "\\aaa"
	rows, err := db.Query(`select app_meta_id,app_meta_name from app_meta where app_meta_name like ?`, name)
	checkErr(err)

	for rows.Next() {
		var appMetaID int
		var appMetaName string
		err = rows.Scan(&appMetaID, &appMetaName)
		checkErr(err)
		fmt.Println(appMetaID, appMetaName)
	}

	rows.Close()

	/*
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}
		defer tx.Rollback()
		stmt, err := tx.Prepare("INSERT INTO foo VALUES (?)")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close() // danger!
		for i := 0; i < 10; i++ {
			_, err = stmt.Exec(i)
			if err != nil {
				log.Fatal(err)
			}
		}
		err = tx.Commit()
		if err != nil {
			log.Fatal(err)
		}
	*/
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
