// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a revision of an ECS task definition to be used in `ecs.Service`.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ecs_task_definition.html.markdown.
type TaskDefinition struct {
	pulumi.CustomResourceState

	// Full ARN of the Task Definition (including both `family` and `revision`).
	Arn pulumi.StringOutput `pulumi:"arn"`

	// A list of valid [container definitions]
	// (http://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html) provided as a
	// single valid JSON document. Please note that you should only provide values that are part of the container
	// definition document. For a detailed description of what parameters are available, see the [Task Definition Parameters]
	// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html) section from the
	// official [Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide).
	ContainerDefinitions pulumi.StringOutput `pulumi:"containerDefinitions"`

	// The number of cpu units used by the task. If the `requiresCompatibilities` is `FARGATE` this field is required.
	Cpu pulumi.StringOutput `pulumi:"cpu"`

	// The Amazon Resource Name (ARN) of the task execution role that the Amazon ECS container agent and the Docker daemon can assume.
	ExecutionRoleArn pulumi.StringOutput `pulumi:"executionRoleArn"`

	// A unique name for your task definition.
	Family pulumi.StringOutput `pulumi:"family"`

	// The IPC resource namespace to be used for the containers in the task The valid values are `host`, `task`, and `none`.
	IpcMode pulumi.StringOutput `pulumi:"ipcMode"`

	// The amount (in MiB) of memory used by the task. If the `requiresCompatibilities` is `FARGATE` this field is required.
	Memory pulumi.StringOutput `pulumi:"memory"`

	// The Docker networking mode to use for the containers in the task. The valid values are `none`, `bridge`, `awsvpc`, and `host`.
	NetworkMode pulumi.StringOutput `pulumi:"networkMode"`

	// The process namespace to use for the containers in the task. The valid values are `host` and `task`.
	PidMode pulumi.StringOutput `pulumi:"pidMode"`

	// A set of placement constraints rules that are taken into consideration during task placement. Maximum number of `placementConstraints` is `10`.
	PlacementConstraints TaskDefinitionPlacementConstraintsArrayOutput `pulumi:"placementConstraints"`

	// The proxy configuration details for the App Mesh proxy.
	ProxyConfiguration TaskDefinitionProxyConfigurationOutput `pulumi:"proxyConfiguration"`

	// A set of launch types required by the task. The valid values are `EC2` and `FARGATE`.
	RequiresCompatibilities pulumi.StringArrayOutput `pulumi:"requiresCompatibilities"`

	// The revision of the task in a particular family.
	Revision pulumi.IntOutput `pulumi:"revision"`

	// Key-value mapping of resource tags
	Tags pulumi.MapOutput `pulumi:"tags"`

	// The ARN of IAM role that allows your Amazon ECS container task to make calls to other AWS services.
	TaskRoleArn pulumi.StringOutput `pulumi:"taskRoleArn"`

	// A set of volume blocks that containers in your task may use.
	Volumes TaskDefinitionVolumesArrayOutput `pulumi:"volumes"`
}

