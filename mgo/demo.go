package main

import (
	"fmt"

	"github.com/globalsign/mgo"
	"gitlab.com/target-digital-transformation/file-manager/db"
)

const (
	mongoURL = "mongodb://root:password@127.0.0.1:27017"
)

func main() {
	session, err := mgo.Dial(mongoURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	d := session.DB("test")
	err = db.CreateIndexes(d)
	if err != nil {
		panic(err)
	}

	s, err := db.StatisInfo(d, "jerry")
	if err != nil {
		panic(err)
	}
	fmt.Println("---->", s.Storage)

	err = db.IncrStatisStorage(d, "jerry", 10)
	if err != nil {
		panic(err)
	}
}
