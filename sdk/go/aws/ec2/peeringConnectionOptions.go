// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage VPC peering connection options.
// 
// > **NOTE on VPC Peering Connections and VPC Peering Connection Options:** This provider provides
// both a standalone VPC Peering Connection Options and a VPC Peering Connection
// resource with `accepter` and `requester` attributes. Do not manage options for the same VPC peering
// connection in both a VPC Peering Connection resource and a VPC Peering Connection Options resource.
// Doing so will cause a conflict of options and will overwrite the options.
// Using a VPC Peering Connection Options resource decouples management of the connection options from
// management of the VPC Peering Connection and allows options to be set correctly in cross-region and
// cross-account scenarios.
// 
// Basic usage:
// 
// 
// Basic cross-account usage:
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/vpc_peering_connection_options.html.markdown.
type PeeringConnectionOptions struct {
	pulumi.CustomResourceState

	// An optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that accepts
	// the peering connection (a maximum of one).
	Accepter PeeringConnectionOptionsAccepterOutput `pulumi:"accepter"`

	// A optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that requests
	// the peering connection (a maximum of one).
	Requester PeeringConnectionOptionsRequesterOutput `pulumi:"requester"`

	// The ID of the requester VPC peering connection.
	VpcPeeringConnectionId pulumi.StringOutput `pulumi:"vpcPeeringConnectionId"`
}

