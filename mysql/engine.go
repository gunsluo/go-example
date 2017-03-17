package main

import (
	"runtime"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gunsluo/xorm"
	//"time"
)

var lock sync.Mutex

type User struct {
	UserID            int       `xorm:"user_id,autoincr pk" json:"userId"`           //用户id
	UserName          string    `xorm:"user_name" json:"userName"`                   //用户名字
	DistinguishedName string    `xorm:"distinguished_name" json:"distinguishedName"` //标识名
	Samaccountname    string    `xorm:"samaccount_name" json:"samaccountname"`       //标识名
	Mail              string    `xorm:"mail" json:"mail"`                            //邮件
	Description       string    `xorm:"description" json:"description"`              //描述
	Role              int       `xorm:"role" json:"role"`                            //1为admin 0为普通用户
	Deleted           int       `xorm:"Deleted" json:"deleted"`                      //用户状态（0表示正常，1表示已删除，2表示禁用）
	CreatedTime       string    `xorm:"created" json:"createdTime"`                  //创建时间
	UpdatedTime       time.Time `xorm:"updated" json:"updatedTime"`                  //更新时间
	DeletedTime       time.Time `xorm:"deleted_time" json:"deletedTime"`             //删除时间
}

func (user *User) TableName() string {
	return "t_user"
}

func test(db *xorm.Engine) {
	//查询数据
	var users []User
	lock.Lock()
	err := db.Sql("SELECT user_id, user_name, distinguished_name, samaccount_name, mail, created_time FROM t_user").Find(&users)
	lock.Unlock()
	checkErr(err)
}

func main() {
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	var count int = 200

	db, err := xorm.NewEngine("mysql", "root:luoji@tcp(192.168.0.7:3306)/example?charset=utf8&loc=Asia%2FShanghai")
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
