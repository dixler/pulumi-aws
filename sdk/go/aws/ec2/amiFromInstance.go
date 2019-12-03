// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The "AMI from instance" resource allows the creation of an Amazon Machine
// Image (AMI) modelled after an existing EBS-backed EC2 instance.
// 
// The created AMI will refer to implicitly-created snapshots of the instance's
// EBS volumes and mimick its assigned block device configuration at the time
// the resource is created.
// 
// This resource is best applied to an instance that is stopped when this instance
// is created, so that the contents of the created image are predictable. When
// applied to an instance that is running, *the instance will be stopped before taking
// the snapshots and then started back up again*, resulting in a period of
// downtime.
// 
// Note that the source instance is inspected only at the initial creation of this
// resource. Ongoing updates to the referenced instance will not be propagated into
// the generated AMI. Users may taint or otherwise recreate the resource in order
// to produce a fresh snapshot.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ami_from_instance.html.markdown.
type AmiFromInstance struct {
	pulumi.CustomResourceState

	// Machine architecture for created instances. Defaults to "x8664".
	Architecture pulumi.StringOutput `pulumi:"architecture"`

	// A longer, human-readable description for the AMI.
	Description pulumi.StringOutput `pulumi:"description"`

	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices AmiFromInstanceEbsBlockDevicesArrayOutput `pulumi:"ebsBlockDevices"`

	// Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
	EnaSupport pulumi.BoolOutput `pulumi:"enaSupport"`

	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices AmiFromInstanceEphemeralBlockDevicesArrayOutput `pulumi:"ephemeralBlockDevices"`

	// Path to an S3 object containing an image manifest, e.g. created
	// by the `ec2-upload-bundle` command in the EC2 command line tools.
	ImageLocation pulumi.StringOutput `pulumi:"imageLocation"`

	// The id of the kernel image (AKI) that will be used as the paravirtual
	// kernel in created instances.
	KernelId pulumi.StringOutput `pulumi:"kernelId"`

	ManageEbsSnapshots pulumi.BoolOutput `pulumi:"manageEbsSnapshots"`

	// A region-unique name for the AMI.
	Name pulumi.StringOutput `pulumi:"name"`

	// The id of an initrd image (ARI) that will be used when booting the
	// created instances.
	RamdiskId pulumi.StringOutput `pulumi:"ramdiskId"`

	// The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
	RootDeviceName pulumi.StringOutput `pulumi:"rootDeviceName"`

	RootSnapshotId pulumi.StringOutput `pulumi:"rootSnapshotId"`

	// Boolean that overrides the behavior of stopping
	// the instance before snapshotting. This is risky since it may cause a snapshot of an
	// inconsistent filesystem state, but can be used to avoid downtime if the user otherwise
	// guarantees that no filesystem writes will be underway at the time of snapshot.
	SnapshotWithoutReboot pulumi.BoolOutput `pulumi:"snapshotWithoutReboot"`

	// The id of the instance to use as the basis of the AMI.
	SourceInstanceId pulumi.StringOutput `pulumi:"sourceInstanceId"`

	// When set to "simple" (the default), enables enhanced networking
	// for created instances. No other value is supported at this time.
	SriovNetSupport pulumi.StringOutput `pulumi:"sriovNetSupport"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// Keyword to choose what virtualization mode created instances
	// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
	// changes the set of further arguments that are required, as described below.
	VirtualizationType pulumi.StringOutput `pulumi:"virtualizationType"`
}

// NewAmiFromInstance registers a new resource with the given unique name, arguments, and options.
func NewAmiFromInstance(ctx *pulumi.Context,
	name string, args *AmiFromInstanceArgs, opts ...pulumi.ResourceOption) (*AmiFromInstance, error) {
	if args == nil || args.SourceInstanceId == nil {
		return nil, errors.New("missing required argument 'SourceInstanceId'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Description; i != nil { inputs["description"] = i.ToStringOutput() }
		if i := args.EbsBlockDevices; i != nil { inputs["ebsBlockDevices"] = i.ToAmiFromInstanceEbsBlockDevicesArrayOutput() }
		if i := args.EphemeralBlockDevices; i != nil { inputs["ephemeralBlockDevices"] = i.ToAmiFromInstanceEphemeralBlockDevicesArrayOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.SnapshotWithoutReboot; i != nil { inputs["snapshotWithoutReboot"] = i.ToBoolOutput() }
		if i := args.SourceInstanceId; i != nil { inputs["sourceInstanceId"] = i.ToStringOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource AmiFromInstance
	err := ctx.RegisterResource("aws:ec2/amiFromInstance:AmiFromInstance", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAmiFromInstance gets an existing AmiFromInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAmiFromInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AmiFromInstanceState, opts ...pulumi.ResourceOption) (*AmiFromInstance, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Architecture; i != nil { inputs["architecture"] = i.ToStringOutput() }
		if i := state.Description; i != nil { inputs["description"] = i.ToStringOutput() }
		if i := state.EbsBlockDevices; i != nil { inputs["ebsBlockDevices"] = i.ToAmiFromInstanceEbsBlockDevicesArrayOutput() }
		if i := state.EnaSupport; i != nil { inputs["enaSupport"] = i.ToBoolOutput() }
		if i := state.EphemeralBlockDevices; i != nil { inputs["ephemeralBlockDevices"] = i.ToAmiFromInstanceEphemeralBlockDevicesArrayOutput() }
		if i := state.ImageLocation; i != nil { inputs["imageLocation"] = i.ToStringOutput() }
		if i := state.KernelId; i != nil { inputs["kernelId"] = i.ToStringOutput() }
		if i := state.ManageEbsSnapshots; i != nil { inputs["manageEbsSnapshots"] = i.ToBoolOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.RamdiskId; i != nil { inputs["ramdiskId"] = i.ToStringOutput() }
		if i := state.RootDeviceName; i != nil { inputs["rootDeviceName"] = i.ToStringOutput() }
		if i := state.RootSnapshotId; i != nil { inputs["rootSnapshotId"] = i.ToStringOutput() }
		if i := state.SnapshotWithoutReboot; i != nil { inputs["snapshotWithoutReboot"] = i.ToBoolOutput() }
		if i := state.SourceInstanceId; i != nil { inputs["sourceInstanceId"] = i.ToStringOutput() }
		if i := state.SriovNetSupport; i != nil { inputs["sriovNetSupport"] = i.ToStringOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := state.VirtualizationType; i != nil { inputs["virtualizationType"] = i.ToStringOutput() }
	}
	var resource AmiFromInstance
	err := ctx.ReadResource("aws:ec2/amiFromInstance:AmiFromInstance", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AmiFromInstance resources.
type AmiFromInstanceState struct {
	// Machine architecture for created instances. Defaults to "x8664".
	Architecture pulumi.StringInput `pulumi:"architecture"`
	// A longer, human-readable description for the AMI.
	Description pulumi.StringInput `pulumi:"description"`
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices AmiFromInstanceEbsBlockDevicesArrayInput `pulumi:"ebsBlockDevices"`
	// Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
	EnaSupport pulumi.BoolInput `pulumi:"enaSupport"`
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices AmiFromInstanceEphemeralBlockDevicesArrayInput `pulumi:"ephemeralBlockDevices"`
	// Path to an S3 object containing an image manifest, e.g. created
	// by the `ec2-upload-bundle` command in the EC2 command line tools.
	ImageLocation pulumi.StringInput `pulumi:"imageLocation"`
	// The id of the kernel image (AKI) that will be used as the paravirtual
	// kernel in created instances.
	KernelId pulumi.StringInput `pulumi:"kernelId"`
	ManageEbsSnapshots pulumi.BoolInput `pulumi:"manageEbsSnapshots"`
	// A region-unique name for the AMI.
	Name pulumi.StringInput `pulumi:"name"`
	// The id of an initrd image (ARI) that will be used when booting the
	// created instances.
	RamdiskId pulumi.StringInput `pulumi:"ramdiskId"`
	// The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
	RootDeviceName pulumi.StringInput `pulumi:"rootDeviceName"`
	RootSnapshotId pulumi.StringInput `pulumi:"rootSnapshotId"`
	// Boolean that overrides the behavior of stopping
	// the instance before snapshotting. This is risky since it may cause a snapshot of an
	// inconsistent filesystem state, but can be used to avoid downtime if the user otherwise
	// guarantees that no filesystem writes will be underway at the time of snapshot.
	SnapshotWithoutReboot pulumi.BoolInput `pulumi:"snapshotWithoutReboot"`
	// The id of the instance to use as the basis of the AMI.
	SourceInstanceId pulumi.StringInput `pulumi:"sourceInstanceId"`
	// When set to "simple" (the default), enables enhanced networking
	// for created instances. No other value is supported at this time.
	SriovNetSupport pulumi.StringInput `pulumi:"sriovNetSupport"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Keyword to choose what virtualization mode created instances
	// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
	// changes the set of further arguments that are required, as described below.
	VirtualizationType pulumi.StringInput `pulumi:"virtualizationType"`
}

