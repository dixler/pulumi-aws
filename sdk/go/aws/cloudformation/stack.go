// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a CloudFormation Stack resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudformation_stack.html.markdown.
type Stack struct {
	pulumi.CustomResourceState

	// A list of capabilities.
	// Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, or `CAPABILITY_AUTO_EXPAND`
	Capabilities pulumi.StringArrayOutput `pulumi:"capabilities"`

	// Set to true to disable rollback of the stack if stack creation failed.
	// Conflicts with `onFailure`.
	DisableRollback pulumi.BoolOutput `pulumi:"disableRollback"`

	// The ARN of an IAM role that AWS CloudFormation assumes to create the stack. If you don't specify a value, AWS CloudFormation uses the role that was previously associated with the stack. If no role is available, AWS CloudFormation uses a temporary session that is generated from your user credentials.
	IamRoleArn pulumi.StringOutput `pulumi:"iamRoleArn"`

	// Stack name.
	Name pulumi.StringOutput `pulumi:"name"`

	// A list of SNS topic ARNs to publish stack related events.
	NotificationArns pulumi.StringArrayOutput `pulumi:"notificationArns"`

	// Action to be taken if stack creation fails. This must be
	// one of: `DO_NOTHING`, `ROLLBACK`, or `DELETE`. Conflicts with `disableRollback`.
	OnFailure pulumi.StringOutput `pulumi:"onFailure"`

	// A map of outputs from the stack.
	Outputs pulumi.MapOutput `pulumi:"outputs"`

	// A map of Parameter structures that specify input parameters for the stack.
	Parameters pulumi.MapOutput `pulumi:"parameters"`

	// Structure containing the stack policy body.
	// Conflicts w/ `policyUrl`.
	PolicyBody pulumi.StringOutput `pulumi:"policyBody"`

	// Location of a file containing the stack policy.
	// Conflicts w/ `policyBody`.
	PolicyUrl pulumi.StringOutput `pulumi:"policyUrl"`

	// A list of tags to associate with this stack.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// Structure containing the template body (max size: 51,200 bytes).
	TemplateBody pulumi.StringOutput `pulumi:"templateBody"`

	// Location of a file containing the template body (max size: 460,800 bytes).
	TemplateUrl pulumi.StringOutput `pulumi:"templateUrl"`

	// The amount of time that can pass before the stack status becomes `CREATE_FAILED`.
	TimeoutInMinutes pulumi.IntOutput `pulumi:"timeoutInMinutes"`
}

