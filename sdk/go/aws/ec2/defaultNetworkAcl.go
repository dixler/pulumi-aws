// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage the default AWS Network ACL. VPC Only.
// 
// Each VPC created in AWS comes with a Default Network ACL that can be managed, but not
// destroyed. **This is an advanced resource**, and has special caveats to be aware
// of when using it. Please read this document in its entirety before using this
// resource.
// 
// The `ec2.DefaultNetworkAcl` behaves differently from normal resources, in that
// this provider does not _create_ this resource, but instead attempts to "adopt" it
// into management. We can do this because each VPC created has a Default Network
// ACL that cannot be destroyed, and is created with a known set of default rules.
// 
// When this provider first adopts the Default Network ACL, it **immediately removes all
// rules in the ACL**. It then proceeds to create any rules specified in the
// configuration. This step is required so that only the rules specified in the
// configuration are created.
// 
// This resource treats its inline rules as absolute; only the rules defined
// inline are created, and any additions/removals external to this resource will
// result in diffs being shown. For these reasons, this resource is incompatible with the
// `ec2.NetworkAclRule` resource.
// 
// For more information about Network ACLs, see the AWS Documentation on
// [Network ACLs][aws-network-acls].
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/default_network_acl.html.markdown.
type DefaultNetworkAcl struct {
	pulumi.CustomResourceState

	// The Network ACL ID to manage. This
	// attribute is exported from `ec2.Vpc`, or manually found via the AWS Console.
	DefaultNetworkAclId pulumi.StringOutput `pulumi:"defaultNetworkAclId"`

	// Specifies an egress rule. Parameters defined below.
	Egress DefaultNetworkAclEgressArrayOutput `pulumi:"egress"`

	// Specifies an ingress rule. Parameters defined below.
	Ingress DefaultNetworkAclIngressArrayOutput `pulumi:"ingress"`

	// The ID of the AWS account that owns the Default Network ACL
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`

	// A list of Subnet IDs to apply the ACL to. See the
	// notes below on managing Subnets in the Default Network ACL
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// The ID of the associated VPC
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewDefaultNetworkAcl registers a new resource with the given unique name, arguments, and options.
func NewDefaultNetworkAcl(ctx *pulumi.Context,
	name string, args *DefaultNetworkAclArgs, opts ...pulumi.ResourceOption) (*DefaultNetworkAcl, error) {
	if args == nil || args.DefaultNetworkAclId == nil {
		return nil, errors.New("missing required argument 'DefaultNetworkAclId'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.DefaultNetworkAclId; i != nil { inputs["defaultNetworkAclId"] = i.ToStringOutput() }
		if i := args.Egress; i != nil { inputs["egress"] = i.ToDefaultNetworkAclEgressArrayOutput() }
		if i := args.Ingress; i != nil { inputs["ingress"] = i.ToDefaultNetworkAclIngressArrayOutput() }
		if i := args.SubnetIds; i != nil { inputs["subnetIds"] = i.ToStringArrayOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource DefaultNetworkAcl
	err := ctx.RegisterResource("aws:ec2/defaultNetworkAcl:DefaultNetworkAcl", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultNetworkAcl gets an existing DefaultNetworkAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultNetworkAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultNetworkAclState, opts ...pulumi.ResourceOption) (*DefaultNetworkAcl, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.DefaultNetworkAclId; i != nil { inputs["defaultNetworkAclId"] = i.ToStringOutput() }
		if i := state.Egress; i != nil { inputs["egress"] = i.ToDefaultNetworkAclEgressArrayOutput() }
		if i := state.Ingress; i != nil { inputs["ingress"] = i.ToDefaultNetworkAclIngressArrayOutput() }
		if i := state.OwnerId; i != nil { inputs["ownerId"] = i.ToStringOutput() }
		if i := state.SubnetIds; i != nil { inputs["subnetIds"] = i.ToStringArrayOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := state.VpcId; i != nil { inputs["vpcId"] = i.ToStringOutput() }
	}
	var resource DefaultNetworkAcl
	err := ctx.ReadResource("aws:ec2/defaultNetworkAcl:DefaultNetworkAcl", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultNetworkAcl resources.
type DefaultNetworkAclState struct {
	// The Network ACL ID to manage. This
	// attribute is exported from `ec2.Vpc`, or manually found via the AWS Console.
	DefaultNetworkAclId pulumi.StringInput `pulumi:"defaultNetworkAclId"`
	// Specifies an egress rule. Parameters defined below.
	Egress DefaultNetworkAclEgressArrayInput `pulumi:"egress"`
	// Specifies an ingress rule. Parameters defined below.
	Ingress DefaultNetworkAclIngressArrayInput `pulumi:"ingress"`
	// The ID of the AWS account that owns the Default Network ACL
	OwnerId pulumi.StringInput `pulumi:"ownerId"`
	// A list of Subnet IDs to apply the ACL to. See the
	// notes below on managing Subnets in the Default Network ACL
	SubnetIds pulumi.StringArrayInput `pulumi:"subnetIds"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The ID of the associated VPC
	VpcId pulumi.StringInput `pulumi:"vpcId"`
}

