// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an AWS Cognito Identity Pool Roles Attachment.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cognito_identity_pool_roles_attachment.html.markdown.
type IdentityPoolRoleAttachment struct {
	pulumi.CustomResourceState

	// An identity pool ID in the format REGION:GUID.
	IdentityPoolId pulumi.StringOutput `pulumi:"identityPoolId"`

	// A List of Role Mapping.
	RoleMappings IdentityPoolRoleAttachmentRoleMappingsArrayOutput `pulumi:"roleMappings"`

	// The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.
	Roles IdentityPoolRoleAttachmentRolesOutput `pulumi:"roles"`
}

// NewIdentityPoolRoleAttachment registers a new resource with the given unique name, arguments, and options.
func NewIdentityPoolRoleAttachment(ctx *pulumi.Context,
	name string, args *IdentityPoolRoleAttachmentArgs, opts ...pulumi.ResourceOption) (*IdentityPoolRoleAttachment, error) {
	if args == nil || args.IdentityPoolId == nil {
		return nil, errors.New("missing required argument 'IdentityPoolId'")
	}
	if args == nil || args.Roles == nil {
		return nil, errors.New("missing required argument 'Roles'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.IdentityPoolId; i != nil { inputs["identityPoolId"] = i.ToStringOutput() }
		if i := args.RoleMappings; i != nil { inputs["roleMappings"] = i.ToIdentityPoolRoleAttachmentRoleMappingsArrayOutput() }
		if i := args.Roles; i != nil { inputs["roles"] = i.ToIdentityPoolRoleAttachmentRolesOutput() }
	}
	var resource IdentityPoolRoleAttachment
	err := ctx.RegisterResource("aws:cognito/identityPoolRoleAttachment:IdentityPoolRoleAttachment", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityPoolRoleAttachment gets an existing IdentityPoolRoleAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityPoolRoleAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityPoolRoleAttachmentState, opts ...pulumi.ResourceOption) (*IdentityPoolRoleAttachment, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.IdentityPoolId; i != nil { inputs["identityPoolId"] = i.ToStringOutput() }
		if i := state.RoleMappings; i != nil { inputs["roleMappings"] = i.ToIdentityPoolRoleAttachmentRoleMappingsArrayOutput() }
		if i := state.Roles; i != nil { inputs["roles"] = i.ToIdentityPoolRoleAttachmentRolesOutput() }
	}
	var resource IdentityPoolRoleAttachment
	err := ctx.ReadResource("aws:cognito/identityPoolRoleAttachment:IdentityPoolRoleAttachment", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityPoolRoleAttachment resources.
type IdentityPoolRoleAttachmentState struct {
	// An identity pool ID in the format REGION:GUID.
	IdentityPoolId pulumi.StringInput `pulumi:"identityPoolId"`
	// A List of Role Mapping.
	RoleMappings IdentityPoolRoleAttachmentRoleMappingsArrayInput `pulumi:"roleMappings"`
	// The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.
	Roles IdentityPoolRoleAttachmentRolesInput `pulumi:"roles"`
}

// The set of arguments for constructing a IdentityPoolRoleAttachment resource.
type IdentityPoolRoleAttachmentArgs struct {
	// An identity pool ID in the format REGION:GUID.
	IdentityPoolId pulumi.StringInput `pulumi:"identityPoolId"`
	// A List of Role Mapping.
	RoleMappings IdentityPoolRoleAttachmentRoleMappingsArrayInput `pulumi:"roleMappings"`
	// The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.
	Roles IdentityPoolRoleAttachmentRolesInput `pulumi:"roles"`
}
type IdentityPoolRoleAttachmentRoleMappings struct {
	// Specifies the action to be taken if either no rules match the claim value for the Rules type, or there is no cognito:preferred_role claim and there are multiple cognito:roles matches for the Token type. `Required` if you specify Token or Rules as the Type.
	AmbiguousRoleResolution *string `pulumi:"ambiguousRoleResolution"`
	// A string identifying the identity provider, for example, "graph.facebook.com" or "cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id".
	IdentityProvider string `pulumi:"identityProvider"`
	// The Rules Configuration to be used for mapping users to roles. You can specify up to 25 rules per identity provider. Rules are evaluated in order. The first one to match specifies the role.
	MappingRules *[]IdentityPoolRoleAttachmentRoleMappingsMappingRules `pulumi:"mappingRules"`
	// The role mapping type.
	Type string `pulumi:"type"`
}
var identityPoolRoleAttachmentRoleMappingsType = reflect.TypeOf((*IdentityPoolRoleAttachmentRoleMappings)(nil)).Elem()

type IdentityPoolRoleAttachmentRoleMappingsInput interface {
	pulumi.Input

	ToIdentityPoolRoleAttachmentRoleMappingsOutput() IdentityPoolRoleAttachmentRoleMappingsOutput
	ToIdentityPoolRoleAttachmentRoleMappingsOutputWithContext(ctx context.Context) IdentityPoolRoleAttachmentRoleMappingsOutput
}

type IdentityPoolRoleAttachmentRoleMappingsArgs struct {
	// Specifies the action to be taken if either no rules match the claim value for the Rules type, or there is no cognito:preferred_role claim and there are multiple cognito:roles matches for the Token type. `Required` if you specify Token or Rules as the Type.
	AmbiguousRoleResolution pulumi.StringInput `pulumi:"ambiguousRoleResolution"`
	// A string identifying the identity provider, for example, "graph.facebook.com" or "cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id".
	IdentityProvider pulumi.StringInput `pulumi:"identityProvider"`
	// The Rules Configuration to be used for mapping users to roles. You can specify up to 25 rules per identity provider. Rules are evaluated in order. The first one to match specifies the role.
	MappingRules IdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayInput `pulumi:"mappingRules"`
	// The role mapping type.
	Type pulumi.StringInput `pulumi:"type"`
}

func (IdentityPoolRoleAttachmentRoleMappingsArgs) ElementType() reflect.Type {
	return identityPoolRoleAttachmentRoleMappingsType
}

func (a IdentityPoolRoleAttachmentRoleMappingsArgs) ToIdentityPoolRoleAttachmentRoleMappingsOutput() IdentityPoolRoleAttachmentRoleMappingsOutput {
	return pulumi.ToOutput(a).(IdentityPoolRoleAttachmentRoleMappingsOutput)
}

func (a IdentityPoolRoleAttachmentRoleMappingsArgs) ToIdentityPoolRoleAttachmentRoleMappingsOutputWithContext(ctx context.Context) IdentityPoolRoleAttachmentRoleMappingsOutput {
	return pulumi.ToOutputWithContext(ctx, a).(IdentityPoolRoleAttachmentRoleMappingsOutput)
}

type IdentityPoolRoleAttachmentRoleMappingsOutput struct { *pulumi.OutputState }

// Specifies the action to be taken if either no rules match the claim value for the Rules type, or there is no cognito:preferred_role claim and there are multiple cognito:roles matches for the Token type. `Required` if you specify Token or Rules as the Type.
func (o IdentityPoolRoleAttachmentRoleMappingsOutput) AmbiguousRoleResolution() pulumi.StringOutput {
	return o.Apply(func(v IdentityPoolRoleAttachmentRoleMappings) string {
		if v.AmbiguousRoleResolution == nil { return *new(string) } else { return *v.AmbiguousRoleResolution }
	}).(pulumi.StringOutput)
}

// A string identifying the identity provider, for example, "graph.facebook.com" or "cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id".
func (o IdentityPoolRoleAttachmentRoleMappingsOutput) IdentityProvider() pulumi.StringOutput {
	return o.Apply(func(v IdentityPoolRoleAttachmentRoleMappings) string {
		return v.IdentityProvider
	}).(pulumi.StringOutput)
}

// The Rules Configuration to be used for mapping users to roles. You can specify up to 25 rules per identity provider. Rules are evaluated in order. The first one to match specifies the role.
func (o IdentityPoolRoleAttachmentRoleMappingsOutput) MappingRules() IdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayOutput {
	return o.Apply(func(v IdentityPoolRoleAttachmentRoleMappings) []IdentityPoolRoleAttachmentRoleMappingsMappingRules {
		if v.MappingRules == nil { return *new([]IdentityPoolRoleAttachmentRoleMappingsMappingRules) } else { return *v.MappingRules }
	}).(IdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayOutput)
}

// The role mapping type.
func (o IdentityPoolRoleAttachmentRoleMappingsOutput) Type() pulumi.StringOutput {
	return o.Apply(func(v IdentityPoolRoleAttachmentRoleMappings) string {
		return v.Type
	}).(pulumi.StringOutput)
}

func (IdentityPoolRoleAttachmentRoleMappingsOutput) ElementType() reflect.Type {
	return identityPoolRoleAttachmentRoleMappingsType
}

func (o IdentityPoolRoleAttachmentRoleMappingsOutput) ToIdentityPoolRoleAttachmentRoleMappingsOutput() IdentityPoolRoleAttachmentRoleMappingsOutput {
	return o
}

func (o IdentityPoolRoleAttachmentRoleMappingsOutput) ToIdentityPoolRoleAttachmentRoleMappingsOutputWithContext(ctx context.Context) IdentityPoolRoleAttachmentRoleMappingsOutput {
	return o
}

func init() { pulumi.RegisterOutputType(IdentityPoolRoleAttachmentRoleMappingsOutput{}) }

var identityPoolRoleAttachmentRoleMappingsArrayType = reflect.TypeOf((*[]IdentityPoolRoleAttachmentRoleMappings)(nil)).Elem()

type IdentityPoolRoleAttachmentRoleMappingsArrayInput interface {
	pulumi.Input

	ToIdentityPoolRoleAttachmentRoleMappingsArrayOutput() IdentityPoolRoleAttachmentRoleMappingsArrayOutput
	ToIdentityPoolRoleAttachmentRoleMappingsArrayOutputWithContext(ctx context.Context) IdentityPoolRoleAttachmentRoleMappingsArrayOutput
}

type IdentityPoolRoleAttachmentRoleMappingsArrayArgs []IdentityPoolRoleAttachmentRoleMappingsInput

func (IdentityPoolRoleAttachmentRoleMappingsArrayArgs) ElementType() reflect.Type {
	return identityPoolRoleAttachmentRoleMappingsArrayType
}

func (a IdentityPoolRoleAttachmentRoleMappingsArrayArgs) ToIdentityPoolRoleAttachmentRoleMappingsArrayOutput() IdentityPoolRoleAttachmentRoleMappingsArrayOutput {
	return pulumi.ToOutput(a).(IdentityPoolRoleAttachmentRoleMappingsArrayOutput)
}

func (a IdentityPoolRoleAttachmentRoleMappingsArrayArgs) ToIdentityPoolRoleAttachmentRoleMappingsArrayOutputWithContext(ctx context.Context) IdentityPoolRoleAttachmentRoleMappingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(IdentityPoolRoleAttachmentRoleMappingsArrayOutput)
}

type IdentityPoolRoleAttachmentRoleMappingsArrayOutput struct { *pulumi.OutputState }

func (o IdentityPoolRoleAttachmentRoleMappingsArrayOutput) Index(i pulumi.IntInput) IdentityPoolRoleAttachmentRoleMappingsOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) IdentityPoolRoleAttachmentRoleMappings {
		return vs[0].([]IdentityPoolRoleAttachmentRoleMappings)[vs[1].(int)]
	}).(IdentityPoolRoleAttachmentRoleMappingsOutput)
}

