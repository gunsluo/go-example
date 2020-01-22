package cli

import (
	"log"

	"github.com/spf13/cobra"
)

func RunMigrate(cmd *cobra.Command, dsn, sqlDir string, args []string) {
	cmdArgs := []string{"migrate"}
	if Verbose(cmd) {
		cmdArgs = append(cmdArgs, "-verbose")
	}

	switch args[0] {
	case "create":
		// migrate create -ext sql -dir ./sql/postgres -seq account
		cmdArgs = append(cmdArgs, "create")
		cmdArgs = append(cmdArgs, "-ext")
		cmdArgs = append(cmdArgs, "sql")
		cmdArgs = append(cmdArgs, "-dir")
		cmdArgs = append(cmdArgs, sqlDir)
		cmdArgs = append(cmdArgs, "-seq")
		cmdArgs = append(cmdArgs, args[1:]...)

		if err := ForkDir(cmd, cmdArgs, "."); err != nil {
			log.Fatalln("Couldn't run migrations:", err)
		}
		return
	case "up", "down", "drop", "force", "goto":
		// migrate -verbose -path model/sql -database 'cockroach://root@localhost:26257/people_hub?sslmode=disable' up
		cmdArgs = append(cmdArgs, "-path")
		cmdArgs = append(cmdArgs, ".")
		cmdArgs = append(cmdArgs, "-database")
		cmdArgs = append(cmdArgs, dsn)
		cmdArgs = append(cmdArgs, args...)
	case "version":
		// migrate -verbose -path model/sql -database 'cockroach://root@localhost:26257/people_hub?sslmode=disable' version
		cmdArgs = append(cmdArgs, "-database")
		cmdArgs = append(cmdArgs, dsn)
		cmdArgs = append(cmdArgs, args...)
	}

	if err := ForkDir(cmd, cmdArgs, sqlDir); err != nil {
		log.Fatalln("Couldn't run migrations:", err)
	}
}
