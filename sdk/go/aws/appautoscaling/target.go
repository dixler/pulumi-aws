// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appautoscaling

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an Application AutoScaling ScalableTarget resource. To manage policies which get attached to the target, see the [`appautoscaling.Policy` resource](https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appautoscaling_target.html.markdown.
type Target struct {
	pulumi.CustomResourceState

	// The max capacity of the scalable target.
	MaxCapacity pulumi.IntOutput `pulumi:"maxCapacity"`

	// The min capacity of the scalable target.
	MinCapacity pulumi.IntOutput `pulumi:"minCapacity"`

	// The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`

	// The ARN of the IAM role that allows Application
	// AutoScaling to modify your scalable target on your behalf.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`

	// The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ScalableDimension pulumi.StringOutput `pulumi:"scalableDimension"`

	// The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ServiceNamespace pulumi.StringOutput `pulumi:"serviceNamespace"`
}

// NewTarget registers a new resource with the given unique name, arguments, and options.
func NewTarget(ctx *pulumi.Context,
	name string, args *TargetArgs, opts ...pulumi.ResourceOption) (*Target, error) {
	if args == nil || args.MaxCapacity == nil {
		return nil, errors.New("missing required argument 'MaxCapacity'")
	}
	if args == nil || args.MinCapacity == nil {
		return nil, errors.New("missing required argument 'MinCapacity'")
	}
	if args == nil || args.ResourceId == nil {
		return nil, errors.New("missing required argument 'ResourceId'")
	}
	if args == nil || args.ScalableDimension == nil {
		return nil, errors.New("missing required argument 'ScalableDimension'")
	}
	if args == nil || args.ServiceNamespace == nil {
		return nil, errors.New("missing required argument 'ServiceNamespace'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.MaxCapacity; i != nil { inputs["maxCapacity"] = i.ToIntOutput() }
		if i := args.MinCapacity; i != nil { inputs["minCapacity"] = i.ToIntOutput() }
		if i := args.ResourceId; i != nil { inputs["resourceId"] = i.ToStringOutput() }
		if i := args.RoleArn; i != nil { inputs["roleArn"] = i.ToStringOutput() }
		if i := args.ScalableDimension; i != nil { inputs["scalableDimension"] = i.ToStringOutput() }
		if i := args.ServiceNamespace; i != nil { inputs["serviceNamespace"] = i.ToStringOutput() }
	}
	var resource Target
	err := ctx.RegisterResource("aws:appautoscaling/target:Target", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTarget gets an existing Target resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetState, opts ...pulumi.ResourceOption) (*Target, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.MaxCapacity; i != nil { inputs["maxCapacity"] = i.ToIntOutput() }
		if i := state.MinCapacity; i != nil { inputs["minCapacity"] = i.ToIntOutput() }
		if i := state.ResourceId; i != nil { inputs["resourceId"] = i.ToStringOutput() }
		if i := state.RoleArn; i != nil { inputs["roleArn"] = i.ToStringOutput() }
		if i := state.ScalableDimension; i != nil { inputs["scalableDimension"] = i.ToStringOutput() }
		if i := state.ServiceNamespace; i != nil { inputs["serviceNamespace"] = i.ToStringOutput() }
	}
	var resource Target
	err := ctx.ReadResource("aws:appautoscaling/target:Target", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Target resources.
type TargetState struct {
	// The max capacity of the scalable target.
	MaxCapacity pulumi.IntInput `pulumi:"maxCapacity"`
	// The min capacity of the scalable target.
	MinCapacity pulumi.IntInput `pulumi:"minCapacity"`
	// The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	// The ARN of the IAM role that allows Application
	// AutoScaling to modify your scalable target on your behalf.
	RoleArn pulumi.StringInput `pulumi:"roleArn"`
	// The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ScalableDimension pulumi.StringInput `pulumi:"scalableDimension"`
	// The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ServiceNamespace pulumi.StringInput `pulumi:"serviceNamespace"`
}

// The set of arguments for constructing a Target resource.
type TargetArgs struct {
	// The max capacity of the scalable target.
	MaxCapacity pulumi.IntInput `pulumi:"maxCapacity"`
	// The min capacity of the scalable target.
	MinCapacity pulumi.IntInput `pulumi:"minCapacity"`
	// The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	// The ARN of the IAM role that allows Application
	// AutoScaling to modify your scalable target on your behalf.
	RoleArn pulumi.StringInput `pulumi:"roleArn"`
	// The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ScalableDimension pulumi.StringInput `pulumi:"scalableDimension"`
	// The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
	ServiceNamespace pulumi.StringInput `pulumi:"serviceNamespace"`
}
