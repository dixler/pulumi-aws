// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a DMS (Data Migration Service) replication task resource. DMS replication tasks can be created, updated, deleted, and imported.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dms_replication_task.html.markdown.
type ReplicationTask struct {
	pulumi.CustomResourceState

	// The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
	CdcStartTime pulumi.StringOutput `pulumi:"cdcStartTime"`

	// The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
	MigrationType pulumi.StringOutput `pulumi:"migrationType"`

	// The Amazon Resource Name (ARN) of the replication instance.
	ReplicationInstanceArn pulumi.StringOutput `pulumi:"replicationInstanceArn"`

	// The Amazon Resource Name (ARN) for the replication task.
	ReplicationTaskArn pulumi.StringOutput `pulumi:"replicationTaskArn"`

	// The replication task identifier.
	ReplicationTaskId pulumi.StringOutput `pulumi:"replicationTaskId"`

	// An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
	ReplicationTaskSettings pulumi.StringOutput `pulumi:"replicationTaskSettings"`

	// The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
	SourceEndpointArn pulumi.StringOutput `pulumi:"sourceEndpointArn"`

	// An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
	TableMappings pulumi.StringOutput `pulumi:"tableMappings"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
	TargetEndpointArn pulumi.StringOutput `pulumi:"targetEndpointArn"`
}

// NewReplicationTask registers a new resource with the given unique name, arguments, and options.
func NewReplicationTask(ctx *pulumi.Context,
	name string, args *ReplicationTaskArgs, opts ...pulumi.ResourceOption) (*ReplicationTask, error) {
	if args == nil || args.MigrationType == nil {
		return nil, errors.New("missing required argument 'MigrationType'")
	}
	if args == nil || args.ReplicationInstanceArn == nil {
		return nil, errors.New("missing required argument 'ReplicationInstanceArn'")
	}
	if args == nil || args.ReplicationTaskId == nil {
		return nil, errors.New("missing required argument 'ReplicationTaskId'")
	}
	if args == nil || args.SourceEndpointArn == nil {
		return nil, errors.New("missing required argument 'SourceEndpointArn'")
	}
	if args == nil || args.TableMappings == nil {
		return nil, errors.New("missing required argument 'TableMappings'")
	}
	if args == nil || args.TargetEndpointArn == nil {
		return nil, errors.New("missing required argument 'TargetEndpointArn'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.CdcStartTime; i != nil { inputs["cdcStartTime"] = i.ToStringOutput() }
		if i := args.MigrationType; i != nil { inputs["migrationType"] = i.ToStringOutput() }
		if i := args.ReplicationInstanceArn; i != nil { inputs["replicationInstanceArn"] = i.ToStringOutput() }
		if i := args.ReplicationTaskId; i != nil { inputs["replicationTaskId"] = i.ToStringOutput() }
		if i := args.ReplicationTaskSettings; i != nil { inputs["replicationTaskSettings"] = i.ToStringOutput() }
		if i := args.SourceEndpointArn; i != nil { inputs["sourceEndpointArn"] = i.ToStringOutput() }
		if i := args.TableMappings; i != nil { inputs["tableMappings"] = i.ToStringOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := args.TargetEndpointArn; i != nil { inputs["targetEndpointArn"] = i.ToStringOutput() }
	}
	var resource ReplicationTask
	err := ctx.RegisterResource("aws:dms/replicationTask:ReplicationTask", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationTask gets an existing ReplicationTask resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationTaskState, opts ...pulumi.ResourceOption) (*ReplicationTask, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.CdcStartTime; i != nil { inputs["cdcStartTime"] = i.ToStringOutput() }
		if i := state.MigrationType; i != nil { inputs["migrationType"] = i.ToStringOutput() }
		if i := state.ReplicationInstanceArn; i != nil { inputs["replicationInstanceArn"] = i.ToStringOutput() }
		if i := state.ReplicationTaskArn; i != nil { inputs["replicationTaskArn"] = i.ToStringOutput() }
		if i := state.ReplicationTaskId; i != nil { inputs["replicationTaskId"] = i.ToStringOutput() }
		if i := state.ReplicationTaskSettings; i != nil { inputs["replicationTaskSettings"] = i.ToStringOutput() }
		if i := state.SourceEndpointArn; i != nil { inputs["sourceEndpointArn"] = i.ToStringOutput() }
		if i := state.TableMappings; i != nil { inputs["tableMappings"] = i.ToStringOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := state.TargetEndpointArn; i != nil { inputs["targetEndpointArn"] = i.ToStringOutput() }
	}
	var resource ReplicationTask
	err := ctx.ReadResource("aws:dms/replicationTask:ReplicationTask", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationTask resources.
type ReplicationTaskState struct {
	// The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
	CdcStartTime pulumi.StringInput `pulumi:"cdcStartTime"`
	// The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
	MigrationType pulumi.StringInput `pulumi:"migrationType"`
	// The Amazon Resource Name (ARN) of the replication instance.
	ReplicationInstanceArn pulumi.StringInput `pulumi:"replicationInstanceArn"`
	// The Amazon Resource Name (ARN) for the replication task.
	ReplicationTaskArn pulumi.StringInput `pulumi:"replicationTaskArn"`
	// The replication task identifier.
	ReplicationTaskId pulumi.StringInput `pulumi:"replicationTaskId"`
	// An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
	ReplicationTaskSettings pulumi.StringInput `pulumi:"replicationTaskSettings"`
	// The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
	SourceEndpointArn pulumi.StringInput `pulumi:"sourceEndpointArn"`
	// An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
	TableMappings pulumi.StringInput `pulumi:"tableMappings"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
	TargetEndpointArn pulumi.StringInput `pulumi:"targetEndpointArn"`
}

// The set of arguments for constructing a ReplicationTask resource.
type ReplicationTaskArgs struct {
	// The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
	CdcStartTime pulumi.StringInput `pulumi:"cdcStartTime"`
	// The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
	MigrationType pulumi.StringInput `pulumi:"migrationType"`
	// The Amazon Resource Name (ARN) of the replication instance.
	ReplicationInstanceArn pulumi.StringInput `pulumi:"replicationInstanceArn"`
	// The replication task identifier.
	ReplicationTaskId pulumi.StringInput `pulumi:"replicationTaskId"`
	// An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
	ReplicationTaskSettings pulumi.StringInput `pulumi:"replicationTaskSettings"`
	// The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
	SourceEndpointArn pulumi.StringInput `pulumi:"sourceEndpointArn"`
	// An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
	TableMappings pulumi.StringInput `pulumi:"tableMappings"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
	TargetEndpointArn pulumi.StringInput `pulumi:"targetEndpointArn"`
}