// The set of arguments for constructing a AmiFromInstance resource.
type AmiFromInstanceArgs struct {
	// A longer, human-readable description for the AMI.
	Description pulumi.StringInput `pulumi:"description"`
	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EbsBlockDevices AmiFromInstanceEbsBlockDevicesArrayInput `pulumi:"ebsBlockDevices"`
	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevices AmiFromInstanceEphemeralBlockDevicesArrayInput `pulumi:"ephemeralBlockDevices"`
	// A region-unique name for the AMI.
	Name pulumi.StringInput `pulumi:"name"`
	// Boolean that overrides the behavior of stopping
	// the instance before snapshotting. This is risky since it may cause a snapshot of an
	// inconsistent filesystem state, but can be used to avoid downtime if the user otherwise
	// guarantees that no filesystem writes will be underway at the time of snapshot.
	SnapshotWithoutReboot pulumi.BoolInput `pulumi:"snapshotWithoutReboot"`
	// The id of the instance to use as the basis of the AMI.
	SourceInstanceId pulumi.StringInput `pulumi:"sourceInstanceId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}
type AmiFromInstanceEbsBlockDevices struct {
	DeleteOnTermination *bool `pulumi:"deleteOnTermination"`
	DeviceName *string `pulumi:"deviceName"`
	Encrypted *bool `pulumi:"encrypted"`
	Iops *int `pulumi:"iops"`
	SnapshotId *string `pulumi:"snapshotId"`
	VolumeSize *int `pulumi:"volumeSize"`
	VolumeType *string `pulumi:"volumeType"`
}
var amiFromInstanceEbsBlockDevicesType = reflect.TypeOf((*AmiFromInstanceEbsBlockDevices)(nil)).Elem()

type AmiFromInstanceEbsBlockDevicesInput interface {
	pulumi.Input

	ToAmiFromInstanceEbsBlockDevicesOutput() AmiFromInstanceEbsBlockDevicesOutput
	ToAmiFromInstanceEbsBlockDevicesOutputWithContext(ctx context.Context) AmiFromInstanceEbsBlockDevicesOutput
}

type AmiFromInstanceEbsBlockDevicesArgs struct {
	DeleteOnTermination pulumi.BoolInput `pulumi:"deleteOnTermination"`
	DeviceName pulumi.StringInput `pulumi:"deviceName"`
	Encrypted pulumi.BoolInput `pulumi:"encrypted"`
	Iops pulumi.IntInput `pulumi:"iops"`
	SnapshotId pulumi.StringInput `pulumi:"snapshotId"`
	VolumeSize pulumi.IntInput `pulumi:"volumeSize"`
	VolumeType pulumi.StringInput `pulumi:"volumeType"`
}

func (AmiFromInstanceEbsBlockDevicesArgs) ElementType() reflect.Type {
	return amiFromInstanceEbsBlockDevicesType
}

func (a AmiFromInstanceEbsBlockDevicesArgs) ToAmiFromInstanceEbsBlockDevicesOutput() AmiFromInstanceEbsBlockDevicesOutput {
	return pulumi.ToOutput(a).(AmiFromInstanceEbsBlockDevicesOutput)
}

func (a AmiFromInstanceEbsBlockDevicesArgs) ToAmiFromInstanceEbsBlockDevicesOutputWithContext(ctx context.Context) AmiFromInstanceEbsBlockDevicesOutput {
	return pulumi.ToOutputWithContext(ctx, a).(AmiFromInstanceEbsBlockDevicesOutput)
}

type AmiFromInstanceEbsBlockDevicesOutput struct { *pulumi.OutputState }

func (o AmiFromInstanceEbsBlockDevicesOutput) DeleteOnTermination() pulumi.BoolOutput {
	return o.Apply(func(v AmiFromInstanceEbsBlockDevices) bool {
		if v.DeleteOnTermination == nil { return *new(bool) } else { return *v.DeleteOnTermination }
	}).(pulumi.BoolOutput)
}

func (o AmiFromInstanceEbsBlockDevicesOutput) DeviceName() pulumi.StringOutput {
	return o.Apply(func(v AmiFromInstanceEbsBlockDevices) string {
		if v.DeviceName == nil { return *new(string) } else { return *v.DeviceName }
	}).(pulumi.StringOutput)
}

func (o AmiFromInstanceEbsBlockDevicesOutput) Encrypted() pulumi.BoolOutput {
	return o.Apply(func(v AmiFromInstanceEbsBlockDevices) bool {
		if v.Encrypted == nil { return *new(bool) } else { return *v.Encrypted }
	}).(pulumi.BoolOutput)
}

func (o AmiFromInstanceEbsBlockDevicesOutput) Iops() pulumi.IntOutput {
	return o.Apply(func(v AmiFromInstanceEbsBlockDevices) int {
		if v.Iops == nil { return *new(int) } else { return *v.Iops }
	}).(pulumi.IntOutput)
}

func (o AmiFromInstanceEbsBlockDevicesOutput) SnapshotId() pulumi.StringOutput {
	return o.Apply(func(v AmiFromInstanceEbsBlockDevices) string {
		if v.SnapshotId == nil { return *new(string) } else { return *v.SnapshotId }
	}).(pulumi.StringOutput)
}

func (o AmiFromInstanceEbsBlockDevicesOutput) VolumeSize() pulumi.IntOutput {
	return o.Apply(func(v AmiFromInstanceEbsBlockDevices) int {
		if v.VolumeSize == nil { return *new(int) } else { return *v.VolumeSize }
	}).(pulumi.IntOutput)
}

func (o AmiFromInstanceEbsBlockDevicesOutput) VolumeType() pulumi.StringOutput {
	return o.Apply(func(v AmiFromInstanceEbsBlockDevices) string {
		if v.VolumeType == nil { return *new(string) } else { return *v.VolumeType }
	}).(pulumi.StringOutput)
}

func (AmiFromInstanceEbsBlockDevicesOutput) ElementType() reflect.Type {
	return amiFromInstanceEbsBlockDevicesType
}

func (o AmiFromInstanceEbsBlockDevicesOutput) ToAmiFromInstanceEbsBlockDevicesOutput() AmiFromInstanceEbsBlockDevicesOutput {
	return o
}

func (o AmiFromInstanceEbsBlockDevicesOutput) ToAmiFromInstanceEbsBlockDevicesOutputWithContext(ctx context.Context) AmiFromInstanceEbsBlockDevicesOutput {
	return o
}

func init() { pulumi.RegisterOutputType(AmiFromInstanceEbsBlockDevicesOutput{}) }

var amiFromInstanceEbsBlockDevicesArrayType = reflect.TypeOf((*[]AmiFromInstanceEbsBlockDevices)(nil)).Elem()

type AmiFromInstanceEbsBlockDevicesArrayInput interface {
	pulumi.Input

	ToAmiFromInstanceEbsBlockDevicesArrayOutput() AmiFromInstanceEbsBlockDevicesArrayOutput
	ToAmiFromInstanceEbsBlockDevicesArrayOutputWithContext(ctx context.Context) AmiFromInstanceEbsBlockDevicesArrayOutput
}

type AmiFromInstanceEbsBlockDevicesArrayArgs []AmiFromInstanceEbsBlockDevicesInput

func (AmiFromInstanceEbsBlockDevicesArrayArgs) ElementType() reflect.Type {
	return amiFromInstanceEbsBlockDevicesArrayType
}

func (a AmiFromInstanceEbsBlockDevicesArrayArgs) ToAmiFromInstanceEbsBlockDevicesArrayOutput() AmiFromInstanceEbsBlockDevicesArrayOutput {
	return pulumi.ToOutput(a).(AmiFromInstanceEbsBlockDevicesArrayOutput)
}

func (a AmiFromInstanceEbsBlockDevicesArrayArgs) ToAmiFromInstanceEbsBlockDevicesArrayOutputWithContext(ctx context.Context) AmiFromInstanceEbsBlockDevicesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(AmiFromInstanceEbsBlockDevicesArrayOutput)
}

type AmiFromInstanceEbsBlockDevicesArrayOutput struct { *pulumi.OutputState }

func (o AmiFromInstanceEbsBlockDevicesArrayOutput) Index(i pulumi.IntInput) AmiFromInstanceEbsBlockDevicesOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) AmiFromInstanceEbsBlockDevices {
		return vs[0].([]AmiFromInstanceEbsBlockDevices)[vs[1].(int)]
	}).(AmiFromInstanceEbsBlockDevicesOutput)
}

func (AmiFromInstanceEbsBlockDevicesArrayOutput) ElementType() reflect.Type {
	return amiFromInstanceEbsBlockDevicesArrayType
}

func (o AmiFromInstanceEbsBlockDevicesArrayOutput) ToAmiFromInstanceEbsBlockDevicesArrayOutput() AmiFromInstanceEbsBlockDevicesArrayOutput {
	return o
}

func (o AmiFromInstanceEbsBlockDevicesArrayOutput) ToAmiFromInstanceEbsBlockDevicesArrayOutputWithContext(ctx context.Context) AmiFromInstanceEbsBlockDevicesArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(AmiFromInstanceEbsBlockDevicesArrayOutput{}) }

type AmiFromInstanceEphemeralBlockDevices struct {
	DeviceName *string `pulumi:"deviceName"`
	VirtualName *string `pulumi:"virtualName"`
}
var amiFromInstanceEphemeralBlockDevicesType = reflect.TypeOf((*AmiFromInstanceEphemeralBlockDevices)(nil)).Elem()

type AmiFromInstanceEphemeralBlockDevicesInput interface {
	pulumi.Input

	ToAmiFromInstanceEphemeralBlockDevicesOutput() AmiFromInstanceEphemeralBlockDevicesOutput
	ToAmiFromInstanceEphemeralBlockDevicesOutputWithContext(ctx context.Context) AmiFromInstanceEphemeralBlockDevicesOutput
}

type AmiFromInstanceEphemeralBlockDevicesArgs struct {
	DeviceName pulumi.StringInput `pulumi:"deviceName"`
	VirtualName pulumi.StringInput `pulumi:"virtualName"`
}

func (AmiFromInstanceEphemeralBlockDevicesArgs) ElementType() reflect.Type {
	return amiFromInstanceEphemeralBlockDevicesType
}

func (a AmiFromInstanceEphemeralBlockDevicesArgs) ToAmiFromInstanceEphemeralBlockDevicesOutput() AmiFromInstanceEphemeralBlockDevicesOutput {
	return pulumi.ToOutput(a).(AmiFromInstanceEphemeralBlockDevicesOutput)
}

func (a AmiFromInstanceEphemeralBlockDevicesArgs) ToAmiFromInstanceEphemeralBlockDevicesOutputWithContext(ctx context.Context) AmiFromInstanceEphemeralBlockDevicesOutput {
	return pulumi.ToOutputWithContext(ctx, a).(AmiFromInstanceEphemeralBlockDevicesOutput)
}

type AmiFromInstanceEphemeralBlockDevicesOutput struct { *pulumi.OutputState }

func (o AmiFromInstanceEphemeralBlockDevicesOutput) DeviceName() pulumi.StringOutput {
	return o.Apply(func(v AmiFromInstanceEphemeralBlockDevices) string {
		if v.DeviceName == nil { return *new(string) } else { return *v.DeviceName }
	}).(pulumi.StringOutput)
}

func (o AmiFromInstanceEphemeralBlockDevicesOutput) VirtualName() pulumi.StringOutput {
	return o.Apply(func(v AmiFromInstanceEphemeralBlockDevices) string {
		if v.VirtualName == nil { return *new(string) } else { return *v.VirtualName }
	}).(pulumi.StringOutput)
}

func (AmiFromInstanceEphemeralBlockDevicesOutput) ElementType() reflect.Type {
	return amiFromInstanceEphemeralBlockDevicesType
}

func (o AmiFromInstanceEphemeralBlockDevicesOutput) ToAmiFromInstanceEphemeralBlockDevicesOutput() AmiFromInstanceEphemeralBlockDevicesOutput {
	return o
}

func (o AmiFromInstanceEphemeralBlockDevicesOutput) ToAmiFromInstanceEphemeralBlockDevicesOutputWithContext(ctx context.Context) AmiFromInstanceEphemeralBlockDevicesOutput {
	return o
}

func init() { pulumi.RegisterOutputType(AmiFromInstanceEphemeralBlockDevicesOutput{}) }

var amiFromInstanceEphemeralBlockDevicesArrayType = reflect.TypeOf((*[]AmiFromInstanceEphemeralBlockDevices)(nil)).Elem()

type AmiFromInstanceEphemeralBlockDevicesArrayInput interface {
	pulumi.Input

	ToAmiFromInstanceEphemeralBlockDevicesArrayOutput() AmiFromInstanceEphemeralBlockDevicesArrayOutput
	ToAmiFromInstanceEphemeralBlockDevicesArrayOutputWithContext(ctx context.Context) AmiFromInstanceEphemeralBlockDevicesArrayOutput
}

type AmiFromInstanceEphemeralBlockDevicesArrayArgs []AmiFromInstanceEphemeralBlockDevicesInput

func (AmiFromInstanceEphemeralBlockDevicesArrayArgs) ElementType() reflect.Type {
	return amiFromInstanceEphemeralBlockDevicesArrayType
}

func (a AmiFromInstanceEphemeralBlockDevicesArrayArgs) ToAmiFromInstanceEphemeralBlockDevicesArrayOutput() AmiFromInstanceEphemeralBlockDevicesArrayOutput {
	return pulumi.ToOutput(a).(AmiFromInstanceEphemeralBlockDevicesArrayOutput)
}

func (a AmiFromInstanceEphemeralBlockDevicesArrayArgs) ToAmiFromInstanceEphemeralBlockDevicesArrayOutputWithContext(ctx context.Context) AmiFromInstanceEphemeralBlockDevicesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(AmiFromInstanceEphemeralBlockDevicesArrayOutput)
}

type AmiFromInstanceEphemeralBlockDevicesArrayOutput struct { *pulumi.OutputState }

func (o AmiFromInstanceEphemeralBlockDevicesArrayOutput) Index(i pulumi.IntInput) AmiFromInstanceEphemeralBlockDevicesOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) AmiFromInstanceEphemeralBlockDevices {
		return vs[0].([]AmiFromInstanceEphemeralBlockDevices)[vs[1].(int)]
	}).(AmiFromInstanceEphemeralBlockDevicesOutput)
}

func (AmiFromInstanceEphemeralBlockDevicesArrayOutput) ElementType() reflect.Type {
	return amiFromInstanceEphemeralBlockDevicesArrayType
}

func (o AmiFromInstanceEphemeralBlockDevicesArrayOutput) ToAmiFromInstanceEphemeralBlockDevicesArrayOutput() AmiFromInstanceEphemeralBlockDevicesArrayOutput {
	return o
}

func (o AmiFromInstanceEphemeralBlockDevicesArrayOutput) ToAmiFromInstanceEphemeralBlockDevicesArrayOutputWithContext(ctx context.Context) AmiFromInstanceEphemeralBlockDevicesArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(AmiFromInstanceEphemeralBlockDevicesArrayOutput{}) }

