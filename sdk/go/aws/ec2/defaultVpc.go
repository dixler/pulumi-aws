// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage the [default AWS VPC](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/default-vpc.html)
// in the current region.
// 
// For AWS accounts created after 2013-12-04, each region comes with a Default VPC.
// **This is an advanced resource**, and has special caveats to be aware of when
// using it. Please read this document in its entirety before using this resource.
// 
// The `ec2.DefaultVpc` behaves differently from normal resources, in that
// this provider does not _create_ this resource, but instead "adopts" it
// into management.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/default_vpc.html.markdown.
type DefaultVpc struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of VPC
	Arn pulumi.StringOutput `pulumi:"arn"`

	// Whether or not an Amazon-provided IPv6 CIDR
	// block with a /56 prefix length for the VPC was assigned
	AssignGeneratedIpv6CidrBlock pulumi.BoolOutput `pulumi:"assignGeneratedIpv6CidrBlock"`

	// The CIDR block of the VPC
	CidrBlock pulumi.StringOutput `pulumi:"cidrBlock"`

	// The ID of the network ACL created by default on VPC creation
	DefaultNetworkAclId pulumi.StringOutput `pulumi:"defaultNetworkAclId"`

	// The ID of the route table created by default on VPC creation
	DefaultRouteTableId pulumi.StringOutput `pulumi:"defaultRouteTableId"`

	// The ID of the security group created by default on VPC creation
	DefaultSecurityGroupId pulumi.StringOutput `pulumi:"defaultSecurityGroupId"`

	DhcpOptionsId pulumi.StringOutput `pulumi:"dhcpOptionsId"`

	// A boolean flag to enable/disable ClassicLink
	// for the VPC. Only valid in regions and accounts that support EC2 Classic.
	// See the [ClassicLink documentation][1] for more information. Defaults false.
	EnableClassiclink pulumi.BoolOutput `pulumi:"enableClassiclink"`

	EnableClassiclinkDnsSupport pulumi.BoolOutput `pulumi:"enableClassiclinkDnsSupport"`

	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
	EnableDnsHostnames pulumi.BoolOutput `pulumi:"enableDnsHostnames"`

	// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
	EnableDnsSupport pulumi.BoolOutput `pulumi:"enableDnsSupport"`

	// Tenancy of instances spin up within VPC.
	InstanceTenancy pulumi.StringOutput `pulumi:"instanceTenancy"`

	// The association ID for the IPv6 CIDR block of the VPC
	Ipv6AssociationId pulumi.StringOutput `pulumi:"ipv6AssociationId"`

	// The IPv6 CIDR block of the VPC
	Ipv6CidrBlock pulumi.StringOutput `pulumi:"ipv6CidrBlock"`

	// The ID of the main route table associated with
	// this VPC. Note that you can change a VPC's main route table by using an
	// [`ec2.MainRouteTableAssociation`](https://www.terraform.io/docs/providers/aws/r/main_route_table_assoc.html)
	MainRouteTableId pulumi.StringOutput `pulumi:"mainRouteTableId"`

	// The ID of the AWS account that owns the VPC.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewDefaultVpc registers a new resource with the given unique name, arguments, and options.
