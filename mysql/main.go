package main

import (
	"database/sql"
	"runtime"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	//"time"
)

var lock sync.Mutex

func test(db *sql.DB) {
	//查询数据
	lock.Lock()
	rows, err := db.Query("SELECT user_id, user_name, distinguished_name, samaccount_name, mail, created_time FROM t_user")
	lock.Unlock()
	checkErr(err)

	for rows.Next() {
		var userID int
		var username string
		var distinguishedName string
		var samaccountName string
		var mail string
		var created string
		err = rows.Scan(&userID, &username, &distinguishedName, &samaccountName, &mail, &created)
		checkErr(err)
		//fmt.Println(userID, username, distinguishedName, samaccountName, mail, created)
	}

	rows.Close()
}

func main() {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	var count int = 200

	db, err := sql.Open("mysql", "root:luoji@tcp(192.168.0.7:3306)/example?charset=utf8")
	checkErr(err)
	db.SetMaxIdleConns(1000)
	db.SetMaxOpenConns(1000)

	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			test(db)
		}()
	}

	wg.Wait()
	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