// NewTaskDefinition registers a new resource with the given unique name, arguments, and options.
func NewTaskDefinition(ctx *pulumi.Context,
	name string, args *TaskDefinitionArgs, opts ...pulumi.ResourceOption) (*TaskDefinition, error) {
	if args == nil || args.ContainerDefinitions == nil {
		return nil, errors.New("missing required argument 'ContainerDefinitions'")
	}
	if args == nil || args.Family == nil {
		return nil, errors.New("missing required argument 'Family'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.ContainerDefinitions; i != nil { inputs["containerDefinitions"] = i.ToStringOutput() }
		if i := args.Cpu; i != nil { inputs["cpu"] = i.ToStringOutput() }
		if i := args.ExecutionRoleArn; i != nil { inputs["executionRoleArn"] = i.ToStringOutput() }
		if i := args.Family; i != nil { inputs["family"] = i.ToStringOutput() }
		if i := args.IpcMode; i != nil { inputs["ipcMode"] = i.ToStringOutput() }
		if i := args.Memory; i != nil { inputs["memory"] = i.ToStringOutput() }
		if i := args.NetworkMode; i != nil { inputs["networkMode"] = i.ToStringOutput() }
		if i := args.PidMode; i != nil { inputs["pidMode"] = i.ToStringOutput() }
		if i := args.PlacementConstraints; i != nil { inputs["placementConstraints"] = i.ToTaskDefinitionPlacementConstraintsArrayOutput() }
		if i := args.ProxyConfiguration; i != nil { inputs["proxyConfiguration"] = i.ToTaskDefinitionProxyConfigurationOutput() }
		if i := args.RequiresCompatibilities; i != nil { inputs["requiresCompatibilities"] = i.ToStringArrayOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := args.TaskRoleArn; i != nil { inputs["taskRoleArn"] = i.ToStringOutput() }
		if i := args.Volumes; i != nil { inputs["volumes"] = i.ToTaskDefinitionVolumesArrayOutput() }
	}
	var resource TaskDefinition
	err := ctx.RegisterResource("aws:ecs/taskDefinition:TaskDefinition", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTaskDefinition gets an existing TaskDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTaskDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TaskDefinitionState, opts ...pulumi.ResourceOption) (*TaskDefinition, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Arn; i != nil { inputs["arn"] = i.ToStringOutput() }
		if i := state.ContainerDefinitions; i != nil { inputs["containerDefinitions"] = i.ToStringOutput() }
		if i := state.Cpu; i != nil { inputs["cpu"] = i.ToStringOutput() }
		if i := state.ExecutionRoleArn; i != nil { inputs["executionRoleArn"] = i.ToStringOutput() }
		if i := state.Family; i != nil { inputs["family"] = i.ToStringOutput() }
		if i := state.IpcMode; i != nil { inputs["ipcMode"] = i.ToStringOutput() }
		if i := state.Memory; i != nil { inputs["memory"] = i.ToStringOutput() }
		if i := state.NetworkMode; i != nil { inputs["networkMode"] = i.ToStringOutput() }
		if i := state.PidMode; i != nil { inputs["pidMode"] = i.ToStringOutput() }
		if i := state.PlacementConstraints; i != nil { inputs["placementConstraints"] = i.ToTaskDefinitionPlacementConstraintsArrayOutput() }
		if i := state.ProxyConfiguration; i != nil { inputs["proxyConfiguration"] = i.ToTaskDefinitionProxyConfigurationOutput() }
		if i := state.RequiresCompatibilities; i != nil { inputs["requiresCompatibilities"] = i.ToStringArrayOutput() }
		if i := state.Revision; i != nil { inputs["revision"] = i.ToIntOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := state.TaskRoleArn; i != nil { inputs["taskRoleArn"] = i.ToStringOutput() }
		if i := state.Volumes; i != nil { inputs["volumes"] = i.ToTaskDefinitionVolumesArrayOutput() }
	}
	var resource TaskDefinition
	err := ctx.ReadResource("aws:ecs/taskDefinition:TaskDefinition", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TaskDefinition resources.
type TaskDefinitionState struct {
	// Full ARN of the Task Definition (including both `family` and `revision`).
	Arn pulumi.StringInput `pulumi:"arn"`
	// A list of valid [container definitions]
	// (http://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html) provided as a
	// single valid JSON document. Please note that you should only provide values that are part of the container
	// definition document. For a detailed description of what parameters are available, see the [Task Definition Parameters]
	// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html) section from the
	// official [Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide).
	ContainerDefinitions pulumi.StringInput `pulumi:"containerDefinitions"`
	// The number of cpu units used by the task. If the `requiresCompatibilities` is `FARGATE` this field is required.
	Cpu pulumi.StringInput `pulumi:"cpu"`
	// The Amazon Resource Name (ARN) of the task execution role that the Amazon ECS container agent and the Docker daemon can assume.
	ExecutionRoleArn pulumi.StringInput `pulumi:"executionRoleArn"`
	// A unique name for your task definition.
	Family pulumi.StringInput `pulumi:"family"`
	// The IPC resource namespace to be used for the containers in the task The valid values are `host`, `task`, and `none`.
	IpcMode pulumi.StringInput `pulumi:"ipcMode"`
	// The amount (in MiB) of memory used by the task. If the `requiresCompatibilities` is `FARGATE` this field is required.
	Memory pulumi.StringInput `pulumi:"memory"`
	// The Docker networking mode to use for the containers in the task. The valid values are `none`, `bridge`, `awsvpc`, and `host`.
	NetworkMode pulumi.StringInput `pulumi:"networkMode"`
	// The process namespace to use for the containers in the task. The valid values are `host` and `task`.
	PidMode pulumi.StringInput `pulumi:"pidMode"`
	// A set of placement constraints rules that are taken into consideration during task placement. Maximum number of `placementConstraints` is `10`.
	PlacementConstraints TaskDefinitionPlacementConstraintsArrayInput `pulumi:"placementConstraints"`
	// The proxy configuration details for the App Mesh proxy.
	ProxyConfiguration TaskDefinitionProxyConfigurationInput `pulumi:"proxyConfiguration"`
	// A set of launch types required by the task. The valid values are `EC2` and `FARGATE`.
	RequiresCompatibilities pulumi.StringArrayInput `pulumi:"requiresCompatibilities"`
	// The revision of the task in a particular family.
	Revision pulumi.IntInput `pulumi:"revision"`
	// Key-value mapping of resource tags
	Tags pulumi.MapInput `pulumi:"tags"`
	// The ARN of IAM role that allows your Amazon ECS container task to make calls to other AWS services.
	TaskRoleArn pulumi.StringInput `pulumi:"taskRoleArn"`
	// A set of volume blocks that containers in your task may use.
	Volumes TaskDefinitionVolumesArrayInput `pulumi:"volumes"`
}

// The set of arguments for constructing a TaskDefinition resource.
type TaskDefinitionArgs struct {
	// A list of valid [container definitions]
	// (http://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html) provided as a
	// single valid JSON document. Please note that you should only provide values that are part of the container
	// definition document. For a detailed description of what parameters are available, see the [Task Definition Parameters]
	// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html) section from the
	// official [Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide).
	ContainerDefinitions pulumi.StringInput `pulumi:"containerDefinitions"`
	// The number of cpu units used by the task. If the `requiresCompatibilities` is `FARGATE` this field is required.
	Cpu pulumi.StringInput `pulumi:"cpu"`
	// The Amazon Resource Name (ARN) of the task execution role that the Amazon ECS container agent and the Docker daemon can assume.
	ExecutionRoleArn pulumi.StringInput `pulumi:"executionRoleArn"`
	// A unique name for your task definition.
	Family pulumi.StringInput `pulumi:"family"`
	// The IPC resource namespace to be used for the containers in the task The valid values are `host`, `task`, and `none`.
	IpcMode pulumi.StringInput `pulumi:"ipcMode"`
	// The amount (in MiB) of memory used by the task. If the `requiresCompatibilities` is `FARGATE` this field is required.
	Memory pulumi.StringInput `pulumi:"memory"`
	// The Docker networking mode to use for the containers in the task. The valid values are `none`, `bridge`, `awsvpc`, and `host`.
	NetworkMode pulumi.StringInput `pulumi:"networkMode"`
	// The process namespace to use for the containers in the task. The valid values are `host` and `task`.
	PidMode pulumi.StringInput `pulumi:"pidMode"`
	// A set of placement constraints rules that are taken into consideration during task placement. Maximum number of `placementConstraints` is `10`.
	PlacementConstraints TaskDefinitionPlacementConstraintsArrayInput `pulumi:"placementConstraints"`
	// The proxy configuration details for the App Mesh proxy.
	ProxyConfiguration TaskDefinitionProxyConfigurationInput `pulumi:"proxyConfiguration"`
	// A set of launch types required by the task. The valid values are `EC2` and `FARGATE`.
	RequiresCompatibilities pulumi.StringArrayInput `pulumi:"requiresCompatibilities"`
	// Key-value mapping of resource tags
	Tags pulumi.MapInput `pulumi:"tags"`
	// The ARN of IAM role that allows your Amazon ECS container task to make calls to other AWS services.
	TaskRoleArn pulumi.StringInput `pulumi:"taskRoleArn"`
	// A set of volume blocks that containers in your task may use.
	Volumes TaskDefinitionVolumesArrayInput `pulumi:"volumes"`
}
type TaskDefinitionPlacementConstraints struct {
	// Cluster Query Language expression to apply to the constraint.
	// For more information, see [Cluster Query Language in the Amazon EC2 Container
	// Service Developer
	// Guide](http://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html).
	Expression *string `pulumi:"expression"`
	// The proxy type. The default value is `APPMESH`. The only supported value is `APPMESH`.
	Type string `pulumi:"type"`
}
var taskDefinitionPlacementConstraintsType = reflect.TypeOf((*TaskDefinitionPlacementConstraints)(nil)).Elem()

type TaskDefinitionPlacementConstraintsInput interface {
	pulumi.Input

	ToTaskDefinitionPlacementConstraintsOutput() TaskDefinitionPlacementConstraintsOutput
	ToTaskDefinitionPlacementConstraintsOutputWithContext(ctx context.Context) TaskDefinitionPlacementConstraintsOutput
}

type TaskDefinitionPlacementConstraintsArgs struct {
	// Cluster Query Language expression to apply to the constraint.
	// For more information, see [Cluster Query Language in the Amazon EC2 Container
	// Service Developer
	// Guide](http://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html).
	Expression pulumi.StringInput `pulumi:"expression"`
	// The proxy type. The default value is `APPMESH`. The only supported value is `APPMESH`.
	Type pulumi.StringInput `pulumi:"type"`
}

func (TaskDefinitionPlacementConstraintsArgs) ElementType() reflect.Type {
	return taskDefinitionPlacementConstraintsType
}

func (a TaskDefinitionPlacementConstraintsArgs) ToTaskDefinitionPlacementConstraintsOutput() TaskDefinitionPlacementConstraintsOutput {
	return pulumi.ToOutput(a).(TaskDefinitionPlacementConstraintsOutput)
}

func (a TaskDefinitionPlacementConstraintsArgs) ToTaskDefinitionPlacementConstraintsOutputWithContext(ctx context.Context) TaskDefinitionPlacementConstraintsOutput {
	return pulumi.ToOutputWithContext(ctx, a).(TaskDefinitionPlacementConstraintsOutput)
}

type TaskDefinitionPlacementConstraintsOutput struct { *pulumi.OutputState }

// Cluster Query Language expression to apply to the constraint.
// For more information, see [Cluster Query Language in the Amazon EC2 Container
// Service Developer
// Guide](http://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html).
func (o TaskDefinitionPlacementConstraintsOutput) Expression() pulumi.StringOutput {
	return o.Apply(func(v TaskDefinitionPlacementConstraints) string {
		if v.Expression == nil { return *new(string) } else { return *v.Expression }
	}).(pulumi.StringOutput)
}

// The proxy type. The default value is `APPMESH`. The only supported value is `APPMESH`.
func (o TaskDefinitionPlacementConstraintsOutput) Type() pulumi.StringOutput {
	return o.Apply(func(v TaskDefinitionPlacementConstraints) string {
		return v.Type
	}).(pulumi.StringOutput)
}

func (TaskDefinitionPlacementConstraintsOutput) ElementType() reflect.Type {
	return taskDefinitionPlacementConstraintsType
}

func (o TaskDefinitionPlacementConstraintsOutput) ToTaskDefinitionPlacementConstraintsOutput() TaskDefinitionPlacementConstraintsOutput {
	return o
}

func (o TaskDefinitionPlacementConstraintsOutput) ToTaskDefinitionPlacementConstraintsOutputWithContext(ctx context.Context) TaskDefinitionPlacementConstraintsOutput {
	return o
}

func init() { pulumi.RegisterOutputType(TaskDefinitionPlacementConstraintsOutput{}) }

var taskDefinitionPlacementConstraintsArrayType = reflect.TypeOf((*[]TaskDefinitionPlacementConstraints)(nil)).Elem()

type TaskDefinitionPlacementConstraintsArrayInput interface {
	pulumi.Input

	ToTaskDefinitionPlacementConstraintsArrayOutput() TaskDefinitionPlacementConstraintsArrayOutput
	ToTaskDefinitionPlacementConstraintsArrayOutputWithContext(ctx context.Context) TaskDefinitionPlacementConstraintsArrayOutput
}

type TaskDefinitionPlacementConstraintsArrayArgs []TaskDefinitionPlacementConstraintsInput

func (TaskDefinitionPlacementConstraintsArrayArgs) ElementType() reflect.Type {
	return taskDefinitionPlacementConstraintsArrayType
}

func (a TaskDefinitionPlacementConstraintsArrayArgs) ToTaskDefinitionPlacementConstraintsArrayOutput() TaskDefinitionPlacementConstraintsArrayOutput {
	return pulumi.ToOutput(a).(TaskDefinitionPlacementConstraintsArrayOutput)
}

func (a TaskDefinitionPlacementConstraintsArrayArgs) ToTaskDefinitionPlacementConstraintsArrayOutputWithContext(ctx context.Context) TaskDefinitionPlacementConstraintsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(TaskDefinitionPlacementConstraintsArrayOutput)
}

type TaskDefinitionPlacementConstraintsArrayOutput struct { *pulumi.OutputState }

func (o TaskDefinitionPlacementConstraintsArrayOutput) Index(i pulumi.IntInput) TaskDefinitionPlacementConstraintsOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) TaskDefinitionPlacementConstraints {
		return vs[0].([]TaskDefinitionPlacementConstraints)[vs[1].(int)]
	}).(TaskDefinitionPlacementConstraintsOutput)
}

func (TaskDefinitionPlacementConstraintsArrayOutput) ElementType() reflect.Type {
	return taskDefinitionPlacementConstraintsArrayType
}

func (o TaskDefinitionPlacementConstraintsArrayOutput) ToTaskDefinitionPlacementConstraintsArrayOutput() TaskDefinitionPlacementConstraintsArrayOutput {
	return o
}

func (o TaskDefinitionPlacementConstraintsArrayOutput) ToTaskDefinitionPlacementConstraintsArrayOutputWithContext(ctx context.Context) TaskDefinitionPlacementConstraintsArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(TaskDefinitionPlacementConstraintsArrayOutput{}) }