func (IdentityPoolRoleAttachmentRoleMappingsArrayOutput) ElementType() reflect.Type {
	return identityPoolRoleAttachmentRoleMappingsArrayType
}

func (o IdentityPoolRoleAttachmentRoleMappingsArrayOutput) ToIdentityPoolRoleAttachmentRoleMappingsArrayOutput() IdentityPoolRoleAttachmentRoleMappingsArrayOutput {
	return o
}

func (o IdentityPoolRoleAttachmentRoleMappingsArrayOutput) ToIdentityPoolRoleAttachmentRoleMappingsArrayOutputWithContext(ctx context.Context) IdentityPoolRoleAttachmentRoleMappingsArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(IdentityPoolRoleAttachmentRoleMappingsArrayOutput{}) }

type IdentityPoolRoleAttachmentRoleMappingsMappingRules struct {
	// The claim name that must be present in the token, for example, "isAdmin" or "paid".
	Claim string `pulumi:"claim"`
	// The match condition that specifies how closely the claim value in the IdP token must match Value.
	MatchType string `pulumi:"matchType"`
	// The role ARN.
	RoleArn string `pulumi:"roleArn"`
	// A brief string that the claim must match, for example, "paid" or "yes".
	Value string `pulumi:"value"`
}
var identityPoolRoleAttachmentRoleMappingsMappingRulesType = reflect.TypeOf((*IdentityPoolRoleAttachmentRoleMappingsMappingRules)(nil)).Elem()

