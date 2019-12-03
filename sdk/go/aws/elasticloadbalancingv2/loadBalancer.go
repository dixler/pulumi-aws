// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancingv2

import (
	"context"
	"reflect"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Load Balancer resource.
// 
// > **Note:** `alb.LoadBalancer` is known as `lb.LoadBalancer`. The functionality is identical.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lb_legacy.html.markdown.
type LoadBalancer struct {
	pulumi.CustomResourceState

	// An Access Logs block. Access Logs documented below.
	AccessLogs LoadBalancerAccessLogsOutput `pulumi:"accessLogs"`

	// The ARN of the load balancer (matches `id`).
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The ARN suffix for use with CloudWatch Metrics.
	ArnSuffix pulumi.StringOutput `pulumi:"arnSuffix"`

	// The DNS name of the load balancer.
	DnsName pulumi.StringOutput `pulumi:"dnsName"`

	// If true, cross-zone load balancing of the load balancer will be enabled.
	// This is a `network` load balancer feature. Defaults to `false`.
	EnableCrossZoneLoadBalancing pulumi.BoolOutput `pulumi:"enableCrossZoneLoadBalancing"`

	// If true, deletion of the load balancer will be disabled via
	// the AWS API. This will prevent this provider from deleting the load balancer. Defaults to `false`.
	EnableDeletionProtection pulumi.BoolOutput `pulumi:"enableDeletionProtection"`

	// Indicates whether HTTP/2 is enabled in `application` load balancers. Defaults to `true`.
	EnableHttp2 pulumi.BoolOutput `pulumi:"enableHttp2"`

	// The time in seconds that the connection is allowed to be idle. Only valid for Load Balancers of type `application`. Default: 60.
	IdleTimeout pulumi.IntOutput `pulumi:"idleTimeout"`

	// If true, the LB will be internal.
	Internal pulumi.BoolOutput `pulumi:"internal"`

	// The type of IP addresses used by the subnets for your load balancer. The possible values are `ipv4` and `dualstack`
	IpAddressType pulumi.StringOutput `pulumi:"ipAddressType"`

	// The type of load balancer to create. Possible values are `application` or `network`. The default value is `application`.
	LoadBalancerType pulumi.StringOutput `pulumi:"loadBalancerType"`

	// The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters,
	// must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen. If not specified,
	// this provider will autogenerate a name beginning with `tf-lb`.
	Name pulumi.StringOutput `pulumi:"name"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`

	// A list of security group IDs to assign to the LB. Only valid for Load Balancers of type `application`.
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`

	// A subnet mapping block as documented below.
	SubnetMappings LoadBalancerSubnetMappingsArrayOutput `pulumi:"subnetMappings"`

	// A list of subnet IDs to attach to the LB. Subnets
	// cannot be updated for Load Balancers of type `network`. Changing this value
	// for load balancers of type `network` will force a recreation of the resource.
	Subnets pulumi.StringArrayOutput `pulumi:"subnets"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`

	VpcId pulumi.StringOutput `pulumi:"vpcId"`

	// The canonical hosted zone ID of the load balancer (to be used in a Route 53 Alias record).
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewLoadBalancer registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancer(ctx *pulumi.Context,
	name string, args *LoadBalancerArgs, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.AccessLogs; i != nil { inputs["accessLogs"] = i.ToLoadBalancerAccessLogsOutput() }
		if i := args.EnableCrossZoneLoadBalancing; i != nil { inputs["enableCrossZoneLoadBalancing"] = i.ToBoolOutput() }
		if i := args.EnableDeletionProtection; i != nil { inputs["enableDeletionProtection"] = i.ToBoolOutput() }
		if i := args.EnableHttp2; i != nil { inputs["enableHttp2"] = i.ToBoolOutput() }
		if i := args.IdleTimeout; i != nil { inputs["idleTimeout"] = i.ToIntOutput() }
		if i := args.Internal; i != nil { inputs["internal"] = i.ToBoolOutput() }
		if i := args.IpAddressType; i != nil { inputs["ipAddressType"] = i.ToStringOutput() }
		if i := args.LoadBalancerType; i != nil { inputs["loadBalancerType"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.NamePrefix; i != nil { inputs["namePrefix"] = i.ToStringOutput() }
		if i := args.SecurityGroups; i != nil { inputs["securityGroups"] = i.ToStringArrayOutput() }
		if i := args.SubnetMappings; i != nil { inputs["subnetMappings"] = i.ToLoadBalancerSubnetMappingsArrayOutput() }
		if i := args.Subnets; i != nil { inputs["subnets"] = i.ToStringArrayOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource LoadBalancer
	err := ctx.RegisterResource("aws:elasticloadbalancingv2/loadBalancer:LoadBalancer", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadBalancer gets an existing LoadBalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadBalancerState, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.AccessLogs; i != nil { inputs["accessLogs"] = i.ToLoadBalancerAccessLogsOutput() }
		if i := state.Arn; i != nil { inputs["arn"] = i.ToStringOutput() }
		if i := state.ArnSuffix; i != nil { inputs["arnSuffix"] = i.ToStringOutput() }
		if i := state.DnsName; i != nil { inputs["dnsName"] = i.ToStringOutput() }
		if i := state.EnableCrossZoneLoadBalancing; i != nil { inputs["enableCrossZoneLoadBalancing"] = i.ToBoolOutput() }
		if i := state.EnableDeletionProtection; i != nil { inputs["enableDeletionProtection"] = i.ToBoolOutput() }
		if i := state.EnableHttp2; i != nil { inputs["enableHttp2"] = i.ToBoolOutput() }
		if i := state.IdleTimeout; i != nil { inputs["idleTimeout"] = i.ToIntOutput() }
		if i := state.Internal; i != nil { inputs["internal"] = i.ToBoolOutput() }
		if i := state.IpAddressType; i != nil { inputs["ipAddressType"] = i.ToStringOutput() }
		if i := state.LoadBalancerType; i != nil { inputs["loadBalancerType"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.NamePrefix; i != nil { inputs["namePrefix"] = i.ToStringOutput() }
		if i := state.SecurityGroups; i != nil { inputs["securityGroups"] = i.ToStringArrayOutput() }
		if i := state.SubnetMappings; i != nil { inputs["subnetMappings"] = i.ToLoadBalancerSubnetMappingsArrayOutput() }
		if i := state.Subnets; i != nil { inputs["subnets"] = i.ToStringArrayOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := state.VpcId; i != nil { inputs["vpcId"] = i.ToStringOutput() }
		if i := state.ZoneId; i != nil { inputs["zoneId"] = i.ToStringOutput() }
	}
	var resource LoadBalancer
	err := ctx.ReadResource("aws:elasticloadbalancingv2/loadBalancer:LoadBalancer", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoadBalancer resources.
type LoadBalancerState struct {
	// An Access Logs block. Access Logs documented below.
	AccessLogs LoadBalancerAccessLogsInput `pulumi:"accessLogs"`
	// The ARN of the load balancer (matches `id`).
	Arn pulumi.StringInput `pulumi:"arn"`
	// The ARN suffix for use with CloudWatch Metrics.
	ArnSuffix pulumi.StringInput `pulumi:"arnSuffix"`
	// The DNS name of the load balancer.
	DnsName pulumi.StringInput `pulumi:"dnsName"`
	// If true, cross-zone load balancing of the load balancer will be enabled.
	// This is a `network` load balancer feature. Defaults to `false`.
	EnableCrossZoneLoadBalancing pulumi.BoolInput `pulumi:"enableCrossZoneLoadBalancing"`
	// If true, deletion of the load balancer will be disabled via
	// the AWS API. This will prevent this provider from deleting the load balancer. Defaults to `false`.
	EnableDeletionProtection pulumi.BoolInput `pulumi:"enableDeletionProtection"`
	// Indicates whether HTTP/2 is enabled in `application` load balancers. Defaults to `true`.
	EnableHttp2 pulumi.BoolInput `pulumi:"enableHttp2"`
	// The time in seconds that the connection is allowed to be idle. Only valid for Load Balancers of type `application`. Default: 60.
	IdleTimeout pulumi.IntInput `pulumi:"idleTimeout"`
	// If true, the LB will be internal.
	Internal pulumi.BoolInput `pulumi:"internal"`
	// The type of IP addresses used by the subnets for your load balancer. The possible values are `ipv4` and `dualstack`
	IpAddressType pulumi.StringInput `pulumi:"ipAddressType"`
	// The type of load balancer to create. Possible values are `application` or `network`. The default value is `application`.
	LoadBalancerType pulumi.StringInput `pulumi:"loadBalancerType"`
	// The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters,
	// must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen. If not specified,
	// this provider will autogenerate a name beginning with `tf-lb`.
	Name pulumi.StringInput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringInput `pulumi:"namePrefix"`
	// A list of security group IDs to assign to the LB. Only valid for Load Balancers of type `application`.
	SecurityGroups pulumi.StringArrayInput `pulumi:"securityGroups"`
	// A subnet mapping block as documented below.
	SubnetMappings LoadBalancerSubnetMappingsArrayInput `pulumi:"subnetMappings"`
	// A list of subnet IDs to attach to the LB. Subnets
	// cannot be updated for Load Balancers of type `network`. Changing this value
	// for load balancers of type `network` will force a recreation of the resource.
	Subnets pulumi.StringArrayInput `pulumi:"subnets"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	VpcId pulumi.StringInput `pulumi:"vpcId"`
	// The canonical hosted zone ID of the load balancer (to be used in a Route 53 Alias record).
	ZoneId pulumi.StringInput `pulumi:"zoneId"`
}

// The set of arguments for constructing a LoadBalancer resource.
type LoadBalancerArgs struct {
	// An Access Logs block. Access Logs documented below.
	AccessLogs LoadBalancerAccessLogsInput `pulumi:"accessLogs"`
	// If true, cross-zone load balancing of the load balancer will be enabled.
	// This is a `network` load balancer feature. Defaults to `false`.
	EnableCrossZoneLoadBalancing pulumi.BoolInput `pulumi:"enableCrossZoneLoadBalancing"`
	// If true, deletion of the load balancer will be disabled via
	// the AWS API. This will prevent this provider from deleting the load balancer. Defaults to `false`.
	EnableDeletionProtection pulumi.BoolInput `pulumi:"enableDeletionProtection"`
	// Indicates whether HTTP/2 is enabled in `application` load balancers. Defaults to `true`.
	EnableHttp2 pulumi.BoolInput `pulumi:"enableHttp2"`
	// The time in seconds that the connection is allowed to be idle. Only valid for Load Balancers of type `application`. Default: 60.
	IdleTimeout pulumi.IntInput `pulumi:"idleTimeout"`
	// If true, the LB will be internal.
	Internal pulumi.BoolInput `pulumi:"internal"`
	// The type of IP addresses used by the subnets for your load balancer. The possible values are `ipv4` and `dualstack`
	IpAddressType pulumi.StringInput `pulumi:"ipAddressType"`
	// The type of load balancer to create. Possible values are `application` or `network`. The default value is `application`.
	LoadBalancerType pulumi.StringInput `pulumi:"loadBalancerType"`
	// The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters,
	// must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen. If not specified,
	// this provider will autogenerate a name beginning with `tf-lb`.
	Name pulumi.StringInput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringInput `pulumi:"namePrefix"`
	// A list of security group IDs to assign to the LB. Only valid for Load Balancers of type `application`.
	SecurityGroups pulumi.StringArrayInput `pulumi:"securityGroups"`
	// A subnet mapping block as documented below.
	SubnetMappings LoadBalancerSubnetMappingsArrayInput `pulumi:"subnetMappings"`
	// A list of subnet IDs to attach to the LB. Subnets
	// cannot be updated for Load Balancers of type `network`. Changing this value
	// for load balancers of type `network` will force a recreation of the resource.
	Subnets pulumi.StringArrayInput `pulumi:"subnets"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}
type LoadBalancerAccessLogs struct {
	// The S3 bucket name to store the logs in.
	Bucket string `pulumi:"bucket"`
	// Boolean to enable / disable `accessLogs`. Defaults to `false`, even when `bucket` is specified.
	Enabled *bool `pulumi:"enabled"`
	// The S3 bucket prefix. Logs are stored in the root if not configured.
	Prefix *string `pulumi:"prefix"`
}
var loadBalancerAccessLogsType = reflect.TypeOf((*LoadBalancerAccessLogs)(nil)).Elem()

type LoadBalancerAccessLogsInput interface {
	pulumi.Input

	ToLoadBalancerAccessLogsOutput() LoadBalancerAccessLogsOutput
	ToLoadBalancerAccessLogsOutputWithContext(ctx context.Context) LoadBalancerAccessLogsOutput
}

type LoadBalancerAccessLogsArgs struct {
	// The S3 bucket name to store the logs in.
	Bucket pulumi.StringInput `pulumi:"bucket"`
	// Boolean to enable / disable `accessLogs`. Defaults to `false`, even when `bucket` is specified.
	Enabled pulumi.BoolInput `pulumi:"enabled"`
	// The S3 bucket prefix. Logs are stored in the root if not configured.
	Prefix pulumi.StringInput `pulumi:"prefix"`
}

func (LoadBalancerAccessLogsArgs) ElementType() reflect.Type {
	return loadBalancerAccessLogsType
}

func (a LoadBalancerAccessLogsArgs) ToLoadBalancerAccessLogsOutput() LoadBalancerAccessLogsOutput {
	return pulumi.ToOutput(a).(LoadBalancerAccessLogsOutput)
}

func (a LoadBalancerAccessLogsArgs) ToLoadBalancerAccessLogsOutputWithContext(ctx context.Context) LoadBalancerAccessLogsOutput {
	return pulumi.ToOutputWithContext(ctx, a).(LoadBalancerAccessLogsOutput)
}

type LoadBalancerAccessLogsOutput struct { *pulumi.OutputState }

// The S3 bucket name to store the logs in.
func (o LoadBalancerAccessLogsOutput) Bucket() pulumi.StringOutput {
	return o.Apply(func(v LoadBalancerAccessLogs) string {
		return v.Bucket
	}).(pulumi.StringOutput)
}

// Boolean to enable / disable `accessLogs`. Defaults to `false`, even when `bucket` is specified.
func (o LoadBalancerAccessLogsOutput) Enabled() pulumi.BoolOutput {
	return o.Apply(func(v LoadBalancerAccessLogs) bool {
		if v.Enabled == nil { return *new(bool) } else { return *v.Enabled }
	}).(pulumi.BoolOutput)
}

// The S3 bucket prefix. Logs are stored in the root if not configured.
func (o LoadBalancerAccessLogsOutput) Prefix() pulumi.StringOutput {
	return o.Apply(func(v LoadBalancerAccessLogs) string {
		if v.Prefix == nil { return *new(string) } else { return *v.Prefix }
	}).(pulumi.StringOutput)
}

func (LoadBalancerAccessLogsOutput) ElementType() reflect.Type {
	return loadBalancerAccessLogsType
}

func (o LoadBalancerAccessLogsOutput) ToLoadBalancerAccessLogsOutput() LoadBalancerAccessLogsOutput {
	return o
}

func (o LoadBalancerAccessLogsOutput) ToLoadBalancerAccessLogsOutputWithContext(ctx context.Context) LoadBalancerAccessLogsOutput {
	return o
}

func init() { pulumi.RegisterOutputType(LoadBalancerAccessLogsOutput{}) }

type LoadBalancerSubnetMappings struct {
	// The allocation ID of the Elastic IP address.
	AllocationId *string `pulumi:"allocationId"`
	// The id of the subnet of which to attach to the load balancer. You can specify only one subnet per Availability Zone.
	SubnetId string `pulumi:"subnetId"`
}
var loadBalancerSubnetMappingsType = reflect.TypeOf((*LoadBalancerSubnetMappings)(nil)).Elem()

type LoadBalancerSubnetMappingsInput interface {
	pulumi.Input

	ToLoadBalancerSubnetMappingsOutput() LoadBalancerSubnetMappingsOutput
	ToLoadBalancerSubnetMappingsOutputWithContext(ctx context.Context) LoadBalancerSubnetMappingsOutput
}

type LoadBalancerSubnetMappingsArgs struct {
	// The allocation ID of the Elastic IP address.
	AllocationId pulumi.StringInput `pulumi:"allocationId"`
	// The id of the subnet of which to attach to the load balancer. You can specify only one subnet per Availability Zone.
	SubnetId pulumi.StringInput `pulumi:"subnetId"`
}

func (LoadBalancerSubnetMappingsArgs) ElementType() reflect.Type {
	return loadBalancerSubnetMappingsType
}

func (a LoadBalancerSubnetMappingsArgs) ToLoadBalancerSubnetMappingsOutput() LoadBalancerSubnetMappingsOutput {
	return pulumi.ToOutput(a).(LoadBalancerSubnetMappingsOutput)
}

func (a LoadBalancerSubnetMappingsArgs) ToLoadBalancerSubnetMappingsOutputWithContext(ctx context.Context) LoadBalancerSubnetMappingsOutput {
	return pulumi.ToOutputWithContext(ctx, a).(LoadBalancerSubnetMappingsOutput)
}

type LoadBalancerSubnetMappingsOutput struct { *pulumi.OutputState }

// The allocation ID of the Elastic IP address.
func (o LoadBalancerSubnetMappingsOutput) AllocationId() pulumi.StringOutput {
	return o.Apply(func(v LoadBalancerSubnetMappings) string {
		if v.AllocationId == nil { return *new(string) } else { return *v.AllocationId }
	}).(pulumi.StringOutput)
}

// The id of the subnet of which to attach to the load balancer. You can specify only one subnet per Availability Zone.
func (o LoadBalancerSubnetMappingsOutput) SubnetId() pulumi.StringOutput {
	return o.Apply(func(v LoadBalancerSubnetMappings) string {
		return v.SubnetId
	}).(pulumi.StringOutput)
}

func (LoadBalancerSubnetMappingsOutput) ElementType() reflect.Type {
	return loadBalancerSubnetMappingsType
}

func (o LoadBalancerSubnetMappingsOutput) ToLoadBalancerSubnetMappingsOutput() LoadBalancerSubnetMappingsOutput {
	return o
}

func (o LoadBalancerSubnetMappingsOutput) ToLoadBalancerSubnetMappingsOutputWithContext(ctx context.Context) LoadBalancerSubnetMappingsOutput {
	return o
}

func init() { pulumi.RegisterOutputType(LoadBalancerSubnetMappingsOutput{}) }

var loadBalancerSubnetMappingsArrayType = reflect.TypeOf((*[]LoadBalancerSubnetMappings)(nil)).Elem()

type LoadBalancerSubnetMappingsArrayInput interface {
	pulumi.Input

	ToLoadBalancerSubnetMappingsArrayOutput() LoadBalancerSubnetMappingsArrayOutput
	ToLoadBalancerSubnetMappingsArrayOutputWithContext(ctx context.Context) LoadBalancerSubnetMappingsArrayOutput
}

type LoadBalancerSubnetMappingsArrayArgs []LoadBalancerSubnetMappingsInput

func (LoadBalancerSubnetMappingsArrayArgs) ElementType() reflect.Type {
	return loadBalancerSubnetMappingsArrayType
}

func (a LoadBalancerSubnetMappingsArrayArgs) ToLoadBalancerSubnetMappingsArrayOutput() LoadBalancerSubnetMappingsArrayOutput {
	return pulumi.ToOutput(a).(LoadBalancerSubnetMappingsArrayOutput)
}

func (a LoadBalancerSubnetMappingsArrayArgs) ToLoadBalancerSubnetMappingsArrayOutputWithContext(ctx context.Context) LoadBalancerSubnetMappingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(LoadBalancerSubnetMappingsArrayOutput)
}

type LoadBalancerSubnetMappingsArrayOutput struct { *pulumi.OutputState }

func (o LoadBalancerSubnetMappingsArrayOutput) Index(i pulumi.IntInput) LoadBalancerSubnetMappingsOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) LoadBalancerSubnetMappings {
		return vs[0].([]LoadBalancerSubnetMappings)[vs[1].(int)]
	}).(LoadBalancerSubnetMappingsOutput)
}

func (LoadBalancerSubnetMappingsArrayOutput) ElementType() reflect.Type {
	return loadBalancerSubnetMappingsArrayType
}

func (o LoadBalancerSubnetMappingsArrayOutput) ToLoadBalancerSubnetMappingsArrayOutput() LoadBalancerSubnetMappingsArrayOutput {
	return o
}

func (o LoadBalancerSubnetMappingsArrayOutput) ToLoadBalancerSubnetMappingsArrayOutputWithContext(ctx context.Context) LoadBalancerSubnetMappingsArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(LoadBalancerSubnetMappingsArrayOutput{}) }

