package rbac

import (
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestRoleStore(t *testing.T) {
	db := createDB(t, func(mock sqlmock.Sqlmock) error {
		mock.ExpectQuery("INSERT INTO role").
			WithArgs("qs:admin", "test").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		mock.ExpectBegin()
		mock.ExpectQuery("INSERT INTO role").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectQuery("INSERT INTO role").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))
		mock.ExpectCommit()

		mock.ExpectQuery("^SELECT (.+) FROM role WHERE name=?").
			WillReturnRows(sqlmock.NewRows([]string{"total"}).AddRow(1))

		mock.ExpectQuery("^SELECT (.+) FROM role WHERE name = ?").
			WillReturnRows(sqlmock.NewRows([]string{"id", "description"}).AddRow(1, 2))

		mock.ExpectQuery("^SELECT (.+) FROM role limit (.+)").
			WithArgs(10, 0).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description"}).
				AddRow(1, "qs:admin", "test").
				AddRow(2, "qs:member", "test").
				AddRow(3, "qs:anonymous", "This is a system reservation role, match all role."))

		mock.ExpectQuery("^SELECT (.+) FROM role (.+)").
			WithArgs("%qs%", 10, 0).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description"}).
				AddRow(1, "qs:admin", "test").
				AddRow(2, "qs:member", "test").
				AddRow(3, "qs:anonymous", "This is a system reservation role, match all role."))

		mock.ExpectQuery("^SELECT (.+) FROM role").
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(3))

		mock.ExpectQuery("^SELECT (.+) FROM role").
			WithArgs("%qs%").
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(3))

		mock.ExpectExec("DELETE FROM role").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("DELETE FROM role").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("DELETE FROM role").WillReturnResult(sqlmock.NewResult(1, 1))

		return nil
	})

	store := newRoleStore()
	err := store.Insert(db, &Role{Name: "qs:admin", Description: "test"})
	if err != nil {
		t.Fatalf("Failed to add role: %s", err)
	}

	tx, err := db.Beginx()
	if err != nil {
		t.Fatalf("Failed to begin transaction: %s", err)
	}

	err = store.Insert(tx, &Role{Name: "qs:member", Description: "test"})
	if err != nil {
		t.Fatalf("Failed to add role: %s", err)
	}

	err = store.Insert(tx, &Role{Name: "qs:anonymous", Description: "test"})
	if err != nil {
		t.Fatalf("Failed to add role: %s", err)
	}

	if err = tx.Commit(); err != nil {
		t.Fatalf("Failed to commit transaction: %s", err)
	}

	exist, err := store.Exist(db, "qs:admin")
	if err != nil || !exist {
		t.Fatalf("Failed to query role: %s", err)
	}

	role, err := store.Get(db, "qs:admin")
	if err != nil {
		t.Fatalf("Failed to get role: %s", err)
	}
	if role == nil {
		t.Fatalf("Failed to get role")
	}

	roles, err := store.GetAll(db, 10, 0)
	if err != nil {
		t.Fatalf("Failed to get all role: %s", err)
	}
	if len(roles) == 0 {
		t.Fatalf("Failed to get all role")
	}

	roles, err = store.GetAll(db, 10, 0, "qs")
	if err != nil {
		t.Fatalf("Failed to get all role: %s", err)
	}
	if len(roles) != 3 {
		t.Fatalf("Failed to get all role")
	}

	count, err := store.Count(db)
	if err != nil {
		t.Fatalf("Failed to count role: %s", err)
	}
	if count == 0 {
		t.Fatalf("Failed to count role")
	}

	count, err = store.Count(db, "qs")
	if err != nil {
		t.Fatalf("Failed to count role: %s", err)
	}
	if count != 3 {
		t.Fatalf("Failed to count role")
	}

	err = store.Delete(db, "qs:admin")
	if err != nil {
		t.Fatalf("Failed to remove role: %s", err)
	}

	err = store.Delete(db, "qs:member")
	if err != nil {
		t.Fatalf("Failed to remove role: %s", err)
	}

	err = store.Delete(db, "qs:anonymous")
	if err != nil {
		t.Fatalf("Failed to remove role: %s", err)
	}

	db.Close()
}
