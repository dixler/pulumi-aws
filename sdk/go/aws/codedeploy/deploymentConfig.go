// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codedeploy

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a CodeDeploy deployment config for an application
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/codedeploy_deployment_config.html.markdown.
type DeploymentConfig struct {
	pulumi.CustomResourceState

	// The compute platform can be `Server`, `Lambda`, or `ECS`. Default is `Server`.
	ComputePlatform pulumi.StringOutput `pulumi:"computePlatform"`

	// The AWS Assigned deployment config id
	DeploymentConfigId pulumi.StringOutput `pulumi:"deploymentConfigId"`

	// The name of the deployment config.
	DeploymentConfigName pulumi.StringOutput `pulumi:"deploymentConfigName"`

	// A minimumHealthyHosts block. Minimum Healthy Hosts are documented below.
	MinimumHealthyHosts DeploymentConfigMinimumHealthyHostsOutput `pulumi:"minimumHealthyHosts"`

	// A trafficRoutingConfig block. Traffic Routing Config is documented below.
	TrafficRoutingConfig DeploymentConfigTrafficRoutingConfigOutput `pulumi:"trafficRoutingConfig"`
}

// NewDeploymentConfig registers a new resource with the given unique name, arguments, and options.
func NewDeploymentConfig(ctx *pulumi.Context,
	name string, args *DeploymentConfigArgs, opts ...pulumi.ResourceOption) (*DeploymentConfig, error) {
	if args == nil || args.DeploymentConfigName == nil {
		return nil, errors.New("missing required argument 'DeploymentConfigName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.ComputePlatform; i != nil { inputs["computePlatform"] = i.ToStringOutput() }
		if i := args.DeploymentConfigName; i != nil { inputs["deploymentConfigName"] = i.ToStringOutput() }
		if i := args.MinimumHealthyHosts; i != nil { inputs["minimumHealthyHosts"] = i.ToDeploymentConfigMinimumHealthyHostsOutput() }
		if i := args.TrafficRoutingConfig; i != nil { inputs["trafficRoutingConfig"] = i.ToDeploymentConfigTrafficRoutingConfigOutput() }
	}
	var resource DeploymentConfig
	err := ctx.RegisterResource("aws:codedeploy/deploymentConfig:DeploymentConfig", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeploymentConfig gets an existing DeploymentConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeploymentConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeploymentConfigState, opts ...pulumi.ResourceOption) (*DeploymentConfig, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.ComputePlatform; i != nil { inputs["computePlatform"] = i.ToStringOutput() }
		if i := state.DeploymentConfigId; i != nil { inputs["deploymentConfigId"] = i.ToStringOutput() }
		if i := state.DeploymentConfigName; i != nil { inputs["deploymentConfigName"] = i.ToStringOutput() }
		if i := state.MinimumHealthyHosts; i != nil { inputs["minimumHealthyHosts"] = i.ToDeploymentConfigMinimumHealthyHostsOutput() }
		if i := state.TrafficRoutingConfig; i != nil { inputs["trafficRoutingConfig"] = i.ToDeploymentConfigTrafficRoutingConfigOutput() }
	}
	var resource DeploymentConfig
	err := ctx.ReadResource("aws:codedeploy/deploymentConfig:DeploymentConfig", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeploymentConfig resources.
type DeploymentConfigState struct {
	// The compute platform can be `Server`, `Lambda`, or `ECS`. Default is `Server`.
	ComputePlatform pulumi.StringInput `pulumi:"computePlatform"`
	// The AWS Assigned deployment config id
	DeploymentConfigId pulumi.StringInput `pulumi:"deploymentConfigId"`
	// The name of the deployment config.
	DeploymentConfigName pulumi.StringInput `pulumi:"deploymentConfigName"`
	// A minimumHealthyHosts block. Minimum Healthy Hosts are documented below.
	MinimumHealthyHosts DeploymentConfigMinimumHealthyHostsInput `pulumi:"minimumHealthyHosts"`
	// A trafficRoutingConfig block. Traffic Routing Config is documented below.
	TrafficRoutingConfig DeploymentConfigTrafficRoutingConfigInput `pulumi:"trafficRoutingConfig"`
}

// The set of arguments for constructing a DeploymentConfig resource.
type DeploymentConfigArgs struct {
	// The compute platform can be `Server`, `Lambda`, or `ECS`. Default is `Server`.
	ComputePlatform pulumi.StringInput `pulumi:"computePlatform"`
	// The name of the deployment config.
	DeploymentConfigName pulumi.StringInput `pulumi:"deploymentConfigName"`
	// A minimumHealthyHosts block. Minimum Healthy Hosts are documented below.
	MinimumHealthyHosts DeploymentConfigMinimumHealthyHostsInput `pulumi:"minimumHealthyHosts"`
	// A trafficRoutingConfig block. Traffic Routing Config is documented below.
	TrafficRoutingConfig DeploymentConfigTrafficRoutingConfigInput `pulumi:"trafficRoutingConfig"`
}
type DeploymentConfigMinimumHealthyHosts struct {
	Type *string `pulumi:"type"`
	Value *int `pulumi:"value"`
}
var deploymentConfigMinimumHealthyHostsType = reflect.TypeOf((*DeploymentConfigMinimumHealthyHosts)(nil)).Elem()

type DeploymentConfigMinimumHealthyHostsInput interface {
	pulumi.Input

	ToDeploymentConfigMinimumHealthyHostsOutput() DeploymentConfigMinimumHealthyHostsOutput
	ToDeploymentConfigMinimumHealthyHostsOutputWithContext(ctx context.Context) DeploymentConfigMinimumHealthyHostsOutput
}

type DeploymentConfigMinimumHealthyHostsArgs struct {
	Type pulumi.StringInput `pulumi:"type"`
	Value pulumi.IntInput `pulumi:"value"`
}

func (DeploymentConfigMinimumHealthyHostsArgs) ElementType() reflect.Type {
	return deploymentConfigMinimumHealthyHostsType
}

func (a DeploymentConfigMinimumHealthyHostsArgs) ToDeploymentConfigMinimumHealthyHostsOutput() DeploymentConfigMinimumHealthyHostsOutput {
	return pulumi.ToOutput(a).(DeploymentConfigMinimumHealthyHostsOutput)
}

func (a DeploymentConfigMinimumHealthyHostsArgs) ToDeploymentConfigMinimumHealthyHostsOutputWithContext(ctx context.Context) DeploymentConfigMinimumHealthyHostsOutput {
	return pulumi.ToOutputWithContext(ctx, a).(DeploymentConfigMinimumHealthyHostsOutput)
}

type DeploymentConfigMinimumHealthyHostsOutput struct { *pulumi.OutputState }

func (o DeploymentConfigMinimumHealthyHostsOutput) Type() pulumi.StringOutput {
	return o.Apply(func(v DeploymentConfigMinimumHealthyHosts) string {
		if v.Type == nil { return *new(string) } else { return *v.Type }
	}).(pulumi.StringOutput)
}

func (o DeploymentConfigMinimumHealthyHostsOutput) Value() pulumi.IntOutput {
	return o.Apply(func(v DeploymentConfigMinimumHealthyHosts) int {
		if v.Value == nil { return *new(int) } else { return *v.Value }
	}).(pulumi.IntOutput)
}

func (DeploymentConfigMinimumHealthyHostsOutput) ElementType() reflect.Type {
	return deploymentConfigMinimumHealthyHostsType
}

func (o DeploymentConfigMinimumHealthyHostsOutput) ToDeploymentConfigMinimumHealthyHostsOutput() DeploymentConfigMinimumHealthyHostsOutput {
	return o
}

func (o DeploymentConfigMinimumHealthyHostsOutput) ToDeploymentConfigMinimumHealthyHostsOutputWithContext(ctx context.Context) DeploymentConfigMinimumHealthyHostsOutput {
	return o
}

func init() { pulumi.RegisterOutputType(DeploymentConfigMinimumHealthyHostsOutput{}) }

type DeploymentConfigTrafficRoutingConfig struct {
	TimeBasedCanary *DeploymentConfigTrafficRoutingConfigTimeBasedCanary `pulumi:"timeBasedCanary"`
	TimeBasedLinear *DeploymentConfigTrafficRoutingConfigTimeBasedLinear `pulumi:"timeBasedLinear"`
	Type *string `pulumi:"type"`
}
var deploymentConfigTrafficRoutingConfigType = reflect.TypeOf((*DeploymentConfigTrafficRoutingConfig)(nil)).Elem()

type DeploymentConfigTrafficRoutingConfigInput interface {
	pulumi.Input

	ToDeploymentConfigTrafficRoutingConfigOutput() DeploymentConfigTrafficRoutingConfigOutput
	ToDeploymentConfigTrafficRoutingConfigOutputWithContext(ctx context.Context) DeploymentConfigTrafficRoutingConfigOutput
}

type DeploymentConfigTrafficRoutingConfigArgs struct {
	TimeBasedCanary DeploymentConfigTrafficRoutingConfigTimeBasedCanaryInput `pulumi:"timeBasedCanary"`
	TimeBasedLinear DeploymentConfigTrafficRoutingConfigTimeBasedLinearInput `pulumi:"timeBasedLinear"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (DeploymentConfigTrafficRoutingConfigArgs) ElementType() reflect.Type {
	return deploymentConfigTrafficRoutingConfigType
}

func (a DeploymentConfigTrafficRoutingConfigArgs) ToDeploymentConfigTrafficRoutingConfigOutput() DeploymentConfigTrafficRoutingConfigOutput {
	return pulumi.ToOutput(a).(DeploymentConfigTrafficRoutingConfigOutput)
}

func (a DeploymentConfigTrafficRoutingConfigArgs) ToDeploymentConfigTrafficRoutingConfigOutputWithContext(ctx context.Context) DeploymentConfigTrafficRoutingConfigOutput {
	return pulumi.ToOutputWithContext(ctx, a).(DeploymentConfigTrafficRoutingConfigOutput)
}

type DeploymentConfigTrafficRoutingConfigOutput struct { *pulumi.OutputState }

func (o DeploymentConfigTrafficRoutingConfigOutput) TimeBasedCanary() DeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutput {
	return o.Apply(func(v DeploymentConfigTrafficRoutingConfig) DeploymentConfigTrafficRoutingConfigTimeBasedCanary {
		if v.TimeBasedCanary == nil { return *new(DeploymentConfigTrafficRoutingConfigTimeBasedCanary) } else { return *v.TimeBasedCanary }
	}).(DeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutput)
}

func (o DeploymentConfigTrafficRoutingConfigOutput) TimeBasedLinear() DeploymentConfigTrafficRoutingConfigTimeBasedLinearOutput {
	return o.Apply(func(v DeploymentConfigTrafficRoutingConfig) DeploymentConfigTrafficRoutingConfigTimeBasedLinear {
		if v.TimeBasedLinear == nil { return *new(DeploymentConfigTrafficRoutingConfigTimeBasedLinear) } else { return *v.TimeBasedLinear }
	}).(DeploymentConfigTrafficRoutingConfigTimeBasedLinearOutput)
}

func (o DeploymentConfigTrafficRoutingConfigOutput) Type() pulumi.StringOutput {
	return o.Apply(func(v DeploymentConfigTrafficRoutingConfig) string {
		if v.Type == nil { return *new(string) } else { return *v.Type }
	}).(pulumi.StringOutput)
}

func (DeploymentConfigTrafficRoutingConfigOutput) ElementType() reflect.Type {
	return deploymentConfigTrafficRoutingConfigType
}

func (o DeploymentConfigTrafficRoutingConfigOutput) ToDeploymentConfigTrafficRoutingConfigOutput() DeploymentConfigTrafficRoutingConfigOutput {
	return o
}

func (o DeploymentConfigTrafficRoutingConfigOutput) ToDeploymentConfigTrafficRoutingConfigOutputWithContext(ctx context.Context) DeploymentConfigTrafficRoutingConfigOutput {
	return o
}

func init() { pulumi.RegisterOutputType(DeploymentConfigTrafficRoutingConfigOutput{}) }

type DeploymentConfigTrafficRoutingConfigTimeBasedCanary struct {
	Interval *int `pulumi:"interval"`
	Percentage *int `pulumi:"percentage"`
}
var deploymentConfigTrafficRoutingConfigTimeBasedCanaryType = reflect.TypeOf((*DeploymentConfigTrafficRoutingConfigTimeBasedCanary)(nil)).Elem()

type DeploymentConfigTrafficRoutingConfigTimeBasedCanaryInput interface {
	pulumi.Input

	ToDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutput() DeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutput
	ToDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputWithContext(ctx context.Context) DeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutput
}

type DeploymentConfigTrafficRoutingConfigTimeBasedCanaryArgs struct {
	Interval pulumi.IntInput `pulumi:"interval"`
	Percentage pulumi.IntInput `pulumi:"percentage"`
}

func (DeploymentConfigTrafficRoutingConfigTimeBasedCanaryArgs) ElementType() reflect.Type {
	return deploymentConfigTrafficRoutingConfigTimeBasedCanaryType
}

func (a DeploymentConfigTrafficRoutingConfigTimeBasedCanaryArgs) ToDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutput() DeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutput {
	return pulumi.ToOutput(a).(DeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutput)
}

func (a DeploymentConfigTrafficRoutingConfigTimeBasedCanaryArgs) ToDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputWithContext(ctx context.Context) DeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutput {
	return pulumi.ToOutputWithContext(ctx, a).(DeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutput)
}

type DeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutput struct { *pulumi.OutputState }

func (o DeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutput) Interval() pulumi.IntOutput {
	return o.Apply(func(v DeploymentConfigTrafficRoutingConfigTimeBasedCanary) int {
		if v.Interval == nil { return *new(int) } else { return *v.Interval }
	}).(pulumi.IntOutput)
}

func (o DeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutput) Percentage() pulumi.IntOutput {
	return o.Apply(func(v DeploymentConfigTrafficRoutingConfigTimeBasedCanary) int {
		if v.Percentage == nil { return *new(int) } else { return *v.Percentage }
	}).(pulumi.IntOutput)
}

func (DeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutput) ElementType() reflect.Type {
	return deploymentConfigTrafficRoutingConfigTimeBasedCanaryType
}

func (o DeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutput) ToDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutput() DeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutput {
	return o
}

func (o DeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutput) ToDeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutputWithContext(ctx context.Context) DeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutput {
	return o
}

func init() { pulumi.RegisterOutputType(DeploymentConfigTrafficRoutingConfigTimeBasedCanaryOutput{}) }

type DeploymentConfigTrafficRoutingConfigTimeBasedLinear struct {
	Interval *int `pulumi:"interval"`
	Percentage *int `pulumi:"percentage"`
}
var deploymentConfigTrafficRoutingConfigTimeBasedLinearType = reflect.TypeOf((*DeploymentConfigTrafficRoutingConfigTimeBasedLinear)(nil)).Elem()

type DeploymentConfigTrafficRoutingConfigTimeBasedLinearInput interface {
	pulumi.Input

	ToDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutput() DeploymentConfigTrafficRoutingConfigTimeBasedLinearOutput
	ToDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputWithContext(ctx context.Context) DeploymentConfigTrafficRoutingConfigTimeBasedLinearOutput
}

type DeploymentConfigTrafficRoutingConfigTimeBasedLinearArgs struct {
	Interval pulumi.IntInput `pulumi:"interval"`
	Percentage pulumi.IntInput `pulumi:"percentage"`
}

func (DeploymentConfigTrafficRoutingConfigTimeBasedLinearArgs) ElementType() reflect.Type {
	return deploymentConfigTrafficRoutingConfigTimeBasedLinearType
}

func (a DeploymentConfigTrafficRoutingConfigTimeBasedLinearArgs) ToDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutput() DeploymentConfigTrafficRoutingConfigTimeBasedLinearOutput {
	return pulumi.ToOutput(a).(DeploymentConfigTrafficRoutingConfigTimeBasedLinearOutput)
}

func (a DeploymentConfigTrafficRoutingConfigTimeBasedLinearArgs) ToDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputWithContext(ctx context.Context) DeploymentConfigTrafficRoutingConfigTimeBasedLinearOutput {
	return pulumi.ToOutputWithContext(ctx, a).(DeploymentConfigTrafficRoutingConfigTimeBasedLinearOutput)
}

type DeploymentConfigTrafficRoutingConfigTimeBasedLinearOutput struct { *pulumi.OutputState }

func (o DeploymentConfigTrafficRoutingConfigTimeBasedLinearOutput) Interval() pulumi.IntOutput {
	return o.Apply(func(v DeploymentConfigTrafficRoutingConfigTimeBasedLinear) int {
		if v.Interval == nil { return *new(int) } else { return *v.Interval }
	}).(pulumi.IntOutput)
}

func (o DeploymentConfigTrafficRoutingConfigTimeBasedLinearOutput) Percentage() pulumi.IntOutput {
	return o.Apply(func(v DeploymentConfigTrafficRoutingConfigTimeBasedLinear) int {
		if v.Percentage == nil { return *new(int) } else { return *v.Percentage }
	}).(pulumi.IntOutput)
}

func (DeploymentConfigTrafficRoutingConfigTimeBasedLinearOutput) ElementType() reflect.Type {
	return deploymentConfigTrafficRoutingConfigTimeBasedLinearType
}

func (o DeploymentConfigTrafficRoutingConfigTimeBasedLinearOutput) ToDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutput() DeploymentConfigTrafficRoutingConfigTimeBasedLinearOutput {
	return o
}

func (o DeploymentConfigTrafficRoutingConfigTimeBasedLinearOutput) ToDeploymentConfigTrafficRoutingConfigTimeBasedLinearOutputWithContext(ctx context.Context) DeploymentConfigTrafficRoutingConfigTimeBasedLinearOutput {
	return o
}

func init() { pulumi.RegisterOutputType(DeploymentConfigTrafficRoutingConfigTimeBasedLinearOutput{}) }

