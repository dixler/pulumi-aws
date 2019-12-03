// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an WAF Regional Rule Resource for use with Application Load Balancer.
// 
// ## Nested Fields
// 
// ### `predicate`
// 
// See the [WAF Documentation](https://docs.aws.amazon.com/waf/latest/APIReference/API_Predicate.html) for more information.
// 
// #### Arguments
// 
// * `type` - (Required) The type of predicate in a rule. Valid values: `ByteMatch`, `GeoMatch`, `IPMatch`, `RegexMatch`, `SizeConstraint`, `SqlInjectionMatch`, or `XssMatch`
// * `dataId` - (Required) The unique identifier of a predicate, such as the ID of a `ByteMatchSet` or `IPSet`.
// * `negated` - (Required) Whether to use the settings or the negated settings that you specified in the objects.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/wafregional_rule.html.markdown.
type Rule struct {
	pulumi.CustomResourceState

	// The ARN of the WAF Regional Rule.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName pulumi.StringOutput `pulumi:"metricName"`

	// The name or description of the rule.
	Name pulumi.StringOutput `pulumi:"name"`

	// The objects to include in a rule (documented below).
	Predicates RulePredicatesArrayOutput `pulumi:"predicates"`

	// Key-value mapping of resource tags
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewRule registers a new resource with the given unique name, arguments, and options.
func NewRule(ctx *pulumi.Context,
	name string, args *RuleArgs, opts ...pulumi.ResourceOption) (*Rule, error) {
	if args == nil || args.MetricName == nil {
		return nil, errors.New("missing required argument 'MetricName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.MetricName; i != nil { inputs["metricName"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.Predicates; i != nil { inputs["predicates"] = i.ToRulePredicatesArrayOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource Rule
	err := ctx.RegisterResource("aws:wafregional/rule:Rule", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRule gets an existing Rule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleState, opts ...pulumi.ResourceOption) (*Rule, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Arn; i != nil { inputs["arn"] = i.ToStringOutput() }
		if i := state.MetricName; i != nil { inputs["metricName"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.Predicates; i != nil { inputs["predicates"] = i.ToRulePredicatesArrayOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource Rule
	err := ctx.ReadResource("aws:wafregional/rule:Rule", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rule resources.
type RuleState struct {
	// The ARN of the WAF Regional Rule.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName pulumi.StringInput `pulumi:"metricName"`
	// The name or description of the rule.
	Name pulumi.StringInput `pulumi:"name"`
	// The objects to include in a rule (documented below).
	Predicates RulePredicatesArrayInput `pulumi:"predicates"`
	// Key-value mapping of resource tags
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a Rule resource.
type RuleArgs struct {
	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName pulumi.StringInput `pulumi:"metricName"`
	// The name or description of the rule.
	Name pulumi.StringInput `pulumi:"name"`
	// The objects to include in a rule (documented below).
	Predicates RulePredicatesArrayInput `pulumi:"predicates"`
	// Key-value mapping of resource tags
	Tags pulumi.MapInput `pulumi:"tags"`
}
type RulePredicates struct {
	DataId string `pulumi:"dataId"`
	Negated bool `pulumi:"negated"`
	Type string `pulumi:"type"`
}
var rulePredicatesType = reflect.TypeOf((*RulePredicates)(nil)).Elem()

type RulePredicatesInput interface {
	pulumi.Input

	ToRulePredicatesOutput() RulePredicatesOutput
	ToRulePredicatesOutputWithContext(ctx context.Context) RulePredicatesOutput
}

type RulePredicatesArgs struct {
	DataId pulumi.StringInput `pulumi:"dataId"`
	Negated pulumi.BoolInput `pulumi:"negated"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (RulePredicatesArgs) ElementType() reflect.Type {
	return rulePredicatesType
}

func (a RulePredicatesArgs) ToRulePredicatesOutput() RulePredicatesOutput {
	return pulumi.ToOutput(a).(RulePredicatesOutput)
}

func (a RulePredicatesArgs) ToRulePredicatesOutputWithContext(ctx context.Context) RulePredicatesOutput {
	return pulumi.ToOutputWithContext(ctx, a).(RulePredicatesOutput)
}

type RulePredicatesOutput struct { *pulumi.OutputState }

func (o RulePredicatesOutput) DataId() pulumi.StringOutput {
	return o.Apply(func(v RulePredicates) string {
		return v.DataId
	}).(pulumi.StringOutput)
}

func (o RulePredicatesOutput) Negated() pulumi.BoolOutput {
	return o.Apply(func(v RulePredicates) bool {
		return v.Negated
	}).(pulumi.BoolOutput)
}

func (o RulePredicatesOutput) Type() pulumi.StringOutput {
	return o.Apply(func(v RulePredicates) string {
		return v.Type
	}).(pulumi.StringOutput)
}

func (RulePredicatesOutput) ElementType() reflect.Type {
	return rulePredicatesType
}

func (o RulePredicatesOutput) ToRulePredicatesOutput() RulePredicatesOutput {
	return o
}

func (o RulePredicatesOutput) ToRulePredicatesOutputWithContext(ctx context.Context) RulePredicatesOutput {
	return o
}

func init() { pulumi.RegisterOutputType(RulePredicatesOutput{}) }

var rulePredicatesArrayType = reflect.TypeOf((*[]RulePredicates)(nil)).Elem()

type RulePredicatesArrayInput interface {
	pulumi.Input

	ToRulePredicatesArrayOutput() RulePredicatesArrayOutput
	ToRulePredicatesArrayOutputWithContext(ctx context.Context) RulePredicatesArrayOutput
}

type RulePredicatesArrayArgs []RulePredicatesInput

func (RulePredicatesArrayArgs) ElementType() reflect.Type {
	return rulePredicatesArrayType
}

func (a RulePredicatesArrayArgs) ToRulePredicatesArrayOutput() RulePredicatesArrayOutput {
	return pulumi.ToOutput(a).(RulePredicatesArrayOutput)
}

func (a RulePredicatesArrayArgs) ToRulePredicatesArrayOutputWithContext(ctx context.Context) RulePredicatesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(RulePredicatesArrayOutput)
}

type RulePredicatesArrayOutput struct { *pulumi.OutputState }

func (o RulePredicatesArrayOutput) Index(i pulumi.IntInput) RulePredicatesOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) RulePredicates {
		return vs[0].([]RulePredicates)[vs[1].(int)]
	}).(RulePredicatesOutput)
}

func (RulePredicatesArrayOutput) ElementType() reflect.Type {
	return rulePredicatesArrayType
}

func (o RulePredicatesArrayOutput) ToRulePredicatesArrayOutput() RulePredicatesArrayOutput {
	return o
}

func (o RulePredicatesArrayOutput) ToRulePredicatesArrayOutputWithContext(ctx context.Context) RulePredicatesArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(RulePredicatesArrayOutput{}) }

