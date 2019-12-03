// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicediscovery

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Service Discovery Private DNS Namespace resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/service_discovery_private_dns_namespace.html.markdown.
type PrivateDnsNamespace struct {
	pulumi.CustomResourceState

	// The ARN that Amazon Route 53 assigns to the namespace when you create it.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The description that you specify for the namespace when you create it.
	Description pulumi.StringOutput `pulumi:"description"`

	// The ID for the hosted zone that Amazon Route 53 creates when you create a namespace.
	HostedZone pulumi.StringOutput `pulumi:"hostedZone"`

	// The name of the namespace.
	Name pulumi.StringOutput `pulumi:"name"`

	// The ID of VPC that you want to associate the namespace with.
	Vpc pulumi.StringOutput `pulumi:"vpc"`
}

// NewPrivateDnsNamespace registers a new resource with the given unique name, arguments, and options.
func NewPrivateDnsNamespace(ctx *pulumi.Context,
	name string, args *PrivateDnsNamespaceArgs, opts ...pulumi.ResourceOption) (*PrivateDnsNamespace, error) {
	if args == nil || args.Vpc == nil {
		return nil, errors.New("missing required argument 'Vpc'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Description; i != nil { inputs["description"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.Vpc; i != nil { inputs["vpc"] = i.ToStringOutput() }
	}
	var resource PrivateDnsNamespace
	err := ctx.RegisterResource("aws:servicediscovery/privateDnsNamespace:PrivateDnsNamespace", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateDnsNamespace gets an existing PrivateDnsNamespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateDnsNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateDnsNamespaceState, opts ...pulumi.ResourceOption) (*PrivateDnsNamespace, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Arn; i != nil { inputs["arn"] = i.ToStringOutput() }
		if i := state.Description; i != nil { inputs["description"] = i.ToStringOutput() }
		if i := state.HostedZone; i != nil { inputs["hostedZone"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.Vpc; i != nil { inputs["vpc"] = i.ToStringOutput() }
	}
	var resource PrivateDnsNamespace
	err := ctx.ReadResource("aws:servicediscovery/privateDnsNamespace:PrivateDnsNamespace", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateDnsNamespace resources.
type PrivateDnsNamespaceState struct {
	// The ARN that Amazon Route 53 assigns to the namespace when you create it.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The description that you specify for the namespace when you create it.
	Description pulumi.StringInput `pulumi:"description"`
	// The ID for the hosted zone that Amazon Route 53 creates when you create a namespace.
	HostedZone pulumi.StringInput `pulumi:"hostedZone"`
	// The name of the namespace.
	Name pulumi.StringInput `pulumi:"name"`
	// The ID of VPC that you want to associate the namespace with.
	Vpc pulumi.StringInput `pulumi:"vpc"`
}

// The set of arguments for constructing a PrivateDnsNamespace resource.
type PrivateDnsNamespaceArgs struct {
	// The description that you specify for the namespace when you create it.
	Description pulumi.StringInput `pulumi:"description"`
	// The name of the namespace.
	Name pulumi.StringInput `pulumi:"name"`
	// The ID of VPC that you want to associate the namespace with.
	Vpc pulumi.StringInput `pulumi:"vpc"`
}
