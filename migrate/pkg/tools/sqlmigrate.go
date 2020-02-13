package tools

import (
	"database/sql"
	"fmt"
	"strings"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
	"github.com/xo/dburl"
)

// MigrateUp use sql-migrate to migrate sql file to databases
// it's support oracle
func SqlMigrateUp(dsn, sqlPath string) error {
	u, err := dburl.Parse(dsn)
	if err != nil {
		return err
	}

	var driver, dataSource string
	if strings.HasPrefix(dsn, "oci8://") {
		// fix wrong oracle driver
		driver = "oci8"
		dataSource = dsn[7:]
	} else {
		driver = u.Driver
		dataSource = u.DSN
	}

	// Read migrations from a folder:
	migrations := &migrate.FileMigrationSource{
		Dir: sqlPath,
	}

	db, err := sql.Open(driver, dataSource)
	if err != nil {
		return err
	}

	fmt.Println(driver, migrations.Dir)
	n, err := migrate.Exec(db, driver, migrations, migrate.Up)
	if err != nil {
		return err
	}
	logrus.WithFields(logrus.Fields{
		"dsn":     dsn,
		"sqlPath": sqlPath,
	}).Infof("Applied %d migrations!", n)

	return nil
}
