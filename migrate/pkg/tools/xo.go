package tools

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/xo/dburl"
	"github.com/xo/xo/cli"
)

// XOArguments is command parameters for xo tool
type XOArguments cli.Arguments

// MMigrateFuncs is a mapping  between specified database and the migrate tool
var MigrateFuncs = map[string]func(string, string) error{
	"postgres": MigrateUp,
	"mssql":    MigrateUp,
	"godror":   SqlMigrateUp,
}

// BuildXO builds db into the .xo.go files
func BuildXO(cmd *cobra.Command, sqlPath string, args XOArguments) {
	if Verbose(cmd) {
		fmt.Println("== Running xo build")
	}

	// step 1: Migrate up
	for _, dsn := range args.DSNS {
		u, err := dburl.Parse(dsn)
		if err != nil {
			log.Fatalln(err)
		}

		migrateUp, ok := MigrateFuncs[u.Driver]
		if !ok {
			log.Fatalln(errors.Errorf("unsupported database %s", u.Driver))
		}

		var subPath string
		if strings.HasPrefix(dsn, "oci8://") {
			subPath = "oracle"
		} else {
			subPath = u.Driver
		}

		err = migrateUp(dsn, path.Join(sqlPath, subPath))
		if err != nil {
			log.Fatalln(err)
		}
	}

	// step 2: run xo generator
	//buildXO(cmd, args)
}

func buildXO(cmd *cobra.Command, args XOArguments) {
	// create a temporary dir to store .xo.go
	tmpDir, err := ioutil.TempDir(args.Out, "buildxo")
	if err != nil {
		log.Fatalln("Couldn't create temporary directory:", err)
	}
	defer os.RemoveAll(tmpDir)

	output := args.Out
	args.Out = tmpDir
	if err := xoGen(cmd, tmpDir, args); err != nil {
		// remove temporary dir
		return
	}

	// move xo.go to output
	if Verbose(cmd) {
		fmt.Println("== moving generated xo back to:", output)
	}

	files, err := ioutil.ReadDir(tmpDir)
	if err != nil {
		fmt.Println("Can't read temporary directory:", err)
		return
	}

	oldFiles, err := ioutil.ReadDir(output)
	if err != nil {
		fmt.Println("Can't read output directory:", err)
		return
	}

	// remove all xo.go
	for _, fi := range oldFiles {
		if fi.IsDir() || !strings.HasSuffix(fi.Name(), ".xo.go") {
			continue
		}

		filePath := filepath.Join(output, fi.Name())
		if err := os.Remove(filePath); err != nil {
			fmt.Println("failed to remove file:", err)
			return
		}
	}

	for _, fi := range files {
		if fi.IsDir() || !strings.HasSuffix(fi.Name(), ".xo.go") {
			continue
		}

		srcFilePath := filepath.Join(tmpDir, fi.Name())
		outputFilePath := filepath.Join(output, fi.Name())

		f, err := os.Open(outputFilePath)
		if err == nil {
			scanner := bufio.NewScanner(f)
			if scanner.Scan() {
				if strings.HasPrefix(scanner.Text(), `// skip`) {
					continue
				}
			}
		}

		if err := os.Rename(srcFilePath, outputFilePath); err != nil {
			fmt.Println("failed to move file:", err)
			return
		}
	}
}

func xoGen(cmd *cobra.Command, tmpDir string, args XOArguments) error {
	if Verbose(cmd) {
		fmt.Println("== Running xo build on dsn ", args.DSNS, " output ", args.Out)
		args.Verbose = true
	}

	// generate the code
	err := cli.Generate(cli.Arguments(args))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"dsn":      args.DSNS,
			"template": args.TemplatePath,
			"output":   tmpDir,
		}).WithError(err).Infoln("Failed to generate xo code")
		return err
	}

	return nil
}
