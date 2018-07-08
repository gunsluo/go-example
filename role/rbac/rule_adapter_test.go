package rbac

import (
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestRuleAdapter(t *testing.T) {
	db := createDB(t, func(mock sqlmock.Sqlmock) error {
		mock.ExpectBegin()
		for i := 0; i < 3; i++ {
			mock.ExpectExec("INSERT INTO rule").WillReturnResult(sqlmock.NewResult(1, 1))
		}
		mock.ExpectCommit()
		mock.ExpectExec("INSERT INTO rule").WillReturnResult(sqlmock.NewResult(1, 1))

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
			mock.ExpectExec("DELETE FROM rule").WillReturnResult(sqlmock.NewResult(1, 1))
		}

		mock.ExpectQuery("^SELECT (.+) FROM rule").
			WillReturnRows(sqlmock.NewRows([]string{"total"}).AddRow(0))

		return nil
	})

	a := newRuleAdapter()
	tx, err := db.Beginx()
	if err != nil {
		t.Fatalf("Failed to begin transaction: %s", err)
	}

	err = a.Insert(tx, "g", []string{"user 1", "member", "U"})
	if err != nil {
		t.Fatalf("Failed to add rule: %s", err)
	}

	err = a.Insert(tx, "g", []string{"user 2", "member", "U"})
	if err != nil {
		t.Fatalf("Failed to add rule: %s", err)
	}

	err = a.Insert(tx, "g", []string{"user 3", "admin", "U"})
	if err != nil {
		t.Fatalf("Failed to add rule: %s", err)
	}

	if err = tx.Commit(); err != nil {
		t.Fatalf("Failed to commit transaction: %s", err)
	}

	err = a.Insert(db, "g", []string{"admin", "member", "R"})
	if err != nil {
		t.Fatalf("Failed to add rule: %s", err)
	}

	exist, err := a.Has(db, "g", []string{"user 3", "admin", "U"})
	if err != nil {
		t.Fatalf("Failed to has rule: %s", err)
	}
	if !exist {
		t.Fatalf("Failed to rule not exists")
	}

	exist, err = a.Has(db, "g", []string{"admin", "member", "R"})
	if err != nil {
		t.Fatalf("Failed to has rule: %s", err)
	}
	if !exist {
		t.Fatalf("Failed to rule not exists")
	}

	rules, err := a.GetFiltered(db, "g", 0)
	if err != nil {
		t.Fatalf("Failed to get filtered rule: %s", err)
	}
	if len(rules) != 4 {
		t.Fatal("failed: get filtered rule")
	}

	rules, err = a.GetFiltered(db, "g", 2, "U")
	if err != nil {
		t.Fatalf("Failed to get filtered rule: %s", err)
	}
	if len(rules) != 3 {
		t.Fatal("failed: get filtered rule")
	}

	rules, err = a.GetFiltered(db, "g", 1, "member", "U")
	if err != nil {
		t.Fatalf("Failed to get filtered rule: %s", err)
	}
	if len(rules) != 2 {
		t.Fatal("failed: get filtered rule")
	}

	rules, err = a.GetFiltered(db, "g", 0, "admin", "member", "R")
	if err != nil {
		t.Fatalf("Failed to get filtered rule: %s", err)
	}
	if len(rules) != 1 {
		t.Fatal("failed: get filtered rule")
	}

	err = a.Delete(db, "g", []string{"admin", "member", "R"})
	if err != nil {
		t.Fatalf("Failed to remove rule: %s", err)
	}

	err = a.RemoveFiltered(db, "g", 1, "member", "U")
	if err != nil {
		t.Fatalf("Failed to remove filtered rule: %s", err)
	}

	err = a.RemoveFiltered(db, "g", 2, "U")
	if err != nil {
		t.Fatalf("Failed to remove filtered rule: %s", err)
	}

	count, err := a.Count(db)
	if err != nil {
		t.Fatalf("Failed to count rule: %s", err)
	}
	if count != 0 {
		t.Fatalf("Failed to count rule: not zero")
	}

	db.Close()
}
