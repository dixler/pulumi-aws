// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// [IPv6 only] Creates an egress-only Internet gateway for your VPC.
// An egress-only Internet gateway is used to enable outbound communication
// over IPv6 from instances in your VPC to the Internet, and prevents hosts
// outside of your VPC from initiating an IPv6 connection with your instance.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/egress_only_internet_gateway.html.markdown.
type EgressOnlyInternetGateway struct {
	pulumi.CustomResourceState

	// The VPC ID to create in.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewEgressOnlyInternetGateway registers a new resource with the given unique name, arguments, and options.
func NewEgressOnlyInternetGateway(ctx *pulumi.Context,
	name string, args *EgressOnlyInternetGatewayArgs, opts ...pulumi.ResourceOption) (*EgressOnlyInternetGateway, error) {
	if args == nil || args.VpcId == nil {
		return nil, errors.New("missing required argument 'VpcId'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.VpcId; i != nil { inputs["vpcId"] = i.ToStringOutput() }
	}
	var resource EgressOnlyInternetGateway
	err := ctx.RegisterResource("aws:ec2/egressOnlyInternetGateway:EgressOnlyInternetGateway", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEgressOnlyInternetGateway gets an existing EgressOnlyInternetGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEgressOnlyInternetGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EgressOnlyInternetGatewayState, opts ...pulumi.ResourceOption) (*EgressOnlyInternetGateway, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.VpcId; i != nil { inputs["vpcId"] = i.ToStringOutput() }
	}
	var resource EgressOnlyInternetGateway
	err := ctx.ReadResource("aws:ec2/egressOnlyInternetGateway:EgressOnlyInternetGateway", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EgressOnlyInternetGateway resources.
type EgressOnlyInternetGatewayState struct {
	// The VPC ID to create in.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
}

// The set of arguments for constructing a EgressOnlyInternetGateway resource.
type EgressOnlyInternetGatewayArgs struct {
	// The VPC ID to create in.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
}
