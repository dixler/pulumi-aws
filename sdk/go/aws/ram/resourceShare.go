// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Resource Access Manager (RAM) Resource Share. To association principals with the share, see the [`ram.PrincipalAssociation` resource](https://www.terraform.io/docs/providers/aws/r/ram_principal_association.html). To associate resources with the share, see the [`ram.ResourceAssociation` resource](https://www.terraform.io/docs/providers/aws/r/ram_resource_association.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ram_resource_share.html.markdown.
type ResourceShare struct {
	pulumi.CustomResourceState

	// Indicates whether principals outside your organization can be associated with a resource share.
	AllowExternalPrincipals pulumi.BoolOutput `pulumi:"allowExternalPrincipals"`

	// The Amazon Resource Name (ARN) of the resource share.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The name of the resource share.
	Name pulumi.StringOutput `pulumi:"name"`

	// A mapping of tags to assign to the resource share.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewResourceShare registers a new resource with the given unique name, arguments, and options.
func NewResourceShare(ctx *pulumi.Context,
	name string, args *ResourceShareArgs, opts ...pulumi.ResourceOption) (*ResourceShare, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.AllowExternalPrincipals; i != nil { inputs["allowExternalPrincipals"] = i.ToBoolOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource ResourceShare
	err := ctx.RegisterResource("aws:ram/resourceShare:ResourceShare", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceShare gets an existing ResourceShare resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceShare(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceShareState, opts ...pulumi.ResourceOption) (*ResourceShare, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.AllowExternalPrincipals; i != nil { inputs["allowExternalPrincipals"] = i.ToBoolOutput() }
		if i := state.Arn; i != nil { inputs["arn"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource ResourceShare
	err := ctx.ReadResource("aws:ram/resourceShare:ResourceShare", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceShare resources.
type ResourceShareState struct {
	// Indicates whether principals outside your organization can be associated with a resource share.
	AllowExternalPrincipals pulumi.BoolInput `pulumi:"allowExternalPrincipals"`
	// The Amazon Resource Name (ARN) of the resource share.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The name of the resource share.
	Name pulumi.StringInput `pulumi:"name"`
	// A mapping of tags to assign to the resource share.
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a ResourceShare resource.
type ResourceShareArgs struct {
	// Indicates whether principals outside your organization can be associated with a resource share.
	AllowExternalPrincipals pulumi.BoolInput `pulumi:"allowExternalPrincipals"`
	// The name of the resource share.
	Name pulumi.StringInput `pulumi:"name"`
	// A mapping of tags to assign to the resource share.
	Tags pulumi.MapInput `pulumi:"tags"`
}
