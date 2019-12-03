// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"context"
	"reflect"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to create an organization.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/organizations_organization.html.markdown.
type Organization struct {
	pulumi.CustomResourceState

	// List of organization accounts including the master account. For a list excluding the master account, see the `nonMasterAccounts` attribute. All elements have these attributes:
	Accounts OrganizationAccountsArrayOutput `pulumi:"accounts"`

	// ARN of the root
	Arn pulumi.StringOutput `pulumi:"arn"`

	// List of AWS service principal names for which you want to enable integration with your organization. This is typically in the form of a URL, such as service-abbreviation.amazonaws.com. Organization must have `featureSet` set to `ALL`. For additional information, see the [AWS Organizations User Guide](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html).
	AwsServiceAccessPrincipals pulumi.StringArrayOutput `pulumi:"awsServiceAccessPrincipals"`

	// List of Organizations policy types to enable in the Organization Root. Organization must have `featureSet` set to `ALL`. For additional information about valid policy types (e.g. `SERVICE_CONTROL_POLICY`), see the [AWS Organizations API Reference](https://docs.aws.amazon.com/organizations/latest/APIReference/API_EnablePolicyType.html).
	EnabledPolicyTypes pulumi.StringArrayOutput `pulumi:"enabledPolicyTypes"`

	// Specify "ALL" (default) or "CONSOLIDATED_BILLING".
	FeatureSet pulumi.StringOutput `pulumi:"featureSet"`

	// ARN of the master account
	MasterAccountArn pulumi.StringOutput `pulumi:"masterAccountArn"`

	// Email address of the master account
	MasterAccountEmail pulumi.StringOutput `pulumi:"masterAccountEmail"`

	// Identifier of the master account
	MasterAccountId pulumi.StringOutput `pulumi:"masterAccountId"`

	// List of organization accounts excluding the master account. For a list including the master account, see the `accounts` attribute. All elements have these attributes:
	NonMasterAccounts OrganizationNonMasterAccountsArrayOutput `pulumi:"nonMasterAccounts"`

	// List of organization roots. All elements have these attributes:
	Roots OrganizationRootsArrayOutput `pulumi:"roots"`
}

