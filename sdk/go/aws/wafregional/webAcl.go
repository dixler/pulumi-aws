// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a WAF Regional Web ACL Resource for use with Application Load Balancer.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/wafregional_web_acl.html.markdown.
type WebAcl struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the WAF Regional WebACL.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
	DefaultAction WebAclDefaultActionOutput `pulumi:"defaultAction"`

	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration WebAclLoggingConfigurationOutput `pulumi:"loggingConfiguration"`

	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName pulumi.StringOutput `pulumi:"metricName"`

	// The name or description of the web ACL.
	Name pulumi.StringOutput `pulumi:"name"`

	// Set of configuration blocks containing rules for the web ACL. Detailed below.
	Rules WebAclRulesArrayOutput `pulumi:"rules"`

	// Key-value mapping of resource tags
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewWebAcl registers a new resource with the given unique name, arguments, and options.
func NewWebAcl(ctx *pulumi.Context,
	name string, args *WebAclArgs, opts ...pulumi.ResourceOption) (*WebAcl, error) {
	if args == nil || args.DefaultAction == nil {
		return nil, errors.New("missing required argument 'DefaultAction'")
	}
	if args == nil || args.MetricName == nil {
		return nil, errors.New("missing required argument 'MetricName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.DefaultAction; i != nil { inputs["defaultAction"] = i.ToWebAclDefaultActionOutput() }
		if i := args.LoggingConfiguration; i != nil { inputs["loggingConfiguration"] = i.ToWebAclLoggingConfigurationOutput() }
		if i := args.MetricName; i != nil { inputs["metricName"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.Rules; i != nil { inputs["rules"] = i.ToWebAclRulesArrayOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource WebAcl
	err := ctx.RegisterResource("aws:wafregional/webAcl:WebAcl", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAcl gets an existing WebAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAclState, opts ...pulumi.ResourceOption) (*WebAcl, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Arn; i != nil { inputs["arn"] = i.ToStringOutput() }
		if i := state.DefaultAction; i != nil { inputs["defaultAction"] = i.ToWebAclDefaultActionOutput() }
		if i := state.LoggingConfiguration; i != nil { inputs["loggingConfiguration"] = i.ToWebAclLoggingConfigurationOutput() }
		if i := state.MetricName; i != nil { inputs["metricName"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.Rules; i != nil { inputs["rules"] = i.ToWebAclRulesArrayOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource WebAcl
	err := ctx.ReadResource("aws:wafregional/webAcl:WebAcl", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAcl resources.
type WebAclState struct {
	// Amazon Resource Name (ARN) of the WAF Regional WebACL.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
	DefaultAction WebAclDefaultActionInput `pulumi:"defaultAction"`
	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration WebAclLoggingConfigurationInput `pulumi:"loggingConfiguration"`
	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName pulumi.StringInput `pulumi:"metricName"`
	// The name or description of the web ACL.
	Name pulumi.StringInput `pulumi:"name"`
	// Set of configuration blocks containing rules for the web ACL. Detailed below.
	Rules WebAclRulesArrayInput `pulumi:"rules"`
	// Key-value mapping of resource tags
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a WebAcl resource.
type WebAclArgs struct {
	// The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
	DefaultAction WebAclDefaultActionInput `pulumi:"defaultAction"`
	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration WebAclLoggingConfigurationInput `pulumi:"loggingConfiguration"`
	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName pulumi.StringInput `pulumi:"metricName"`
	// The name or description of the web ACL.
	Name pulumi.StringInput `pulumi:"name"`
	// Set of configuration blocks containing rules for the web ACL. Detailed below.
	Rules WebAclRulesArrayInput `pulumi:"rules"`
	// Key-value mapping of resource tags
	Tags pulumi.MapInput `pulumi:"tags"`
}
type WebAclDefaultAction struct {
	// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`
	Type string `pulumi:"type"`
}
var webAclDefaultActionType = reflect.TypeOf((*WebAclDefaultAction)(nil)).Elem()

type WebAclDefaultActionInput interface {
	pulumi.Input

	ToWebAclDefaultActionOutput() WebAclDefaultActionOutput
	ToWebAclDefaultActionOutputWithContext(ctx context.Context) WebAclDefaultActionOutput
}

type WebAclDefaultActionArgs struct {
	// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`
	Type pulumi.StringInput `pulumi:"type"`
}

func (WebAclDefaultActionArgs) ElementType() reflect.Type {
	return webAclDefaultActionType
}

func (a WebAclDefaultActionArgs) ToWebAclDefaultActionOutput() WebAclDefaultActionOutput {
	return pulumi.ToOutput(a).(WebAclDefaultActionOutput)
}

func (a WebAclDefaultActionArgs) ToWebAclDefaultActionOutputWithContext(ctx context.Context) WebAclDefaultActionOutput {
	return pulumi.ToOutputWithContext(ctx, a).(WebAclDefaultActionOutput)
}

type WebAclDefaultActionOutput struct { *pulumi.OutputState }

// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`
func (o WebAclDefaultActionOutput) Type() pulumi.StringOutput {
	return o.Apply(func(v WebAclDefaultAction) string {
		return v.Type
	}).(pulumi.StringOutput)
}

func (WebAclDefaultActionOutput) ElementType() reflect.Type {
	return webAclDefaultActionType
}

func (o WebAclDefaultActionOutput) ToWebAclDefaultActionOutput() WebAclDefaultActionOutput {
	return o
}

func (o WebAclDefaultActionOutput) ToWebAclDefaultActionOutputWithContext(ctx context.Context) WebAclDefaultActionOutput {
	return o
}

func init() { pulumi.RegisterOutputType(WebAclDefaultActionOutput{}) }

type WebAclLoggingConfiguration struct {
	// Amazon Resource Name (ARN) of Kinesis Firehose Delivery Stream
	LogDestination string `pulumi:"logDestination"`
	// Configuration block containing parts of the request that you want redacted from the logs. Detailed below.
	RedactedFields *WebAclLoggingConfigurationRedactedFields `pulumi:"redactedFields"`
}
var webAclLoggingConfigurationType = reflect.TypeOf((*WebAclLoggingConfiguration)(nil)).Elem()

type WebAclLoggingConfigurationInput interface {
	pulumi.Input

	ToWebAclLoggingConfigurationOutput() WebAclLoggingConfigurationOutput
	ToWebAclLoggingConfigurationOutputWithContext(ctx context.Context) WebAclLoggingConfigurationOutput
}

type WebAclLoggingConfigurationArgs struct {
	// Amazon Resource Name (ARN) of Kinesis Firehose Delivery Stream
	LogDestination pulumi.StringInput `pulumi:"logDestination"`
	// Configuration block containing parts of the request that you want redacted from the logs. Detailed below.
	RedactedFields WebAclLoggingConfigurationRedactedFieldsInput `pulumi:"redactedFields"`
}

func (WebAclLoggingConfigurationArgs) ElementType() reflect.Type {
	return webAclLoggingConfigurationType
}

func (a WebAclLoggingConfigurationArgs) ToWebAclLoggingConfigurationOutput() WebAclLoggingConfigurationOutput {
	return pulumi.ToOutput(a).(WebAclLoggingConfigurationOutput)
}

func (a WebAclLoggingConfigurationArgs) ToWebAclLoggingConfigurationOutputWithContext(ctx context.Context) WebAclLoggingConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, a).(WebAclLoggingConfigurationOutput)
}

type WebAclLoggingConfigurationOutput struct { *pulumi.OutputState }

// Amazon Resource Name (ARN) of Kinesis Firehose Delivery Stream
func (o WebAclLoggingConfigurationOutput) LogDestination() pulumi.StringOutput {
	return o.Apply(func(v WebAclLoggingConfiguration) string {
		return v.LogDestination
	}).(pulumi.StringOutput)
}

// Configuration block containing parts of the request that you want redacted from the logs. Detailed below.
func (o WebAclLoggingConfigurationOutput) RedactedFields() WebAclLoggingConfigurationRedactedFieldsOutput {
	return o.Apply(func(v WebAclLoggingConfiguration) WebAclLoggingConfigurationRedactedFields {
		if v.RedactedFields == nil { return *new(WebAclLoggingConfigurationRedactedFields) } else { return *v.RedactedFields }
	}).(WebAclLoggingConfigurationRedactedFieldsOutput)
}

func (WebAclLoggingConfigurationOutput) ElementType() reflect.Type {
	return webAclLoggingConfigurationType
}

func (o WebAclLoggingConfigurationOutput) ToWebAclLoggingConfigurationOutput() WebAclLoggingConfigurationOutput {
	return o
}

func (o WebAclLoggingConfigurationOutput) ToWebAclLoggingConfigurationOutputWithContext(ctx context.Context) WebAclLoggingConfigurationOutput {
	return o
}

func init() { pulumi.RegisterOutputType(WebAclLoggingConfigurationOutput{}) }

type WebAclLoggingConfigurationRedactedFields struct {
	// Set of configuration blocks for fields to redact. Detailed below.
	FieldToMatches []WebAclLoggingConfigurationRedactedFieldsFieldToMatches `pulumi:"fieldToMatches"`
}
var webAclLoggingConfigurationRedactedFieldsType = reflect.TypeOf((*WebAclLoggingConfigurationRedactedFields)(nil)).Elem()

type WebAclLoggingConfigurationRedactedFieldsInput interface {
	pulumi.Input

	ToWebAclLoggingConfigurationRedactedFieldsOutput() WebAclLoggingConfigurationRedactedFieldsOutput
	ToWebAclLoggingConfigurationRedactedFieldsOutputWithContext(ctx context.Context) WebAclLoggingConfigurationRedactedFieldsOutput
}

type WebAclLoggingConfigurationRedactedFieldsArgs struct {
	// Set of configuration blocks for fields to redact. Detailed below.
	FieldToMatches WebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayInput `pulumi:"fieldToMatches"`
}

func (WebAclLoggingConfigurationRedactedFieldsArgs) ElementType() reflect.Type {
	return webAclLoggingConfigurationRedactedFieldsType
}

func (a WebAclLoggingConfigurationRedactedFieldsArgs) ToWebAclLoggingConfigurationRedactedFieldsOutput() WebAclLoggingConfigurationRedactedFieldsOutput {
	return pulumi.ToOutput(a).(WebAclLoggingConfigurationRedactedFieldsOutput)
}

func (a WebAclLoggingConfigurationRedactedFieldsArgs) ToWebAclLoggingConfigurationRedactedFieldsOutputWithContext(ctx context.Context) WebAclLoggingConfigurationRedactedFieldsOutput {
	return pulumi.ToOutputWithContext(ctx, a).(WebAclLoggingConfigurationRedactedFieldsOutput)
}

type WebAclLoggingConfigurationRedactedFieldsOutput struct { *pulumi.OutputState }

// Set of configuration blocks for fields to redact. Detailed below.
func (o WebAclLoggingConfigurationRedactedFieldsOutput) FieldToMatches() WebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayOutput {
	return o.Apply(func(v WebAclLoggingConfigurationRedactedFields) []WebAclLoggingConfigurationRedactedFieldsFieldToMatches {
		return v.FieldToMatches
	}).(WebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayOutput)
}

func (WebAclLoggingConfigurationRedactedFieldsOutput) ElementType() reflect.Type {
	return webAclLoggingConfigurationRedactedFieldsType
}

func (o WebAclLoggingConfigurationRedactedFieldsOutput) ToWebAclLoggingConfigurationRedactedFieldsOutput() WebAclLoggingConfigurationRedactedFieldsOutput {
	return o
}

func (o WebAclLoggingConfigurationRedactedFieldsOutput) ToWebAclLoggingConfigurationRedactedFieldsOutputWithContext(ctx context.Context) WebAclLoggingConfigurationRedactedFieldsOutput {
	return o
}

func init() { pulumi.RegisterOutputType(WebAclLoggingConfigurationRedactedFieldsOutput{}) }

type WebAclLoggingConfigurationRedactedFieldsFieldToMatches struct {
	// When the value of `type` is `HEADER`, enter the name of the header that you want the WAF to search, for example, `User-Agent` or `Referer`. If the value of `type` is any other value, omit `data`.
	Data *string `pulumi:"data"`
	// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`
	Type string `pulumi:"type"`
}
var webAclLoggingConfigurationRedactedFieldsFieldToMatchesType = reflect.TypeOf((*WebAclLoggingConfigurationRedactedFieldsFieldToMatches)(nil)).Elem()

type WebAclLoggingConfigurationRedactedFieldsFieldToMatchesInput interface {
	pulumi.Input

	ToWebAclLoggingConfigurationRedactedFieldsFieldToMatchesOutput() WebAclLoggingConfigurationRedactedFieldsFieldToMatchesOutput
	ToWebAclLoggingConfigurationRedactedFieldsFieldToMatchesOutputWithContext(ctx context.Context) WebAclLoggingConfigurationRedactedFieldsFieldToMatchesOutput
}

type WebAclLoggingConfigurationRedactedFieldsFieldToMatchesArgs struct {
	// When the value of `type` is `HEADER`, enter the name of the header that you want the WAF to search, for example, `User-Agent` or `Referer`. If the value of `type` is any other value, omit `data`.
	Data pulumi.StringInput `pulumi:"data"`
	// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`
	Type pulumi.StringInput `pulumi:"type"`
}

func (WebAclLoggingConfigurationRedactedFieldsFieldToMatchesArgs) ElementType() reflect.Type {
	return webAclLoggingConfigurationRedactedFieldsFieldToMatchesType
}

func (a WebAclLoggingConfigurationRedactedFieldsFieldToMatchesArgs) ToWebAclLoggingConfigurationRedactedFieldsFieldToMatchesOutput() WebAclLoggingConfigurationRedactedFieldsFieldToMatchesOutput {
	return pulumi.ToOutput(a).(WebAclLoggingConfigurationRedactedFieldsFieldToMatchesOutput)
}

func (a WebAclLoggingConfigurationRedactedFieldsFieldToMatchesArgs) ToWebAclLoggingConfigurationRedactedFieldsFieldToMatchesOutputWithContext(ctx context.Context) WebAclLoggingConfigurationRedactedFieldsFieldToMatchesOutput {
	return pulumi.ToOutputWithContext(ctx, a).(WebAclLoggingConfigurationRedactedFieldsFieldToMatchesOutput)
}

type WebAclLoggingConfigurationRedactedFieldsFieldToMatchesOutput struct { *pulumi.OutputState }

// When the value of `type` is `HEADER`, enter the name of the header that you want the WAF to search, for example, `User-Agent` or `Referer`. If the value of `type` is any other value, omit `data`.
func (o WebAclLoggingConfigurationRedactedFieldsFieldToMatchesOutput) Data() pulumi.StringOutput {
	return o.Apply(func(v WebAclLoggingConfigurationRedactedFieldsFieldToMatches) string {
		if v.Data == nil { return *new(string) } else { return *v.Data }
	}).(pulumi.StringOutput)
}

// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`
func (o WebAclLoggingConfigurationRedactedFieldsFieldToMatchesOutput) Type() pulumi.StringOutput {
	return o.Apply(func(v WebAclLoggingConfigurationRedactedFieldsFieldToMatches) string {
		return v.Type
	}).(pulumi.StringOutput)
}

func (WebAclLoggingConfigurationRedactedFieldsFieldToMatchesOutput) ElementType() reflect.Type {
	return webAclLoggingConfigurationRedactedFieldsFieldToMatchesType
}

func (o WebAclLoggingConfigurationRedactedFieldsFieldToMatchesOutput) ToWebAclLoggingConfigurationRedactedFieldsFieldToMatchesOutput() WebAclLoggingConfigurationRedactedFieldsFieldToMatchesOutput {
	return o
}

func (o WebAclLoggingConfigurationRedactedFieldsFieldToMatchesOutput) ToWebAclLoggingConfigurationRedactedFieldsFieldToMatchesOutputWithContext(ctx context.Context) WebAclLoggingConfigurationRedactedFieldsFieldToMatchesOutput {
	return o
}

func init() { pulumi.RegisterOutputType(WebAclLoggingConfigurationRedactedFieldsFieldToMatchesOutput{}) }

var webAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayType = reflect.TypeOf((*[]WebAclLoggingConfigurationRedactedFieldsFieldToMatches)(nil)).Elem()

type WebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayInput interface {
	pulumi.Input

	ToWebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayOutput() WebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayOutput
	ToWebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayOutputWithContext(ctx context.Context) WebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayOutput
}

type WebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayArgs []WebAclLoggingConfigurationRedactedFieldsFieldToMatchesInput

func (WebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayArgs) ElementType() reflect.Type {
	return webAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayType
}

func (a WebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayArgs) ToWebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayOutput() WebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayOutput {
	return pulumi.ToOutput(a).(WebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayOutput)
}

func (a WebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayArgs) ToWebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayOutputWithContext(ctx context.Context) WebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(WebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayOutput)
}

type WebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayOutput struct { *pulumi.OutputState }

func (o WebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayOutput) Index(i pulumi.IntInput) WebAclLoggingConfigurationRedactedFieldsFieldToMatchesOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) WebAclLoggingConfigurationRedactedFieldsFieldToMatches {
		return vs[0].([]WebAclLoggingConfigurationRedactedFieldsFieldToMatches)[vs[1].(int)]
	}).(WebAclLoggingConfigurationRedactedFieldsFieldToMatchesOutput)
}

func (WebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayOutput) ElementType() reflect.Type {
	return webAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayType
}

func (o WebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayOutput) ToWebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayOutput() WebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayOutput {
	return o
}

func (o WebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayOutput) ToWebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayOutputWithContext(ctx context.Context) WebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(WebAclLoggingConfigurationRedactedFieldsFieldToMatchesArrayOutput{}) }

type WebAclRules struct {
	// Configuration block of the action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule.  Not used if `type` is `GROUP`. Detailed below.
	Action *WebAclRulesAction `pulumi:"action"`
	// Configuration block of the override the action that a group requests CloudFront or AWS WAF takes when a web request matches the conditions in the rule.  Only used if `type` is `GROUP`. Detailed below.
	OverrideAction *WebAclRulesOverrideAction `pulumi:"overrideAction"`
	// Specifies the order in which the rules in a WebACL are evaluated.
	// Rules with a lower value are evaluated before rules with a higher value.
	Priority int `pulumi:"priority"`
	// ID of the associated WAF (Regional) rule (e.g. [`wafregional.Rule`](https://www.terraform.io/docs/providers/aws/r/wafregional_rule.html)). WAF (Global) rules cannot be used.
	RuleId string `pulumi:"ruleId"`
	// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`
	Type *string `pulumi:"type"`
}
var webAclRulesType = reflect.TypeOf((*WebAclRules)(nil)).Elem()

type WebAclRulesInput interface {
	pulumi.Input

	ToWebAclRulesOutput() WebAclRulesOutput
	ToWebAclRulesOutputWithContext(ctx context.Context) WebAclRulesOutput
}

type WebAclRulesArgs struct {
	// Configuration block of the action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule.  Not used if `type` is `GROUP`. Detailed below.
	Action WebAclRulesActionInput `pulumi:"action"`
	// Configuration block of the override the action that a group requests CloudFront or AWS WAF takes when a web request matches the conditions in the rule.  Only used if `type` is `GROUP`. Detailed below.
	OverrideAction WebAclRulesOverrideActionInput `pulumi:"overrideAction"`
	// Specifies the order in which the rules in a WebACL are evaluated.
	// Rules with a lower value are evaluated before rules with a higher value.
	Priority pulumi.IntInput `pulumi:"priority"`
	// ID of the associated WAF (Regional) rule (e.g. [`wafregional.Rule`](https://www.terraform.io/docs/providers/aws/r/wafregional_rule.html)). WAF (Global) rules cannot be used.
	RuleId pulumi.StringInput `pulumi:"ruleId"`
	// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`
	Type pulumi.StringInput `pulumi:"type"`
}

func (WebAclRulesArgs) ElementType() reflect.Type {
	return webAclRulesType
}

func (a WebAclRulesArgs) ToWebAclRulesOutput() WebAclRulesOutput {
	return pulumi.ToOutput(a).(WebAclRulesOutput)
}

func (a WebAclRulesArgs) ToWebAclRulesOutputWithContext(ctx context.Context) WebAclRulesOutput {
	return pulumi.ToOutputWithContext(ctx, a).(WebAclRulesOutput)
}

type WebAclRulesOutput struct { *pulumi.OutputState }

// Configuration block of the action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule.  Not used if `type` is `GROUP`. Detailed below.
func (o WebAclRulesOutput) Action() WebAclRulesActionOutput {
	return o.Apply(func(v WebAclRules) WebAclRulesAction {
		if v.Action == nil { return *new(WebAclRulesAction) } else { return *v.Action }
	}).(WebAclRulesActionOutput)
}

// Configuration block of the override the action that a group requests CloudFront or AWS WAF takes when a web request matches the conditions in the rule.  Only used if `type` is `GROUP`. Detailed below.
func (o WebAclRulesOutput) OverrideAction() WebAclRulesOverrideActionOutput {
	return o.Apply(func(v WebAclRules) WebAclRulesOverrideAction {
		if v.OverrideAction == nil { return *new(WebAclRulesOverrideAction) } else { return *v.OverrideAction }
	}).(WebAclRulesOverrideActionOutput)
}

// Specifies the order in which the rules in a WebACL are evaluated.
// Rules with a lower value are evaluated before rules with a higher value.
func (o WebAclRulesOutput) Priority() pulumi.IntOutput {
	return o.Apply(func(v WebAclRules) int {
		return v.Priority
	}).(pulumi.IntOutput)
}

// ID of the associated WAF (Regional) rule (e.g. [`wafregional.Rule`](https://www.terraform.io/docs/providers/aws/r/wafregional_rule.html)). WAF (Global) rules cannot be used.
func (o WebAclRulesOutput) RuleId() pulumi.StringOutput {
	return o.Apply(func(v WebAclRules) string {
		return v.RuleId
	}).(pulumi.StringOutput)
}

// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`
func (o WebAclRulesOutput) Type() pulumi.StringOutput {
	return o.Apply(func(v WebAclRules) string {
		if v.Type == nil { return *new(string) } else { return *v.Type }
	}).(pulumi.StringOutput)
}

func (WebAclRulesOutput) ElementType() reflect.Type {
	return webAclRulesType
}

func (o WebAclRulesOutput) ToWebAclRulesOutput() WebAclRulesOutput {
	return o
}

func (o WebAclRulesOutput) ToWebAclRulesOutputWithContext(ctx context.Context) WebAclRulesOutput {
	return o
}

func init() { pulumi.RegisterOutputType(WebAclRulesOutput{}) }

type WebAclRulesAction struct {
	// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`
	Type string `pulumi:"type"`
}
var webAclRulesActionType = reflect.TypeOf((*WebAclRulesAction)(nil)).Elem()

type WebAclRulesActionInput interface {
	pulumi.Input

	ToWebAclRulesActionOutput() WebAclRulesActionOutput
	ToWebAclRulesActionOutputWithContext(ctx context.Context) WebAclRulesActionOutput
}

type WebAclRulesActionArgs struct {
	// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`
	Type pulumi.StringInput `pulumi:"type"`
}

func (WebAclRulesActionArgs) ElementType() reflect.Type {
	return webAclRulesActionType
}

func (a WebAclRulesActionArgs) ToWebAclRulesActionOutput() WebAclRulesActionOutput {
	return pulumi.ToOutput(a).(WebAclRulesActionOutput)
}

func (a WebAclRulesActionArgs) ToWebAclRulesActionOutputWithContext(ctx context.Context) WebAclRulesActionOutput {
	return pulumi.ToOutputWithContext(ctx, a).(WebAclRulesActionOutput)
}

type WebAclRulesActionOutput struct { *pulumi.OutputState }

// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`
func (o WebAclRulesActionOutput) Type() pulumi.StringOutput {
	return o.Apply(func(v WebAclRulesAction) string {
		return v.Type
	}).(pulumi.StringOutput)
}

func (WebAclRulesActionOutput) ElementType() reflect.Type {
	return webAclRulesActionType
}

func (o WebAclRulesActionOutput) ToWebAclRulesActionOutput() WebAclRulesActionOutput {
	return o
}

func (o WebAclRulesActionOutput) ToWebAclRulesActionOutputWithContext(ctx context.Context) WebAclRulesActionOutput {
	return o
}

func init() { pulumi.RegisterOutputType(WebAclRulesActionOutput{}) }

var webAclRulesArrayType = reflect.TypeOf((*[]WebAclRules)(nil)).Elem()

type WebAclRulesArrayInput interface {
	pulumi.Input

	ToWebAclRulesArrayOutput() WebAclRulesArrayOutput
	ToWebAclRulesArrayOutputWithContext(ctx context.Context) WebAclRulesArrayOutput
}

type WebAclRulesArrayArgs []WebAclRulesInput

func (WebAclRulesArrayArgs) ElementType() reflect.Type {
	return webAclRulesArrayType
}

func (a WebAclRulesArrayArgs) ToWebAclRulesArrayOutput() WebAclRulesArrayOutput {
	return pulumi.ToOutput(a).(WebAclRulesArrayOutput)
}

func (a WebAclRulesArrayArgs) ToWebAclRulesArrayOutputWithContext(ctx context.Context) WebAclRulesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(WebAclRulesArrayOutput)
}

type WebAclRulesArrayOutput struct { *pulumi.OutputState }

func (o WebAclRulesArrayOutput) Index(i pulumi.IntInput) WebAclRulesOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) WebAclRules {
		return vs[0].([]WebAclRules)[vs[1].(int)]
	}).(WebAclRulesOutput)
}

func (WebAclRulesArrayOutput) ElementType() reflect.Type {
	return webAclRulesArrayType
}

func (o WebAclRulesArrayOutput) ToWebAclRulesArrayOutput() WebAclRulesArrayOutput {
	return o
}

func (o WebAclRulesArrayOutput) ToWebAclRulesArrayOutputWithContext(ctx context.Context) WebAclRulesArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(WebAclRulesArrayOutput{}) }

