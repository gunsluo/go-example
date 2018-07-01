package rbac

import (
	"os"
	"strings"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/xo/dburl"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var (
	dsn = "postgres://postgres:password@127.0.0.1:5432/ladon?sslmode=disable"
)

func createDB(t *testing.T, dbMockSetupOpts ...func(mock sqlmock.Sqlmock) error) *sqlx.DB {
	var db *sqlx.DB

	mode := os.Getenv("LADON_TEST_MODE")
	if strings.ToUpper(mode) == "DB" {
		db = createRealDB(t)
		tdsn := os.Getenv("LADON_TEST_DSN")
		if tdsn != "" {
			dsn = tdsn
		}
	} else {
		db = createMockDB(t, dbMockSetupOpts...)
	}

	return db
}

func createMockDB(t *testing.T, dbMockSetupOpts ...func(mock sqlmock.Sqlmock) error) *sqlx.DB {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Could not create mock database: %s", err)
	}

	for _, setupOptFn := range dbMockSetupOpts {
		if err = setupOptFn(mock); err != nil {
			t.Fatalf("run mock setup: %s", err)
		}
	}

	db := sqlx.NewDb(mockDB, "pq")
	return db
}

func createRealDB(t *testing.T) *sqlx.DB {
	// The database manager expects a sqlx.DB object
	dbURL, err := dburl.Parse(dsn)
	if err != nil {
		t.Fatalf("dsn invalid: %s", err)
	}

	dbname := strings.TrimPrefix(dbURL.Path, "/")
	if dbname != "" {
		dsnNoDB := strings.Replace(dbURL.DSN, "dbname="+dbname, "", 1)

		var tdb *sqlx.DB
		tdb, err = sqlx.Open(dbURL.Driver, dsnNoDB)
		if err != nil {
			t.Fatalf("Could not connect to database: %s", err)
		}

		_, err = tdb.Exec("CREATE DATABASE " + dbname)
		if err != nil {
			t.Logf("Create database[%s]: %s", dbname, err)
		}

		tdb.Close()
	}

	db, err := sqlx.Open(dbURL.Driver, dbURL.DSN)
	if err != nil {
		t.Fatalf("Could not connect to database: %s", err)
	}

	return db
}

func isAllInArray(array []string, params ...string) bool {
	return ArrayEquals(array, params)
}

func containArray(array []string, params ...string) bool {
	if len(array) < len(params) {
		return false
	}

	var found bool
	for i := range params {
		for j := range array {
			if params[i] == array[j] {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}