// NewOrganization registers a new resource with the given unique name, arguments, and options.
func NewOrganization(ctx *pulumi.Context,
	name string, args *OrganizationArgs, opts ...pulumi.ResourceOption) (*Organization, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.AwsServiceAccessPrincipals; i != nil { inputs["awsServiceAccessPrincipals"] = i.ToStringArrayOutput() }
		if i := args.EnabledPolicyTypes; i != nil { inputs["enabledPolicyTypes"] = i.ToStringArrayOutput() }
		if i := args.FeatureSet; i != nil { inputs["featureSet"] = i.ToStringOutput() }
	}
	var resource Organization
	err := ctx.RegisterResource("aws:organizations/organization:Organization", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganization gets an existing Organization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationState, opts ...pulumi.ResourceOption) (*Organization, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Accounts; i != nil { inputs["accounts"] = i.ToOrganizationAccountsArrayOutput() }
		if i := state.Arn; i != nil { inputs["arn"] = i.ToStringOutput() }
		if i := state.AwsServiceAccessPrincipals; i != nil { inputs["awsServiceAccessPrincipals"] = i.ToStringArrayOutput() }
		if i := state.EnabledPolicyTypes; i != nil { inputs["enabledPolicyTypes"] = i.ToStringArrayOutput() }
		if i := state.FeatureSet; i != nil { inputs["featureSet"] = i.ToStringOutput() }
		if i := state.MasterAccountArn; i != nil { inputs["masterAccountArn"] = i.ToStringOutput() }
		if i := state.MasterAccountEmail; i != nil { inputs["masterAccountEmail"] = i.ToStringOutput() }
		if i := state.MasterAccountId; i != nil { inputs["masterAccountId"] = i.ToStringOutput() }
		if i := state.NonMasterAccounts; i != nil { inputs["nonMasterAccounts"] = i.ToOrganizationNonMasterAccountsArrayOutput() }
		if i := state.Roots; i != nil { inputs["roots"] = i.ToOrganizationRootsArrayOutput() }
	}
	var resource Organization
	err := ctx.ReadResource("aws:organizations/organization:Organization", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Organization resources.
type OrganizationState struct {
	// List of organization accounts including the master account. For a list excluding the master account, see the `nonMasterAccounts` attribute. All elements have these attributes:
	Accounts OrganizationAccountsArrayInput `pulumi:"accounts"`
	// ARN of the root
	Arn pulumi.StringInput `pulumi:"arn"`
	// List of AWS service principal names for which you want to enable integration with your organization. This is typically in the form of a URL, such as service-abbreviation.amazonaws.com. Organization must have `featureSet` set to `ALL`. For additional information, see the [AWS Organizations User Guide](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html).
	AwsServiceAccessPrincipals pulumi.StringArrayInput `pulumi:"awsServiceAccessPrincipals"`
	// List of Organizations policy types to enable in the Organization Root. Organization must have `featureSet` set to `ALL`. For additional information about valid policy types (e.g. `SERVICE_CONTROL_POLICY`), see the [AWS Organizations API Reference](https://docs.aws.amazon.com/organizations/latest/APIReference/API_EnablePolicyType.html).
	EnabledPolicyTypes pulumi.StringArrayInput `pulumi:"enabledPolicyTypes"`
	// Specify "ALL" (default) or "CONSOLIDATED_BILLING".
	FeatureSet pulumi.StringInput `pulumi:"featureSet"`
	// ARN of the master account
	MasterAccountArn pulumi.StringInput `pulumi:"masterAccountArn"`
	// Email address of the master account
	MasterAccountEmail pulumi.StringInput `pulumi:"masterAccountEmail"`
	// Identifier of the master account
	MasterAccountId pulumi.StringInput `pulumi:"masterAccountId"`
	// List of organization accounts excluding the master account. For a list including the master account, see the `accounts` attribute. All elements have these attributes:
	NonMasterAccounts OrganizationNonMasterAccountsArrayInput `pulumi:"nonMasterAccounts"`
	// List of organization roots. All elements have these attributes:
	Roots OrganizationRootsArrayInput `pulumi:"roots"`
}

// The set of arguments for constructing a Organization resource.
type OrganizationArgs struct {
	// List of AWS service principal names for which you want to enable integration with your organization. This is typically in the form of a URL, such as service-abbreviation.amazonaws.com. Organization must have `featureSet` set to `ALL`. For additional information, see the [AWS Organizations User Guide](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html).
	AwsServiceAccessPrincipals pulumi.StringArrayInput `pulumi:"awsServiceAccessPrincipals"`
	// List of Organizations policy types to enable in the Organization Root. Organization must have `featureSet` set to `ALL`. For additional information about valid policy types (e.g. `SERVICE_CONTROL_POLICY`), see the [AWS Organizations API Reference](https://docs.aws.amazon.com/organizations/latest/APIReference/API_EnablePolicyType.html).
	EnabledPolicyTypes pulumi.StringArrayInput `pulumi:"enabledPolicyTypes"`
	// Specify "ALL" (default) or "CONSOLIDATED_BILLING".
	FeatureSet pulumi.StringInput `pulumi:"featureSet"`
}
type OrganizationAccounts struct {
	// ARN of the root
	Arn string `pulumi:"arn"`
	// Email of the account
	Email string `pulumi:"email"`
	// Identifier of the root
	Id string `pulumi:"id"`
	// The name of the policy type
	Name string `pulumi:"name"`
}
var organizationAccountsType = reflect.TypeOf((*OrganizationAccounts)(nil)).Elem()

type OrganizationAccountsInput interface {
	pulumi.Input

	ToOrganizationAccountsOutput() OrganizationAccountsOutput
	ToOrganizationAccountsOutputWithContext(ctx context.Context) OrganizationAccountsOutput
}

type OrganizationAccountsArgs struct {
	// ARN of the root
	Arn pulumi.StringInput `pulumi:"arn"`
	// Email of the account
	Email pulumi.StringInput `pulumi:"email"`
	// Identifier of the root
	Id pulumi.StringInput `pulumi:"id"`
	// The name of the policy type
	Name pulumi.StringInput `pulumi:"name"`
}

func (OrganizationAccountsArgs) ElementType() reflect.Type {
	return organizationAccountsType
}

func (a OrganizationAccountsArgs) ToOrganizationAccountsOutput() OrganizationAccountsOutput {
	return pulumi.ToOutput(a).(OrganizationAccountsOutput)
}

func (a OrganizationAccountsArgs) ToOrganizationAccountsOutputWithContext(ctx context.Context) OrganizationAccountsOutput {
	return pulumi.ToOutputWithContext(ctx, a).(OrganizationAccountsOutput)
}

type OrganizationAccountsOutput struct { *pulumi.OutputState }

// ARN of the root
func (o OrganizationAccountsOutput) Arn() pulumi.StringOutput {
	return o.Apply(func(v OrganizationAccounts) string {
		return v.Arn
	}).(pulumi.StringOutput)
}

// Email of the account
func (o OrganizationAccountsOutput) Email() pulumi.StringOutput {
	return o.Apply(func(v OrganizationAccounts) string {
		return v.Email
	}).(pulumi.StringOutput)
}

// Identifier of the root
func (o OrganizationAccountsOutput) Id() pulumi.StringOutput {
	return o.Apply(func(v OrganizationAccounts) string {
		return v.Id
	}).(pulumi.StringOutput)
}

// The name of the policy type
func (o OrganizationAccountsOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v OrganizationAccounts) string {
		return v.Name
	}).(pulumi.StringOutput)
}

