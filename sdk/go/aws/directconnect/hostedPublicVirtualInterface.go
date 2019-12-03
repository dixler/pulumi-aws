// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Direct Connect hosted public virtual interface resource. This resource represents the allocator's side of the hosted virtual interface.
// A hosted virtual interface is a virtual interface that is owned by another AWS account.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dx_hosted_public_virtual_interface.html.markdown.
type HostedPublicVirtualInterface struct {
	pulumi.CustomResourceState

	// The address family for the BGP peer. `ipv4 ` or `ipv6`.
	AddressFamily pulumi.StringOutput `pulumi:"addressFamily"`

	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress pulumi.StringOutput `pulumi:"amazonAddress"`

	// The ARN of the virtual interface.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The Direct Connect endpoint on which the virtual interface terminates.
	AwsDevice pulumi.StringOutput `pulumi:"awsDevice"`

	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn pulumi.IntOutput `pulumi:"bgpAsn"`

	// The authentication key for BGP configuration.
	BgpAuthKey pulumi.StringOutput `pulumi:"bgpAuthKey"`

	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	ConnectionId pulumi.StringOutput `pulumi:"connectionId"`

	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress pulumi.StringOutput `pulumi:"customerAddress"`

	// The name for the virtual interface.
	Name pulumi.StringOutput `pulumi:"name"`

	// The AWS account that will own the new virtual interface.
	OwnerAccountId pulumi.StringOutput `pulumi:"ownerAccountId"`

	// A list of routes to be advertised to the AWS network in this region.
	RouteFilterPrefixes pulumi.StringArrayOutput `pulumi:"routeFilterPrefixes"`

	// The VLAN ID.
	Vlan pulumi.IntOutput `pulumi:"vlan"`
}

