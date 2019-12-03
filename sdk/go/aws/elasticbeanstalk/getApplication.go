// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticbeanstalk

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Retrieve information about an Elastic Beanstalk Application.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/elastic_beanstalk_application.html.markdown.
func LookupApplication(ctx *pulumi.Context, args *GetApplicationArgs, opts ...pulumi.InvokeOption) (*GetApplicationResult, error) {
	var rv GetApplicationResult
	err := ctx.Invoke("aws:elasticbeanstalk/getApplication:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApplication.
type GetApplicationArgs struct {
	// The name of the application
	Name string `pulumi:"name"`
}

// A collection of values returned by getApplication.
type GetApplicationResult struct {
	AppversionLifecycle GetApplicationAppversionLifecycleResult `pulumi:"appversionLifecycle"`
	// The Amazon Resource Name (ARN) of the application.
	Arn string `pulumi:"arn"`
	// Short description of the application
	Description string `pulumi:"description"`
	Name string `pulumi:"name"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
type GetApplicationAppversionLifecycleResult struct {
	// Specifies whether delete a version's source bundle from S3 when the application version is deleted.
	DeleteSourceFromS3 bool `pulumi:"deleteSourceFromS3"`
	// The number of days to retain an application version.
	MaxAgeInDays int `pulumi:"maxAgeInDays"`
	// The maximum number of application versions to retain.
	MaxCount int `pulumi:"maxCount"`
	// The ARN of an IAM service role under which the application version is deleted.  Elastic Beanstalk must have permission to assume this role.
	ServiceRole string `pulumi:"serviceRole"`
}