func (OrganizationAccountsOutput) ElementType() reflect.Type {
	return organizationAccountsType
}

func (o OrganizationAccountsOutput) ToOrganizationAccountsOutput() OrganizationAccountsOutput {
	return o
}

func (o OrganizationAccountsOutput) ToOrganizationAccountsOutputWithContext(ctx context.Context) OrganizationAccountsOutput {
	return o
}

func init() { pulumi.RegisterOutputType(OrganizationAccountsOutput{}) }

var organizationAccountsArrayType = reflect.TypeOf((*[]OrganizationAccounts)(nil)).Elem()

type OrganizationAccountsArrayInput interface {
	pulumi.Input

	ToOrganizationAccountsArrayOutput() OrganizationAccountsArrayOutput
	ToOrganizationAccountsArrayOutputWithContext(ctx context.Context) OrganizationAccountsArrayOutput
}

type OrganizationAccountsArrayArgs []OrganizationAccountsInput

func (OrganizationAccountsArrayArgs) ElementType() reflect.Type {
	return organizationAccountsArrayType
}

func (a OrganizationAccountsArrayArgs) ToOrganizationAccountsArrayOutput() OrganizationAccountsArrayOutput {
	return pulumi.ToOutput(a).(OrganizationAccountsArrayOutput)
}

func (a OrganizationAccountsArrayArgs) ToOrganizationAccountsArrayOutputWithContext(ctx context.Context) OrganizationAccountsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(OrganizationAccountsArrayOutput)
}

type OrganizationAccountsArrayOutput struct { *pulumi.OutputState }

func (o OrganizationAccountsArrayOutput) Index(i pulumi.IntInput) OrganizationAccountsOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) OrganizationAccounts {
		return vs[0].([]OrganizationAccounts)[vs[1].(int)]
	}).(OrganizationAccountsOutput)
}

func (OrganizationAccountsArrayOutput) ElementType() reflect.Type {
	return organizationAccountsArrayType
}

func (o OrganizationAccountsArrayOutput) ToOrganizationAccountsArrayOutput() OrganizationAccountsArrayOutput {
	return o
}

