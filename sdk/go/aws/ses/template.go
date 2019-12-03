// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to create a SES template.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ses_template.html.markdown.
type Template struct {
	pulumi.CustomResourceState

	// The HTML body of the email. Must be less than 500KB in size, including both the text and HTML parts.
	Html pulumi.StringOutput `pulumi:"html"`

	// The name of the template. Cannot exceed 64 characters. You will refer to this name when you send email.
	Name pulumi.StringOutput `pulumi:"name"`

	// The subject line of the email.
	Subject pulumi.StringOutput `pulumi:"subject"`

	// The email body that will be visible to recipients whose email clients do not display HTML. Must be less than 500KB in size, including both the text and HTML parts.
	Text pulumi.StringOutput `pulumi:"text"`
}

// NewTemplate registers a new resource with the given unique name, arguments, and options.
func NewTemplate(ctx *pulumi.Context,
	name string, args *TemplateArgs, opts ...pulumi.ResourceOption) (*Template, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Html; i != nil { inputs["html"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.Subject; i != nil { inputs["subject"] = i.ToStringOutput() }
		if i := args.Text; i != nil { inputs["text"] = i.ToStringOutput() }
	}
	var resource Template
	err := ctx.RegisterResource("aws:ses/template:Template", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTemplate gets an existing Template resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TemplateState, opts ...pulumi.ResourceOption) (*Template, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Html; i != nil { inputs["html"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.Subject; i != nil { inputs["subject"] = i.ToStringOutput() }
		if i := state.Text; i != nil { inputs["text"] = i.ToStringOutput() }
	}
	var resource Template
	err := ctx.ReadResource("aws:ses/template:Template", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Template resources.
type TemplateState struct {
	// The HTML body of the email. Must be less than 500KB in size, including both the text and HTML parts.
	Html pulumi.StringInput `pulumi:"html"`
	// The name of the template. Cannot exceed 64 characters. You will refer to this name when you send email.
	Name pulumi.StringInput `pulumi:"name"`
	// The subject line of the email.
	Subject pulumi.StringInput `pulumi:"subject"`
	// The email body that will be visible to recipients whose email clients do not display HTML. Must be less than 500KB in size, including both the text and HTML parts.
	Text pulumi.StringInput `pulumi:"text"`
}

// The set of arguments for constructing a Template resource.
type TemplateArgs struct {
	// The HTML body of the email. Must be less than 500KB in size, including both the text and HTML parts.
	Html pulumi.StringInput `pulumi:"html"`
	// The name of the template. Cannot exceed 64 characters. You will refer to this name when you send email.
	Name pulumi.StringInput `pulumi:"name"`
	// The subject line of the email.
	Subject pulumi.StringInput `pulumi:"subject"`
	// The email body that will be visible to recipients whose email clients do not display HTML. Must be less than 500KB in size, including both the text and HTML parts.
	Text pulumi.StringInput `pulumi:"text"`
}
