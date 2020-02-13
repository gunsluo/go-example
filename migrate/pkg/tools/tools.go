package tools

import (
	"log"

	"github.com/spf13/cobra"
)

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
