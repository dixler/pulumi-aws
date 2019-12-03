// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecr

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The ECR Repository data source allows the ARN, Repository URI and Registry ID to be retrieved for an ECR repository.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ecr_repository.html.markdown.
func LookupRepository(ctx *pulumi.Context, args *GetRepositoryArgs, opts ...pulumi.InvokeOption) (*GetRepositoryResult, error) {
	var rv GetRepositoryResult
	err := ctx.Invoke("aws:ecr/getRepository:getRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRepository.
type GetRepositoryArgs struct {
	// The name of the ECR Repository.
	Name string `pulumi:"name"`
	Tags *map[string]string `pulumi:"tags"`
}

// A collection of values returned by getRepository.
type GetRepositoryResult struct {
	// Full ARN of the repository.
	Arn string `pulumi:"arn"`
	Name string `pulumi:"name"`
	// The registry ID where the repository was created.
	RegistryId string `pulumi:"registryId"`
	// The URL of the repository (in the form `aws_account_id.dkr.ecr.region.amazonaws.com/repositoryName`).
	RepositoryUrl string `pulumi:"repositoryUrl"`
	// A mapping of tags assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
