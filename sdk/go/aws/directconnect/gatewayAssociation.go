// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Associates a Direct Connect Gateway with a VGW or transit gateway.
// 
// To create a cross-account association, create an [`directconnect.GatewayAssociationProposal` resource](https://www.terraform.io/docs/providers/aws/r/dx_gateway_association_proposal.html)
// in the AWS account that owns the VGW or transit gateway and then accept the proposal in the AWS account that owns the Direct Connect Gateway
// by creating an `directconnect.GatewayAssociation` resource with the `proposalId` and `associatedGatewayOwnerAccountId` attributes set.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dx_gateway_association.html.markdown.
type GatewayAssociation struct {
	pulumi.CustomResourceState

	// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
	AllowedPrefixes pulumi.StringArrayOutput `pulumi:"allowedPrefixes"`

	// The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for single account Direct Connect gateway associations.
	AssociatedGatewayId pulumi.StringOutput `pulumi:"associatedGatewayId"`

	// The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for cross-account Direct Connect gateway associations.
	AssociatedGatewayOwnerAccountId pulumi.StringOutput `pulumi:"associatedGatewayOwnerAccountId"`

	// The type of the associated gateway, `transitGateway` or `virtualPrivateGateway`.
	AssociatedGatewayType pulumi.StringOutput `pulumi:"associatedGatewayType"`

	// The ID of the Direct Connect gateway association.
	DxGatewayAssociationId pulumi.StringOutput `pulumi:"dxGatewayAssociationId"`

	// The ID of the Direct Connect gateway.
	DxGatewayId pulumi.StringOutput `pulumi:"dxGatewayId"`

	// The ID of the AWS account that owns the Direct Connect gateway.
	DxGatewayOwnerAccountId pulumi.StringOutput `pulumi:"dxGatewayOwnerAccountId"`

	// The ID of the Direct Connect gateway association proposal.
	// Used for cross-account Direct Connect gateway associations.
	ProposalId pulumi.StringOutput `pulumi:"proposalId"`

	// *Deprecated:* Use `associatedGatewayId` instead. The ID of the VGW with which to associate the gateway.
	// Used for single account Direct Connect gateway associations.
	VpnGatewayId pulumi.StringOutput `pulumi:"vpnGatewayId"`
}

// NewGatewayAssociation registers a new resource with the given unique name, arguments, and options.
func NewGatewayAssociation(ctx *pulumi.Context,
	name string, args *GatewayAssociationArgs, opts ...pulumi.ResourceOption) (*GatewayAssociation, error) {
	if args == nil || args.DxGatewayId == nil {
		return nil, errors.New("missing required argument 'DxGatewayId'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.AllowedPrefixes; i != nil { inputs["allowedPrefixes"] = i.ToStringArrayOutput() }
		if i := args.AssociatedGatewayId; i != nil { inputs["associatedGatewayId"] = i.ToStringOutput() }
		if i := args.AssociatedGatewayOwnerAccountId; i != nil { inputs["associatedGatewayOwnerAccountId"] = i.ToStringOutput() }
		if i := args.DxGatewayId; i != nil { inputs["dxGatewayId"] = i.ToStringOutput() }
		if i := args.ProposalId; i != nil { inputs["proposalId"] = i.ToStringOutput() }
		if i := args.VpnGatewayId; i != nil { inputs["vpnGatewayId"] = i.ToStringOutput() }
	}
	var resource GatewayAssociation
	err := ctx.RegisterResource("aws:directconnect/gatewayAssociation:GatewayAssociation", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayAssociation gets an existing GatewayAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayAssociationState, opts ...pulumi.ResourceOption) (*GatewayAssociation, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.AllowedPrefixes; i != nil { inputs["allowedPrefixes"] = i.ToStringArrayOutput() }
		if i := state.AssociatedGatewayId; i != nil { inputs["associatedGatewayId"] = i.ToStringOutput() }
		if i := state.AssociatedGatewayOwnerAccountId; i != nil { inputs["associatedGatewayOwnerAccountId"] = i.ToStringOutput() }
		if i := state.AssociatedGatewayType; i != nil { inputs["associatedGatewayType"] = i.ToStringOutput() }
		if i := state.DxGatewayAssociationId; i != nil { inputs["dxGatewayAssociationId"] = i.ToStringOutput() }
		if i := state.DxGatewayId; i != nil { inputs["dxGatewayId"] = i.ToStringOutput() }
		if i := state.DxGatewayOwnerAccountId; i != nil { inputs["dxGatewayOwnerAccountId"] = i.ToStringOutput() }
		if i := state.ProposalId; i != nil { inputs["proposalId"] = i.ToStringOutput() }
		if i := state.VpnGatewayId; i != nil { inputs["vpnGatewayId"] = i.ToStringOutput() }
	}
	var resource GatewayAssociation
	err := ctx.ReadResource("aws:directconnect/gatewayAssociation:GatewayAssociation", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayAssociation resources.
type GatewayAssociationState struct {
	// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
	AllowedPrefixes pulumi.StringArrayInput `pulumi:"allowedPrefixes"`
	// The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for single account Direct Connect gateway associations.
	AssociatedGatewayId pulumi.StringInput `pulumi:"associatedGatewayId"`
	// The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for cross-account Direct Connect gateway associations.
	AssociatedGatewayOwnerAccountId pulumi.StringInput `pulumi:"associatedGatewayOwnerAccountId"`
	// The type of the associated gateway, `transitGateway` or `virtualPrivateGateway`.
	AssociatedGatewayType pulumi.StringInput `pulumi:"associatedGatewayType"`
	// The ID of the Direct Connect gateway association.
	DxGatewayAssociationId pulumi.StringInput `pulumi:"dxGatewayAssociationId"`
	// The ID of the Direct Connect gateway.
	DxGatewayId pulumi.StringInput `pulumi:"dxGatewayId"`
	// The ID of the AWS account that owns the Direct Connect gateway.
	DxGatewayOwnerAccountId pulumi.StringInput `pulumi:"dxGatewayOwnerAccountId"`
	// The ID of the Direct Connect gateway association proposal.
	// Used for cross-account Direct Connect gateway associations.
	ProposalId pulumi.StringInput `pulumi:"proposalId"`
	// *Deprecated:* Use `associatedGatewayId` instead. The ID of the VGW with which to associate the gateway.
	// Used for single account Direct Connect gateway associations.
	VpnGatewayId pulumi.StringInput `pulumi:"vpnGatewayId"`
}

// The set of arguments for constructing a GatewayAssociation resource.
type GatewayAssociationArgs struct {
	// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
	AllowedPrefixes pulumi.StringArrayInput `pulumi:"allowedPrefixes"`
	// The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for single account Direct Connect gateway associations.
	AssociatedGatewayId pulumi.StringInput `pulumi:"associatedGatewayId"`
	// The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for cross-account Direct Connect gateway associations.
	AssociatedGatewayOwnerAccountId pulumi.StringInput `pulumi:"associatedGatewayOwnerAccountId"`
	// The ID of the Direct Connect gateway.
	DxGatewayId pulumi.StringInput `pulumi:"dxGatewayId"`
	// The ID of the Direct Connect gateway association proposal.
	// Used for cross-account Direct Connect gateway associations.
	ProposalId pulumi.StringInput `pulumi:"proposalId"`
	// *Deprecated:* Use `associatedGatewayId` instead. The ID of the VGW with which to associate the gateway.
	// Used for single account Direct Connect gateway associations.
	VpnGatewayId pulumi.StringInput `pulumi:"vpnGatewayId"`
}
