// Package version oututs the version.
package createrole

import (
	//"fmt"
	"errors"
	"github.com/tj/cobra"
	"github.com/aws/aws-sdk-go/service/iam"

  "github.com/ssh2003/apex/roleinit"
	"github.com/ssh2003/apex/cmd/apex/root"

)

var credentialsError = `

  AWS region missing, are your credentials set up? Try:

  $ export AWS_PROFILE=myapp-stage
  $ apex init

  Visit http://apex.run/#aws-credentials for more details on
  setting up AWS credentials and specifying which profile to
  use.

`

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
func run(c *cobra.Command, args []string) error {
	if err := root.Prepare(c, args); err != nil {
		return err
	}

	region := root.Config.Region
	if region == nil {
		return errors.New(credentialsError)
	}

	r := roleinit.RoleInit {
		IAM:    iam.New(root.Session),
		Region: *region,
	}
	return r.RInit()

}
