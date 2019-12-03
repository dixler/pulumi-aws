// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package macie

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Associates an AWS account with Amazon Macie as a member account.
// 
// > **NOTE:** Before using Amazon Macie for the first time it must be enabled manually. Instructions are [here](https://docs.aws.amazon.com/macie/latest/userguide/macie-setting-up.html#macie-setting-up-enable).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/macie_member_account_association.html.markdown.
type MemberAccountAssociation struct {
	pulumi.CustomResourceState

	// The ID of the AWS account that you want to associate with Amazon Macie as a member account.
	MemberAccountId pulumi.StringOutput `pulumi:"memberAccountId"`
}

// NewMemberAccountAssociation registers a new resource with the given unique name, arguments, and options.
func NewMemberAccountAssociation(ctx *pulumi.Context,
	name string, args *MemberAccountAssociationArgs, opts ...pulumi.ResourceOption) (*MemberAccountAssociation, error) {
	if args == nil || args.MemberAccountId == nil {
		return nil, errors.New("missing required argument 'MemberAccountId'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.MemberAccountId; i != nil { inputs["memberAccountId"] = i.ToStringOutput() }
	}
	var resource MemberAccountAssociation
	err := ctx.RegisterResource("aws:macie/memberAccountAssociation:MemberAccountAssociation", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMemberAccountAssociation gets an existing MemberAccountAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMemberAccountAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MemberAccountAssociationState, opts ...pulumi.ResourceOption) (*MemberAccountAssociation, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.MemberAccountId; i != nil { inputs["memberAccountId"] = i.ToStringOutput() }
	}
	var resource MemberAccountAssociation
	err := ctx.ReadResource("aws:macie/memberAccountAssociation:MemberAccountAssociation", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MemberAccountAssociation resources.
type MemberAccountAssociationState struct {
	// The ID of the AWS account that you want to associate with Amazon Macie as a member account.
	MemberAccountId pulumi.StringInput `pulumi:"memberAccountId"`
}

// The set of arguments for constructing a MemberAccountAssociation resource.
type MemberAccountAssociationArgs struct {
	// The ID of the AWS account that you want to associate with Amazon Macie as a member account.
	MemberAccountId pulumi.StringInput `pulumi:"memberAccountId"`
}
