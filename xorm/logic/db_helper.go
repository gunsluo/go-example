package logic

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

const (
	envTestDBDSN     = "TEST_DB_DSN"
	defaultTestDBDSN = "postgres://postgres:password@127.0.0.1:5432/mock_temporary?sslmode=disable"
)

func getDSN() string {
	dsn := os.Getenv(envTestDBDSN)
	if dsn == "" {
		dsn = defaultTestDBDSN
	}

	return dsn
}

/*
func initDB() (*sqlx.DB, error) {
	dsn := getDSN()
	migrationPath := ""

	// Set up database connection.
	if err := sqlz.CreateDBIfNotExist(dsn); err != nil {
		return nil, fmt.Errorf("couldn't create db: %s, %w", dsn, err)
	}

	var options []sqlz.DBSetupOpt
	if migrationPath != `` {
		options = append(options, sqlz.Migrate(dsn, migrationPath))
	}

	db, err := sqlz.Connect(dsn, 1, options...)
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to db: %s, %w", dsn, err)
	}

	return db, nil
}
*/

func newEngine() *xorm.Engine {
	dsn := getDSN()
	engine, err := xorm.NewEngine("postgres", dsn)
	if err != nil {
		panic(err)
	}
	engine.DB().DB.Close()

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	//db.SetMaxOpenConns(1)
	//db.SetMaxIdleConns(1)

	ctx := context.Background()
	_, err = db.ExecContext(ctx, "SET search_path TO pg_temp_1")
	if err != nil {
		panic(err)
	}
	_, err = db.ExecContext(ctx, `CREATE TEMP TABLE
		IF
			NOT EXISTS inspector (
				ID SERIAL PRIMARY KEY NOT NULL,
				username VARCHAR ( 256 ) NOT NULL DEFAULT '',
				password VARCHAR ( 256 ) NOT NULL DEFAULT '',
				created TIMESTAMP NOT NULL DEFAULT NOW()
			);`)

	if err != nil {
		panic(err)
	}

	rs, err := db.QueryContext(ctx, `SELECT "id", "username", "password", "created" FROM "inspector" WHERE (id = $1)`, 1)
	fmt.Println("-->", rs, err)

	//tx.Commit()
	//inspector := &models.Inspector{}
	//has, err := session.Where("id = ?", 1).Get(inspector)
	//fmt.Println("-->", has, err)

	/*
			rs, err := db.Exec("SET search_path TO pg_temp_1;")
			fmt.Println("-->", rs, err)
			db.Exec(`CREATE TABLE
		IF
			NOT EXISTS inspector (
				ID SERIAL PRIMARY KEY NOT NULL,
				username VARCHAR ( 256 ) NOT NULL DEFAULT '',
				password VARCHAR ( 256 ) NOT NULL DEFAULT '',
				created TIMESTAMP NOT NULL DEFAULT NOW()
			);`)
	*/

	engine.DB().DB = db
	engine.ShowSQL(true)

	return engine
}

func newSession() *xorm.Session {
	engine := newEngine()
	session := engine.NewSession()

	return session
}
