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

	s, err := db.GetShareRelation(d, "filename1", "share")
	if err != nil {
		panic(err)
	}
	fmt.Println("---->", s.Beneficiaries)

	err = db.AddBeneficiaries(d, "filename1", "jerry", "share", []string{"luoji", "gunsluo", "abc"})
	if err != nil {
		panic(err)
	}

	err = db.RemoveBeneficiaries(d, "filename1", "share", []string{"luoji", "abc"})
	if err != nil {
		panic(err)
	}
}
