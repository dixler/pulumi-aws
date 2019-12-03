// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dynamodb

import (
	"context"
	"reflect"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a DynamoDB table resource
// 
// > **Note:** It is recommended to use `lifecycle` [`ignoreChanges`](https://www.terraform.io/docs/configuration/resources.html#ignore_changes) for `readCapacity` and/or `writeCapacity` if there's [autoscaling policy](https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy.html) attached to the table.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dynamodb_table.html.markdown.
type Table struct {
	pulumi.CustomResourceState

	// The arn of the table
	Arn pulumi.StringOutput `pulumi:"arn"`

	// List of nested attribute definitions. Only required for `hashKey` and `rangeKey` attributes. Each attribute has two properties:
	Attributes TableAttributesArrayOutput `pulumi:"attributes"`

	// Controls how you are charged for read and write throughput and how you manage capacity. The valid values are `PROVISIONED` and `PAY_PER_REQUEST`. Defaults to `PROVISIONED`.
	BillingMode pulumi.StringOutput `pulumi:"billingMode"`

	// Describe a GSI for the table;
	// subject to the normal limits on the number of GSIs, projected
	// attributes, etc.
	GlobalSecondaryIndexes TableGlobalSecondaryIndexesArrayOutput `pulumi:"globalSecondaryIndexes"`

	// The name of the hash key in the index; must be
	// defined as an attribute in the resource.
	HashKey pulumi.StringOutput `pulumi:"hashKey"`

	// Describe an LSI on the table;
	// these can only be allocated *at creation* so you cannot change this
	// definition after you have created the resource.
	LocalSecondaryIndexes TableLocalSecondaryIndexesArrayOutput `pulumi:"localSecondaryIndexes"`

	// The name of the index
	Name pulumi.StringOutput `pulumi:"name"`

	// Point-in-time recovery options.
	PointInTimeRecovery TablePointInTimeRecoveryOutput `pulumi:"pointInTimeRecovery"`

	// The name of the range key; must be defined
	RangeKey pulumi.StringOutput `pulumi:"rangeKey"`

	// The number of read units for this index. Must be set if billingMode is set to PROVISIONED.
	ReadCapacity pulumi.IntOutput `pulumi:"readCapacity"`

	// Encryption at rest options. AWS DynamoDB tables are automatically encrypted at rest with an AWS owned Customer Master Key if this argument isn't specified.
	ServerSideEncryption TableServerSideEncryptionOutput `pulumi:"serverSideEncryption"`

	// The ARN of the Table Stream. Only available when `streamEnabled = true`
	StreamArn pulumi.StringOutput `pulumi:"streamArn"`

	// Indicates whether Streams are to be enabled (true) or disabled (false).
	StreamEnabled pulumi.BoolOutput `pulumi:"streamEnabled"`

	// A timestamp, in ISO 8601 format, for this stream. Note that this timestamp is not
	// a unique identifier for the stream on its own. However, the combination of AWS customer ID,
	// table name and this field is guaranteed to be unique.
	// It can be used for creating CloudWatch Alarms. Only available when `streamEnabled = true`
	StreamLabel pulumi.StringOutput `pulumi:"streamLabel"`

	// When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are `KEYS_ONLY`, `NEW_IMAGE`, `OLD_IMAGE`, `NEW_AND_OLD_IMAGES`.
	StreamViewType pulumi.StringOutput `pulumi:"streamViewType"`

	// A map of tags to populate on the created table.
	Tags pulumi.MapOutput `pulumi:"tags"`

	// Defines ttl, has two properties, and can only be specified once:
	Ttl TableTtlOutput `pulumi:"ttl"`

	// The number of write units for this index. Must be set if billingMode is set to PROVISIONED.
	WriteCapacity pulumi.IntOutput `pulumi:"writeCapacity"`
}

// NewTable registers a new resource with the given unique name, arguments, and options.
func NewTable(ctx *pulumi.Context,
	name string, args *TableArgs, opts ...pulumi.ResourceOption) (*Table, error) {
	if args == nil || args.Attributes == nil {
		return nil, errors.New("missing required argument 'Attributes'")
	}
	if args == nil || args.HashKey == nil {
		return nil, errors.New("missing required argument 'HashKey'")
	}
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.Attributes; i != nil { inputs["attributes"] = i.ToTableAttributesArrayOutput() }
		if i := args.BillingMode; i != nil { inputs["billingMode"] = i.ToStringOutput() }
		if i := args.GlobalSecondaryIndexes; i != nil { inputs["globalSecondaryIndexes"] = i.ToTableGlobalSecondaryIndexesArrayOutput() }
		if i := args.HashKey; i != nil { inputs["hashKey"] = i.ToStringOutput() }
		if i := args.LocalSecondaryIndexes; i != nil { inputs["localSecondaryIndexes"] = i.ToTableLocalSecondaryIndexesArrayOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.PointInTimeRecovery; i != nil { inputs["pointInTimeRecovery"] = i.ToTablePointInTimeRecoveryOutput() }
		if i := args.RangeKey; i != nil { inputs["rangeKey"] = i.ToStringOutput() }
		if i := args.ReadCapacity; i != nil { inputs["readCapacity"] = i.ToIntOutput() }
		if i := args.ServerSideEncryption; i != nil { inputs["serverSideEncryption"] = i.ToTableServerSideEncryptionOutput() }
		if i := args.StreamEnabled; i != nil { inputs["streamEnabled"] = i.ToBoolOutput() }
		if i := args.StreamViewType; i != nil { inputs["streamViewType"] = i.ToStringOutput() }
		if i := args.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := args.Ttl; i != nil { inputs["ttl"] = i.ToTableTtlOutput() }
		if i := args.WriteCapacity; i != nil { inputs["writeCapacity"] = i.ToIntOutput() }
	}
	var resource Table
	err := ctx.RegisterResource("aws:dynamodb/table:Table", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTable gets an existing Table resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableState, opts ...pulumi.ResourceOption) (*Table, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.Arn; i != nil { inputs["arn"] = i.ToStringOutput() }
		if i := state.Attributes; i != nil { inputs["attributes"] = i.ToTableAttributesArrayOutput() }
		if i := state.BillingMode; i != nil { inputs["billingMode"] = i.ToStringOutput() }
		if i := state.GlobalSecondaryIndexes; i != nil { inputs["globalSecondaryIndexes"] = i.ToTableGlobalSecondaryIndexesArrayOutput() }
		if i := state.HashKey; i != nil { inputs["hashKey"] = i.ToStringOutput() }
		if i := state.LocalSecondaryIndexes; i != nil { inputs["localSecondaryIndexes"] = i.ToTableLocalSecondaryIndexesArrayOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.PointInTimeRecovery; i != nil { inputs["pointInTimeRecovery"] = i.ToTablePointInTimeRecoveryOutput() }
		if i := state.RangeKey; i != nil { inputs["rangeKey"] = i.ToStringOutput() }
		if i := state.ReadCapacity; i != nil { inputs["readCapacity"] = i.ToIntOutput() }
		if i := state.ServerSideEncryption; i != nil { inputs["serverSideEncryption"] = i.ToTableServerSideEncryptionOutput() }
		if i := state.StreamArn; i != nil { inputs["streamArn"] = i.ToStringOutput() }
		if i := state.StreamEnabled; i != nil { inputs["streamEnabled"] = i.ToBoolOutput() }
		if i := state.StreamLabel; i != nil { inputs["streamLabel"] = i.ToStringOutput() }
		if i := state.StreamViewType; i != nil { inputs["streamViewType"] = i.ToStringOutput() }
		if i := state.Tags; i != nil { inputs["tags"] = i.ToMapOutput() }
		if i := state.Ttl; i != nil { inputs["ttl"] = i.ToTableTtlOutput() }
		if i := state.WriteCapacity; i != nil { inputs["writeCapacity"] = i.ToIntOutput() }
	}
	var resource Table
	err := ctx.ReadResource("aws:dynamodb/table:Table", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Table resources.
type TableState struct {
	// The arn of the table
	Arn pulumi.StringInput `pulumi:"arn"`
	// List of nested attribute definitions. Only required for `hashKey` and `rangeKey` attributes. Each attribute has two properties:
	Attributes TableAttributesArrayInput `pulumi:"attributes"`
	// Controls how you are charged for read and write throughput and how you manage capacity. The valid values are `PROVISIONED` and `PAY_PER_REQUEST`. Defaults to `PROVISIONED`.
	BillingMode pulumi.StringInput `pulumi:"billingMode"`
	// Describe a GSI for the table;
	// subject to the normal limits on the number of GSIs, projected
	// attributes, etc.
	GlobalSecondaryIndexes TableGlobalSecondaryIndexesArrayInput `pulumi:"globalSecondaryIndexes"`
	// The name of the hash key in the index; must be
	// defined as an attribute in the resource.
	HashKey pulumi.StringInput `pulumi:"hashKey"`
	// Describe an LSI on the table;
	// these can only be allocated *at creation* so you cannot change this
	// definition after you have created the resource.
	LocalSecondaryIndexes TableLocalSecondaryIndexesArrayInput `pulumi:"localSecondaryIndexes"`
	// The name of the index
	Name pulumi.StringInput `pulumi:"name"`
	// Point-in-time recovery options.
	PointInTimeRecovery TablePointInTimeRecoveryInput `pulumi:"pointInTimeRecovery"`
	// The name of the range key; must be defined
	RangeKey pulumi.StringInput `pulumi:"rangeKey"`
	// The number of read units for this index. Must be set if billingMode is set to PROVISIONED.
	ReadCapacity pulumi.IntInput `pulumi:"readCapacity"`
	// Encryption at rest options. AWS DynamoDB tables are automatically encrypted at rest with an AWS owned Customer Master Key if this argument isn't specified.
	ServerSideEncryption TableServerSideEncryptionInput `pulumi:"serverSideEncryption"`
	// The ARN of the Table Stream. Only available when `streamEnabled = true`
	StreamArn pulumi.StringInput `pulumi:"streamArn"`
	// Indicates whether Streams are to be enabled (true) or disabled (false).
	StreamEnabled pulumi.BoolInput `pulumi:"streamEnabled"`
	// A timestamp, in ISO 8601 format, for this stream. Note that this timestamp is not
	// a unique identifier for the stream on its own. However, the combination of AWS customer ID,
	// table name and this field is guaranteed to be unique.
	// It can be used for creating CloudWatch Alarms. Only available when `streamEnabled = true`
	StreamLabel pulumi.StringInput `pulumi:"streamLabel"`
	// When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are `KEYS_ONLY`, `NEW_IMAGE`, `OLD_IMAGE`, `NEW_AND_OLD_IMAGES`.
	StreamViewType pulumi.StringInput `pulumi:"streamViewType"`
	// A map of tags to populate on the created table.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Defines ttl, has two properties, and can only be specified once:
	Ttl TableTtlInput `pulumi:"ttl"`
	// The number of write units for this index. Must be set if billingMode is set to PROVISIONED.
	WriteCapacity pulumi.IntInput `pulumi:"writeCapacity"`
}

// The set of arguments for constructing a Table resource.
type TableArgs struct {
	// List of nested attribute definitions. Only required for `hashKey` and `rangeKey` attributes. Each attribute has two properties:
	Attributes TableAttributesArrayInput `pulumi:"attributes"`
	// Controls how you are charged for read and write throughput and how you manage capacity. The valid values are `PROVISIONED` and `PAY_PER_REQUEST`. Defaults to `PROVISIONED`.
	BillingMode pulumi.StringInput `pulumi:"billingMode"`
	// Describe a GSI for the table;
	// subject to the normal limits on the number of GSIs, projected
	// attributes, etc.
	GlobalSecondaryIndexes TableGlobalSecondaryIndexesArrayInput `pulumi:"globalSecondaryIndexes"`
	// The name of the hash key in the index; must be
	// defined as an attribute in the resource.
	HashKey pulumi.StringInput `pulumi:"hashKey"`
	// Describe an LSI on the table;
	// these can only be allocated *at creation* so you cannot change this
	// definition after you have created the resource.
	LocalSecondaryIndexes TableLocalSecondaryIndexesArrayInput `pulumi:"localSecondaryIndexes"`
	// The name of the index
	Name pulumi.StringInput `pulumi:"name"`
	// Point-in-time recovery options.
	PointInTimeRecovery TablePointInTimeRecoveryInput `pulumi:"pointInTimeRecovery"`
	// The name of the range key; must be defined
	RangeKey pulumi.StringInput `pulumi:"rangeKey"`
	// The number of read units for this index. Must be set if billingMode is set to PROVISIONED.
	ReadCapacity pulumi.IntInput `pulumi:"readCapacity"`
	// Encryption at rest options. AWS DynamoDB tables are automatically encrypted at rest with an AWS owned Customer Master Key if this argument isn't specified.
	ServerSideEncryption TableServerSideEncryptionInput `pulumi:"serverSideEncryption"`
	// Indicates whether Streams are to be enabled (true) or disabled (false).
	StreamEnabled pulumi.BoolInput `pulumi:"streamEnabled"`
	// When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are `KEYS_ONLY`, `NEW_IMAGE`, `OLD_IMAGE`, `NEW_AND_OLD_IMAGES`.
	StreamViewType pulumi.StringInput `pulumi:"streamViewType"`
	// A map of tags to populate on the created table.
	Tags pulumi.MapInput `pulumi:"tags"`
	// Defines ttl, has two properties, and can only be specified once:
	Ttl TableTtlInput `pulumi:"ttl"`
	// The number of write units for this index. Must be set if billingMode is set to PROVISIONED.
	WriteCapacity pulumi.IntInput `pulumi:"writeCapacity"`
}
type TableAttributes struct {
	// The name of the index
	Name string `pulumi:"name"`
	// Attribute type, which must be a scalar type: `S`, `N`, or `B` for (S)tring, (N)umber or (B)inary data
	Type string `pulumi:"type"`
}
var tableAttributesType = reflect.TypeOf((*TableAttributes)(nil)).Elem()

type TableAttributesInput interface {
	pulumi.Input

	ToTableAttributesOutput() TableAttributesOutput
	ToTableAttributesOutputWithContext(ctx context.Context) TableAttributesOutput
}

type TableAttributesArgs struct {
	// The name of the index
	Name pulumi.StringInput `pulumi:"name"`
	// Attribute type, which must be a scalar type: `S`, `N`, or `B` for (S)tring, (N)umber or (B)inary data
	Type pulumi.StringInput `pulumi:"type"`
}

func (TableAttributesArgs) ElementType() reflect.Type {
	return tableAttributesType
}

func (a TableAttributesArgs) ToTableAttributesOutput() TableAttributesOutput {
	return pulumi.ToOutput(a).(TableAttributesOutput)
}

func (a TableAttributesArgs) ToTableAttributesOutputWithContext(ctx context.Context) TableAttributesOutput {
	return pulumi.ToOutputWithContext(ctx, a).(TableAttributesOutput)
}

type TableAttributesOutput struct { *pulumi.OutputState }

// The name of the index
func (o TableAttributesOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v TableAttributes) string {
		return v.Name
	}).(pulumi.StringOutput)
}

// Attribute type, which must be a scalar type: `S`, `N`, or `B` for (S)tring, (N)umber or (B)inary data
func (o TableAttributesOutput) Type() pulumi.StringOutput {
	return o.Apply(func(v TableAttributes) string {
		return v.Type
	}).(pulumi.StringOutput)
}

func (TableAttributesOutput) ElementType() reflect.Type {
	return tableAttributesType
}

func (o TableAttributesOutput) ToTableAttributesOutput() TableAttributesOutput {
	return o
}

func (o TableAttributesOutput) ToTableAttributesOutputWithContext(ctx context.Context) TableAttributesOutput {
	return o
}

func init() { pulumi.RegisterOutputType(TableAttributesOutput{}) }

var tableAttributesArrayType = reflect.TypeOf((*[]TableAttributes)(nil)).Elem()

type TableAttributesArrayInput interface {
	pulumi.Input

	ToTableAttributesArrayOutput() TableAttributesArrayOutput
	ToTableAttributesArrayOutputWithContext(ctx context.Context) TableAttributesArrayOutput
}

type TableAttributesArrayArgs []TableAttributesInput

func (TableAttributesArrayArgs) ElementType() reflect.Type {
	return tableAttributesArrayType
}

func (a TableAttributesArrayArgs) ToTableAttributesArrayOutput() TableAttributesArrayOutput {
	return pulumi.ToOutput(a).(TableAttributesArrayOutput)
}

func (a TableAttributesArrayArgs) ToTableAttributesArrayOutputWithContext(ctx context.Context) TableAttributesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(TableAttributesArrayOutput)
}

type TableAttributesArrayOutput struct { *pulumi.OutputState }

func (o TableAttributesArrayOutput) Index(i pulumi.IntInput) TableAttributesOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) TableAttributes {
		return vs[0].([]TableAttributes)[vs[1].(int)]
	}).(TableAttributesOutput)
}