// The set of arguments for constructing a DefaultNetworkAcl resource.
type DefaultNetworkAclArgs struct {
	// The Network ACL ID to manage. This
	// attribute is exported from `ec2.Vpc`, or manually found via the AWS Console.
	DefaultNetworkAclId pulumi.StringInput `pulumi:"defaultNetworkAclId"`
	// Specifies an egress rule. Parameters defined below.
	Egress DefaultNetworkAclEgressArrayInput `pulumi:"egress"`
	// Specifies an ingress rule. Parameters defined below.
	Ingress DefaultNetworkAclIngressArrayInput `pulumi:"ingress"`
	// A list of Subnet IDs to apply the ACL to. See the
	// notes below on managing Subnets in the Default Network ACL
	SubnetIds pulumi.StringArrayInput `pulumi:"subnetIds"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}
type DefaultNetworkAclEgress struct {
	// The action to take.
	Action string `pulumi:"action"`
	// The CIDR block to match. This must be a
	// valid network mask.
	CidrBlock *string `pulumi:"cidrBlock"`
	// The from port to match.
	FromPort int `pulumi:"fromPort"`
	// The ICMP type code to be used. Default 0.
	IcmpCode *int `pulumi:"icmpCode"`
	// The ICMP type to be used. Default 0.
	IcmpType *int `pulumi:"icmpType"`
	// The IPv6 CIDR block.
	Ipv6CidrBlock *string `pulumi:"ipv6CidrBlock"`
	// The protocol to match. If using the -1 'all'
	// protocol, you must specify a from and to port of 0.
	Protocol string `pulumi:"protocol"`
	// The rule number. Used for ordering.
	RuleNo int `pulumi:"ruleNo"`
	// The to port to match.
	ToPort int `pulumi:"toPort"`
}
var defaultNetworkAclEgressType = reflect.TypeOf((*DefaultNetworkAclEgress)(nil)).Elem()

type DefaultNetworkAclEgressInput interface {
	pulumi.Input

	ToDefaultNetworkAclEgressOutput() DefaultNetworkAclEgressOutput
	ToDefaultNetworkAclEgressOutputWithContext(ctx context.Context) DefaultNetworkAclEgressOutput
}

type DefaultNetworkAclEgressArgs struct {
	// The action to take.
	Action pulumi.StringInput `pulumi:"action"`
	// The CIDR block to match. This must be a
	// valid network mask.
	CidrBlock pulumi.StringInput `pulumi:"cidrBlock"`
	// The from port to match.
	FromPort pulumi.IntInput `pulumi:"fromPort"`
	// The ICMP type code to be used. Default 0.
	IcmpCode pulumi.IntInput `pulumi:"icmpCode"`
	// The ICMP type to be used. Default 0.
	IcmpType pulumi.IntInput `pulumi:"icmpType"`
	// The IPv6 CIDR block.
	Ipv6CidrBlock pulumi.StringInput `pulumi:"ipv6CidrBlock"`
	// The protocol to match. If using the -1 'all'
	// protocol, you must specify a from and to port of 0.
	Protocol pulumi.StringInput `pulumi:"protocol"`
	// The rule number. Used for ordering.
	RuleNo pulumi.IntInput `pulumi:"ruleNo"`
	// The to port to match.
	ToPort pulumi.IntInput `pulumi:"toPort"`
}

func (DefaultNetworkAclEgressArgs) ElementType() reflect.Type {
	return defaultNetworkAclEgressType
}

func (a DefaultNetworkAclEgressArgs) ToDefaultNetworkAclEgressOutput() DefaultNetworkAclEgressOutput {
	return pulumi.ToOutput(a).(DefaultNetworkAclEgressOutput)
}

func (a DefaultNetworkAclEgressArgs) ToDefaultNetworkAclEgressOutputWithContext(ctx context.Context) DefaultNetworkAclEgressOutput {
	return pulumi.ToOutputWithContext(ctx, a).(DefaultNetworkAclEgressOutput)
}

type DefaultNetworkAclEgressOutput struct { *pulumi.OutputState }

// The action to take.
func (o DefaultNetworkAclEgressOutput) Action() pulumi.StringOutput {
	return o.Apply(func(v DefaultNetworkAclEgress) string {
		return v.Action
	}).(pulumi.StringOutput)
}

// The CIDR block to match. This must be a
// valid network mask.
func (o DefaultNetworkAclEgressOutput) CidrBlock() pulumi.StringOutput {
	return o.Apply(func(v DefaultNetworkAclEgress) string {
		if v.CidrBlock == nil { return *new(string) } else { return *v.CidrBlock }
	}).(pulumi.StringOutput)
}

// The from port to match.
func (o DefaultNetworkAclEgressOutput) FromPort() pulumi.IntOutput {
	return o.Apply(func(v DefaultNetworkAclEgress) int {
		return v.FromPort
	}).(pulumi.IntOutput)
}

// The ICMP type code to be used. Default 0.
func (o DefaultNetworkAclEgressOutput) IcmpCode() pulumi.IntOutput {
	return o.Apply(func(v DefaultNetworkAclEgress) int {
		if v.IcmpCode == nil { return *new(int) } else { return *v.IcmpCode }
	}).(pulumi.IntOutput)
}

// The ICMP type to be used. Default 0.
func (o DefaultNetworkAclEgressOutput) IcmpType() pulumi.IntOutput {
	return o.Apply(func(v DefaultNetworkAclEgress) int {
		if v.IcmpType == nil { return *new(int) } else { return *v.IcmpType }
	}).(pulumi.IntOutput)
}

// The IPv6 CIDR block.
func (o DefaultNetworkAclEgressOutput) Ipv6CidrBlock() pulumi.StringOutput {
	return o.Apply(func(v DefaultNetworkAclEgress) string {
		if v.Ipv6CidrBlock == nil { return *new(string) } else { return *v.Ipv6CidrBlock }
	}).(pulumi.StringOutput)
}

// The protocol to match. If using the -1 'all'
// protocol, you must specify a from and to port of 0.
func (o DefaultNetworkAclEgressOutput) Protocol() pulumi.StringOutput {
	return o.Apply(func(v DefaultNetworkAclEgress) string {
		return v.Protocol
	}).(pulumi.StringOutput)
}

// The rule number. Used for ordering.
func (o DefaultNetworkAclEgressOutput) RuleNo() pulumi.IntOutput {
	return o.Apply(func(v DefaultNetworkAclEgress) int {
		return v.RuleNo
	}).(pulumi.IntOutput)
}

// The to port to match.
func (o DefaultNetworkAclEgressOutput) ToPort() pulumi.IntOutput {
	return o.Apply(func(v DefaultNetworkAclEgress) int {
		return v.ToPort
	}).(pulumi.IntOutput)
}

func (DefaultNetworkAclEgressOutput) ElementType() reflect.Type {
	return defaultNetworkAclEgressType
}

func (o DefaultNetworkAclEgressOutput) ToDefaultNetworkAclEgressOutput() DefaultNetworkAclEgressOutput {
	return o
}

func (o DefaultNetworkAclEgressOutput) ToDefaultNetworkAclEgressOutputWithContext(ctx context.Context) DefaultNetworkAclEgressOutput {
	return o
}

func init() { pulumi.RegisterOutputType(DefaultNetworkAclEgressOutput{}) }

var defaultNetworkAclEgressArrayType = reflect.TypeOf((*[]DefaultNetworkAclEgress)(nil)).Elem()

type DefaultNetworkAclEgressArrayInput interface {
	pulumi.Input

	ToDefaultNetworkAclEgressArrayOutput() DefaultNetworkAclEgressArrayOutput
	ToDefaultNetworkAclEgressArrayOutputWithContext(ctx context.Context) DefaultNetworkAclEgressArrayOutput
}

type DefaultNetworkAclEgressArrayArgs []DefaultNetworkAclEgressInput

func (DefaultNetworkAclEgressArrayArgs) ElementType() reflect.Type {
	return defaultNetworkAclEgressArrayType
}

func (a DefaultNetworkAclEgressArrayArgs) ToDefaultNetworkAclEgressArrayOutput() DefaultNetworkAclEgressArrayOutput {
	return pulumi.ToOutput(a).(DefaultNetworkAclEgressArrayOutput)
}

func (a DefaultNetworkAclEgressArrayArgs) ToDefaultNetworkAclEgressArrayOutputWithContext(ctx context.Context) DefaultNetworkAclEgressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(DefaultNetworkAclEgressArrayOutput)
}

type DefaultNetworkAclEgressArrayOutput struct { *pulumi.OutputState }

func (o DefaultNetworkAclEgressArrayOutput) Index(i pulumi.IntInput) DefaultNetworkAclEgressOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) DefaultNetworkAclEgress {
		return vs[0].([]DefaultNetworkAclEgress)[vs[1].(int)]
	}).(DefaultNetworkAclEgressOutput)
}

func (DefaultNetworkAclEgressArrayOutput) ElementType() reflect.Type {
	return defaultNetworkAclEgressArrayType
}

func (o DefaultNetworkAclEgressArrayOutput) ToDefaultNetworkAclEgressArrayOutput() DefaultNetworkAclEgressArrayOutput {
	return o
}

func (o DefaultNetworkAclEgressArrayOutput) ToDefaultNetworkAclEgressArrayOutputWithContext(ctx context.Context) DefaultNetworkAclEgressArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(DefaultNetworkAclEgressArrayOutput{}) }

type DefaultNetworkAclIngress struct {
	// The action to take.
	Action string `pulumi:"action"`
	// The CIDR block to match. This must be a
	// valid network mask.
	CidrBlock *string `pulumi:"cidrBlock"`
	// The from port to match.
	FromPort int `pulumi:"fromPort"`
	// The ICMP type code to be used. Default 0.
	IcmpCode *int `pulumi:"icmpCode"`
	// The ICMP type to be used. Default 0.
	IcmpType *int `pulumi:"icmpType"`
	// The IPv6 CIDR block.
	Ipv6CidrBlock *string `pulumi:"ipv6CidrBlock"`
	// The protocol to match. If using the -1 'all'
	// protocol, you must specify a from and to port of 0.
	Protocol string `pulumi:"protocol"`
	// The rule number. Used for ordering.
	RuleNo int `pulumi:"ruleNo"`
	// The to port to match.
	ToPort int `pulumi:"toPort"`
}
var defaultNetworkAclIngressType = reflect.TypeOf((*DefaultNetworkAclIngress)(nil)).Elem()

type DefaultNetworkAclIngressInput interface {
	pulumi.Input

	ToDefaultNetworkAclIngressOutput() DefaultNetworkAclIngressOutput
	ToDefaultNetworkAclIngressOutputWithContext(ctx context.Context) DefaultNetworkAclIngressOutput
}

type DefaultNetworkAclIngressArgs struct {
	// The action to take.
	Action pulumi.StringInput `pulumi:"action"`
	// The CIDR block to match. This must be a
	// valid network mask.
	CidrBlock pulumi.StringInput `pulumi:"cidrBlock"`
	// The from port to match.
	FromPort pulumi.IntInput `pulumi:"fromPort"`
	// The ICMP type code to be used. Default 0.
	IcmpCode pulumi.IntInput `pulumi:"icmpCode"`
	// The ICMP type to be used. Default 0.
	IcmpType pulumi.IntInput `pulumi:"icmpType"`
	// The IPv6 CIDR block.
	Ipv6CidrBlock pulumi.StringInput `pulumi:"ipv6CidrBlock"`
	// The protocol to match. If using the -1 'all'
	// protocol, you must specify a from and to port of 0.
	Protocol pulumi.StringInput `pulumi:"protocol"`
	// The rule number. Used for ordering.
	RuleNo pulumi.IntInput `pulumi:"ruleNo"`
	// The to port to match.
	ToPort pulumi.IntInput `pulumi:"toPort"`
}

func (DefaultNetworkAclIngressArgs) ElementType() reflect.Type {
	return defaultNetworkAclIngressType
}

func (a DefaultNetworkAclIngressArgs) ToDefaultNetworkAclIngressOutput() DefaultNetworkAclIngressOutput {
	return pulumi.ToOutput(a).(DefaultNetworkAclIngressOutput)
}

func (a DefaultNetworkAclIngressArgs) ToDefaultNetworkAclIngressOutputWithContext(ctx context.Context) DefaultNetworkAclIngressOutput {
	return pulumi.ToOutputWithContext(ctx, a).(DefaultNetworkAclIngressOutput)
}

type DefaultNetworkAclIngressOutput struct { *pulumi.OutputState }

// The action to take.
func (o DefaultNetworkAclIngressOutput) Action() pulumi.StringOutput {
	return o.Apply(func(v DefaultNetworkAclIngress) string {
		return v.Action
	}).(pulumi.StringOutput)
}

// The CIDR block to match. This must be a
// valid network mask.
func (o DefaultNetworkAclIngressOutput) CidrBlock() pulumi.StringOutput {
	return o.Apply(func(v DefaultNetworkAclIngress) string {
		if v.CidrBlock == nil { return *new(string) } else { return *v.CidrBlock }
	}).(pulumi.StringOutput)
}

// The from port to match.
func (o DefaultNetworkAclIngressOutput) FromPort() pulumi.IntOutput {
	return o.Apply(func(v DefaultNetworkAclIngress) int {
		return v.FromPort
	}).(pulumi.IntOutput)
}

// The ICMP type code to be used. Default 0.
func (o DefaultNetworkAclIngressOutput) IcmpCode() pulumi.IntOutput {
	return o.Apply(func(v DefaultNetworkAclIngress) int {
		if v.IcmpCode == nil { return *new(int) } else { return *v.IcmpCode }
	}).(pulumi.IntOutput)
}

// The ICMP type to be used. Default 0.
func (o DefaultNetworkAclIngressOutput) IcmpType() pulumi.IntOutput {
	return o.Apply(func(v DefaultNetworkAclIngress) int {
		if v.IcmpType == nil { return *new(int) } else { return *v.IcmpType }
	}).(pulumi.IntOutput)
}

// The IPv6 CIDR block.
func (o DefaultNetworkAclIngressOutput) Ipv6CidrBlock() pulumi.StringOutput {
	return o.Apply(func(v DefaultNetworkAclIngress) string {
		if v.Ipv6CidrBlock == nil { return *new(string) } else { return *v.Ipv6CidrBlock }
	}).(pulumi.StringOutput)
}

// The protocol to match. If using the -1 'all'
// protocol, you must specify a from and to port of 0.
func (o DefaultNetworkAclIngressOutput) Protocol() pulumi.StringOutput {
	return o.Apply(func(v DefaultNetworkAclIngress) string {
		return v.Protocol
	}).(pulumi.StringOutput)
}

// The rule number. Used for ordering.
func (o DefaultNetworkAclIngressOutput) RuleNo() pulumi.IntOutput {
	return o.Apply(func(v DefaultNetworkAclIngress) int {
		return v.RuleNo
	}).(pulumi.IntOutput)
}

// The to port to match.
func (o DefaultNetworkAclIngressOutput) ToPort() pulumi.IntOutput {
	return o.Apply(func(v DefaultNetworkAclIngress) int {
		return v.ToPort
	}).(pulumi.IntOutput)
}

func (DefaultNetworkAclIngressOutput) ElementType() reflect.Type {
	return defaultNetworkAclIngressType
}

func (o DefaultNetworkAclIngressOutput) ToDefaultNetworkAclIngressOutput() DefaultNetworkAclIngressOutput {
	return o
}

func (o DefaultNetworkAclIngressOutput) ToDefaultNetworkAclIngressOutputWithContext(ctx context.Context) DefaultNetworkAclIngressOutput {
	return o
}

func init() { pulumi.RegisterOutputType(DefaultNetworkAclIngressOutput{}) }

var defaultNetworkAclIngressArrayType = reflect.TypeOf((*[]DefaultNetworkAclIngress)(nil)).Elem()

type DefaultNetworkAclIngressArrayInput interface {
	pulumi.Input

	ToDefaultNetworkAclIngressArrayOutput() DefaultNetworkAclIngressArrayOutput
	ToDefaultNetworkAclIngressArrayOutputWithContext(ctx context.Context) DefaultNetworkAclIngressArrayOutput
}

type DefaultNetworkAclIngressArrayArgs []DefaultNetworkAclIngressInput

func (DefaultNetworkAclIngressArrayArgs) ElementType() reflect.Type {
	return defaultNetworkAclIngressArrayType
}

func (a DefaultNetworkAclIngressArrayArgs) ToDefaultNetworkAclIngressArrayOutput() DefaultNetworkAclIngressArrayOutput {
	return pulumi.ToOutput(a).(DefaultNetworkAclIngressArrayOutput)
}

func (a DefaultNetworkAclIngressArrayArgs) ToDefaultNetworkAclIngressArrayOutputWithContext(ctx context.Context) DefaultNetworkAclIngressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(DefaultNetworkAclIngressArrayOutput)
}

type DefaultNetworkAclIngressArrayOutput struct { *pulumi.OutputState }

func (o DefaultNetworkAclIngressArrayOutput) Index(i pulumi.IntInput) DefaultNetworkAclIngressOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) DefaultNetworkAclIngress {
		return vs[0].([]DefaultNetworkAclIngress)[vs[1].(int)]
	}).(DefaultNetworkAclIngressOutput)
}

func (DefaultNetworkAclIngressArrayOutput) ElementType() reflect.Type {
	return defaultNetworkAclIngressArrayType
}

func (o DefaultNetworkAclIngressArrayOutput) ToDefaultNetworkAclIngressArrayOutput() DefaultNetworkAclIngressArrayOutput {
	return o
}

func (o DefaultNetworkAclIngressArrayOutput) ToDefaultNetworkAclIngressArrayOutputWithContext(ctx context.Context) DefaultNetworkAclIngressArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(DefaultNetworkAclIngressArrayOutput{}) }

