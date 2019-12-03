// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package inspector

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Inspector assessment target
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/inspector_assessment_target.html.markdown.
type AssessmentTarget struct {
	pulumi.CustomResourceState

	// The target assessment ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`

	// The name of the assessment target.
	Name pulumi.StringOutput `pulumi:"name"`

	// Inspector Resource Group Amazon Resource Name (ARN) stating tags for instance matching. If not specified, all EC2 instances in the current AWS account and region are included in the assessment target.
	ResourceGroupArn pulumi.StringOutput `pulumi:"resourceGroupArn"`
}

// NewAssessmentTarget registers a new resource with the given unique name, arguments, and options.
func NewAssessmentTarget(ctx *pulumi.Context,
	name string, args *AssessmentTargetArgs, opts ...pulumi.ResourceOption) (*AssessmentTarget, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.ResourceGroupArn; i != nil { inputs["resourceGroupArn"] = i.ToStringOutput() }
	}
	var resource AssessmentTarget
	err := ctx.RegisterResource("aws:inspector/assessmentTarget:AssessmentTarget", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssessmentTarget gets an existing AssessmentTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssessmentTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssessmentTargetState, opts ...pulumi.ResourceOption) (*AssessmentTarget, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Arn; i != nil { inputs["arn"] = i.ToStringOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.ResourceGroupArn; i != nil { inputs["resourceGroupArn"] = i.ToStringOutput() }
	}
	var resource AssessmentTarget
	err := ctx.ReadResource("aws:inspector/assessmentTarget:AssessmentTarget", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AssessmentTarget resources.
type AssessmentTargetState struct {
	// The target assessment ARN.
	Arn pulumi.StringInput `pulumi:"arn"`
	// The name of the assessment target.
	Name pulumi.StringInput `pulumi:"name"`
	// Inspector Resource Group Amazon Resource Name (ARN) stating tags for instance matching. If not specified, all EC2 instances in the current AWS account and region are included in the assessment target.
	ResourceGroupArn pulumi.StringInput `pulumi:"resourceGroupArn"`
}

// The set of arguments for constructing a AssessmentTarget resource.
type AssessmentTargetArgs struct {
	// The name of the assessment target.
	Name pulumi.StringInput `pulumi:"name"`
	// Inspector Resource Group Amazon Resource Name (ARN) stating tags for instance matching. If not specified, all EC2 instances in the current AWS account and region are included in the assessment target.
	ResourceGroupArn pulumi.StringInput `pulumi:"resourceGroupArn"`
}
