// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elb

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a load balancer SSL negotiation policy, which allows an ELB to control the ciphers and protocols that are supported during SSL negotiations between a client and a load balancer.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lb_ssl_negotiation_policy.html.markdown.
type SslNegotiationPolicy struct {
	pulumi.CustomResourceState

	// An SSL Negotiation policy attribute. Each has two properties:
	Attributes SslNegotiationPolicyAttributesArrayOutput `pulumi:"attributes"`

	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort pulumi.IntOutput `pulumi:"lbPort"`

	// The load balancer to which the policy
	// should be attached.
	LoadBalancer pulumi.StringOutput `pulumi:"loadBalancer"`

	// The name of the attribute
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewSslNegotiationPolicy registers a new resource with the given unique name, arguments, and options.
func NewSslNegotiationPolicy(ctx *pulumi.Context,
	name string, args *SslNegotiationPolicyArgs, opts ...pulumi.ResourceOption) (*SslNegotiationPolicy, error) {
	if args == nil || args.LbPort == nil {
		return nil, errors.New("missing required argument 'LbPort'")
	}
	if args == nil || args.LoadBalancer == nil {
		return nil, errors.New("missing required argument 'LoadBalancer'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Attributes; i != nil { inputs["attributes"] = i.ToSslNegotiationPolicyAttributesArrayOutput() }
		if i := args.LbPort; i != nil { inputs["lbPort"] = i.ToIntOutput() }
		if i := args.LoadBalancer; i != nil { inputs["loadBalancer"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
	}
	var resource SslNegotiationPolicy
	err := ctx.RegisterResource("aws:elb/sslNegotiationPolicy:SslNegotiationPolicy", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSslNegotiationPolicy gets an existing SslNegotiationPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSslNegotiationPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SslNegotiationPolicyState, opts ...pulumi.ResourceOption) (*SslNegotiationPolicy, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Attributes; i != nil { inputs["attributes"] = i.ToSslNegotiationPolicyAttributesArrayOutput() }
		if i := state.LbPort; i != nil { inputs["lbPort"] = i.ToIntOutput() }
		if i := state.LoadBalancer; i != nil { inputs["loadBalancer"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
	}
	var resource SslNegotiationPolicy
	err := ctx.ReadResource("aws:elb/sslNegotiationPolicy:SslNegotiationPolicy", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SslNegotiationPolicy resources.
type SslNegotiationPolicyState struct {
	// An SSL Negotiation policy attribute. Each has two properties:
	Attributes SslNegotiationPolicyAttributesArrayInput `pulumi:"attributes"`
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort pulumi.IntInput `pulumi:"lbPort"`
	// The load balancer to which the policy
	// should be attached.
	LoadBalancer pulumi.StringInput `pulumi:"loadBalancer"`
	// The name of the attribute
	Name pulumi.StringInput `pulumi:"name"`
}

// The set of arguments for constructing a SslNegotiationPolicy resource.
type SslNegotiationPolicyArgs struct {
	// An SSL Negotiation policy attribute. Each has two properties:
	Attributes SslNegotiationPolicyAttributesArrayInput `pulumi:"attributes"`
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort pulumi.IntInput `pulumi:"lbPort"`
	// The load balancer to which the policy
	// should be attached.
	LoadBalancer pulumi.StringInput `pulumi:"loadBalancer"`
	// The name of the attribute
	Name pulumi.StringInput `pulumi:"name"`
}
type SslNegotiationPolicyAttributes struct {
	// The name of the attribute
	Name string `pulumi:"name"`
	// The value of the attribute
	Value string `pulumi:"value"`
}
var sslNegotiationPolicyAttributesType = reflect.TypeOf((*SslNegotiationPolicyAttributes)(nil)).Elem()

type SslNegotiationPolicyAttributesInput interface {
	pulumi.Input

	ToSslNegotiationPolicyAttributesOutput() SslNegotiationPolicyAttributesOutput
	ToSslNegotiationPolicyAttributesOutputWithContext(ctx context.Context) SslNegotiationPolicyAttributesOutput
}

type SslNegotiationPolicyAttributesArgs struct {
	// The name of the attribute
	Name pulumi.StringInput `pulumi:"name"`
	// The value of the attribute
	Value pulumi.StringInput `pulumi:"value"`
}

func (SslNegotiationPolicyAttributesArgs) ElementType() reflect.Type {
	return sslNegotiationPolicyAttributesType
}

func (a SslNegotiationPolicyAttributesArgs) ToSslNegotiationPolicyAttributesOutput() SslNegotiationPolicyAttributesOutput {
	return pulumi.ToOutput(a).(SslNegotiationPolicyAttributesOutput)
}

func (a SslNegotiationPolicyAttributesArgs) ToSslNegotiationPolicyAttributesOutputWithContext(ctx context.Context) SslNegotiationPolicyAttributesOutput {
	return pulumi.ToOutputWithContext(ctx, a).(SslNegotiationPolicyAttributesOutput)
}

type SslNegotiationPolicyAttributesOutput struct { *pulumi.OutputState }

// The name of the attribute
func (o SslNegotiationPolicyAttributesOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v SslNegotiationPolicyAttributes) string {
		return v.Name
	}).(pulumi.StringOutput)
}

// The value of the attribute
func (o SslNegotiationPolicyAttributesOutput) Value() pulumi.StringOutput {
	return o.Apply(func(v SslNegotiationPolicyAttributes) string {
		return v.Value
	}).(pulumi.StringOutput)
}

func (SslNegotiationPolicyAttributesOutput) ElementType() reflect.Type {
	return sslNegotiationPolicyAttributesType
}

func (o SslNegotiationPolicyAttributesOutput) ToSslNegotiationPolicyAttributesOutput() SslNegotiationPolicyAttributesOutput {
	return o
}

func (o SslNegotiationPolicyAttributesOutput) ToSslNegotiationPolicyAttributesOutputWithContext(ctx context.Context) SslNegotiationPolicyAttributesOutput {
	return o
}

func init() { pulumi.RegisterOutputType(SslNegotiationPolicyAttributesOutput{}) }

var sslNegotiationPolicyAttributesArrayType = reflect.TypeOf((*[]SslNegotiationPolicyAttributes)(nil)).Elem()

type SslNegotiationPolicyAttributesArrayInput interface {
	pulumi.Input

	ToSslNegotiationPolicyAttributesArrayOutput() SslNegotiationPolicyAttributesArrayOutput
	ToSslNegotiationPolicyAttributesArrayOutputWithContext(ctx context.Context) SslNegotiationPolicyAttributesArrayOutput
}

type SslNegotiationPolicyAttributesArrayArgs []SslNegotiationPolicyAttributesInput

func (SslNegotiationPolicyAttributesArrayArgs) ElementType() reflect.Type {
	return sslNegotiationPolicyAttributesArrayType
}

func (a SslNegotiationPolicyAttributesArrayArgs) ToSslNegotiationPolicyAttributesArrayOutput() SslNegotiationPolicyAttributesArrayOutput {
	return pulumi.ToOutput(a).(SslNegotiationPolicyAttributesArrayOutput)
}

func (a SslNegotiationPolicyAttributesArrayArgs) ToSslNegotiationPolicyAttributesArrayOutputWithContext(ctx context.Context) SslNegotiationPolicyAttributesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(SslNegotiationPolicyAttributesArrayOutput)
}

type SslNegotiationPolicyAttributesArrayOutput struct { *pulumi.OutputState }

func (o SslNegotiationPolicyAttributesArrayOutput) Index(i pulumi.IntInput) SslNegotiationPolicyAttributesOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) SslNegotiationPolicyAttributes {
		return vs[0].([]SslNegotiationPolicyAttributes)[vs[1].(int)]
	}).(SslNegotiationPolicyAttributesOutput)
}

func (SslNegotiationPolicyAttributesArrayOutput) ElementType() reflect.Type {
	return sslNegotiationPolicyAttributesArrayType
}

func (o SslNegotiationPolicyAttributesArrayOutput) ToSslNegotiationPolicyAttributesArrayOutput() SslNegotiationPolicyAttributesArrayOutput {
	return o
}

func (o SslNegotiationPolicyAttributesArrayOutput) ToSslNegotiationPolicyAttributesArrayOutputWithContext(ctx context.Context) SslNegotiationPolicyAttributesArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(SslNegotiationPolicyAttributesArrayOutput{}) }

