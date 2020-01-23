package cli

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/xo/dburl"
)

// BuildGo build a golang project
func BuildGo(cmd *cobra.Command, pathShort string) {
	if Verbose(cmd) {
		fmt.Println("== Running Go build")
	}

	if err := ForkSplit(cmd, fmt.Sprintf("go install %s", pathShort)); err != nil {
		log.Fatalln("couldn't go install:", err)
	}
}

// BuildProto builds the protobuf to Go files
func BuildProto(cmd *cobra.Command, rootDir, pbDocShortPath string, protoShortPaths, includePaths []string) {
	BuildProtoWithMock(cmd, rootDir, pbDocShortPath, protoShortPaths, nil, includePaths)
}

// BuildProtoWithMock is same as BuildProto, but generates mock as well
func BuildProtoWithMock(cmd *cobra.Command, rootDir, pbDocShortPath string, protoShortPaths, mockShortPaths, includePaths []string) {
	if Verbose(cmd) {
		fmt.Println("== Running proto build")
	}

	Install(cmd, "protoc", "")
	Install(cmd, "protoc-gen-gogoslick", "")
	Install(cmd, "protoc-gen-doc", "")

	if len(mockShortPaths) != 0 {
		Install(cmd, "mockgen", "")
	}

	pbDocPath := filepath.Join(rootDir, pbDocShortPath)
	if err := mkdirIfNotExist(pbDocPath); err != nil {
		log.Fatalln("Couldn't mkdir:", pbDocPath)
	}

	for i, protoPath := range protoShortPaths {
		if Verbose(cmd) {
			fmt.Println("== Generating pb.go for", protoPath)
		}

		absPbDir := filepath.Join(rootDir, protoPath)

		if DryRun(cmd) {
			fmt.Println("== Reading contents of", absPbDir, "and running protoc")
		}

		// mock
		var mockAbsPath string
		if len(mockShortPaths) != 0 {
			mockAbsPath = filepath.Join(rootDir, mockShortPaths[i])
			os.MkdirAll(mockAbsPath, 0755)
		}

		files, err := ioutil.ReadDir(absPbDir)
		if err != nil {
			log.Fatalln("Couldn't read dir:", absPbDir)
		}

		// args
		args := []string{
			"protoc",
			"-I=.",
			"-I=$GOPATH/src",
		}
		for _, ipath := range includePaths {
			args = append(args, "-I="+ipath)
		}

		for _, pbFile := range files {
			if pbFile.IsDir() || !strings.HasSuffix(pbFile.Name(), ".proto") {
				continue
			}

			pbArgs := append(args, "--gogoslick_out=plugins=grpc:.")
			pbArgs = append(pbArgs, pbFile.Name())
			if err := ForkDir(cmd, pbArgs, absPbDir); err != nil {
				log.Fatalf("Couldn't run protoc on %s: %v\n", pbFile.Name(), err)
			}

			// Delete pesky "import gogoproto" line
			pbGoFile := strings.Replace(pbFile.Name(), ".proto", ".pb.go", -1)
			protoCleanup(cmd, filepath.Join(absPbDir, pbGoFile))

			// build proto doc
			pbMdFile := strings.Replace(pbFile.Name(), ".proto", ".md", -1)
			docArgs := append(args, "--doc_out="+pbDocPath)
			docArgs = append(docArgs, "--doc_opt=markdown,"+pbMdFile)
			docArgs = append(docArgs, pbFile.Name())
			if err := ForkDir(cmd, docArgs, absPbDir); err != nil {
				log.Fatalf("Couldn't run protoc on %s: %v\n", pbFile.Name(), err)
			}
			// protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf --doc_out=. --doc_opt=markdown,access-control.md access-control.proto

			// build mock
			if mockAbsPath != "" {
				fn := strings.Replace(pbFile.Name(), ".proto", ".mock.go", -1)
				mockDest := filepath.Join(mockAbsPath, fn)
				mockArgs := []string{
					"mockgen",
					"-source",
					filepath.Join(absPbDir, pbGoFile),
					"-destination",
					mockDest,
				}
				if err := ForkDir(cmd, mockArgs, mockAbsPath); err != nil {
					log.Fatal("failed to run mockgen")
				}
			}
		}
	}
}