// NewStack registers a new resource with the given unique name, arguments, and options.
func NewStack(ctx *pulumi.Context,
	name string, args *StackArgs, opts ...pulumi.ResourceOption) (*Stack, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Capabilities; i != nil { inputs["capabilities"] = i.ToStringArrayOutput() }
		if i := args.DisableRollback; i != nil { inputs["disableRollback"] = i.ToBoolOutput() }
		if i := args.IamRoleArn; i != nil { inputs["iamRoleArn"] = i.ToStringOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.NotificationArns; i != nil { inputs["notificationArns"] = i.ToStringArrayOutput() }
		if i := args.OnFailure; i != nil { inputs["onFailure"] = i.ToStringOutput() }
		if i := args.Parameters; i != nil { inputs["parameters"] = i.ToMapOutput() }
		if i := args.PolicyBody; i != nil { inputs["policyBody"] = i.ToStringOutput() }
		if i := args.PolicyUrl; i != nil { inputs["policyUrl"] = i.ToStringOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := args.TemplateBody; i != nil { inputs["templateBody"] = i.ToStringOutput() }
		if i := args.TemplateUrl; i != nil { inputs["templateUrl"] = i.ToStringOutput() }
		if i := args.TimeoutInMinutes; i != nil { inputs["timeoutInMinutes"] = i.ToIntOutput() }
	}
	var resource Stack
	err := ctx.RegisterResource("aws:cloudformation/stack:Stack", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStack gets an existing Stack resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StackState, opts ...pulumi.ResourceOption) (*Stack, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Capabilities; i != nil { inputs["capabilities"] = i.ToStringArrayOutput() }
		if i := state.DisableRollback; i != nil { inputs["disableRollback"] = i.ToBoolOutput() }
		if i := state.IamRoleArn; i != nil { inputs["iamRoleArn"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.NotificationArns; i != nil { inputs["notificationArns"] = i.ToStringArrayOutput() }
		if i := state.OnFailure; i != nil { inputs["onFailure"] = i.ToStringOutput() }
		if i := state.Outputs; i != nil { inputs["outputs"] = i.ToMapOutput() }
		if i := state.Parameters; i != nil { inputs["parameters"] = i.ToMapOutput() }
		if i := state.PolicyBody; i != nil { inputs["policyBody"] = i.ToStringOutput() }
		if i := state.PolicyUrl; i != nil { inputs["policyUrl"] = i.ToStringOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := state.TemplateBody; i != nil { inputs["templateBody"] = i.ToStringOutput() }
		if i := state.TemplateUrl; i != nil { inputs["templateUrl"] = i.ToStringOutput() }
		if i := state.TimeoutInMinutes; i != nil { inputs["timeoutInMinutes"] = i.ToIntOutput() }
	}
	var resource Stack
	err := ctx.ReadResource("aws:cloudformation/stack:Stack", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Stack resources.
type StackState struct {
	// A list of capabilities.
	// Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, or `CAPABILITY_AUTO_EXPAND`
	Capabilities pulumi.StringArrayInput `pulumi:"capabilities"`
	// Set to true to disable rollback of the stack if stack creation failed.
	// Conflicts with `onFailure`.
	DisableRollback pulumi.BoolInput `pulumi:"disableRollback"`
	// The ARN of an IAM role that AWS CloudFormation assumes to create the stack. If you don't specify a value, AWS CloudFormation uses the role that was previously associated with the stack. If no role is available, AWS CloudFormation uses a temporary session that is generated from your user credentials.
	IamRoleArn pulumi.StringInput `pulumi:"iamRoleArn"`
	// Stack name.
	Name pulumi.StringInput `pulumi:"name"`
	// A list of SNS topic ARNs to publish stack related events.
	NotificationArns pulumi.StringArrayInput `pulumi:"notificationArns"`
	// Action to be taken if stack creation fails. This must be
	// one of: `DO_NOTHING`, `ROLLBACK`, or `DELETE`. Conflicts with `disableRollback`.
	OnFailure pulumi.StringInput `pulumi:"onFailure"`
	// A map of outputs from the stack.
	Outputs pulumi.MapInput `pulumi:"outputs"`
	// A map of Parameter structures that specify input parameters for the stack.
	Parameters pulumi.MapInput `pulumi:"parameters"`
	// Structure containing the stack policy body.
	// Conflicts w/ `policyUrl`.
	PolicyBody pulumi.StringInput `pulumi:"policyBody"`
	// Location of a file containing the stack policy.
	// Conflicts w/ `policyBody`.
	PolicyUrl pulumi.StringInput `pulumi:"policyUrl"`
	// A list of tags to associate with this stack.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Structure containing the template body (max size: 51,200 bytes).
	TemplateBody pulumi.StringInput `pulumi:"templateBody"`
	// Location of a file containing the template body (max size: 460,800 bytes).
	TemplateUrl pulumi.StringInput `pulumi:"templateUrl"`
	// The amount of time that can pass before the stack status becomes `CREATE_FAILED`.
	TimeoutInMinutes pulumi.IntInput `pulumi:"timeoutInMinutes"`
}

// The set of arguments for constructing a Stack resource.
type StackArgs struct {
	// A list of capabilities.
	// Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, or `CAPABILITY_AUTO_EXPAND`
	Capabilities pulumi.StringArrayInput `pulumi:"capabilities"`
	// Set to true to disable rollback of the stack if stack creation failed.
	// Conflicts with `onFailure`.
	DisableRollback pulumi.BoolInput `pulumi:"disableRollback"`
	// The ARN of an IAM role that AWS CloudFormation assumes to create the stack. If you don't specify a value, AWS CloudFormation uses the role that was previously associated with the stack. If no role is available, AWS CloudFormation uses a temporary session that is generated from your user credentials.
	IamRoleArn pulumi.StringInput `pulumi:"iamRoleArn"`
	// Stack name.
	Name pulumi.StringInput `pulumi:"name"`
	// A list of SNS topic ARNs to publish stack related events.
	NotificationArns pulumi.StringArrayInput `pulumi:"notificationArns"`
	// Action to be taken if stack creation fails. This must be
	// one of: `DO_NOTHING`, `ROLLBACK`, or `DELETE`. Conflicts with `disableRollback`.
	OnFailure pulumi.StringInput `pulumi:"onFailure"`
	// A map of Parameter structures that specify input parameters for the stack.
	Parameters pulumi.MapInput `pulumi:"parameters"`
	// Structure containing the stack policy body.
	// Conflicts w/ `policyUrl`.
	PolicyBody pulumi.StringInput `pulumi:"policyBody"`
	// Location of a file containing the stack policy.
	// Conflicts w/ `policyBody`.
	PolicyUrl pulumi.StringInput `pulumi:"policyUrl"`
	// A list of tags to associate with this stack.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Structure containing the template body (max size: 51,200 bytes).
	TemplateBody pulumi.StringInput `pulumi:"templateBody"`
	// Location of a file containing the template body (max size: 460,800 bytes).
	TemplateUrl pulumi.StringInput `pulumi:"templateUrl"`
	// The amount of time that can pass before the stack status becomes `CREATE_FAILED`.
	TimeoutInMinutes pulumi.IntInput `pulumi:"timeoutInMinutes"`
}
