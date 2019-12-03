// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package neptune

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// A Cluster Instance Resource defines attributes that are specific to a single instance in a Neptune Cluster.
// 
// You can simply add neptune instances and Neptune manages the replication. You can use the [count][1]
// meta-parameter to make multiple instances and join them all to the same Neptune Cluster, or you may specify different Cluster Instance resources with various `instanceClass` sizes.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/neptune_cluster_instance.html.markdown.
type ClusterInstance struct {
	pulumi.CustomResourceState

	// The hostname of the instance. See also `endpoint` and `port`.
	Address pulumi.StringOutput `pulumi:"address"`

	// Specifies whether any instance modifications
	// are applied immediately, or during the next maintenance window. Default is`false`.
	ApplyImmediately pulumi.BoolOutput `pulumi:"applyImmediately"`

	// Amazon Resource Name (ARN) of neptune instance
	Arn pulumi.StringOutput `pulumi:"arn"`

	// Indicates that minor engine upgrades will be applied automatically to the instance during the maintenance window. Default is `true`.
	AutoMinorVersionUpgrade pulumi.BoolOutput `pulumi:"autoMinorVersionUpgrade"`

	// The EC2 Availability Zone that the neptune instance is created in.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`

	// The identifier of the [`neptune.Cluster`](https://www.terraform.io/docs/providers/aws/r/neptune_cluster.html) in which to launch this instance.
	ClusterIdentifier pulumi.StringOutput `pulumi:"clusterIdentifier"`

	// The region-unique, immutable identifier for the neptune instance.
	DbiResourceId pulumi.StringOutput `pulumi:"dbiResourceId"`

	// The connection endpoint in `address:port` format.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`

	// The name of the database engine to be used for the neptune instance. Defaults to `neptune`. Valid Values: `neptune`.
	Engine pulumi.StringOutput `pulumi:"engine"`

	// The neptune engine version.
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`

	// The indentifier for the neptune instance, if omitted, this provider will assign a random, unique identifier.
	Identifier pulumi.StringOutput `pulumi:"identifier"`

	// Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
	IdentifierPrefix pulumi.StringOutput `pulumi:"identifierPrefix"`

	// The instance class to use.
	InstanceClass pulumi.StringOutput `pulumi:"instanceClass"`

	// The ARN for the KMS encryption key if one is set to the neptune cluster.
	KmsKeyArn pulumi.StringOutput `pulumi:"kmsKeyArn"`

	// The name of the neptune parameter group to associate with this instance.
	NeptuneParameterGroupName pulumi.StringOutput `pulumi:"neptuneParameterGroupName"`

	// A subnet group to associate with this neptune instance. **NOTE:** This must match the `neptuneSubnetGroupName` of the attached [`neptune.Cluster`](https://www.terraform.io/docs/providers/aws/r/neptune_cluster.html).
	NeptuneSubnetGroupName pulumi.StringOutput `pulumi:"neptuneSubnetGroupName"`

	// The port on which the DB accepts connections. Defaults to `8182`.
	Port pulumi.IntOutput `pulumi:"port"`

	// The daily time range during which automated backups are created if automated backups are enabled. Eg: "04:00-09:00"
	PreferredBackupWindow pulumi.StringOutput `pulumi:"preferredBackupWindow"`

	// The window to perform maintenance in.
	// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	PreferredMaintenanceWindow pulumi.StringOutput `pulumi:"preferredMaintenanceWindow"`

	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
	PromotionTier pulumi.IntOutput `pulumi:"promotionTier"`

	// Bool to control if instance is publicly accessible. Default is `false`.
	PubliclyAccessible pulumi.BoolOutput `pulumi:"publiclyAccessible"`

	// Specifies whether the neptune cluster is encrypted.
	StorageEncrypted pulumi.BoolOutput `pulumi:"storageEncrypted"`

	// A mapping of tags to assign to the instance.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// Boolean indicating if this instance is writable. `False` indicates this instance is a read replica.
	Writer pulumi.BoolOutput `pulumi:"writer"`
}

// NewClusterInstance registers a new resource with the given unique name, arguments, and options.
func NewClusterInstance(ctx *pulumi.Context,
	name string, args *ClusterInstanceArgs, opts ...pulumi.ResourceOption) (*ClusterInstance, error) {
	if args == nil || args.ClusterIdentifier == nil {
		return nil, errors.New("missing required argument 'ClusterIdentifier'")
	}
	if args == nil || args.InstanceClass == nil {
		return nil, errors.New("missing required argument 'InstanceClass'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.ApplyImmediately; i != nil { inputs["applyImmediately"] = i.ToBoolOutput() }
		if i := args.AutoMinorVersionUpgrade; i != nil { inputs["autoMinorVersionUpgrade"] = i.ToBoolOutput() }
		if i := args.AvailabilityZone; i != nil { inputs["availabilityZone"] = i.ToStringOutput() }
		if i := args.ClusterIdentifier; i != nil { inputs["clusterIdentifier"] = i.ToStringOutput() }
		if i := args.Engine; i != nil { inputs["engine"] = i.ToStringOutput() }
		if i := args.EngineVersion; i != nil { inputs["engineVersion"] = i.ToStringOutput() }
		if i := args.Identifier; i != nil { inputs["identifier"] = i.ToStringOutput() }
		if i := args.IdentifierPrefix; i != nil { inputs["identifierPrefix"] = i.ToStringOutput() }
		if i := args.InstanceClass; i != nil { inputs["instanceClass"] = i.ToStringOutput() }
		if i := args.NeptuneParameterGroupName; i != nil { inputs["neptuneParameterGroupName"] = i.ToStringOutput() }
		if i := args.NeptuneSubnetGroupName; i != nil { inputs["neptuneSubnetGroupName"] = i.ToStringOutput() }
		if i := args.Port; i != nil { inputs["port"] = i.ToIntOutput() }
		if i := args.PreferredBackupWindow; i != nil { inputs["preferredBackupWindow"] = i.ToStringOutput() }
		if i := args.PreferredMaintenanceWindow; i != nil { inputs["preferredMaintenanceWindow"] = i.ToStringOutput() }
		if i := args.PromotionTier; i != nil { inputs["promotionTier"] = i.ToIntOutput() }
		if i := args.PubliclyAccessible; i != nil { inputs["publiclyAccessible"] = i.ToBoolOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource ClusterInstance
	err := ctx.RegisterResource("aws:neptune/clusterInstance:ClusterInstance", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterInstance gets an existing ClusterInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterInstanceState, opts ...pulumi.ResourceOption) (*ClusterInstance, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Address; i != nil { inputs["address"] = i.ToStringOutput() }
		if i := state.ApplyImmediately; i != nil { inputs["applyImmediately"] = i.ToBoolOutput() }
		if i := state.Arn; i != nil { inputs["arn"] = i.ToStringOutput() }
		if i := state.AutoMinorVersionUpgrade; i != nil { inputs["autoMinorVersionUpgrade"] = i.ToBoolOutput() }
		if i := state.AvailabilityZone; i != nil { inputs["availabilityZone"] = i.ToStringOutput() }
		if i := state.ClusterIdentifier; i != nil { inputs["clusterIdentifier"] = i.ToStringOutput() }
		if i := state.DbiResourceId; i != nil { inputs["dbiResourceId"] = i.ToStringOutput() }
		if i := state.Endpoint; i != nil { inputs["endpoint"] = i.ToStringOutput() }
		if i := state.Engine; i != nil { inputs["engine"] = i.ToStringOutput() }
		if i := state.EngineVersion; i != nil { inputs["engineVersion"] = i.ToStringOutput() }
		if i := state.Identifier; i != nil { inputs["identifier"] = i.ToStringOutput() }
		if i := state.IdentifierPrefix; i != nil { inputs["identifierPrefix"] = i.ToStringOutput() }
		if i := state.InstanceClass; i != nil { inputs["instanceClass"] = i.ToStringOutput() }
		if i := state.KmsKeyArn; i != nil { inputs["kmsKeyArn"] = i.ToStringOutput() }
		if i := state.NeptuneParameterGroupName; i != nil { inputs["neptuneParameterGroupName"] = i.ToStringOutput() }
		if i := state.NeptuneSubnetGroupName; i != nil { inputs["neptuneSubnetGroupName"] = i.ToStringOutput() }
		if i := state.Port; i != nil { inputs["port"] = i.ToIntOutput() }
		if i := state.PreferredBackupWindow; i != nil { inputs["preferredBackupWindow"] = i.ToStringOutput() }
		if i := state.PreferredMaintenanceWindow; i != nil { inputs["preferredMaintenanceWindow"] = i.ToStringOutput() }
		if i := state.PromotionTier; i != nil { inputs["promotionTier"] = i.ToIntOutput() }
		if i := state.PubliclyAccessible; i != nil { inputs["publiclyAccessible"] = i.ToBoolOutput() }
		if i := state.StorageEncrypted; i != nil { inputs["storageEncrypted"] = i.ToBoolOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := state.Writer; i != nil { inputs["writer"] = i.ToBoolOutput() }
	}
	var resource ClusterInstance
	err := ctx.ReadResource("aws:neptune/clusterInstance:ClusterInstance", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterInstance resources.
type ClusterInstanceState struct {
	// The hostname of the instance. See also `endpoint` and `port`.
	Address pulumi.StringInput `pulumi:"address"`
	// Specifies whether any instance modifications
	// are applied immediately, or during the next maintenance window. Default is`false`.
	ApplyImmediately pulumi.BoolInput `pulumi:"applyImmediately"`
	// Amazon Resource Name (ARN) of neptune instance
	Arn pulumi.StringInput `pulumi:"arn"`
	// Indicates that minor engine upgrades will be applied automatically to the instance during the maintenance window. Default is `true`.
	AutoMinorVersionUpgrade pulumi.BoolInput `pulumi:"autoMinorVersionUpgrade"`
	// The EC2 Availability Zone that the neptune instance is created in.
	AvailabilityZone pulumi.StringInput `pulumi:"availabilityZone"`
	// The identifier of the [`neptune.Cluster`](https://www.terraform.io/docs/providers/aws/r/neptune_cluster.html) in which to launch this instance.
	ClusterIdentifier pulumi.StringInput `pulumi:"clusterIdentifier"`
	// The region-unique, immutable identifier for the neptune instance.
	DbiResourceId pulumi.StringInput `pulumi:"dbiResourceId"`
	// The connection endpoint in `address:port` format.
	Endpoint pulumi.StringInput `pulumi:"endpoint"`
	// The name of the database engine to be used for the neptune instance. Defaults to `neptune`. Valid Values: `neptune`.
	Engine pulumi.StringInput `pulumi:"engine"`
	// The neptune engine version.
	EngineVersion pulumi.StringInput `pulumi:"engineVersion"`
	// The indentifier for the neptune instance, if omitted, this provider will assign a random, unique identifier.
	Identifier pulumi.StringInput `pulumi:"identifier"`
	// Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
	IdentifierPrefix pulumi.StringInput `pulumi:"identifierPrefix"`
	// The instance class to use.
	InstanceClass pulumi.StringInput `pulumi:"instanceClass"`
	// The ARN for the KMS encryption key if one is set to the neptune cluster.
	KmsKeyArn pulumi.StringInput `pulumi:"kmsKeyArn"`
	// The name of the neptune parameter group to associate with this instance.
	NeptuneParameterGroupName pulumi.StringInput `pulumi:"neptuneParameterGroupName"`
	// A subnet group to associate with this neptune instance. **NOTE:** This must match the `neptuneSubnetGroupName` of the attached [`neptune.Cluster`](https://www.terraform.io/docs/providers/aws/r/neptune_cluster.html).
	NeptuneSubnetGroupName pulumi.StringInput `pulumi:"neptuneSubnetGroupName"`
	// The port on which the DB accepts connections. Defaults to `8182`.
	Port pulumi.IntInput `pulumi:"port"`
	// The daily time range during which automated backups are created if automated backups are enabled. Eg: "04:00-09:00"
	PreferredBackupWindow pulumi.StringInput `pulumi:"preferredBackupWindow"`
	// The window to perform maintenance in.
	// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	PreferredMaintenanceWindow pulumi.StringInput `pulumi:"preferredMaintenanceWindow"`
	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
	PromotionTier pulumi.IntInput `pulumi:"promotionTier"`
	// Bool to control if instance is publicly accessible. Default is `false`.
	PubliclyAccessible pulumi.BoolInput `pulumi:"publiclyAccessible"`
	// Specifies whether the neptune cluster is encrypted.
	StorageEncrypted pulumi.BoolInput `pulumi:"storageEncrypted"`
	// A mapping of tags to assign to the instance.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Boolean indicating if this instance is writable. `False` indicates this instance is a read replica.
	Writer pulumi.BoolInput `pulumi:"writer"`
}

// The set of arguments for constructing a ClusterInstance resource.
type ClusterInstanceArgs struct {
	// Specifies whether any instance modifications
	// are applied immediately, or during the next maintenance window. Default is`false`.
	ApplyImmediately pulumi.BoolInput `pulumi:"applyImmediately"`
	// Indicates that minor engine upgrades will be applied automatically to the instance during the maintenance window. Default is `true`.
	AutoMinorVersionUpgrade pulumi.BoolInput `pulumi:"autoMinorVersionUpgrade"`
	// The EC2 Availability Zone that the neptune instance is created in.
	AvailabilityZone pulumi.StringInput `pulumi:"availabilityZone"`
	// The identifier of the [`neptune.Cluster`](https://www.terraform.io/docs/providers/aws/r/neptune_cluster.html) in which to launch this instance.
	ClusterIdentifier pulumi.StringInput `pulumi:"clusterIdentifier"`
	// The name of the database engine to be used for the neptune instance. Defaults to `neptune`. Valid Values: `neptune`.
	Engine pulumi.StringInput `pulumi:"engine"`
	// The neptune engine version.
	EngineVersion pulumi.StringInput `pulumi:"engineVersion"`
	// The indentifier for the neptune instance, if omitted, this provider will assign a random, unique identifier.
	Identifier pulumi.StringInput `pulumi:"identifier"`
	// Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
	IdentifierPrefix pulumi.StringInput `pulumi:"identifierPrefix"`
	// The instance class to use.
	InstanceClass pulumi.StringInput `pulumi:"instanceClass"`
	// The name of the neptune parameter group to associate with this instance.
	NeptuneParameterGroupName pulumi.StringInput `pulumi:"neptuneParameterGroupName"`
	// A subnet group to associate with this neptune instance. **NOTE:** This must match the `neptuneSubnetGroupName` of the attached [`neptune.Cluster`](https://www.terraform.io/docs/providers/aws/r/neptune_cluster.html).
	NeptuneSubnetGroupName pulumi.StringInput `pulumi:"neptuneSubnetGroupName"`
	// The port on which the DB accepts connections. Defaults to `8182`.
	Port pulumi.IntInput `pulumi:"port"`
	// The daily time range during which automated backups are created if automated backups are enabled. Eg: "04:00-09:00"
	PreferredBackupWindow pulumi.StringInput `pulumi:"preferredBackupWindow"`
	// The window to perform maintenance in.
	// Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
	PreferredMaintenanceWindow pulumi.StringInput `pulumi:"preferredMaintenanceWindow"`
	// Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
	PromotionTier pulumi.IntInput `pulumi:"promotionTier"`
	// Bool to control if instance is publicly accessible. Default is `false`.
	PubliclyAccessible pulumi.BoolInput `pulumi:"publiclyAccessible"`
	// A mapping of tags to assign to the instance.
	Tags pulumi.MapInput `pulumi:"tags"`
}
