package migrate

import (
	"log"
	"path"

	"github.com/gunsluo/go-example/migrate/pkg/cli"
	"github.com/spf13/cobra"
	"github.com/xo/dburl"
)

// Cmd is the exported command.
var Cmd *cobra.Command

func init() {
	Cmd = &cobra.Command{
		Use:   "migrate command tool",
		Short: "runs migrate command",
		Long:  "Runs migrations for the selected service. --dsn requires a driver",
		Run:   migrate,
		Args:  cobra.MinimumNArgs(1),
	}

	Cmd.Flags().String("dsn", "postgres://postgres:password@localhost:5432/db?sslmode=disable", "run migrations on a specific database dsn")
	Cmd.Flags().String("path", "storage/sql", "specific path of sql file")
	Cmd.Flags().BoolP("dry-run", "n", false, "don't execute anything, only show the commands")
}

func migrate(cmd *cobra.Command, args []string) {
	dsn, err := cmd.Flags().GetString("dsn")
	if err != nil {
		log.Fatalln(err)
	}

	sqlPath, err := cmd.Flags().GetString("path")
	if err != nil {
		log.Fatalln(err)
	}

	db, err := dburl.Parse(dsn)
	if err != nil {
		log.Fatalln(err)
	}

	cli.RunMigrate(cmd, dsn, path.Join(sqlPath, db.Driver), args)
}
