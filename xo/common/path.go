package common

import (
	"os"
	"path/filepath"

	"gitlab.com/tesgo/kit/pkg/cli"
)

const (
	// PathShort is the `go get`table project name
	PathShort   = "github.com/gunsluo/go-example/xo"
	xoPathShort = "gitlab.com/tesgo/kit/pkg/xo"
)

// GetProjectPath returns the root path of reach.
func GetProjectPath() string {
	return cli.GetProjectPath(PathShort)
}

// GetXOPath returns template path of xo
func GetXOPath() string {
	xoPath := filepath.Join(GetProjectPath(), "/vendor/", xoPathShort)
	if _, err := os.Stat(xoPath); err == nil {
		return xoPath
	}

	xoPath = cli.GetProjectPath(xoPathShort)
	return xoPath
}
