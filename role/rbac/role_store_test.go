package rbac

import (
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestRoleStore(t *testing.T) {
	db := createDB(t, func(mock sqlmock.Sqlmock) error {
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO role").
			WithArgs("qs:admin", "test").
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO role").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("INSERT INTO role").WillReturnResult(sqlmock.NewResult(2, 2))
		mock.ExpectCommit()

		mock.ExpectQuery("^select (.+) from role where name=?").
			WillReturnRows(sqlmock.NewRows([]string{"total"}).AddRow(1))

		mock.ExpectQuery("^select (.+) from role where name = ?").
			WillReturnRows(sqlmock.NewRows([]string{"id", "description"}).AddRow(1, 2))

		mock.ExpectQuery("^select (.+) from role where id = ?").
			WillReturnRows(sqlmock.NewRows([]string{"id", "description"}).AddRow(1, 2))

		mock.ExpectQuery("^select (.+) from role limit (.+)").
			WithArgs(10, 0).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description"}).
				AddRow(1, "qs:admin", "test").
				AddRow(2, "qs:member", "test").
				AddRow(3, "qs:anonymous", "This is a system reservation role, match all role."))

		mock.ExpectQuery("^select (.+) from role (.+)").
			WithArgs("%qs%", 10, 0).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "description"}).
				AddRow(1, "qs:admin", "test").
				AddRow(2, "qs:member", "test").
				AddRow(3, "qs:anonymous", "This is a system reservation role, match all role."))

		mock.ExpectQuery("^select (.+) from role").
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(3))

		mock.ExpectQuery("^select (.+) from role").
			WithArgs("%qs%").
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(3))

		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM role").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM role").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM role").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		return nil
	})
	defer db.Close()

	rs := newRoleStore(db)
	err := rs.Create("qs:admin", "test")
	if err != nil {
		t.Fatalf("Failed to add role: %s", err)
	}
	err = rs.BatchCreate([]*Role{
		&Role{
			Name:        "qs:member",
			Description: "test",
		},
		&Role{
			Name:        "qs:anonymous",
			Description: "This is a system reservation role, match all role.",
		},
	})
	if err != nil {
		t.Fatalf("Failed to add role: %s", err)
	}

	exist, err := rs.Exist("qs:admin")
	if err != nil || !exist {
		t.Fatalf("Failed to get role: %s", err)
	}

	role, err := rs.GetByName("qs:admin")
	if err != nil {
		t.Fatalf("Failed to get all role: %s", err)
	}
	if role == nil {
		t.Fatalf("Failed to get role")
	}

	nrole, err := rs.Get(role.ID)
	if err != nil {
		t.Fatalf("Failed to get all role: %s", err)
	}
	if nrole == nil {
		t.Fatalf("Failed to get role")
	}

	roles, err := rs.GetAll(10, 0)
	if err != nil {
		t.Fatalf("Failed to get all role: %s", err)
	}
	if len(roles) == 0 {
		t.Fatalf("Failed to get all role")
	}

	roles, err = rs.GetAll(10, 0, "qs")
	if err != nil {
		t.Fatalf("Failed to get all role: %s", err)
	}
	if len(roles) != 3 {
		t.Fatalf("Failed to get all role")
	}

	count, err := rs.Count()
	if err != nil {
		t.Fatalf("Failed to count role: %s", err)
	}
	if count == 0 {
		t.Fatalf("Failed to count role")
	}

	count, err = rs.Count("qs")
	if err != nil {
		t.Fatalf("Failed to count role: %s", err)
	}
	if count != 3 {
		t.Fatalf("Failed to count role")
	}

	err = rs.Delete("qs:admin")
	if err != nil {
		t.Fatalf("Failed to remove role: %s", err)
	}

	err = rs.Delete("qs:member")
	if err != nil {
		t.Fatalf("Failed to remove role: %s", err)
	}

	err = rs.Delete("qs:anonymous")
	if err != nil {
		t.Fatalf("Failed to remove role: %s", err)
	}

	db.Close()
}
