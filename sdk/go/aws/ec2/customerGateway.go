// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a customer gateway inside a VPC. These objects can be connected to VPN gateways via VPN connections, and allow you to establish tunnels between your network and the VPC.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/customer_gateway.html.markdown.
type CustomerGateway struct {
	pulumi.CustomResourceState

	// The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
	BgpAsn pulumi.IntOutput `pulumi:"bgpAsn"`

	// The IP address of the gateway's Internet-routable external interface.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`

	// Tags to apply to the gateway.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// The type of customer gateway. The only type AWS
	// supports at this time is "ipsec.1".
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCustomerGateway registers a new resource with the given unique name, arguments, and options.
func NewCustomerGateway(ctx *pulumi.Context,
	name string, args *CustomerGatewayArgs, opts ...pulumi.ResourceOption) (*CustomerGateway, error) {
	if args == nil || args.BgpAsn == nil {
		return nil, errors.New("missing required argument 'BgpAsn'")
	}
	if args == nil || args.IpAddress == nil {
		return nil, errors.New("missing required argument 'IpAddress'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.BgpAsn; i != nil { inputs["bgpAsn"] = i.ToIntOutput() }
		if i := args.IpAddress; i != nil { inputs["ipAddress"] = i.ToStringOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := args.Type; i != nil { inputs["type"] = i.ToStringOutput() }
	}
	var resource CustomerGateway
	err := ctx.RegisterResource("aws:ec2/customerGateway:CustomerGateway", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomerGateway gets an existing CustomerGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomerGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomerGatewayState, opts ...pulumi.ResourceOption) (*CustomerGateway, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.BgpAsn; i != nil { inputs["bgpAsn"] = i.ToIntOutput() }
		if i := state.IpAddress; i != nil { inputs["ipAddress"] = i.ToStringOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := state.Type; i != nil { inputs["type"] = i.ToStringOutput() }
	}
	var resource CustomerGateway
	err := ctx.ReadResource("aws:ec2/customerGateway:CustomerGateway", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomerGateway resources.
type CustomerGatewayState struct {
	// The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
	BgpAsn pulumi.IntInput `pulumi:"bgpAsn"`
	// The IP address of the gateway's Internet-routable external interface.
	IpAddress pulumi.StringInput `pulumi:"ipAddress"`
	// Tags to apply to the gateway.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The type of customer gateway. The only type AWS
	// supports at this time is "ipsec.1".
	Type pulumi.StringInput `pulumi:"type"`
}

// The set of arguments for constructing a CustomerGateway resource.
type CustomerGatewayArgs struct {
	// The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
	BgpAsn pulumi.IntInput `pulumi:"bgpAsn"`
	// The IP address of the gateway's Internet-routable external interface.
	IpAddress pulumi.StringInput `pulumi:"ipAddress"`
	// Tags to apply to the gateway.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The type of customer gateway. The only type AWS
	// supports at this time is "ipsec.1".
	Type pulumi.StringInput `pulumi:"type"`
}
