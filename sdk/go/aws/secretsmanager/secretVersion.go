// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package secretsmanager

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage AWS Secrets Manager secret version including its secret value. To manage secret metadata, see the [`secretsmanager.Secret` resource](https://www.terraform.io/docs/providers/aws/r/secretsmanager_secret.html).
// 
// > **NOTE:** If the `AWSCURRENT` staging label is present on this version during resource deletion, that label cannot be removed and will be skipped to prevent errors when fully deleting the secret. That label will leave this secret version active even after the resource is deleted from this provider unless the secret itself is deleted. Move the `AWSCURRENT` staging label before or after deleting this resource from this provider to fully trigger version deprecation if necessary.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/secretsmanager_secret_version.html.markdown.
type SecretVersion struct {
	pulumi.CustomResourceState

	// The ARN of the secret.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// Specifies binary data that you want to encrypt and store in this version of the secret. This is required if secretString is not set. Needs to be encoded to base64.
	SecretBinary pulumi.StringOutput `pulumi:"secretBinary"`

	// Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
	SecretId pulumi.StringOutput `pulumi:"secretId"`

	// Specifies text data that you want to encrypt and store in this version of the secret. This is required if secretBinary is not set.
	SecretString pulumi.StringOutput `pulumi:"secretString"`

	// The unique identifier of the version of the secret.
	VersionId pulumi.StringOutput `pulumi:"versionId"`

	// Specifies a list of staging labels that are attached to this version of the secret. A staging label must be unique to a single version of the secret. If you specify a staging label that's already associated with a different version of the same secret then that staging label is automatically removed from the other version and attached to this version. If you do not specify a value, then AWS Secrets Manager automatically moves the staging label `AWSCURRENT` to this new version on creation.
	VersionStages pulumi.StringArrayOutput `pulumi:"versionStages"`
}

// NewSecretVersion registers a new resource with the given unique name, arguments, and options.
func NewSecretVersion(ctx *pulumi.Context,
	name string, args *SecretVersionArgs, opts ...pulumi.ResourceOption) (*SecretVersion, error) {
	if args == nil || args.SecretId == nil {
		return nil, errors.New("missing required argument 'SecretId'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.SecretBinary; i != nil { inputs["secretBinary"] = i.ToStringOutput() }
		if i := args.SecretId; i != nil { inputs["secretId"] = i.ToStringOutput() }
		if i := args.SecretString; i != nil { inputs["secretString"] = i.ToStringOutput() }
		if i := args.VersionStages; i != nil { inputs["versionStages"] = i.ToStringArrayOutput() }
	}
	var resource SecretVersion
	err := ctx.RegisterResource("aws:secretsmanager/secretVersion:SecretVersion", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretVersion gets an existing SecretVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretVersionState, opts ...pulumi.ResourceOption) (*SecretVersion, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Arn; i != nil { inputs["arn"] = i.ToStringOutput() }
		if i := state.SecretBinary; i != nil { inputs["secretBinary"] = i.ToStringOutput() }
		if i := state.SecretId; i != nil { inputs["secretId"] = i.ToStringOutput() }
		if i := state.SecretString; i != nil { inputs["secretString"] = i.ToStringOutput() }
		if i := state.VersionId; i != nil { inputs["versionId"] = i.ToStringOutput() }
		if i := state.VersionStages; i != nil { inputs["versionStages"] = i.ToStringArrayOutput() }
	}
	var resource SecretVersion
	err := ctx.ReadResource("aws:secretsmanager/secretVersion:SecretVersion", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretVersion resources.
type SecretVersionState struct {
	// The ARN of the secret.
	Arn pulumi.StringInput `pulumi:"arn"`
	// Specifies binary data that you want to encrypt and store in this version of the secret. This is required if secretString is not set. Needs to be encoded to base64.
	SecretBinary pulumi.StringInput `pulumi:"secretBinary"`
	// Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
	SecretId pulumi.StringInput `pulumi:"secretId"`
	// Specifies text data that you want to encrypt and store in this version of the secret. This is required if secretBinary is not set.
	SecretString pulumi.StringInput `pulumi:"secretString"`
	// The unique identifier of the version of the secret.
	VersionId pulumi.StringInput `pulumi:"versionId"`
	// Specifies a list of staging labels that are attached to this version of the secret. A staging label must be unique to a single version of the secret. If you specify a staging label that's already associated with a different version of the same secret then that staging label is automatically removed from the other version and attached to this version. If you do not specify a value, then AWS Secrets Manager automatically moves the staging label `AWSCURRENT` to this new version on creation.
	VersionStages pulumi.StringArrayInput `pulumi:"versionStages"`
}

// The set of arguments for constructing a SecretVersion resource.
type SecretVersionArgs struct {
	// Specifies binary data that you want to encrypt and store in this version of the secret. This is required if secretString is not set. Needs to be encoded to base64.
	SecretBinary pulumi.StringInput `pulumi:"secretBinary"`
	// Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
	SecretId pulumi.StringInput `pulumi:"secretId"`
	// Specifies text data that you want to encrypt and store in this version of the secret. This is required if secretBinary is not set.
	SecretString pulumi.StringInput `pulumi:"secretString"`
	// Specifies a list of staging labels that are attached to this version of the secret. A staging label must be unique to a single version of the secret. If you specify a staging label that's already associated with a different version of the same secret then that staging label is automatically removed from the other version and attached to this version. If you do not specify a value, then AWS Secrets Manager automatically moves the staging label `AWSCURRENT` to this new version on creation.
	VersionStages pulumi.StringArrayInput `pulumi:"versionStages"`
}
