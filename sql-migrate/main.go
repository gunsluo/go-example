package main

import (
	"database/sql"
	"fmt"
	"path"
	"strings"

	_ "github.com/godror/godror"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/xo/dburl"
)

func main() {
	//dsn := "postgres://postgres:password@localhost:5432/db?sslmode=disable"
	dsn := "oracle://c##admin/password@127.0.0.1:1521/ORCLCDB"

	var driver, dataSourceName, subDir string
	if strings.HasPrefix(dsn, "oracle://") {
		driver = "godror"
		dataSourceName = dsn[9:]
		subDir = "oracle"
	} else {
		u, err := dburl.Parse(dsn)
		if err != nil {
			panic(err)
		}
		driver = u.Driver
		dataSourceName = u.DSN
		subDir = u.Driver
	}

	// Read migrations from a folder:
	migrations := &migrate.FileMigrationSource{
		Dir: path.Join("migrations", subDir),
	}

	db, err := sql.Open(driver, dataSourceName)
	if err != nil {
		panic(err)
	}

	fmt.Println(driver, migrations.Dir)
	n, err := migrate.Exec(db, "godror", migrations, migrate.Up)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
