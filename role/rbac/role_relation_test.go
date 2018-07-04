package rbac

import (
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestRoleRelation(t *testing.T) {
	db := createDB(t, func(mock sqlmock.Sqlmock) error {
		return nil
	})

	relation := newRoleRelation(db)
	if relation == nil {
		t.Fatalf("Failed: role relation is nil ")
	}
	relation.buildRoleLinks()

	err := relation.AddRoleForUser("user 1", "member")
	if err != nil {
		t.Fatalf("Failed to add role for user: %s", err)
	}

	err = relation.AddRoleForUser("user 2", "member")
	if err != nil {
		t.Fatalf("Failed to add role for user: %s", err)
	}

	err = relation.AddRoleForUser("user 3", "admin")
	if err != nil {
		t.Fatalf("Failed to add role for user: %s", err)
	}

	/*
		err = relation.AddRoleForUser("user 1", "member")
		if err != nil {
			t.Fatalf("Failed to add role for user: %s", err)
		}
	*/
}
