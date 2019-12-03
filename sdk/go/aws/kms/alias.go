// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an alias for a KMS customer master key. AWS Console enforces 1-to-1 mapping between aliases & keys,
// but API (hence this provider too) allows you to create as many aliases as
// the [account limits](http://docs.aws.amazon.com/kms/latest/developerguide/limits.html) allow you.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/kms_alias.html.markdown.
type Alias struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the key alias.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The display name of the alias. The name must start with the word "alias" followed by a forward slash (alias/)
	Name pulumi.StringOutput `pulumi:"name"`

	// Creates an unique alias beginning with the specified prefix.
	// The name must start with the word "alias" followed by a forward slash (alias/).  Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`

	// The Amazon Resource Name (ARN) of the target key identifier.
	TargetKeyArn pulumi.StringOutput `pulumi:"targetKeyArn"`

	// Identifier for the key for which the alias is for, can be either an ARN or key_id.
	TargetKeyId pulumi.StringOutput `pulumi:"targetKeyId"`
}

// NewAlias registers a new resource with the given unique name, arguments, and options.
func NewAlias(ctx *pulumi.Context,
	name string, args *AliasArgs, opts ...pulumi.ResourceOption) (*Alias, error) {
	if args == nil || args.TargetKeyId == nil {
		return nil, errors.New("missing required argument 'TargetKeyId'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.NamePrefix; i != nil { inputs["namePrefix"] = i.ToStringOutput() }
		if i := args.TargetKeyId; i != nil { inputs["targetKeyId"] = i.ToStringOutput() }
	}
	var resource Alias
	err := ctx.RegisterResource("aws:kms/alias:Alias", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlias gets an existing Alias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AliasState, opts ...pulumi.ResourceOption) (*Alias, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Arn; i != nil { inputs["arn"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.NamePrefix; i != nil { inputs["namePrefix"] = i.ToStringOutput() }
		if i := state.TargetKeyArn; i != nil { inputs["targetKeyArn"] = i.ToStringOutput() }
		if i := state.TargetKeyId; i != nil { inputs["targetKeyId"] = i.ToStringOutput() }
	}
	var resource Alias
	err := ctx.ReadResource("aws:kms/alias:Alias", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Alias resources.
type AliasState struct {
	// The Amazon Resource Name (ARN) of the key alias.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The display name of the alias. The name must start with the word "alias" followed by a forward slash (alias/)
	Name pulumi.StringInput `pulumi:"name"`
	// Creates an unique alias beginning with the specified prefix.
	// The name must start with the word "alias" followed by a forward slash (alias/).  Conflicts with `name`.
	NamePrefix pulumi.StringInput `pulumi:"namePrefix"`
	// The Amazon Resource Name (ARN) of the target key identifier.
	TargetKeyArn pulumi.StringInput `pulumi:"targetKeyArn"`
	// Identifier for the key for which the alias is for, can be either an ARN or key_id.
	TargetKeyId pulumi.StringInput `pulumi:"targetKeyId"`
}

// The set of arguments for constructing a Alias resource.
type AliasArgs struct {
	// The display name of the alias. The name must start with the word "alias" followed by a forward slash (alias/)
	Name pulumi.StringInput `pulumi:"name"`
	// Creates an unique alias beginning with the specified prefix.
	// The name must start with the word "alias" followed by a forward slash (alias/).  Conflicts with `name`.
	NamePrefix pulumi.StringInput `pulumi:"namePrefix"`
	// Identifier for the key for which the alias is for, can be either an ARN or key_id.
	TargetKeyId pulumi.StringInput `pulumi:"targetKeyId"`
}
