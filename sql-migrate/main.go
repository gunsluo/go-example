package main

import (
	"database/sql"
	"fmt"
	"path"
	"strings"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-oci8"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/xo/dburl"
)

func main() {
	//dsn := "postgres://postgres:password@localhost:5432/db?sslmode=disable"
	dsn := "oci8://c##admin/password@127.0.0.1:1521/ORCLCDB"

	var driver, dataSourceName, subDir string
	if strings.HasPrefix(dsn, "oci8://") {
		driver = "oci8"
		dataSourceName = dsn[7:]
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
	n, err := migrate.Exec(db, driver, migrations, migrate.Up)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
