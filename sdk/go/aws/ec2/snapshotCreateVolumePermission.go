// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Adds permission to create volumes off of a given EBS Snapshot.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/snapshot_create_volume_permission.html.markdown.
type SnapshotCreateVolumePermission struct {
	pulumi.CustomResourceState

	// An AWS Account ID to add create volume permissions
	AccountId pulumi.StringOutput `pulumi:"accountId"`

	// A snapshot ID
	SnapshotId pulumi.StringOutput `pulumi:"snapshotId"`
}

// NewSnapshotCreateVolumePermission registers a new resource with the given unique name, arguments, and options.
func NewSnapshotCreateVolumePermission(ctx *pulumi.Context,
	name string, args *SnapshotCreateVolumePermissionArgs, opts ...pulumi.ResourceOption) (*SnapshotCreateVolumePermission, error) {
	if args == nil || args.AccountId == nil {
		return nil, errors.New("missing required argument 'AccountId'")
	}
	if args == nil || args.SnapshotId == nil {
		return nil, errors.New("missing required argument 'SnapshotId'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.AccountId; i != nil { inputs["accountId"] = i.ToStringOutput() }
		if i := args.SnapshotId; i != nil { inputs["snapshotId"] = i.ToStringOutput() }
	}
	var resource SnapshotCreateVolumePermission
	err := ctx.RegisterResource("aws:ec2/snapshotCreateVolumePermission:SnapshotCreateVolumePermission", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshotCreateVolumePermission gets an existing SnapshotCreateVolumePermission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshotCreateVolumePermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotCreateVolumePermissionState, opts ...pulumi.ResourceOption) (*SnapshotCreateVolumePermission, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.AccountId; i != nil { inputs["accountId"] = i.ToStringOutput() }
		if i := state.SnapshotId; i != nil { inputs["snapshotId"] = i.ToStringOutput() }
	}
	var resource SnapshotCreateVolumePermission
	err := ctx.ReadResource("aws:ec2/snapshotCreateVolumePermission:SnapshotCreateVolumePermission", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnapshotCreateVolumePermission resources.
type SnapshotCreateVolumePermissionState struct {
	// An AWS Account ID to add create volume permissions
	AccountId pulumi.StringInput `pulumi:"accountId"`
	// A snapshot ID
	SnapshotId pulumi.StringInput `pulumi:"snapshotId"`
}

// The set of arguments for constructing a SnapshotCreateVolumePermission resource.
type SnapshotCreateVolumePermissionArgs struct {
	// An AWS Account ID to add create volume permissions
	AccountId pulumi.StringInput `pulumi:"accountId"`
	// A snapshot ID
	SnapshotId pulumi.StringInput `pulumi:"snapshotId"`
}
