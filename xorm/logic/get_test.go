package logic

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gunsluo/go-example/xorm/models"
)

func TestGetInspectorById(t *testing.T) {
	expect := &models.Inspector{
		Id:       1,
		Username: "luoji",
		Password: "luoji",
		Created:  time.Now(),
	}

	mockfn := func(mock sqlmock.Sqlmock) {
		mock.ExpectQuery(`SELECT (.+) FROM "inspector"`).
			WithArgs(expect.Id).
			WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password", "created"}).
				AddRow(expect.Id, expect.Username, expect.Password, expect.Created))
	}
	session := newMockSession(mockfn)

	inspector := &models.Inspector{}
	has, err := GetInspectorById(session, inspector, 1)
	if err != nil {
		t.Fatal(err)
	}

	if !has {
		t.Fatal("not found")
	}

	if inspector.Id != expect.Id {
		t.Fatal("failed GetInsinspectorById")
	}
}

func Test2GetInspectorById(t *testing.T) {
	expect := &models.Inspector{
		Id:       1,
		Username: "luoji",
		Password: "luoji",
		Created:  time.Now(),
	}

	session := newSession()

	inspector := &models.Inspector{}
	has, err := GetInspectorById(session, inspector, 1)
	if err != nil {
		t.Fatal(err)
	}

	if !has {
		t.Fatal("not found")
	}

	if inspector.Id != expect.Id {
		t.Fatal("failed GetInsinspectorById")
	}
}
