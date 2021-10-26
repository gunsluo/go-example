package main

import (
	"database/sql"
	"embed"
	"fmt"

	//_ "github.com/godror/godror"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/xo/dburl"
)

//go:embed migrations/postgres/*
var f embed.FS

func main() {
	dsn := "postgres://postgres:password@localhost:5432/db?sslmode=disable"

	var driver, dataSourceName string
	u, err := dburl.Parse(dsn)
	if err != nil {
		panic(err)
	}
	driver = u.Driver
	dataSourceName = u.DSN

	/*
			dirs, err := f.ReadDir(".")
			if err != nil {
				panic(err)
			}
			for _, d := range dirs {
				fmt.Println("==>", d.Name(), d.IsDir())
			}

				data, err := f.ReadFile("migrations/postgres/20200212105815-001.sql")
				fmt.Println(string(data), err)

		dir := http.FS(f)
		findMigrations(dir, "/migrations/postgres")
	*/

	// Read migrations from a embed fs:
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: f,
		Root:       "/migrations/postgres",
	}

	db, err := sql.Open(driver, dataSourceName)
	if err != nil {
		panic(err)
	}

	n, err := migrate.Exec(db, driver, migrations, migrate.Up)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