type TaskDefinitionProxyConfiguration struct {
	// The name of the container that will serve as the App Mesh proxy.
	ContainerName string `pulumi:"containerName"`
	// The set of network configuration parameters to provide the Container Network Interface (CNI) plugin, specified a key-value mapping.
	Properties *map[string]string `pulumi:"properties"`
	// The proxy type. The default value is `APPMESH`. The only supported value is `APPMESH`.
	Type *string `pulumi:"type"`
}
var taskDefinitionProxyConfigurationType = reflect.TypeOf((*TaskDefinitionProxyConfiguration)(nil)).Elem()

type TaskDefinitionProxyConfigurationInput interface {
	pulumi.Input

	ToTaskDefinitionProxyConfigurationOutput() TaskDefinitionProxyConfigurationOutput
	ToTaskDefinitionProxyConfigurationOutputWithContext(ctx context.Context) TaskDefinitionProxyConfigurationOutput
}

type TaskDefinitionProxyConfigurationArgs struct {
	// The name of the container that will serve as the App Mesh proxy.
	ContainerName pulumi.StringInput `pulumi:"containerName"`
	// The set of network configuration parameters to provide the Container Network Interface (CNI) plugin, specified a key-value mapping.
	Properties pulumi.StringMapInput `pulumi:"properties"`
	// The proxy type. The default value is `APPMESH`. The only supported value is `APPMESH`.
	Type pulumi.StringInput `pulumi:"type"`
}

