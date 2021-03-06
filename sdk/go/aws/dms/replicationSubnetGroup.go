// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package dms

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a DMS (Data Migration Service) replication subnet group resource. DMS replication subnet groups can be created, updated, deleted, and imported.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dms_replication_subnet_group.html.markdown.
type ReplicationSubnetGroup struct {
	pulumi.CustomResourceState

	ReplicationSubnetGroupArn pulumi.StringOutput `pulumi:"replicationSubnetGroupArn"`
	// The description for the subnet group.
	ReplicationSubnetGroupDescription pulumi.StringOutput `pulumi:"replicationSubnetGroupDescription"`
	// The name for the replication subnet group. This value is stored as a lowercase string.
	ReplicationSubnetGroupId pulumi.StringOutput `pulumi:"replicationSubnetGroupId"`
	// A list of the EC2 subnet IDs for the subnet group.
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The ID of the VPC the subnet group is in.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewReplicationSubnetGroup registers a new resource with the given unique name, arguments, and options.
func NewReplicationSubnetGroup(ctx *pulumi.Context,
	name string, args *ReplicationSubnetGroupArgs, opts ...pulumi.ResourceOption) (*ReplicationSubnetGroup, error) {
	if args == nil || args.ReplicationSubnetGroupDescription == nil {
		return nil, errors.New("missing required argument 'ReplicationSubnetGroupDescription'")
	}
	if args == nil || args.ReplicationSubnetGroupId == nil {
		return nil, errors.New("missing required argument 'ReplicationSubnetGroupId'")
	}
	if args == nil || args.SubnetIds == nil {
		return nil, errors.New("missing required argument 'SubnetIds'")
	}
	if args == nil {
		args = &ReplicationSubnetGroupArgs{}
	}
	var resource ReplicationSubnetGroup
	err := ctx.RegisterResource("aws:dms/replicationSubnetGroup:ReplicationSubnetGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationSubnetGroup gets an existing ReplicationSubnetGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationSubnetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationSubnetGroupState, opts ...pulumi.ResourceOption) (*ReplicationSubnetGroup, error) {
	var resource ReplicationSubnetGroup
	err := ctx.ReadResource("aws:dms/replicationSubnetGroup:ReplicationSubnetGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationSubnetGroup resources.
type replicationSubnetGroupState struct {
	ReplicationSubnetGroupArn *string `pulumi:"replicationSubnetGroupArn"`
	// The description for the subnet group.
	ReplicationSubnetGroupDescription *string `pulumi:"replicationSubnetGroupDescription"`
	// The name for the replication subnet group. This value is stored as a lowercase string.
	ReplicationSubnetGroupId *string `pulumi:"replicationSubnetGroupId"`
	// A list of the EC2 subnet IDs for the subnet group.
	SubnetIds []string `pulumi:"subnetIds"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the VPC the subnet group is in.
	VpcId *string `pulumi:"vpcId"`
}

type ReplicationSubnetGroupState struct {
	ReplicationSubnetGroupArn pulumi.StringPtrInput
	// The description for the subnet group.
	ReplicationSubnetGroupDescription pulumi.StringPtrInput
	// The name for the replication subnet group. This value is stored as a lowercase string.
	ReplicationSubnetGroupId pulumi.StringPtrInput
	// A list of the EC2 subnet IDs for the subnet group.
	SubnetIds pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The ID of the VPC the subnet group is in.
	VpcId pulumi.StringPtrInput
}

func (ReplicationSubnetGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationSubnetGroupState)(nil)).Elem()
}

type replicationSubnetGroupArgs struct {
	// The description for the subnet group.
	ReplicationSubnetGroupDescription string `pulumi:"replicationSubnetGroupDescription"`
	// The name for the replication subnet group. This value is stored as a lowercase string.
	ReplicationSubnetGroupId string `pulumi:"replicationSubnetGroupId"`
	// A list of the EC2 subnet IDs for the subnet group.
	SubnetIds []string `pulumi:"subnetIds"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a ReplicationSubnetGroup resource.
type ReplicationSubnetGroupArgs struct {
	// The description for the subnet group.
	ReplicationSubnetGroupDescription pulumi.StringInput
	// The name for the replication subnet group. This value is stored as a lowercase string.
	ReplicationSubnetGroupId pulumi.StringInput
	// A list of the EC2 subnet IDs for the subnet group.
	SubnetIds pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (ReplicationSubnetGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationSubnetGroupArgs)(nil)).Elem()
}