func protoCleanup(cmd *cobra.Command, pbGoPath string) {
	if Verbose(cmd) {
		fmt.Println("== Cleaning up after protoc build")
	}

	subs := []string{
		`s golang/protobuf/ptypes/[[:alnum:]-]* gogo/protobuf/types g`,
		`s ^import\ _\ "gogoproto\"$  g`,
		`s _\ "gogoproto\"$  g`,
	}

	args := []string{"sed", "-i.bak", "-E", "regex goes here", pbGoPath}
	for _, sub := range subs {
		args[3] = sub
		if err := Fork(cmd, args); err != nil {
			log.Fatalln("sed failed:", err)
		}
		if err := Fork(cmd, []string{"rm", pbGoPath + ".bak"}); err != nil {
			log.Fatalln("rm failed:", err)
		}
	}
}

func mkdirIfNotExist(path string) error {
	if _, err := os.Stat(path); err == nil {
		return nil
	}

	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}

	return nil
}

// XOCommand is command parameters for xo tool
type XOCommand struct {
	DSN          string
	TemplatePath string
	OutputPath   string
	Args         []string
}

// BuildXO builds db into the .xo.go files
func BuildXO(cmd *cobra.Command, xoPath, sqlPath string, c XOCommand) {
	if Verbose(cmd) {
		fmt.Println("== Running xo build")
	}

	Install(cmd, "xo", xoPath)

	// step 1: Migrate up
	dsns := strings.Split(c.DSN, ",")
	for _, dsn := range dsns {
		db, err := dburl.Parse(dsn)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("migrate up:", dsn)
		migrateUp(cmd, dsn, path.Join(sqlPath, db.Driver))
	}

	// step 2: run xo generator
	buildXO(cmd, c)
}

func migrateUp(cmd *cobra.Command, dsn, sqlPath string) {
	if Verbose(cmd) {
		fmt.Println("== Running up migrations on", dsn)
	}

	args := []string{"migrate"}

	if Verbose(cmd) {
		args = append(args, "-verbose")
	}

	args = append(args, "-path")
	args = append(args, ".")
	args = append(args, "-database")
	args = append(args, dsn)
	args = append(args, "up")

	if err := ForkDir(cmd, args, sqlPath); err != nil {
		log.Fatalf("Failed to run migrations on dsn %s: %v", dsn, err)
	}
}

func buildXO(cmd *cobra.Command, c XOCommand) {
	// create a temporary dir to store .xo.go
	tmpDir, err := ioutil.TempDir(c.OutputPath, "buildxo")
	if err != nil {
		log.Fatalln("Couldn't create temporary directory:", err)
	}
	defer os.RemoveAll(tmpDir)

	if err := xoGen(cmd, tmpDir, c); err != nil {
		// remove temporary dir
		return
	}

	// move xo.go to output
	if Verbose(cmd) {
		fmt.Println("== moving generated xo back to:", c.OutputPath)
	}

	files, err := ioutil.ReadDir(tmpDir)
	if err != nil {
		fmt.Println("Can't read temporary directory:", err)
		return
	}

	oldFiles, err := ioutil.ReadDir(c.OutputPath)
	if err != nil {
		fmt.Println("Can't read output directory:", err)
		return
	}

	// remove all xo.go
	for _, fi := range oldFiles {
		if fi.IsDir() || !strings.HasSuffix(fi.Name(), ".xo.go") {
			continue
		}

		filePath := filepath.Join(c.OutputPath, fi.Name())
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
		outputFilePath := filepath.Join(c.OutputPath, fi.Name())

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

func xoGen(cmd *cobra.Command, tmpDir string, c XOCommand) error {
	args := []string{"xo"}
	if Verbose(cmd) {
		fmt.Println("== Running xo build on dsn ", c.DSN, " output ", c.OutputPath)
		args = append(args, "-v")
	}

	args = append(args, c.DSN)
	args = append(args, "--template-path")
	args = append(args, c.TemplatePath)
	args = append(args, "-o")
	args = append(args, ".")
	args = append(args, c.Args...)

	if err := ForkDir(cmd, args, tmpDir); err != nil {
		logrus.WithFields(logrus.Fields{
			"args":   args,
			"output": tmpDir,
		}).Infoln("Failed to generate xo code")
		return errors.New("Failed to generate xo code")
	}

	return nil
}
