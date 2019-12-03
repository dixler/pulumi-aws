// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This data source can be used to fetch information about a specific
// IAM user. By using this data source, you can reference IAM user
// properties without having to hard code ARNs or unique IDs as input.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/iam_user.html.markdown.
func LookupUser(ctx *pulumi.Context, args *GetUserArgs, opts ...pulumi.InvokeOption) (*GetUserResult, error) {
	var rv GetUserResult
	err := ctx.Invoke("aws:iam/getUser:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUser.
type GetUserArgs struct {
	// The friendly IAM user name to match.
	UserName string `pulumi:"userName"`
}

// A collection of values returned by getUser.
type GetUserResult struct {
	// The Amazon Resource Name (ARN) assigned by AWS for this user.
	Arn string `pulumi:"arn"`
	// Path in which this user was created.
	Path string `pulumi:"path"`
	// The ARN of the policy that is used to set the permissions boundary for the user.
	PermissionsBoundary string `pulumi:"permissionsBoundary"`
	// The unique ID assigned by AWS for this user.
	UserId string `pulumi:"userId"`
	// The name associated to this User
	UserName string `pulumi:"userName"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
