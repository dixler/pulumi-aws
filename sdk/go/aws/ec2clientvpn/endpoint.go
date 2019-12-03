// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2clientvpn

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an AWS Client VPN endpoint for OpenVPN clients. For more information on usage, please see the
// [AWS Client VPN Administrator's Guide](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/what-is.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_client_vpn_endpoint.html.markdown.
type Endpoint struct {
	pulumi.CustomResourceState

	// Information about the authentication method to be used to authenticate clients.
	AuthenticationOptions EndpointAuthenticationOptionsOutput `pulumi:"authenticationOptions"`

	// The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
	ClientCidrBlock pulumi.StringOutput `pulumi:"clientCidrBlock"`

	// Information about the client connection logging options.
	ConnectionLogOptions EndpointConnectionLogOptionsOutput `pulumi:"connectionLogOptions"`

	// Name of the repository.
	Description pulumi.StringOutput `pulumi:"description"`

	// The DNS name to be used by clients when establishing their VPN session.
	DnsName pulumi.StringOutput `pulumi:"dnsName"`

	// Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the VPC that is to be associated with Client VPN endpoint is used as the DNS server.
	DnsServers pulumi.StringArrayOutput `pulumi:"dnsServers"`

	// The ARN of the ACM server certificate.
	ServerCertificateArn pulumi.StringOutput `pulumi:"serverCertificateArn"`

	// Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
	SplitTunnel pulumi.BoolOutput `pulumi:"splitTunnel"`

	// The current state of the Client VPN endpoint.
	Status pulumi.StringOutput `pulumi:"status"`

	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// The transport protocol to be used by the VPN session. Default value is `udp`.
	TransportProtocol pulumi.StringOutput `pulumi:"transportProtocol"`
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil || args.AuthenticationOptions == nil {
		return nil, errors.New("missing required argument 'AuthenticationOptions'")
	}
	if args == nil || args.ClientCidrBlock == nil {
		return nil, errors.New("missing required argument 'ClientCidrBlock'")
	}
	if args == nil || args.ConnectionLogOptions == nil {
		return nil, errors.New("missing required argument 'ConnectionLogOptions'")
	}
	if args == nil || args.ServerCertificateArn == nil {
		return nil, errors.New("missing required argument 'ServerCertificateArn'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.AuthenticationOptions; i != nil { inputs["authenticationOptions"] = i.ToEndpointAuthenticationOptionsOutput() }
		if i := args.ClientCidrBlock; i != nil { inputs["clientCidrBlock"] = i.ToStringOutput() }
		if i := args.ConnectionLogOptions; i != nil { inputs["connectionLogOptions"] = i.ToEndpointConnectionLogOptionsOutput() }
		if i := args.Description; i != nil { inputs["description"] = i.ToStringOutput() }
		if i := args.DnsServers; i != nil { inputs["dnsServers"] = i.ToStringArrayOutput() }
		if i := args.ServerCertificateArn; i != nil { inputs["serverCertificateArn"] = i.ToStringOutput() }
		if i := args.SplitTunnel; i != nil { inputs["splitTunnel"] = i.ToBoolOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := args.TransportProtocol; i != nil { inputs["transportProtocol"] = i.ToStringOutput() }
	}
	var resource Endpoint
	err := ctx.RegisterResource("aws:ec2clientvpn/endpoint:Endpoint", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.AuthenticationOptions; i != nil { inputs["authenticationOptions"] = i.ToEndpointAuthenticationOptionsOutput() }
		if i := state.ClientCidrBlock; i != nil { inputs["clientCidrBlock"] = i.ToStringOutput() }
		if i := state.ConnectionLogOptions; i != nil { inputs["connectionLogOptions"] = i.ToEndpointConnectionLogOptionsOutput() }
		if i := state.Description; i != nil { inputs["description"] = i.ToStringOutput() }
		if i := state.DnsName; i != nil { inputs["dnsName"] = i.ToStringOutput() }
		if i := state.DnsServers; i != nil { inputs["dnsServers"] = i.ToStringArrayOutput() }
		if i := state.ServerCertificateArn; i != nil { inputs["serverCertificateArn"] = i.ToStringOutput() }
		if i := state.SplitTunnel; i != nil { inputs["splitTunnel"] = i.ToBoolOutput() }
		if i := state.Status; i != nil { inputs["status"] = i.ToStringOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := state.TransportProtocol; i != nil { inputs["transportProtocol"] = i.ToStringOutput() }
	}
	var resource Endpoint
	err := ctx.ReadResource("aws:ec2clientvpn/endpoint:Endpoint", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type EndpointState struct {
	// Information about the authentication method to be used to authenticate clients.
	AuthenticationOptions EndpointAuthenticationOptionsInput `pulumi:"authenticationOptions"`
	// The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
	ClientCidrBlock pulumi.StringInput `pulumi:"clientCidrBlock"`
	// Information about the client connection logging options.
	ConnectionLogOptions EndpointConnectionLogOptionsInput `pulumi:"connectionLogOptions"`
	// Name of the repository.
	Description pulumi.StringInput `pulumi:"description"`
	// The DNS name to be used by clients when establishing their VPN session.
	DnsName pulumi.StringInput `pulumi:"dnsName"`
	// Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the VPC that is to be associated with Client VPN endpoint is used as the DNS server.
	DnsServers pulumi.StringArrayInput `pulumi:"dnsServers"`
	// The ARN of the ACM server certificate.
	ServerCertificateArn pulumi.StringInput `pulumi:"serverCertificateArn"`
	// Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
	SplitTunnel pulumi.BoolInput `pulumi:"splitTunnel"`
	// The current state of the Client VPN endpoint.
	Status pulumi.StringInput `pulumi:"status"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The transport protocol to be used by the VPN session. Default value is `udp`.
	TransportProtocol pulumi.StringInput `pulumi:"transportProtocol"`
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	// Information about the authentication method to be used to authenticate clients.
	AuthenticationOptions EndpointAuthenticationOptionsInput `pulumi:"authenticationOptions"`
	// The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. The CIDR block should be /22 or greater.
	ClientCidrBlock pulumi.StringInput `pulumi:"clientCidrBlock"`
	// Information about the client connection logging options.
	ConnectionLogOptions EndpointConnectionLogOptionsInput `pulumi:"connectionLogOptions"`
	// Name of the repository.
	Description pulumi.StringInput `pulumi:"description"`
	// Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address of the VPC that is to be associated with Client VPN endpoint is used as the DNS server.
	DnsServers pulumi.StringArrayInput `pulumi:"dnsServers"`
	// The ARN of the ACM server certificate.
	ServerCertificateArn pulumi.StringInput `pulumi:"serverCertificateArn"`
	// Indicates whether split-tunnel is enabled on VPN endpoint. Default value is `false`.
	SplitTunnel pulumi.BoolInput `pulumi:"splitTunnel"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput `pulumi:"tags"`
	// The transport protocol to be used by the VPN session. Default value is `udp`.
	TransportProtocol pulumi.StringInput `pulumi:"transportProtocol"`
}
type EndpointAuthenticationOptions struct {
	// The ID of the Active Directory to be used for authentication if type is `directory-service-authentication`.
	ActiveDirectoryId *string `pulumi:"activeDirectoryId"`
	// The ARN of the client certificate. The certificate must be signed by a certificate authority (CA) and it must be provisioned in AWS Certificate Manager (ACM). Only necessary when type is set to `certificate-authentication`.
	RootCertificateChainArn *string `pulumi:"rootCertificateChainArn"`
	// The type of client authentication to be used. Specify `certificate-authentication` to use certificate-based authentication, or `directory-service-authentication` to use Active Directory authentication.
	Type string `pulumi:"type"`
}
var endpointAuthenticationOptionsType = reflect.TypeOf((*EndpointAuthenticationOptions)(nil)).Elem()

type EndpointAuthenticationOptionsInput interface {
	pulumi.Input

	ToEndpointAuthenticationOptionsOutput() EndpointAuthenticationOptionsOutput
	ToEndpointAuthenticationOptionsOutputWithContext(ctx context.Context) EndpointAuthenticationOptionsOutput
}

type EndpointAuthenticationOptionsArgs struct {
	// The ID of the Active Directory to be used for authentication if type is `directory-service-authentication`.
	ActiveDirectoryId pulumi.StringInput `pulumi:"activeDirectoryId"`
	// The ARN of the client certificate. The certificate must be signed by a certificate authority (CA) and it must be provisioned in AWS Certificate Manager (ACM). Only necessary when type is set to `certificate-authentication`.
	RootCertificateChainArn pulumi.StringInput `pulumi:"rootCertificateChainArn"`
	// The type of client authentication to be used. Specify `certificate-authentication` to use certificate-based authentication, or `directory-service-authentication` to use Active Directory authentication.
	Type pulumi.StringInput `pulumi:"type"`
}

func (EndpointAuthenticationOptionsArgs) ElementType() reflect.Type {
	return endpointAuthenticationOptionsType
}

func (a EndpointAuthenticationOptionsArgs) ToEndpointAuthenticationOptionsOutput() EndpointAuthenticationOptionsOutput {
	return pulumi.ToOutput(a).(EndpointAuthenticationOptionsOutput)
}

func (a EndpointAuthenticationOptionsArgs) ToEndpointAuthenticationOptionsOutputWithContext(ctx context.Context) EndpointAuthenticationOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, a).(EndpointAuthenticationOptionsOutput)
}

type EndpointAuthenticationOptionsOutput struct { *pulumi.OutputState }

// The ID of the Active Directory to be used for authentication if type is `directory-service-authentication`.
func (o EndpointAuthenticationOptionsOutput) ActiveDirectoryId() pulumi.StringOutput {
	return o.Apply(func(v EndpointAuthenticationOptions) string {
		if v.ActiveDirectoryId == nil { return *new(string) } else { return *v.ActiveDirectoryId }
	}).(pulumi.StringOutput)
}

// The ARN of the client certificate. The certificate must be signed by a certificate authority (CA) and it must be provisioned in AWS Certificate Manager (ACM). Only necessary when type is set to `certificate-authentication`.
func (o EndpointAuthenticationOptionsOutput) RootCertificateChainArn() pulumi.StringOutput {
	return o.Apply(func(v EndpointAuthenticationOptions) string {
		if v.RootCertificateChainArn == nil { return *new(string) } else { return *v.RootCertificateChainArn }
	}).(pulumi.StringOutput)
}

// The type of client authentication to be used. Specify `certificate-authentication` to use certificate-based authentication, or `directory-service-authentication` to use Active Directory authentication.
func (o EndpointAuthenticationOptionsOutput) Type() pulumi.StringOutput {
	return o.Apply(func(v EndpointAuthenticationOptions) string {
		return v.Type
	}).(pulumi.StringOutput)
}

func (EndpointAuthenticationOptionsOutput) ElementType() reflect.Type {
	return endpointAuthenticationOptionsType
}

func (o EndpointAuthenticationOptionsOutput) ToEndpointAuthenticationOptionsOutput() EndpointAuthenticationOptionsOutput {
	return o
}

func (o EndpointAuthenticationOptionsOutput) ToEndpointAuthenticationOptionsOutputWithContext(ctx context.Context) EndpointAuthenticationOptionsOutput {
	return o
}

func init() { pulumi.RegisterOutputType(EndpointAuthenticationOptionsOutput{}) }

type EndpointConnectionLogOptions struct {
	// The name of the CloudWatch Logs log group.
	CloudwatchLogGroup *string `pulumi:"cloudwatchLogGroup"`
	// The name of the CloudWatch Logs log stream to which the connection data is published.
	CloudwatchLogStream *string `pulumi:"cloudwatchLogStream"`
	// Indicates whether connection logging is enabled.
	Enabled bool `pulumi:"enabled"`
}
var endpointConnectionLogOptionsType = reflect.TypeOf((*EndpointConnectionLogOptions)(nil)).Elem()

type EndpointConnectionLogOptionsInput interface {
	pulumi.Input

	ToEndpointConnectionLogOptionsOutput() EndpointConnectionLogOptionsOutput
	ToEndpointConnectionLogOptionsOutputWithContext(ctx context.Context) EndpointConnectionLogOptionsOutput
}

type EndpointConnectionLogOptionsArgs struct {
	// The name of the CloudWatch Logs log group.
	CloudwatchLogGroup pulumi.StringInput `pulumi:"cloudwatchLogGroup"`
	// The name of the CloudWatch Logs log stream to which the connection data is published.
	CloudwatchLogStream pulumi.StringInput `pulumi:"cloudwatchLogStream"`
	// Indicates whether connection logging is enabled.
	Enabled pulumi.BoolInput `pulumi:"enabled"`
}

func (EndpointConnectionLogOptionsArgs) ElementType() reflect.Type {
	return endpointConnectionLogOptionsType
}

func (a EndpointConnectionLogOptionsArgs) ToEndpointConnectionLogOptionsOutput() EndpointConnectionLogOptionsOutput {
	return pulumi.ToOutput(a).(EndpointConnectionLogOptionsOutput)
}

func (a EndpointConnectionLogOptionsArgs) ToEndpointConnectionLogOptionsOutputWithContext(ctx context.Context) EndpointConnectionLogOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, a).(EndpointConnectionLogOptionsOutput)
}

type EndpointConnectionLogOptionsOutput struct { *pulumi.OutputState }

// The name of the CloudWatch Logs log group.
func (o EndpointConnectionLogOptionsOutput) CloudwatchLogGroup() pulumi.StringOutput {
	return o.Apply(func(v EndpointConnectionLogOptions) string {
		if v.CloudwatchLogGroup == nil { return *new(string) } else { return *v.CloudwatchLogGroup }
	}).(pulumi.StringOutput)
}

// The name of the CloudWatch Logs log stream to which the connection data is published.
func (o EndpointConnectionLogOptionsOutput) CloudwatchLogStream() pulumi.StringOutput {
	return o.Apply(func(v EndpointConnectionLogOptions) string {
		if v.CloudwatchLogStream == nil { return *new(string) } else { return *v.CloudwatchLogStream }
	}).(pulumi.StringOutput)
}

// Indicates whether connection logging is enabled.
func (o EndpointConnectionLogOptionsOutput) Enabled() pulumi.BoolOutput {
	return o.Apply(func(v EndpointConnectionLogOptions) bool {
		return v.Enabled
	}).(pulumi.BoolOutput)
}

func (EndpointConnectionLogOptionsOutput) ElementType() reflect.Type {
	return endpointConnectionLogOptionsType
}

func (o EndpointConnectionLogOptionsOutput) ToEndpointConnectionLogOptionsOutput() EndpointConnectionLogOptionsOutput {
	return o
}

func (o EndpointConnectionLogOptionsOutput) ToEndpointConnectionLogOptionsOutputWithContext(ctx context.Context) EndpointConnectionLogOptionsOutput {
	return o
}

func init() { pulumi.RegisterOutputType(EndpointConnectionLogOptionsOutput{}) }

