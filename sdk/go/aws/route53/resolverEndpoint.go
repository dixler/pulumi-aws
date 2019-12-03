// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Route 53 Resolver endpoint resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/route53_resolver_endpoint.html.markdown.
type ResolverEndpoint struct {
	pulumi.CustomResourceState

	// The ARN of the Route 53 Resolver endpoint.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The direction of DNS queries to or from the Route 53 Resolver endpoint.
	// Valid values are `INBOUND` (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
	// or `OUTBOUND` (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
	Direction pulumi.StringOutput `pulumi:"direction"`

	// The ID of the VPC that you want to create the resolver endpoint in.
	HostVpcId pulumi.StringOutput `pulumi:"hostVpcId"`

	// The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
	// to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
	IpAddresses ResolverEndpointIpAddressesArrayOutput `pulumi:"ipAddresses"`

	// The friendly name of the Route 53 Resolver endpoint.
	Name pulumi.StringOutput `pulumi:"name"`

	// The ID of one or more security groups that you want to use to control access to this VPC.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewResolverEndpoint registers a new resource with the given unique name, arguments, and options.
func NewResolverEndpoint(ctx *pulumi.Context,
	name string, args *ResolverEndpointArgs, opts ...pulumi.ResourceOption) (*ResolverEndpoint, error) {
	if args == nil || args.Direction == nil {
		return nil, errors.New("missing required argument 'Direction'")
	}
	if args == nil || args.IpAddresses == nil {
		return nil, errors.New("missing required argument 'IpAddresses'")
	}
	if args == nil || args.SecurityGroupIds == nil {
		return nil, errors.New("missing required argument 'SecurityGroupIds'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Direction; i != nil { inputs["direction"] = i.ToStringOutput() }
		if i := args.IpAddresses; i != nil { inputs["ipAddresses"] = i.ToResolverEndpointIpAddressesArrayOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.SecurityGroupIds; i != nil { inputs["securityGroupIds"] = i.ToStringArrayOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource ResolverEndpoint
	err := ctx.RegisterResource("aws:route53/resolverEndpoint:ResolverEndpoint", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResolverEndpoint gets an existing ResolverEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolverEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResolverEndpointState, opts ...pulumi.ResourceOption) (*ResolverEndpoint, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Arn; i != nil { inputs["arn"] = i.ToStringOutput() }
		if i := state.Direction; i != nil { inputs["direction"] = i.ToStringOutput() }
		if i := state.HostVpcId; i != nil { inputs["hostVpcId"] = i.ToStringOutput() }
		if i := state.IpAddresses; i != nil { inputs["ipAddresses"] = i.ToResolverEndpointIpAddressesArrayOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.SecurityGroupIds; i != nil { inputs["securityGroupIds"] = i.ToStringArrayOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource ResolverEndpoint
	err := ctx.ReadResource("aws:route53/resolverEndpoint:ResolverEndpoint", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResolverEndpoint resources.
type ResolverEndpointState struct {
	// The ARN of the Route 53 Resolver endpoint.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The direction of DNS queries to or from the Route 53 Resolver endpoint.
	// Valid values are `INBOUND` (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
	// or `OUTBOUND` (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
	Direction pulumi.StringInput `pulumi:"direction"`
	// The ID of the VPC that you want to create the resolver endpoint in.
	HostVpcId pulumi.StringInput `pulumi:"hostVpcId"`
	// The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
	// to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
	IpAddresses ResolverEndpointIpAddressesArrayInput `pulumi:"ipAddresses"`
	// The friendly name of the Route 53 Resolver endpoint.
	Name pulumi.StringInput `pulumi:"name"`
	// The ID of one or more security groups that you want to use to control access to this VPC.
	SecurityGroupIds pulumi.StringArrayInput `pulumi:"securityGroupIds"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a ResolverEndpoint resource.
type ResolverEndpointArgs struct {
	// The direction of DNS queries to or from the Route 53 Resolver endpoint.
	// Valid values are `INBOUND` (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
	// or `OUTBOUND` (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
	Direction pulumi.StringInput `pulumi:"direction"`
	// The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
	// to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
	IpAddresses ResolverEndpointIpAddressesArrayInput `pulumi:"ipAddresses"`
	// The friendly name of the Route 53 Resolver endpoint.
	Name pulumi.StringInput `pulumi:"name"`
	// The ID of one or more security groups that you want to use to control access to this VPC.
	SecurityGroupIds pulumi.StringArrayInput `pulumi:"securityGroupIds"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}
type ResolverEndpointIpAddresses struct {
	// The IP address in the subnet that you want to use for DNS queries.
	Ip *string `pulumi:"ip"`
	IpId *string `pulumi:"ipId"`
	// The ID of the subnet that contains the IP address.
	SubnetId string `pulumi:"subnetId"`
}
var resolverEndpointIpAddressesType = reflect.TypeOf((*ResolverEndpointIpAddresses)(nil)).Elem()

type ResolverEndpointIpAddressesInput interface {
	pulumi.Input

	ToResolverEndpointIpAddressesOutput() ResolverEndpointIpAddressesOutput
	ToResolverEndpointIpAddressesOutputWithContext(ctx context.Context) ResolverEndpointIpAddressesOutput
}

type ResolverEndpointIpAddressesArgs struct {
	// The IP address in the subnet that you want to use for DNS queries.
	Ip pulumi.StringInput `pulumi:"ip"`
	IpId pulumi.StringInput `pulumi:"ipId"`
	// The ID of the subnet that contains the IP address.
	SubnetId pulumi.StringInput `pulumi:"subnetId"`
}

func (ResolverEndpointIpAddressesArgs) ElementType() reflect.Type {
	return resolverEndpointIpAddressesType
}

func (a ResolverEndpointIpAddressesArgs) ToResolverEndpointIpAddressesOutput() ResolverEndpointIpAddressesOutput {
	return pulumi.ToOutput(a).(ResolverEndpointIpAddressesOutput)
}

func (a ResolverEndpointIpAddressesArgs) ToResolverEndpointIpAddressesOutputWithContext(ctx context.Context) ResolverEndpointIpAddressesOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ResolverEndpointIpAddressesOutput)
}

type ResolverEndpointIpAddressesOutput struct { *pulumi.OutputState }

// The IP address in the subnet that you want to use for DNS queries.
func (o ResolverEndpointIpAddressesOutput) Ip() pulumi.StringOutput {
	return o.Apply(func(v ResolverEndpointIpAddresses) string {
		if v.Ip == nil { return *new(string) } else { return *v.Ip }
	}).(pulumi.StringOutput)
}

func (o ResolverEndpointIpAddressesOutput) IpId() pulumi.StringOutput {
	return o.Apply(func(v ResolverEndpointIpAddresses) string {
		if v.IpId == nil { return *new(string) } else { return *v.IpId }
	}).(pulumi.StringOutput)
}

// The ID of the subnet that contains the IP address.
func (o ResolverEndpointIpAddressesOutput) SubnetId() pulumi.StringOutput {
	return o.Apply(func(v ResolverEndpointIpAddresses) string {
		return v.SubnetId
	}).(pulumi.StringOutput)
}

func (ResolverEndpointIpAddressesOutput) ElementType() reflect.Type {
	return resolverEndpointIpAddressesType
}

func (o ResolverEndpointIpAddressesOutput) ToResolverEndpointIpAddressesOutput() ResolverEndpointIpAddressesOutput {
	return o
}

func (o ResolverEndpointIpAddressesOutput) ToResolverEndpointIpAddressesOutputWithContext(ctx context.Context) ResolverEndpointIpAddressesOutput {
	return o
}

func init() { pulumi.RegisterOutputType(ResolverEndpointIpAddressesOutput{}) }

var resolverEndpointIpAddressesArrayType = reflect.TypeOf((*[]ResolverEndpointIpAddresses)(nil)).Elem()

type ResolverEndpointIpAddressesArrayInput interface {
	pulumi.Input

	ToResolverEndpointIpAddressesArrayOutput() ResolverEndpointIpAddressesArrayOutput
	ToResolverEndpointIpAddressesArrayOutputWithContext(ctx context.Context) ResolverEndpointIpAddressesArrayOutput
}

type ResolverEndpointIpAddressesArrayArgs []ResolverEndpointIpAddressesInput

func (ResolverEndpointIpAddressesArrayArgs) ElementType() reflect.Type {
	return resolverEndpointIpAddressesArrayType
}

func (a ResolverEndpointIpAddressesArrayArgs) ToResolverEndpointIpAddressesArrayOutput() ResolverEndpointIpAddressesArrayOutput {
	return pulumi.ToOutput(a).(ResolverEndpointIpAddressesArrayOutput)
}

func (a ResolverEndpointIpAddressesArrayArgs) ToResolverEndpointIpAddressesArrayOutputWithContext(ctx context.Context) ResolverEndpointIpAddressesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ResolverEndpointIpAddressesArrayOutput)
}

type ResolverEndpointIpAddressesArrayOutput struct { *pulumi.OutputState }

func (o ResolverEndpointIpAddressesArrayOutput) Index(i pulumi.IntInput) ResolverEndpointIpAddressesOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) ResolverEndpointIpAddresses {
		return vs[0].([]ResolverEndpointIpAddresses)[vs[1].(int)]
	}).(ResolverEndpointIpAddressesOutput)
}

func (ResolverEndpointIpAddressesArrayOutput) ElementType() reflect.Type {
	return resolverEndpointIpAddressesArrayType
}

func (o ResolverEndpointIpAddressesArrayOutput) ToResolverEndpointIpAddressesArrayOutput() ResolverEndpointIpAddressesArrayOutput {
	return o
}

func (o ResolverEndpointIpAddressesArrayOutput) ToResolverEndpointIpAddressesArrayOutputWithContext(ctx context.Context) ResolverEndpointIpAddressesArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(ResolverEndpointIpAddressesArrayOutput{}) }