func (TaskDefinitionProxyConfigurationArgs) ElementType() reflect.Type {
	return taskDefinitionProxyConfigurationType
}

func (a TaskDefinitionProxyConfigurationArgs) ToTaskDefinitionProxyConfigurationOutput() TaskDefinitionProxyConfigurationOutput {
	return pulumi.ToOutput(a).(TaskDefinitionProxyConfigurationOutput)
}

func (a TaskDefinitionProxyConfigurationArgs) ToTaskDefinitionProxyConfigurationOutputWithContext(ctx context.Context) TaskDefinitionProxyConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, a).(TaskDefinitionProxyConfigurationOutput)
}

type TaskDefinitionProxyConfigurationOutput struct { *pulumi.OutputState }

// The name of the container that will serve as the App Mesh proxy.
func (o TaskDefinitionProxyConfigurationOutput) ContainerName() pulumi.StringOutput {
	return o.Apply(func(v TaskDefinitionProxyConfiguration) string {
		return v.ContainerName
	}).(pulumi.StringOutput)
}

// The set of network configuration parameters to provide the Container Network Interface (CNI) plugin, specified a key-value mapping.
func (o TaskDefinitionProxyConfigurationOutput) Properties() pulumi.StringMapOutput {
	return o.Apply(func(v TaskDefinitionProxyConfiguration) map[string]string {
		if v.Properties == nil { return *new(map[string]string) } else { return *v.Properties }
	}).(pulumi.StringMapOutput)
}