func NewDefaultVpc(ctx *pulumi.Context,
	name string, args *DefaultVpcArgs, opts ...pulumi.ResourceOption) (*DefaultVpc, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.EnableClassiclink; i != nil { inputs["enableClassiclink"] = i.ToBoolOutput() }
		if i := args.EnableClassiclinkDnsSupport; i != nil { inputs["enableClassiclinkDnsSupport"] = i.ToBoolOutput() }
		if i := args.EnableDnsHostnames; i != nil { inputs["enableDnsHostnames"] = i.ToBoolOutput() }
		if i := args.EnableDnsSupport; i != nil { inputs["enableDnsSupport"] = i.ToBoolOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource DefaultVpc
	err := ctx.RegisterResource("aws:ec2/defaultVpc:DefaultVpc", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultVpc gets an existing DefaultVpc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultVpc(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultVpcState, opts ...pulumi.ResourceOption) (*DefaultVpc, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Arn; i != nil { inputs["arn"] = i.ToStringOutput() }
		if i := state.AssignGeneratedIpv6CidrBlock; i != nil { inputs["assignGeneratedIpv6CidrBlock"] = i.ToBoolOutput() }
		if i := state.CidrBlock; i != nil { inputs["cidrBlock"] = i.ToStringOutput() }
		if i := state.DefaultNetworkAclId; i != nil { inputs["defaultNetworkAclId"] = i.ToStringOutput() }
		if i := state.DefaultRouteTableId; i != nil { inputs["defaultRouteTableId"] = i.ToStringOutput() }
		if i := state.DefaultSecurityGroupId; i != nil { inputs["defaultSecurityGroupId"] = i.ToStringOutput() }
		if i := state.DhcpOptionsId; i != nil { inputs["dhcpOptionsId"] = i.ToStringOutput() }
		if i := state.EnableClassiclink; i != nil { inputs["enableClassiclink"] = i.ToBoolOutput() }
		if i := state.EnableClassiclinkDnsSupport; i != nil { inputs["enableClassiclinkDnsSupport"] = i.ToBoolOutput() }
		if i := state.EnableDnsHostnames; i != nil { inputs["enableDnsHostnames"] = i.ToBoolOutput() }
		if i := state.EnableDnsSupport; i != nil { inputs["enableDnsSupport"] = i.ToBoolOutput() }
		if i := state.InstanceTenancy; i != nil { inputs["instanceTenancy"] = i.ToStringOutput() }
		if i := state.Ipv6AssociationId; i != nil { inputs["ipv6AssociationId"] = i.ToStringOutput() }
		if i := state.Ipv6CidrBlock; i != nil { inputs["ipv6CidrBlock"] = i.ToStringOutput() }
		if i := state.MainRouteTableId; i != nil { inputs["mainRouteTableId"] = i.ToStringOutput() }
		if i := state.OwnerId; i != nil { inputs["ownerId"] = i.ToStringOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource DefaultVpc
	err := ctx.ReadResource("aws:ec2/defaultVpc:DefaultVpc", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultVpc resources.
type DefaultVpcState struct {
	// Amazon Resource Name (ARN) of VPC
	Arn pulumi.StringInput `pulumi:"arn"`
	// Whether or not an Amazon-provided IPv6 CIDR
	// block with a /56 prefix length for the VPC was assigned
	AssignGeneratedIpv6CidrBlock pulumi.BoolInput `pulumi:"assignGeneratedIpv6CidrBlock"`
	// The CIDR block of the VPC
	CidrBlock pulumi.StringInput `pulumi:"cidrBlock"`
	// The ID of the network ACL created by default on VPC creation
	DefaultNetworkAclId pulumi.StringInput `pulumi:"defaultNetworkAclId"`
	// The ID of the route table created by default on VPC creation
	DefaultRouteTableId pulumi.StringInput `pulumi:"defaultRouteTableId"`
	// The ID of the security group created by default on VPC creation
	DefaultSecurityGroupId pulumi.StringInput `pulumi:"defaultSecurityGroupId"`
	DhcpOptionsId pulumi.StringInput `pulumi:"dhcpOptionsId"`
	// A boolean flag to enable/disable ClassicLink
	// for the VPC. Only valid in regions and accounts that support EC2 Classic.
	// See the [ClassicLink documentation][1] for more information. Defaults false.
	EnableClassiclink pulumi.BoolInput `pulumi:"enableClassiclink"`
	EnableClassiclinkDnsSupport pulumi.BoolInput `pulumi:"enableClassiclinkDnsSupport"`
	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
	EnableDnsHostnames pulumi.BoolInput `pulumi:"enableDnsHostnames"`
	// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
	EnableDnsSupport pulumi.BoolInput `pulumi:"enableDnsSupport"`
	// Tenancy of instances spin up within VPC.
	InstanceTenancy pulumi.StringInput `pulumi:"instanceTenancy"`
	// The association ID for the IPv6 CIDR block of the VPC
	Ipv6AssociationId pulumi.StringInput `pulumi:"ipv6AssociationId"`
	// The IPv6 CIDR block of the VPC
	Ipv6CidrBlock pulumi.StringInput `pulumi:"ipv6CidrBlock"`
	// The ID of the main route table associated with
	// this VPC. Note that you can change a VPC's main route table by using an
	// [`ec2.MainRouteTableAssociation`](https://www.terraform.io/docs/providers/aws/r/main_route_table_assoc.html)
	MainRouteTableId pulumi.StringInput `pulumi:"mainRouteTableId"`
	// The ID of the AWS account that owns the VPC.
	OwnerId pulumi.StringInput `pulumi:"ownerId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a DefaultVpc resource.
type DefaultVpcArgs struct {
	// A boolean flag to enable/disable ClassicLink
	// for the VPC. Only valid in regions and accounts that support EC2 Classic.
	// See the [ClassicLink documentation][1] for more information. Defaults false.
	EnableClassiclink pulumi.BoolInput `pulumi:"enableClassiclink"`
	EnableClassiclinkDnsSupport pulumi.BoolInput `pulumi:"enableClassiclinkDnsSupport"`
	// A boolean flag to enable/disable DNS hostnames in the VPC. Defaults false.
	EnableDnsHostnames pulumi.BoolInput `pulumi:"enableDnsHostnames"`
	// A boolean flag to enable/disable DNS support in the VPC. Defaults true.
	EnableDnsSupport pulumi.BoolInput `pulumi:"enableDnsSupport"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}