func (TableAttributesArrayOutput) ElementType() reflect.Type {
	return tableAttributesArrayType
}

func (o TableAttributesArrayOutput) ToTableAttributesArrayOutput() TableAttributesArrayOutput {
	return o
}

func (o TableAttributesArrayOutput) ToTableAttributesArrayOutputWithContext(ctx context.Context) TableAttributesArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(TableAttributesArrayOutput{}) }

type TableGlobalSecondaryIndexes struct {
	// The name of the hash key in the index; must be
	// defined as an attribute in the resource.
	HashKey string `pulumi:"hashKey"`
	// The name of the index
	Name string `pulumi:"name"`
	// Only required with `INCLUDE` as a
	// projection type; a list of attributes to project into the index. These
	// do not need to be defined as attributes on the table.
	NonKeyAttributes *[]string `pulumi:"nonKeyAttributes"`
	// One of `ALL`, `INCLUDE` or `KEYS_ONLY`
	// where `ALL` projects every attribute into the index, `KEYS_ONLY`
	// projects just the hash and range key into the index, and `INCLUDE`
	// projects only the keys specified in the _non_key_attributes_
	// parameter.
	ProjectionType string `pulumi:"projectionType"`
	// The name of the range key; must be defined
	RangeKey *string `pulumi:"rangeKey"`
	// The number of read units for this index. Must be set if billingMode is set to PROVISIONED.
	ReadCapacity *int `pulumi:"readCapacity"`
	// The number of write units for this index. Must be set if billingMode is set to PROVISIONED.
	WriteCapacity *int `pulumi:"writeCapacity"`
}
var tableGlobalSecondaryIndexesType = reflect.TypeOf((*TableGlobalSecondaryIndexes)(nil)).Elem()