type IdentityPoolRoleAttachmentRoleMappingsMappingRulesInput interface {
	pulumi.Input

	ToIdentityPoolRoleAttachmentRoleMappingsMappingRulesOutput() IdentityPoolRoleAttachmentRoleMappingsMappingRulesOutput
	ToIdentityPoolRoleAttachmentRoleMappingsMappingRulesOutputWithContext(ctx context.Context) IdentityPoolRoleAttachmentRoleMappingsMappingRulesOutput
}

type IdentityPoolRoleAttachmentRoleMappingsMappingRulesArgs struct {
	// The claim name that must be present in the token, for example, "isAdmin" or "paid".
	Claim pulumi.StringInput `pulumi:"claim"`
	// The match condition that specifies how closely the claim value in the IdP token must match Value.
	MatchType pulumi.StringInput `pulumi:"matchType"`
	// The role ARN.
	RoleArn pulumi.StringInput `pulumi:"roleArn"`
	// A brief string that the claim must match, for example, "paid" or "yes".
	Value pulumi.StringInput `pulumi:"value"`
}

func (IdentityPoolRoleAttachmentRoleMappingsMappingRulesArgs) ElementType() reflect.Type {
	return identityPoolRoleAttachmentRoleMappingsMappingRulesType
}

func (a IdentityPoolRoleAttachmentRoleMappingsMappingRulesArgs) ToIdentityPoolRoleAttachmentRoleMappingsMappingRulesOutput() IdentityPoolRoleAttachmentRoleMappingsMappingRulesOutput {
	return pulumi.ToOutput(a).(IdentityPoolRoleAttachmentRoleMappingsMappingRulesOutput)
}

