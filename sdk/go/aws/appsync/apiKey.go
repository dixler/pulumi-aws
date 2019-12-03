// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appsync

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an AppSync API Key.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appsync_api_key.html.markdown.
type ApiKey struct {
	pulumi.CustomResourceState

	// The ID of the associated AppSync API
	ApiId pulumi.StringOutput `pulumi:"apiId"`

	// The API key description. Defaults to "Managed by Pulumi".
	Description pulumi.StringOutput `pulumi:"description"`

	// RFC3339 string representation of the expiry date. Rounded down to nearest hour. By default, it is 7 days from the date of creation.
	Expires pulumi.StringOutput `pulumi:"expires"`

	// The API key
	Key pulumi.StringOutput `pulumi:"key"`
}

// NewApiKey registers a new resource with the given unique name, arguments, and options.
func NewApiKey(ctx *pulumi.Context,
	name string, args *ApiKeyArgs, opts ...pulumi.ResourceOption) (*ApiKey, error) {
	if args == nil || args.ApiId == nil {
		return nil, errors.New("missing required argument 'ApiId'")
	}
	inputs := map[string]pulumi.Input{}
	inputs["description"] = pulumi.Any("Managed by Pulumi")
	if args != nil {
		if i := args.ApiId; i != nil { inputs["apiId"] = i.ToStringOutput() }
		if i := args.Description; i != nil { inputs["description"] = i.ToStringOutput() }
		if i := args.Expires; i != nil { inputs["expires"] = i.ToStringOutput() }
	}
	var resource ApiKey
	err := ctx.RegisterResource("aws:appsync/apiKey:ApiKey", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiKey gets an existing ApiKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiKeyState, opts ...pulumi.ResourceOption) (*ApiKey, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.ApiId; i != nil { inputs["apiId"] = i.ToStringOutput() }
		if i := state.Description; i != nil { inputs["description"] = i.ToStringOutput() }
		if i := state.Expires; i != nil { inputs["expires"] = i.ToStringOutput() }
		if i := state.Key; i != nil { inputs["key"] = i.ToStringOutput() }
	}
	var resource ApiKey
	err := ctx.ReadResource("aws:appsync/apiKey:ApiKey", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiKey resources.
type ApiKeyState struct {
	// The ID of the associated AppSync API
	ApiId pulumi.StringInput `pulumi:"apiId"`
	// The API key description. Defaults to "Managed by Pulumi".
	Description pulumi.StringInput `pulumi:"description"`
	// RFC3339 string representation of the expiry date. Rounded down to nearest hour. By default, it is 7 days from the date of creation.
	Expires pulumi.StringInput `pulumi:"expires"`
	// The API key
	Key pulumi.StringInput `pulumi:"key"`
}

// The set of arguments for constructing a ApiKey resource.
type ApiKeyArgs struct {
	// The ID of the associated AppSync API
	ApiId pulumi.StringInput `pulumi:"apiId"`
	// The API key description. Defaults to "Managed by Pulumi".
	Description pulumi.StringInput `pulumi:"description"`
	// RFC3339 string representation of the expiry date. Rounded down to nearest hour. By default, it is 7 days from the date of creation.
	Expires pulumi.StringInput `pulumi:"expires"`
}