type TableGlobalSecondaryIndexesInput interface {
	pulumi.Input

	ToTableGlobalSecondaryIndexesOutput() TableGlobalSecondaryIndexesOutput
	ToTableGlobalSecondaryIndexesOutputWithContext(ctx context.Context) TableGlobalSecondaryIndexesOutput
}

type TableGlobalSecondaryIndexesArgs struct {
	// The name of the hash key in the index; must be
	// defined as an attribute in the resource.
	HashKey pulumi.StringInput `pulumi:"hashKey"`
	// The name of the index
	Name pulumi.StringInput `pulumi:"name"`
	// Only required with `INCLUDE` as a
	// projection type; a list of attributes to project into the index. These
	// do not need to be defined as attributes on the table.
	NonKeyAttributes pulumi.StringArrayInput `pulumi:"nonKeyAttributes"`
	// One of `ALL`, `INCLUDE` or `KEYS_ONLY`
	// where `ALL` projects every attribute into the index, `KEYS_ONLY`
	// projects just the hash and range key into the index, and `INCLUDE`
	// projects only the keys specified in the _non_key_attributes_
	// parameter.
	ProjectionType pulumi.StringInput `pulumi:"projectionType"`
	// The name of the range key; must be defined
	RangeKey pulumi.StringInput `pulumi:"rangeKey"`
	// The number of read units for this index. Must be set if billingMode is set to PROVISIONED.
	ReadCapacity pulumi.IntInput `pulumi:"readCapacity"`
	// The number of write units for this index. Must be set if billingMode is set to PROVISIONED.
	WriteCapacity pulumi.IntInput `pulumi:"writeCapacity"`
}

