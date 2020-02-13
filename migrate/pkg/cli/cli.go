// deprecated, please use tools package
package cli

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// GetPackagePath returns the path of golang package.
func GetPackagePath(pathShort string) string {
	if gopath := os.Getenv("GOPATH"); gopath != "" {
		return filepath.Join(gopath, "src", pathShort)
	}

	panic("Go is not installed")
}

// ForkSplit runs a command represented by a space-separated string.
func ForkSplit(c *cobra.Command, command string, env ...string) error {
	return Fork(c, strings.Split(command, " "), env...)
}

// Fork runs a command represented by a strings slice using std in/out/err,
// and searching for the first argument in PATH.
func Fork(c *cobra.Command, args []string, env ...string) error {
	if len(args) == 0 {
		return errors.New("no args")
	}

	cmd := &exec.Cmd{
		Args:   args,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	if len(env) > 0 {
		cmd.Env = append(os.Environ(), env...)
	}

	return fork(c, cmd)
}

// ForkDir runs a new command with a different working directory.
func ForkDir(c *cobra.Command, args []string, dir string, env ...string) error {
	if len(args) == 0 {
		return errors.New("no args")
	}

	cmd := &exec.Cmd{
		Args:   args,
		Dir:    dir,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	if len(env) > 0 {
		cmd.Env = append(os.Environ(), env...)
	}

	return fork(c, cmd)
}

func fork(c *cobra.Command, cmd *exec.Cmd) error {

	// Look for exec
	cmdPath, err := exec.LookPath(cmd.Args[0])
	if err != nil {
		return errors.New("command not found: " + cmd.Args[0])
	}

	// Logging
	if Verbose(c) || DryRun(c) {
		if Verbose(c) {
			fmt.Println("== Running command", cmd.Args[0])
		}
		fmt.Printf("$ {WD: %s} ", cmd.Dir)
		fmt.Println(strings.Join(cmd.Args, " "))
	}
	if DryRun(c) {
		return nil
	}

	// expand env
	for i := range cmd.Args {
		cmd.Args[i] = os.ExpandEnv(cmd.Args[i])
	}

	// Set up command path
	cmd.Path = cmdPath
	cmd.Args[0] = cmdPath

	return cmd.Run()
}

// Verbose returns the persistent flag verbose.
func Verbose(c *cobra.Command) bool {
	if c == nil {
		return false
	}

	verbose, err := c.Flags().GetBool("verbose")
	if err != nil && c.HasAvailableInheritedFlags() {
		verbose, err = c.InheritedFlags().GetBool("verbose")
	}

	if err != nil {
		log.Fatalln(err)
	}
	return verbose
}

// DryRun returns the persistent flag verbose.
func DryRun(c *cobra.Command) bool {
	if c == nil {
		return false
	}

	dryRun, err := c.Flags().GetBool("dry-run")
	if err != nil && c.HasAvailableInheritedFlags() {
		dryRun, err = c.InheritedFlags().GetBool("dry-run")
	}

	if err != nil {
		log.Fatalln(err)
	}
	return dryRun
}
