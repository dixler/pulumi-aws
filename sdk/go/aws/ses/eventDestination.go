// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an SES event destination
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ses_event_destination.html.markdown.
type EventDestination struct {
	pulumi.CustomResourceState

	// CloudWatch destination for the events
	CloudwatchDestinations EventDestinationCloudwatchDestinationsArrayOutput `pulumi:"cloudwatchDestinations"`

	// The name of the configuration set
	ConfigurationSetName pulumi.StringOutput `pulumi:"configurationSetName"`

	// If true, the event destination will be enabled
	Enabled pulumi.BoolOutput `pulumi:"enabled"`

	// Send the events to a kinesis firehose destination
	KinesisDestination EventDestinationKinesisDestinationOutput `pulumi:"kinesisDestination"`

	// A list of matching types. May be any of `"send"`, `"reject"`, `"bounce"`, `"complaint"`, `"delivery"`, `"open"`, `"click"`, or `"renderingFailure"`.
	MatchingTypes pulumi.StringArrayOutput `pulumi:"matchingTypes"`

	// The name of the event destination
	Name pulumi.StringOutput `pulumi:"name"`

	// Send the events to an SNS Topic destination
	SnsDestination EventDestinationSnsDestinationOutput `pulumi:"snsDestination"`
}

// NewEventDestination registers a new resource with the given unique name, arguments, and options.
func NewEventDestination(ctx *pulumi.Context,
	name string, args *EventDestinationArgs, opts ...pulumi.ResourceOption) (*EventDestination, error) {
	if args == nil || args.ConfigurationSetName == nil {
		return nil, errors.New("missing required argument 'ConfigurationSetName'")
	}
	if args == nil || args.MatchingTypes == nil {
		return nil, errors.New("missing required argument 'MatchingTypes'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.CloudwatchDestinations; i != nil { inputs["cloudwatchDestinations"] = i.ToEventDestinationCloudwatchDestinationsArrayOutput() }
		if i := args.ConfigurationSetName; i != nil { inputs["configurationSetName"] = i.ToStringOutput() }
		if i := args.Enabled; i != nil { inputs["enabled"] = i.ToBoolOutput() }
		if i := args.KinesisDestination; i != nil { inputs["kinesisDestination"] = i.ToEventDestinationKinesisDestinationOutput() }
		if i := args.MatchingTypes; i != nil { inputs["matchingTypes"] = i.ToStringArrayOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.SnsDestination; i != nil { inputs["snsDestination"] = i.ToEventDestinationSnsDestinationOutput() }
	}
	var resource EventDestination
	err := ctx.RegisterResource("aws:ses/eventDestination:EventDestination", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventDestination gets an existing EventDestination resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventDestination(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventDestinationState, opts ...pulumi.ResourceOption) (*EventDestination, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.CloudwatchDestinations; i != nil { inputs["cloudwatchDestinations"] = i.ToEventDestinationCloudwatchDestinationsArrayOutput() }
		if i := state.ConfigurationSetName; i != nil { inputs["configurationSetName"] = i.ToStringOutput() }
		if i := state.Enabled; i != nil { inputs["enabled"] = i.ToBoolOutput() }
		if i := state.KinesisDestination; i != nil { inputs["kinesisDestination"] = i.ToEventDestinationKinesisDestinationOutput() }
		if i := state.MatchingTypes; i != nil { inputs["matchingTypes"] = i.ToStringArrayOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.SnsDestination; i != nil { inputs["snsDestination"] = i.ToEventDestinationSnsDestinationOutput() }
	}
	var resource EventDestination
	err := ctx.ReadResource("aws:ses/eventDestination:EventDestination", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventDestination resources.
type EventDestinationState struct {
	// CloudWatch destination for the events
	CloudwatchDestinations EventDestinationCloudwatchDestinationsArrayInput `pulumi:"cloudwatchDestinations"`
	// The name of the configuration set
	ConfigurationSetName pulumi.StringInput `pulumi:"configurationSetName"`
	// If true, the event destination will be enabled
	Enabled pulumi.BoolInput `pulumi:"enabled"`
	// Send the events to a kinesis firehose destination
	KinesisDestination EventDestinationKinesisDestinationInput `pulumi:"kinesisDestination"`
	// A list of matching types. May be any of `"send"`, `"reject"`, `"bounce"`, `"complaint"`, `"delivery"`, `"open"`, `"click"`, or `"renderingFailure"`.
	MatchingTypes pulumi.StringArrayInput `pulumi:"matchingTypes"`
	// The name of the event destination
	Name pulumi.StringInput `pulumi:"name"`
	// Send the events to an SNS Topic destination
	SnsDestination EventDestinationSnsDestinationInput `pulumi:"snsDestination"`
}

// The set of arguments for constructing a EventDestination resource.
type EventDestinationArgs struct {
	// CloudWatch destination for the events
	CloudwatchDestinations EventDestinationCloudwatchDestinationsArrayInput `pulumi:"cloudwatchDestinations"`
	// The name of the configuration set
	ConfigurationSetName pulumi.StringInput `pulumi:"configurationSetName"`
	// If true, the event destination will be enabled
	Enabled pulumi.BoolInput `pulumi:"enabled"`
	// Send the events to a kinesis firehose destination
	KinesisDestination EventDestinationKinesisDestinationInput `pulumi:"kinesisDestination"`
	// A list of matching types. May be any of `"send"`, `"reject"`, `"bounce"`, `"complaint"`, `"delivery"`, `"open"`, `"click"`, or `"renderingFailure"`.
	MatchingTypes pulumi.StringArrayInput `pulumi:"matchingTypes"`
	// The name of the event destination
	Name pulumi.StringInput `pulumi:"name"`
	// Send the events to an SNS Topic destination
	SnsDestination EventDestinationSnsDestinationInput `pulumi:"snsDestination"`
}
type EventDestinationCloudwatchDestinations struct {
	// The default value for the event
	DefaultValue string `pulumi:"defaultValue"`
	// The name for the dimension
	DimensionName string `pulumi:"dimensionName"`
	// The source for the value. It can be either `"messageTag"` or `"emailHeader"`
	ValueSource string `pulumi:"valueSource"`
}
var eventDestinationCloudwatchDestinationsType = reflect.TypeOf((*EventDestinationCloudwatchDestinations)(nil)).Elem()

type EventDestinationCloudwatchDestinationsInput interface {
	pulumi.Input

	ToEventDestinationCloudwatchDestinationsOutput() EventDestinationCloudwatchDestinationsOutput
	ToEventDestinationCloudwatchDestinationsOutputWithContext(ctx context.Context) EventDestinationCloudwatchDestinationsOutput
}

type EventDestinationCloudwatchDestinationsArgs struct {
	// The default value for the event
	DefaultValue pulumi.StringInput `pulumi:"defaultValue"`
	// The name for the dimension
	DimensionName pulumi.StringInput `pulumi:"dimensionName"`
	// The source for the value. It can be either `"messageTag"` or `"emailHeader"`
	ValueSource pulumi.StringInput `pulumi:"valueSource"`
}

func (EventDestinationCloudwatchDestinationsArgs) ElementType() reflect.Type {
	return eventDestinationCloudwatchDestinationsType
}

func (a EventDestinationCloudwatchDestinationsArgs) ToEventDestinationCloudwatchDestinationsOutput() EventDestinationCloudwatchDestinationsOutput {
	return pulumi.ToOutput(a).(EventDestinationCloudwatchDestinationsOutput)
}

func (a EventDestinationCloudwatchDestinationsArgs) ToEventDestinationCloudwatchDestinationsOutputWithContext(ctx context.Context) EventDestinationCloudwatchDestinationsOutput {
	return pulumi.ToOutputWithContext(ctx, a).(EventDestinationCloudwatchDestinationsOutput)
}

type EventDestinationCloudwatchDestinationsOutput struct { *pulumi.OutputState }

// The default value for the event
func (o EventDestinationCloudwatchDestinationsOutput) DefaultValue() pulumi.StringOutput {
	return o.Apply(func(v EventDestinationCloudwatchDestinations) string {
		return v.DefaultValue
	}).(pulumi.StringOutput)
}

// The name for the dimension
func (o EventDestinationCloudwatchDestinationsOutput) DimensionName() pulumi.StringOutput {
	return o.Apply(func(v EventDestinationCloudwatchDestinations) string {
		return v.DimensionName
	}).(pulumi.StringOutput)
}

// The source for the value. It can be either `"messageTag"` or `"emailHeader"`
func (o EventDestinationCloudwatchDestinationsOutput) ValueSource() pulumi.StringOutput {
	return o.Apply(func(v EventDestinationCloudwatchDestinations) string {
		return v.ValueSource
	}).(pulumi.StringOutput)
}

func (EventDestinationCloudwatchDestinationsOutput) ElementType() reflect.Type {
	return eventDestinationCloudwatchDestinationsType
}

func (o EventDestinationCloudwatchDestinationsOutput) ToEventDestinationCloudwatchDestinationsOutput() EventDestinationCloudwatchDestinationsOutput {
	return o
}

func (o EventDestinationCloudwatchDestinationsOutput) ToEventDestinationCloudwatchDestinationsOutputWithContext(ctx context.Context) EventDestinationCloudwatchDestinationsOutput {
	return o
}

func init() { pulumi.RegisterOutputType(EventDestinationCloudwatchDestinationsOutput{}) }

var eventDestinationCloudwatchDestinationsArrayType = reflect.TypeOf((*[]EventDestinationCloudwatchDestinations)(nil)).Elem()

type EventDestinationCloudwatchDestinationsArrayInput interface {
	pulumi.Input

	ToEventDestinationCloudwatchDestinationsArrayOutput() EventDestinationCloudwatchDestinationsArrayOutput
	ToEventDestinationCloudwatchDestinationsArrayOutputWithContext(ctx context.Context) EventDestinationCloudwatchDestinationsArrayOutput
}

type EventDestinationCloudwatchDestinationsArrayArgs []EventDestinationCloudwatchDestinationsInput

func (EventDestinationCloudwatchDestinationsArrayArgs) ElementType() reflect.Type {
	return eventDestinationCloudwatchDestinationsArrayType
}

func (a EventDestinationCloudwatchDestinationsArrayArgs) ToEventDestinationCloudwatchDestinationsArrayOutput() EventDestinationCloudwatchDestinationsArrayOutput {
	return pulumi.ToOutput(a).(EventDestinationCloudwatchDestinationsArrayOutput)
}

func (a EventDestinationCloudwatchDestinationsArrayArgs) ToEventDestinationCloudwatchDestinationsArrayOutputWithContext(ctx context.Context) EventDestinationCloudwatchDestinationsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(EventDestinationCloudwatchDestinationsArrayOutput)
}

type EventDestinationCloudwatchDestinationsArrayOutput struct { *pulumi.OutputState }

func (o EventDestinationCloudwatchDestinationsArrayOutput) Index(i pulumi.IntInput) EventDestinationCloudwatchDestinationsOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) EventDestinationCloudwatchDestinations {
		return vs[0].([]EventDestinationCloudwatchDestinations)[vs[1].(int)]
	}).(EventDestinationCloudwatchDestinationsOutput)
}

func (EventDestinationCloudwatchDestinationsArrayOutput) ElementType() reflect.Type {
	return eventDestinationCloudwatchDestinationsArrayType
}

func (o EventDestinationCloudwatchDestinationsArrayOutput) ToEventDestinationCloudwatchDestinationsArrayOutput() EventDestinationCloudwatchDestinationsArrayOutput {
	return o
}

func (o EventDestinationCloudwatchDestinationsArrayOutput) ToEventDestinationCloudwatchDestinationsArrayOutputWithContext(ctx context.Context) EventDestinationCloudwatchDestinationsArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(EventDestinationCloudwatchDestinationsArrayOutput{}) }

type EventDestinationKinesisDestination struct {
	// The ARN of the role that has permissions to access the Kinesis Stream
	RoleArn string `pulumi:"roleArn"`
	// The ARN of the Kinesis Stream
	StreamArn string `pulumi:"streamArn"`
}
var eventDestinationKinesisDestinationType = reflect.TypeOf((*EventDestinationKinesisDestination)(nil)).Elem()

type EventDestinationKinesisDestinationInput interface {
	pulumi.Input

	ToEventDestinationKinesisDestinationOutput() EventDestinationKinesisDestinationOutput
	ToEventDestinationKinesisDestinationOutputWithContext(ctx context.Context) EventDestinationKinesisDestinationOutput
}

type EventDestinationKinesisDestinationArgs struct {
	// The ARN of the role that has permissions to access the Kinesis Stream
	RoleArn pulumi.StringInput `pulumi:"roleArn"`
	// The ARN of the Kinesis Stream
	StreamArn pulumi.StringInput `pulumi:"streamArn"`
}

func (EventDestinationKinesisDestinationArgs) ElementType() reflect.Type {
	return eventDestinationKinesisDestinationType
}

func (a EventDestinationKinesisDestinationArgs) ToEventDestinationKinesisDestinationOutput() EventDestinationKinesisDestinationOutput {
	return pulumi.ToOutput(a).(EventDestinationKinesisDestinationOutput)
}

func (a EventDestinationKinesisDestinationArgs) ToEventDestinationKinesisDestinationOutputWithContext(ctx context.Context) EventDestinationKinesisDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, a).(EventDestinationKinesisDestinationOutput)
}

