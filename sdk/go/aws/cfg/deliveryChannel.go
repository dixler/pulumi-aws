// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cfg

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an AWS Config Delivery Channel.
// 
// > **Note:** Delivery Channel requires a [Configuration Recorder](https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder.html) to be present. Use of `dependsOn` (as shown below) is recommended to avoid race conditions.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/config_delivery_channel.html.markdown.
type DeliveryChannel struct {
	pulumi.CustomResourceState

	// The name of the delivery channel. Defaults to `default`. Changing it recreates the resource.
	Name pulumi.StringOutput `pulumi:"name"`

	// The name of the S3 bucket used to store the configuration history.
	S3BucketName pulumi.StringOutput `pulumi:"s3BucketName"`

	// The prefix for the specified S3 bucket.
	S3KeyPrefix pulumi.StringOutput `pulumi:"s3KeyPrefix"`

	// Options for how AWS Config delivers configuration snapshots. See below
	SnapshotDeliveryProperties DeliveryChannelSnapshotDeliveryPropertiesOutput `pulumi:"snapshotDeliveryProperties"`

	// The ARN of the SNS topic that AWS Config delivers notifications to.
	SnsTopicArn pulumi.StringOutput `pulumi:"snsTopicArn"`
}

// NewDeliveryChannel registers a new resource with the given unique name, arguments, and options.
func NewDeliveryChannel(ctx *pulumi.Context,
	name string, args *DeliveryChannelArgs, opts ...pulumi.ResourceOption) (*DeliveryChannel, error) {
	if args == nil || args.S3BucketName == nil {
		return nil, errors.New("missing required argument 'S3BucketName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.S3BucketName; i != nil { inputs["s3BucketName"] = i.ToStringOutput() }
		if i := args.S3KeyPrefix; i != nil { inputs["s3KeyPrefix"] = i.ToStringOutput() }
		if i := args.SnapshotDeliveryProperties; i != nil { inputs["snapshotDeliveryProperties"] = i.ToDeliveryChannelSnapshotDeliveryPropertiesOutput() }
		if i := args.SnsTopicArn; i != nil { inputs["snsTopicArn"] = i.ToStringOutput() }
	}
	var resource DeliveryChannel
	err := ctx.RegisterResource("aws:cfg/deliveryChannel:DeliveryChannel", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeliveryChannel gets an existing DeliveryChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeliveryChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeliveryChannelState, opts ...pulumi.ResourceOption) (*DeliveryChannel, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.S3BucketName; i != nil { inputs["s3BucketName"] = i.ToStringOutput() }
		if i := state.S3KeyPrefix; i != nil { inputs["s3KeyPrefix"] = i.ToStringOutput() }
		if i := state.SnapshotDeliveryProperties; i != nil { inputs["snapshotDeliveryProperties"] = i.ToDeliveryChannelSnapshotDeliveryPropertiesOutput() }
		if i := state.SnsTopicArn; i != nil { inputs["snsTopicArn"] = i.ToStringOutput() }
	}
	var resource DeliveryChannel
	err := ctx.ReadResource("aws:cfg/deliveryChannel:DeliveryChannel", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeliveryChannel resources.
type DeliveryChannelState struct {
	// The name of the delivery channel. Defaults to `default`. Changing it recreates the resource.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the S3 bucket used to store the configuration history.
	S3BucketName pulumi.StringInput `pulumi:"s3BucketName"`
	// The prefix for the specified S3 bucket.
	S3KeyPrefix pulumi.StringInput `pulumi:"s3KeyPrefix"`
	// Options for how AWS Config delivers configuration snapshots. See below
	SnapshotDeliveryProperties DeliveryChannelSnapshotDeliveryPropertiesInput `pulumi:"snapshotDeliveryProperties"`
	// The ARN of the SNS topic that AWS Config delivers notifications to.
	SnsTopicArn pulumi.StringInput `pulumi:"snsTopicArn"`
}

// The set of arguments for constructing a DeliveryChannel resource.
type DeliveryChannelArgs struct {
	// The name of the delivery channel. Defaults to `default`. Changing it recreates the resource.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the S3 bucket used to store the configuration history.
	S3BucketName pulumi.StringInput `pulumi:"s3BucketName"`
	// The prefix for the specified S3 bucket.
	S3KeyPrefix pulumi.StringInput `pulumi:"s3KeyPrefix"`
	// Options for how AWS Config delivers configuration snapshots. See below
	SnapshotDeliveryProperties DeliveryChannelSnapshotDeliveryPropertiesInput `pulumi:"snapshotDeliveryProperties"`
	// The ARN of the SNS topic that AWS Config delivers notifications to.
	SnsTopicArn pulumi.StringInput `pulumi:"snsTopicArn"`
}
type DeliveryChannelSnapshotDeliveryProperties struct {
	// - The frequency with which AWS Config recurringly delivers configuration snapshots.
	// e.g. `One_Hour` or `Three_Hours`.
	// Valid values are listed [here](https://docs.aws.amazon.com/config/latest/APIReference/API_ConfigSnapshotDeliveryProperties.html#API_ConfigSnapshotDeliveryProperties_Contents).
	DeliveryFrequency *string `pulumi:"deliveryFrequency"`
}
var deliveryChannelSnapshotDeliveryPropertiesType = reflect.TypeOf((*DeliveryChannelSnapshotDeliveryProperties)(nil)).Elem()

type DeliveryChannelSnapshotDeliveryPropertiesInput interface {
	pulumi.Input

	ToDeliveryChannelSnapshotDeliveryPropertiesOutput() DeliveryChannelSnapshotDeliveryPropertiesOutput
	ToDeliveryChannelSnapshotDeliveryPropertiesOutputWithContext(ctx context.Context) DeliveryChannelSnapshotDeliveryPropertiesOutput
}

type DeliveryChannelSnapshotDeliveryPropertiesArgs struct {
	// - The frequency with which AWS Config recurringly delivers configuration snapshots.
	// e.g. `One_Hour` or `Three_Hours`.
	// Valid values are listed [here](https://docs.aws.amazon.com/config/latest/APIReference/API_ConfigSnapshotDeliveryProperties.html#API_ConfigSnapshotDeliveryProperties_Contents).
	DeliveryFrequency pulumi.StringInput `pulumi:"deliveryFrequency"`
}

func (DeliveryChannelSnapshotDeliveryPropertiesArgs) ElementType() reflect.Type {
	return deliveryChannelSnapshotDeliveryPropertiesType
}

func (a DeliveryChannelSnapshotDeliveryPropertiesArgs) ToDeliveryChannelSnapshotDeliveryPropertiesOutput() DeliveryChannelSnapshotDeliveryPropertiesOutput {
	return pulumi.ToOutput(a).(DeliveryChannelSnapshotDeliveryPropertiesOutput)
}

func (a DeliveryChannelSnapshotDeliveryPropertiesArgs) ToDeliveryChannelSnapshotDeliveryPropertiesOutputWithContext(ctx context.Context) DeliveryChannelSnapshotDeliveryPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, a).(DeliveryChannelSnapshotDeliveryPropertiesOutput)
}

type DeliveryChannelSnapshotDeliveryPropertiesOutput struct { *pulumi.OutputState }

// - The frequency with which AWS Config recurringly delivers configuration snapshots.
// e.g. `One_Hour` or `Three_Hours`.
// Valid values are listed [here](https://docs.aws.amazon.com/config/latest/APIReference/API_ConfigSnapshotDeliveryProperties.html#API_ConfigSnapshotDeliveryProperties_Contents).
func (o DeliveryChannelSnapshotDeliveryPropertiesOutput) DeliveryFrequency() pulumi.StringOutput {
	return o.Apply(func(v DeliveryChannelSnapshotDeliveryProperties) string {
		if v.DeliveryFrequency == nil { return *new(string) } else { return *v.DeliveryFrequency }
	}).(pulumi.StringOutput)
}

func (DeliveryChannelSnapshotDeliveryPropertiesOutput) ElementType() reflect.Type {
	return deliveryChannelSnapshotDeliveryPropertiesType
}

func (o DeliveryChannelSnapshotDeliveryPropertiesOutput) ToDeliveryChannelSnapshotDeliveryPropertiesOutput() DeliveryChannelSnapshotDeliveryPropertiesOutput {
	return o
}

func (o DeliveryChannelSnapshotDeliveryPropertiesOutput) ToDeliveryChannelSnapshotDeliveryPropertiesOutputWithContext(ctx context.Context) DeliveryChannelSnapshotDeliveryPropertiesOutput {
	return o
}

func init() { pulumi.RegisterOutputType(DeliveryChannelSnapshotDeliveryPropertiesOutput{}) }