func (TableGlobalSecondaryIndexesArgs) ElementType() reflect.Type {
	return tableGlobalSecondaryIndexesType
}

func (a TableGlobalSecondaryIndexesArgs) ToTableGlobalSecondaryIndexesOutput() TableGlobalSecondaryIndexesOutput {
	return pulumi.ToOutput(a).(TableGlobalSecondaryIndexesOutput)
}

func (a TableGlobalSecondaryIndexesArgs) ToTableGlobalSecondaryIndexesOutputWithContext(ctx context.Context) TableGlobalSecondaryIndexesOutput {
	return pulumi.ToOutputWithContext(ctx, a).(TableGlobalSecondaryIndexesOutput)
}

type TableGlobalSecondaryIndexesOutput struct { *pulumi.OutputState }

// The name of the hash key in the index; must be
// defined as an attribute in the resource.
func (o TableGlobalSecondaryIndexesOutput) HashKey() pulumi.StringOutput {
	return o.Apply(func(v TableGlobalSecondaryIndexes) string {
		return v.HashKey
	}).(pulumi.StringOutput)
}

// The name of the index
func (o TableGlobalSecondaryIndexesOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v TableGlobalSecondaryIndexes) string {
		return v.Name
	}).(pulumi.StringOutput)
}

// Only required with `INCLUDE` as a
// projection type; a list of attributes to project into the index. These
// do not need to be defined as attributes on the table.
func (o TableGlobalSecondaryIndexesOutput) NonKeyAttributes() pulumi.StringArrayOutput {
	return o.Apply(func(v TableGlobalSecondaryIndexes) []string {
		if v.NonKeyAttributes == nil { return *new([]string) } else { return *v.NonKeyAttributes }
	}).(pulumi.StringArrayOutput)
}

