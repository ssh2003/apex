// Package version oututs the version.
package version

import (
	"fmt"

	"github.com/tj/cobra"

	"github.com/ssh2003/apex/cmd/apex/root"
)

// Version of program.
const Version = "1.0.0-ssh2003-alpha0"

// Command config.
var Command = &cobra.Command{
	Use:              "version",
	Short:            "Print version of Apex",
	PersistentPreRun: root.PreRunNoop,
	Run:              run,
}

// Initialize.
func init() {
	root.Register(Command)
}

// Run command.
func run(c *cobra.Command, args []string) {
	fmt.Printf("Apex version %s\n", Version)
}
