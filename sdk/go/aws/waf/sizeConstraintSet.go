// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a WAF Size Constraint Set Resource
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/waf_size_constraint_set.html.markdown.
type SizeConstraintSet struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN)
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The name or description of the Size Constraint Set.
	Name pulumi.StringOutput `pulumi:"name"`

	// Specifies the parts of web requests that you want to inspect the size of.
	SizeConstraints SizeConstraintSetSizeConstraintsArrayOutput `pulumi:"sizeConstraints"`
}

// NewSizeConstraintSet registers a new resource with the given unique name, arguments, and options.
func NewSizeConstraintSet(ctx *pulumi.Context,
	name string, args *SizeConstraintSetArgs, opts ...pulumi.ResourceOption) (*SizeConstraintSet, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.SizeConstraints; i != nil { inputs["sizeConstraints"] = i.ToSizeConstraintSetSizeConstraintsArrayOutput() }
	}
	var resource SizeConstraintSet
	err := ctx.RegisterResource("aws:waf/sizeConstraintSet:SizeConstraintSet", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSizeConstraintSet gets an existing SizeConstraintSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSizeConstraintSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SizeConstraintSetState, opts ...pulumi.ResourceOption) (*SizeConstraintSet, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Arn; i != nil { inputs["arn"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.SizeConstraints; i != nil { inputs["sizeConstraints"] = i.ToSizeConstraintSetSizeConstraintsArrayOutput() }
	}
	var resource SizeConstraintSet
	err := ctx.ReadResource("aws:waf/sizeConstraintSet:SizeConstraintSet", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SizeConstraintSet resources.
type SizeConstraintSetState struct {
	// Amazon Resource Name (ARN)
	Arn pulumi.StringInput `pulumi:"arn"`
	// The name or description of the Size Constraint Set.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the parts of web requests that you want to inspect the size of.
	SizeConstraints SizeConstraintSetSizeConstraintsArrayInput `pulumi:"sizeConstraints"`
}

// The set of arguments for constructing a SizeConstraintSet resource.
type SizeConstraintSetArgs struct {
	// The name or description of the Size Constraint Set.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the parts of web requests that you want to inspect the size of.
	SizeConstraints SizeConstraintSetSizeConstraintsArrayInput `pulumi:"sizeConstraints"`
}
type SizeConstraintSetSizeConstraints struct {
	// The type of comparison you want to perform.
	// e.g. `EQ`, `NE`, `LT`, `GT`.
	// See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_SizeConstraint.html#WAF-Type-SizeConstraint-ComparisonOperator) for all supported values.
	ComparisonOperator string `pulumi:"comparisonOperator"`
	// Specifies where in a web request to look for the size constraint.
	FieldToMatch SizeConstraintSetSizeConstraintsFieldToMatch `pulumi:"fieldToMatch"`
	// The size in bytes that you want to compare against the size of the specified `fieldToMatch`.
	// Valid values are between 0 - 21474836480 bytes (0 - 20 GB).
	Size int `pulumi:"size"`
	// Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
	// If you specify a transformation, AWS WAF performs the transformation on `fieldToMatch` before inspecting a request for a match.
	// e.g. `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
	// See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_SizeConstraint.html#WAF-Type-SizeConstraint-TextTransformation)
	// for all supported values.
	// **Note:** if you choose `BODY` as `type`, you must choose `NONE` because CloudFront forwards only the first 8192 bytes for inspection.
	TextTransformation string `pulumi:"textTransformation"`
}
var sizeConstraintSetSizeConstraintsType = reflect.TypeOf((*SizeConstraintSetSizeConstraints)(nil)).Elem()

type SizeConstraintSetSizeConstraintsInput interface {
	pulumi.Input

	ToSizeConstraintSetSizeConstraintsOutput() SizeConstraintSetSizeConstraintsOutput
	ToSizeConstraintSetSizeConstraintsOutputWithContext(ctx context.Context) SizeConstraintSetSizeConstraintsOutput
}

type SizeConstraintSetSizeConstraintsArgs struct {
	// The type of comparison you want to perform.
	// e.g. `EQ`, `NE`, `LT`, `GT`.
	// See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_SizeConstraint.html#WAF-Type-SizeConstraint-ComparisonOperator) for all supported values.
	ComparisonOperator pulumi.StringInput `pulumi:"comparisonOperator"`
	// Specifies where in a web request to look for the size constraint.
	FieldToMatch SizeConstraintSetSizeConstraintsFieldToMatchInput `pulumi:"fieldToMatch"`
	// The size in bytes that you want to compare against the size of the specified `fieldToMatch`.
	// Valid values are between 0 - 21474836480 bytes (0 - 20 GB).
	Size pulumi.IntInput `pulumi:"size"`
	// Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
	// If you specify a transformation, AWS WAF performs the transformation on `fieldToMatch` before inspecting a request for a match.
	// e.g. `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
	// See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_SizeConstraint.html#WAF-Type-SizeConstraint-TextTransformation)
	// for all supported values.
	// **Note:** if you choose `BODY` as `type`, you must choose `NONE` because CloudFront forwards only the first 8192 bytes for inspection.
	TextTransformation pulumi.StringInput `pulumi:"textTransformation"`
}

func (SizeConstraintSetSizeConstraintsArgs) ElementType() reflect.Type {
	return sizeConstraintSetSizeConstraintsType
}

func (a SizeConstraintSetSizeConstraintsArgs) ToSizeConstraintSetSizeConstraintsOutput() SizeConstraintSetSizeConstraintsOutput {
	return pulumi.ToOutput(a).(SizeConstraintSetSizeConstraintsOutput)
}

func (a SizeConstraintSetSizeConstraintsArgs) ToSizeConstraintSetSizeConstraintsOutputWithContext(ctx context.Context) SizeConstraintSetSizeConstraintsOutput {
	return pulumi.ToOutputWithContext(ctx, a).(SizeConstraintSetSizeConstraintsOutput)
}

type SizeConstraintSetSizeConstraintsOutput struct { *pulumi.OutputState }

// The type of comparison you want to perform.
// e.g. `EQ`, `NE`, `LT`, `GT`.
// See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_SizeConstraint.html#WAF-Type-SizeConstraint-ComparisonOperator) for all supported values.
func (o SizeConstraintSetSizeConstraintsOutput) ComparisonOperator() pulumi.StringOutput {
	return o.Apply(func(v SizeConstraintSetSizeConstraints) string {
		return v.ComparisonOperator
	}).(pulumi.StringOutput)
}

// Specifies where in a web request to look for the size constraint.
func (o SizeConstraintSetSizeConstraintsOutput) FieldToMatch() SizeConstraintSetSizeConstraintsFieldToMatchOutput {
	return o.Apply(func(v SizeConstraintSetSizeConstraints) SizeConstraintSetSizeConstraintsFieldToMatch {
		return v.FieldToMatch
	}).(SizeConstraintSetSizeConstraintsFieldToMatchOutput)
}

// The size in bytes that you want to compare against the size of the specified `fieldToMatch`.
// Valid values are between 0 - 21474836480 bytes (0 - 20 GB).
func (o SizeConstraintSetSizeConstraintsOutput) Size() pulumi.IntOutput {
	return o.Apply(func(v SizeConstraintSetSizeConstraints) int {
		return v.Size
	}).(pulumi.IntOutput)
}

// Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
// If you specify a transformation, AWS WAF performs the transformation on `fieldToMatch` before inspecting a request for a match.
// e.g. `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
// See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_SizeConstraint.html#WAF-Type-SizeConstraint-TextTransformation)
// for all supported values.
// **Note:** if you choose `BODY` as `type`, you must choose `NONE` because CloudFront forwards only the first 8192 bytes for inspection.
func (o SizeConstraintSetSizeConstraintsOutput) TextTransformation() pulumi.StringOutput {
	return o.Apply(func(v SizeConstraintSetSizeConstraints) string {
		return v.TextTransformation
	}).(pulumi.StringOutput)
}

func (SizeConstraintSetSizeConstraintsOutput) ElementType() reflect.Type {
	return sizeConstraintSetSizeConstraintsType
}

func (o SizeConstraintSetSizeConstraintsOutput) ToSizeConstraintSetSizeConstraintsOutput() SizeConstraintSetSizeConstraintsOutput {
	return o
}

func (o SizeConstraintSetSizeConstraintsOutput) ToSizeConstraintSetSizeConstraintsOutputWithContext(ctx context.Context) SizeConstraintSetSizeConstraintsOutput {
	return o
}

func init() { pulumi.RegisterOutputType(SizeConstraintSetSizeConstraintsOutput{}) }

var sizeConstraintSetSizeConstraintsArrayType = reflect.TypeOf((*[]SizeConstraintSetSizeConstraints)(nil)).Elem()

type SizeConstraintSetSizeConstraintsArrayInput interface {
	pulumi.Input

	ToSizeConstraintSetSizeConstraintsArrayOutput() SizeConstraintSetSizeConstraintsArrayOutput
	ToSizeConstraintSetSizeConstraintsArrayOutputWithContext(ctx context.Context) SizeConstraintSetSizeConstraintsArrayOutput
}

type SizeConstraintSetSizeConstraintsArrayArgs []SizeConstraintSetSizeConstraintsInput

func (SizeConstraintSetSizeConstraintsArrayArgs) ElementType() reflect.Type {
	return sizeConstraintSetSizeConstraintsArrayType
}

func (a SizeConstraintSetSizeConstraintsArrayArgs) ToSizeConstraintSetSizeConstraintsArrayOutput() SizeConstraintSetSizeConstraintsArrayOutput {
	return pulumi.ToOutput(a).(SizeConstraintSetSizeConstraintsArrayOutput)
}

func (a SizeConstraintSetSizeConstraintsArrayArgs) ToSizeConstraintSetSizeConstraintsArrayOutputWithContext(ctx context.Context) SizeConstraintSetSizeConstraintsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(SizeConstraintSetSizeConstraintsArrayOutput)
}

type SizeConstraintSetSizeConstraintsArrayOutput struct { *pulumi.OutputState }

func (o SizeConstraintSetSizeConstraintsArrayOutput) Index(i pulumi.IntInput) SizeConstraintSetSizeConstraintsOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) SizeConstraintSetSizeConstraints {
		return vs[0].([]SizeConstraintSetSizeConstraints)[vs[1].(int)]
	}).(SizeConstraintSetSizeConstraintsOutput)
}

