package sql

import (
	"database/sql"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/xo/dburl"
)

// DBSetupOpt is a function to perform some setup task on a database.
type DBSetupOpt func(*logrus.Logger, *sql.DB) error

// CreateDB creates and "USE"s a database.
func CreateDB(dbname string) DBSetupOpt {
	return func(logger *logrus.Logger, db *sql.DB) error {
		// TODO regex to only match valid identifiers
		// CRDB and/or database/sql are stupid and don't let you use db names as parameters.
		_, err := db.Exec("CREATE DATABASE " + dbname)
		if err == nil {
			logger.Infof("Created empty database %s, as it did not exist", dbname)
		}
		//_, err = db.Exec("USE " + dbname)
		//if err != nil {
		//logger.Fatalln("Failed to USE db:", err)
		//}

		return nil
	}
}

func CreateDBIfNotExist(logger *logrus.Logger, dsn string) error {
	db, err := dburl.Parse(dsn)
	if err != nil {
		return errors.Wrap(err, "couldn't parse database address")
	}

	// create db
	dbname := strings.TrimPrefix(db.Path, "/")
	if dbname != "" {
		dsnNoDB := strings.Replace(db.DSN, "dbname="+dbname, "", 1)
		if cdb, err := Connect(logger, db.Driver, dsnNoDB, 1); err != nil {
			return errors.Wrap(err, "Couldn't connect to db ")
		} else {
			_, err = cdb.Exec("CREATE DATABASE " + dbname)
			if err == nil {
				logger.Infof("Created empty database %s, as it did not exist", dbname)
			}
			cdb.Close()
		}
	}

	return nil
}

func CloseDB() DBSetupOpt {
	return func(logger *logrus.Logger, db *sql.DB) error {
		db.Close()
		return nil
	}
}

/*
// CreateSystemUser creates a system user.
func CreateSystemUser(_ *logrus.Logger, db *sql.DB) error {
	_, err := db.Exec(`INSERT INTO "user" (
		id,
		username,
		email,
		capabilities
	) VALUES (
		1,
		'system',
		'system@localhost',
		'*'
	) ON CONFLICT (id) DO NOTHING;`)

	return err
}

// CreateSystemOAuthClient creates a system OAuth client.
func CreateSystemOAuthClient(_ *logrus.Logger, db *sql.DB) error {
	_, err := db.Exec(`INSERT INTO "oauth_client" (
		id,
		secret,
		user_id,
		name,
		redirect_uri,
		grant_type,
		is_system,
		scope
	) VALUES (
		1,
		'` + crypto.RandN(32) + `',
		1,
		'System Password Auth Client',
		'http://localhost',
		'password',
		true,
		'*'
	) ON CONFLICT (id) DO NOTHING;`)

	return err
}
*/
