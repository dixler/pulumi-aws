// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Registers a custom domain name for use with AWS API Gateway. Additional information about this functionality
// can be found in the [API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html).
// 
// This resource just establishes ownership of and the TLS settings for
// a particular domain name. An API can be attached to a particular path
// under the registered domain name using
// the `apigateway.BasePathMapping` resource.
// 
// API Gateway domains can be defined as either 'edge-optimized' or 'regional'.  In an edge-optimized configuration,
// API Gateway internally creates and manages a CloudFront distribution to route requests on the given hostname. In
// addition to this resource it's necessary to create a DNS record corresponding to the given domain name which is an alias
// (either Route53 alias or traditional CNAME) to the Cloudfront domain name exported in the `cloudfrontDomainName`
// attribute.
// 
// In a regional configuration, API Gateway does not create a CloudFront distribution to route requests to the API, though
// a distribution can be created if needed. In either case, it is necessary to create a DNS record corresponding to the
// given domain name which is an alias (either Route53 alias or traditional CNAME) to the regional domain name exported in
// the `regionalDomainName` attribute.
// 
// > **Note:** API Gateway requires the use of AWS Certificate Manager (ACM) certificates instead of Identity and Access Management (IAM) certificates in regions that support ACM. Regions that support ACM can be found in the [Regions and Endpoints Documentation](https://docs.aws.amazon.com/general/latest/gr/rande.html#acm_region). To import an existing private key and certificate into ACM or request an ACM certificate, see the [`acm.Certificate` resource](https://www.terraform.io/docs/providers/aws/r/acm_certificate.html).
// 
// > **Note:** The `apigateway.DomainName` resource expects dependency on the `acm.CertificateValidation` as 
// only verified certificates can be used. This can be made either explicitly by adding the 
// `dependsOn = [aws_acm_certificate_validation.cert]` attribute. Or implicitly by referring certificate ARN 
// from the validation resource where it will be available after the resource creation: 
// `regionalCertificateArn = aws_acm_certificate_validation.cert.certificate_arn`.
// 
// > **Note:** All arguments including the private key will be stored in the raw state as plain-text.
// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_domain_name.html.markdown.
type DomainName struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN)
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when an edge-optimized domain name is desired. Conflicts with `certificateName`, `certificateBody`, `certificateChain`, `certificatePrivateKey`, `regionalCertificateArn`, and `regionalCertificateName`.
	CertificateArn pulumi.StringOutput `pulumi:"certificateArn"`

	// The certificate issued for the domain name
	// being registered, in PEM format. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`, `regionalCertificateArn`, and
	// `regionalCertificateName`.
	CertificateBody pulumi.StringOutput `pulumi:"certificateBody"`

	// The certificate for the CA that issued the
	// certificate, along with any intermediate CA certificates required to
	// create an unbroken chain to a certificate trusted by the intended API clients. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`,
	// `regionalCertificateArn`, and `regionalCertificateName`.
	CertificateChain pulumi.StringOutput `pulumi:"certificateChain"`

	// The unique name to use when registering this
	// certificate as an IAM server certificate. Conflicts with `certificateArn`, `regionalCertificateArn`, and
	// `regionalCertificateName`. Required if `certificateArn` is not set.
	CertificateName pulumi.StringOutput `pulumi:"certificateName"`

	// The private key associated with the
	// domain certificate given in `certificateBody`. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`, `regionalCertificateArn`, and `regionalCertificateName`.
	CertificatePrivateKey pulumi.StringOutput `pulumi:"certificatePrivateKey"`

	// The upload date associated with the domain certificate.
	CertificateUploadDate pulumi.StringOutput `pulumi:"certificateUploadDate"`

	// The hostname created by Cloudfront to represent
	// the distribution that implements this domain name mapping.
	CloudfrontDomainName pulumi.StringOutput `pulumi:"cloudfrontDomainName"`

	// For convenience, the hosted zone ID (`Z2FDTNDATAQYW2`)
	// that can be used to create a Route53 alias record for the distribution.
	CloudfrontZoneId pulumi.StringOutput `pulumi:"cloudfrontZoneId"`

	// The fully-qualified domain name to register
	DomainName pulumi.StringOutput `pulumi:"domainName"`

	// Configuration block defining API endpoint information including type. Defined below.
	EndpointConfiguration DomainNameEndpointConfigurationOutput `pulumi:"endpointConfiguration"`

	// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when a regional domain name is desired. Conflicts with `certificateArn`, `certificateName`, `certificateBody`, `certificateChain`, and `certificatePrivateKey`.
	RegionalCertificateArn pulumi.StringOutput `pulumi:"regionalCertificateArn"`

	// The user-friendly name of the certificate that will be used by regional endpoint for this domain name. Conflicts with `certificateArn`, `certificateName`, `certificateBody`, `certificateChain`, and
	// `certificatePrivateKey`.
	RegionalCertificateName pulumi.StringOutput `pulumi:"regionalCertificateName"`

	// The hostname for the custom domain's regional endpoint.
	RegionalDomainName pulumi.StringOutput `pulumi:"regionalDomainName"`

	// The hosted zone ID that can be used to create a Route53 alias record for the regional endpoint.
	RegionalZoneId pulumi.StringOutput `pulumi:"regionalZoneId"`

	// The Transport Layer Security (TLS) version + cipher suite for this DomainName. The valid values are `TLS_1_0` and `TLS_1_2`. Must be configured to perform drift detection.
	SecurityPolicy pulumi.StringOutput `pulumi:"securityPolicy"`

	// Key-value mapping of resource tags
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewDomainName registers a new resource with the given unique name, arguments, and options.
func NewDomainName(ctx *pulumi.Context,
	name string, args *DomainNameArgs, opts ...pulumi.ResourceOption) (*DomainName, error) {
	if args == nil || args.DomainName == nil {
		return nil, errors.New("missing required argument 'DomainName'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.CertificateArn; i != nil { inputs["certificateArn"] = i.ToStringOutput() }
		if i := args.CertificateBody; i != nil { inputs["certificateBody"] = i.ToStringOutput() }
		if i := args.CertificateChain; i != nil { inputs["certificateChain"] = i.ToStringOutput() }
		if i := args.CertificateName; i != nil { inputs["certificateName"] = i.ToStringOutput() }
		if i := args.CertificatePrivateKey; i != nil { inputs["certificatePrivateKey"] = i.ToStringOutput() }
		if i := args.DomainName; i != nil { inputs["domainName"] = i.ToStringOutput() }
		if i := args.EndpointConfiguration; i != nil { inputs["endpointConfiguration"] = i.ToDomainNameEndpointConfigurationOutput() }
		if i := args.RegionalCertificateArn; i != nil { inputs["regionalCertificateArn"] = i.ToStringOutput() }
		if i := args.RegionalCertificateName; i != nil { inputs["regionalCertificateName"] = i.ToStringOutput() }
		if i := args.SecurityPolicy; i != nil { inputs["securityPolicy"] = i.ToStringOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource DomainName
	err := ctx.RegisterResource("aws:apigateway/domainName:DomainName", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainName gets an existing DomainName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainNameState, opts ...pulumi.ResourceOption) (*DomainName, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Arn; i != nil { inputs["arn"] = i.ToStringOutput() }
		if i := state.CertificateArn; i != nil { inputs["certificateArn"] = i.ToStringOutput() }
		if i := state.CertificateBody; i != nil { inputs["certificateBody"] = i.ToStringOutput() }
		if i := state.CertificateChain; i != nil { inputs["certificateChain"] = i.ToStringOutput() }
		if i := state.CertificateName; i != nil { inputs["certificateName"] = i.ToStringOutput() }
		if i := state.CertificatePrivateKey; i != nil { inputs["certificatePrivateKey"] = i.ToStringOutput() }
		if i := state.CertificateUploadDate; i != nil { inputs["certificateUploadDate"] = i.ToStringOutput() }
		if i := state.CloudfrontDomainName; i != nil { inputs["cloudfrontDomainName"] = i.ToStringOutput() }
		if i := state.CloudfrontZoneId; i != nil { inputs["cloudfrontZoneId"] = i.ToStringOutput() }
		if i := state.DomainName; i != nil { inputs["domainName"] = i.ToStringOutput() }
		if i := state.EndpointConfiguration; i != nil { inputs["endpointConfiguration"] = i.ToDomainNameEndpointConfigurationOutput() }
		if i := state.RegionalCertificateArn; i != nil { inputs["regionalCertificateArn"] = i.ToStringOutput() }
		if i := state.RegionalCertificateName; i != nil { inputs["regionalCertificateName"] = i.ToStringOutput() }
		if i := state.RegionalDomainName; i != nil { inputs["regionalDomainName"] = i.ToStringOutput() }
		if i := state.RegionalZoneId; i != nil { inputs["regionalZoneId"] = i.ToStringOutput() }
		if i := state.SecurityPolicy; i != nil { inputs["securityPolicy"] = i.ToStringOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
	}
	var resource DomainName
	err := ctx.ReadResource("aws:apigateway/domainName:DomainName", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainName resources.
type DomainNameState struct {
	// Amazon Resource Name (ARN)
	Arn pulumi.StringInput `pulumi:"arn"`
	// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when an edge-optimized domain name is desired. Conflicts with `certificateName`, `certificateBody`, `certificateChain`, `certificatePrivateKey`, `regionalCertificateArn`, and `regionalCertificateName`.
	CertificateArn pulumi.StringInput `pulumi:"certificateArn"`
	// The certificate issued for the domain name
	// being registered, in PEM format. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`, `regionalCertificateArn`, and
	// `regionalCertificateName`.
	CertificateBody pulumi.StringInput `pulumi:"certificateBody"`
	// The certificate for the CA that issued the
	// certificate, along with any intermediate CA certificates required to
	// create an unbroken chain to a certificate trusted by the intended API clients. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`,
	// `regionalCertificateArn`, and `regionalCertificateName`.
	CertificateChain pulumi.StringInput `pulumi:"certificateChain"`
	// The unique name to use when registering this
	// certificate as an IAM server certificate. Conflicts with `certificateArn`, `regionalCertificateArn`, and
	// `regionalCertificateName`. Required if `certificateArn` is not set.
	CertificateName pulumi.StringInput `pulumi:"certificateName"`
	// The private key associated with the
	// domain certificate given in `certificateBody`. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`, `regionalCertificateArn`, and `regionalCertificateName`.
	CertificatePrivateKey pulumi.StringInput `pulumi:"certificatePrivateKey"`
	// The upload date associated with the domain certificate.
	CertificateUploadDate pulumi.StringInput `pulumi:"certificateUploadDate"`
	// The hostname created by Cloudfront to represent
	// the distribution that implements this domain name mapping.
	CloudfrontDomainName pulumi.StringInput `pulumi:"cloudfrontDomainName"`
	// For convenience, the hosted zone ID (`Z2FDTNDATAQYW2`)
	// that can be used to create a Route53 alias record for the distribution.
	CloudfrontZoneId pulumi.StringInput `pulumi:"cloudfrontZoneId"`
	// The fully-qualified domain name to register
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// Configuration block defining API endpoint information including type. Defined below.
	EndpointConfiguration DomainNameEndpointConfigurationInput `pulumi:"endpointConfiguration"`
	// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when a regional domain name is desired. Conflicts with `certificateArn`, `certificateName`, `certificateBody`, `certificateChain`, and `certificatePrivateKey`.
	RegionalCertificateArn pulumi.StringInput `pulumi:"regionalCertificateArn"`
	// The user-friendly name of the certificate that will be used by regional endpoint for this domain name. Conflicts with `certificateArn`, `certificateName`, `certificateBody`, `certificateChain`, and
	// `certificatePrivateKey`.
	RegionalCertificateName pulumi.StringInput `pulumi:"regionalCertificateName"`
	// The hostname for the custom domain's regional endpoint.
	RegionalDomainName pulumi.StringInput `pulumi:"regionalDomainName"`
	// The hosted zone ID that can be used to create a Route53 alias record for the regional endpoint.
	RegionalZoneId pulumi.StringInput `pulumi:"regionalZoneId"`
	// The Transport Layer Security (TLS) version + cipher suite for this DomainName. The valid values are `TLS_1_0` and `TLS_1_2`. Must be configured to perform drift detection.
	SecurityPolicy pulumi.StringInput `pulumi:"securityPolicy"`
	// Key-value mapping of resource tags
	Tags pulumi.MapInput `pulumi:"tags"`
}

// The set of arguments for constructing a DomainName resource.
type DomainNameArgs struct {
	// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when an edge-optimized domain name is desired. Conflicts with `certificateName`, `certificateBody`, `certificateChain`, `certificatePrivateKey`, `regionalCertificateArn`, and `regionalCertificateName`.
	CertificateArn pulumi.StringInput `pulumi:"certificateArn"`
	// The certificate issued for the domain name
	// being registered, in PEM format. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`, `regionalCertificateArn`, and
	// `regionalCertificateName`.
	CertificateBody pulumi.StringInput `pulumi:"certificateBody"`
	// The certificate for the CA that issued the
	// certificate, along with any intermediate CA certificates required to
	// create an unbroken chain to a certificate trusted by the intended API clients. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`,
	// `regionalCertificateArn`, and `regionalCertificateName`.
	CertificateChain pulumi.StringInput `pulumi:"certificateChain"`
	// The unique name to use when registering this
	// certificate as an IAM server certificate. Conflicts with `certificateArn`, `regionalCertificateArn`, and
	// `regionalCertificateName`. Required if `certificateArn` is not set.
	CertificateName pulumi.StringInput `pulumi:"certificateName"`
	// The private key associated with the
	// domain certificate given in `certificateBody`. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificateArn`, `regionalCertificateArn`, and `regionalCertificateName`.
	CertificatePrivateKey pulumi.StringInput `pulumi:"certificatePrivateKey"`
	// The fully-qualified domain name to register
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// Configuration block defining API endpoint information including type. Defined below.
	EndpointConfiguration DomainNameEndpointConfigurationInput `pulumi:"endpointConfiguration"`
	// The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when a regional domain name is desired. Conflicts with `certificateArn`, `certificateName`, `certificateBody`, `certificateChain`, and `certificatePrivateKey`.
	RegionalCertificateArn pulumi.StringInput `pulumi:"regionalCertificateArn"`
	// The user-friendly name of the certificate that will be used by regional endpoint for this domain name. Conflicts with `certificateArn`, `certificateName`, `certificateBody`, `certificateChain`, and
	// `certificatePrivateKey`.
	RegionalCertificateName pulumi.StringInput `pulumi:"regionalCertificateName"`
	// The Transport Layer Security (TLS) version + cipher suite for this DomainName. The valid values are `TLS_1_0` and `TLS_1_2`. Must be configured to perform drift detection.
	SecurityPolicy pulumi.StringInput `pulumi:"securityPolicy"`
	// Key-value mapping of resource tags
	Tags pulumi.MapInput `pulumi:"tags"`
}
type DomainNameEndpointConfiguration struct {
	// A list of endpoint types. This resource currently only supports managing a single value. Valid values: `EDGE` or `REGIONAL`. If unspecified, defaults to `EDGE`. Must be declared as `REGIONAL` in non-Commercial partitions. Refer to the [documentation](https://docs.aws.amazon.com/apigateway/latest/developerguide/create-regional-api.html) for more information on the difference between edge-optimized and regional APIs.
	Types string `pulumi:"types"`
}
var domainNameEndpointConfigurationType = reflect.TypeOf((*DomainNameEndpointConfiguration)(nil)).Elem()

type DomainNameEndpointConfigurationInput interface {
	pulumi.Input

	ToDomainNameEndpointConfigurationOutput() DomainNameEndpointConfigurationOutput
	ToDomainNameEndpointConfigurationOutputWithContext(ctx context.Context) DomainNameEndpointConfigurationOutput
}

type DomainNameEndpointConfigurationArgs struct {
	// A list of endpoint types. This resource currently only supports managing a single value. Valid values: `EDGE` or `REGIONAL`. If unspecified, defaults to `EDGE`. Must be declared as `REGIONAL` in non-Commercial partitions. Refer to the [documentation](https://docs.aws.amazon.com/apigateway/latest/developerguide/create-regional-api.html) for more information on the difference between edge-optimized and regional APIs.
	Types pulumi.StringInput `pulumi:"types"`
}

func (DomainNameEndpointConfigurationArgs) ElementType() reflect.Type {
	return domainNameEndpointConfigurationType
}

func (a DomainNameEndpointConfigurationArgs) ToDomainNameEndpointConfigurationOutput() DomainNameEndpointConfigurationOutput {
	return pulumi.ToOutput(a).(DomainNameEndpointConfigurationOutput)
}

func (a DomainNameEndpointConfigurationArgs) ToDomainNameEndpointConfigurationOutputWithContext(ctx context.Context) DomainNameEndpointConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, a).(DomainNameEndpointConfigurationOutput)
}

type DomainNameEndpointConfigurationOutput struct { *pulumi.OutputState }

// A list of endpoint types. This resource currently only supports managing a single value. Valid values: `EDGE` or `REGIONAL`. If unspecified, defaults to `EDGE`. Must be declared as `REGIONAL` in non-Commercial partitions. Refer to the [documentation](https://docs.aws.amazon.com/apigateway/latest/developerguide/create-regional-api.html) for more information on the difference between edge-optimized and regional APIs.
func (o DomainNameEndpointConfigurationOutput) Types() pulumi.StringOutput {
	return o.Apply(func(v DomainNameEndpointConfiguration) string {
		return v.Types
	}).(pulumi.StringOutput)
}

func (DomainNameEndpointConfigurationOutput) ElementType() reflect.Type {
	return domainNameEndpointConfigurationType
}

func (o DomainNameEndpointConfigurationOutput) ToDomainNameEndpointConfigurationOutput() DomainNameEndpointConfigurationOutput {
	return o
}

func (o DomainNameEndpointConfigurationOutput) ToDomainNameEndpointConfigurationOutputWithContext(ctx context.Context) DomainNameEndpointConfigurationOutput {
	return o
}

func init() { pulumi.RegisterOutputType(DomainNameEndpointConfigurationOutput{}) }

