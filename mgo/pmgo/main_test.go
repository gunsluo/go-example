package main

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/globalsign/mgo/bson"
	"github.com/golang/mock/gomock"
	"github.com/percona/pmgo"
	"github.com/percona/pmgo/pmgomock"
)

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := User{
		ID:   1,
		Name: "Zapp Brannigan",
	}

	// Mock up a database, session, collection and a query and set
	// expected/returned values for each type
	query := pmgomock.NewMockQueryManager(ctrl)
	query.EXPECT().One(gomock.Any()).SetArg(0, user).Return(nil)

	collection := pmgomock.NewMockCollectionManager(ctrl)
	collection.EXPECT().Find(bson.M{"id": 1}).Return(query)

	database := pmgomock.NewMockDatabaseManager(ctrl)
	database.EXPECT().C("testc").Return(collection)

	session := pmgomock.NewMockSessionManager(ctrl)
	session.EXPECT().DB("test").Return(database)

	// Call the function we want to test. It will use the mocked interfaces
	readUser, err := getUser(session, 1)

	if err != nil {
		t.Errorf("getUser returned an error: %s\n", err.Error())
	}

	if !reflect.DeepEqual(*readUser, user) {
		t.Errorf("Users don't match. Got %+v, want %+v\n", readUser, user)
	}
}

var Server pmgo.DBTestServer

func TestIntegration(t *testing.T) {
	setup()

	readUser, err := getUser(Server.Session(), 1)
	if err != nil {
		t.Errorf("getUser returned an error: %s\n", err.Error())
	}

	if !reflect.DeepEqual(*readUser, mockUser()) {
		t.Errorf("Users don't match. Got %+v, want %+v\n", readUser, mockUser())
	}

	tearDown()
}

func setup() {
	os.Setenv("CHECK_SESSIONS", "0")
	tempDir, err := ioutil.TempDir("", "testing")
	if err != nil {
		panic(err)
	}

	Server = pmgo.NewDBServer()
	Server.SetPath(tempDir)

	session := Server.Session()
	// load some fake data into the db
	session.DB("test").C("testc").Insert(mockUser())
}

func mockUser() User {
	return User{
		ID:   1,
		Name: "Zapp Brannigan",
	}

}

func tearDown() {
	Server.Session().Close()
	Server.Session().DB("samples").DropDatabase()
	Server.Stop()
}
