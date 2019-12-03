// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datasync

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an AWS DataSync EFS Location.
// 
// > **NOTE:** The EFS File System must have a mounted EFS Mount Target before creating this resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/datasync_location_efs.html.markdown.
type EfsLocation struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// Configuration block containing EC2 configurations for connecting to the EFS File System.
	Ec2Config EfsLocationEc2ConfigOutput `pulumi:"ec2Config"`

	// Amazon Resource Name (ARN) of EFS File System.
	EfsFileSystemArn pulumi.StringOutput `pulumi:"efsFileSystemArn"`

	// Subdirectory to perform actions as source or destination. Default `/`.
	Subdirectory pulumi.StringOutput `pulumi:"subdirectory"`

	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags pulumi.StringMapOutput `pulumi:"tags"`

	Uri pulumi.StringOutput `pulumi:"uri"`
}

// NewEfsLocation registers a new resource with the given unique name, arguments, and options.
func NewEfsLocation(ctx *pulumi.Context,
	name string, args *EfsLocationArgs, opts ...pulumi.ResourceOption) (*EfsLocation, error) {
	if args == nil || args.Ec2Config == nil {
		return nil, errors.New("missing required argument 'Ec2Config'")
	}
	if args == nil || args.EfsFileSystemArn == nil {
		return nil, errors.New("missing required argument 'EfsFileSystemArn'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Ec2Config; i != nil { inputs["ec2Config"] = i.ToEfsLocationEc2ConfigOutput() }
		if i := args.EfsFileSystemArn; i != nil { inputs["efsFileSystemArn"] = i.ToStringOutput() }
		if i := args.Subdirectory; i != nil { inputs["subdirectory"] = i.ToStringOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToStringMapOutput() }
	}
	var resource EfsLocation
	err := ctx.RegisterResource("aws:datasync/efsLocation:EfsLocation", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEfsLocation gets an existing EfsLocation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEfsLocation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EfsLocationState, opts ...pulumi.ResourceOption) (*EfsLocation, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Arn; i != nil { inputs["arn"] = i.ToStringOutput() }
		if i := state.Ec2Config; i != nil { inputs["ec2Config"] = i.ToEfsLocationEc2ConfigOutput() }
		if i := state.EfsFileSystemArn; i != nil { inputs["efsFileSystemArn"] = i.ToStringOutput() }
		if i := state.Subdirectory; i != nil { inputs["subdirectory"] = i.ToStringOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToStringMapOutput() }
		if i := state.Uri; i != nil { inputs["uri"] = i.ToStringOutput() }
	}
	var resource EfsLocation
	err := ctx.ReadResource("aws:datasync/efsLocation:EfsLocation", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EfsLocation resources.
type EfsLocationState struct {
	// Amazon Resource Name (ARN) of the DataSync Location.
	Arn pulumi.StringInput `pulumi:"arn"`
	// Configuration block containing EC2 configurations for connecting to the EFS File System.
	Ec2Config EfsLocationEc2ConfigInput `pulumi:"ec2Config"`
	// Amazon Resource Name (ARN) of EFS File System.
	EfsFileSystemArn pulumi.StringInput `pulumi:"efsFileSystemArn"`
	// Subdirectory to perform actions as source or destination. Default `/`.
	Subdirectory pulumi.StringInput `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags pulumi.StringMapInput `pulumi:"tags"`
	Uri pulumi.StringInput `pulumi:"uri"`
}

// The set of arguments for constructing a EfsLocation resource.
type EfsLocationArgs struct {
	// Configuration block containing EC2 configurations for connecting to the EFS File System.
	Ec2Config EfsLocationEc2ConfigInput `pulumi:"ec2Config"`
	// Amazon Resource Name (ARN) of EFS File System.
	EfsFileSystemArn pulumi.StringInput `pulumi:"efsFileSystemArn"`
	// Subdirectory to perform actions as source or destination. Default `/`.
	Subdirectory pulumi.StringInput `pulumi:"subdirectory"`
	// Key-value pairs of resource tags to assign to the DataSync Location.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}
type EfsLocationEc2Config struct {
	// List of Amazon Resource Names (ARNs) of the EC2 Security Groups that are associated with the EFS Mount Target.
	SecurityGroupArns []string `pulumi:"securityGroupArns"`
	// Amazon Resource Name (ARN) of the EC2 Subnet that is associated with the EFS Mount Target.
	SubnetArn string `pulumi:"subnetArn"`
}
var efsLocationEc2ConfigType = reflect.TypeOf((*EfsLocationEc2Config)(nil)).Elem()

type EfsLocationEc2ConfigInput interface {
	pulumi.Input

	ToEfsLocationEc2ConfigOutput() EfsLocationEc2ConfigOutput
	ToEfsLocationEc2ConfigOutputWithContext(ctx context.Context) EfsLocationEc2ConfigOutput
}

type EfsLocationEc2ConfigArgs struct {
	// List of Amazon Resource Names (ARNs) of the EC2 Security Groups that are associated with the EFS Mount Target.
	SecurityGroupArns pulumi.StringArrayInput `pulumi:"securityGroupArns"`
	// Amazon Resource Name (ARN) of the EC2 Subnet that is associated with the EFS Mount Target.
	SubnetArn pulumi.StringInput `pulumi:"subnetArn"`
}

func (EfsLocationEc2ConfigArgs) ElementType() reflect.Type {
	return efsLocationEc2ConfigType
}

func (a EfsLocationEc2ConfigArgs) ToEfsLocationEc2ConfigOutput() EfsLocationEc2ConfigOutput {
	return pulumi.ToOutput(a).(EfsLocationEc2ConfigOutput)
}

func (a EfsLocationEc2ConfigArgs) ToEfsLocationEc2ConfigOutputWithContext(ctx context.Context) EfsLocationEc2ConfigOutput {
	return pulumi.ToOutputWithContext(ctx, a).(EfsLocationEc2ConfigOutput)
}

type EfsLocationEc2ConfigOutput struct { *pulumi.OutputState }

// List of Amazon Resource Names (ARNs) of the EC2 Security Groups that are associated with the EFS Mount Target.
func (o EfsLocationEc2ConfigOutput) SecurityGroupArns() pulumi.StringArrayOutput {
	return o.Apply(func(v EfsLocationEc2Config) []string {
		return v.SecurityGroupArns
	}).(pulumi.StringArrayOutput)
}

// Amazon Resource Name (ARN) of the EC2 Subnet that is associated with the EFS Mount Target.
func (o EfsLocationEc2ConfigOutput) SubnetArn() pulumi.StringOutput {
	return o.Apply(func(v EfsLocationEc2Config) string {
		return v.SubnetArn
	}).(pulumi.StringOutput)
}

func (EfsLocationEc2ConfigOutput) ElementType() reflect.Type {
	return efsLocationEc2ConfigType
}

func (o EfsLocationEc2ConfigOutput) ToEfsLocationEc2ConfigOutput() EfsLocationEc2ConfigOutput {
	return o
}

func (o EfsLocationEc2ConfigOutput) ToEfsLocationEc2ConfigOutputWithContext(ctx context.Context) EfsLocationEc2ConfigOutput {
	return o
}

func init() { pulumi.RegisterOutputType(EfsLocationEc2ConfigOutput{}) }