// One of `ALL`, `INCLUDE` or `KEYS_ONLY`
// where `ALL` projects every attribute into the index, `KEYS_ONLY`
// projects just the hash and range key into the index, and `INCLUDE`
// projects only the keys specified in the _non_key_attributes_
// parameter.
func (o TableGlobalSecondaryIndexesOutput) ProjectionType() pulumi.StringOutput {
	return o.Apply(func(v TableGlobalSecondaryIndexes) string {
		return v.ProjectionType
	}).(pulumi.StringOutput)
}

// The name of the range key; must be defined
func (o TableGlobalSecondaryIndexesOutput) RangeKey() pulumi.StringOutput {
	return o.Apply(func(v TableGlobalSecondaryIndexes) string {
		if v.RangeKey == nil { return *new(string) } else { return *v.RangeKey }
	}).(pulumi.StringOutput)
}

// The number of read units for this index. Must be set if billingMode is set to PROVISIONED.
func (o TableGlobalSecondaryIndexesOutput) ReadCapacity() pulumi.IntOutput {
	return o.Apply(func(v TableGlobalSecondaryIndexes) int {
		if v.ReadCapacity == nil { return *new(int) } else { return *v.ReadCapacity }
	}).(pulumi.IntOutput)
}

// The number of write units for this index. Must be set if billingMode is set to PROVISIONED.
func (o TableGlobalSecondaryIndexesOutput) WriteCapacity() pulumi.IntOutput {
	return o.Apply(func(v TableGlobalSecondaryIndexes) int {
		if v.WriteCapacity == nil { return *new(int) } else { return *v.WriteCapacity }
	}).(pulumi.IntOutput)
}

func (TableGlobalSecondaryIndexesOutput) ElementType() reflect.Type {
	return tableGlobalSecondaryIndexesType
}

func (o TableGlobalSecondaryIndexesOutput) ToTableGlobalSecondaryIndexesOutput() TableGlobalSecondaryIndexesOutput {
	return o
}

func (o TableGlobalSecondaryIndexesOutput) ToTableGlobalSecondaryIndexesOutputWithContext(ctx context.Context) TableGlobalSecondaryIndexesOutput {
	return o
}

func init() { pulumi.RegisterOutputType(TableGlobalSecondaryIndexesOutput{}) }

var tableGlobalSecondaryIndexesArrayType = reflect.TypeOf((*[]TableGlobalSecondaryIndexes)(nil)).Elem()

type TableGlobalSecondaryIndexesArrayInput interface {
	pulumi.Input

	ToTableGlobalSecondaryIndexesArrayOutput() TableGlobalSecondaryIndexesArrayOutput
	ToTableGlobalSecondaryIndexesArrayOutputWithContext(ctx context.Context) TableGlobalSecondaryIndexesArrayOutput
}

type TableGlobalSecondaryIndexesArrayArgs []TableGlobalSecondaryIndexesInput

func (TableGlobalSecondaryIndexesArrayArgs) ElementType() reflect.Type {
	return tableGlobalSecondaryIndexesArrayType
}

func (a TableGlobalSecondaryIndexesArrayArgs) ToTableGlobalSecondaryIndexesArrayOutput() TableGlobalSecondaryIndexesArrayOutput {
	return pulumi.ToOutput(a).(TableGlobalSecondaryIndexesArrayOutput)
}

func (a TableGlobalSecondaryIndexesArrayArgs) ToTableGlobalSecondaryIndexesArrayOutputWithContext(ctx context.Context) TableGlobalSecondaryIndexesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(TableGlobalSecondaryIndexesArrayOutput)
}

type TableGlobalSecondaryIndexesArrayOutput struct { *pulumi.OutputState }

func (o TableGlobalSecondaryIndexesArrayOutput) Index(i pulumi.IntInput) TableGlobalSecondaryIndexesOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) TableGlobalSecondaryIndexes {
		return vs[0].([]TableGlobalSecondaryIndexes)[vs[1].(int)]
	}).(TableGlobalSecondaryIndexesOutput)
}

