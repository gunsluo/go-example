package main

import (
	"reflect"
	"testing"

	"github.com/globalsign/mgo/bson"
	"github.com/golang/mock/gomock"
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
