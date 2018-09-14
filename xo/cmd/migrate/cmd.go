package migrate

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gunsluo/go-example/xo/common"
	"github.com/spf13/cobra"
	"gitlab.com/tesgo/kit/pkg/cli"
)

const (
	crDSN = "cockroach://root@localhost:26257/%s?sslmode=disable"
	pgDSN = "postgres://%s:%s@localhost:5432/%s?sslmode=disable"
)

var (
	sqlPaths = map[string]string{
		"demo": "model/sql",
	}
)

// Cmd is the exported command.
var Cmd *cobra.Command

func init() {
	Cmd = &cobra.Command{
		Use:   "migrate SERVICE COMMAND",
		Short: "runs migrations",
		Example: "  cmd migrate auth up\n" +
			" cmd migrate kadir -v down\n" +
			"  cmd migrate kadir version\n" +
			"  cmd migrate --db reach --cr auth up\n" +
			"  cmd migrate --pg --dsn postgres://postgres:password@12.34.56.78:1337/db",
		Long: "Runs migrations for the selected service. --dsn requires a driver " +
			"to be selected (--cr or --pg), and running it without a driver will " +
			"migrate all dbs (CR and PG).",
		Args: cobra.MinimumNArgs(2),
		Run:  migrate,
	}

	Cmd.Flags().Bool("cr", false, "run migrations on cockroachdb")
	Cmd.Flags().Bool("pg", false, "run migrations on postgres")
	Cmd.Flags().String("dsn", "", "run migrations on a specific database dsn")
	Cmd.Flags().String("db", "", "run migrations on a specific database name")

	Cmd.Flags().BoolP("dry-run", "n", false, "don't execute anything, only show the commands")
}

func migrate(cmd *cobra.Command, args []string) {
	var cr, pg bool
	var dsn, dbname string
	var err error

	// Set database URL and whether we run migrate on cr, pg, or both
	cr, err = cmd.Flags().GetBool("cr")
	if err != nil {
		log.Fatalln(err)
	}
	pg, err = cmd.Flags().GetBool("pg")
	if err != nil {
		log.Fatalln(err)
	}
	dsn, err = cmd.Flags().GetString("dsn")
	if err != nil {
		log.Fatalln(err)
	}
	dbname, err = cmd.Flags().GetString("db")
	if err != nil {
		log.Fatalln(err)
	}

	// Set service
	svc := args[0]
	if _, ok := sqlPaths[svc]; !ok {
		log.Fatalf("Service %s not supported\n", svc)
	}
	if cli.Verbose(cmd) {
		fmt.Println("== Service selected:", svc)
	}

	// Set custom db url
	if dsn != "" {
		if !cr && !pg {
			log.Fatalln("A custom db URL needs to specify db type")
		}
		fmt.Println("== Custom db selected:", dsn)
	}

	if !cr && !pg {
		cr = true
		pg = true
	}
	if cr {
		migrateCR(cmd, dsn, dbname, svc, args)
	}
	if pg {
		migratePG(cmd, dsn, dbname, svc, args)
	}
}

func migrateCR(cmd *cobra.Command, dsn, dbname, svc string, args []string) {
	if dsn == "" {
		if dbname == "" {
			dbname = svc
		}

		dsn = fmt.Sprintf(crDSN, dbname)
	}

	sqlDir := sqlPaths[svc]
	if cli.Verbose(cmd) {
		fmt.Println("== Running CR migrations for service", svc, "on db", dsn)
	}

	sqlAbsPath := filepath.Join(common.GetProjectPath(), sqlDir)
	cli.RunMigrate(cmd, dsn, sqlAbsPath, args)
}

func migratePG(cmd *cobra.Command, dsn, dbname, svc string, args []string) {
	if dsn == "" {
		if dbname == "" {
			dbname = svc
		}

		dsn = fmt.Sprintf(pgDSN, os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), dbname)
	}

	sqlDir := sqlPaths[svc]
	if cli.Verbose(cmd) {
		fmt.Println("== Running PG migrations for service", svc, "on db", dsn)
	}

	sqlAbsPath := filepath.Join(common.GetProjectPath(), sqlDir)
	cli.RunMigrate(cmd, dsn, sqlAbsPath, args)
}