func (TableGlobalSecondaryIndexesArrayOutput) ElementType() reflect.Type {
	return tableGlobalSecondaryIndexesArrayType
}

func (o TableGlobalSecondaryIndexesArrayOutput) ToTableGlobalSecondaryIndexesArrayOutput() TableGlobalSecondaryIndexesArrayOutput {
	return o
}

func (o TableGlobalSecondaryIndexesArrayOutput) ToTableGlobalSecondaryIndexesArrayOutputWithContext(ctx context.Context) TableGlobalSecondaryIndexesArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(TableGlobalSecondaryIndexesArrayOutput{}) }

type TableLocalSecondaryIndexes struct {
	// The name of the index
	Name string `pulumi:"name"`
	// Only required with `INCLUDE` as a
	// projection type; a list of attributes to project into the index. These
	// do not need to be defined as attributes on the table.
	NonKeyAttributes *[]string `pulumi:"nonKeyAttributes"`
	// One of `ALL`, `INCLUDE` or `KEYS_ONLY`
	// where `ALL` projects every attribute into the index, `KEYS_ONLY`
	// projects just the hash and range key into the index, and `INCLUDE`
	// projects only the keys specified in the _non_key_attributes_
	// parameter.
	ProjectionType string `pulumi:"projectionType"`
	// The name of the range key; must be defined
	RangeKey string `pulumi:"rangeKey"`
}
var tableLocalSecondaryIndexesType = reflect.TypeOf((*TableLocalSecondaryIndexes)(nil)).Elem()

type TableLocalSecondaryIndexesInput interface {
	pulumi.Input

	ToTableLocalSecondaryIndexesOutput() TableLocalSecondaryIndexesOutput
	ToTableLocalSecondaryIndexesOutputWithContext(ctx context.Context) TableLocalSecondaryIndexesOutput
}

type TableLocalSecondaryIndexesArgs struct {
	// The name of the index
	Name pulumi.StringInput `pulumi:"name"`
	// Only required with `INCLUDE` as a
	// projection type; a list of attributes to project into the index. These
	// do not need to be defined as attributes on the table.
	NonKeyAttributes pulumi.StringArrayInput `pulumi:"nonKeyAttributes"`
	// One of `ALL`, `INCLUDE` or `KEYS_ONLY`
	// where `ALL` projects every attribute into the index, `KEYS_ONLY`
	// projects just the hash and range key into the index, and `INCLUDE`
	// projects only the keys specified in the _non_key_attributes_
	// parameter.
	ProjectionType pulumi.StringInput `pulumi:"projectionType"`
	// The name of the range key; must be defined
	RangeKey pulumi.StringInput `pulumi:"rangeKey"`
}

func (TableLocalSecondaryIndexesArgs) ElementType() reflect.Type {
	return tableLocalSecondaryIndexesType
}

func (a TableLocalSecondaryIndexesArgs) ToTableLocalSecondaryIndexesOutput() TableLocalSecondaryIndexesOutput {
	return pulumi.ToOutput(a).(TableLocalSecondaryIndexesOutput)
}

func (a TableLocalSecondaryIndexesArgs) ToTableLocalSecondaryIndexesOutputWithContext(ctx context.Context) TableLocalSecondaryIndexesOutput {
	return pulumi.ToOutputWithContext(ctx, a).(TableLocalSecondaryIndexesOutput)
}

type TableLocalSecondaryIndexesOutput struct { *pulumi.OutputState }

// The name of the index
func (o TableLocalSecondaryIndexesOutput) Name() pulumi.StringOutput {
	return o.Apply(func(v TableLocalSecondaryIndexes) string {
		return v.Name
	}).(pulumi.StringOutput)
}

// Only required with `INCLUDE` as a
// projection type; a list of attributes to project into the index. These
// do not need to be defined as attributes on the table.
func (o TableLocalSecondaryIndexesOutput) NonKeyAttributes() pulumi.StringArrayOutput {
	return o.Apply(func(v TableLocalSecondaryIndexes) []string {
		if v.NonKeyAttributes == nil { return *new([]string) } else { return *v.NonKeyAttributes }
	}).(pulumi.StringArrayOutput)
}

// One of `ALL`, `INCLUDE` or `KEYS_ONLY`
// where `ALL` projects every attribute into the index, `KEYS_ONLY`
// projects just the hash and range key into the index, and `INCLUDE`
// projects only the keys specified in the _non_key_attributes_
// parameter.
func (o TableLocalSecondaryIndexesOutput) ProjectionType() pulumi.StringOutput {
	return o.Apply(func(v TableLocalSecondaryIndexes) string {
		return v.ProjectionType
	}).(pulumi.StringOutput)
}

// The name of the range key; must be defined
func (o TableLocalSecondaryIndexesOutput) RangeKey() pulumi.StringOutput {
	return o.Apply(func(v TableLocalSecondaryIndexes) string {
		return v.RangeKey
	}).(pulumi.StringOutput)
}

func (TableLocalSecondaryIndexesOutput) ElementType() reflect.Type {
	return tableLocalSecondaryIndexesType
}

func (o TableLocalSecondaryIndexesOutput) ToTableLocalSecondaryIndexesOutput() TableLocalSecondaryIndexesOutput {
	return o
}

