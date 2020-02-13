package sql

import (
	"database/sql"
	"path"
	"strings"

	"github.com/cenk/backoff"
	"github.com/gunsluo/go-example/migrate/pkg/tools"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/xo/dburl"
)

const (
	migrateRetries = 12
)

// MMigrateFuncs is a mapping  between specified database and the migrate tool
var MigrateFuncs = map[string]func(string, string) error{
	"postgres": tools.MigrateUp,
	"godror":   tools.SqlMigrateUp,
}

// Migrate runs migrations. It's supposed to run in production,
// so assume the target database is CockroachDB.
func Migrate(dsn, service string, sqlPath string) DBSetupOpt {
	return func(logger *logrus.Logger, _ *sql.DB) error {
		if dsn == "" {
			return errors.New("no dsn specified")
		}

		u, err := dburl.Parse(dsn)
		if err != nil {
			return errors.Wrap(err, "couldn't parse database address")
		}

		migrateUp, ok := MigrateFuncs[u.Driver]
		if !ok {
			return errors.Errorf("unsupported database %s", u.Driver)
		}

		var subPath string
		if strings.HasPrefix(dsn, "oci8://") {
			subPath = "oracle"
		} else {
			subPath = u.Driver
		}

		// This is needed because the migrate tool recognizes 'postgres/postgresql'
		// driver part of the DSN and applies the Postgres driver instead of the Cockroach one.
		// dsn = strings.Replace(dsn, "postgresql", "cockroach", 1)
		// dsn = strings.Replace(dsn, "postgres", "cockroach", 1)

		sqlPath = path.Join(sqlPath, subPath)
		logger.WithFields(logrus.Fields{"dsn": dsn, "sql-path": sqlPath}).Infoln("Running migrations")

		var retry int
		migrate := func() error {
			err := migrateUp(dsn, sqlPath)
			if err != nil {
				logger.WithFields(logrus.Fields{
					"err":     err,
					"retries": migrateRetries - retry,
				}).Infoln("Retrying to run migrations")
				retry++
			}

			return err
		}

		return backoff.Retry(migrate,
			backoff.WithMaxRetries(backoff.NewExponentialBackOff(), migrateRetries))
	}
}
