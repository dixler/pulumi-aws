// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codepipeline

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a CodePipeline Webhook.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/codepipeline_webhook.html.markdown.
type Webhook struct {
	pulumi.CustomResourceState

	// The type of authentication  to use. One of `IP`, `GITHUB_HMAC`, or `UNAUTHENTICATED`.
	Authentication pulumi.StringOutput `pulumi:"authentication"`

	// An `auth` block. Required for `IP` and `GITHUB_HMAC`. Auth blocks are documented below.
	AuthenticationConfiguration WebhookAuthenticationConfigurationOutput `pulumi:"authenticationConfiguration"`

	// One or more `filter` blocks. Filter blocks are documented below.
	Filters WebhookFiltersArrayOutput `pulumi:"filters"`

	// The name of the webhook.
	Name pulumi.StringOutput `pulumi:"name"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// The name of the action in a pipeline you want to connect to the webhook. The action must be from the source (first) stage of the pipeline.
	TargetAction pulumi.StringOutput `pulumi:"targetAction"`

	// The name of the pipeline.
	TargetPipeline pulumi.StringOutput `pulumi:"targetPipeline"`

	// The CodePipeline webhook's URL. POST events to this endpoint to trigger the target.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewWebhook registers a new resource with the given unique name, arguments, and options.
func NewWebhook(ctx *pulumi.Context,
	name string, args *WebhookArgs, opts ...pulumi.ResourceOption) (*Webhook, error) {
	if args == nil || args.Authentication == nil {
		return nil, errors.New("missing required argument 'Authentication'")
	}
	if args == nil || args.Filters == nil {
		return nil, errors.New("missing required argument 'Filters'")
	}
	if args == nil || args.TargetAction == nil {
		return nil, errors.New("missing required argument 'TargetAction'")
	}
	if args == nil || args.TargetPipeline == nil {
		return nil, errors.New("missing required argument 'TargetPipeline'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Authentication; i != nil { inputs["authentication"] = i.ToStringOutput() }
		if i := args.AuthenticationConfiguration; i != nil { inputs["authenticationConfiguration"] = i.ToWebhookAuthenticationConfigurationOutput() }
		if i := args.Filters; i != nil { inputs["filters"] = i.ToWebhookFiltersArrayOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := args.TargetAction; i != nil { inputs["targetAction"] = i.ToStringOutput() }
		if i := args.TargetPipeline; i != nil { inputs["targetPipeline"] = i.ToStringOutput() }
	}
	var resource Webhook
	err := ctx.RegisterResource("aws:codepipeline/webhook:Webhook", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebhook gets an existing Webhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebhookState, opts ...pulumi.ResourceOption) (*Webhook, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Authentication; i != nil { inputs["authentication"] = i.ToStringOutput() }
		if i := state.AuthenticationConfiguration; i != nil { inputs["authenticationConfiguration"] = i.ToWebhookAuthenticationConfigurationOutput() }
		if i := state.Filters; i != nil { inputs["filters"] = i.ToWebhookFiltersArrayOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := state.TargetAction; i != nil { inputs["targetAction"] = i.ToStringOutput() }
		if i := state.TargetPipeline; i != nil { inputs["targetPipeline"] = i.ToStringOutput() }
		if i := state.Url; i != nil { inputs["url"] = i.ToStringOutput() }
	}
	var resource Webhook
	err := ctx.ReadResource("aws:codepipeline/webhook:Webhook", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Webhook resources.
type WebhookState struct {
	// The type of authentication  to use. One of `IP`, `GITHUB_HMAC`, or `UNAUTHENTICATED`.
	Authentication pulumi.StringInput `pulumi:"authentication"`
	// An `auth` block. Required for `IP` and `GITHUB_HMAC`. Auth blocks are documented below.
	AuthenticationConfiguration WebhookAuthenticationConfigurationInput `pulumi:"authenticationConfiguration"`
	// One or more `filter` blocks. Filter blocks are documented below.
	Filters WebhookFiltersArrayInput `pulumi:"filters"`
	// The name of the webhook.
	Name pulumi.StringInput `pulumi:"name"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The name of the action in a pipeline you want to connect to the webhook. The action must be from the source (first) stage of the pipeline.
	TargetAction pulumi.StringInput `pulumi:"targetAction"`
	// The name of the pipeline.
	TargetPipeline pulumi.StringInput `pulumi:"targetPipeline"`
	// The CodePipeline webhook's URL. POST events to this endpoint to trigger the target.
	Url pulumi.StringInput `pulumi:"url"`
}

