package roleinit

import (
	"fmt"

	//"strings"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
)

type RoleInit struct {
	IAM    iamiface.IAMAPI
	Region string

	name        string
	description string
}

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


func (r *RoleInit) RInit(name string) error {
	fmt.Println("RInit")
	fmt.Println(name)

	_, err := r.createRole(name)
	return err
}

func (r *RoleInit) createRole(name string) (string, error) {
	roleName := fmt.Sprintf("%s_lambda_function", name)
	policyName := fmt.Sprintf("%s_lambda_logs", name)

	logf("creating IAM %s role", roleName)
	role, err := r.IAM.CreateRole(&iam.CreateRoleInput{
		RoleName:                 &roleName,
		AssumeRolePolicyDocument: aws.String(iamAssumeRolePolicy),
	})

	if err != nil {
		return "", fmt.Errorf("creating role: %s", err)
	}

	logf("creating IAM %s policy", policyName)
	policy, err := r.IAM.CreatePolicy(&iam.CreatePolicyInput{
		PolicyName:     &policyName,
		Description:    aws.String("Allow lambda_function to utilize CloudWatchLogs. Created by apex(1)."),
		PolicyDocument: aws.String(iamLogsPolicy),
	})

	if err != nil {
		return "", fmt.Errorf("creating policy: %s", err)
	}

	logf("attaching policy to lambda_function role.")
	_, err = r.IAM.AttachRolePolicy(&iam.AttachRolePolicyInput{
		RoleName:  &roleName,
		PolicyArn: policy.Policy.Arn,
	})

	if err != nil {
		return "", fmt.Errorf("creating policy: %s", err)
	}

	return *role.Role.Arn, nil
}

func logf(s string, v ...interface{}) {
	fmt.Printf("  \033[34m[+]\033[0m %s\n", fmt.Sprintf(s, v...))
}

/*func (b *Bootstrapper) createRole() (string, error) {
	roleName := fmt.Sprintf("%s_lambda_function", b.name)
	policyName := fmt.Sprintf("%s_lambda_logs", b.name)

	logf("creating IAM %s role", roleName)
	role, err := b.IAM.CreateRole(&iam.CreateRoleInput{
		RoleName:                 &roleName,
		AssumeRolePolicyDocument: aws.String(iamAssumeRolePolicy),
	})

	if err != nil {
		return "", fmt.Errorf("creating role: %s", err)
	}

	logf("creating IAM %s policy", policyName)
	policy, err := b.IAM.CreatePolicy(&iam.CreatePolicyInput{
		PolicyName:     &policyName,
		Description:    aws.String("Allow lambda_function to utilize CloudWatchLogs. Created by apex(1)."),
		PolicyDocument: aws.String(iamLogsPolicy),
	})

	if err != nil {
		return "", fmt.Errorf("creating policy: %s", err)
	}

	logf("attaching policy to lambda_function role.")
	_, err = b.IAM.AttachRolePolicy(&iam.AttachRolePolicyInput{
		RoleName:  &roleName,
		PolicyArn: policy.Policy.Arn,
	})

	if err != nil {
		return "", fmt.Errorf("creating policy: %s", err)
	}

	return *role.Role.Arn, nil
}
*/