func (o OrganizationAccountsArrayOutput) ToOrganizationAccountsArrayOutputWithContext(ctx context.Context) OrganizationAccountsArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(OrganizationAccountsArrayOutput{}) }

type OrganizationNonMasterAccounts struct {
	// ARN of the root
	Arn string `pulumi:"arn"`
	// Email of the account
	Email string `pulumi:"email"`
	// Identifier of the root
	Id string `pulumi:"id"`
	// The name of the policy type
	Name string `pulumi:"name"`
}
var organizationNonMasterAccountsType = reflect.TypeOf((*OrganizationNonMasterAccounts)(nil)).Elem()

type OrganizationNonMasterAccountsInput interface {
	pulumi.Input

	ToOrganizationNonMasterAccountsOutput() OrganizationNonMasterAccountsOutput
	ToOrganizationNonMasterAccountsOutputWithContext(ctx context.Context) OrganizationNonMasterAccountsOutput
}

type OrganizationNonMasterAccountsArgs struct {
	// ARN of the root
	Arn pulumi.StringInput `pulumi:"arn"`
	// Email of the account
	Email pulumi.StringInput `pulumi:"email"`
	// Identifier of the root
	Id pulumi.StringInput `pulumi:"id"`
	// The name of the policy type
	Name pulumi.StringInput `pulumi:"name"`
}

func (OrganizationNonMasterAccountsArgs) ElementType() reflect.Type {
	return organizationNonMasterAccountsType
}

func (a OrganizationNonMasterAccountsArgs) ToOrganizationNonMasterAccountsOutput() OrganizationNonMasterAccountsOutput {
	return pulumi.ToOutput(a).(OrganizationNonMasterAccountsOutput)
}

func (a OrganizationNonMasterAccountsArgs) ToOrganizationNonMasterAccountsOutputWithContext(ctx context.Context) OrganizationNonMasterAccountsOutput {
	return pulumi.ToOutputWithContext(ctx, a).(OrganizationNonMasterAccountsOutput)
}

type OrganizationNonMasterAccountsOutput struct { *pulumi.OutputState }

// ARN of the root
func (o OrganizationNonMasterAccountsOutput) Arn() pulumi.StringOutput {
	return o.Apply(func(v OrganizationNonMasterAccounts) string {
		return v.Arn
	}).(pulumi.StringOutput)
}

// Email of the account
func (o OrganizationNonMasterAccountsOutput) Email() pulumi.StringOutput {
	return o.Apply(func(v OrganizationNonMasterAccounts) string {
		return v.Email
	}).(pulumi.StringOutput)
}

// Identifier of the root
func (o OrganizationNonMasterAccountsOutput) Id() pulumi.StringOutput {
	return o.Apply(func(v OrganizationNonMasterAccounts) string {
		return v.Id
	}).(pulumi.StringOutput)
}

// The name of the policy type
func (o OrganizationNonMasterAccountsOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v OrganizationNonMasterAccounts) string {
		return v.Name
	}).(pulumi.StringOutput)
}

func (OrganizationNonMasterAccountsOutput) ElementType() reflect.Type {
	return organizationNonMasterAccountsType
}

func (o OrganizationNonMasterAccountsOutput) ToOrganizationNonMasterAccountsOutput() OrganizationNonMasterAccountsOutput {
	return o
}

func (o OrganizationNonMasterAccountsOutput) ToOrganizationNonMasterAccountsOutputWithContext(ctx context.Context) OrganizationNonMasterAccountsOutput {
	return o
}

func init() { pulumi.RegisterOutputType(OrganizationNonMasterAccountsOutput{}) }

var organizationNonMasterAccountsArrayType = reflect.TypeOf((*[]OrganizationNonMasterAccounts)(nil)).Elem()

type OrganizationNonMasterAccountsArrayInput interface {
	pulumi.Input

	ToOrganizationNonMasterAccountsArrayOutput() OrganizationNonMasterAccountsArrayOutput
	ToOrganizationNonMasterAccountsArrayOutputWithContext(ctx context.Context) OrganizationNonMasterAccountsArrayOutput
}

