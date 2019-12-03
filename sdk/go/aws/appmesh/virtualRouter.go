// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an AWS App Mesh virtual router resource.
// 
// ## Breaking Changes
// 
// Because of backward incompatible API changes (read [here](https://github.com/awslabs/aws-app-mesh-examples/issues/92) and [here](https://github.com/awslabs/aws-app-mesh-examples/issues/94)), `appmesh.VirtualRouter` resource definitions created with provider versions earlier than v2.3.0 will need to be modified:
// 
// * Remove service `serviceNames` from the `spec` argument.
// AWS has created a `appmesh.VirtualService` resource for each of service names.
// These resource can be imported using `import`.
// 
// * Add a `listener` configuration block to the `spec` argument.
// 
// The state associated with existing resources will automatically be migrated.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appmesh_virtual_router.html.markdown.
type VirtualRouter struct {
	pulumi.CustomResourceState

	// The ARN of the virtual router.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The creation date of the virtual router.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`

	// The last update date of the virtual router.
	LastUpdatedDate pulumi.StringOutput `pulumi:"lastUpdatedDate"`

	// The name of the service mesh in which to create the virtual router.
	MeshName pulumi.StringOutput `pulumi:"meshName"`

	// The name to use for the virtual router.
	Name pulumi.StringOutput `pulumi:"name"`

	// The virtual router specification to apply.
	Spec VirtualRouterSpecOutput `pulumi:"spec"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewVirtualRouter registers a new resource with the given unique name, arguments, and options.
func NewVirtualRouter(ctx *pulumi.Context,
	name string, args *VirtualRouterArgs, opts ...pulumi.ResourceOption) (*VirtualRouter, error) {
	if args == nil || args.MeshName == nil {
		return nil, errors.New("missing required argument 'MeshName'")
	}
	if args == nil || args.Spec == nil {
		return nil, errors.New("missing required argument 'Spec'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.MeshName; i != nil { inputs["meshName"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.Spec; i != nil { inputs["spec"] = i.ToVirtualRouterSpecOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource VirtualRouter
	err := ctx.RegisterResource("aws:appmesh/virtualRouter:VirtualRouter", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualRouter gets an existing VirtualRouter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualRouter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualRouterState, opts ...pulumi.ResourceOption) (*VirtualRouter, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Arn; i != nil { inputs["arn"] = i.ToStringOutput() }
		if i := state.CreatedDate; i != nil { inputs["createdDate"] = i.ToStringOutput() }
		if i := state.LastUpdatedDate; i != nil { inputs["lastUpdatedDate"] = i.ToStringOutput() }
		if i := state.MeshName; i != nil { inputs["meshName"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.Spec; i != nil { inputs["spec"] = i.ToVirtualRouterSpecOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource VirtualRouter
	err := ctx.ReadResource("aws:appmesh/virtualRouter:VirtualRouter", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualRouter resources.
type VirtualRouterState struct {
	// The ARN of the virtual router.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The creation date of the virtual router.
	CreatedDate pulumi.StringInput `pulumi:"createdDate"`
	// The last update date of the virtual router.
	LastUpdatedDate pulumi.StringInput `pulumi:"lastUpdatedDate"`
	// The name of the service mesh in which to create the virtual router.
	MeshName pulumi.StringInput `pulumi:"meshName"`
	// The name to use for the virtual router.
	Name pulumi.StringInput `pulumi:"name"`
	// The virtual router specification to apply.
	Spec VirtualRouterSpecInput `pulumi:"spec"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a VirtualRouter resource.
type VirtualRouterArgs struct {
	// The name of the service mesh in which to create the virtual router.
	MeshName pulumi.StringInput `pulumi:"meshName"`
	// The name to use for the virtual router.
	Name pulumi.StringInput `pulumi:"name"`
	// The virtual router specification to apply.
	Spec VirtualRouterSpecInput `pulumi:"spec"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
}
type VirtualRouterSpec struct {
	// The listeners that the virtual router is expected to receive inbound traffic from.
	// Currently only one listener is supported per virtual router.
	Listener VirtualRouterSpecListener `pulumi:"listener"`
}
var virtualRouterSpecType = reflect.TypeOf((*VirtualRouterSpec)(nil)).Elem()

type VirtualRouterSpecInput interface {
	pulumi.Input

	ToVirtualRouterSpecOutput() VirtualRouterSpecOutput
	ToVirtualRouterSpecOutputWithContext(ctx context.Context) VirtualRouterSpecOutput
}

type VirtualRouterSpecArgs struct {
	// The listeners that the virtual router is expected to receive inbound traffic from.
	// Currently only one listener is supported per virtual router.
	Listener VirtualRouterSpecListenerInput `pulumi:"listener"`
}

func (VirtualRouterSpecArgs) ElementType() reflect.Type {
	return virtualRouterSpecType
}

func (a VirtualRouterSpecArgs) ToVirtualRouterSpecOutput() VirtualRouterSpecOutput {
	return pulumi.ToOutput(a).(VirtualRouterSpecOutput)
}

func (a VirtualRouterSpecArgs) ToVirtualRouterSpecOutputWithContext(ctx context.Context) VirtualRouterSpecOutput {
	return pulumi.ToOutputWithContext(ctx, a).(VirtualRouterSpecOutput)
}

type VirtualRouterSpecOutput struct { *pulumi.OutputState }

// The listeners that the virtual router is expected to receive inbound traffic from.
// Currently only one listener is supported per virtual router.
func (o VirtualRouterSpecOutput) Listener() VirtualRouterSpecListenerOutput {
	return o.Apply(func(v VirtualRouterSpec) VirtualRouterSpecListener {
		return v.Listener
	}).(VirtualRouterSpecListenerOutput)
}

func (VirtualRouterSpecOutput) ElementType() reflect.Type {
	return virtualRouterSpecType
}

func (o VirtualRouterSpecOutput) ToVirtualRouterSpecOutput() VirtualRouterSpecOutput {
	return o
}

func (o VirtualRouterSpecOutput) ToVirtualRouterSpecOutputWithContext(ctx context.Context) VirtualRouterSpecOutput {
	return o
}

func init() { pulumi.RegisterOutputType(VirtualRouterSpecOutput{}) }

type VirtualRouterSpecListener struct {
	// The port mapping information for the listener.
	PortMapping VirtualRouterSpecListenerPortMapping `pulumi:"portMapping"`
}
var virtualRouterSpecListenerType = reflect.TypeOf((*VirtualRouterSpecListener)(nil)).Elem()

type VirtualRouterSpecListenerInput interface {
	pulumi.Input

	ToVirtualRouterSpecListenerOutput() VirtualRouterSpecListenerOutput
	ToVirtualRouterSpecListenerOutputWithContext(ctx context.Context) VirtualRouterSpecListenerOutput
}

type VirtualRouterSpecListenerArgs struct {
	// The port mapping information for the listener.
	PortMapping VirtualRouterSpecListenerPortMappingInput `pulumi:"portMapping"`
}

func (VirtualRouterSpecListenerArgs) ElementType() reflect.Type {
	return virtualRouterSpecListenerType
}

func (a VirtualRouterSpecListenerArgs) ToVirtualRouterSpecListenerOutput() VirtualRouterSpecListenerOutput {
	return pulumi.ToOutput(a).(VirtualRouterSpecListenerOutput)
}

func (a VirtualRouterSpecListenerArgs) ToVirtualRouterSpecListenerOutputWithContext(ctx context.Context) VirtualRouterSpecListenerOutput {
	return pulumi.ToOutputWithContext(ctx, a).(VirtualRouterSpecListenerOutput)
}

type VirtualRouterSpecListenerOutput struct { *pulumi.OutputState }

// The port mapping information for the listener.
func (o VirtualRouterSpecListenerOutput) PortMapping() VirtualRouterSpecListenerPortMappingOutput {
	return o.Apply(func(v VirtualRouterSpecListener) VirtualRouterSpecListenerPortMapping {
		return v.PortMapping
	}).(VirtualRouterSpecListenerPortMappingOutput)
}

func (VirtualRouterSpecListenerOutput) ElementType() reflect.Type {
	return virtualRouterSpecListenerType
}

func (o VirtualRouterSpecListenerOutput) ToVirtualRouterSpecListenerOutput() VirtualRouterSpecListenerOutput {
	return o
}

func (o VirtualRouterSpecListenerOutput) ToVirtualRouterSpecListenerOutputWithContext(ctx context.Context) VirtualRouterSpecListenerOutput {
	return o
}

func init() { pulumi.RegisterOutputType(VirtualRouterSpecListenerOutput{}) }

type VirtualRouterSpecListenerPortMapping struct {
	// The port used for the port mapping.
	Port int `pulumi:"port"`
	// The protocol used for the port mapping. Valid values are `http` and `tcp`.
	Protocol string `pulumi:"protocol"`
}
var virtualRouterSpecListenerPortMappingType = reflect.TypeOf((*VirtualRouterSpecListenerPortMapping)(nil)).Elem()

type VirtualRouterSpecListenerPortMappingInput interface {
	pulumi.Input

	ToVirtualRouterSpecListenerPortMappingOutput() VirtualRouterSpecListenerPortMappingOutput
	ToVirtualRouterSpecListenerPortMappingOutputWithContext(ctx context.Context) VirtualRouterSpecListenerPortMappingOutput
}

type VirtualRouterSpecListenerPortMappingArgs struct {
	// The port used for the port mapping.
	Port pulumi.IntInput `pulumi:"port"`
	// The protocol used for the port mapping. Valid values are `http` and `tcp`.
	Protocol pulumi.StringInput `pulumi:"protocol"`
}

func (VirtualRouterSpecListenerPortMappingArgs) ElementType() reflect.Type {
	return virtualRouterSpecListenerPortMappingType
}

func (a VirtualRouterSpecListenerPortMappingArgs) ToVirtualRouterSpecListenerPortMappingOutput() VirtualRouterSpecListenerPortMappingOutput {
	return pulumi.ToOutput(a).(VirtualRouterSpecListenerPortMappingOutput)
}

func (a VirtualRouterSpecListenerPortMappingArgs) ToVirtualRouterSpecListenerPortMappingOutputWithContext(ctx context.Context) VirtualRouterSpecListenerPortMappingOutput {
	return pulumi.ToOutputWithContext(ctx, a).(VirtualRouterSpecListenerPortMappingOutput)
}

type VirtualRouterSpecListenerPortMappingOutput struct { *pulumi.OutputState }

// The port used for the port mapping.
func (o VirtualRouterSpecListenerPortMappingOutput) Port() pulumi.IntOutput {
	return o.Apply(func(v VirtualRouterSpecListenerPortMapping) int {
		return v.Port
	}).(pulumi.IntOutput)
}

// The protocol used for the port mapping. Valid values are `http` and `tcp`.
func (o VirtualRouterSpecListenerPortMappingOutput) Protocol() pulumi.StringOutput {
	return o.Apply(func(v VirtualRouterSpecListenerPortMapping) string {
		return v.Protocol
	}).(pulumi.StringOutput)
}

func (VirtualRouterSpecListenerPortMappingOutput) ElementType() reflect.Type {
	return virtualRouterSpecListenerPortMappingType
}

func (o VirtualRouterSpecListenerPortMappingOutput) ToVirtualRouterSpecListenerPortMappingOutput() VirtualRouterSpecListenerPortMappingOutput {
	return o
}

func (o VirtualRouterSpecListenerPortMappingOutput) ToVirtualRouterSpecListenerPortMappingOutputWithContext(ctx context.Context) VirtualRouterSpecListenerPortMappingOutput {
	return o
}

func init() { pulumi.RegisterOutputType(VirtualRouterSpecListenerPortMappingOutput{}) }

