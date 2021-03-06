// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package iot

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an IoT role alias.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iot_role_alias.html.markdown.
type RoleAlias struct {
	pulumi.CustomResourceState

	// The name of the role alias.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// The ARN assigned by AWS to this role alias.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The duration of the credential, in seconds. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 900 seconds (15 minutes) to 3600 seconds (60 minutes).
	CredentialDuration pulumi.IntPtrOutput `pulumi:"credentialDuration"`
	// The identity of the role to which the alias refers.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
}

// NewRoleAlias registers a new resource with the given unique name, arguments, and options.
func NewRoleAlias(ctx *pulumi.Context,
	name string, args *RoleAliasArgs, opts ...pulumi.ResourceOption) (*RoleAlias, error) {
	if args == nil || args.Alias == nil {
		return nil, errors.New("missing required argument 'Alias'")
	}
	if args == nil || args.RoleArn == nil {
		return nil, errors.New("missing required argument 'RoleArn'")
	}
	if args == nil {
		args = &RoleAliasArgs{}
	}
	var resource RoleAlias
	err := ctx.RegisterResource("aws:iot/roleAlias:RoleAlias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoleAlias gets an existing RoleAlias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoleAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleAliasState, opts ...pulumi.ResourceOption) (*RoleAlias, error) {
	var resource RoleAlias
	err := ctx.ReadResource("aws:iot/roleAlias:RoleAlias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoleAlias resources.
type roleAliasState struct {
	// The name of the role alias.
	Alias *string `pulumi:"alias"`
	// The ARN assigned by AWS to this role alias.
	Arn *string `pulumi:"arn"`
	// The duration of the credential, in seconds. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 900 seconds (15 minutes) to 3600 seconds (60 minutes).
	CredentialDuration *int `pulumi:"credentialDuration"`
	// The identity of the role to which the alias refers.
	RoleArn *string `pulumi:"roleArn"`
}

type RoleAliasState struct {
	// The name of the role alias.
	Alias pulumi.StringPtrInput
	// The ARN assigned by AWS to this role alias.
	Arn pulumi.StringPtrInput
	// The duration of the credential, in seconds. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 900 seconds (15 minutes) to 3600 seconds (60 minutes).
	CredentialDuration pulumi.IntPtrInput
	// The identity of the role to which the alias refers.
	RoleArn pulumi.StringPtrInput
}

func (RoleAliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAliasState)(nil)).Elem()
}

type roleAliasArgs struct {
	// The name of the role alias.
	Alias string `pulumi:"alias"`
	// The duration of the credential, in seconds. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 900 seconds (15 minutes) to 3600 seconds (60 minutes).
	CredentialDuration *int `pulumi:"credentialDuration"`
	// The identity of the role to which the alias refers.
	RoleArn string `pulumi:"roleArn"`
}

// The set of arguments for constructing a RoleAlias resource.
type RoleAliasArgs struct {
	// The name of the role alias.
	Alias pulumi.StringInput
	// The duration of the credential, in seconds. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 900 seconds (15 minutes) to 3600 seconds (60 minutes).
	CredentialDuration pulumi.IntPtrInput
	// The identity of the role to which the alias refers.
	RoleArn pulumi.StringInput
}

func (RoleAliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAliasArgs)(nil)).Elem()
}

