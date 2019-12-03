// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a [Route53 Delegation Set](https://docs.aws.amazon.com/Route53/latest/APIReference/API-actions-by-function.html#actions-by-function-reusable-delegation-sets) resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/route53_delegation_set.html.markdown.
type DelegationSet struct {
	pulumi.CustomResourceState

	// A list of authoritative name servers for the hosted zone
	// (effectively a list of NS records).
	NameServers pulumi.StringArrayOutput `pulumi:"nameServers"`

	// This is a reference name used in Caller Reference
	// (helpful for identifying single delegation set amongst others)
	ReferenceName pulumi.StringOutput `pulumi:"referenceName"`
}

// NewDelegationSet registers a new resource with the given unique name, arguments, and options.
func NewDelegationSet(ctx *pulumi.Context,
	name string, args *DelegationSetArgs, opts ...pulumi.ResourceOption) (*DelegationSet, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.ReferenceName; i != nil { inputs["referenceName"] = i.ToStringOutput() }
	}
	var resource DelegationSet
	err := ctx.RegisterResource("aws:route53/delegationSet:DelegationSet", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDelegationSet gets an existing DelegationSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDelegationSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DelegationSetState, opts ...pulumi.ResourceOption) (*DelegationSet, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.NameServers; i != nil { inputs["nameServers"] = i.ToStringArrayOutput() }
		if i := state.ReferenceName; i != nil { inputs["referenceName"] = i.ToStringOutput() }
	}
	var resource DelegationSet
	err := ctx.ReadResource("aws:route53/delegationSet:DelegationSet", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DelegationSet resources.
type DelegationSetState struct {
	// A list of authoritative name servers for the hosted zone
	// (effectively a list of NS records).
	NameServers pulumi.StringArrayInput `pulumi:"nameServers"`
	// This is a reference name used in Caller Reference
	// (helpful for identifying single delegation set amongst others)
	ReferenceName pulumi.StringInput `pulumi:"referenceName"`
}

// The set of arguments for constructing a DelegationSet resource.
type DelegationSetArgs struct {
	// This is a reference name used in Caller Reference
	// (helpful for identifying single delegation set amongst others)
	ReferenceName pulumi.StringInput `pulumi:"referenceName"`
}
