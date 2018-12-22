package main

import (
	"context"
	"fmt"

	"github.com/gunsluo/go-example/sqlboiler/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

//go:generate sqlboiler --wipe psql
func main() {
	//db, err := sql.Open("postgres", "dbname=test host=localhost user=postgres password=password sslmode=disable")
	dsn := "postgres://postgres:password@localhost:5432/test?sslmode=disable"
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	count, err := models.Pilots().Count(ctx, db)
	if err != nil {
		panic(err)
	}

	fmt.Println("count:", count)

	rows, err := models.NewQuery(qm.From("pilots")).QueryContext(ctx, db)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		//pilot := models.Pilot{}
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			panic(err)
		}
		fmt.Println("pilot:", id, name)
	}

	type NameAndLanguage struct {
		Name     string
		Language string
	}

	var nl NameAndLanguage
	err = queries.Raw("SELECT p.name,l.language FROM pilots as p LEFT JOIN pilot_languages as pl ON p.id=pl.pilot_id LEFT JOIN languages as l ON pl.language_id=l.id WHERE p.name=$1", "luoji").
		Bind(ctx, db, &nl)
	if err != nil {
		panic(err)
	}
	fmt.Println("name and language:", nl)

	var nls []NameAndLanguage
	err = queries.Raw("SELECT p.name,l.language FROM pilots as p LEFT JOIN pilot_languages as pl ON p.id=pl.pilot_id LEFT JOIN languages as l ON pl.language_id=l.id").
		Bind(ctx, db, &nls)
	if err != nil {
		panic(err)
	}
	for _, nl := range nls {
		fmt.Println("name and language:", nl)
	}
}
