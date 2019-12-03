// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a RDS database instance snapshot. For managing RDS database cluster snapshots, see the [`rds.ClusterSnapshot` resource](https://www.terraform.io/docs/providers/aws/r/db_cluster_snapshot.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/db_snapshot.html.markdown.
type Snapshot struct {
	pulumi.CustomResourceState

	// Specifies the allocated storage size in gigabytes (GB).
	AllocatedStorage pulumi.IntOutput `pulumi:"allocatedStorage"`

	// Specifies the name of the Availability Zone the DB instance was located in at the time of the DB snapshot.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`

	// The DB Instance Identifier from which to take the snapshot.
	DbInstanceIdentifier pulumi.StringOutput `pulumi:"dbInstanceIdentifier"`

	// The Amazon Resource Name (ARN) for the DB snapshot.
	DbSnapshotArn pulumi.StringOutput `pulumi:"dbSnapshotArn"`

	// The Identifier for the snapshot.
	DbSnapshotIdentifier pulumi.StringOutput `pulumi:"dbSnapshotIdentifier"`

	// Specifies whether the DB snapshot is encrypted.
	Encrypted pulumi.BoolOutput `pulumi:"encrypted"`

	// Specifies the name of the database engine.
	Engine pulumi.StringOutput `pulumi:"engine"`

	// Specifies the version of the database engine.
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`

	// Specifies the Provisioned IOPS (I/O operations per second) value of the DB instance at the time of the snapshot.
	Iops pulumi.IntOutput `pulumi:"iops"`

	// The ARN for the KMS encryption key.
	KmsKeyId pulumi.StringOutput `pulumi:"kmsKeyId"`

	// License model information for the restored DB instance.
	LicenseModel pulumi.StringOutput `pulumi:"licenseModel"`

	// Provides the option group name for the DB snapshot.
	OptionGroupName pulumi.StringOutput `pulumi:"optionGroupName"`

	Port pulumi.IntOutput `pulumi:"port"`

	SnapshotType pulumi.StringOutput `pulumi:"snapshotType"`

	// The DB snapshot Arn that the DB snapshot was copied from. It only has value in case of cross customer or cross region copy.
	SourceDbSnapshotIdentifier pulumi.StringOutput `pulumi:"sourceDbSnapshotIdentifier"`

	// The region that the DB snapshot was created in or copied from.
	SourceRegion pulumi.StringOutput `pulumi:"sourceRegion"`

	// Specifies the status of this DB snapshot.
	Status pulumi.StringOutput `pulumi:"status"`

	// Specifies the storage type associated with DB snapshot.
	StorageType pulumi.StringOutput `pulumi:"storageType"`

	// Key-value mapping of resource tags
	Tags pulumi.MapOutput `pulumi:"tags"`

	// Specifies the storage type associated with DB snapshot.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil || args.DbInstanceIdentifier == nil {
		return nil, errors.New("missing required argument 'DbInstanceIdentifier'")
	}
	if args == nil || args.DbSnapshotIdentifier == nil {
		return nil, errors.New("missing required argument 'DbSnapshotIdentifier'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.DbInstanceIdentifier; i != nil { inputs["dbInstanceIdentifier"] = i.ToStringOutput() }
		if i := args.DbSnapshotIdentifier; i != nil { inputs["dbSnapshotIdentifier"] = i.ToStringOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource Snapshot
	err := ctx.RegisterResource("aws:rds/snapshot:Snapshot", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshot gets an existing Snapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.AllocatedStorage; i != nil { inputs["allocatedStorage"] = i.ToIntOutput() }
		if i := state.AvailabilityZone; i != nil { inputs["availabilityZone"] = i.ToStringOutput() }
		if i := state.DbInstanceIdentifier; i != nil { inputs["dbInstanceIdentifier"] = i.ToStringOutput() }
		if i := state.DbSnapshotArn; i != nil { inputs["dbSnapshotArn"] = i.ToStringOutput() }
		if i := state.DbSnapshotIdentifier; i != nil { inputs["dbSnapshotIdentifier"] = i.ToStringOutput() }
		if i := state.Encrypted; i != nil { inputs["encrypted"] = i.ToBoolOutput() }
		if i := state.Engine; i != nil { inputs["engine"] = i.ToStringOutput() }
		if i := state.EngineVersion; i != nil { inputs["engineVersion"] = i.ToStringOutput() }
		if i := state.Iops; i != nil { inputs["iops"] = i.ToIntOutput() }
		if i := state.KmsKeyId; i != nil { inputs["kmsKeyId"] = i.ToStringOutput() }
		if i := state.LicenseModel; i != nil { inputs["licenseModel"] = i.ToStringOutput() }
		if i := state.OptionGroupName; i != nil { inputs["optionGroupName"] = i.ToStringOutput() }
		if i := state.Port; i != nil { inputs["port"] = i.ToIntOutput() }
		if i := state.SnapshotType; i != nil { inputs["snapshotType"] = i.ToStringOutput() }
		if i := state.SourceDbSnapshotIdentifier; i != nil { inputs["sourceDbSnapshotIdentifier"] = i.ToStringOutput() }
		if i := state.SourceRegion; i != nil { inputs["sourceRegion"] = i.ToStringOutput() }
		if i := state.Status; i != nil { inputs["status"] = i.ToStringOutput() }
		if i := state.StorageType; i != nil { inputs["storageType"] = i.ToStringOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := state.VpcId; i != nil { inputs["vpcId"] = i.ToStringOutput() }
	}
	var resource Snapshot
	err := ctx.ReadResource("aws:rds/snapshot:Snapshot", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type SnapshotState struct {
	// Specifies the allocated storage size in gigabytes (GB).
	AllocatedStorage pulumi.IntInput `pulumi:"allocatedStorage"`
	// Specifies the name of the Availability Zone the DB instance was located in at the time of the DB snapshot.
	AvailabilityZone pulumi.StringInput `pulumi:"availabilityZone"`
	// The DB Instance Identifier from which to take the snapshot.
	DbInstanceIdentifier pulumi.StringInput `pulumi:"dbInstanceIdentifier"`
	// The Amazon Resource Name (ARN) for the DB snapshot.
	DbSnapshotArn pulumi.StringInput `pulumi:"dbSnapshotArn"`
	// The Identifier for the snapshot.
	DbSnapshotIdentifier pulumi.StringInput `pulumi:"dbSnapshotIdentifier"`
	// Specifies whether the DB snapshot is encrypted.
	Encrypted pulumi.BoolInput `pulumi:"encrypted"`
	// Specifies the name of the database engine.
	Engine pulumi.StringInput `pulumi:"engine"`
	// Specifies the version of the database engine.
	EngineVersion pulumi.StringInput `pulumi:"engineVersion"`
	// Specifies the Provisioned IOPS (I/O operations per second) value of the DB instance at the time of the snapshot.
	Iops pulumi.IntInput `pulumi:"iops"`
	// The ARN for the KMS encryption key.
	KmsKeyId pulumi.StringInput `pulumi:"kmsKeyId"`
	// License model information for the restored DB instance.
	LicenseModel pulumi.StringInput `pulumi:"licenseModel"`
	// Provides the option group name for the DB snapshot.
	OptionGroupName pulumi.StringInput `pulumi:"optionGroupName"`
	Port pulumi.IntInput `pulumi:"port"`
	SnapshotType pulumi.StringInput `pulumi:"snapshotType"`
	// The DB snapshot Arn that the DB snapshot was copied from. It only has value in case of cross customer or cross region copy.
	SourceDbSnapshotIdentifier pulumi.StringInput `pulumi:"sourceDbSnapshotIdentifier"`
	// The region that the DB snapshot was created in or copied from.
	SourceRegion pulumi.StringInput `pulumi:"sourceRegion"`
	// Specifies the status of this DB snapshot.
	Status pulumi.StringInput `pulumi:"status"`
	// Specifies the storage type associated with DB snapshot.
	StorageType pulumi.StringInput `pulumi:"storageType"`
	// Key-value mapping of resource tags
	Tags pulumi.MapInput `pulumi:"tags"`
	// Specifies the storage type associated with DB snapshot.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// The DB Instance Identifier from which to take the snapshot.
	DbInstanceIdentifier pulumi.StringInput `pulumi:"dbInstanceIdentifier"`
	// The Identifier for the snapshot.
	DbSnapshotIdentifier pulumi.StringInput `pulumi:"dbSnapshotIdentifier"`
	// Key-value mapping of resource tags
	Tags pulumi.MapInput `pulumi:"tags"`
}
