// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Glue Classifier resource.
// 
// > **NOTE:** It is only valid to create one type of classifier (csv, grok, JSON, or XML). Changing classifier types will recreate the classifier.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/glue_classifier.html.markdown.
type Classifier struct {
	pulumi.CustomResourceState

	// A classifier for Csv content. Defined below.
	CsvClassifier ClassifierCsvClassifierOutput `pulumi:"csvClassifier"`

	// A classifier that uses grok patterns. Defined below.
	GrokClassifier ClassifierGrokClassifierOutput `pulumi:"grokClassifier"`

	// A classifier for JSON content. Defined below.
	JsonClassifier ClassifierJsonClassifierOutput `pulumi:"jsonClassifier"`

	// The name of the classifier.
	Name pulumi.StringOutput `pulumi:"name"`

	// A classifier for XML content. Defined below.
	XmlClassifier ClassifierXmlClassifierOutput `pulumi:"xmlClassifier"`
}

// NewClassifier registers a new resource with the given unique name, arguments, and options.
func NewClassifier(ctx *pulumi.Context,
	name string, args *ClassifierArgs, opts ...pulumi.ResourceOption) (*Classifier, error) {
	inputs := map[string]pulumi.Input{}
	if args != nil {
		if i := args.CsvClassifier; i != nil { inputs["csvClassifier"] = i.ToClassifierCsvClassifierOutput() }
		if i := args.GrokClassifier; i != nil { inputs["grokClassifier"] = i.ToClassifierGrokClassifierOutput() }
		if i := args.JsonClassifier; i != nil { inputs["jsonClassifier"] = i.ToClassifierJsonClassifierOutput() }
		if i := args.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := args.XmlClassifier; i != nil { inputs["xmlClassifier"] = i.ToClassifierXmlClassifierOutput() }
	}
	var resource Classifier
	err := ctx.RegisterResource("aws:glue/classifier:Classifier", name, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClassifier gets an existing Classifier resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClassifier(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClassifierState, opts ...pulumi.ResourceOption) (*Classifier, error) {
	inputs := map[string]pulumi.Input{}
	if state != nil {
		if i := state.CsvClassifier; i != nil { inputs["csvClassifier"] = i.ToClassifierCsvClassifierOutput() }
		if i := state.GrokClassifier; i != nil { inputs["grokClassifier"] = i.ToClassifierGrokClassifierOutput() }
		if i := state.JsonClassifier; i != nil { inputs["jsonClassifier"] = i.ToClassifierJsonClassifierOutput() }
		if i := state.Name; i != nil { inputs["name"] = i.ToStringOutput() }
		if i := state.XmlClassifier; i != nil { inputs["xmlClassifier"] = i.ToClassifierXmlClassifierOutput() }
	}
	var resource Classifier
	err := ctx.ReadResource("aws:glue/classifier:Classifier", name, id, inputs, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Classifier resources.
type ClassifierState struct {
	// A classifier for Csv content. Defined below.
	CsvClassifier ClassifierCsvClassifierInput `pulumi:"csvClassifier"`
	// A classifier that uses grok patterns. Defined below.
	GrokClassifier ClassifierGrokClassifierInput `pulumi:"grokClassifier"`
	// A classifier for JSON content. Defined below.
	JsonClassifier ClassifierJsonClassifierInput `pulumi:"jsonClassifier"`
	// The name of the classifier.
	Name pulumi.StringInput `pulumi:"name"`
	// A classifier for XML content. Defined below.
	XmlClassifier ClassifierXmlClassifierInput `pulumi:"xmlClassifier"`
}

// The set of arguments for constructing a Classifier resource.
type ClassifierArgs struct {
	// A classifier for Csv content. Defined below.
	CsvClassifier ClassifierCsvClassifierInput `pulumi:"csvClassifier"`
	// A classifier that uses grok patterns. Defined below.
	GrokClassifier ClassifierGrokClassifierInput `pulumi:"grokClassifier"`
	// A classifier for JSON content. Defined below.
	JsonClassifier ClassifierJsonClassifierInput `pulumi:"jsonClassifier"`
	// The name of the classifier.
	Name pulumi.StringInput `pulumi:"name"`
	// A classifier for XML content. Defined below.
	XmlClassifier ClassifierXmlClassifierInput `pulumi:"xmlClassifier"`
}
type ClassifierCsvClassifier struct {
	// Enables the processing of files that contain only one column.
	AllowSingleColumn *bool `pulumi:"allowSingleColumn"`
	// Indicates whether the CSV file contains a header. This can be one of "ABSENT", "PRESENT", or "UNKNOWN".
	ContainsHeader *string `pulumi:"containsHeader"`
	// The delimiter used in the Csv to separate columns.
	Delimiter *string `pulumi:"delimiter"`
	// Specifies whether to trim column values. 
	DisableValueTrimming *bool `pulumi:"disableValueTrimming"`
	// A list of strings representing column names.
	Headers *[]string `pulumi:"headers"`
	// A custom symbol to denote what combines content into a single column value. It must be different from the column delimiter.
	QuoteSymbol *string `pulumi:"quoteSymbol"`
}
var classifierCsvClassifierType = reflect.TypeOf((*ClassifierCsvClassifier)(nil)).Elem()

type ClassifierCsvClassifierInput interface {
	pulumi.Input

	ToClassifierCsvClassifierOutput() ClassifierCsvClassifierOutput
	ToClassifierCsvClassifierOutputWithContext(ctx context.Context) ClassifierCsvClassifierOutput
}

type ClassifierCsvClassifierArgs struct {
	// Enables the processing of files that contain only one column.
	AllowSingleColumn pulumi.BoolInput `pulumi:"allowSingleColumn"`
	// Indicates whether the CSV file contains a header. This can be one of "ABSENT", "PRESENT", or "UNKNOWN".
	ContainsHeader pulumi.StringInput `pulumi:"containsHeader"`
	// The delimiter used in the Csv to separate columns.
	Delimiter pulumi.StringInput `pulumi:"delimiter"`
	// Specifies whether to trim column values. 
	DisableValueTrimming pulumi.BoolInput `pulumi:"disableValueTrimming"`
	// A list of strings representing column names.
	Headers pulumi.StringArrayInput `pulumi:"headers"`
	// A custom symbol to denote what combines content into a single column value. It must be different from the column delimiter.
	QuoteSymbol pulumi.StringInput `pulumi:"quoteSymbol"`
}

func (ClassifierCsvClassifierArgs) ElementType() reflect.Type {
	return classifierCsvClassifierType
}

func (a ClassifierCsvClassifierArgs) ToClassifierCsvClassifierOutput() ClassifierCsvClassifierOutput {
	return pulumi.ToOutput(a).(ClassifierCsvClassifierOutput)
}

func (a ClassifierCsvClassifierArgs) ToClassifierCsvClassifierOutputWithContext(ctx context.Context) ClassifierCsvClassifierOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ClassifierCsvClassifierOutput)
}

type ClassifierCsvClassifierOutput struct { *pulumi.OutputState }

// Enables the processing of files that contain only one column.
func (o ClassifierCsvClassifierOutput) AllowSingleColumn() pulumi.BoolOutput {
	return o.Apply(func(v ClassifierCsvClassifier) bool {
		if v.AllowSingleColumn == nil { return *new(bool) } else { return *v.AllowSingleColumn }
	}).(pulumi.BoolOutput)
}

// Indicates whether the CSV file contains a header. This can be one of "ABSENT", "PRESENT", or "UNKNOWN".
func (o ClassifierCsvClassifierOutput) ContainsHeader() pulumi.StringOutput {
	return o.Apply(func(v ClassifierCsvClassifier) string {
		if v.ContainsHeader == nil { return *new(string) } else { return *v.ContainsHeader }
	}).(pulumi.StringOutput)
}

// The delimiter used in the Csv to separate columns.
func (o ClassifierCsvClassifierOutput) Delimiter() pulumi.StringOutput {
	return o.Apply(func(v ClassifierCsvClassifier) string {
		if v.Delimiter == nil { return *new(string) } else { return *v.Delimiter }
	}).(pulumi.StringOutput)
}

// Specifies whether to trim column values. 
func (o ClassifierCsvClassifierOutput) DisableValueTrimming() pulumi.BoolOutput {
	return o.Apply(func(v ClassifierCsvClassifier) bool {
		if v.DisableValueTrimming == nil { return *new(bool) } else { return *v.DisableValueTrimming }
	}).(pulumi.BoolOutput)
}

// A list of strings representing column names.
func (o ClassifierCsvClassifierOutput) Headers() pulumi.StringArrayOutput {
	return o.Apply(func(v ClassifierCsvClassifier) []string {
		if v.Headers == nil { return *new([]string) } else { return *v.Headers }
	}).(pulumi.StringArrayOutput)
}

// A custom symbol to denote what combines content into a single column value. It must be different from the column delimiter.
func (o ClassifierCsvClassifierOutput) QuoteSymbol() pulumi.StringOutput {
	return o.Apply(func(v ClassifierCsvClassifier) string {
		if v.QuoteSymbol == nil { return *new(string) } else { return *v.QuoteSymbol }
	}).(pulumi.StringOutput)
}

func (ClassifierCsvClassifierOutput) ElementType() reflect.Type {
	return classifierCsvClassifierType
}

func (o ClassifierCsvClassifierOutput) ToClassifierCsvClassifierOutput() ClassifierCsvClassifierOutput {
	return o
}

func (o ClassifierCsvClassifierOutput) ToClassifierCsvClassifierOutputWithContext(ctx context.Context) ClassifierCsvClassifierOutput {
	return o
}

func init() { pulumi.RegisterOutputType(ClassifierCsvClassifierOutput{}) }

type ClassifierGrokClassifier struct {
	// An identifier of the data format that the classifier matches.
	Classification string `pulumi:"classification"`
	// Custom grok patterns used by this classifier.
	CustomPatterns *string `pulumi:"customPatterns"`
	// The grok pattern used by this classifier.
	GrokPattern string `pulumi:"grokPattern"`
}
var classifierGrokClassifierType = reflect.TypeOf((*ClassifierGrokClassifier)(nil)).Elem()

type ClassifierGrokClassifierInput interface {
	pulumi.Input

	ToClassifierGrokClassifierOutput() ClassifierGrokClassifierOutput
	ToClassifierGrokClassifierOutputWithContext(ctx context.Context) ClassifierGrokClassifierOutput
}

type ClassifierGrokClassifierArgs struct {
	// An identifier of the data format that the classifier matches.
	Classification pulumi.StringInput `pulumi:"classification"`
	// Custom grok patterns used by this classifier.
	CustomPatterns pulumi.StringInput `pulumi:"customPatterns"`
	// The grok pattern used by this classifier.
	GrokPattern pulumi.StringInput `pulumi:"grokPattern"`
}

func (ClassifierGrokClassifierArgs) ElementType() reflect.Type {
	return classifierGrokClassifierType
}

func (a ClassifierGrokClassifierArgs) ToClassifierGrokClassifierOutput() ClassifierGrokClassifierOutput {
	return pulumi.ToOutput(a).(ClassifierGrokClassifierOutput)
}

func (a ClassifierGrokClassifierArgs) ToClassifierGrokClassifierOutputWithContext(ctx context.Context) ClassifierGrokClassifierOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ClassifierGrokClassifierOutput)
}

