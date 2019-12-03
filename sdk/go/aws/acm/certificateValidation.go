// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package acm

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// This resource represents a successful validation of an ACM certificate in concert
// with other resources.
// 
// Most commonly, this resource is used together with `route53.Record` and
// `acm.Certificate` to request a DNS validated certificate,
// deploy the required validation records and wait for validation to complete.
// 
// > **WARNING:** This resource implements a part of the validation workflow. It does not represent a real-world entity in AWS, therefore changing or deleting this resource on its own has no immediate effect.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/acm_certificate_validation.html.markdown.
type CertificateValidation struct {
	pulumi.CustomResourceState

	// The ARN of the certificate that is being validated.
	CertificateArn pulumi.StringOutput `pulumi:"certificateArn"`

	// List of FQDNs that implement the validation. Only valid for DNS validation method ACM certificates. If this is set, the resource can implement additional sanity checks and has an explicit dependency on the resource that is implementing the validation
	ValidationRecordFqdns pulumi.StringArrayOutput `pulumi:"validationRecordFqdns"`
}

// NewCertificateValidation registers a new resource with the given unique name, arguments, and options.
func NewCertificateValidation(ctx *pulumi.Context,
	name string, args *CertificateValidationArgs, opts ...pulumi.ResourceOption) (*CertificateValidation, error) {
	if args == nil || args.CertificateArn == nil {
		return nil, errors.New("missing required argument 'CertificateArn'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.CertificateArn; i != nil { inputs["certificateArn"] = i.ToStringOutput() }
		if i := args.ValidationRecordFqdns; i != nil { inputs["validationRecordFqdns"] = i.ToStringArrayOutput() }
	}
	var resource CertificateValidation
	err := ctx.RegisterResource("aws:acm/certificateValidation:CertificateValidation", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateValidation gets an existing CertificateValidation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateValidation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateValidationState, opts ...pulumi.ResourceOption) (*CertificateValidation, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.CertificateArn; i != nil { inputs["certificateArn"] = i.ToStringOutput() }
		if i := state.ValidationRecordFqdns; i != nil { inputs["validationRecordFqdns"] = i.ToStringArrayOutput() }
	}
	var resource CertificateValidation
	err := ctx.ReadResource("aws:acm/certificateValidation:CertificateValidation", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateValidation resources.
type CertificateValidationState struct {
	// The ARN of the certificate that is being validated.
	CertificateArn pulumi.StringInput `pulumi:"certificateArn"`
	// List of FQDNs that implement the validation. Only valid for DNS validation method ACM certificates. If this is set, the resource can implement additional sanity checks and has an explicit dependency on the resource that is implementing the validation
	ValidationRecordFqdns pulumi.StringArrayInput `pulumi:"validationRecordFqdns"`
}

// The set of arguments for constructing a CertificateValidation resource.
type CertificateValidationArgs struct {
	// The ARN of the certificate that is being validated.
	CertificateArn pulumi.StringInput `pulumi:"certificateArn"`
	// List of FQDNs that implement the validation. Only valid for DNS validation method ACM certificates. If this is set, the resource can implement additional sanity checks and has an explicit dependency on the resource that is implementing the validation
	ValidationRecordFqdns pulumi.StringArrayInput `pulumi:"validationRecordFqdns"`
}
