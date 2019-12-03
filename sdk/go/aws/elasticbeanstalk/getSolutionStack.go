// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticbeanstalk

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get the name of a elastic beanstalk solution stack.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/elastic_beanstalk_solution_stack.html.markdown.
func LookupSolutionStack(ctx *pulumi.Context, args *GetSolutionStackArgs, opts ...pulumi.InvokeOption) (*GetSolutionStackResult, error) {
	var rv GetSolutionStackResult
	err := ctx.Invoke("aws:elasticbeanstalk/getSolutionStack:getSolutionStack", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSolutionStack.
type GetSolutionStackArgs struct {
	// If more than one result is returned, use the most
	// recent solution stack.
	MostRecent *bool `pulumi:"mostRecent"`
	// A regex string to apply to the solution stack list returned
	// by AWS. See [Elastic Beanstalk Supported Platforms][beanstalk-platforms] from
	// AWS documentation for reference solution stack names.
	NameRegex string `pulumi:"nameRegex"`
}

// A collection of values returned by getSolutionStack.
type GetSolutionStackResult struct {
	MostRecent *bool `pulumi:"mostRecent"`
	// The name of the solution stack.
	Name string `pulumi:"name"`
	NameRegex string `pulumi:"nameRegex"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
