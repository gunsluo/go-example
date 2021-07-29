package unittest

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/xo/dburl"
)

func CreateDBIfNotExist(dsn string) error {
	u, err := dburl.Parse(dsn)
	if err != nil {
		return fmt.Errorf("couldn't parse database address, %w", err)
	}

	// create db
	dbname := strings.TrimPrefix(u.Path, "/")
	if dbname != "" {
		u.Path = ""
		if cdb, err := Connect(u.String()); err != nil {
			return fmt.Errorf("Couldn't connect to db with dsn: %v, %w", u.String(), err)
		} else {
			_, err = cdb.Exec("CREATE DATABASE " + dbname)
			if err == nil {
				fmt.Printf("Created empty database %s, as it did not exist\n", dbname)
			}
			cdb.Close()
		}
	}

	return nil
}

func Connect(dsn string) (*sqlx.DB, error) {
	u, err := dburl.Parse(dsn)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse database address, %w", err)
	}
	driver, dataSource := u.Driver, u.DSN

	// Check if DSN is in order. If not, return nil and check err value.
	db, err := sqlx.Open(driver, dataSource)
	if err != nil {
		// Bad DSN, we quit immediately
		err = fmt.Errorf("bad dsn, %w", err)
		return nil, err
	}

	if dbErr := db.Ping(); dbErr != nil {
		return nil, fmt.Errorf("couldn't ping db, %w", dbErr)
	}

	return db, nil
}

func DropDBIfExist(dsn string) error {
	u, err := dburl.Parse(dsn)
	if err != nil {
		return fmt.Errorf("couldn't parse database address, %w", err)
	}

	// create db
	dbname := strings.TrimPrefix(u.Path, "/")
	if dbname != "" {
		u.Path = ""
		if cdb, err := Connect(u.String()); err != nil {
			return fmt.Errorf("Couldn't connect to db with dsn: %v, %w", u.String(), err)
		} else {
			_, err = cdb.Exec("DROP DATABASE " + dbname)
			if err == nil {
				fmt.Printf("Drop database %s, as it did exist\n", dbname)
			}
			cdb.Close()
		}
	}

	return nil
}