func (o TableLocalSecondaryIndexesOutput) ToTableLocalSecondaryIndexesOutputWithContext(ctx context.Context) TableLocalSecondaryIndexesOutput {
	return o
}

func init() { pulumi.RegisterOutputType(TableLocalSecondaryIndexesOutput{}) }

var tableLocalSecondaryIndexesArrayType = reflect.TypeOf((*[]TableLocalSecondaryIndexes)(nil)).Elem()

type TableLocalSecondaryIndexesArrayInput interface {
	pulumi.Input

	ToTableLocalSecondaryIndexesArrayOutput() TableLocalSecondaryIndexesArrayOutput
	ToTableLocalSecondaryIndexesArrayOutputWithContext(ctx context.Context) TableLocalSecondaryIndexesArrayOutput
}

type TableLocalSecondaryIndexesArrayArgs []TableLocalSecondaryIndexesInput

func (TableLocalSecondaryIndexesArrayArgs) ElementType() reflect.Type {
	return tableLocalSecondaryIndexesArrayType
}

func (a TableLocalSecondaryIndexesArrayArgs) ToTableLocalSecondaryIndexesArrayOutput() TableLocalSecondaryIndexesArrayOutput {
	return pulumi.ToOutput(a).(TableLocalSecondaryIndexesArrayOutput)
}

func (a TableLocalSecondaryIndexesArrayArgs) ToTableLocalSecondaryIndexesArrayOutputWithContext(ctx context.Context) TableLocalSecondaryIndexesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, a).(TableLocalSecondaryIndexesArrayOutput)
}

type TableLocalSecondaryIndexesArrayOutput struct { *pulumi.OutputState }

func (o TableLocalSecondaryIndexesArrayOutput) Index(i pulumi.IntInput) TableLocalSecondaryIndexesOutput {
	return pulumi.All(o, i).Apply(func(vs []interface{}) TableLocalSecondaryIndexes {
		return vs[0].([]TableLocalSecondaryIndexes)[vs[1].(int)]
	}).(TableLocalSecondaryIndexesOutput)
}

func (TableLocalSecondaryIndexesArrayOutput) ElementType() reflect.Type {
	return tableLocalSecondaryIndexesArrayType
}

func (o TableLocalSecondaryIndexesArrayOutput) ToTableLocalSecondaryIndexesArrayOutput() TableLocalSecondaryIndexesArrayOutput {
	return o
}

func (o TableLocalSecondaryIndexesArrayOutput) ToTableLocalSecondaryIndexesArrayOutputWithContext(ctx context.Context) TableLocalSecondaryIndexesArrayOutput {
	return o
}

func init() { pulumi.RegisterOutputType(TableLocalSecondaryIndexesArrayOutput{}) }

type TablePointInTimeRecovery struct {
	// Whether to enable point-in-time recovery - note that it can take up to 10 minutes to enable for new tables. If the `pointInTimeRecovery` block is not provided then this defaults to `false`.
	Enabled bool `pulumi:"enabled"`
}
var tablePointInTimeRecoveryType = reflect.TypeOf((*TablePointInTimeRecovery)(nil)).Elem()

type TablePointInTimeRecoveryInput interface {
	pulumi.Input

	ToTablePointInTimeRecoveryOutput() TablePointInTimeRecoveryOutput
	ToTablePointInTimeRecoveryOutputWithContext(ctx context.Context) TablePointInTimeRecoveryOutput
}

type TablePointInTimeRecoveryArgs struct {
	// Whether to enable point-in-time recovery - note that it can take up to 10 minutes to enable for new tables. If the `pointInTimeRecovery` block is not provided then this defaults to `false`.
	Enabled pulumi.BoolInput `pulumi:"enabled"`
}

func (TablePointInTimeRecoveryArgs) ElementType() reflect.Type {
	return tablePointInTimeRecoveryType
}

func (a TablePointInTimeRecoveryArgs) ToTablePointInTimeRecoveryOutput() TablePointInTimeRecoveryOutput {
	return pulumi.ToOutput(a).(TablePointInTimeRecoveryOutput)
}

func (a TablePointInTimeRecoveryArgs) ToTablePointInTimeRecoveryOutputWithContext(ctx context.Context) TablePointInTimeRecoveryOutput {
	return pulumi.ToOutputWithContext(ctx, a).(TablePointInTimeRecoveryOutput)
}

type TablePointInTimeRecoveryOutput struct { *pulumi.OutputState }

// Whether to enable point-in-time recovery - note that it can take up to 10 minutes to enable for new tables. If the `pointInTimeRecovery` block is not provided then this defaults to `false`.
func (o TablePointInTimeRecoveryOutput) Enabled() pulumi.BoolOutput {
	return o.Apply(func(v TablePointInTimeRecovery) bool {
		return v.Enabled
	}).(pulumi.BoolOutput)
}

func (TablePointInTimeRecoveryOutput) ElementType() reflect.Type {
	return tablePointInTimeRecoveryType
}

func (o TablePointInTimeRecoveryOutput) ToTablePointInTimeRecoveryOutput() TablePointInTimeRecoveryOutput {
	return o
}

func (o TablePointInTimeRecoveryOutput) ToTablePointInTimeRecoveryOutputWithContext(ctx context.Context) TablePointInTimeRecoveryOutput {
	return o
}

