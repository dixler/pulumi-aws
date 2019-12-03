// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to create a new launch configuration, used for autoscaling groups.
// 
// ## Block devices
// 
// Each of the `*_block_device` attributes controls a portion of the AWS
// Launch Configuration's "Block Device Mapping". It's a good idea to familiarize yourself with [AWS's Block Device
// Mapping docs](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html)
// to understand the implications of using these attributes.
// 
// The `rootBlockDevice` mapping supports the following:
// 
// * `volumeType` - (Optional) The type of volume. Can be `"standard"`, `"gp2"`,
//   or `"io1"`. (Default: `"standard"`).
// * `volumeSize` - (Optional) The size of the volume in gigabytes.
// * `iops` - (Optional) The amount of provisioned
//   [IOPS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html).
//   This must be set with a `volumeType` of `"io1"`.
// * `deleteOnTermination` - (Optional) Whether the volume should be destroyed
//   on instance termination (Default: `true`).
// * `encrypted` - (Optional) Whether the volume should be encrypted or not. (Default: `false`).
// 
// Modifying any of the `rootBlockDevice` settings requires resource
// replacement.
// 
// Each `ebsBlockDevice` supports the following:
// 
// * `deviceName` - (Required) The name of the device to mount.
// * `snapshotId` - (Optional) The Snapshot ID to mount.
// * `volumeType` - (Optional) The type of volume. Can be `"standard"`, `"gp2"`,
//   or `"io1"`. (Default: `"standard"`).
// * `volumeSize` - (Optional) The size of the volume in gigabytes.
// * `iops` - (Optional) The amount of provisioned
//   [IOPS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html).
//   This must be set with a `volumeType` of `"io1"`.
// * `deleteOnTermination` - (Optional) Whether the volume should be destroyed
//   on instance termination (Default: `true`).
// * `encrypted` - (Optional) Whether the volume should be encrypted or not. Do not use this option if you are using `snapshotId` as the encrypted flag will be determined by the snapshot. (Default: `false`).
// 
// Modifying any `ebsBlockDevice` currently requires resource replacement.
// 
// Each `ephemeralBlockDevice` supports the following:
// 
// * `deviceName` - The name of the block device to mount on the instance.
// * `virtualName` - The [Instance Store Device
//   Name](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#InstanceStoreDeviceNames)
//   (e.g. `"ephemeral0"`)
// 
// Each AWS Instance type has a different set of Instance Store block devices
// available for attachment. AWS [publishes a
// list](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#StorageOnInstanceTypes)
// of which ephemeral devices are available on each type. The devices are always
// identified by the `virtualName` in the format `"ephemeral{0..N}"`.
// 
// > **NOTE:** Changes to `*_block_device` configuration of _existing_ resources
// cannot currently be detected by this provider. After updating to block device
// configuration, resource recreation can be manually triggered by using the
// [`taint` command](https://www.terraform.io/docs/commands/taint.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/launch_configuration.html.markdown.
type LaunchConfiguration struct {
	pulumi.CustomResourceState

	// Associate a public ip address with an instance in a VPC.
	AssociatePublicIpAddress pulumi.BoolOutput `pulumi:"associatePublicIpAddress"`

	// Additional EBS block devices to attach to the
	// instance.  See Block Devices below for details.
	EbsBlockDevices LaunchConfigurationEbsBlockDevicesArrayOutput `pulumi:"ebsBlockDevices"`

	// If true, the launched EC2 instance will be EBS-optimized.
	EbsOptimized pulumi.BoolOutput `pulumi:"ebsOptimized"`

	// Enables/disables detailed monitoring. This is enabled by default.
	EnableMonitoring pulumi.BoolOutput `pulumi:"enableMonitoring"`

	// Customize Ephemeral (also known as
	// "Instance Store") volumes on the instance. See Block Devices below for details.
	EphemeralBlockDevices LaunchConfigurationEphemeralBlockDevicesArrayOutput `pulumi:"ephemeralBlockDevices"`

	// The name attribute of the IAM instance profile to associate
	// with launched instances.
	IamInstanceProfile pulumi.StringOutput `pulumi:"iamInstanceProfile"`

	// The EC2 image ID to launch.
	ImageId pulumi.StringOutput `pulumi:"imageId"`

	// The size of instance to launch.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`

	// The key name that should be used for the instance.
	KeyName pulumi.StringOutput `pulumi:"keyName"`

	// The name of the launch configuration. If you leave
	// this blank, this provider will auto-generate a unique name.
	Name pulumi.StringOutput `pulumi:"name"`

	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`

	// The tenancy of the instance. Valid values are
	// `"default"` or `"dedicated"`, see [AWS's Create Launch Configuration](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_CreateLaunchConfiguration.html)
	// for more details
	PlacementTenancy pulumi.StringOutput `pulumi:"placementTenancy"`

	// Customize details about the root block
	// device of the instance. See Block Devices below for details.
	RootBlockDevice LaunchConfigurationRootBlockDeviceOutput `pulumi:"rootBlockDevice"`

	// A list of associated security group IDS.
	SecurityGroups pulumi.StringArrayOutput `pulumi:"securityGroups"`

	// The maximum price to use for reserving spot instances.
	SpotPrice pulumi.StringOutput `pulumi:"spotPrice"`

	// The user data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see `userDataBase64` instead.
	UserData pulumi.StringOutput `pulumi:"userData"`

	// Can be used instead of `userData` to pass base64-encoded binary data directly. Use this instead of `userData` whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption.
	UserDataBase64 pulumi.StringOutput `pulumi:"userDataBase64"`

	// The ID of a ClassicLink-enabled VPC. Only applies to EC2-Classic instances. (eg. `vpc-2730681a`)
	VpcClassicLinkId pulumi.StringOutput `pulumi:"vpcClassicLinkId"`

	// The IDs of one or more security groups for the specified ClassicLink-enabled VPC (eg. `sg-46ae3d11`).
	VpcClassicLinkSecurityGroups pulumi.StringArrayOutput `pulumi:"vpcClassicLinkSecurityGroups"`
}

// NewLaunchConfiguration registers a new resource with the given unique name, arguments, and options.
func NewLaunchConfiguration(ctx *pulumi.Context,
	name string, args *LaunchConfigurationArgs, opts ...pulumi.ResourceOption) (*LaunchConfiguration, error) {
	if args == nil || args.ImageId == nil {
		return nil, errors.New("missing required argument 'ImageId'")
	}
	if args == nil || args.InstanceType == nil {
		return nil, errors.New("missing required argument 'InstanceType'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.AssociatePublicIpAddress; i != nil { inputs["associatePublicIpAddress"] = i.ToBoolOutput() }
		if i := args.EbsBlockDevices; i != nil { inputs["ebsBlockDevices"] = i.ToLaunchConfigurationEbsBlockDevicesArrayOutput() }
		if i := args.EbsOptimized; i != nil { inputs["ebsOptimized"] = i.ToBoolOutput() }
		if i := args.EnableMonitoring; i != nil { inputs["enableMonitoring"] = i.ToBoolOutput() }
		if i := args.EphemeralBlockDevices; i != nil { inputs["ephemeralBlockDevices"] = i.ToLaunchConfigurationEphemeralBlockDevicesArrayOutput() }
		if i := args.IamInstanceProfile; i != nil { inputs["iamInstanceProfile"] = i.ToStringOutput() }
		if i := args.ImageId; i != nil { inputs["imageId"] = i.ToStringOutput() }
		if i := args.InstanceType; i != nil { inputs["instanceType"] = i.ToStringOutput() }
		if i := args.KeyName; i != nil { inputs["keyName"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.NamePrefix; i != nil { inputs["namePrefix"] = i.ToStringOutput() }
		if i := args.PlacementTenancy; i != nil { inputs["placementTenancy"] = i.ToStringOutput() }
		if i := args.RootBlockDevice; i != nil { inputs["rootBlockDevice"] = i.ToLaunchConfigurationRootBlockDeviceOutput() }
		if i := args.SecurityGroups; i != nil { inputs["securityGroups"] = i.ToStringArrayOutput() }
		if i := args.SpotPrice; i != nil { inputs["spotPrice"] = i.ToStringOutput() }
		if i := args.UserData; i != nil { inputs["userData"] = i.ToStringOutput() }
		if i := args.UserDataBase64; i != nil { inputs["userDataBase64"] = i.ToStringOutput() }
		if i := args.VpcClassicLinkId; i != nil { inputs["vpcClassicLinkId"] = i.ToStringOutput() }
		if i := args.VpcClassicLinkSecurityGroups; i != nil { inputs["vpcClassicLinkSecurityGroups"] = i.ToStringArrayOutput() }
	}
	var resource LaunchConfiguration
	err := ctx.RegisterResource("aws:ec2/launchConfiguration:LaunchConfiguration", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLaunchConfiguration gets an existing LaunchConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLaunchConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LaunchConfigurationState, opts ...pulumi.ResourceOption) (*LaunchConfiguration, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.AssociatePublicIpAddress; i != nil { inputs["associatePublicIpAddress"] = i.ToBoolOutput() }
		if i := state.EbsBlockDevices; i != nil { inputs["ebsBlockDevices"] = i.ToLaunchConfigurationEbsBlockDevicesArrayOutput() }
		if i := state.EbsOptimized; i != nil { inputs["ebsOptimized"] = i.ToBoolOutput() }
		if i := state.EnableMonitoring; i != nil { inputs["enableMonitoring"] = i.ToBoolOutput() }
		if i := state.EphemeralBlockDevices; i != nil { inputs["ephemeralBlockDevices"] = i.ToLaunchConfigurationEphemeralBlockDevicesArrayOutput() }
		if i := state.IamInstanceProfile; i != nil { inputs["iamInstanceProfile"] = i.ToStringOutput() }
		if i := state.ImageId; i != nil { inputs["imageId"] = i.ToStringOutput() }
		if i := state.InstanceType; i != nil { inputs["instanceType"] = i.ToStringOutput() }
		if i := state.KeyName; i != nil { inputs["keyName"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.NamePrefix; i != nil { inputs["namePrefix"] = i.ToStringOutput() }
		if i := state.PlacementTenancy; i != nil { inputs["placementTenancy"] = i.ToStringOutput() }
		if i := state.RootBlockDevice; i != nil { inputs["rootBlockDevice"] = i.ToLaunchConfigurationRootBlockDeviceOutput() }
		if i := state.SecurityGroups; i != nil { inputs["securityGroups"] = i.ToStringArrayOutput() }
		if i := state.SpotPrice; i != nil { inputs["spotPrice"] = i.ToStringOutput() }
		if i := state.UserData; i != nil { inputs["userData"] = i.ToStringOutput() }
		if i := state.UserDataBase64; i != nil { inputs["userDataBase64"] = i.ToStringOutput() }
		if i := state.VpcClassicLinkId; i != nil { inputs["vpcClassicLinkId"] = i.ToStringOutput() }
		if i := state.VpcClassicLinkSecurityGroups; i != nil { inputs["vpcClassicLinkSecurityGroups"] = i.ToStringArrayOutput() }
	}
	var resource LaunchConfiguration
	err := ctx.ReadResource("aws:ec2/launchConfiguration:LaunchConfiguration", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LaunchConfiguration resources.
type LaunchConfigurationState struct {
	// Associate a public ip address with an instance in a VPC.
	AssociatePublicIpAddress pulumi.BoolInput `pulumi:"associatePublicIpAddress"`
	// Additional EBS block devices to attach to the
	// instance.  See Block Devices below for details.
	EbsBlockDevices LaunchConfigurationEbsBlockDevicesArrayInput `pulumi:"ebsBlockDevices"`
	// If true, the launched EC2 instance will be EBS-optimized.
	EbsOptimized pulumi.BoolInput `pulumi:"ebsOptimized"`
	// Enables/disables detailed monitoring. This is enabled by default.
	EnableMonitoring pulumi.BoolInput `pulumi:"enableMonitoring"`
	// Customize Ephemeral (also known as
	// "Instance Store") volumes on the instance. See Block Devices below for details.
	EphemeralBlockDevices LaunchConfigurationEphemeralBlockDevicesArrayInput `pulumi:"ephemeralBlockDevices"`
	// The name attribute of the IAM instance profile to associate
	// with launched instances.
	IamInstanceProfile pulumi.StringInput `pulumi:"iamInstanceProfile"`
	// The EC2 image ID to launch.
	ImageId pulumi.StringInput `pulumi:"imageId"`
	// The size of instance to launch.
	InstanceType pulumi.StringInput `pulumi:"instanceType"`
	// The key name that should be used for the instance.
	KeyName pulumi.StringInput `pulumi:"keyName"`
	// The name of the launch configuration. If you leave
	// this blank, this provider will auto-generate a unique name.
	Name pulumi.StringInput `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringInput `pulumi:"namePrefix"`
	// The tenancy of the instance. Valid values are
	// `"default"` or `"dedicated"`, see [AWS's Create Launch Configuration](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_CreateLaunchConfiguration.html)
	// for more details
	PlacementTenancy pulumi.StringInput `pulumi:"placementTenancy"`
	// Customize details about the root block
	// device of the instance. See Block Devices below for details.
	RootBlockDevice LaunchConfigurationRootBlockDeviceInput `pulumi:"rootBlockDevice"`
	// A list of associated security group IDS.
	SecurityGroups pulumi.StringArrayInput `pulumi:"securityGroups"`
	// The maximum price to use for reserving spot instances.
	SpotPrice pulumi.StringInput `pulumi:"spotPrice"`
	// The user data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see `userDataBase64` instead.
	UserData pulumi.StringInput `pulumi:"userData"`
	// Can be used instead of `userData` to pass base64-encoded binary data directly. Use this instead of `userData` whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption.
	UserDataBase64 pulumi.StringInput `pulumi:"userDataBase64"`
	// The ID of a ClassicLink-enabled VPC. Only applies to EC2-Classic instances. (eg. `vpc-2730681a`)
	VpcClassicLinkId pulumi.StringInput `pulumi:"vpcClassicLinkId"`
	// The IDs of one or more security groups for the specified ClassicLink-enabled VPC (eg. `sg-46ae3d11`).
	VpcClassicLinkSecurityGroups pulumi.StringArrayInput `pulumi:"vpcClassicLinkSecurityGroups"`
}

