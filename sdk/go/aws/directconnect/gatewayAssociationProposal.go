// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Direct Connect Gateway Association Proposal, typically for enabling cross-account associations. For single account associations, see the [`directconnect.GatewayAssociation` resource](https://www.terraform.io/docs/providers/aws/r/dx_gateway_association.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dx_gateway_association_proposal.html.markdown.
type GatewayAssociationProposal struct {
	pulumi.CustomResourceState

	// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
	AllowedPrefixes pulumi.StringArrayOutput `pulumi:"allowedPrefixes"`

	// The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
	AssociatedGatewayId pulumi.StringOutput `pulumi:"associatedGatewayId"`

	// The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
	AssociatedGatewayOwnerAccountId pulumi.StringOutput `pulumi:"associatedGatewayOwnerAccountId"`

	// The type of the associated gateway, `transitGateway` or `virtualPrivateGateway`.
	AssociatedGatewayType pulumi.StringOutput `pulumi:"associatedGatewayType"`

	// Direct Connect Gateway identifier.
	DxGatewayId pulumi.StringOutput `pulumi:"dxGatewayId"`

	// AWS Account identifier of the Direct Connect Gateway's owner.
	DxGatewayOwnerAccountId pulumi.StringOutput `pulumi:"dxGatewayOwnerAccountId"`

	// *Deprecated:* Use `associatedGatewayId` instead. Virtual Gateway identifier to associate with the Direct Connect Gateway.
	VpnGatewayId pulumi.StringOutput `pulumi:"vpnGatewayId"`
}

// NewGatewayAssociationProposal registers a new resource with the given unique name, arguments, and options.
func NewGatewayAssociationProposal(ctx *pulumi.Context,
	name string, args *GatewayAssociationProposalArgs, opts ...pulumi.ResourceOption) (*GatewayAssociationProposal, error) {
	if args == nil || args.DxGatewayId == nil {
		return nil, errors.New("missing required argument 'DxGatewayId'")
	}
	if args == nil || args.DxGatewayOwnerAccountId == nil {
		return nil, errors.New("missing required argument 'DxGatewayOwnerAccountId'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.AllowedPrefixes; i != nil { inputs["allowedPrefixes"] = i.ToStringArrayOutput() }
		if i := args.AssociatedGatewayId; i != nil { inputs["associatedGatewayId"] = i.ToStringOutput() }
		if i := args.DxGatewayId; i != nil { inputs["dxGatewayId"] = i.ToStringOutput() }
		if i := args.DxGatewayOwnerAccountId; i != nil { inputs["dxGatewayOwnerAccountId"] = i.ToStringOutput() }
		if i := args.VpnGatewayId; i != nil { inputs["vpnGatewayId"] = i.ToStringOutput() }
	}
	var resource GatewayAssociationProposal
	err := ctx.RegisterResource("aws:directconnect/gatewayAssociationProposal:GatewayAssociationProposal", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayAssociationProposal gets an existing GatewayAssociationProposal resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayAssociationProposal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayAssociationProposalState, opts ...pulumi.ResourceOption) (*GatewayAssociationProposal, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.AllowedPrefixes; i != nil { inputs["allowedPrefixes"] = i.ToStringArrayOutput() }
		if i := state.AssociatedGatewayId; i != nil { inputs["associatedGatewayId"] = i.ToStringOutput() }
		if i := state.AssociatedGatewayOwnerAccountId; i != nil { inputs["associatedGatewayOwnerAccountId"] = i.ToStringOutput() }
		if i := state.AssociatedGatewayType; i != nil { inputs["associatedGatewayType"] = i.ToStringOutput() }
		if i := state.DxGatewayId; i != nil { inputs["dxGatewayId"] = i.ToStringOutput() }
		if i := state.DxGatewayOwnerAccountId; i != nil { inputs["dxGatewayOwnerAccountId"] = i.ToStringOutput() }
		if i := state.VpnGatewayId; i != nil { inputs["vpnGatewayId"] = i.ToStringOutput() }
	}
	var resource GatewayAssociationProposal
	err := ctx.ReadResource("aws:directconnect/gatewayAssociationProposal:GatewayAssociationProposal", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayAssociationProposal resources.
type GatewayAssociationProposalState struct {
	// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
	AllowedPrefixes pulumi.StringArrayInput `pulumi:"allowedPrefixes"`
	// The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
	AssociatedGatewayId pulumi.StringInput `pulumi:"associatedGatewayId"`
	// The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
	AssociatedGatewayOwnerAccountId pulumi.StringInput `pulumi:"associatedGatewayOwnerAccountId"`
	// The type of the associated gateway, `transitGateway` or `virtualPrivateGateway`.
	AssociatedGatewayType pulumi.StringInput `pulumi:"associatedGatewayType"`
	// Direct Connect Gateway identifier.
	DxGatewayId pulumi.StringInput `pulumi:"dxGatewayId"`
	// AWS Account identifier of the Direct Connect Gateway's owner.
	DxGatewayOwnerAccountId pulumi.StringInput `pulumi:"dxGatewayOwnerAccountId"`
	// *Deprecated:* Use `associatedGatewayId` instead. Virtual Gateway identifier to associate with the Direct Connect Gateway.
	VpnGatewayId pulumi.StringInput `pulumi:"vpnGatewayId"`
}

// The set of arguments for constructing a GatewayAssociationProposal resource.
type GatewayAssociationProposalArgs struct {
	// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
	AllowedPrefixes pulumi.StringArrayInput `pulumi:"allowedPrefixes"`
	// The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
	AssociatedGatewayId pulumi.StringInput `pulumi:"associatedGatewayId"`
	// Direct Connect Gateway identifier.
	DxGatewayId pulumi.StringInput `pulumi:"dxGatewayId"`
	// AWS Account identifier of the Direct Connect Gateway's owner.
	DxGatewayOwnerAccountId pulumi.StringInput `pulumi:"dxGatewayOwnerAccountId"`
	// *Deprecated:* Use `associatedGatewayId` instead. Virtual Gateway identifier to associate with the Direct Connect Gateway.
	VpnGatewayId pulumi.StringInput `pulumi:"vpnGatewayId"`
}
