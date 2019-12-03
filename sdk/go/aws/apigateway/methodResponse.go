// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an HTTP Method Response for an API Gateway Resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_method_response.html.markdown.
type MethodResponse struct {
	pulumi.CustomResourceState

	// The HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
	HttpMethod pulumi.StringOutput `pulumi:"httpMethod"`

	// The API resource ID
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`

	// A map of the API models used for the response's content type
	ResponseModels pulumi.StringMapOutput `pulumi:"responseModels"`

	// A map of response parameters that can be sent to the caller.
	// For example: `responseParameters = { "method.response.header.X-Some-Header" = true }`
	// would define that the header `X-Some-Header` can be provided on the response.
	ResponseParameters pulumi.BoolMapOutput `pulumi:"responseParameters"`

	// The ID of the associated REST API
	RestApi pulumi.StringOutput `pulumi:"restApi"`

	// The HTTP status code
	StatusCode pulumi.StringOutput `pulumi:"statusCode"`
}

// NewMethodResponse registers a new resource with the given unique name, arguments, and options.
func NewMethodResponse(ctx *pulumi.Context,
	name string, args *MethodResponseArgs, opts ...pulumi.ResourceOption) (*MethodResponse, error) {
	if args == nil || args.HttpMethod == nil {
		return nil, errors.New("missing required argument 'HttpMethod'")
	}
	if args == nil || args.ResourceId == nil {
		return nil, errors.New("missing required argument 'ResourceId'")
	}
	if args == nil || args.RestApi == nil {
		return nil, errors.New("missing required argument 'RestApi'")
	}
	if args == nil || args.StatusCode == nil {
		return nil, errors.New("missing required argument 'StatusCode'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.HttpMethod; i != nil { inputs["httpMethod"] = i.ToStringOutput() }
		if i := args.ResourceId; i != nil { inputs["resourceId"] = i.ToStringOutput() }
		if i := args.ResponseModels; i != nil { inputs["responseModels"] = i.ToStringMapOutput() }
		if i := args.ResponseParameters; i != nil { inputs["responseParameters"] = i.ToBoolMapOutput() }
		if i := args.RestApi; i != nil { inputs["restApi"] = i.ToStringOutput() }
		if i := args.StatusCode; i != nil { inputs["statusCode"] = i.ToStringOutput() }
	}
	var resource MethodResponse
	err := ctx.RegisterResource("aws:apigateway/methodResponse:MethodResponse", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMethodResponse gets an existing MethodResponse resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMethodResponse(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MethodResponseState, opts ...pulumi.ResourceOption) (*MethodResponse, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.HttpMethod; i != nil { inputs["httpMethod"] = i.ToStringOutput() }
		if i := state.ResourceId; i != nil { inputs["resourceId"] = i.ToStringOutput() }
		if i := state.ResponseModels; i != nil { inputs["responseModels"] = i.ToStringMapOutput() }
		if i := state.ResponseParameters; i != nil { inputs["responseParameters"] = i.ToBoolMapOutput() }
		if i := state.RestApi; i != nil { inputs["restApi"] = i.ToStringOutput() }
		if i := state.StatusCode; i != nil { inputs["statusCode"] = i.ToStringOutput() }
	}
	var resource MethodResponse
	err := ctx.ReadResource("aws:apigateway/methodResponse:MethodResponse", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MethodResponse resources.
type MethodResponseState struct {
	// The HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
	HttpMethod pulumi.StringInput `pulumi:"httpMethod"`
	// The API resource ID
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	// A map of the API models used for the response's content type
	ResponseModels pulumi.StringMapInput `pulumi:"responseModels"`
	// A map of response parameters that can be sent to the caller.
	// For example: `responseParameters = { "method.response.header.X-Some-Header" = true }`
	// would define that the header `X-Some-Header` can be provided on the response.
	ResponseParameters pulumi.BoolMapInput `pulumi:"responseParameters"`
	// The ID of the associated REST API
	RestApi pulumi.StringInput `pulumi:"restApi"`
	// The HTTP status code
	StatusCode pulumi.StringInput `pulumi:"statusCode"`
}

// The set of arguments for constructing a MethodResponse resource.
type MethodResponseArgs struct {
	// The HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
	HttpMethod pulumi.StringInput `pulumi:"httpMethod"`
	// The API resource ID
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	// A map of the API models used for the response's content type
	ResponseModels pulumi.StringMapInput `pulumi:"responseModels"`
	// A map of response parameters that can be sent to the caller.
	// For example: `responseParameters = { "method.response.header.X-Some-Header" = true }`
	// would define that the header `X-Some-Header` can be provided on the response.
	ResponseParameters pulumi.BoolMapInput `pulumi:"responseParameters"`
	// The ID of the associated REST API
	RestApi pulumi.StringInput `pulumi:"restApi"`
	// The HTTP status code
	StatusCode pulumi.StringInput `pulumi:"statusCode"`
}
