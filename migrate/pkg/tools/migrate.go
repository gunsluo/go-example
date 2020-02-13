package tools

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/sirupsen/logrus"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/sqlserver"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// MigrateUp use golang-migrate to migrate sql file to databases
// it's not support oracle
func MigrateUp(dsn, sqlPath string) error {
	m, err := migrate.New("file://"+sqlPath, dsn)
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

	return nil
}
