package tools

import (
	"database/sql"
	"strings"

	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/golang-migrate/migrate/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/xo/dburl"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// MigrateUp use golang-migrate to migrate sql file to databases
// it's not support oracle
func MigrateUp(dsn, sqlPath string) error {
	u, err := dburl.Parse(dsn)
	if err != nil {
		return err
	}

	switch u.Driver {
	case "postgres":
		db, err := sql.Open(u.Driver, u.DSN)
		if err != nil {
			return err
		}

		instance, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			return err
		}

		dbname := strings.TrimPrefix(u.Path, "/")
		m, err := migrate.NewWithDatabaseInstance("file://"+sqlPath, dbname, instance)
		if err != nil {
			return err
		}

		err = m.Up()
		if err != nil {
			if err != migrate.ErrNoChange {
				return err
			}

			logrus.WithFields(logrus.Fields{
				"dsn":     dsn,
				"sqlPath": sqlPath,
			}).Infoln("migrate up: no change")
		} else {
			logrus.WithFields(logrus.Fields{
				"dsn":     dsn,
				"sqlPath": sqlPath,
			}).Infoln("migrate up: success")
		}
	default:
		return errors.Errorf("unsupport driver type %s", u.Driver)
	}

	return nil
}