type WebAclRulesOverrideAction struct {
	// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`
	Type string `pulumi:"type"`
}
var webAclRulesOverrideActionType = reflect.TypeOf((*WebAclRulesOverrideAction)(nil)).Elem()

type WebAclRulesOverrideActionInput interface {
	pulumi.Input

	ToWebAclRulesOverrideActionOutput() WebAclRulesOverrideActionOutput
	ToWebAclRulesOverrideActionOutputWithContext(ctx context.Context) WebAclRulesOverrideActionOutput
}

type WebAclRulesOverrideActionArgs struct {
	// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`
	Type pulumi.StringInput `pulumi:"type"`
}

func (WebAclRulesOverrideActionArgs) ElementType() reflect.Type {
	return webAclRulesOverrideActionType
}

func (a WebAclRulesOverrideActionArgs) ToWebAclRulesOverrideActionOutput() WebAclRulesOverrideActionOutput {
	return pulumi.ToOutput(a).(WebAclRulesOverrideActionOutput)
}

func (a WebAclRulesOverrideActionArgs) ToWebAclRulesOverrideActionOutputWithContext(ctx context.Context) WebAclRulesOverrideActionOutput {
	return pulumi.ToOutputWithContext(ctx, a).(WebAclRulesOverrideActionOutput)
}

type WebAclRulesOverrideActionOutput struct { *pulumi.OutputState }

// Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule. e.g. `ALLOW`, `BLOCK` or `COUNT`
func (o WebAclRulesOverrideActionOutput) Type() pulumi.StringOutput {
	return o.Apply(func(v WebAclRulesOverrideAction) string {
		return v.Type
	}).(pulumi.StringOutput)
}

func (WebAclRulesOverrideActionOutput) ElementType() reflect.Type {
	return webAclRulesOverrideActionType
}

func (o WebAclRulesOverrideActionOutput) ToWebAclRulesOverrideActionOutput() WebAclRulesOverrideActionOutput {
	return o
}

func (o WebAclRulesOverrideActionOutput) ToWebAclRulesOverrideActionOutputWithContext(ctx context.Context) WebAclRulesOverrideActionOutput {
	return o
}

func init() { pulumi.RegisterOutputType(WebAclRulesOverrideActionOutput{}) }

