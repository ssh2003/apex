package roleinit

import (
	"fmt"
	//"io/ioutil"
	//"os"
	//"strings"
	//"github.com/aws/aws-sdk-go/aws"
	//"github.com/aws/aws-sdk-go/service/iam"
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
	return nil
}
