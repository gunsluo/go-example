package logic

import (
	"testing"
	"time"

	"github.com/gunsluo/go-example/xorm/models"
)

func TestGetInspectorById(t *testing.T) {
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
