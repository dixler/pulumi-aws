// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package xray

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates and manages an AWS XRay Sampling Rule.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/xray_sampling_rule.html.markdown.
type SamplingRule struct {
	pulumi.CustomResourceState

	// The ARN of the sampling rule.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// Matches attributes derived from the request.
	Attributes pulumi.StringMapOutput `pulumi:"attributes"`

	// The percentage of matching requests to instrument, after the reservoir is exhausted.
	FixedRate pulumi.Float64Output `pulumi:"fixedRate"`

	// Matches the hostname from a request URL.
	Host pulumi.StringOutput `pulumi:"host"`

	// Matches the HTTP method of a request.
	HttpMethod pulumi.StringOutput `pulumi:"httpMethod"`

	// The priority of the sampling rule.
	Priority pulumi.IntOutput `pulumi:"priority"`

	// A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.
	ReservoirSize pulumi.IntOutput `pulumi:"reservoirSize"`

	// Matches the ARN of the AWS resource on which the service runs.
	ResourceArn pulumi.StringOutput `pulumi:"resourceArn"`

	// The name of the sampling rule.
	RuleName pulumi.StringOutput `pulumi:"ruleName"`

	// Matches the `name` that the service uses to identify itself in segments.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`

	// Matches the `origin` that the service uses to identify its type in segments.
	ServiceType pulumi.StringOutput `pulumi:"serviceType"`

	// Matches the path from a request URL.
	UrlPath pulumi.StringOutput `pulumi:"urlPath"`

	// The version of the sampling rule format (`1` )
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewSamplingRule registers a new resource with the given unique name, arguments, and options.
func NewSamplingRule(ctx *pulumi.Context,
	name string, args *SamplingRuleArgs, opts ...pulumi.ResourceOption) (*SamplingRule, error) {
	if args == nil || args.FixedRate == nil {
		return nil, errors.New("missing required argument 'FixedRate'")
	}
	if args == nil || args.Host == nil {
		return nil, errors.New("missing required argument 'Host'")
	}
	if args == nil || args.HttpMethod == nil {
		return nil, errors.New("missing required argument 'HttpMethod'")
	}
	if args == nil || args.Priority == nil {
		return nil, errors.New("missing required argument 'Priority'")
	}
	if args == nil || args.ReservoirSize == nil {
		return nil, errors.New("missing required argument 'ReservoirSize'")
	}
	if args == nil || args.ResourceArn == nil {
		return nil, errors.New("missing required argument 'ResourceArn'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil || args.ServiceType == nil {
		return nil, errors.New("missing required argument 'ServiceType'")
	}
	if args == nil || args.UrlPath == nil {
		return nil, errors.New("missing required argument 'UrlPath'")
	}
	if args == nil || args.Version == nil {
		return nil, errors.New("missing required argument 'Version'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Attributes; i != nil { inputs["attributes"] = i.ToStringMapOutput() }
		if i := args.FixedRate; i != nil { inputs["fixedRate"] = i.ToFloat64Output() }
		if i := args.Host; i != nil { inputs["host"] = i.ToStringOutput() }
		if i := args.HttpMethod; i != nil { inputs["httpMethod"] = i.ToStringOutput() }
		if i := args.Priority; i != nil { inputs["priority"] = i.ToIntOutput() }
		if i := args.ReservoirSize; i != nil { inputs["reservoirSize"] = i.ToIntOutput() }
		if i := args.ResourceArn; i != nil { inputs["resourceArn"] = i.ToStringOutput() }
		if i := args.RuleName; i != nil { inputs["ruleName"] = i.ToStringOutput() }
		if i := args.ServiceName; i != nil { inputs["serviceName"] = i.ToStringOutput() }
		if i := args.ServiceType; i != nil { inputs["serviceType"] = i.ToStringOutput() }
		if i := args.UrlPath; i != nil { inputs["urlPath"] = i.ToStringOutput() }
		if i := args.Version; i != nil { inputs["version"] = i.ToIntOutput() }
	}
	var resource SamplingRule
	err := ctx.RegisterResource("aws:xray/samplingRule:SamplingRule", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSamplingRule gets an existing SamplingRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSamplingRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SamplingRuleState, opts ...pulumi.ResourceOption) (*SamplingRule, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Arn; i != nil { inputs["arn"] = i.ToStringOutput() }
		if i := state.Attributes; i != nil { inputs["attributes"] = i.ToStringMapOutput() }
		if i := state.FixedRate; i != nil { inputs["fixedRate"] = i.ToFloat64Output() }
		if i := state.Host; i != nil { inputs["host"] = i.ToStringOutput() }
		if i := state.HttpMethod; i != nil { inputs["httpMethod"] = i.ToStringOutput() }
		if i := state.Priority; i != nil { inputs["priority"] = i.ToIntOutput() }
		if i := state.ReservoirSize; i != nil { inputs["reservoirSize"] = i.ToIntOutput() }
		if i := state.ResourceArn; i != nil { inputs["resourceArn"] = i.ToStringOutput() }
		if i := state.RuleName; i != nil { inputs["ruleName"] = i.ToStringOutput() }
		if i := state.ServiceName; i != nil { inputs["serviceName"] = i.ToStringOutput() }
		if i := state.ServiceType; i != nil { inputs["serviceType"] = i.ToStringOutput() }
		if i := state.UrlPath; i != nil { inputs["urlPath"] = i.ToStringOutput() }
		if i := state.Version; i != nil { inputs["version"] = i.ToIntOutput() }
	}
	var resource SamplingRule
	err := ctx.ReadResource("aws:xray/samplingRule:SamplingRule", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SamplingRule resources.
type SamplingRuleState struct {
	// The ARN of the sampling rule.
	Arn pulumi.StringInput `pulumi:"arn"`
	// Matches attributes derived from the request.
	Attributes pulumi.StringMapInput `pulumi:"attributes"`
	// The percentage of matching requests to instrument, after the reservoir is exhausted.
	FixedRate pulumi.Float64Input `pulumi:"fixedRate"`
	// Matches the hostname from a request URL.
	Host pulumi.StringInput `pulumi:"host"`
	// Matches the HTTP method of a request.
	HttpMethod pulumi.StringInput `pulumi:"httpMethod"`
	// The priority of the sampling rule.
	Priority pulumi.IntInput `pulumi:"priority"`
	// A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.
	ReservoirSize pulumi.IntInput `pulumi:"reservoirSize"`
	// Matches the ARN of the AWS resource on which the service runs.
	ResourceArn pulumi.StringInput `pulumi:"resourceArn"`
	// The name of the sampling rule.
	RuleName pulumi.StringInput `pulumi:"ruleName"`
	// Matches the `name` that the service uses to identify itself in segments.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Matches the `origin` that the service uses to identify its type in segments.
	ServiceType pulumi.StringInput `pulumi:"serviceType"`
	// Matches the path from a request URL.
	UrlPath pulumi.StringInput `pulumi:"urlPath"`
	// The version of the sampling rule format (`1` )
	Version pulumi.IntInput `pulumi:"version"`
}

// The set of arguments for constructing a SamplingRule resource.
type SamplingRuleArgs struct {
	// Matches attributes derived from the request.
	Attributes pulumi.StringMapInput `pulumi:"attributes"`
	// The percentage of matching requests to instrument, after the reservoir is exhausted.
	FixedRate pulumi.Float64Input `pulumi:"fixedRate"`
	// Matches the hostname from a request URL.
	Host pulumi.StringInput `pulumi:"host"`
	// Matches the HTTP method of a request.
	HttpMethod pulumi.StringInput `pulumi:"httpMethod"`
	// The priority of the sampling rule.
	Priority pulumi.IntInput `pulumi:"priority"`
	// A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.
	ReservoirSize pulumi.IntInput `pulumi:"reservoirSize"`
	// Matches the ARN of the AWS resource on which the service runs.
	ResourceArn pulumi.StringInput `pulumi:"resourceArn"`
	// The name of the sampling rule.
	RuleName pulumi.StringInput `pulumi:"ruleName"`
	// Matches the `name` that the service uses to identify itself in segments.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Matches the `origin` that the service uses to identify its type in segments.
	ServiceType pulumi.StringInput `pulumi:"serviceType"`
	// Matches the path from a request URL.
	UrlPath pulumi.StringInput `pulumi:"urlPath"`
	// The version of the sampling rule format (`1` )
	Version pulumi.IntInput `pulumi:"version"`
}
