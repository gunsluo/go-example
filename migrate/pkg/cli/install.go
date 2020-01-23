package cli

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

type installer func(*cobra.Command, string) error

var (
	knownTools = map[string]installer{
		"xo":                   installXo,
		"protoc":               installGogoproto,
		"protoc-gen-gogoslick": installGogoproto,
		"protoc-gen-doc":       installGogoproto,
		"mockgen":              installGoMock,
	}
)

// Install will install a Go tool.
func Install(c *cobra.Command, cmd, dpath string) string {
	if Verbose(c) {
		fmt.Println("== Looking up in PATH: ", cmd)
	}
	if DryRun(c) {
		return cmd
	}

	if fpath, err := exec.LookPath(cmd); err == nil {
		return fpath
	}

	if Verbose(c) {
		fmt.Println("== Installing: ", cmd)
	}

	install, ok := knownTools[cmd]
	if !ok {
		log.Fatalln("command not installed and not known:", cmd)
	}

	if err := install(c, dpath); err != nil {
		log.Fatalf("couldn't install %s: %v\n", cmd, err)
	}

	fpath, err := exec.LookPath(cmd)
	if err != nil {
		log.Fatalf("couldn't find exec %s after installing: %v\n", cmd, err)
	}

	return fpath
}

func installGoRepo(repo string) installer {
	return func(cmd *cobra.Command, dpath string) error {
		return Fork(cmd, []string{"go", "get", "-u", repo})
	}
}

// // Could do the install here but it's not worth spending the time
// const (
// 	protocLinux = "https://github.com/google/protobuf/releases/download/v3.5.0/protoc-3.5.0-linux-x86_64.zip"
// 	protocMacOS = "https://github.com/google/protobuf/releases/download/v3.5.0/protoc-3.5.0-osx-x86_64.zip"
// )

func installXo(c *cobra.Command, dpath string) error {
	// go install -tags "mysql postgres cockroach"
	fmt.Println("===", []string{"go", "install", "-tags", "mysql postgres mssql"}, dpath)
	return ForkDir(c, []string{"go", "install", "-tags", "mysql postgres mssql"}, dpath)
}

func installGogoproto(c *cobra.Command, dpath string) error {
	var err error

	if _, err = exec.LookPath("protoc"); err != nil {
		fmt.Print(`protoc not found
Please install the protobuf compiler:
	https://github.com/golang/protobuf#installation
Extract the files then add the bin directory to your PATH.
`)
		return err
	}

	if _, err = exec.LookPath("protoc-gen-gogoslick"); err != nil {
		if err = ForkDir(c, []string{"go", "get", "-u", "github.com/gogo/protobuf/..."}, dpath); err != nil {
			return err
		}
	}

	if _, err = exec.LookPath("protoc-gen-doc"); err != nil {
		if err = ForkDir(c, []string{"go", "get", "-u", "github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc"}, dpath); err != nil {
			return err
		}
	}

	return nil
}

func installGoMock(c *cobra.Command, dpath string) error {
	var err error

	if _, err = exec.LookPath("mockgen"); err != nil {
		if err = ForkDir(c, []string{"go", "get", "-u", "github.com/golang/mock/gomock"}, dpath); err != nil {
			return err
		}
		if err = ForkDir(c, []string{"go", "install", "github.com/golang/mock/mockgen"}, dpath); err != nil {
			return err
		}
	}

	return nil
}
