// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package globalaccelerator

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Global Accelerator endpoint group.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/globalaccelerator_endpoint_group.html.markdown.
type EndpointGroup struct {
	pulumi.CustomResourceState

	// The list of endpoint objects. Fields documented below.
	EndpointConfigurations EndpointGroupEndpointConfigurationsArrayOutput `pulumi:"endpointConfigurations"`

	EndpointGroupRegion pulumi.StringOutput `pulumi:"endpointGroupRegion"`

	// The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
	HealthCheckIntervalSeconds pulumi.IntOutput `pulumi:"healthCheckIntervalSeconds"`

	// If the protocol is HTTP/S, then this specifies the path that is the destination for health check targets. The default value is slash (/).
	HealthCheckPath pulumi.StringOutput `pulumi:"healthCheckPath"`

	// The port that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default port is the listener port that this endpoint group is associated with. If listener port is a list of ports, Global Accelerator uses the first port in the list.
	HealthCheckPort pulumi.IntOutput `pulumi:"healthCheckPort"`

	// The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
	HealthCheckProtocol pulumi.StringOutput `pulumi:"healthCheckProtocol"`

	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn pulumi.StringOutput `pulumi:"listenerArn"`

	// The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
	ThresholdCount pulumi.IntOutput `pulumi:"thresholdCount"`

	// The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.
	TrafficDialPercentage pulumi.Float64Output `pulumi:"trafficDialPercentage"`
}