type ClassifierGrokClassifierOutput struct { *pulumi.OutputState }

// An identifier of the data format that the classifier matches.
func (o ClassifierGrokClassifierOutput) Classification() pulumi.StringOutput {
	return o.Apply(func(v ClassifierGrokClassifier) string {
		return v.Classification
	}).(pulumi.StringOutput)
}

// Custom grok patterns used by this classifier.
func (o ClassifierGrokClassifierOutput) CustomPatterns() pulumi.StringOutput {
	return o.Apply(func(v ClassifierGrokClassifier) string {
		if v.CustomPatterns == nil { return *new(string) } else { return *v.CustomPatterns }
	}).(pulumi.StringOutput)
}

// The grok pattern used by this classifier.
func (o ClassifierGrokClassifierOutput) GrokPattern() pulumi.StringOutput {
	return o.Apply(func(v ClassifierGrokClassifier) string {
		return v.GrokPattern
	}).(pulumi.StringOutput)
}

func (ClassifierGrokClassifierOutput) ElementType() reflect.Type {
	return classifierGrokClassifierType
}

func (o ClassifierGrokClassifierOutput) ToClassifierGrokClassifierOutput() ClassifierGrokClassifierOutput {
	return o
}

func (o ClassifierGrokClassifierOutput) ToClassifierGrokClassifierOutputWithContext(ctx context.Context) ClassifierGrokClassifierOutput {
	return o
}

