package serve

import (
	"log"
	"path"

	"github.com/gunsluo/go-example/migrate/common"
	"github.com/gunsluo/go-example/migrate/pkg/sql"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "serve",
	Short: "Start a service and run migrate sql",
	Run:   run,
}

func init() {
	Cmd.Flags().String("dsn", "postgres://postgres:password@localhost:5432/db?sslmode=disable", "a specific database dsn")
}

func run(cmd *cobra.Command, _ []string) {
	dsn, err := cmd.Flags().GetString("dsn")
	if err != nil {
		log.Fatalln(err)
	}

	logger := logrus.New()
	// Set up database connection.
	if err := sql.CreateDBIfNotExist(logger, dsn); err != nil {
		log.Fatalln(errors.Wrap(err, "couldn't create db "+dsn))
	}

	sqlPath := path.Join(common.GetProjectPath(), "storage/migrations")
	db, err := sql.Connect(logger, dsn, 6, sql.Migrate(dsn, "test", sqlPath))
	if err != nil {
		log.Fatalln(errors.Wrap(err, "couldn't connect to db "+dsn))
	}

	db.Close()
}
