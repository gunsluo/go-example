package main

import (
	"fmt"
	"log"

	"github.com/globalsign/mgo/bson"
	"github.com/gunsluo/pmgo"
)

// User is  a user information
type User struct {
	ID   int    `bson:"id"`
	Name string `bson:"name"`
}

func main() {
	dialer := pmgo.NewDialer()
	session, err := dialer.Dial("mongodb://root:password@localhost:27017")
	if err != nil {
		print(err)
		return
	}

	if err := addUser(session, &User{ID: 1, Name: "jerry"}); err != nil {
		log.Printf("error add the user to the db: %s", err.Error())
		return
	}

	user, err := getUser(session, 1)
	if err != nil {
		log.Printf("error reading the user from the db: %s", err.Error())
		return
	}
	fmt.Printf("User: %+v\n", user)
}

func addUser(session pmgo.SessionManager, user *User) error {
	return session.DB("test").C("testc").Insert(user)
}

func getUser(session pmgo.SessionManager, id int) (*User, error) {
	var user User
	err := session.DB("test").C("testc").Find(bson.M{"id": id}).One(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
