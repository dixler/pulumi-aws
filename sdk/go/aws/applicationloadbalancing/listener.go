// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package applicationloadbalancing

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Load Balancer Listener resource.
// 
// > **Note:** `aws_alb_listener` is known as `aws_lb_listener`. The functionality is identical.
type Listener struct {
	s *pulumi.ResourceState
}

// NewListener registers a new resource with the given unique name, arguments, and options.
func NewListener(ctx *pulumi.Context,
	name string, args *ListenerArgs, opts ...pulumi.ResourceOpt) (*Listener, error) {
	if args == nil || args.DefaultActions == nil {
		return nil, errors.New("missing required argument 'DefaultActions'")
	}
	if args == nil || args.LoadBalancerArn == nil {
		return nil, errors.New("missing required argument 'LoadBalancerArn'")
	}
	if args == nil || args.Port == nil {
		return nil, errors.New("missing required argument 'Port'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["certificateArn"] = nil
		inputs["defaultActions"] = nil
		inputs["loadBalancerArn"] = nil
		inputs["port"] = nil
		inputs["protocol"] = nil
		inputs["sslPolicy"] = nil
	} else {
		inputs["certificateArn"] = args.CertificateArn
		inputs["defaultActions"] = args.DefaultActions
		inputs["loadBalancerArn"] = args.LoadBalancerArn
		inputs["port"] = args.Port
		inputs["protocol"] = args.Protocol
		inputs["sslPolicy"] = args.SslPolicy
	}
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:applicationloadbalancing/listener:Listener", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Listener{s: s}, nil
}

// GetListener gets an existing Listener resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetListener(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ListenerState, opts ...pulumi.ResourceOpt) (*Listener, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["certificateArn"] = state.CertificateArn
		inputs["defaultActions"] = state.DefaultActions
		inputs["loadBalancerArn"] = state.LoadBalancerArn
		inputs["port"] = state.Port
		inputs["protocol"] = state.Protocol
		inputs["sslPolicy"] = state.SslPolicy
	}
	s, err := ctx.ReadResource("aws:applicationloadbalancing/listener:Listener", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Listener{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Listener) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Listener) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The ARN of the listener (matches `id`)
func (r *Listener) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the [`aws_lb_listener_certificate` resource](https://www.terraform.io/docs/providers/aws/r/lb_listener_certificate.html).
func (r *Listener) CertificateArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["certificateArn"])
}

// An Action block. Action blocks are documented below.
func (r *Listener) DefaultActions() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["defaultActions"])
}

// The ARN of the load balancer.
func (r *Listener) LoadBalancerArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["loadBalancerArn"])
}

// The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
func (r *Listener) Port() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["port"])
}

// The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
func (r *Listener) Protocol() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["protocol"])
}

// The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
func (r *Listener) SslPolicy() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sslPolicy"])
}

// Input properties used for looking up and filtering Listener resources.
type ListenerState struct {
	// The ARN of the listener (matches `id`)
	Arn interface{}
	// The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the [`aws_lb_listener_certificate` resource](https://www.terraform.io/docs/providers/aws/r/lb_listener_certificate.html).
	CertificateArn interface{}
	// An Action block. Action blocks are documented below.
	DefaultActions interface{}
	// The ARN of the load balancer.
	LoadBalancerArn interface{}
	// The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
	Port interface{}
	// The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
	Protocol interface{}
	// The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
	SslPolicy interface{}
}

// The set of arguments for constructing a Listener resource.
type ListenerArgs struct {
	// The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the [`aws_lb_listener_certificate` resource](https://www.terraform.io/docs/providers/aws/r/lb_listener_certificate.html).
	CertificateArn interface{}
	// An Action block. Action blocks are documented below.
	DefaultActions interface{}
	// The ARN of the load balancer.
	LoadBalancerArn interface{}
	// The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
	Port interface{}
	// The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
	Protocol interface{}
	// The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
	SslPolicy interface{}
}