func init() { pulumi.RegisterOutputType(ClassifierGrokClassifierOutput{}) }

type ClassifierJsonClassifier struct {
	// A `JsonPath` string defining the JSON data for the classifier to classify. AWS Glue supports a subset of `JsonPath`, as described in [Writing JsonPath Custom Classifiers](https://docs.aws.amazon.com/glue/latest/dg/custom-classifier.html#custom-classifier-json).
	JsonPath string `pulumi:"jsonPath"`
}
var classifierJsonClassifierType = reflect.TypeOf((*ClassifierJsonClassifier)(nil)).Elem()

type ClassifierJsonClassifierInput interface {
	pulumi.Input

	ToClassifierJsonClassifierOutput() ClassifierJsonClassifierOutput
	ToClassifierJsonClassifierOutputWithContext(ctx context.Context) ClassifierJsonClassifierOutput
}

type ClassifierJsonClassifierArgs struct {
	// A `JsonPath` string defining the JSON data for the classifier to classify. AWS Glue supports a subset of `JsonPath`, as described in [Writing JsonPath Custom Classifiers](https://docs.aws.amazon.com/glue/latest/dg/custom-classifier.html#custom-classifier-json).
	JsonPath pulumi.StringInput `pulumi:"jsonPath"`
}

func (ClassifierJsonClassifierArgs) ElementType() reflect.Type {
	return classifierJsonClassifierType
}

