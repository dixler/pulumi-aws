// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancingv2

import (
	"context"
	"reflect"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Target Group resource for use with Load Balancer resources.
// 
// > **Note:** `alb.TargetGroup` is known as `lb.TargetGroup`. The functionality is identical.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lb_target_group_legacy.html.markdown.
type TargetGroup struct {
	pulumi.CustomResourceState

	// The ARN of the Target Group (matches `id`)
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The ARN suffix for use with CloudWatch Metrics.
	ArnSuffix pulumi.StringOutput `pulumi:"arnSuffix"`

	// The amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused. The range is 0-3600 seconds. The default value is 300 seconds.
	DeregistrationDelay pulumi.IntOutput `pulumi:"deregistrationDelay"`

	// A Health Check block. Health Check blocks are documented below.
	HealthCheck TargetGroupHealthCheckOutput `pulumi:"healthCheck"`

	// Boolean whether the request and response headers exchanged between the load balancer and the Lambda function include arrays of values or strings. Only applies when `targetType` is `lambda`.
	LambdaMultiValueHeadersEnabled pulumi.BoolOutput `pulumi:"lambdaMultiValueHeadersEnabled"`

	// The name of the target group. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringOutput `pulumi:"name"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`. Cannot be longer than 6 characters.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`

	// The port to use to connect with the target. Valid values are either ports 1-65536, or `traffic-port`. Defaults to `traffic-port`.
	Port pulumi.IntOutput `pulumi:"port"`

	// The protocol to use to connect with the target. Defaults to `HTTP`. Not applicable when `targetType` is `lambda`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`

	// Boolean to enable / disable support for proxy protocol v2 on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#proxy-protocol) for more information.
	ProxyProtocolV2 pulumi.BoolOutput `pulumi:"proxyProtocolV2"`

	// The amount time for targets to warm up before the load balancer sends them a full share of requests. The range is 30-900 seconds or 0 to disable. The default value is 0 seconds.
	SlowStart pulumi.IntOutput `pulumi:"slowStart"`

	// A Stickiness block. Stickiness blocks are documented below. `stickiness` is only valid if used with Load Balancers of type `Application`
	Stickiness TargetGroupStickinessOutput `pulumi:"stickiness"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// The type of target that you must specify when registering targets with this target group.
	// The possible values are `instance` (targets are specified by instance ID) or `ip` (targets are specified by IP address) or `lambda` (targets are specified by lambda arn).
	// The default is `instance`. Note that you can't specify targets for a target group using both instance IDs and IP addresses.
	// If the target type is `ip`, specify IP addresses from the subnets of the virtual private cloud (VPC) for the target group,
	// the RFC 1918 range (10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16), and the RFC 6598 range (100.64.0.0/10).
	// You can't specify publicly routable IP addresses.
	TargetType pulumi.StringOutput `pulumi:"targetType"`

	// The identifier of the VPC in which to create the target group. Required when `targetType` is `instance` or `ip`. Does not apply when `targetType` is `lambda`.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewTargetGroup registers a new resource with the given unique name, arguments, and options.
func NewTargetGroup(ctx *pulumi.Context,
	name string, args *TargetGroupArgs, opts ...pulumi.ResourceOption) (*TargetGroup, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.DeregistrationDelay; i != nil { inputs["deregistrationDelay"] = i.ToIntOutput() }
		if i := args.HealthCheck; i != nil { inputs["healthCheck"] = i.ToTargetGroupHealthCheckOutput() }
		if i := args.LambdaMultiValueHeadersEnabled; i != nil { inputs["lambdaMultiValueHeadersEnabled"] = i.ToBoolOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.NamePrefix; i != nil { inputs["namePrefix"] = i.ToStringOutput() }
		if i := args.Port; i != nil { inputs["port"] = i.ToIntOutput() }
		if i := args.Protocol; i != nil { inputs["protocol"] = i.ToStringOutput() }
		if i := args.ProxyProtocolV2; i != nil { inputs["proxyProtocolV2"] = i.ToBoolOutput() }
		if i := args.SlowStart; i != nil { inputs["slowStart"] = i.ToIntOutput() }
		if i := args.Stickiness; i != nil { inputs["stickiness"] = i.ToTargetGroupStickinessOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := args.TargetType; i != nil { inputs["targetType"] = i.ToStringOutput() }
		if i := args.VpcId; i != nil { inputs["vpcId"] = i.ToStringOutput() }
	}
	var resource TargetGroup
	err := ctx.RegisterResource("aws:elasticloadbalancingv2/targetGroup:TargetGroup", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetGroup gets an existing TargetGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetGroupState, opts ...pulumi.ResourceOption) (*TargetGroup, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Arn; i != nil { inputs["arn"] = i.ToStringOutput() }
		if i := state.ArnSuffix; i != nil { inputs["arnSuffix"] = i.ToStringOutput() }
		if i := state.DeregistrationDelay; i != nil { inputs["deregistrationDelay"] = i.ToIntOutput() }
		if i := state.HealthCheck; i != nil { inputs["healthCheck"] = i.ToTargetGroupHealthCheckOutput() }
		if i := state.LambdaMultiValueHeadersEnabled; i != nil { inputs["lambdaMultiValueHeadersEnabled"] = i.ToBoolOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.NamePrefix; i != nil { inputs["namePrefix"] = i.ToStringOutput() }
		if i := state.Port; i != nil { inputs["port"] = i.ToIntOutput() }
		if i := state.Protocol; i != nil { inputs["protocol"] = i.ToStringOutput() }
		if i := state.ProxyProtocolV2; i != nil { inputs["proxyProtocolV2"] = i.ToBoolOutput() }
		if i := state.SlowStart; i != nil { inputs["slowStart"] = i.ToIntOutput() }
		if i := state.Stickiness; i != nil { inputs["stickiness"] = i.ToTargetGroupStickinessOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := state.TargetType; i != nil { inputs["targetType"] = i.ToStringOutput() }
		if i := state.VpcId; i != nil { inputs["vpcId"] = i.ToStringOutput() }
	}
	var resource TargetGroup
	err := ctx.ReadResource("aws:elasticloadbalancingv2/targetGroup:TargetGroup", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetGroup resources.
type TargetGroupState struct {
	// The ARN of the Target Group (matches `id`)
	Arn pulumi.StringInput `pulumi:"arn"`
	// The ARN suffix for use with CloudWatch Metrics.
	ArnSuffix pulumi.StringInput `pulumi:"arnSuffix"`
	// The amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused. The range is 0-3600 seconds. The default value is 300 seconds.
	DeregistrationDelay pulumi.IntInput `pulumi:"deregistrationDelay"`
	// A Health Check block. Health Check blocks are documented below.
	HealthCheck TargetGroupHealthCheckInput `pulumi:"healthCheck"`
	// Boolean whether the request and response headers exchanged between the load balancer and the Lambda function include arrays of values or strings. Only applies when `targetType` is `lambda`.
	LambdaMultiValueHeadersEnabled pulumi.BoolInput `pulumi:"lambdaMultiValueHeadersEnabled"`
	// The name of the target group. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringInput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`. Cannot be longer than 6 characters.
	NamePrefix pulumi.StringInput `pulumi:"namePrefix"`
	// The port to use to connect with the target. Valid values are either ports 1-65536, or `traffic-port`. Defaults to `traffic-port`.
	Port pulumi.IntInput `pulumi:"port"`
	// The protocol to use to connect with the target. Defaults to `HTTP`. Not applicable when `targetType` is `lambda`.
	Protocol pulumi.StringInput `pulumi:"protocol"`
	// Boolean to enable / disable support for proxy protocol v2 on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#proxy-protocol) for more information.
	ProxyProtocolV2 pulumi.BoolInput `pulumi:"proxyProtocolV2"`
	// The amount time for targets to warm up before the load balancer sends them a full share of requests. The range is 30-900 seconds or 0 to disable. The default value is 0 seconds.
	SlowStart pulumi.IntInput `pulumi:"slowStart"`
	// A Stickiness block. Stickiness blocks are documented below. `stickiness` is only valid if used with Load Balancers of type `Application`
	Stickiness TargetGroupStickinessInput `pulumi:"stickiness"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The type of target that you must specify when registering targets with this target group.
	// The possible values are `instance` (targets are specified by instance ID) or `ip` (targets are specified by IP address) or `lambda` (targets are specified by lambda arn).
	// The default is `instance`. Note that you can't specify targets for a target group using both instance IDs and IP addresses.
	// If the target type is `ip`, specify IP addresses from the subnets of the virtual private cloud (VPC) for the target group,
	// the RFC 1918 range (10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16), and the RFC 6598 range (100.64.0.0/10).
	// You can't specify publicly routable IP addresses.
	TargetType pulumi.StringInput `pulumi:"targetType"`
	// The identifier of the VPC in which to create the target group. Required when `targetType` is `instance` or `ip`. Does not apply when `targetType` is `lambda`.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
}

// The set of arguments for constructing a TargetGroup resource.
type TargetGroupArgs struct {
	// The amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused. The range is 0-3600 seconds. The default value is 300 seconds.
	DeregistrationDelay pulumi.IntInput `pulumi:"deregistrationDelay"`
	// A Health Check block. Health Check blocks are documented below.
	HealthCheck TargetGroupHealthCheckInput `pulumi:"healthCheck"`
	// Boolean whether the request and response headers exchanged between the load balancer and the Lambda function include arrays of values or strings. Only applies when `targetType` is `lambda`.
	LambdaMultiValueHeadersEnabled pulumi.BoolInput `pulumi:"lambdaMultiValueHeadersEnabled"`
	// The name of the target group. If omitted, this provider will assign a random, unique name.
	Name pulumi.StringInput `pulumi:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`. Cannot be longer than 6 characters.
	NamePrefix pulumi.StringInput `pulumi:"namePrefix"`
	// The port to use to connect with the target. Valid values are either ports 1-65536, or `traffic-port`. Defaults to `traffic-port`.
	Port pulumi.IntInput `pulumi:"port"`
	// The protocol to use to connect with the target. Defaults to `HTTP`. Not applicable when `targetType` is `lambda`.
	Protocol pulumi.StringInput `pulumi:"protocol"`
	// Boolean to enable / disable support for proxy protocol v2 on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#proxy-protocol) for more information.
	ProxyProtocolV2 pulumi.BoolInput `pulumi:"proxyProtocolV2"`
	// The amount time for targets to warm up before the load balancer sends them a full share of requests. The range is 30-900 seconds or 0 to disable. The default value is 0 seconds.
	SlowStart pulumi.IntInput `pulumi:"slowStart"`
	// A Stickiness block. Stickiness blocks are documented below. `stickiness` is only valid if used with Load Balancers of type `Application`
	Stickiness TargetGroupStickinessInput `pulumi:"stickiness"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The type of target that you must specify when registering targets with this target group.
	// The possible values are `instance` (targets are specified by instance ID) or `ip` (targets are specified by IP address) or `lambda` (targets are specified by lambda arn).
	// The default is `instance`. Note that you can't specify targets for a target group using both instance IDs and IP addresses.
	// If the target type is `ip`, specify IP addresses from the subnets of the virtual private cloud (VPC) for the target group,
	// the RFC 1918 range (10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16), and the RFC 6598 range (100.64.0.0/10).
	// You can't specify publicly routable IP addresses.
	TargetType pulumi.StringInput `pulumi:"targetType"`
	// The identifier of the VPC in which to create the target group. Required when `targetType` is `instance` or `ip`. Does not apply when `targetType` is `lambda`.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
}
type TargetGroupHealthCheck struct {
	// Indicates whether  health checks are enabled. Defaults to true.
	Enabled *bool `pulumi:"enabled"`
	// The number of consecutive health checks successes required before considering an unhealthy target healthy. Defaults to 3.
	HealthyThreshold *int `pulumi:"healthyThreshold"`
	// The approximate amount of time, in seconds, between health checks of an individual target. Minimum value 5 seconds, Maximum value 300 seconds. For `lambda` target groups, it needs to be greater as the `timeout` of the underlying `lambda`. Default 30 seconds.
	Interval *int `pulumi:"interval"`
	Matcher *string `pulumi:"matcher"`
	// The destination for the health check request. Applies to Application Load Balancers only (HTTP/HTTPS), not Network Load Balancers (TCP).
	Path *string `pulumi:"path"`
	// The port to use to connect with the target. Valid values are either ports 1-65536, or `traffic-port`. Defaults to `traffic-port`.
	Port *string `pulumi:"port"`
	// The protocol to use to connect with the target. Defaults to `HTTP`. Not applicable when `targetType` is `lambda`.
	Protocol *string `pulumi:"protocol"`
	// The amount of time, in seconds, during which no response means a failed health check. For Application Load Balancers, the range is 2 to 120 seconds, and the default is 5 seconds for the `instance` target type and 30 seconds for the `lambda` target type. For Network Load Balancers, you cannot set a custom value, and the default is 10 seconds for TCP and HTTPS health checks and 6 seconds for HTTP health checks.
	Timeout *int `pulumi:"timeout"`
	// The number of consecutive health check failures required before considering the target unhealthy . For Network Load Balancers, this value must be the same as the `healthyThreshold`. Defaults to 3.
	// * `matcher` (Required for HTTP/HTTPS ALB) The HTTP codes to use when checking for a successful response from a target. You can specify multiple values (for example, "200,202") or a range of values (for example, "200-299"). Applies to Application Load Balancers only (HTTP/HTTPS), not Network Load Balancers (TCP).
	UnhealthyThreshold *int `pulumi:"unhealthyThreshold"`
}
var targetGroupHealthCheckType = reflect.TypeOf((*TargetGroupHealthCheck)(nil)).Elem()

type TargetGroupHealthCheckInput interface {
	pulumi.Input

	ToTargetGroupHealthCheckOutput() TargetGroupHealthCheckOutput
	ToTargetGroupHealthCheckOutputWithContext(ctx context.Context) TargetGroupHealthCheckOutput
}

type TargetGroupHealthCheckArgs struct {
	// Indicates whether  health checks are enabled. Defaults to true.
	Enabled pulumi.BoolInput `pulumi:"enabled"`
	// The number of consecutive health checks successes required before considering an unhealthy target healthy. Defaults to 3.
	HealthyThreshold pulumi.IntInput `pulumi:"healthyThreshold"`
	// The approximate amount of time, in seconds, between health checks of an individual target. Minimum value 5 seconds, Maximum value 300 seconds. For `lambda` target groups, it needs to be greater as the `timeout` of the underlying `lambda`. Default 30 seconds.
	Interval pulumi.IntInput `pulumi:"interval"`
	Matcher pulumi.StringInput `pulumi:"matcher"`
	// The destination for the health check request. Applies to Application Load Balancers only (HTTP/HTTPS), not Network Load Balancers (TCP).
	Path pulumi.StringInput `pulumi:"path"`
	// The port to use to connect with the target. Valid values are either ports 1-65536, or `traffic-port`. Defaults to `traffic-port`.
	Port pulumi.StringInput `pulumi:"port"`
	// The protocol to use to connect with the target. Defaults to `HTTP`. Not applicable when `targetType` is `lambda`.
	Protocol pulumi.StringInput `pulumi:"protocol"`
	// The amount of time, in seconds, during which no response means a failed health check. For Application Load Balancers, the range is 2 to 120 seconds, and the default is 5 seconds for the `instance` target type and 30 seconds for the `lambda` target type. For Network Load Balancers, you cannot set a custom value, and the default is 10 seconds for TCP and HTTPS health checks and 6 seconds for HTTP health checks.
	Timeout pulumi.IntInput `pulumi:"timeout"`
	// The number of consecutive health check failures required before considering the target unhealthy . For Network Load Balancers, this value must be the same as the `healthyThreshold`. Defaults to 3.
	// * `matcher` (Required for HTTP/HTTPS ALB) The HTTP codes to use when checking for a successful response from a target. You can specify multiple values (for example, "200,202") or a range of values (for example, "200-299"). Applies to Application Load Balancers only (HTTP/HTTPS), not Network Load Balancers (TCP).
	UnhealthyThreshold pulumi.IntInput `pulumi:"unhealthyThreshold"`
}

func (TargetGroupHealthCheckArgs) ElementType() reflect.Type {
	return targetGroupHealthCheckType
}

func (a TargetGroupHealthCheckArgs) ToTargetGroupHealthCheckOutput() TargetGroupHealthCheckOutput {
	return pulumi.ToOutput(a).(TargetGroupHealthCheckOutput)
}

func (a TargetGroupHealthCheckArgs) ToTargetGroupHealthCheckOutputWithContext(ctx context.Context) TargetGroupHealthCheckOutput {
	return pulumi.ToOutputWithContext(ctx, a).(TargetGroupHealthCheckOutput)
}

type TargetGroupHealthCheckOutput struct { *pulumi.OutputState }

// Indicates whether  health checks are enabled. Defaults to true.
func (o TargetGroupHealthCheckOutput) Enabled() pulumi.BoolOutput {
	return o.Apply(func(v TargetGroupHealthCheck) bool {
		if v.Enabled == nil { return *new(bool) } else { return *v.Enabled }
	}).(pulumi.BoolOutput)
}

// The number of consecutive health checks successes required before considering an unhealthy target healthy. Defaults to 3.
func (o TargetGroupHealthCheckOutput) HealthyThreshold() pulumi.IntOutput {
	return o.Apply(func(v TargetGroupHealthCheck) int {
		if v.HealthyThreshold == nil { return *new(int) } else { return *v.HealthyThreshold }
	}).(pulumi.IntOutput)
}

// The approximate amount of time, in seconds, between health checks of an individual target. Minimum value 5 seconds, Maximum value 300 seconds. For `lambda` target groups, it needs to be greater as the `timeout` of the underlying `lambda`. Default 30 seconds.
func (o TargetGroupHealthCheckOutput) Interval() pulumi.IntOutput {
	return o.Apply(func(v TargetGroupHealthCheck) int {
		if v.Interval == nil { return *new(int) } else { return *v.Interval }
	}).(pulumi.IntOutput)
}

func (o TargetGroupHealthCheckOutput) Matcher() pulumi.StringOutput {
	return o.Apply(func(v TargetGroupHealthCheck) string {
		if v.Matcher == nil { return *new(string) } else { return *v.Matcher }
	}).(pulumi.StringOutput)
}

// The destination for the health check request. Applies to Application Load Balancers only (HTTP/HTTPS), not Network Load Balancers (TCP).
func (o TargetGroupHealthCheckOutput) Path() pulumi.StringOutput {
	return o.Apply(func(v TargetGroupHealthCheck) string {
		if v.Path == nil { return *new(string) } else { return *v.Path }
	}).(pulumi.StringOutput)
}

// The port to use to connect with the target. Valid values are either ports 1-65536, or `traffic-port`. Defaults to `traffic-port`.
func (o TargetGroupHealthCheckOutput) Port() pulumi.StringOutput {
	return o.Apply(func(v TargetGroupHealthCheck) string {
		if v.Port == nil { return *new(string) } else { return *v.Port }
	}).(pulumi.StringOutput)
}

// The protocol to use to connect with the target. Defaults to `HTTP`. Not applicable when `targetType` is `lambda`.
func (o TargetGroupHealthCheckOutput) Protocol() pulumi.StringOutput {
	return o.Apply(func(v TargetGroupHealthCheck) string {
		if v.Protocol == nil { return *new(string) } else { return *v.Protocol }
	}).(pulumi.StringOutput)
}

// The amount of time, in seconds, during which no response means a failed health check. For Application Load Balancers, the range is 2 to 120 seconds, and the default is 5 seconds for the `instance` target type and 30 seconds for the `lambda` target type. For Network Load Balancers, you cannot set a custom value, and the default is 10 seconds for TCP and HTTPS health checks and 6 seconds for HTTP health checks.
func (o TargetGroupHealthCheckOutput) Timeout() pulumi.IntOutput {
	return o.Apply(func(v TargetGroupHealthCheck) int {
		if v.Timeout == nil { return *new(int) } else { return *v.Timeout }
	}).(pulumi.IntOutput)
}

// The number of consecutive health check failures required before considering the target unhealthy . For Network Load Balancers, this value must be the same as the `healthyThreshold`. Defaults to 3.
// * `matcher` (Required for HTTP/HTTPS ALB) The HTTP codes to use when checking for a successful response from a target. You can specify multiple values (for example, "200,202") or a range of values (for example, "200-299"). Applies to Application Load Balancers only (HTTP/HTTPS), not Network Load Balancers (TCP).
func (o TargetGroupHealthCheckOutput) UnhealthyThreshold() pulumi.IntOutput {
	return o.Apply(func(v TargetGroupHealthCheck) int {
		if v.UnhealthyThreshold == nil { return *new(int) } else { return *v.UnhealthyThreshold }
	}).(pulumi.IntOutput)
}

func (TargetGroupHealthCheckOutput) ElementType() reflect.Type {
	return targetGroupHealthCheckType
}

func (o TargetGroupHealthCheckOutput) ToTargetGroupHealthCheckOutput() TargetGroupHealthCheckOutput {
	return o
}

func (o TargetGroupHealthCheckOutput) ToTargetGroupHealthCheckOutputWithContext(ctx context.Context) TargetGroupHealthCheckOutput {
	return o
}

func init() { pulumi.RegisterOutputType(TargetGroupHealthCheckOutput{}) }

type TargetGroupStickiness struct {
	// The time period, in seconds, during which requests from a client should be routed to the same target. After this time period expires, the load balancer-generated cookie is considered stale. The range is 1 second to 1 week (604800 seconds). The default value is 1 day (86400 seconds).
	CookieDuration *int `pulumi:"cookieDuration"`
	// Indicates whether  health checks are enabled. Defaults to true.
	Enabled *bool `pulumi:"enabled"`
	// The type of sticky sessions. The only current possible value is `lbCookie`.
	Type string `pulumi:"type"`
}
var targetGroupStickinessType = reflect.TypeOf((*TargetGroupStickiness)(nil)).Elem()

type TargetGroupStickinessInput interface {
	pulumi.Input

	ToTargetGroupStickinessOutput() TargetGroupStickinessOutput
	ToTargetGroupStickinessOutputWithContext(ctx context.Context) TargetGroupStickinessOutput
}

type TargetGroupStickinessArgs struct {
	// The time period, in seconds, during which requests from a client should be routed to the same target. After this time period expires, the load balancer-generated cookie is considered stale. The range is 1 second to 1 week (604800 seconds). The default value is 1 day (86400 seconds).
	CookieDuration pulumi.IntInput `pulumi:"cookieDuration"`
	// Indicates whether  health checks are enabled. Defaults to true.
	Enabled pulumi.BoolInput `pulumi:"enabled"`
	// The type of sticky sessions. The only current possible value is `lbCookie`.
	Type pulumi.StringInput `pulumi:"type"`
}

func (TargetGroupStickinessArgs) ElementType() reflect.Type {
	return targetGroupStickinessType
}

func (a TargetGroupStickinessArgs) ToTargetGroupStickinessOutput() TargetGroupStickinessOutput {
	return pulumi.ToOutput(a).(TargetGroupStickinessOutput)
}

func (a TargetGroupStickinessArgs) ToTargetGroupStickinessOutputWithContext(ctx context.Context) TargetGroupStickinessOutput {
	return pulumi.ToOutputWithContext(ctx, a).(TargetGroupStickinessOutput)
}

type TargetGroupStickinessOutput struct { *pulumi.OutputState }

// The time period, in seconds, during which requests from a client should be routed to the same target. After this time period expires, the load balancer-generated cookie is considered stale. The range is 1 second to 1 week (604800 seconds). The default value is 1 day (86400 seconds).
func (o TargetGroupStickinessOutput) CookieDuration() pulumi.IntOutput {
	return o.Apply(func(v TargetGroupStickiness) int {
		if v.CookieDuration == nil { return *new(int) } else { return *v.CookieDuration }
	}).(pulumi.IntOutput)
}

// Indicates whether  health checks are enabled. Defaults to true.
func (o TargetGroupStickinessOutput) Enabled() pulumi.BoolOutput {
	return o.Apply(func(v TargetGroupStickiness) bool {
		if v.Enabled == nil { return *new(bool) } else { return *v.Enabled }
	}).(pulumi.BoolOutput)
}

// The type of sticky sessions. The only current possible value is `lbCookie`.
func (o TargetGroupStickinessOutput) Type() pulumi.StringOutput {
	return o.Apply(func(v TargetGroupStickiness) string {
		return v.Type
	}).(pulumi.StringOutput)
}

func (TargetGroupStickinessOutput) ElementType() reflect.Type {
	return targetGroupStickinessType
}

func (o TargetGroupStickinessOutput) ToTargetGroupStickinessOutput() TargetGroupStickinessOutput {
	return o
}

func (o TargetGroupStickinessOutput) ToTargetGroupStickinessOutputWithContext(ctx context.Context) TargetGroupStickinessOutput {
	return o
}

func init() { pulumi.RegisterOutputType(TargetGroupStickinessOutput{}) }