// NewEndpointGroup registers a new resource with the given unique name, arguments, and options.
func NewEndpointGroup(ctx *pulumi.Context,
	name string, args *EndpointGroupArgs, opts ...pulumi.ResourceOption) (*EndpointGroup, error) {
	if args == nil || args.ListenerArn == nil {
		return nil, errors.New("missing required argument 'ListenerArn'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.EndpointConfigurations; i != nil { inputs["endpointConfigurations"] = i.ToEndpointGroupEndpointConfigurationsArrayOutput() }
		if i := args.EndpointGroupRegion; i != nil { inputs["endpointGroupRegion"] = i.ToStringOutput() }
		if i := args.HealthCheckIntervalSeconds; i != nil { inputs["healthCheckIntervalSeconds"] = i.ToIntOutput() }
		if i := args.HealthCheckPath; i != nil { inputs["healthCheckPath"] = i.ToStringOutput() }
		if i := args.HealthCheckPort; i != nil { inputs["healthCheckPort"] = i.ToIntOutput() }
		if i := args.HealthCheckProtocol; i != nil { inputs["healthCheckProtocol"] = i.ToStringOutput() }
		if i := args.ListenerArn; i != nil { inputs["listenerArn"] = i.ToStringOutput() }
		if i := args.ThresholdCount; i != nil { inputs["thresholdCount"] = i.ToIntOutput() }
		if i := args.TrafficDialPercentage; i != nil { inputs["trafficDialPercentage"] = i.ToFloat64Output() }
	}
	var resource EndpointGroup
	err := ctx.RegisterResource("aws:globalaccelerator/endpointGroup:EndpointGroup", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointGroup gets an existing EndpointGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointGroupState, opts ...pulumi.ResourceOption) (*EndpointGroup, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.EndpointConfigurations; i != nil { inputs["endpointConfigurations"] = i.ToEndpointGroupEndpointConfigurationsArrayOutput() }
		if i := state.EndpointGroupRegion; i != nil { inputs["endpointGroupRegion"] = i.ToStringOutput() }
		if i := state.HealthCheckIntervalSeconds; i != nil { inputs["healthCheckIntervalSeconds"] = i.ToIntOutput() }
		if i := state.HealthCheckPath; i != nil { inputs["healthCheckPath"] = i.ToStringOutput() }
		if i := state.HealthCheckPort; i != nil { inputs["healthCheckPort"] = i.ToIntOutput() }
		if i := state.HealthCheckProtocol; i != nil { inputs["healthCheckProtocol"] = i.ToStringOutput() }
		if i := state.ListenerArn; i != nil { inputs["listenerArn"] = i.ToStringOutput() }
		if i := state.ThresholdCount; i != nil { inputs["thresholdCount"] = i.ToIntOutput() }
		if i := state.TrafficDialPercentage; i != nil { inputs["trafficDialPercentage"] = i.ToFloat64Output() }
	}
	var resource EndpointGroup
	err := ctx.ReadResource("aws:globalaccelerator/endpointGroup:EndpointGroup", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointGroup resources.
type EndpointGroupState struct {
	// The list of endpoint objects. Fields documented below.
	EndpointConfigurations EndpointGroupEndpointConfigurationsArrayInput `pulumi:"endpointConfigurations"`
	EndpointGroupRegion pulumi.StringInput `pulumi:"endpointGroupRegion"`
	// The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
	HealthCheckIntervalSeconds pulumi.IntInput `pulumi:"healthCheckIntervalSeconds"`
	// If the protocol is HTTP/S, then this specifies the path that is the destination for health check targets. The default value is slash (/).
	HealthCheckPath pulumi.StringInput `pulumi:"healthCheckPath"`
	// The port that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default port is the listener port that this endpoint group is associated with. If listener port is a list of ports, Global Accelerator uses the first port in the list.
	HealthCheckPort pulumi.IntInput `pulumi:"healthCheckPort"`
	// The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
	HealthCheckProtocol pulumi.StringInput `pulumi:"healthCheckProtocol"`
	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn pulumi.StringInput `pulumi:"listenerArn"`
	// The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
	ThresholdCount pulumi.IntInput `pulumi:"thresholdCount"`
	// The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.
	TrafficDialPercentage pulumi.Float64Input `pulumi:"trafficDialPercentage"`
}

// The set of arguments for constructing a EndpointGroup resource.
type EndpointGroupArgs struct {
	// The list of endpoint objects. Fields documented below.
	EndpointConfigurations EndpointGroupEndpointConfigurationsArrayInput `pulumi:"endpointConfigurations"`
	EndpointGroupRegion pulumi.StringInput `pulumi:"endpointGroupRegion"`
	// The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
	HealthCheckIntervalSeconds pulumi.IntInput `pulumi:"healthCheckIntervalSeconds"`
	// If the protocol is HTTP/S, then this specifies the path that is the destination for health check targets. The default value is slash (/).
	HealthCheckPath pulumi.StringInput `pulumi:"healthCheckPath"`
	// The port that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default port is the listener port that this endpoint group is associated with. If listener port is a list of ports, Global Accelerator uses the first port in the list.
	HealthCheckPort pulumi.IntInput `pulumi:"healthCheckPort"`
	// The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
	HealthCheckProtocol pulumi.StringInput `pulumi:"healthCheckProtocol"`
	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn pulumi.StringInput `pulumi:"listenerArn"`
	// The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
	ThresholdCount pulumi.IntInput `pulumi:"thresholdCount"`
	// The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.
	TrafficDialPercentage pulumi.Float64Input `pulumi:"trafficDialPercentage"`
}
type EndpointGroupEndpointConfigurations struct {
	// An ID for the endpoint. If the endpoint is a Network Load Balancer or Application Load Balancer, this is the Amazon Resource Name (ARN) of the resource. If the endpoint is an Elastic IP address, this is the Elastic IP address allocation ID.
	EndpointId *string `pulumi:"endpointId"`
	// The weight associated with the endpoint. When you add weights to endpoints, you configure AWS Global Accelerator to route traffic based on proportions that you specify. 
	Weight *int `pulumi:"weight"`
}
var endpointGroupEndpointConfigurationsType = reflect.TypeOf((*EndpointGroupEndpointConfigurations)(nil)).Elem()

type EndpointGroupEndpointConfigurationsInput interface {
	pulumi.Input

	ToEndpointGroupEndpointConfigurationsOutput() EndpointGroupEndpointConfigurationsOutput
	ToEndpointGroupEndpointConfigurationsOutputWithContext(ctx context.Context) EndpointGroupEndpointConfigurationsOutput
}

type EndpointGroupEndpointConfigurationsArgs struct {
	// An ID for the endpoint. If the endpoint is a Network Load Balancer or Application Load Balancer, this is the Amazon Resource Name (ARN) of the resource. If the endpoint is an Elastic IP address, this is the Elastic IP address allocation ID.
	EndpointId pulumi.StringInput `pulumi:"endpointId"`
	// The weight associated with the endpoint. When you add weights to endpoints, you configure AWS Global Accelerator to route traffic based on proportions that you specify. 
	Weight pulumi.IntInput `pulumi:"weight"`
}

func (EndpointGroupEndpointConfigurationsArgs) ElementType() reflect.Type {
	return endpointGroupEndpointConfigurationsType
}

func (a EndpointGroupEndpointConfigurationsArgs) ToEndpointGroupEndpointConfigurationsOutput() EndpointGroupEndpointConfigurationsOutput {
	return pulumi.ToOutput(a).(EndpointGroupEndpointConfigurationsOutput)
}

func (a EndpointGroupEndpointConfigurationsArgs) ToEndpointGroupEndpointConfigurationsOutputWithContext(ctx context.Context) EndpointGroupEndpointConfigurationsOutput {
	return pulumi.ToOutputWithContext(ctx, a).(EndpointGroupEndpointConfigurationsOutput)
}

type EndpointGroupEndpointConfigurationsOutput struct { *pulumi.OutputState }

// An ID for the endpoint. If the endpoint is a Network Load Balancer or Application Load Balancer, this is the Amazon Resource Name (ARN) of the resource. If the endpoint is an Elastic IP address, this is the Elastic IP address allocation ID.
func (o EndpointGroupEndpointConfigurationsOutput) EndpointId() pulumi.StringOutput {
	return o.Apply(func(v EndpointGroupEndpointConfigurations) string {
		if v.EndpointId == nil { return *new(string) } else { return *v.EndpointId }
	}).(pulumi.StringOutput)
}

// The weight associated with the endpoint. When you add weights to endpoints, you configure AWS Global Accelerator to route traffic based on proportions that you specify. 
func (o EndpointGroupEndpointConfigurationsOutput) Weight() pulumi.IntOutput {
	return o.Apply(func(v EndpointGroupEndpointConfigurations) int {
		if v.Weight == nil { return *new(int) } else { return *v.Weight }
	}).(pulumi.IntOutput)
}

func (EndpointGroupEndpointConfigurationsOutput) ElementType() reflect.Type {
	return endpointGroupEndpointConfigurationsType
}

func (o EndpointGroupEndpointConfigurationsOutput) ToEndpointGroupEndpointConfigurationsOutput() EndpointGroupEndpointConfigurationsOutput {
	return o
}

func (o EndpointGroupEndpointConfigurationsOutput) ToEndpointGroupEndpointConfigurationsOutputWithContext(ctx context.Context) EndpointGroupEndpointConfigurationsOutput {
	return o
}

func init() { pulumi.RegisterOutputType(EndpointGroupEndpointConfigurationsOutput{}) }

var endpointGroupEndpointConfigurationsArrayType = reflect.TypeOf((*[]EndpointGroupEndpointConfigurations)(nil)).Elem()

type EndpointGroupEndpointConfigurationsArrayInput interface {
	pulumi.Input

	ToEndpointGroupEndpointConfigurationsArrayOutput() EndpointGroupEndpointConfigurationsArrayOutput
	ToEndpointGroupEndpointConfigurationsArrayOutputWithContext(ctx context.Context) EndpointGroupEndpointConfigurationsArrayOutput
}

type EndpointGroupEndpointConfigurationsArrayArgs []EndpointGroupEndpointConfigurationsInput

func (EndpointGroupEndpointConfigurationsArrayArgs) ElementType() reflect.Type {
	return endpointGroupEndpointConfigurationsArrayType
}

func (a EndpointGroupEndpointConfigurationsArrayArgs) ToEndpointGroupEndpointConfigurationsArrayOutput() EndpointGroupEndpointConfigurationsArrayOutput {
	return pulumi.ToOutput(a).(EndpointGroupEndpointConfigurationsArrayOutput)
}

func (a EndpointGroupEndpointConfigurationsArrayArgs) ToEndpointGroupEndpointConfigurationsArrayOutputWithContext(ctx context.Context) EndpointGroupEndpointConfigurationsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(EndpointGroupEndpointConfigurationsArrayOutput)
}

type EndpointGroupEndpointConfigurationsArrayOutput struct { *pulumi.OutputState }

func (o EndpointGroupEndpointConfigurationsArrayOutput) Index(i pulumi.IntInput) EndpointGroupEndpointConfigurationsOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) EndpointGroupEndpointConfigurations {
		return vs[0].([]EndpointGroupEndpointConfigurations)[vs[1].(int)]
	}).(EndpointGroupEndpointConfigurationsOutput)
}

func (EndpointGroupEndpointConfigurationsArrayOutput) ElementType() reflect.Type {
	return endpointGroupEndpointConfigurationsArrayType
}

func (o EndpointGroupEndpointConfigurationsArrayOutput) ToEndpointGroupEndpointConfigurationsArrayOutput() EndpointGroupEndpointConfigurationsArrayOutput {
	return o
}

func (o EndpointGroupEndpointConfigurationsArrayOutput) ToEndpointGroupEndpointConfigurationsArrayOutputWithContext(ctx context.Context) EndpointGroupEndpointConfigurationsArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(EndpointGroupEndpointConfigurationsArrayOutput{}) }

