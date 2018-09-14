package build

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gunsluo/go-example/xo/common"
	"github.com/spf13/cobra"
	"gitlab.com/tesgo/kit/pkg/cli"
)

// Cmd is the exported command.
var Cmd *cobra.Command
var (
	argAll    bool
	argGo     bool
	argXo     bool
	argDriver string
	argDbAddr string
	argModel  string
	//argProto  bool

	verboseXO bool
)

const (
	templatesPath  = "pgtemplates"
	extraRulesPath = "model/sql/extrarules.yaml"
	//protoDocPath  = "docs/grpc"
	//protoPath     = "proto/pb"
)

var (
	sqlMapping = map[string]map[string]string{
		"demo": map[string]string{
			"model/sql": "model",
		},
	}
)

func init() {
	Cmd = &cobra.Command{
		Use:   "build",
		Short: "builds go code",
		Run:   build,
	}

	Cmd.Flags().BoolVarP(&argAll, "all", "a", false, "run all build tools")
	Cmd.Flags().BoolVar(&argGo, "go", true, "run go build")
	//Cmd.Flags().BoolVar(&argProto, "proto", false, "run protoc build")
	Cmd.Flags().BoolVar(&argXo, "xo", false, "run xo build")
	Cmd.Flags().StringVar(&argDriver, "driver", "postgres", "target db engine") //postgres
	Cmd.Flags().StringVar(&argDbAddr, "db", "localhost:5432", "target db address, e.g host:port")
	availableModels := make([]string, 0, len(sqlMapping))
	for k := range sqlMapping {
		availableModels = append(availableModels, k)
	}
	Cmd.Flags().StringVarP(&argModel, "model", "m", "", "the model to run xo generation on, default is all, options: "+strings.Join(availableModels, ", "))
	Cmd.Flags().BoolP("dry-run", "n", false, "don't execute anything, only show the commands")
	Cmd.Flags().BoolVar(&verboseXO, "vxo", false, "verbose xo logging")
}

func build(cmd *cobra.Command, _ []string) {
	if argAll {
		argGo = true
		argXo = true
		//argProto = true
	}

	//if argProto {
	//	cli.BuildProto(cmd, common.GetProjectPath(), protoDocPath, []string{protoPath}, []string{"../"})
	//}
	if argXo {
		buildXO(cmd)
	}
	if argGo {
		cli.BuildGo(cmd, filepath.Join(common.PathShort, "/cmd"))
	}
}

func buildXO(cmd *cobra.Command) {
	xoPath := common.GetXOPath()
	absTemplatesPath := filepath.Join(xoPath, templatesPath)
	absExtraRulesFile := filepath.Join(common.GetProjectPath(), extraRulesPath)
	sqlPathsMapping := map[string]map[string]string{}
	for dbname, vmap := range sqlMapping {
		if argModel != "" && dbname != argModel {
			continue
		}
		nmap := make(map[string]string, len(vmap))
		for k, v := range vmap {
			k = filepath.Join(common.GetProjectPath(), k)
			v = filepath.Join(common.GetProjectPath(), v)
			nmap[k] = v
		}
		sqlPathsMapping[dbname] = nmap
	}

	fmt.Println("-->", absTemplatesPath)
	cli.BuildXO(cmd, xoPath, absTemplatesPath, absExtraRulesFile, sqlPathsMapping, argDriver, argDbAddr, verboseXO)
}
