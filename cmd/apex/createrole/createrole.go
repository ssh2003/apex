// Package version oututs the version.
package createrole

import (
	"fmt"

	"github.com/tj/cobra"

  "github.com/ssh2003/apex/roleinit"
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
	if err := root.Prepare(c, args); err != nil {
		return err
	}

	region := root.Config.Region
	if region == nil {
		return errors.New(credentialsError)
	}

	r := roleinit.RoleInit{
		IAM:    iam.New(root.Session),
		Region: *region,
	}
	r.RInit()

}