type EventDestinationKinesisDestinationOutput struct { *pulumi.OutputState }

// The ARN of the role that has permissions to access the Kinesis Stream
func (o EventDestinationKinesisDestinationOutput) RoleArn() pulumi.StringOutput {
	return o.Apply(func(v EventDestinationKinesisDestination) string {
		return v.RoleArn
	}).(pulumi.StringOutput)
}

// The ARN of the Kinesis Stream
func (o EventDestinationKinesisDestinationOutput) StreamArn() pulumi.StringOutput {
	return o.Apply(func(v EventDestinationKinesisDestination) string {
		return v.StreamArn
	}).(pulumi.StringOutput)
}

func (EventDestinationKinesisDestinationOutput) ElementType() reflect.Type {
	return eventDestinationKinesisDestinationType
}

func (o EventDestinationKinesisDestinationOutput) ToEventDestinationKinesisDestinationOutput() EventDestinationKinesisDestinationOutput {
	return o
}

func (o EventDestinationKinesisDestinationOutput) ToEventDestinationKinesisDestinationOutputWithContext(ctx context.Context) EventDestinationKinesisDestinationOutput {
	return o
}

func init() { pulumi.RegisterOutputType(EventDestinationKinesisDestinationOutput{}) }

type EventDestinationSnsDestination struct {
	// The ARN of the SNS topic
	TopicArn string `pulumi:"topicArn"`
}
var eventDestinationSnsDestinationType = reflect.TypeOf((*EventDestinationSnsDestination)(nil)).Elem()