// NewHostedPublicVirtualInterface registers a new resource with the given unique name, arguments, and options.
func NewHostedPublicVirtualInterface(ctx *pulumi.Context,
	name string, args *HostedPublicVirtualInterfaceArgs, opts ...pulumi.ResourceOption) (*HostedPublicVirtualInterface, error) {
	if args == nil || args.AddressFamily == nil {
		return nil, errors.New("missing required argument 'AddressFamily'")
	}
	if args == nil || args.BgpAsn == nil {
		return nil, errors.New("missing required argument 'BgpAsn'")
	}
	if args == nil || args.ConnectionId == nil {
		return nil, errors.New("missing required argument 'ConnectionId'")
	}
	if args == nil || args.OwnerAccountId == nil {
		return nil, errors.New("missing required argument 'OwnerAccountId'")
	}
	if args == nil || args.RouteFilterPrefixes == nil {
		return nil, errors.New("missing required argument 'RouteFilterPrefixes'")
	}
	if args == nil || args.Vlan == nil {
		return nil, errors.New("missing required argument 'Vlan'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.AddressFamily; i != nil { inputs["addressFamily"] = i.ToStringOutput() }
		if i := args.AmazonAddress; i != nil { inputs["amazonAddress"] = i.ToStringOutput() }
		if i := args.BgpAsn; i != nil { inputs["bgpAsn"] = i.ToIntOutput() }
		if i := args.BgpAuthKey; i != nil { inputs["bgpAuthKey"] = i.ToStringOutput() }
		if i := args.ConnectionId; i != nil { inputs["connectionId"] = i.ToStringOutput() }
		if i := args.CustomerAddress; i != nil { inputs["customerAddress"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.OwnerAccountId; i != nil { inputs["ownerAccountId"] = i.ToStringOutput() }
		if i := args.RouteFilterPrefixes; i != nil { inputs["routeFilterPrefixes"] = i.ToStringArrayOutput() }
		if i := args.Vlan; i != nil { inputs["vlan"] = i.ToIntOutput() }
	}
	var resource HostedPublicVirtualInterface
	err := ctx.RegisterResource("aws:directconnect/hostedPublicVirtualInterface:HostedPublicVirtualInterface", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostedPublicVirtualInterface gets an existing HostedPublicVirtualInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostedPublicVirtualInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostedPublicVirtualInterfaceState, opts ...pulumi.ResourceOption) (*HostedPublicVirtualInterface, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.AddressFamily; i != nil { inputs["addressFamily"] = i.ToStringOutput() }
		if i := state.AmazonAddress; i != nil { inputs["amazonAddress"] = i.ToStringOutput() }
		if i := state.Arn; i != nil { inputs["arn"] = i.ToStringOutput() }
		if i := state.AwsDevice; i != nil { inputs["awsDevice"] = i.ToStringOutput() }
		if i := state.BgpAsn; i != nil { inputs["bgpAsn"] = i.ToIntOutput() }
		if i := state.BgpAuthKey; i != nil { inputs["bgpAuthKey"] = i.ToStringOutput() }
		if i := state.ConnectionId; i != nil { inputs["connectionId"] = i.ToStringOutput() }
		if i := state.CustomerAddress; i != nil { inputs["customerAddress"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.OwnerAccountId; i != nil { inputs["ownerAccountId"] = i.ToStringOutput() }
		if i := state.RouteFilterPrefixes; i != nil { inputs["routeFilterPrefixes"] = i.ToStringArrayOutput() }
		if i := state.Vlan; i != nil { inputs["vlan"] = i.ToIntOutput() }
	}
	var resource HostedPublicVirtualInterface
	err := ctx.ReadResource("aws:directconnect/hostedPublicVirtualInterface:HostedPublicVirtualInterface", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostedPublicVirtualInterface resources.
type HostedPublicVirtualInterfaceState struct {
	// The address family for the BGP peer. `ipv4 ` or `ipv6`.
	AddressFamily pulumi.StringInput `pulumi:"addressFamily"`
	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress pulumi.StringInput `pulumi:"amazonAddress"`
	// The ARN of the virtual interface.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The Direct Connect endpoint on which the virtual interface terminates.
	AwsDevice pulumi.StringInput `pulumi:"awsDevice"`
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn pulumi.IntInput `pulumi:"bgpAsn"`
	// The authentication key for BGP configuration.
	BgpAuthKey pulumi.StringInput `pulumi:"bgpAuthKey"`
	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	ConnectionId pulumi.StringInput `pulumi:"connectionId"`
	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress pulumi.StringInput `pulumi:"customerAddress"`
	// The name for the virtual interface.
	Name pulumi.StringInput `pulumi:"name"`
	// The AWS account that will own the new virtual interface.
	OwnerAccountId pulumi.StringInput `pulumi:"ownerAccountId"`
	// A list of routes to be advertised to the AWS network in this region.
	RouteFilterPrefixes pulumi.StringArrayInput `pulumi:"routeFilterPrefixes"`
	// The VLAN ID.
	Vlan pulumi.IntInput `pulumi:"vlan"`
}

// The set of arguments for constructing a HostedPublicVirtualInterface resource.
type HostedPublicVirtualInterfaceArgs struct {
	// The address family for the BGP peer. `ipv4 ` or `ipv6`.
	AddressFamily pulumi.StringInput `pulumi:"addressFamily"`
	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress pulumi.StringInput `pulumi:"amazonAddress"`
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn pulumi.IntInput `pulumi:"bgpAsn"`
	// The authentication key for BGP configuration.
	BgpAuthKey pulumi.StringInput `pulumi:"bgpAuthKey"`
	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	ConnectionId pulumi.StringInput `pulumi:"connectionId"`
	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress pulumi.StringInput `pulumi:"customerAddress"`
	// The name for the virtual interface.
	Name pulumi.StringInput `pulumi:"name"`
	// The AWS account that will own the new virtual interface.
	OwnerAccountId pulumi.StringInput `pulumi:"ownerAccountId"`
	// A list of routes to be advertised to the AWS network in this region.
	RouteFilterPrefixes pulumi.StringArrayInput `pulumi:"routeFilterPrefixes"`
	// The VLAN ID.
	Vlan pulumi.IntInput `pulumi:"vlan"`
}