type OrganizationNonMasterAccountsArrayArgs []OrganizationNonMasterAccountsInput

func (OrganizationNonMasterAccountsArrayArgs) ElementType() reflect.Type {
	return organizationNonMasterAccountsArrayType
}

func (a OrganizationNonMasterAccountsArrayArgs) ToOrganizationNonMasterAccountsArrayOutput() OrganizationNonMasterAccountsArrayOutput {
	return pulumi.ToOutput(a).(OrganizationNonMasterAccountsArrayOutput)
}

func (a OrganizationNonMasterAccountsArrayArgs) ToOrganizationNonMasterAccountsArrayOutputWithContext(ctx context.Context) OrganizationNonMasterAccountsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(OrganizationNonMasterAccountsArrayOutput)
}

type OrganizationNonMasterAccountsArrayOutput struct { *pulumi.OutputState }

func (o OrganizationNonMasterAccountsArrayOutput) Index(i pulumi.IntInput) OrganizationNonMasterAccountsOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) OrganizationNonMasterAccounts {
		return vs[0].([]OrganizationNonMasterAccounts)[vs[1].(int)]
	}).(OrganizationNonMasterAccountsOutput)
}

func (OrganizationNonMasterAccountsArrayOutput) ElementType() reflect.Type {
	return organizationNonMasterAccountsArrayType
}

func (o OrganizationNonMasterAccountsArrayOutput) ToOrganizationNonMasterAccountsArrayOutput() OrganizationNonMasterAccountsArrayOutput {
	return o
}

func (o OrganizationNonMasterAccountsArrayOutput) ToOrganizationNonMasterAccountsArrayOutputWithContext(ctx context.Context) OrganizationNonMasterAccountsArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(OrganizationNonMasterAccountsArrayOutput{}) }

type OrganizationRoots struct {
	// ARN of the root
	Arn string `pulumi:"arn"`
	// Identifier of the root
	Id string `pulumi:"id"`
	// The name of the policy type
	Name string `pulumi:"name"`
	// List of policy types enabled for this root. All elements have these attributes:
	PolicyTypes []OrganizationRootsPolicyTypes `pulumi:"policyTypes"`
}
var organizationRootsType = reflect.TypeOf((*OrganizationRoots)(nil)).Elem()

type OrganizationRootsInput interface {
	pulumi.Input

	ToOrganizationRootsOutput() OrganizationRootsOutput
	ToOrganizationRootsOutputWithContext(ctx context.Context) OrganizationRootsOutput
}

type OrganizationRootsArgs struct {
	// ARN of the root
	Arn pulumi.StringInput `pulumi:"arn"`
	// Identifier of the root
	Id pulumi.StringInput `pulumi:"id"`
	// The name of the policy type
	Name pulumi.StringInput `pulumi:"name"`
	// List of policy types enabled for this root. All elements have these attributes:
	PolicyTypes OrganizationRootsPolicyTypesArrayInput `pulumi:"policyTypes"`
}

func (OrganizationRootsArgs) ElementType() reflect.Type {
	return organizationRootsType
}

func (a OrganizationRootsArgs) ToOrganizationRootsOutput() OrganizationRootsOutput {
	return pulumi.ToOutput(a).(OrganizationRootsOutput)
}

func (a OrganizationRootsArgs) ToOrganizationRootsOutputWithContext(ctx context.Context) OrganizationRootsOutput {
	return pulumi.ToOutputWithContext(ctx, a).(OrganizationRootsOutput)
}

type OrganizationRootsOutput struct { *pulumi.OutputState }

// ARN of the root
func (o OrganizationRootsOutput) Arn() pulumi.StringOutput {
	return o.Apply(func(v OrganizationRoots) string {
		return v.Arn
	}).(pulumi.StringOutput)
}

