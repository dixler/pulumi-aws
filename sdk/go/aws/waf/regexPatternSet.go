// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a WAF Regex Pattern Set Resource
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/waf_regex_pattern_set.html.markdown.
type RegexPatternSet struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN)
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The name or description of the Regex Pattern Set.
	Name pulumi.StringOutput `pulumi:"name"`

	// A list of regular expression (regex) patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`.
	RegexPatternStrings pulumi.StringArrayOutput `pulumi:"regexPatternStrings"`
}

// NewRegexPatternSet registers a new resource with the given unique name, arguments, and options.
func NewRegexPatternSet(ctx *pulumi.Context,
	name string, args *RegexPatternSetArgs, opts ...pulumi.ResourceOption) (*RegexPatternSet, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.RegexPatternStrings; i != nil { inputs["regexPatternStrings"] = i.ToStringArrayOutput() }
	}
	var resource RegexPatternSet
	err := ctx.RegisterResource("aws:waf/regexPatternSet:RegexPatternSet", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegexPatternSet gets an existing RegexPatternSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegexPatternSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegexPatternSetState, opts ...pulumi.ResourceOption) (*RegexPatternSet, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Arn; i != nil { inputs["arn"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.RegexPatternStrings; i != nil { inputs["regexPatternStrings"] = i.ToStringArrayOutput() }
	}
	var resource RegexPatternSet
	err := ctx.ReadResource("aws:waf/regexPatternSet:RegexPatternSet", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegexPatternSet resources.
type RegexPatternSetState struct {
	// Amazon Resource Name (ARN)
	Arn pulumi.StringInput `pulumi:"arn"`
	// The name or description of the Regex Pattern Set.
	Name pulumi.StringInput `pulumi:"name"`
	// A list of regular expression (regex) patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`.
	RegexPatternStrings pulumi.StringArrayInput `pulumi:"regexPatternStrings"`
}

// The set of arguments for constructing a RegexPatternSet resource.
type RegexPatternSetArgs struct {
	// The name or description of the Regex Pattern Set.
	Name pulumi.StringInput `pulumi:"name"`
	// A list of regular expression (regex) patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`.
	RegexPatternStrings pulumi.StringArrayInput `pulumi:"regexPatternStrings"`
}
