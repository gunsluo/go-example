package sql

import (
	"database/sql"

	"github.com/cenk/backoff"
	"github.com/gunsluo/go-example/migrate/pkg/cli"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	migrateRetries = 12
)

// Migrate runs migrations. It's supposed to run in production,
// so assume the target database is CockroachDB.
func Migrate(dsn, service string, sqlPath string) DBSetupOpt {
	return func(logger *logrus.Logger, _ *sql.DB) error {
		if dsn == "" {
			return errors.New("no dsn specified")
		}

		// This is needed because the migrate tool recognizes 'postgres/postgresql'
		// driver part of the DSN and applies the Postgres driver instead of the Cockroach one.
		// dsn = strings.Replace(dsn, "postgresql", "cockroach", 1)
		// dsn = strings.Replace(dsn, "postgres", "cockroach", 1)

		logger.WithFields(logrus.Fields{"dsn": dsn, "sql-path": sqlPath}).Infoln("Running migrations")

		var retry int
		//migrate -database 'sqlserver://SA:password@localhost:1433?database=db&encrypt=disable' -path ./sql/mssql up
		migrate := func() error {
			err := cli.ForkDir(nil, []string{"migrate",
				"-verbose",
				"-database", dsn,
				// "-lock-timeout", "30", // until it's fixed we use backoff
				"-path", ".",
				"up"},
				sqlPath)

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
			backoff.WithMaxRetries(backoff.NewExponentialBackOff(), uint64(10)))
	}
}
