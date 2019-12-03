// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a HTTP Method for an API Gateway Resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_method.html.markdown.
type Method struct {
	pulumi.CustomResourceState

	// Specify if the method requires an API key
	ApiKeyRequired pulumi.BoolOutput `pulumi:"apiKeyRequired"`

	// The type of authorization used for the method (`NONE`, `CUSTOM`, `AWS_IAM`, `COGNITO_USER_POOLS`)
	Authorization pulumi.StringOutput `pulumi:"authorization"`

	// The authorization scopes used when the authorization is `COGNITO_USER_POOLS`
	AuthorizationScopes pulumi.StringArrayOutput `pulumi:"authorizationScopes"`

	// The authorizer id to be used when the authorization is `CUSTOM` or `COGNITO_USER_POOLS`
	AuthorizerId pulumi.StringOutput `pulumi:"authorizerId"`

	// The HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
	HttpMethod pulumi.StringOutput `pulumi:"httpMethod"`

	// A map of the API models used for the request's content type
	// where key is the content type (e.g. `application/json`)
	// and value is either `Error`, `Empty` (built-in models) or `apigateway.Model`'s `name`.
	RequestModels pulumi.StringMapOutput `pulumi:"requestModels"`

	// A map of request parameters (from the path, query string and headers) that should be passed to the integration. The boolean value indicates whether the parameter is required (`true`) or optional (`false`).
	// For example: `requestParameters = {"method.request.header.X-Some-Header" = true "method.request.querystring.some-query-param" = true}` would define that the header `X-Some-Header` and the query string `some-query-param` must be provided in the request.
	RequestParameters pulumi.BoolMapOutput `pulumi:"requestParameters"`

	// The ID of a `apigateway.RequestValidator`
	RequestValidatorId pulumi.StringOutput `pulumi:"requestValidatorId"`

	// The API resource ID
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`

	// The ID of the associated REST API
	RestApi pulumi.StringOutput `pulumi:"restApi"`
}

// NewMethod registers a new resource with the given unique name, arguments, and options.
func NewMethod(ctx *pulumi.Context,
	name string, args *MethodArgs, opts ...pulumi.ResourceOption) (*Method, error) {
	if args == nil || args.Authorization == nil {
		return nil, errors.New("missing required argument 'Authorization'")
	}
	if args == nil || args.HttpMethod == nil {
		return nil, errors.New("missing required argument 'HttpMethod'")
	}
	if args == nil || args.ResourceId == nil {
		return nil, errors.New("missing required argument 'ResourceId'")
	}
	if args == nil || args.RestApi == nil {
		return nil, errors.New("missing required argument 'RestApi'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.ApiKeyRequired; i != nil { inputs["apiKeyRequired"] = i.ToBoolOutput() }
		if i := args.Authorization; i != nil { inputs["authorization"] = i.ToStringOutput() }
		if i := args.AuthorizationScopes; i != nil { inputs["authorizationScopes"] = i.ToStringArrayOutput() }
		if i := args.AuthorizerId; i != nil { inputs["authorizerId"] = i.ToStringOutput() }
		if i := args.HttpMethod; i != nil { inputs["httpMethod"] = i.ToStringOutput() }
		if i := args.RequestModels; i != nil { inputs["requestModels"] = i.ToStringMapOutput() }
		if i := args.RequestParameters; i != nil { inputs["requestParameters"] = i.ToBoolMapOutput() }
		if i := args.RequestValidatorId; i != nil { inputs["requestValidatorId"] = i.ToStringOutput() }
		if i := args.ResourceId; i != nil { inputs["resourceId"] = i.ToStringOutput() }
		if i := args.RestApi; i != nil { inputs["restApi"] = i.ToStringOutput() }
	}
	var resource Method
	err := ctx.RegisterResource("aws:apigateway/method:Method", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMethod gets an existing Method resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMethod(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MethodState, opts ...pulumi.ResourceOption) (*Method, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.ApiKeyRequired; i != nil { inputs["apiKeyRequired"] = i.ToBoolOutput() }
		if i := state.Authorization; i != nil { inputs["authorization"] = i.ToStringOutput() }
		if i := state.AuthorizationScopes; i != nil { inputs["authorizationScopes"] = i.ToStringArrayOutput() }
		if i := state.AuthorizerId; i != nil { inputs["authorizerId"] = i.ToStringOutput() }
		if i := state.HttpMethod; i != nil { inputs["httpMethod"] = i.ToStringOutput() }
		if i := state.RequestModels; i != nil { inputs["requestModels"] = i.ToStringMapOutput() }
		if i := state.RequestParameters; i != nil { inputs["requestParameters"] = i.ToBoolMapOutput() }
		if i := state.RequestValidatorId; i != nil { inputs["requestValidatorId"] = i.ToStringOutput() }
		if i := state.ResourceId; i != nil { inputs["resourceId"] = i.ToStringOutput() }
		if i := state.RestApi; i != nil { inputs["restApi"] = i.ToStringOutput() }
	}
	var resource Method
	err := ctx.ReadResource("aws:apigateway/method:Method", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Method resources.
type MethodState struct {
	// Specify if the method requires an API key
	ApiKeyRequired pulumi.BoolInput `pulumi:"apiKeyRequired"`
	// The type of authorization used for the method (`NONE`, `CUSTOM`, `AWS_IAM`, `COGNITO_USER_POOLS`)
	Authorization pulumi.StringInput `pulumi:"authorization"`
	// The authorization scopes used when the authorization is `COGNITO_USER_POOLS`
	AuthorizationScopes pulumi.StringArrayInput `pulumi:"authorizationScopes"`
	// The authorizer id to be used when the authorization is `CUSTOM` or `COGNITO_USER_POOLS`
	AuthorizerId pulumi.StringInput `pulumi:"authorizerId"`
	// The HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
	HttpMethod pulumi.StringInput `pulumi:"httpMethod"`
	// A map of the API models used for the request's content type
	// where key is the content type (e.g. `application/json`)
	// and value is either `Error`, `Empty` (built-in models) or `apigateway.Model`'s `name`.
	RequestModels pulumi.StringMapInput `pulumi:"requestModels"`
	// A map of request parameters (from the path, query string and headers) that should be passed to the integration. The boolean value indicates whether the parameter is required (`true`) or optional (`false`).
	// For example: `requestParameters = {"method.request.header.X-Some-Header" = true "method.request.querystring.some-query-param" = true}` would define that the header `X-Some-Header` and the query string `some-query-param` must be provided in the request.
	RequestParameters pulumi.BoolMapInput `pulumi:"requestParameters"`
	// The ID of a `apigateway.RequestValidator`
	RequestValidatorId pulumi.StringInput `pulumi:"requestValidatorId"`
	// The API resource ID
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	// The ID of the associated REST API
	RestApi pulumi.StringInput `pulumi:"restApi"`
}

// The set of arguments for constructing a Method resource.
type MethodArgs struct {
	// Specify if the method requires an API key
	ApiKeyRequired pulumi.BoolInput `pulumi:"apiKeyRequired"`
	// The type of authorization used for the method (`NONE`, `CUSTOM`, `AWS_IAM`, `COGNITO_USER_POOLS`)
	Authorization pulumi.StringInput `pulumi:"authorization"`
	// The authorization scopes used when the authorization is `COGNITO_USER_POOLS`
	AuthorizationScopes pulumi.StringArrayInput `pulumi:"authorizationScopes"`
	// The authorizer id to be used when the authorization is `CUSTOM` or `COGNITO_USER_POOLS`
	AuthorizerId pulumi.StringInput `pulumi:"authorizerId"`
	// The HTTP Method (`GET`, `POST`, `PUT`, `DELETE`, `HEAD`, `OPTIONS`, `ANY`)
	HttpMethod pulumi.StringInput `pulumi:"httpMethod"`
	// A map of the API models used for the request's content type
	// where key is the content type (e.g. `application/json`)
	// and value is either `Error`, `Empty` (built-in models) or `apigateway.Model`'s `name`.
	RequestModels pulumi.StringMapInput `pulumi:"requestModels"`
	// A map of request parameters (from the path, query string and headers) that should be passed to the integration. The boolean value indicates whether the parameter is required (`true`) or optional (`false`).
	// For example: `requestParameters = {"method.request.header.X-Some-Header" = true "method.request.querystring.some-query-param" = true}` would define that the header `X-Some-Header` and the query string `some-query-param` must be provided in the request.
	RequestParameters pulumi.BoolMapInput `pulumi:"requestParameters"`
	// The ID of a `apigateway.RequestValidator`
	RequestValidatorId pulumi.StringInput `pulumi:"requestValidatorId"`
	// The API resource ID
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	// The ID of the associated REST API
	RestApi pulumi.StringInput `pulumi:"restApi"`
}