func init() { pulumi.RegisterOutputType(TablePointInTimeRecoveryOutput{}) }

type TableServerSideEncryption struct {
	// Whether to enable point-in-time recovery - note that it can take up to 10 minutes to enable for new tables. If the `pointInTimeRecovery` block is not provided then this defaults to `false`.
	Enabled bool `pulumi:"enabled"`
}
var tableServerSideEncryptionType = reflect.TypeOf((*TableServerSideEncryption)(nil)).Elem()

type TableServerSideEncryptionInput interface {
	pulumi.Input

	ToTableServerSideEncryptionOutput() TableServerSideEncryptionOutput
	ToTableServerSideEncryptionOutputWithContext(ctx context.Context) TableServerSideEncryptionOutput
}

type TableServerSideEncryptionArgs struct {
	// Whether to enable point-in-time recovery - note that it can take up to 10 minutes to enable for new tables. If the `pointInTimeRecovery` block is not provided then this defaults to `false`.
	Enabled pulumi.BoolInput `pulumi:"enabled"`
}

func (TableServerSideEncryptionArgs) ElementType() reflect.Type {
	return tableServerSideEncryptionType
}

func (a TableServerSideEncryptionArgs) ToTableServerSideEncryptionOutput() TableServerSideEncryptionOutput {
	return pulumi.ToOutput(a).(TableServerSideEncryptionOutput)
}

func (a TableServerSideEncryptionArgs) ToTableServerSideEncryptionOutputWithContext(ctx context.Context) TableServerSideEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, a).(TableServerSideEncryptionOutput)
}

type TableServerSideEncryptionOutput struct { *pulumi.OutputState }

// Whether to enable point-in-time recovery - note that it can take up to 10 minutes to enable for new tables. If the `pointInTimeRecovery` block is not provided then this defaults to `false`.
func (o TableServerSideEncryptionOutput) Enabled() pulumi.BoolOutput {
	return o.Apply(func(v TableServerSideEncryption) bool {
		return v.Enabled
	}).(pulumi.BoolOutput)
}

func (TableServerSideEncryptionOutput) ElementType() reflect.Type {
	return tableServerSideEncryptionType
}

func (o TableServerSideEncryptionOutput) ToTableServerSideEncryptionOutput() TableServerSideEncryptionOutput {
	return o
}

func (o TableServerSideEncryptionOutput) ToTableServerSideEncryptionOutputWithContext(ctx context.Context) TableServerSideEncryptionOutput {
	return o
}

func init() { pulumi.RegisterOutputType(TableServerSideEncryptionOutput{}) }

type TableTtl struct {
	// The name of the table attribute to store the TTL timestamp in.
	AttributeName string `pulumi:"attributeName"`
	// Whether to enable point-in-time recovery - note that it can take up to 10 minutes to enable for new tables. If the `pointInTimeRecovery` block is not provided then this defaults to `false`.
	Enabled *bool `pulumi:"enabled"`
}
var tableTtlType = reflect.TypeOf((*TableTtl)(nil)).Elem()

type TableTtlInput interface {
	pulumi.Input

	ToTableTtlOutput() TableTtlOutput
	ToTableTtlOutputWithContext(ctx context.Context) TableTtlOutput
}

type TableTtlArgs struct {
	// The name of the table attribute to store the TTL timestamp in.
	AttributeName pulumi.StringInput `pulumi:"attributeName"`
	// Whether to enable point-in-time recovery - note that it can take up to 10 minutes to enable for new tables. If the `pointInTimeRecovery` block is not provided then this defaults to `false`.
	Enabled pulumi.BoolInput `pulumi:"enabled"`
}

func (TableTtlArgs) ElementType() reflect.Type {
	return tableTtlType
}

func (a TableTtlArgs) ToTableTtlOutput() TableTtlOutput {
	return pulumi.ToOutput(a).(TableTtlOutput)
}

func (a TableTtlArgs) ToTableTtlOutputWithContext(ctx context.Context) TableTtlOutput {
	return pulumi.ToOutputWithContext(ctx, a).(TableTtlOutput)
}

type TableTtlOutput struct { *pulumi.OutputState }

// The name of the table attribute to store the TTL timestamp in.
func (o TableTtlOutput) AttributeName() pulumi.StringOutput {
	return o.Apply(func(v TableTtl) string {
		return v.AttributeName
	}).(pulumi.StringOutput)
}

// Whether to enable point-in-time recovery - note that it can take up to 10 minutes to enable for new tables. If the `pointInTimeRecovery` block is not provided then this defaults to `false`.
func (o TableTtlOutput) Enabled() pulumi.BoolOutput {
	return o.Apply(func(v TableTtl) bool {
		if v.Enabled == nil { return *new(bool) } else { return *v.Enabled }
	}).(pulumi.BoolOutput)
}

func (TableTtlOutput) ElementType() reflect.Type {
	return tableTtlType
}

func (o TableTtlOutput) ToTableTtlOutput() TableTtlOutput {
	return o
}

func (o TableTtlOutput) ToTableTtlOutputWithContext(ctx context.Context) TableTtlOutput {
	return o
}

func init() { pulumi.RegisterOutputType(TableTtlOutput{}) }

