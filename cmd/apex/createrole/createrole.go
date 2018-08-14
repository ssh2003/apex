// Package version oututs the version.
package createrole

import (
	"fmt"

	"github.com/tj/cobra"

	"github.com/ssh2003/apex/cmd/apex/root"

	//"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	//"github.com/aws/aws-sdk-go/service/iam/iamiface"

)
/*
var iamAssumeRolePolicy = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}`

var iamLogsPolicy = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "logs:*"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}`
*/





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
	svc := iam.New()
	fmt.Println(svc)
}
