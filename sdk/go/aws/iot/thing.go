// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates and manages an AWS IoT Thing.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iot_thing.html.markdown.
type Thing struct {
	pulumi.CustomResourceState

	// The ARN of the thing.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// Map of attributes of the thing.
	Attributes pulumi.MapOutput `pulumi:"attributes"`

	// The default client ID.
	DefaultClientId pulumi.StringOutput `pulumi:"defaultClientId"`

	// The name of the thing.
	Name pulumi.StringOutput `pulumi:"name"`

	// The thing type name.
	ThingTypeName pulumi.StringOutput `pulumi:"thingTypeName"`

	// The current version of the thing record in the registry.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewThing registers a new resource with the given unique name, arguments, and options.
func NewThing(ctx *pulumi.Context,
	name string, args *ThingArgs, opts ...pulumi.ResourceOption) (*Thing, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Attributes; i != nil { inputs["attributes"] = i.ToMapOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.ThingTypeName; i != nil { inputs["thingTypeName"] = i.ToStringOutput() }
	}
	var resource Thing
	err := ctx.RegisterResource("aws:iot/thing:Thing", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetThing gets an existing Thing resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetThing(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ThingState, opts ...pulumi.ResourceOption) (*Thing, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Arn; i != nil { inputs["arn"] = i.ToStringOutput() }
		if i := state.Attributes; i != nil { inputs["attributes"] = i.ToMapOutput() }
		if i := state.DefaultClientId; i != nil { inputs["defaultClientId"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.ThingTypeName; i != nil { inputs["thingTypeName"] = i.ToStringOutput() }
		if i := state.Version; i != nil { inputs["version"] = i.ToIntOutput() }
	}
	var resource Thing
	err := ctx.ReadResource("aws:iot/thing:Thing", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Thing resources.
type ThingState struct {
	// The ARN of the thing.
	Arn pulumi.StringInput `pulumi:"arn"`
	// Map of attributes of the thing.
	Attributes pulumi.MapInput `pulumi:"attributes"`
	// The default client ID.
	DefaultClientId pulumi.StringInput `pulumi:"defaultClientId"`
	// The name of the thing.
	Name pulumi.StringInput `pulumi:"name"`
	// The thing type name.
	ThingTypeName pulumi.StringInput `pulumi:"thingTypeName"`
	// The current version of the thing record in the registry.
	Version pulumi.IntInput `pulumi:"version"`
}

// The set of arguments for constructing a Thing resource.
type ThingArgs struct {
	// Map of attributes of the thing.
	Attributes pulumi.MapInput `pulumi:"attributes"`
	// The name of the thing.
	Name pulumi.StringInput `pulumi:"name"`
	// The thing type name.
	ThingTypeName pulumi.StringInput `pulumi:"thingTypeName"`
}