// The proxy type. The default value is `APPMESH`. The only supported value is `APPMESH`.
func (o TaskDefinitionProxyConfigurationOutput) Type() pulumi.StringOutput {
	return o.Apply(func(v TaskDefinitionProxyConfiguration) string {
		if v.Type == nil { return *new(string) } else { return *v.Type }
	}).(pulumi.StringOutput)
}

func (TaskDefinitionProxyConfigurationOutput) ElementType() reflect.Type {
	return taskDefinitionProxyConfigurationType
}

func (o TaskDefinitionProxyConfigurationOutput) ToTaskDefinitionProxyConfigurationOutput() TaskDefinitionProxyConfigurationOutput {
	return o
}

func (o TaskDefinitionProxyConfigurationOutput) ToTaskDefinitionProxyConfigurationOutputWithContext(ctx context.Context) TaskDefinitionProxyConfigurationOutput {
	return o
}

func init() { pulumi.RegisterOutputType(TaskDefinitionProxyConfigurationOutput{}) }

type TaskDefinitionVolumes struct {
	// Used to configure a docker volume
	DockerVolumeConfiguration *TaskDefinitionVolumesDockerVolumeConfiguration `pulumi:"dockerVolumeConfiguration"`
	// The path on the host container instance that is presented to the container. If not set, ECS will create a nonpersistent data volume that starts empty and is deleted after the task has finished.
	HostPath *string `pulumi:"hostPath"`
	// The name of the volume. This name is referenced in the `sourceVolume`
	// parameter of container definition in the `mountPoints` section.
	Name string `pulumi:"name"`
}
var taskDefinitionVolumesType = reflect.TypeOf((*TaskDefinitionVolumes)(nil)).Elem()

