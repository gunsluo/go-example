package rbac

import (
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestRuleAdapter(t *testing.T) {
	db := createDB(t, func(mock sqlmock.Sqlmock) error {
		for i := 0; i < 4; i++ {
			mock.ExpectBegin()
			mock.ExpectExec("INSERT INTO rule").WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()
		}

		for i := 0; i < 2; i++ {
			mock.ExpectQuery("^SELECT (.+) FROM rule (.+)").
				WillReturnRows(sqlmock.NewRows([]string{"total"}).AddRow(1))
		}

		mock.ExpectQuery("^SELECT (.+) FROM rule (.+)").
			WillReturnRows(sqlmock.NewRows([]string{"id", "p_type", "v0", "v1", "v2", "v3", "v4", "v5"}).
				AddRow(1, 2, 3, 4, 5, 6, 7, 8).
				AddRow(1, 2, 3, 4, 5, 6, 7, 8).
				AddRow(1, 2, 3, 4, 5, 6, 7, 8).
				AddRow(1, 2, 3, 4, 5, 6, 7, 8))

		mock.ExpectQuery("^SELECT (.+) FROM rule (.+)").
			WillReturnRows(sqlmock.NewRows([]string{"id", "p_type", "v0", "v1", "v2", "v3", "v4", "v5"}).
				AddRow(1, 2, 3, 4, 5, 6, 7, 8).
				AddRow(1, 2, 3, 4, 5, 6, 7, 8).
				AddRow(1, 2, 3, 4, 5, 6, 7, 8))

		mock.ExpectQuery("^SELECT (.+) FROM rule (.+)").
			WillReturnRows(sqlmock.NewRows([]string{"id", "p_type", "v0", "v1", "v2", "v3", "v4", "v5"}).
				AddRow(1, 2, 3, 4, 5, 6, 7, 8).
				AddRow(1, 2, 3, 4, 5, 6, 7, 8))

		mock.ExpectQuery("^SELECT (.+) FROM rule (.+)").
			WillReturnRows(sqlmock.NewRows([]string{"id", "p_type", "v0", "v1", "v2", "v3", "v4", "v5"}).
				AddRow(1, 2, 3, 4, 5, 6, 7, 8))

		for i := 0; i < 3; i++ {
			mock.ExpectBegin()
			mock.ExpectExec("DELETE FROM rule").WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()
		}

		mock.ExpectQuery("^SELECT (.+) FROM rule").
			WillReturnRows(sqlmock.NewRows([]string{"total"}).AddRow(0))

		return nil
	})

	a := newRuleAdapter(db)
	err := a.Add("g", []string{"user 1", "member", "S"})
	if err != nil {
		t.Fatalf("Failed to add rule: %s", err)
	}

	err = a.Add("g", []string{"user 2", "member", "S"})
	if err != nil {
		t.Fatalf("Failed to add rule: %s", err)
	}

	err = a.Add("g", []string{"user 3", "admin", "S"})
	if err != nil {
		t.Fatalf("Failed to add rule: %s", err)
	}

	err = a.Add("g", []string{"admin", "member", "R"})
	if err != nil {
		t.Fatalf("Failed to add rule: %s", err)
	}

	exist, err := a.Has("g", []string{"user 3", "admin", "S"})
	if err != nil {
		t.Fatalf("Failed to has rule: %s", err)
	}
	if !exist {
		t.Fatalf("Failed to rule not exists")
	}

	exist, err = a.Has("g", []string{"admin", "member", "R"})
	if err != nil {
		t.Fatalf("Failed to has rule: %s", err)
	}
	if !exist {
		t.Fatalf("Failed to rule not exists")
	}

	rules, err := a.GetFiltered("g", 0)
	if err != nil {
		t.Fatalf("Failed to get filtered rule: %s", err)
	}
	if len(rules) != 4 {
		t.Fatal("failed: get filtered rule")
	}

	rules, err = a.GetFiltered("g", 2, "S")
	if err != nil {
		t.Fatalf("Failed to get filtered rule: %s", err)
	}
	if len(rules) != 3 {
		t.Fatal("failed: get filtered rule")
	}

	rules, err = a.GetFiltered("g", 1, "member", "S")
	if err != nil {
		t.Fatalf("Failed to get filtered rule: %s", err)
	}
	if len(rules) != 2 {
		t.Fatal("failed: get filtered rule")
	}

	rules, err = a.GetFiltered("g", 0, "admin", "member", "R")
	if err != nil {
		t.Fatalf("Failed to get filtered rule: %s", err)
	}
	if len(rules) != 1 {
		t.Fatal("failed: get filtered rule")
	}

	err = a.Remove("g", []string{"admin", "member", "R"})
	if err != nil {
		t.Fatalf("Failed to remove rule: %s", err)
	}

	err = a.RemoveFiltered("g", 1, "member", "S")
	if err != nil {
		t.Fatalf("Failed to remove filtered rule: %s", err)
	}

	err = a.RemoveFiltered("g", 2, "S")
	if err != nil {
		t.Fatalf("Failed to remove filtered rule: %s", err)
	}

	count, err := a.Count()
	if err != nil {
		t.Fatalf("Failed to count rule: %s", err)
	}
	if count != 0 {
		t.Fatalf("Failed to count rule: not zero")
	}

	db.Close()
}
