package build

import (
	"log"
	"path/filepath"

	"github.com/gunsluo/go-example/migrate/common"
	"github.com/gunsluo/go-example/migrate/pkg/cli"
	"github.com/spf13/cobra"
)

// Cmd is the exported command.
var Cmd *cobra.Command

func init() {
	Cmd = &cobra.Command{
		Use:   "build",
		Short: "builds go code",
		Run:   build,
	}

	Cmd.Flags().Bool("xo", false, "run xo build")
	Cmd.Flags().String("dsn", "postgres://postgres:password@localhost:5432/db?sslmode=disable,sqlserver://SA:password@localhost:1433?database=db&encrypt=disable",
		"run building on a specific database dsn")
	Cmd.Flags().String("template", "", "path of templates file")
	Cmd.Flags().String("path", "", "specific path of sql file")
	Cmd.Flags().StringP("output", "o", "storage", "the path of generating code file")
	Cmd.Flags().BoolP("dry-run", "n", false, "don't execute anything, only show the commands")
}

func build(cmd *cobra.Command, args []string) {
	xo, err := cmd.Flags().GetBool("xo")
	if err != nil {
		log.Fatalln(err)
	}

	dsn, err := cmd.Flags().GetString("dsn")
	if err != nil {
		log.Fatalln(err)
	}

	templatePath, err := cmd.Flags().GetString("template")
	if err != nil {
		log.Fatalln(err)
	}

	sqlPath, err := cmd.Flags().GetString("path")
	if err != nil {
		log.Fatalln(err)
	}

	outputPath, err := cmd.Flags().GetString("output")
	if err != nil {
		log.Fatalln(err)
	}

	if xo {
		buildXO(cmd, dsn, templatePath, sqlPath, outputPath)
	}
}

func buildXO(cmd *cobra.Command, dsn, templatePath, sqlPath, outputPath string) {
	xoPath := common.GetXOPath()
	if templatePath == "" {
		templatePath = filepath.Join(xoPath, "templates")
	}

	if sqlPath == "" {
		sqlPath = filepath.Join(common.GetProjectPath(), "storage/sql")
	}

	cli.BuildXO(cmd, xoPath, sqlPath, cli.XOCommand{
		DSN:          dsn,
		TemplatePath: templatePath,
		OutputPath:   outputPath,
		Args: []string{
			"--package", "storage",
			"--ignore-tables", "schema_migrations", "schema_lock",
			"--enable-ac",
			"--enable-extension",
		},
	})
}