// NewPeeringConnectionOptions registers a new resource with the given unique name, arguments, and options.
func NewPeeringConnectionOptions(ctx *pulumi.Context,
	name string, args *PeeringConnectionOptionsArgs, opts ...pulumi.ResourceOption) (*PeeringConnectionOptions, error) {
	if args == nil || args.VpcPeeringConnectionId == nil {
		return nil, errors.New("missing required argument 'VpcPeeringConnectionId'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Accepter; i != nil { inputs["accepter"] = i.ToPeeringConnectionOptionsAccepterOutput() }
		if i := args.Requester; i != nil { inputs["requester"] = i.ToPeeringConnectionOptionsRequesterOutput() }
		if i := args.VpcPeeringConnectionId; i != nil { inputs["vpcPeeringConnectionId"] = i.ToStringOutput() }
	}
	var resource PeeringConnectionOptions
	err := ctx.RegisterResource("aws:ec2/peeringConnectionOptions:PeeringConnectionOptions", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPeeringConnectionOptions gets an existing PeeringConnectionOptions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPeeringConnectionOptions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PeeringConnectionOptionsState, opts ...pulumi.ResourceOption) (*PeeringConnectionOptions, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Accepter; i != nil { inputs["accepter"] = i.ToPeeringConnectionOptionsAccepterOutput() }
		if i := state.Requester; i != nil { inputs["requester"] = i.ToPeeringConnectionOptionsRequesterOutput() }
		if i := state.VpcPeeringConnectionId; i != nil { inputs["vpcPeeringConnectionId"] = i.ToStringOutput() }
	}
	var resource PeeringConnectionOptions
	err := ctx.ReadResource("aws:ec2/peeringConnectionOptions:PeeringConnectionOptions", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PeeringConnectionOptions resources.
type PeeringConnectionOptionsState struct {
	// An optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that accepts
	// the peering connection (a maximum of one).
	Accepter PeeringConnectionOptionsAccepterInput `pulumi:"accepter"`
	// A optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that requests
	// the peering connection (a maximum of one).
	Requester PeeringConnectionOptionsRequesterInput `pulumi:"requester"`
	// The ID of the requester VPC peering connection.
	VpcPeeringConnectionId pulumi.StringInput `pulumi:"vpcPeeringConnectionId"`
}

// The set of arguments for constructing a PeeringConnectionOptions resource.
type PeeringConnectionOptionsArgs struct {
	// An optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that accepts
	// the peering connection (a maximum of one).
	Accepter PeeringConnectionOptionsAccepterInput `pulumi:"accepter"`
	// A optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that requests
	// the peering connection (a maximum of one).
	Requester PeeringConnectionOptionsRequesterInput `pulumi:"requester"`
	// The ID of the requester VPC peering connection.
	VpcPeeringConnectionId pulumi.StringInput `pulumi:"vpcPeeringConnectionId"`
}
type PeeringConnectionOptionsAccepter struct {
	// Allow a local linked EC2-Classic instance to communicate
	// with instances in a peer VPC. This enables an outbound communication from the local ClassicLink connection
	// to the remote VPC. This option is not supported for inter-region VPC peering.
	AllowClassicLinkToRemoteVpc *bool `pulumi:"allowClassicLinkToRemoteVpc"`
	// Allow a local VPC to resolve public DNS hostnames to
	// private IP addresses when queried from instances in the peer VPC.
	AllowRemoteVpcDnsResolution *bool `pulumi:"allowRemoteVpcDnsResolution"`
	// Allow a local VPC to communicate with a linked EC2-Classic
	// instance in a peer VPC. This enables an outbound communication from the local VPC to the remote ClassicLink
	// connection. This option is not supported for inter-region VPC peering.
	AllowVpcToRemoteClassicLink *bool `pulumi:"allowVpcToRemoteClassicLink"`
}
var peeringConnectionOptionsAccepterType = reflect.TypeOf((*PeeringConnectionOptionsAccepter)(nil)).Elem()

type PeeringConnectionOptionsAccepterInput interface {
	pulumi.Input

	ToPeeringConnectionOptionsAccepterOutput() PeeringConnectionOptionsAccepterOutput
	ToPeeringConnectionOptionsAccepterOutputWithContext(ctx context.Context) PeeringConnectionOptionsAccepterOutput
}

type PeeringConnectionOptionsAccepterArgs struct {
	// Allow a local linked EC2-Classic instance to communicate
	// with instances in a peer VPC. This enables an outbound communication from the local ClassicLink connection
	// to the remote VPC. This option is not supported for inter-region VPC peering.
	AllowClassicLinkToRemoteVpc pulumi.BoolInput `pulumi:"allowClassicLinkToRemoteVpc"`
	// Allow a local VPC to resolve public DNS hostnames to
	// private IP addresses when queried from instances in the peer VPC.
	AllowRemoteVpcDnsResolution pulumi.BoolInput `pulumi:"allowRemoteVpcDnsResolution"`
	// Allow a local VPC to communicate with a linked EC2-Classic
	// instance in a peer VPC. This enables an outbound communication from the local VPC to the remote ClassicLink
	// connection. This option is not supported for inter-region VPC peering.
	AllowVpcToRemoteClassicLink pulumi.BoolInput `pulumi:"allowVpcToRemoteClassicLink"`
}

func (PeeringConnectionOptionsAccepterArgs) ElementType() reflect.Type {
	return peeringConnectionOptionsAccepterType
}

func (a PeeringConnectionOptionsAccepterArgs) ToPeeringConnectionOptionsAccepterOutput() PeeringConnectionOptionsAccepterOutput {
	return pulumi.ToOutput(a).(PeeringConnectionOptionsAccepterOutput)
}

func (a PeeringConnectionOptionsAccepterArgs) ToPeeringConnectionOptionsAccepterOutputWithContext(ctx context.Context) PeeringConnectionOptionsAccepterOutput {
	return pulumi.ToOutputWithContext(ctx, a).(PeeringConnectionOptionsAccepterOutput)
}

type PeeringConnectionOptionsAccepterOutput struct { *pulumi.OutputState }

// Allow a local linked EC2-Classic instance to communicate
// with instances in a peer VPC. This enables an outbound communication from the local ClassicLink connection
// to the remote VPC. This option is not supported for inter-region VPC peering.
func (o PeeringConnectionOptionsAccepterOutput) AllowClassicLinkToRemoteVpc() pulumi.BoolOutput {
	return o.Apply(func(v PeeringConnectionOptionsAccepter) bool {
		if v.AllowClassicLinkToRemoteVpc == nil { return *new(bool) } else { return *v.AllowClassicLinkToRemoteVpc }
	}).(pulumi.BoolOutput)
}

// Allow a local VPC to resolve public DNS hostnames to
// private IP addresses when queried from instances in the peer VPC.
func (o PeeringConnectionOptionsAccepterOutput) AllowRemoteVpcDnsResolution() pulumi.BoolOutput {
	return o.Apply(func(v PeeringConnectionOptionsAccepter) bool {
		if v.AllowRemoteVpcDnsResolution == nil { return *new(bool) } else { return *v.AllowRemoteVpcDnsResolution }
	}).(pulumi.BoolOutput)
}

// Allow a local VPC to communicate with a linked EC2-Classic
// instance in a peer VPC. This enables an outbound communication from the local VPC to the remote ClassicLink
// connection. This option is not supported for inter-region VPC peering.
func (o PeeringConnectionOptionsAccepterOutput) AllowVpcToRemoteClassicLink() pulumi.BoolOutput {
	return o.Apply(func(v PeeringConnectionOptionsAccepter) bool {
		if v.AllowVpcToRemoteClassicLink == nil { return *new(bool) } else { return *v.AllowVpcToRemoteClassicLink }
	}).(pulumi.BoolOutput)
}

func (PeeringConnectionOptionsAccepterOutput) ElementType() reflect.Type {
	return peeringConnectionOptionsAccepterType
}

func (o PeeringConnectionOptionsAccepterOutput) ToPeeringConnectionOptionsAccepterOutput() PeeringConnectionOptionsAccepterOutput {
	return o
}

func (o PeeringConnectionOptionsAccepterOutput) ToPeeringConnectionOptionsAccepterOutputWithContext(ctx context.Context) PeeringConnectionOptionsAccepterOutput {
	return o
}

func init() { pulumi.RegisterOutputType(PeeringConnectionOptionsAccepterOutput{}) }

type PeeringConnectionOptionsRequester struct {
	// Allow a local linked EC2-Classic instance to communicate
	// with instances in a peer VPC. This enables an outbound communication from the local ClassicLink connection
	// to the remote VPC. This option is not supported for inter-region VPC peering.
	AllowClassicLinkToRemoteVpc *bool `pulumi:"allowClassicLinkToRemoteVpc"`
	// Allow a local VPC to resolve public DNS hostnames to
	// private IP addresses when queried from instances in the peer VPC.
	AllowRemoteVpcDnsResolution *bool `pulumi:"allowRemoteVpcDnsResolution"`
	// Allow a local VPC to communicate with a linked EC2-Classic
	// instance in a peer VPC. This enables an outbound communication from the local VPC to the remote ClassicLink
	// connection. This option is not supported for inter-region VPC peering.
	AllowVpcToRemoteClassicLink *bool `pulumi:"allowVpcToRemoteClassicLink"`
}
var peeringConnectionOptionsRequesterType = reflect.TypeOf((*PeeringConnectionOptionsRequester)(nil)).Elem()

type PeeringConnectionOptionsRequesterInput interface {
	pulumi.Input

	ToPeeringConnectionOptionsRequesterOutput() PeeringConnectionOptionsRequesterOutput
	ToPeeringConnectionOptionsRequesterOutputWithContext(ctx context.Context) PeeringConnectionOptionsRequesterOutput
}

type PeeringConnectionOptionsRequesterArgs struct {
	// Allow a local linked EC2-Classic instance to communicate
	// with instances in a peer VPC. This enables an outbound communication from the local ClassicLink connection
	// to the remote VPC. This option is not supported for inter-region VPC peering.
	AllowClassicLinkToRemoteVpc pulumi.BoolInput `pulumi:"allowClassicLinkToRemoteVpc"`
	// Allow a local VPC to resolve public DNS hostnames to
	// private IP addresses when queried from instances in the peer VPC.
	AllowRemoteVpcDnsResolution pulumi.BoolInput `pulumi:"allowRemoteVpcDnsResolution"`
	// Allow a local VPC to communicate with a linked EC2-Classic
	// instance in a peer VPC. This enables an outbound communication from the local VPC to the remote ClassicLink
	// connection. This option is not supported for inter-region VPC peering.
	AllowVpcToRemoteClassicLink pulumi.BoolInput `pulumi:"allowVpcToRemoteClassicLink"`
}

func (PeeringConnectionOptionsRequesterArgs) ElementType() reflect.Type {
	return peeringConnectionOptionsRequesterType
}

func (a PeeringConnectionOptionsRequesterArgs) ToPeeringConnectionOptionsRequesterOutput() PeeringConnectionOptionsRequesterOutput {
	return pulumi.ToOutput(a).(PeeringConnectionOptionsRequesterOutput)
}

func (a PeeringConnectionOptionsRequesterArgs) ToPeeringConnectionOptionsRequesterOutputWithContext(ctx context.Context) PeeringConnectionOptionsRequesterOutput {
	return pulumi.ToOutputWithContext(ctx, a).(PeeringConnectionOptionsRequesterOutput)
}

type PeeringConnectionOptionsRequesterOutput struct { *pulumi.OutputState }

// Allow a local linked EC2-Classic instance to communicate
// with instances in a peer VPC. This enables an outbound communication from the local ClassicLink connection
// to the remote VPC. This option is not supported for inter-region VPC peering.
func (o PeeringConnectionOptionsRequesterOutput) AllowClassicLinkToRemoteVpc() pulumi.BoolOutput {
	return o.Apply(func(v PeeringConnectionOptionsRequester) bool {
		if v.AllowClassicLinkToRemoteVpc == nil { return *new(bool) } else { return *v.AllowClassicLinkToRemoteVpc }
	}).(pulumi.BoolOutput)
}

// Allow a local VPC to resolve public DNS hostnames to
// private IP addresses when queried from instances in the peer VPC.
func (o PeeringConnectionOptionsRequesterOutput) AllowRemoteVpcDnsResolution() pulumi.BoolOutput {
	return o.Apply(func(v PeeringConnectionOptionsRequester) bool {
		if v.AllowRemoteVpcDnsResolution == nil { return *new(bool) } else { return *v.AllowRemoteVpcDnsResolution }
	}).(pulumi.BoolOutput)
}

// Allow a local VPC to communicate with a linked EC2-Classic
// instance in a peer VPC. This enables an outbound communication from the local VPC to the remote ClassicLink
// connection. This option is not supported for inter-region VPC peering.
func (o PeeringConnectionOptionsRequesterOutput) AllowVpcToRemoteClassicLink() pulumi.BoolOutput {
	return o.Apply(func(v PeeringConnectionOptionsRequester) bool {
		if v.AllowVpcToRemoteClassicLink == nil { return *new(bool) } else { return *v.AllowVpcToRemoteClassicLink }
	}).(pulumi.BoolOutput)
}

func (PeeringConnectionOptionsRequesterOutput) ElementType() reflect.Type {
	return peeringConnectionOptionsRequesterType
}

func (o PeeringConnectionOptionsRequesterOutput) ToPeeringConnectionOptionsRequesterOutput() PeeringConnectionOptionsRequesterOutput {
	return o
}

func (o PeeringConnectionOptionsRequesterOutput) ToPeeringConnectionOptionsRequesterOutputWithContext(ctx context.Context) PeeringConnectionOptionsRequesterOutput {
	return o
}

func init() { pulumi.RegisterOutputType(PeeringConnectionOptionsRequesterOutput{}) }