func (SizeConstraintSetSizeConstraintsArrayOutput) ElementType() reflect.Type {
	return sizeConstraintSetSizeConstraintsArrayType
}

func (o SizeConstraintSetSizeConstraintsArrayOutput) ToSizeConstraintSetSizeConstraintsArrayOutput() SizeConstraintSetSizeConstraintsArrayOutput {
	return o
}

func (o SizeConstraintSetSizeConstraintsArrayOutput) ToSizeConstraintSetSizeConstraintsArrayOutputWithContext(ctx context.Context) SizeConstraintSetSizeConstraintsArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(SizeConstraintSetSizeConstraintsArrayOutput{}) }

type SizeConstraintSetSizeConstraintsFieldToMatch struct {
	// When `type` is `HEADER`, enter the name of the header that you want to search, e.g. `User-Agent` or `Referer`.
	// If `type` is any other value, omit this field.
	Data *string `pulumi:"data"`
	// The part of the web request that you want AWS WAF to search for a specified string.
	// e.g. `HEADER`, `METHOD` or `BODY`.
	// See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_FieldToMatch.html)
	// for all supported values.
	Type string `pulumi:"type"`
}
var sizeConstraintSetSizeConstraintsFieldToMatchType = reflect.TypeOf((*SizeConstraintSetSizeConstraintsFieldToMatch)(nil)).Elem()