func (a ClassifierJsonClassifierArgs) ToClassifierJsonClassifierOutput() ClassifierJsonClassifierOutput {
	return pulumi.ToOutput(a).(ClassifierJsonClassifierOutput)
}

func (a ClassifierJsonClassifierArgs) ToClassifierJsonClassifierOutputWithContext(ctx context.Context) ClassifierJsonClassifierOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ClassifierJsonClassifierOutput)
}

type ClassifierJsonClassifierOutput struct { *pulumi.OutputState }

// A `JsonPath` string defining the JSON data for the classifier to classify. AWS Glue supports a subset of `JsonPath`, as described in [Writing JsonPath Custom Classifiers](https://docs.aws.amazon.com/glue/latest/dg/custom-classifier.html#custom-classifier-json).
func (o ClassifierJsonClassifierOutput) JsonPath() pulumi.StringOutput {
	return o.Apply(func(v ClassifierJsonClassifier) string {
		return v.JsonPath
	}).(pulumi.StringOutput)
}

func (ClassifierJsonClassifierOutput) ElementType() reflect.Type {
	return classifierJsonClassifierType
}

func (o ClassifierJsonClassifierOutput) ToClassifierJsonClassifierOutput() ClassifierJsonClassifierOutput {
	return o
}

func (o ClassifierJsonClassifierOutput) ToClassifierJsonClassifierOutputWithContext(ctx context.Context) ClassifierJsonClassifierOutput {
	return o
}

