package roleinit

import (
	"fmt"
	//"io/ioutil"
	"os"
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

func (r *RoleInit) RInit() error {
	fmt.Println("RInit")
	fmt.Println(os.Args)
	result, err := r.IAM.ListUsers(&iam.ListUsersInput{
        MaxItems: aws.Int64(10),
    })
	if err != nil {
	        fmt.Println("Error", err)
	        return err
	}
	for i, user := range result.Users {
        if user == nil {
            continue
        }
        fmt.Printf("%d user %s created %v\n", i, *user.UserName, user.CreateDate)
  }
	return nil
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