type SizeConstraintSetSizeConstraintsFieldToMatchInput interface {
	pulumi.Input

	ToSizeConstraintSetSizeConstraintsFieldToMatchOutput() SizeConstraintSetSizeConstraintsFieldToMatchOutput
	ToSizeConstraintSetSizeConstraintsFieldToMatchOutputWithContext(ctx context.Context) SizeConstraintSetSizeConstraintsFieldToMatchOutput
}

type SizeConstraintSetSizeConstraintsFieldToMatchArgs struct {
	// When `type` is `HEADER`, enter the name of the header that you want to search, e.g. `User-Agent` or `Referer`.
	// If `type` is any other value, omit this field.
	Data pulumi.StringInput `pulumi:"data"`
	// The part of the web request that you want AWS WAF to search for a specified string.
	// e.g. `HEADER`, `METHOD` or `BODY`.
	// See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_FieldToMatch.html)
	// for all supported values.
	Type pulumi.StringInput `pulumi:"type"`
}

func (SizeConstraintSetSizeConstraintsFieldToMatchArgs) ElementType() reflect.Type {
	return sizeConstraintSetSizeConstraintsFieldToMatchType
}

func (a SizeConstraintSetSizeConstraintsFieldToMatchArgs) ToSizeConstraintSetSizeConstraintsFieldToMatchOutput() SizeConstraintSetSizeConstraintsFieldToMatchOutput {
	return pulumi.ToOutput(a).(SizeConstraintSetSizeConstraintsFieldToMatchOutput)
}

