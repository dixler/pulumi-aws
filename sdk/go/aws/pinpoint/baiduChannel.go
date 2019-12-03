// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Pinpoint Baidu Channel resource.
// 
// > **Note:** All arguments including the Api Key and Secret Key will be stored in the raw state as plain-text.
// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/pinpoint_baidu_channel.html.markdown.
type BaiduChannel struct {
	pulumi.CustomResourceState

	// Platform credential API key from Baidu.
	ApiKey pulumi.StringOutput `pulumi:"apiKey"`

	// The application ID.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`

	// Specifies whether to enable the channel. Defaults to `true`.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`

	// Platform credential Secret key from Baidu.
	SecretKey pulumi.StringOutput `pulumi:"secretKey"`
}

// NewBaiduChannel registers a new resource with the given unique name, arguments, and options.
func NewBaiduChannel(ctx *pulumi.Context,
	name string, args *BaiduChannelArgs, opts ...pulumi.ResourceOption) (*BaiduChannel, error) {
	if args == nil || args.ApiKey == nil {
		return nil, errors.New("missing required argument 'ApiKey'")
	}
	if args == nil || args.ApplicationId == nil {
		return nil, errors.New("missing required argument 'ApplicationId'")
	}
	if args == nil || args.SecretKey == nil {
		return nil, errors.New("missing required argument 'SecretKey'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.ApiKey; i != nil { inputs["apiKey"] = i.ToStringOutput() }
		if i := args.ApplicationId; i != nil { inputs["applicationId"] = i.ToStringOutput() }
		if i := args.Enabled; i != nil { inputs["enabled"] = i.ToBoolOutput() }
		if i := args.SecretKey; i != nil { inputs["secretKey"] = i.ToStringOutput() }
	}
	var resource BaiduChannel
	err := ctx.RegisterResource("aws:pinpoint/baiduChannel:BaiduChannel", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBaiduChannel gets an existing BaiduChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBaiduChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BaiduChannelState, opts ...pulumi.ResourceOption) (*BaiduChannel, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.ApiKey; i != nil { inputs["apiKey"] = i.ToStringOutput() }
		if i := state.ApplicationId; i != nil { inputs["applicationId"] = i.ToStringOutput() }
		if i := state.Enabled; i != nil { inputs["enabled"] = i.ToBoolOutput() }
		if i := state.SecretKey; i != nil { inputs["secretKey"] = i.ToStringOutput() }
	}
	var resource BaiduChannel
	err := ctx.ReadResource("aws:pinpoint/baiduChannel:BaiduChannel", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BaiduChannel resources.
type BaiduChannelState struct {
	// Platform credential API key from Baidu.
	ApiKey pulumi.StringInput `pulumi:"apiKey"`
	// The application ID.
	ApplicationId pulumi.StringInput `pulumi:"applicationId"`
	// Specifies whether to enable the channel. Defaults to `true`.
	Enabled pulumi.BoolInput `pulumi:"enabled"`
	// Platform credential Secret key from Baidu.
	SecretKey pulumi.StringInput `pulumi:"secretKey"`
}

// The set of arguments for constructing a BaiduChannel resource.
type BaiduChannelArgs struct {
	// Platform credential API key from Baidu.
	ApiKey pulumi.StringInput `pulumi:"apiKey"`
	// The application ID.
	ApplicationId pulumi.StringInput `pulumi:"applicationId"`
	// Specifies whether to enable the channel. Defaults to `true`.
	Enabled pulumi.BoolInput `pulumi:"enabled"`
	// Platform credential Secret key from Baidu.
	SecretKey pulumi.StringInput `pulumi:"secretKey"`
}