// The set of arguments for constructing a Webhook resource.
type WebhookArgs struct {
	// The type of authentication  to use. One of `IP`, `GITHUB_HMAC`, or `UNAUTHENTICATED`.
	Authentication pulumi.StringInput `pulumi:"authentication"`
	// An `auth` block. Required for `IP` and `GITHUB_HMAC`. Auth blocks are documented below.
	AuthenticationConfiguration WebhookAuthenticationConfigurationInput `pulumi:"authenticationConfiguration"`
	// One or more `filter` blocks. Filter blocks are documented below.
	Filters WebhookFiltersArrayInput `pulumi:"filters"`
	// The name of the webhook.
	Name pulumi.StringInput `pulumi:"name"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The name of the action in a pipeline you want to connect to the webhook. The action must be from the source (first) stage of the pipeline.
	TargetAction pulumi.StringInput `pulumi:"targetAction"`
	// The name of the pipeline.
	TargetPipeline pulumi.StringInput `pulumi:"targetPipeline"`
}
type WebhookAuthenticationConfiguration struct {
	AllowedIpRange *string `pulumi:"allowedIpRange"`
	SecretToken *string `pulumi:"secretToken"`
}
var webhookAuthenticationConfigurationType = reflect.TypeOf((*WebhookAuthenticationConfiguration)(nil)).Elem()

type WebhookAuthenticationConfigurationInput interface {
	pulumi.Input

	ToWebhookAuthenticationConfigurationOutput() WebhookAuthenticationConfigurationOutput
	ToWebhookAuthenticationConfigurationOutputWithContext(ctx context.Context) WebhookAuthenticationConfigurationOutput
}

type WebhookAuthenticationConfigurationArgs struct {
	AllowedIpRange pulumi.StringInput `pulumi:"allowedIpRange"`
	SecretToken pulumi.StringInput `pulumi:"secretToken"`
}

func (WebhookAuthenticationConfigurationArgs) ElementType() reflect.Type {
	return webhookAuthenticationConfigurationType
}

func (a WebhookAuthenticationConfigurationArgs) ToWebhookAuthenticationConfigurationOutput() WebhookAuthenticationConfigurationOutput {
	return pulumi.ToOutput(a).(WebhookAuthenticationConfigurationOutput)
}

func (a WebhookAuthenticationConfigurationArgs) ToWebhookAuthenticationConfigurationOutputWithContext(ctx context.Context) WebhookAuthenticationConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, a).(WebhookAuthenticationConfigurationOutput)
}

type WebhookAuthenticationConfigurationOutput struct { *pulumi.OutputState }

func (o WebhookAuthenticationConfigurationOutput) AllowedIpRange() pulumi.StringOutput {
	return o.Apply(func(v WebhookAuthenticationConfiguration) string {
		if v.AllowedIpRange == nil { return *new(string) } else { return *v.AllowedIpRange }
	}).(pulumi.StringOutput)
}

func (o WebhookAuthenticationConfigurationOutput) SecretToken() pulumi.StringOutput {
	return o.Apply(func(v WebhookAuthenticationConfiguration) string {
		if v.SecretToken == nil { return *new(string) } else { return *v.SecretToken }
	}).(pulumi.StringOutput)
}

func (WebhookAuthenticationConfigurationOutput) ElementType() reflect.Type {
	return webhookAuthenticationConfigurationType
}

func (o WebhookAuthenticationConfigurationOutput) ToWebhookAuthenticationConfigurationOutput() WebhookAuthenticationConfigurationOutput {
	return o
}

func (o WebhookAuthenticationConfigurationOutput) ToWebhookAuthenticationConfigurationOutputWithContext(ctx context.Context) WebhookAuthenticationConfigurationOutput {
	return o
}

func init() { pulumi.RegisterOutputType(WebhookAuthenticationConfigurationOutput{}) }

type WebhookFilters struct {
	JsonPath string `pulumi:"jsonPath"`
	MatchEquals string `pulumi:"matchEquals"`
}
var webhookFiltersType = reflect.TypeOf((*WebhookFilters)(nil)).Elem()

type WebhookFiltersInput interface {
	pulumi.Input

	ToWebhookFiltersOutput() WebhookFiltersOutput
	ToWebhookFiltersOutputWithContext(ctx context.Context) WebhookFiltersOutput
}

type WebhookFiltersArgs struct {
	JsonPath pulumi.StringInput `pulumi:"jsonPath"`
	MatchEquals pulumi.StringInput `pulumi:"matchEquals"`
}

func (WebhookFiltersArgs) ElementType() reflect.Type {
	return webhookFiltersType
}

func (a WebhookFiltersArgs) ToWebhookFiltersOutput() WebhookFiltersOutput {
	return pulumi.ToOutput(a).(WebhookFiltersOutput)
}

func (a WebhookFiltersArgs) ToWebhookFiltersOutputWithContext(ctx context.Context) WebhookFiltersOutput {
	return pulumi.ToOutputWithContext(ctx, a).(WebhookFiltersOutput)
}

type WebhookFiltersOutput struct { *pulumi.OutputState }

func (o WebhookFiltersOutput) JsonPath() pulumi.StringOutput {
	return o.Apply(func(v WebhookFilters) string {
		return v.JsonPath
	}).(pulumi.StringOutput)
}

func (o WebhookFiltersOutput) MatchEquals() pulumi.StringOutput {
	return o.Apply(func(v WebhookFilters) string {
		return v.MatchEquals
	}).(pulumi.StringOutput)
}

func (WebhookFiltersOutput) ElementType() reflect.Type {
	return webhookFiltersType
}

func (o WebhookFiltersOutput) ToWebhookFiltersOutput() WebhookFiltersOutput {
	return o
}

func (o WebhookFiltersOutput) ToWebhookFiltersOutputWithContext(ctx context.Context) WebhookFiltersOutput {
	return o
}

func init() { pulumi.RegisterOutputType(WebhookFiltersOutput{}) }

var webhookFiltersArrayType = reflect.TypeOf((*[]WebhookFilters)(nil)).Elem()

type WebhookFiltersArrayInput interface {
	pulumi.Input

	ToWebhookFiltersArrayOutput() WebhookFiltersArrayOutput
	ToWebhookFiltersArrayOutputWithContext(ctx context.Context) WebhookFiltersArrayOutput
}

type WebhookFiltersArrayArgs []WebhookFiltersInput

func (WebhookFiltersArrayArgs) ElementType() reflect.Type {
	return webhookFiltersArrayType
}

func (a WebhookFiltersArrayArgs) ToWebhookFiltersArrayOutput() WebhookFiltersArrayOutput {
	return pulumi.ToOutput(a).(WebhookFiltersArrayOutput)
}

func (a WebhookFiltersArrayArgs) ToWebhookFiltersArrayOutputWithContext(ctx context.Context) WebhookFiltersArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(WebhookFiltersArrayOutput)
}

type WebhookFiltersArrayOutput struct { *pulumi.OutputState }

func (o WebhookFiltersArrayOutput) Index(i pulumi.IntInput) WebhookFiltersOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) WebhookFilters {
		return vs[0].([]WebhookFilters)[vs[1].(int)]
	}).(WebhookFiltersOutput)
}

func (WebhookFiltersArrayOutput) ElementType() reflect.Type {
	return webhookFiltersArrayType
}

func (o WebhookFiltersArrayOutput) ToWebhookFiltersArrayOutput() WebhookFiltersArrayOutput {
	return o
}

func (o WebhookFiltersArrayOutput) ToWebhookFiltersArrayOutputWithContext(ctx context.Context) WebhookFiltersArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(WebhookFiltersArrayOutput{}) }

