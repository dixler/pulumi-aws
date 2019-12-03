// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ebs

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage whether default EBS encryption is enabled for your AWS account in the current AWS region. To manage the default KMS key for the region, see the [`ebs.DefaultKmsKey` resource](https://www.terraform.io/docs/providers/aws/r/ebs_default_kms_key.html).
// 
// > **NOTE:** Removing this resource disables default EBS encryption.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ebs_encryption_by_default.html.markdown.
type EncryptionByDefault struct {
	pulumi.CustomResourceState

	// Whether or not default EBS encryption is enabled. Valid values are `true` or `false`. Defaults to `true`.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
}

// NewEncryptionByDefault registers a new resource with the given unique name, arguments, and options.
func NewEncryptionByDefault(ctx *pulumi.Context,
	name string, args *EncryptionByDefaultArgs, opts ...pulumi.ResourceOption) (*EncryptionByDefault, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Enabled; i != nil { inputs["enabled"] = i.ToBoolOutput() }
	}
	var resource EncryptionByDefault
	err := ctx.RegisterResource("aws:ebs/encryptionByDefault:EncryptionByDefault", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEncryptionByDefault gets an existing EncryptionByDefault resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEncryptionByDefault(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EncryptionByDefaultState, opts ...pulumi.ResourceOption) (*EncryptionByDefault, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Enabled; i != nil { inputs["enabled"] = i.ToBoolOutput() }
	}
	var resource EncryptionByDefault
	err := ctx.ReadResource("aws:ebs/encryptionByDefault:EncryptionByDefault", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EncryptionByDefault resources.
type EncryptionByDefaultState struct {
	// Whether or not default EBS encryption is enabled. Valid values are `true` or `false`. Defaults to `true`.
	Enabled pulumi.BoolInput `pulumi:"enabled"`
}

// The set of arguments for constructing a EncryptionByDefault resource.
type EncryptionByDefaultArgs struct {
	// Whether or not default EBS encryption is enabled. Valid values are `true` or `false`. Defaults to `true`.
	Enabled pulumi.BoolInput `pulumi:"enabled"`
}