// Identifier of the root
func (o OrganizationRootsOutput) Id() pulumi.StringOutput {
	return o.Apply(func(v OrganizationRoots) string {
		return v.Id
	}).(pulumi.StringOutput)
}

// The name of the policy type
func (o OrganizationRootsOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v OrganizationRoots) string {
		return v.Name
	}).(pulumi.StringOutput)
}

// List of policy types enabled for this root. All elements have these attributes:
func (o OrganizationRootsOutput) PolicyTypes() OrganizationRootsPolicyTypesArrayOutput {
	return o.Apply(func(v OrganizationRoots) []OrganizationRootsPolicyTypes {
		return v.PolicyTypes
	}).(OrganizationRootsPolicyTypesArrayOutput)
}

func (OrganizationRootsOutput) ElementType() reflect.Type {
	return organizationRootsType
}

func (o OrganizationRootsOutput) ToOrganizationRootsOutput() OrganizationRootsOutput {
	return o
}

func (o OrganizationRootsOutput) ToOrganizationRootsOutputWithContext(ctx context.Context) OrganizationRootsOutput {
	return o
}

func init() { pulumi.RegisterOutputType(OrganizationRootsOutput{}) }

var organizationRootsArrayType = reflect.TypeOf((*[]OrganizationRoots)(nil)).Elem()

type OrganizationRootsArrayInput interface {
	pulumi.Input

	ToOrganizationRootsArrayOutput() OrganizationRootsArrayOutput
	ToOrganizationRootsArrayOutputWithContext(ctx context.Context) OrganizationRootsArrayOutput
}

type OrganizationRootsArrayArgs []OrganizationRootsInput

func (OrganizationRootsArrayArgs) ElementType() reflect.Type {
	return organizationRootsArrayType
}

func (a OrganizationRootsArrayArgs) ToOrganizationRootsArrayOutput() OrganizationRootsArrayOutput {
	return pulumi.ToOutput(a).(OrganizationRootsArrayOutput)
}

func (a OrganizationRootsArrayArgs) ToOrganizationRootsArrayOutputWithContext(ctx context.Context) OrganizationRootsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(OrganizationRootsArrayOutput)
}

type OrganizationRootsArrayOutput struct { *pulumi.OutputState }

func (o OrganizationRootsArrayOutput) Index(i pulumi.IntInput) OrganizationRootsOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) OrganizationRoots {
		return vs[0].([]OrganizationRoots)[vs[1].(int)]
	}).(OrganizationRootsOutput)
}

func (OrganizationRootsArrayOutput) ElementType() reflect.Type {
	return organizationRootsArrayType
}

func (o OrganizationRootsArrayOutput) ToOrganizationRootsArrayOutput() OrganizationRootsArrayOutput {
	return o
}

func (o OrganizationRootsArrayOutput) ToOrganizationRootsArrayOutputWithContext(ctx context.Context) OrganizationRootsArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(OrganizationRootsArrayOutput{}) }

type OrganizationRootsPolicyTypes struct {
	// The status of the policy type as it relates to the associated root
	Status string `pulumi:"status"`
	Type string `pulumi:"type"`
}
var organizationRootsPolicyTypesType = reflect.TypeOf((*OrganizationRootsPolicyTypes)(nil)).Elem()

type OrganizationRootsPolicyTypesInput interface {
	pulumi.Input

	ToOrganizationRootsPolicyTypesOutput() OrganizationRootsPolicyTypesOutput
	ToOrganizationRootsPolicyTypesOutputWithContext(ctx context.Context) OrganizationRootsPolicyTypesOutput
}

