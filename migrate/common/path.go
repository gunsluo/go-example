package common

import (
	"os"
	"path/filepath"

	"github.com/gunsluo/go-example/migrate/pkg/cli"
)

const (
	// ProjectPathShort is the `go get`table project name
	PathShort   = "github.com/gunsluo/go-example/migrate"
	xoPathShort = "gitlab.com/target-digital-transformation/kit/pkg/xo"
)

// GetProjectPath returns the root path of reach.
func GetProjectPath() string {
	return cli.GetPackagePath(PathShort)
}

func GetXOPath() string {
	xoPath := filepath.Join(GetProjectPath(), "/vendor/", xoPathShort)
	if _, err := os.Stat(xoPath); err == nil {
		return xoPath
	}

	xoPath = cli.GetPackagePath(xoPathShort)
	return xoPath
}