// The set of arguments for constructing a LaunchConfiguration resource.
type LaunchConfigurationArgs struct {
	// Associate a public ip address with an instance in a VPC.
	AssociatePublicIpAddress pulumi.BoolInput `pulumi:"associatePublicIpAddress"`
	// Additional EBS block devices to attach to the
	// instance.  See Block Devices below for details.
	EbsBlockDevices LaunchConfigurationEbsBlockDevicesArrayInput `pulumi:"ebsBlockDevices"`
	// If true, the launched EC2 instance will be EBS-optimized.
	EbsOptimized pulumi.BoolInput `pulumi:"ebsOptimized"`
	// Enables/disables detailed monitoring. This is enabled by default.
	EnableMonitoring pulumi.BoolInput `pulumi:"enableMonitoring"`
	// Customize Ephemeral (also known as
	// "Instance Store") volumes on the instance. See Block Devices below for details.
	EphemeralBlockDevices LaunchConfigurationEphemeralBlockDevicesArrayInput `pulumi:"ephemeralBlockDevices"`
	// The name attribute of the IAM instance profile to associate
	// with launched instances.
	IamInstanceProfile pulumi.StringInput `pulumi:"iamInstanceProfile"`
	// The EC2 image ID to launch.
	ImageId pulumi.StringInput `pulumi:"imageId"`
	// The size of instance to launch.
	InstanceType pulumi.StringInput `pulumi:"instanceType"`
	// The key name that should be used for the instance.
	KeyName pulumi.StringInput `pulumi:"keyName"`
	// The name of the launch configuration. If you leave
	// this blank, this provider will auto-generate a unique name.
	Name pulumi.StringInput `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringInput `pulumi:"namePrefix"`
	// The tenancy of the instance. Valid values are
	// `"default"` or `"dedicated"`, see [AWS's Create Launch Configuration](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_CreateLaunchConfiguration.html)
	// for more details
	PlacementTenancy pulumi.StringInput `pulumi:"placementTenancy"`
	// Customize details about the root block
	// device of the instance. See Block Devices below for details.
	RootBlockDevice LaunchConfigurationRootBlockDeviceInput `pulumi:"rootBlockDevice"`
	// A list of associated security group IDS.
	SecurityGroups pulumi.StringArrayInput `pulumi:"securityGroups"`
	// The maximum price to use for reserving spot instances.
	SpotPrice pulumi.StringInput `pulumi:"spotPrice"`
	// The user data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see `userDataBase64` instead.
	UserData pulumi.StringInput `pulumi:"userData"`
	// Can be used instead of `userData` to pass base64-encoded binary data directly. Use this instead of `userData` whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption.
	UserDataBase64 pulumi.StringInput `pulumi:"userDataBase64"`
	// The ID of a ClassicLink-enabled VPC. Only applies to EC2-Classic instances. (eg. `vpc-2730681a`)
	VpcClassicLinkId pulumi.StringInput `pulumi:"vpcClassicLinkId"`
	// The IDs of one or more security groups for the specified ClassicLink-enabled VPC (eg. `sg-46ae3d11`).
	VpcClassicLinkSecurityGroups pulumi.StringArrayInput `pulumi:"vpcClassicLinkSecurityGroups"`
}
type LaunchConfigurationEbsBlockDevices struct {
	DeleteOnTermination *bool `pulumi:"deleteOnTermination"`
	DeviceName string `pulumi:"deviceName"`
	Encrypted *bool `pulumi:"encrypted"`
	Iops *int `pulumi:"iops"`
	NoDevice *bool `pulumi:"noDevice"`
	SnapshotId *string `pulumi:"snapshotId"`
	VolumeSize *int `pulumi:"volumeSize"`
	VolumeType *string `pulumi:"volumeType"`
}
var launchConfigurationEbsBlockDevicesType = reflect.TypeOf((*LaunchConfigurationEbsBlockDevices)(nil)).Elem()

type LaunchConfigurationEbsBlockDevicesInput interface {
	pulumi.Input

	ToLaunchConfigurationEbsBlockDevicesOutput() LaunchConfigurationEbsBlockDevicesOutput
	ToLaunchConfigurationEbsBlockDevicesOutputWithContext(ctx context.Context) LaunchConfigurationEbsBlockDevicesOutput
}

type LaunchConfigurationEbsBlockDevicesArgs struct {
	DeleteOnTermination pulumi.BoolInput `pulumi:"deleteOnTermination"`
	DeviceName pulumi.StringInput `pulumi:"deviceName"`
	Encrypted pulumi.BoolInput `pulumi:"encrypted"`
	Iops pulumi.IntInput `pulumi:"iops"`
	NoDevice pulumi.BoolInput `pulumi:"noDevice"`
	SnapshotId pulumi.StringInput `pulumi:"snapshotId"`
	VolumeSize pulumi.IntInput `pulumi:"volumeSize"`
	VolumeType pulumi.StringInput `pulumi:"volumeType"`
}

func (LaunchConfigurationEbsBlockDevicesArgs) ElementType() reflect.Type {
	return launchConfigurationEbsBlockDevicesType
}

func (a LaunchConfigurationEbsBlockDevicesArgs) ToLaunchConfigurationEbsBlockDevicesOutput() LaunchConfigurationEbsBlockDevicesOutput {
	return pulumi.ToOutput(a).(LaunchConfigurationEbsBlockDevicesOutput)
}

func (a LaunchConfigurationEbsBlockDevicesArgs) ToLaunchConfigurationEbsBlockDevicesOutputWithContext(ctx context.Context) LaunchConfigurationEbsBlockDevicesOutput {
	return pulumi.ToOutputWithContext(ctx, a).(LaunchConfigurationEbsBlockDevicesOutput)
}

type LaunchConfigurationEbsBlockDevicesOutput struct { *pulumi.OutputState }

func (o LaunchConfigurationEbsBlockDevicesOutput) DeleteOnTermination() pulumi.BoolOutput {
	return o.Apply(func(v LaunchConfigurationEbsBlockDevices) bool {
		if v.DeleteOnTermination == nil { return *new(bool) } else { return *v.DeleteOnTermination }
	}).(pulumi.BoolOutput)
}

func (o LaunchConfigurationEbsBlockDevicesOutput) DeviceName() pulumi.StringOutput {
	return o.Apply(func(v LaunchConfigurationEbsBlockDevices) string {
		return v.DeviceName
	}).(pulumi.StringOutput)
}

func (o LaunchConfigurationEbsBlockDevicesOutput) Encrypted() pulumi.BoolOutput {
	return o.Apply(func(v LaunchConfigurationEbsBlockDevices) bool {
		if v.Encrypted == nil { return *new(bool) } else { return *v.Encrypted }
	}).(pulumi.BoolOutput)
}

func (o LaunchConfigurationEbsBlockDevicesOutput) Iops() pulumi.IntOutput {
	return o.Apply(func(v LaunchConfigurationEbsBlockDevices) int {
		if v.Iops == nil { return *new(int) } else { return *v.Iops }
	}).(pulumi.IntOutput)
}

func (o LaunchConfigurationEbsBlockDevicesOutput) NoDevice() pulumi.BoolOutput {
	return o.Apply(func(v LaunchConfigurationEbsBlockDevices) bool {
		if v.NoDevice == nil { return *new(bool) } else { return *v.NoDevice }
	}).(pulumi.BoolOutput)
}

func (o LaunchConfigurationEbsBlockDevicesOutput) SnapshotId() pulumi.StringOutput {
	return o.Apply(func(v LaunchConfigurationEbsBlockDevices) string {
		if v.SnapshotId == nil { return *new(string) } else { return *v.SnapshotId }
	}).(pulumi.StringOutput)
}

func (o LaunchConfigurationEbsBlockDevicesOutput) VolumeSize() pulumi.IntOutput {
	return o.Apply(func(v LaunchConfigurationEbsBlockDevices) int {
		if v.VolumeSize == nil { return *new(int) } else { return *v.VolumeSize }
	}).(pulumi.IntOutput)
}

func (o LaunchConfigurationEbsBlockDevicesOutput) VolumeType() pulumi.StringOutput {
	return o.Apply(func(v LaunchConfigurationEbsBlockDevices) string {
		if v.VolumeType == nil { return *new(string) } else { return *v.VolumeType }
	}).(pulumi.StringOutput)
}

func (LaunchConfigurationEbsBlockDevicesOutput) ElementType() reflect.Type {
	return launchConfigurationEbsBlockDevicesType
}

func (o LaunchConfigurationEbsBlockDevicesOutput) ToLaunchConfigurationEbsBlockDevicesOutput() LaunchConfigurationEbsBlockDevicesOutput {
	return o
}

func (o LaunchConfigurationEbsBlockDevicesOutput) ToLaunchConfigurationEbsBlockDevicesOutputWithContext(ctx context.Context) LaunchConfigurationEbsBlockDevicesOutput {
	return o
}

func init() { pulumi.RegisterOutputType(LaunchConfigurationEbsBlockDevicesOutput{}) }

var launchConfigurationEbsBlockDevicesArrayType = reflect.TypeOf((*[]LaunchConfigurationEbsBlockDevices)(nil)).Elem()

type LaunchConfigurationEbsBlockDevicesArrayInput interface {
	pulumi.Input

	ToLaunchConfigurationEbsBlockDevicesArrayOutput() LaunchConfigurationEbsBlockDevicesArrayOutput
	ToLaunchConfigurationEbsBlockDevicesArrayOutputWithContext(ctx context.Context) LaunchConfigurationEbsBlockDevicesArrayOutput
}

type LaunchConfigurationEbsBlockDevicesArrayArgs []LaunchConfigurationEbsBlockDevicesInput

func (LaunchConfigurationEbsBlockDevicesArrayArgs) ElementType() reflect.Type {
	return launchConfigurationEbsBlockDevicesArrayType
}

func (a LaunchConfigurationEbsBlockDevicesArrayArgs) ToLaunchConfigurationEbsBlockDevicesArrayOutput() LaunchConfigurationEbsBlockDevicesArrayOutput {
	return pulumi.ToOutput(a).(LaunchConfigurationEbsBlockDevicesArrayOutput)
}

func (a LaunchConfigurationEbsBlockDevicesArrayArgs) ToLaunchConfigurationEbsBlockDevicesArrayOutputWithContext(ctx context.Context) LaunchConfigurationEbsBlockDevicesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(LaunchConfigurationEbsBlockDevicesArrayOutput)
}

type LaunchConfigurationEbsBlockDevicesArrayOutput struct { *pulumi.OutputState }

func (o LaunchConfigurationEbsBlockDevicesArrayOutput) Index(i pulumi.IntInput) LaunchConfigurationEbsBlockDevicesOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) LaunchConfigurationEbsBlockDevices {
		return vs[0].([]LaunchConfigurationEbsBlockDevices)[vs[1].(int)]
	}).(LaunchConfigurationEbsBlockDevicesOutput)
}

func (LaunchConfigurationEbsBlockDevicesArrayOutput) ElementType() reflect.Type {
	return launchConfigurationEbsBlockDevicesArrayType
}

func (o LaunchConfigurationEbsBlockDevicesArrayOutput) ToLaunchConfigurationEbsBlockDevicesArrayOutput() LaunchConfigurationEbsBlockDevicesArrayOutput {
	return o
}

func (o LaunchConfigurationEbsBlockDevicesArrayOutput) ToLaunchConfigurationEbsBlockDevicesArrayOutputWithContext(ctx context.Context) LaunchConfigurationEbsBlockDevicesArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(LaunchConfigurationEbsBlockDevicesArrayOutput{}) }

type LaunchConfigurationEphemeralBlockDevices struct {
	DeviceName string `pulumi:"deviceName"`
	VirtualName string `pulumi:"virtualName"`
}
var launchConfigurationEphemeralBlockDevicesType = reflect.TypeOf((*LaunchConfigurationEphemeralBlockDevices)(nil)).Elem()

type LaunchConfigurationEphemeralBlockDevicesInput interface {
	pulumi.Input

	ToLaunchConfigurationEphemeralBlockDevicesOutput() LaunchConfigurationEphemeralBlockDevicesOutput
	ToLaunchConfigurationEphemeralBlockDevicesOutputWithContext(ctx context.Context) LaunchConfigurationEphemeralBlockDevicesOutput
}

type LaunchConfigurationEphemeralBlockDevicesArgs struct {
	DeviceName pulumi.StringInput `pulumi:"deviceName"`
	VirtualName pulumi.StringInput `pulumi:"virtualName"`
}

func (LaunchConfigurationEphemeralBlockDevicesArgs) ElementType() reflect.Type {
	return launchConfigurationEphemeralBlockDevicesType
}

func (a LaunchConfigurationEphemeralBlockDevicesArgs) ToLaunchConfigurationEphemeralBlockDevicesOutput() LaunchConfigurationEphemeralBlockDevicesOutput {
	return pulumi.ToOutput(a).(LaunchConfigurationEphemeralBlockDevicesOutput)
}

func (a LaunchConfigurationEphemeralBlockDevicesArgs) ToLaunchConfigurationEphemeralBlockDevicesOutputWithContext(ctx context.Context) LaunchConfigurationEphemeralBlockDevicesOutput {
	return pulumi.ToOutputWithContext(ctx, a).(LaunchConfigurationEphemeralBlockDevicesOutput)
}

type LaunchConfigurationEphemeralBlockDevicesOutput struct { *pulumi.OutputState }

func (o LaunchConfigurationEphemeralBlockDevicesOutput) DeviceName() pulumi.StringOutput {
	return o.Apply(func(v LaunchConfigurationEphemeralBlockDevices) string {
		return v.DeviceName
	}).(pulumi.StringOutput)
}

func (o LaunchConfigurationEphemeralBlockDevicesOutput) VirtualName() pulumi.StringOutput {
	return o.Apply(func(v LaunchConfigurationEphemeralBlockDevices) string {
		return v.VirtualName
	}).(pulumi.StringOutput)
}

func (LaunchConfigurationEphemeralBlockDevicesOutput) ElementType() reflect.Type {
	return launchConfigurationEphemeralBlockDevicesType
}

func (o LaunchConfigurationEphemeralBlockDevicesOutput) ToLaunchConfigurationEphemeralBlockDevicesOutput() LaunchConfigurationEphemeralBlockDevicesOutput {
	return o
}

func (o LaunchConfigurationEphemeralBlockDevicesOutput) ToLaunchConfigurationEphemeralBlockDevicesOutputWithContext(ctx context.Context) LaunchConfigurationEphemeralBlockDevicesOutput {
	return o
}

func init() { pulumi.RegisterOutputType(LaunchConfigurationEphemeralBlockDevicesOutput{}) }

var launchConfigurationEphemeralBlockDevicesArrayType = reflect.TypeOf((*[]LaunchConfigurationEphemeralBlockDevices)(nil)).Elem()

type LaunchConfigurationEphemeralBlockDevicesArrayInput interface {
	pulumi.Input

	ToLaunchConfigurationEphemeralBlockDevicesArrayOutput() LaunchConfigurationEphemeralBlockDevicesArrayOutput
	ToLaunchConfigurationEphemeralBlockDevicesArrayOutputWithContext(ctx context.Context) LaunchConfigurationEphemeralBlockDevicesArrayOutput
}

type LaunchConfigurationEphemeralBlockDevicesArrayArgs []LaunchConfigurationEphemeralBlockDevicesInput

func (LaunchConfigurationEphemeralBlockDevicesArrayArgs) ElementType() reflect.Type {
	return launchConfigurationEphemeralBlockDevicesArrayType
}

func (a LaunchConfigurationEphemeralBlockDevicesArrayArgs) ToLaunchConfigurationEphemeralBlockDevicesArrayOutput() LaunchConfigurationEphemeralBlockDevicesArrayOutput {
	return pulumi.ToOutput(a).(LaunchConfigurationEphemeralBlockDevicesArrayOutput)
}

func (a LaunchConfigurationEphemeralBlockDevicesArrayArgs) ToLaunchConfigurationEphemeralBlockDevicesArrayOutputWithContext(ctx context.Context) LaunchConfigurationEphemeralBlockDevicesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(LaunchConfigurationEphemeralBlockDevicesArrayOutput)
}

type LaunchConfigurationEphemeralBlockDevicesArrayOutput struct { *pulumi.OutputState }

func (o LaunchConfigurationEphemeralBlockDevicesArrayOutput) Index(i pulumi.IntInput) LaunchConfigurationEphemeralBlockDevicesOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) LaunchConfigurationEphemeralBlockDevices {
		return vs[0].([]LaunchConfigurationEphemeralBlockDevices)[vs[1].(int)]
	}).(LaunchConfigurationEphemeralBlockDevicesOutput)
}

func (LaunchConfigurationEphemeralBlockDevicesArrayOutput) ElementType() reflect.Type {
	return launchConfigurationEphemeralBlockDevicesArrayType
}

func (o LaunchConfigurationEphemeralBlockDevicesArrayOutput) ToLaunchConfigurationEphemeralBlockDevicesArrayOutput() LaunchConfigurationEphemeralBlockDevicesArrayOutput {
	return o
}

func (o LaunchConfigurationEphemeralBlockDevicesArrayOutput) ToLaunchConfigurationEphemeralBlockDevicesArrayOutputWithContext(ctx context.Context) LaunchConfigurationEphemeralBlockDevicesArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(LaunchConfigurationEphemeralBlockDevicesArrayOutput{}) }

type LaunchConfigurationRootBlockDevice struct {
	DeleteOnTermination *bool `pulumi:"deleteOnTermination"`
	Encrypted *bool `pulumi:"encrypted"`
	Iops *int `pulumi:"iops"`
	VolumeSize *int `pulumi:"volumeSize"`
	VolumeType *string `pulumi:"volumeType"`
}
var launchConfigurationRootBlockDeviceType = reflect.TypeOf((*LaunchConfigurationRootBlockDevice)(nil)).Elem()

type LaunchConfigurationRootBlockDeviceInput interface {
	pulumi.Input

	ToLaunchConfigurationRootBlockDeviceOutput() LaunchConfigurationRootBlockDeviceOutput
	ToLaunchConfigurationRootBlockDeviceOutputWithContext(ctx context.Context) LaunchConfigurationRootBlockDeviceOutput
}

type LaunchConfigurationRootBlockDeviceArgs struct {
	DeleteOnTermination pulumi.BoolInput `pulumi:"deleteOnTermination"`
	Encrypted pulumi.BoolInput `pulumi:"encrypted"`
	Iops pulumi.IntInput `pulumi:"iops"`
	VolumeSize pulumi.IntInput `pulumi:"volumeSize"`
	VolumeType pulumi.StringInput `pulumi:"volumeType"`
}

func (LaunchConfigurationRootBlockDeviceArgs) ElementType() reflect.Type {
	return launchConfigurationRootBlockDeviceType
}

func (a LaunchConfigurationRootBlockDeviceArgs) ToLaunchConfigurationRootBlockDeviceOutput() LaunchConfigurationRootBlockDeviceOutput {
	return pulumi.ToOutput(a).(LaunchConfigurationRootBlockDeviceOutput)
}

func (a LaunchConfigurationRootBlockDeviceArgs) ToLaunchConfigurationRootBlockDeviceOutputWithContext(ctx context.Context) LaunchConfigurationRootBlockDeviceOutput {
	return pulumi.ToOutputWithContext(ctx, a).(LaunchConfigurationRootBlockDeviceOutput)
}

type LaunchConfigurationRootBlockDeviceOutput struct { *pulumi.OutputState }

func (o LaunchConfigurationRootBlockDeviceOutput) DeleteOnTermination() pulumi.BoolOutput {
	return o.Apply(func(v LaunchConfigurationRootBlockDevice) bool {
		if v.DeleteOnTermination == nil { return *new(bool) } else { return *v.DeleteOnTermination }
	}).(pulumi.BoolOutput)
}

func (o LaunchConfigurationRootBlockDeviceOutput) Encrypted() pulumi.BoolOutput {
	return o.Apply(func(v LaunchConfigurationRootBlockDevice) bool {
		if v.Encrypted == nil { return *new(bool) } else { return *v.Encrypted }
	}).(pulumi.BoolOutput)
}

func (o LaunchConfigurationRootBlockDeviceOutput) Iops() pulumi.IntOutput {
	return o.Apply(func(v LaunchConfigurationRootBlockDevice) int {
		if v.Iops == nil { return *new(int) } else { return *v.Iops }
	}).(pulumi.IntOutput)
}

func (o LaunchConfigurationRootBlockDeviceOutput) VolumeSize() pulumi.IntOutput {
	return o.Apply(func(v LaunchConfigurationRootBlockDevice) int {
		if v.VolumeSize == nil { return *new(int) } else { return *v.VolumeSize }
	}).(pulumi.IntOutput)
}

func (o LaunchConfigurationRootBlockDeviceOutput) VolumeType() pulumi.StringOutput {
	return o.Apply(func(v LaunchConfigurationRootBlockDevice) string {
		if v.VolumeType == nil { return *new(string) } else { return *v.VolumeType }
	}).(pulumi.StringOutput)
}

func (LaunchConfigurationRootBlockDeviceOutput) ElementType() reflect.Type {
	return launchConfigurationRootBlockDeviceType
}

func (o LaunchConfigurationRootBlockDeviceOutput) ToLaunchConfigurationRootBlockDeviceOutput() LaunchConfigurationRootBlockDeviceOutput {
	return o
}

func (o LaunchConfigurationRootBlockDeviceOutput) ToLaunchConfigurationRootBlockDeviceOutputWithContext(ctx context.Context) LaunchConfigurationRootBlockDeviceOutput {
	return o
}

func init() { pulumi.RegisterOutputType(LaunchConfigurationRootBlockDeviceOutput{}) }