type TaskDefinitionVolumesInput interface {
	pulumi.Input

	ToTaskDefinitionVolumesOutput() TaskDefinitionVolumesOutput
	ToTaskDefinitionVolumesOutputWithContext(ctx context.Context) TaskDefinitionVolumesOutput
}

type TaskDefinitionVolumesArgs struct {
	// Used to configure a docker volume
	DockerVolumeConfiguration TaskDefinitionVolumesDockerVolumeConfigurationInput `pulumi:"dockerVolumeConfiguration"`
	// The path on the host container instance that is presented to the container. If not set, ECS will create a nonpersistent data volume that starts empty and is deleted after the task has finished.
	HostPath pulumi.StringInput `pulumi:"hostPath"`
	// The name of the volume. This name is referenced in the `sourceVolume`
	// parameter of container definition in the `mountPoints` section.
	Name pulumi.StringInput `pulumi:"name"`
}

func (TaskDefinitionVolumesArgs) ElementType() reflect.Type {
	return taskDefinitionVolumesType
}

func (a TaskDefinitionVolumesArgs) ToTaskDefinitionVolumesOutput() TaskDefinitionVolumesOutput {
	return pulumi.ToOutput(a).(TaskDefinitionVolumesOutput)
}

func (a TaskDefinitionVolumesArgs) ToTaskDefinitionVolumesOutputWithContext(ctx context.Context) TaskDefinitionVolumesOutput {
	return pulumi.ToOutputWithContext(ctx, a).(TaskDefinitionVolumesOutput)
}

type TaskDefinitionVolumesOutput struct { *pulumi.OutputState }

// Used to configure a docker volume
func (o TaskDefinitionVolumesOutput) DockerVolumeConfiguration() TaskDefinitionVolumesDockerVolumeConfigurationOutput {
	return o.Apply(func(v TaskDefinitionVolumes) TaskDefinitionVolumesDockerVolumeConfiguration {
		if v.DockerVolumeConfiguration == nil { return *new(TaskDefinitionVolumesDockerVolumeConfiguration) } else { return *v.DockerVolumeConfiguration }
	}).(TaskDefinitionVolumesDockerVolumeConfigurationOutput)
}

// The path on the host container instance that is presented to the container. If not set, ECS will create a nonpersistent data volume that starts empty and is deleted after the task has finished.
func (o TaskDefinitionVolumesOutput) HostPath() pulumi.StringOutput {
	return o.Apply(func(v TaskDefinitionVolumes) string {
		if v.HostPath == nil { return *new(string) } else { return *v.HostPath }
	}).(pulumi.StringOutput)
}

// The name of the volume. This name is referenced in the `sourceVolume`
// parameter of container definition in the `mountPoints` section.
func (o TaskDefinitionVolumesOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v TaskDefinitionVolumes) string {
		return v.Name
	}).(pulumi.StringOutput)
}

func (TaskDefinitionVolumesOutput) ElementType() reflect.Type {
	return taskDefinitionVolumesType
}

func (o TaskDefinitionVolumesOutput) ToTaskDefinitionVolumesOutput() TaskDefinitionVolumesOutput {
	return o
}

func (o TaskDefinitionVolumesOutput) ToTaskDefinitionVolumesOutputWithContext(ctx context.Context) TaskDefinitionVolumesOutput {
	return o
}

