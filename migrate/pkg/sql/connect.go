package sql

import (
	"log"

	"github.com/cenk/backoff"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const defaultRetries = 6

// Connect tries to connect to a database up to n times.
// There are `DBSetupOpt`s to do things like run migrations, create system objects, etc.
func Connect(logger *logrus.Logger, driver, dsn string, n int, opts ...DBSetupOpt) (*sqlx.DB, error) {
	var db *sqlx.DB
	var err error
	retries := 0

	if n < 1 {
		n = defaultRetries - 1
	}

	try := func() error {
		logger.WithFields(logrus.Fields{"retries-left": n - retries, "dsn": dsn}).Infoln("Connecting to db")

		// Check if DSN is in order. If not, return nil and check err value.
		db, err = sqlx.Open(driver, dsn)
		if err != nil {
			// Bad DSN, we quit immediately
			err = errors.Wrap(err, "bad dsn")
			return nil
		}

		if dbErr := db.Ping(); dbErr != nil {

			retries++
			return errors.Wrap(dbErr, "couldn't ping db")
		}

		return nil
	}

	boff := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), uint64(n))
	errBackoff := backoff.Retry(try, boff)

	// Bad dsn.
	if err != nil {
		return nil, err
	}

	// Couldn't connect after n attempts.
	if errBackoff != nil {
		return nil, errBackoff
	}

	for i, opt := range opts {
		if err := opt(logger, db.DB); err != nil {
			log.Fatalf("Failed running setup step %d: %v\n", i, err)
		}
	}

	return db, nil
}