type EventDestinationSnsDestinationInput interface {
	pulumi.Input

	ToEventDestinationSnsDestinationOutput() EventDestinationSnsDestinationOutput
	ToEventDestinationSnsDestinationOutputWithContext(ctx context.Context) EventDestinationSnsDestinationOutput
}

type EventDestinationSnsDestinationArgs struct {
	// The ARN of the SNS topic
	TopicArn pulumi.StringInput `pulumi:"topicArn"`
}

func (EventDestinationSnsDestinationArgs) ElementType() reflect.Type {
	return eventDestinationSnsDestinationType
}

func (a EventDestinationSnsDestinationArgs) ToEventDestinationSnsDestinationOutput() EventDestinationSnsDestinationOutput {
	return pulumi.ToOutput(a).(EventDestinationSnsDestinationOutput)
}

func (a EventDestinationSnsDestinationArgs) ToEventDestinationSnsDestinationOutputWithContext(ctx context.Context) EventDestinationSnsDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, a).(EventDestinationSnsDestinationOutput)
}

type EventDestinationSnsDestinationOutput struct { *pulumi.OutputState }

// The ARN of the SNS topic
func (o EventDestinationSnsDestinationOutput) TopicArn() pulumi.StringOutput {
	return o.Apply(func(v EventDestinationSnsDestination) string {
		return v.TopicArn
	}).(pulumi.StringOutput)
}

func (EventDestinationSnsDestinationOutput) ElementType() reflect.Type {
	return eventDestinationSnsDestinationType
}

func (o EventDestinationSnsDestinationOutput) ToEventDestinationSnsDestinationOutput() EventDestinationSnsDestinationOutput {
	return o
}

func (o EventDestinationSnsDestinationOutput) ToEventDestinationSnsDestinationOutputWithContext(ctx context.Context) EventDestinationSnsDestinationOutput {
	return o
}

func init() { pulumi.RegisterOutputType(EventDestinationSnsDestinationOutput{}) }

