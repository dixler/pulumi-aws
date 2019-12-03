// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancing

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the Account ID of the [AWS Elastic Load Balancing Service Account](http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-access-logs.html#attach-bucket-policy)
// in a given region for the purpose of whitelisting in S3 bucket policy.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/elb_service_account_legacy.html.markdown.
func LookupServiceAccount(ctx *pulumi.Context, args *GetServiceAccountArgs, opts ...pulumi.InvokeOption) (*GetServiceAccountResult, error) {
	var rv GetServiceAccountResult
	err := ctx.Invoke("aws:elasticloadbalancing/getServiceAccount:getServiceAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServiceAccount.
type GetServiceAccountArgs struct {
	// Name of the region whose AWS ELB account ID is desired.
	// Defaults to the region from the AWS provider configuration.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getServiceAccount.
type GetServiceAccountResult struct {
	// The ARN of the AWS ELB service account in the selected region.
	Arn string `pulumi:"arn"`
	Region *string `pulumi:"region"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