func init() { pulumi.RegisterOutputType(TaskDefinitionVolumesOutput{}) }

var taskDefinitionVolumesArrayType = reflect.TypeOf((*[]TaskDefinitionVolumes)(nil)).Elem()

type TaskDefinitionVolumesArrayInput interface {
	pulumi.Input

	ToTaskDefinitionVolumesArrayOutput() TaskDefinitionVolumesArrayOutput
	ToTaskDefinitionVolumesArrayOutputWithContext(ctx context.Context) TaskDefinitionVolumesArrayOutput
}

type TaskDefinitionVolumesArrayArgs []TaskDefinitionVolumesInput

func (TaskDefinitionVolumesArrayArgs) ElementType() reflect.Type {
	return taskDefinitionVolumesArrayType
}

func (a TaskDefinitionVolumesArrayArgs) ToTaskDefinitionVolumesArrayOutput() TaskDefinitionVolumesArrayOutput {
	return pulumi.ToOutput(a).(TaskDefinitionVolumesArrayOutput)
}

func (a TaskDefinitionVolumesArrayArgs) ToTaskDefinitionVolumesArrayOutputWithContext(ctx context.Context) TaskDefinitionVolumesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(TaskDefinitionVolumesArrayOutput)
}

type TaskDefinitionVolumesArrayOutput struct { *pulumi.OutputState }

func (o TaskDefinitionVolumesArrayOutput) Index(i pulumi.IntInput) TaskDefinitionVolumesOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) TaskDefinitionVolumes {
		return vs[0].([]TaskDefinitionVolumes)[vs[1].(int)]
	}).(TaskDefinitionVolumesOutput)
}

func (TaskDefinitionVolumesArrayOutput) ElementType() reflect.Type {
	return taskDefinitionVolumesArrayType
}

func (o TaskDefinitionVolumesArrayOutput) ToTaskDefinitionVolumesArrayOutput() TaskDefinitionVolumesArrayOutput {
	return o
}

func (o TaskDefinitionVolumesArrayOutput) ToTaskDefinitionVolumesArrayOutputWithContext(ctx context.Context) TaskDefinitionVolumesArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(TaskDefinitionVolumesArrayOutput{}) }

type TaskDefinitionVolumesDockerVolumeConfiguration struct {
	// If this value is `true`, the Docker volume is created if it does not already exist. *Note*: This field is only used if the scope is `shared`.
	Autoprovision *bool `pulumi:"autoprovision"`
	// The Docker volume driver to use. The driver value must match the driver name provided by Docker because it is used for task placement.
	Driver *string `pulumi:"driver"`
	// A map of Docker driver specific options.
	DriverOpts *map[string]string `pulumi:"driverOpts"`
	// A map of custom metadata to add to your Docker volume.
	Labels *map[string]string `pulumi:"labels"`
	// The scope for the Docker volume, which determines its lifecycle, either `task` or `shared`.  Docker volumes that are scoped to a `task` are automatically provisioned when the task starts and destroyed when the task stops. Docker volumes that are `scoped` as shared persist after the task stops.
	Scope *string `pulumi:"scope"`
}
var taskDefinitionVolumesDockerVolumeConfigurationType = reflect.TypeOf((*TaskDefinitionVolumesDockerVolumeConfiguration)(nil)).Elem()

type TaskDefinitionVolumesDockerVolumeConfigurationInput interface {
	pulumi.Input

	ToTaskDefinitionVolumesDockerVolumeConfigurationOutput() TaskDefinitionVolumesDockerVolumeConfigurationOutput
	ToTaskDefinitionVolumesDockerVolumeConfigurationOutputWithContext(ctx context.Context) TaskDefinitionVolumesDockerVolumeConfigurationOutput
}

type TaskDefinitionVolumesDockerVolumeConfigurationArgs struct {
	// If this value is `true`, the Docker volume is created if it does not already exist. *Note*: This field is only used if the scope is `shared`.
	Autoprovision pulumi.BoolInput `pulumi:"autoprovision"`
	// The Docker volume driver to use. The driver value must match the driver name provided by Docker because it is used for task placement.
	Driver pulumi.StringInput `pulumi:"driver"`
	// A map of Docker driver specific options.
	DriverOpts pulumi.StringMapInput `pulumi:"driverOpts"`
	// A map of custom metadata to add to your Docker volume.
	Labels pulumi.StringMapInput `pulumi:"labels"`
	// The scope for the Docker volume, which determines its lifecycle, either `task` or `shared`.  Docker volumes that are scoped to a `task` are automatically provisioned when the task starts and destroyed when the task stops. Docker volumes that are `scoped` as shared persist after the task stops.
	Scope pulumi.StringInput `pulumi:"scope"`
}