func (a IdentityPoolRoleAttachmentRoleMappingsMappingRulesArgs) ToIdentityPoolRoleAttachmentRoleMappingsMappingRulesOutputWithContext(ctx context.Context) IdentityPoolRoleAttachmentRoleMappingsMappingRulesOutput {
	return pulumi.ToOutputWithContext(ctx, a).(IdentityPoolRoleAttachmentRoleMappingsMappingRulesOutput)
}

type IdentityPoolRoleAttachmentRoleMappingsMappingRulesOutput struct { *pulumi.OutputState }

// The claim name that must be present in the token, for example, "isAdmin" or "paid".
func (o IdentityPoolRoleAttachmentRoleMappingsMappingRulesOutput) Claim() pulumi.StringOutput {
	return o.Apply(func(v IdentityPoolRoleAttachmentRoleMappingsMappingRules) string {
		return v.Claim
	}).(pulumi.StringOutput)
}

// The match condition that specifies how closely the claim value in the IdP token must match Value.
func (o IdentityPoolRoleAttachmentRoleMappingsMappingRulesOutput) MatchType() pulumi.StringOutput {
	return o.Apply(func(v IdentityPoolRoleAttachmentRoleMappingsMappingRules) string {
		return v.MatchType
	}).(pulumi.StringOutput)
}

// The role ARN.
func (o IdentityPoolRoleAttachmentRoleMappingsMappingRulesOutput) RoleArn() pulumi.StringOutput {
	return o.Apply(func(v IdentityPoolRoleAttachmentRoleMappingsMappingRules) string {
		return v.RoleArn
	}).(pulumi.StringOutput)
}

// A brief string that the claim must match, for example, "paid" or "yes".
func (o IdentityPoolRoleAttachmentRoleMappingsMappingRulesOutput) Value() pulumi.StringOutput {
	return o.Apply(func(v IdentityPoolRoleAttachmentRoleMappingsMappingRules) string {
		return v.Value
	}).(pulumi.StringOutput)
}

func (IdentityPoolRoleAttachmentRoleMappingsMappingRulesOutput) ElementType() reflect.Type {
	return identityPoolRoleAttachmentRoleMappingsMappingRulesType
}

func (o IdentityPoolRoleAttachmentRoleMappingsMappingRulesOutput) ToIdentityPoolRoleAttachmentRoleMappingsMappingRulesOutput() IdentityPoolRoleAttachmentRoleMappingsMappingRulesOutput {
	return o
}

func (o IdentityPoolRoleAttachmentRoleMappingsMappingRulesOutput) ToIdentityPoolRoleAttachmentRoleMappingsMappingRulesOutputWithContext(ctx context.Context) IdentityPoolRoleAttachmentRoleMappingsMappingRulesOutput {
	return o
}

func init() { pulumi.RegisterOutputType(IdentityPoolRoleAttachmentRoleMappingsMappingRulesOutput{}) }

var identityPoolRoleAttachmentRoleMappingsMappingRulesArrayType = reflect.TypeOf((*[]IdentityPoolRoleAttachmentRoleMappingsMappingRules)(nil)).Elem()

type IdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayInput interface {
	pulumi.Input

	ToIdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayOutput() IdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayOutput
	ToIdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayOutputWithContext(ctx context.Context) IdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayOutput
}

type IdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayArgs []IdentityPoolRoleAttachmentRoleMappingsMappingRulesInput

func (IdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayArgs) ElementType() reflect.Type {
	return identityPoolRoleAttachmentRoleMappingsMappingRulesArrayType
}

func (a IdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayArgs) ToIdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayOutput() IdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayOutput {
	return pulumi.ToOutput(a).(IdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayOutput)
}

func (a IdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayArgs) ToIdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayOutputWithContext(ctx context.Context) IdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(IdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayOutput)
}

type IdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayOutput struct { *pulumi.OutputState }

func (o IdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayOutput) Index(i pulumi.IntInput) IdentityPoolRoleAttachmentRoleMappingsMappingRulesOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) IdentityPoolRoleAttachmentRoleMappingsMappingRules {
		return vs[0].([]IdentityPoolRoleAttachmentRoleMappingsMappingRules)[vs[1].(int)]
	}).(IdentityPoolRoleAttachmentRoleMappingsMappingRulesOutput)
}

func (IdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayOutput) ElementType() reflect.Type {
	return identityPoolRoleAttachmentRoleMappingsMappingRulesArrayType
}

func (o IdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayOutput) ToIdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayOutput() IdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayOutput {
	return o
}

func (o IdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayOutput) ToIdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayOutputWithContext(ctx context.Context) IdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(IdentityPoolRoleAttachmentRoleMappingsMappingRulesArrayOutput{}) }

type IdentityPoolRoleAttachmentRoles struct {
	Authenticated *string `pulumi:"authenticated"`
	Unauthenticated *string `pulumi:"unauthenticated"`
}
var identityPoolRoleAttachmentRolesType = reflect.TypeOf((*IdentityPoolRoleAttachmentRoles)(nil)).Elem()

type IdentityPoolRoleAttachmentRolesInput interface {
	pulumi.Input

	ToIdentityPoolRoleAttachmentRolesOutput() IdentityPoolRoleAttachmentRolesOutput
	ToIdentityPoolRoleAttachmentRolesOutputWithContext(ctx context.Context) IdentityPoolRoleAttachmentRolesOutput
}

type IdentityPoolRoleAttachmentRolesArgs struct {
	Authenticated pulumi.StringInput `pulumi:"authenticated"`
	Unauthenticated pulumi.StringInput `pulumi:"unauthenticated"`
}

func (IdentityPoolRoleAttachmentRolesArgs) ElementType() reflect.Type {
	return identityPoolRoleAttachmentRolesType
}

func (a IdentityPoolRoleAttachmentRolesArgs) ToIdentityPoolRoleAttachmentRolesOutput() IdentityPoolRoleAttachmentRolesOutput {
	return pulumi.ToOutput(a).(IdentityPoolRoleAttachmentRolesOutput)
}

func (a IdentityPoolRoleAttachmentRolesArgs) ToIdentityPoolRoleAttachmentRolesOutputWithContext(ctx context.Context) IdentityPoolRoleAttachmentRolesOutput {
	return pulumi.ToOutputWithContext(ctx, a).(IdentityPoolRoleAttachmentRolesOutput)
}

type IdentityPoolRoleAttachmentRolesOutput struct { *pulumi.OutputState }

func (o IdentityPoolRoleAttachmentRolesOutput) Authenticated() pulumi.StringOutput {
	return o.Apply(func(v IdentityPoolRoleAttachmentRoles) string {
		if v.Authenticated == nil { return *new(string) } else { return *v.Authenticated }
	}).(pulumi.StringOutput)
}

func (o IdentityPoolRoleAttachmentRolesOutput) Unauthenticated() pulumi.StringOutput {
	return o.Apply(func(v IdentityPoolRoleAttachmentRoles) string {
		if v.Unauthenticated == nil { return *new(string) } else { return *v.Unauthenticated }
	}).(pulumi.StringOutput)
}

func (IdentityPoolRoleAttachmentRolesOutput) ElementType() reflect.Type {
	return identityPoolRoleAttachmentRolesType
}

func (o IdentityPoolRoleAttachmentRolesOutput) ToIdentityPoolRoleAttachmentRolesOutput() IdentityPoolRoleAttachmentRolesOutput {
	return o
}

func (o IdentityPoolRoleAttachmentRolesOutput) ToIdentityPoolRoleAttachmentRolesOutputWithContext(ctx context.Context) IdentityPoolRoleAttachmentRolesOutput {
	return o
}

func init() { pulumi.RegisterOutputType(IdentityPoolRoleAttachmentRolesOutput{}) }