type OrganizationRootsPolicyTypesArgs struct {
	// The status of the policy type as it relates to the associated root
	Status pulumi.StringInput `pulumi:"status"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (OrganizationRootsPolicyTypesArgs) ElementType() reflect.Type {
	return organizationRootsPolicyTypesType
}

func (a OrganizationRootsPolicyTypesArgs) ToOrganizationRootsPolicyTypesOutput() OrganizationRootsPolicyTypesOutput {
	return pulumi.ToOutput(a).(OrganizationRootsPolicyTypesOutput)
}

func (a OrganizationRootsPolicyTypesArgs) ToOrganizationRootsPolicyTypesOutputWithContext(ctx context.Context) OrganizationRootsPolicyTypesOutput {
	return pulumi.ToOutputWithContext(ctx, a).(OrganizationRootsPolicyTypesOutput)
}

type OrganizationRootsPolicyTypesOutput struct { *pulumi.OutputState }

// The status of the policy type as it relates to the associated root
func (o OrganizationRootsPolicyTypesOutput) Status() pulumi.StringOutput {
	return o.Apply(func(v OrganizationRootsPolicyTypes) string {
		return v.Status
	}).(pulumi.StringOutput)
}

func (o OrganizationRootsPolicyTypesOutput) Type() pulumi.StringOutput {
	return o.Apply(func(v OrganizationRootsPolicyTypes) string {
		return v.Type
	}).(pulumi.StringOutput)
}

func (OrganizationRootsPolicyTypesOutput) ElementType() reflect.Type {
	return organizationRootsPolicyTypesType
}

func (o OrganizationRootsPolicyTypesOutput) ToOrganizationRootsPolicyTypesOutput() OrganizationRootsPolicyTypesOutput {
	return o
}

func (o OrganizationRootsPolicyTypesOutput) ToOrganizationRootsPolicyTypesOutputWithContext(ctx context.Context) OrganizationRootsPolicyTypesOutput {
	return o
}

func init() { pulumi.RegisterOutputType(OrganizationRootsPolicyTypesOutput{}) }

var organizationRootsPolicyTypesArrayType = reflect.TypeOf((*[]OrganizationRootsPolicyTypes)(nil)).Elem()

type OrganizationRootsPolicyTypesArrayInput interface {
	pulumi.Input

	ToOrganizationRootsPolicyTypesArrayOutput() OrganizationRootsPolicyTypesArrayOutput
	ToOrganizationRootsPolicyTypesArrayOutputWithContext(ctx context.Context) OrganizationRootsPolicyTypesArrayOutput
}

type OrganizationRootsPolicyTypesArrayArgs []OrganizationRootsPolicyTypesInput

func (OrganizationRootsPolicyTypesArrayArgs) ElementType() reflect.Type {
	return organizationRootsPolicyTypesArrayType
}

func (a OrganizationRootsPolicyTypesArrayArgs) ToOrganizationRootsPolicyTypesArrayOutput() OrganizationRootsPolicyTypesArrayOutput {
	return pulumi.ToOutput(a).(OrganizationRootsPolicyTypesArrayOutput)
}

func (a OrganizationRootsPolicyTypesArrayArgs) ToOrganizationRootsPolicyTypesArrayOutputWithContext(ctx context.Context) OrganizationRootsPolicyTypesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(OrganizationRootsPolicyTypesArrayOutput)
}

type OrganizationRootsPolicyTypesArrayOutput struct { *pulumi.OutputState }

func (o OrganizationRootsPolicyTypesArrayOutput) Index(i pulumi.IntInput) OrganizationRootsPolicyTypesOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) OrganizationRootsPolicyTypes {
		return vs[0].([]OrganizationRootsPolicyTypes)[vs[1].(int)]
	}).(OrganizationRootsPolicyTypesOutput)
}

func (OrganizationRootsPolicyTypesArrayOutput) ElementType() reflect.Type {
	return organizationRootsPolicyTypesArrayType
}

func (o OrganizationRootsPolicyTypesArrayOutput) ToOrganizationRootsPolicyTypesArrayOutput() OrganizationRootsPolicyTypesArrayOutput {
	return o
}

func (o OrganizationRootsPolicyTypesArrayOutput) ToOrganizationRootsPolicyTypesArrayOutputWithContext(ctx context.Context) OrganizationRootsPolicyTypesArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(OrganizationRootsPolicyTypesArrayOutput{}) }