func (a SizeConstraintSetSizeConstraintsFieldToMatchArgs) ToSizeConstraintSetSizeConstraintsFieldToMatchOutputWithContext(ctx context.Context) SizeConstraintSetSizeConstraintsFieldToMatchOutput {
	return pulumi.ToOutputWithContext(ctx, a).(SizeConstraintSetSizeConstraintsFieldToMatchOutput)
}

type SizeConstraintSetSizeConstraintsFieldToMatchOutput struct { *pulumi.OutputState }

// When `type` is `HEADER`, enter the name of the header that you want to search, e.g. `User-Agent` or `Referer`.
// If `type` is any other value, omit this field.
func (o SizeConstraintSetSizeConstraintsFieldToMatchOutput) Data() pulumi.StringOutput {
	return o.Apply(func(v SizeConstraintSetSizeConstraintsFieldToMatch) string {
		if v.Data == nil { return *new(string) } else { return *v.Data }
	}).(pulumi.StringOutput)
}

// The part of the web request that you want AWS WAF to search for a specified string.
// e.g. `HEADER`, `METHOD` or `BODY`.
// See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_FieldToMatch.html)
// for all supported values.
func (o SizeConstraintSetSizeConstraintsFieldToMatchOutput) Type() pulumi.StringOutput {
	return o.Apply(func(v SizeConstraintSetSizeConstraintsFieldToMatch) string {
		return v.Type
	}).(pulumi.StringOutput)
}

func (SizeConstraintSetSizeConstraintsFieldToMatchOutput) ElementType() reflect.Type {
	return sizeConstraintSetSizeConstraintsFieldToMatchType
}

func (o SizeConstraintSetSizeConstraintsFieldToMatchOutput) ToSizeConstraintSetSizeConstraintsFieldToMatchOutput() SizeConstraintSetSizeConstraintsFieldToMatchOutput {
	return o
}

func (o SizeConstraintSetSizeConstraintsFieldToMatchOutput) ToSizeConstraintSetSizeConstraintsFieldToMatchOutputWithContext(ctx context.Context) SizeConstraintSetSizeConstraintsFieldToMatchOutput {
	return o
}

func init() { pulumi.RegisterOutputType(SizeConstraintSetSizeConstraintsFieldToMatchOutput{}) }