func init() { pulumi.RegisterOutputType(ClassifierJsonClassifierOutput{}) }

type ClassifierXmlClassifier struct {
	// An identifier of the data format that the classifier matches.
	Classification string `pulumi:"classification"`
	// The XML tag designating the element that contains each record in an XML document being parsed. Note that this cannot identify a self-closing element (closed by `/>`). An empty row element that contains only attributes can be parsed as long as it ends with a closing tag (for example, `<row item_a="A" item_b="B"></row>` is okay, but `<row item_a="A" item_b="B" />` is not).
	RowTag string `pulumi:"rowTag"`
}
var classifierXmlClassifierType = reflect.TypeOf((*ClassifierXmlClassifier)(nil)).Elem()

type ClassifierXmlClassifierInput interface {
	pulumi.Input

	ToClassifierXmlClassifierOutput() ClassifierXmlClassifierOutput
	ToClassifierXmlClassifierOutputWithContext(ctx context.Context) ClassifierXmlClassifierOutput
}

type ClassifierXmlClassifierArgs struct {
	// An identifier of the data format that the classifier matches.
	Classification pulumi.StringInput `pulumi:"classification"`
	// The XML tag designating the element that contains each record in an XML document being parsed. Note that this cannot identify a self-closing element (closed by `/>`). An empty row element that contains only attributes can be parsed as long as it ends with a closing tag (for example, `<row item_a="A" item_b="B"></row>` is okay, but `<row item_a="A" item_b="B" />` is not).
	RowTag pulumi.StringInput `pulumi:"rowTag"`
}

func (ClassifierXmlClassifierArgs) ElementType() reflect.Type {
	return classifierXmlClassifierType
}

func (a ClassifierXmlClassifierArgs) ToClassifierXmlClassifierOutput() ClassifierXmlClassifierOutput {
	return pulumi.ToOutput(a).(ClassifierXmlClassifierOutput)
}

func (a ClassifierXmlClassifierArgs) ToClassifierXmlClassifierOutputWithContext(ctx context.Context) ClassifierXmlClassifierOutput {
	return pulumi.ToOutputWithContext(ctx, a).(ClassifierXmlClassifierOutput)
}

type ClassifierXmlClassifierOutput struct { *pulumi.OutputState }

// An identifier of the data format that the classifier matches.
func (o ClassifierXmlClassifierOutput) Classification() pulumi.StringOutput {
	return o.Apply(func(v ClassifierXmlClassifier) string {
		return v.Classification
	}).(pulumi.StringOutput)
}

// The XML tag designating the element that contains each record in an XML document being parsed. Note that this cannot identify a self-closing element (closed by `/>`). An empty row element that contains only attributes can be parsed as long as it ends with a closing tag (for example, `<row item_a="A" item_b="B"></row>` is okay, but `<row item_a="A" item_b="B" />` is not).
func (o ClassifierXmlClassifierOutput) RowTag() pulumi.StringOutput {
	return o.Apply(func(v ClassifierXmlClassifier) string {
		return v.RowTag
	}).(pulumi.StringOutput)
}

func (ClassifierXmlClassifierOutput) ElementType() reflect.Type {
	return classifierXmlClassifierType
}

func (o ClassifierXmlClassifierOutput) ToClassifierXmlClassifierOutput() ClassifierXmlClassifierOutput {
	return o
}

func (o ClassifierXmlClassifierOutput) ToClassifierXmlClassifierOutputWithContext(ctx context.Context) ClassifierXmlClassifierOutput {
	return o
}

func init() { pulumi.RegisterOutputType(ClassifierXmlClassifierOutput{}) }

