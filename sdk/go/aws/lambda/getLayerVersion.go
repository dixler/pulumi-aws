// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides information about a Lambda Layer Version.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/lambda_layer_version.html.markdown.
func LookupLayerVersion(ctx *pulumi.Context, args *GetLayerVersionArgs, opts ...pulumi.InvokeOption) (*GetLayerVersionResult, error) {
	var rv GetLayerVersionResult
	err := ctx.Invoke("aws:lambda/getLayerVersion:getLayerVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLayerVersion.
type GetLayerVersionArgs struct {
	// Specific runtime the layer version must support. Conflicts with `version`. If specified, the latest available layer version supporting the provided runtime will be used.
	CompatibleRuntime *string `pulumi:"compatibleRuntime"`
	// Name of the lambda layer.
	LayerName string `pulumi:"layerName"`
	// Specific layer version. Conflicts with `compatibleRuntime`. If omitted, the latest available layer version will be used.
	Version *int `pulumi:"version"`
}

// A collection of values returned by getLayerVersion.
type GetLayerVersionResult struct {
	// The Amazon Resource Name (ARN) of the Lambda Layer with version.
	Arn string `pulumi:"arn"`
	CompatibleRuntime *string `pulumi:"compatibleRuntime"`
	// A list of [Runtimes][1] the specific Lambda Layer version is compatible with.
	CompatibleRuntimes []string `pulumi:"compatibleRuntimes"`
	// The date this resource was created.
	CreatedDate string `pulumi:"createdDate"`
	// Description of the specific Lambda Layer version.
	Description string `pulumi:"description"`
	// The Amazon Resource Name (ARN) of the Lambda Layer without version.
	LayerArn string `pulumi:"layerArn"`
	LayerName string `pulumi:"layerName"`
	// License info associated with the specific Lambda Layer version.
	LicenseInfo string `pulumi:"licenseInfo"`
	// Base64-encoded representation of raw SHA-256 sum of the zip file.
	SourceCodeHash string `pulumi:"sourceCodeHash"`
	// The size in bytes of the function .zip file.
	SourceCodeSize int `pulumi:"sourceCodeSize"`
	// This Lamba Layer version.
	Version int `pulumi:"version"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
