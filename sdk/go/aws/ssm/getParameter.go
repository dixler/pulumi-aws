// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an SSM Parameter data source.
func LookupParameter(ctx *pulumi.Context, args *GetParameterArgs) (*GetParameterResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
		inputs["withDecryption"] = args.WithDecryption
	}
	outputs, err := ctx.Invoke("aws:ssm/getParameter:getParameter", inputs)
	if err != nil {
		return nil, err
	}
	return &GetParameterResult{
		Arn: outputs["arn"],
		Type: outputs["type"],
		Value: outputs["value"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getParameter.
type GetParameterArgs struct {
	// The name of the parameter.
	Name interface{}
	// Whether to return decrypted `SecureString` value. Defaults to `true`.
	WithDecryption interface{}
}

// A collection of values returned by getParameter.
type GetParameterResult struct {
	Arn interface{}
	Type interface{}
	Value interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
