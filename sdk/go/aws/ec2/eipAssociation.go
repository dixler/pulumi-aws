// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an AWS EIP Association as a top level resource, to associate and
// disassociate Elastic IPs from AWS Instances and Network Interfaces.
// 
// > **NOTE:** Do not use this resource to associate an EIP to `lb.LoadBalancer` or `ec2.NatGateway` resources. Instead use the `allocationId` available in those resources to allow AWS to manage the association, otherwise you will see `AuthFailure` errors.
// 
// > **NOTE:** `ec2.EipAssociation` is useful in scenarios where EIPs are either
// pre-existing or distributed to customers or users and therefore cannot be changed.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/eip_association.html.markdown.
type EipAssociation struct {
	pulumi.CustomResourceState

	// The allocation ID. This is required for EC2-VPC.
	AllocationId pulumi.StringOutput `pulumi:"allocationId"`

	// Whether to allow an Elastic IP to
	// be re-associated. Defaults to `true` in VPC.
	AllowReassociation pulumi.BoolOutput `pulumi:"allowReassociation"`

	// The ID of the instance. This is required for
	// EC2-Classic. For EC2-VPC, you can specify either the instance ID or the
	// network interface ID, but not both. The operation fails if you specify an
	// instance ID unless exactly one network interface is attached.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`

	// The ID of the network interface. If the
	// instance has more than one network interface, you must specify a network
	// interface ID.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`

	// The primary or secondary private IP address
	// to associate with the Elastic IP address. If no private IP address is
	// specified, the Elastic IP address is associated with the primary private IP
	// address.
	PrivateIpAddress pulumi.StringOutput `pulumi:"privateIpAddress"`

	// The Elastic IP address. This is required for EC2-Classic.
	PublicIp pulumi.StringOutput `pulumi:"publicIp"`
}

// NewEipAssociation registers a new resource with the given unique name, arguments, and options.
func NewEipAssociation(ctx *pulumi.Context,
	name string, args *EipAssociationArgs, opts ...pulumi.ResourceOption) (*EipAssociation, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.AllocationId; i != nil { inputs["allocationId"] = i.ToStringOutput() }
		if i := args.AllowReassociation; i != nil { inputs["allowReassociation"] = i.ToBoolOutput() }
		if i := args.InstanceId; i != nil { inputs["instanceId"] = i.ToStringOutput() }
		if i := args.NetworkInterfaceId; i != nil { inputs["networkInterfaceId"] = i.ToStringOutput() }
		if i := args.PrivateIpAddress; i != nil { inputs["privateIpAddress"] = i.ToStringOutput() }
		if i := args.PublicIp; i != nil { inputs["publicIp"] = i.ToStringOutput() }
	}
	var resource EipAssociation
	err := ctx.RegisterResource("aws:ec2/eipAssociation:EipAssociation", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEipAssociation gets an existing EipAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEipAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EipAssociationState, opts ...pulumi.ResourceOption) (*EipAssociation, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.AllocationId; i != nil { inputs["allocationId"] = i.ToStringOutput() }
		if i := state.AllowReassociation; i != nil { inputs["allowReassociation"] = i.ToBoolOutput() }
		if i := state.InstanceId; i != nil { inputs["instanceId"] = i.ToStringOutput() }
		if i := state.NetworkInterfaceId; i != nil { inputs["networkInterfaceId"] = i.ToStringOutput() }
		if i := state.PrivateIpAddress; i != nil { inputs["privateIpAddress"] = i.ToStringOutput() }
		if i := state.PublicIp; i != nil { inputs["publicIp"] = i.ToStringOutput() }
	}
	var resource EipAssociation
	err := ctx.ReadResource("aws:ec2/eipAssociation:EipAssociation", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EipAssociation resources.
type EipAssociationState struct {
	// The allocation ID. This is required for EC2-VPC.
	AllocationId pulumi.StringInput `pulumi:"allocationId"`
	// Whether to allow an Elastic IP to
	// be re-associated. Defaults to `true` in VPC.
	AllowReassociation pulumi.BoolInput `pulumi:"allowReassociation"`
	// The ID of the instance. This is required for
	// EC2-Classic. For EC2-VPC, you can specify either the instance ID or the
	// network interface ID, but not both. The operation fails if you specify an
	// instance ID unless exactly one network interface is attached.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// The ID of the network interface. If the
	// instance has more than one network interface, you must specify a network
	// interface ID.
	NetworkInterfaceId pulumi.StringInput `pulumi:"networkInterfaceId"`
	// The primary or secondary private IP address
	// to associate with the Elastic IP address. If no private IP address is
	// specified, the Elastic IP address is associated with the primary private IP
	// address.
	PrivateIpAddress pulumi.StringInput `pulumi:"privateIpAddress"`
	// The Elastic IP address. This is required for EC2-Classic.
	PublicIp pulumi.StringInput `pulumi:"publicIp"`
}

// The set of arguments for constructing a EipAssociation resource.
type EipAssociationArgs struct {
	// The allocation ID. This is required for EC2-VPC.
	AllocationId pulumi.StringInput `pulumi:"allocationId"`
	// Whether to allow an Elastic IP to
	// be re-associated. Defaults to `true` in VPC.
	AllowReassociation pulumi.BoolInput `pulumi:"allowReassociation"`
	// The ID of the instance. This is required for
	// EC2-Classic. For EC2-VPC, you can specify either the instance ID or the
	// network interface ID, but not both. The operation fails if you specify an
	// instance ID unless exactly one network interface is attached.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// The ID of the network interface. If the
	// instance has more than one network interface, you must specify a network
	// interface ID.
	NetworkInterfaceId pulumi.StringInput `pulumi:"networkInterfaceId"`
	// The primary or secondary private IP address
	// to associate with the Elastic IP address. If no private IP address is
	// specified, the Elastic IP address is associated with the primary private IP
	// address.
	PrivateIpAddress pulumi.StringInput `pulumi:"privateIpAddress"`
	// The Elastic IP address. This is required for EC2-Classic.
	PublicIp pulumi.StringInput `pulumi:"publicIp"`
}