func (TaskDefinitionVolumesDockerVolumeConfigurationArgs) ElementType() reflect.Type {
	return taskDefinitionVolumesDockerVolumeConfigurationType
}

func (a TaskDefinitionVolumesDockerVolumeConfigurationArgs) ToTaskDefinitionVolumesDockerVolumeConfigurationOutput() TaskDefinitionVolumesDockerVolumeConfigurationOutput {
	return pulumi.ToOutput(a).(TaskDefinitionVolumesDockerVolumeConfigurationOutput)
}

func (a TaskDefinitionVolumesDockerVolumeConfigurationArgs) ToTaskDefinitionVolumesDockerVolumeConfigurationOutputWithContext(ctx context.Context) TaskDefinitionVolumesDockerVolumeConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, a).(TaskDefinitionVolumesDockerVolumeConfigurationOutput)
}

type TaskDefinitionVolumesDockerVolumeConfigurationOutput struct { *pulumi.OutputState }

// If this value is `true`, the Docker volume is created if it does not already exist. *Note*: This field is only used if the scope is `shared`.
func (o TaskDefinitionVolumesDockerVolumeConfigurationOutput) Autoprovision() pulumi.BoolOutput {
	return o.Apply(func(v TaskDefinitionVolumesDockerVolumeConfiguration) bool {
		if v.Autoprovision == nil { return *new(bool) } else { return *v.Autoprovision }
	}).(pulumi.BoolOutput)
}

// The Docker volume driver to use. The driver value must match the driver name provided by Docker because it is used for task placement.
func (o TaskDefinitionVolumesDockerVolumeConfigurationOutput) Driver() pulumi.StringOutput {
	return o.Apply(func(v TaskDefinitionVolumesDockerVolumeConfiguration) string {
		if v.Driver == nil { return *new(string) } else { return *v.Driver }
	}).(pulumi.StringOutput)
}

// A map of Docker driver specific options.
func (o TaskDefinitionVolumesDockerVolumeConfigurationOutput) DriverOpts() pulumi.StringMapOutput {
	return o.Apply(func(v TaskDefinitionVolumesDockerVolumeConfiguration) map[string]string {
		if v.DriverOpts == nil { return *new(map[string]string) } else { return *v.DriverOpts }
	}).(pulumi.StringMapOutput)
}

// A map of custom metadata to add to your Docker volume.
func (o TaskDefinitionVolumesDockerVolumeConfigurationOutput) Labels() pulumi.StringMapOutput {
	return o.Apply(func(v TaskDefinitionVolumesDockerVolumeConfiguration) map[string]string {
		if v.Labels == nil { return *new(map[string]string) } else { return *v.Labels }
	}).(pulumi.StringMapOutput)
}

// The scope for the Docker volume, which determines its lifecycle, either `task` or `shared`.  Docker volumes that are scoped to a `task` are automatically provisioned when the task starts and destroyed when the task stops. Docker volumes that are `scoped` as shared persist after the task stops.
func (o TaskDefinitionVolumesDockerVolumeConfigurationOutput) Scope() pulumi.StringOutput {
	return o.Apply(func(v TaskDefinitionVolumesDockerVolumeConfiguration) string {
		if v.Scope == nil { return *new(string) } else { return *v.Scope }
	}).(pulumi.StringOutput)
}

func (TaskDefinitionVolumesDockerVolumeConfigurationOutput) ElementType() reflect.Type {
	return taskDefinitionVolumesDockerVolumeConfigurationType
}

func (o TaskDefinitionVolumesDockerVolumeConfigurationOutput) ToTaskDefinitionVolumesDockerVolumeConfigurationOutput() TaskDefinitionVolumesDockerVolumeConfigurationOutput {
	return o
}

func (o TaskDefinitionVolumesDockerVolumeConfigurationOutput) ToTaskDefinitionVolumesDockerVolumeConfigurationOutputWithContext(ctx context.Context) TaskDefinitionVolumesDockerVolumeConfigurationOutput {
	return o
}

func init() { pulumi.RegisterOutputType(TaskDefinitionVolumesDockerVolumeConfigurationOutput{}) }

