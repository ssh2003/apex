// Package version oututs the version.
package createrole

import (
	"fmt"

	"github.com/tj/cobra"

	"github.com/ssh2003/apex/cmd/apex/root"
)

// Command config.
var Command = &cobra.Command{
	Use:              "createrole",
	Short:            "Create base role for lambda function",
	PersistentPreRun: root.PreRunNoop,
	Run:              run,
}

// Initialize.
func init() {
	root.Register(Command)
}

// Run command.
func run(c *cobra.Command, args []string) {
	fmt.Println("Create Role")
}
