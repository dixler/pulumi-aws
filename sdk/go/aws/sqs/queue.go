// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sqs

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type Queue struct {
	s *pulumi.ResourceState
}

// NewQueue registers a new resource with the given unique name, arguments, and options.
func NewQueue(ctx *pulumi.Context,
	name string, args *QueueArgs, opts ...pulumi.ResourceOpt) (*Queue, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["contentBasedDeduplication"] = nil
		inputs["delaySeconds"] = nil
		inputs["fifoQueue"] = nil
		inputs["kmsDataKeyReusePeriodSeconds"] = nil
		inputs["kmsMasterKeyId"] = nil
		inputs["maxMessageSize"] = nil
		inputs["messageRetentionSeconds"] = nil
		inputs["name"] = nil
		inputs["namePrefix"] = nil
		inputs["policy"] = nil
		inputs["receiveWaitTimeSeconds"] = nil
		inputs["redrivePolicy"] = nil
		inputs["tags"] = nil
		inputs["visibilityTimeoutSeconds"] = nil
	} else {
		inputs["contentBasedDeduplication"] = args.ContentBasedDeduplication
		inputs["delaySeconds"] = args.DelaySeconds
		inputs["fifoQueue"] = args.FifoQueue
		inputs["kmsDataKeyReusePeriodSeconds"] = args.KmsDataKeyReusePeriodSeconds
		inputs["kmsMasterKeyId"] = args.KmsMasterKeyId
		inputs["maxMessageSize"] = args.MaxMessageSize
		inputs["messageRetentionSeconds"] = args.MessageRetentionSeconds
		inputs["name"] = args.Name
		inputs["namePrefix"] = args.NamePrefix
		inputs["policy"] = args.Policy
		inputs["receiveWaitTimeSeconds"] = args.ReceiveWaitTimeSeconds
		inputs["redrivePolicy"] = args.RedrivePolicy
		inputs["tags"] = args.Tags
		inputs["visibilityTimeoutSeconds"] = args.VisibilityTimeoutSeconds
	}
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:sqs/queue:Queue", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Queue{s: s}, nil
}

// GetQueue gets an existing Queue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueue(ctx *pulumi.Context,
	name string, id pulumi.ID, state *QueueState, opts ...pulumi.ResourceOpt) (*Queue, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["contentBasedDeduplication"] = state.ContentBasedDeduplication
		inputs["delaySeconds"] = state.DelaySeconds
		inputs["fifoQueue"] = state.FifoQueue
		inputs["kmsDataKeyReusePeriodSeconds"] = state.KmsDataKeyReusePeriodSeconds
		inputs["kmsMasterKeyId"] = state.KmsMasterKeyId
		inputs["maxMessageSize"] = state.MaxMessageSize
		inputs["messageRetentionSeconds"] = state.MessageRetentionSeconds
		inputs["name"] = state.Name
		inputs["namePrefix"] = state.NamePrefix
		inputs["policy"] = state.Policy
		inputs["receiveWaitTimeSeconds"] = state.ReceiveWaitTimeSeconds
		inputs["redrivePolicy"] = state.RedrivePolicy
		inputs["tags"] = state.Tags
		inputs["visibilityTimeoutSeconds"] = state.VisibilityTimeoutSeconds
	}
	s, err := ctx.ReadResource("aws:sqs/queue:Queue", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Queue{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Queue) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Queue) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The ARN of the SQS queue
func (r *Queue) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// Enables content-based deduplication for FIFO queues. For more information, see the [related documentation](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-exactly-once-processing)
func (r *Queue) ContentBasedDeduplication() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["contentBasedDeduplication"])
}

// The time in seconds that the delivery of all messages in the queue will be delayed. An integer from 0 to 900 (15 minutes). The default for this attribute is 0 seconds.
func (r *Queue) DelaySeconds() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["delaySeconds"])
}

// Boolean designating a FIFO queue. If not set, it defaults to `false` making it standard.
func (r *Queue) FifoQueue() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["fifoQueue"])
}

// The length of time, in seconds, for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. An integer representing seconds, between 60 seconds (1 minute) and 86,400 seconds (24 hours). The default is 300 (5 minutes).
func (r *Queue) KmsDataKeyReusePeriodSeconds() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["kmsDataKeyReusePeriodSeconds"])
}

// The ID of an AWS-managed customer master key (CMK) for Amazon SQS or a custom CMK. For more information, see [Key Terms](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-sse-key-terms).
func (r *Queue) KmsMasterKeyId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["kmsMasterKeyId"])
}

// The limit of how many bytes a message can contain before Amazon SQS rejects it. An integer from 1024 bytes (1 KiB) up to 262144 bytes (256 KiB). The default for this attribute is 262144 (256 KiB).
func (r *Queue) MaxMessageSize() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["maxMessageSize"])
}

// The number of seconds Amazon SQS retains a message. Integer representing seconds, from 60 (1 minute) to 1209600 (14 days). The default for this attribute is 345600 (4 days).
func (r *Queue) MessageRetentionSeconds() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["messageRetentionSeconds"])
}

// This is the human-readable name of the queue. If omitted, this provider will assign a random name.
func (r *Queue) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
func (r *Queue) NamePrefix() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["namePrefix"])
}

func (r *Queue) Policy() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["policy"])
}

// The time for which a ReceiveMessage call will wait for a message to arrive (long polling) before returning. An integer from 0 to 20 (seconds). The default for this attribute is 0, meaning that the call will return immediately.
func (r *Queue) ReceiveWaitTimeSeconds() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["receiveWaitTimeSeconds"])
}

// The JSON policy to set up the Dead Letter Queue, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/SQSDeadLetterQueue.html). **Note:** when specifying `maxReceiveCount`, you must specify it as an integer (`5`), and not a string (`"5"`).
func (r *Queue) RedrivePolicy() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["redrivePolicy"])
}

// A mapping of tags to assign to the queue.
func (r *Queue) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// The visibility timeout for the queue. An integer from 0 to 43200 (12 hours). The default for this attribute is 30. For more information about visibility timeout, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/AboutVT.html).
func (r *Queue) VisibilityTimeoutSeconds() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["visibilityTimeoutSeconds"])
}

// Input properties used for looking up and filtering Queue resources.
type QueueState struct {
	// The ARN of the SQS queue
	Arn interface{}
	// Enables content-based deduplication for FIFO queues. For more information, see the [related documentation](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-exactly-once-processing)
	ContentBasedDeduplication interface{}
	// The time in seconds that the delivery of all messages in the queue will be delayed. An integer from 0 to 900 (15 minutes). The default for this attribute is 0 seconds.
	DelaySeconds interface{}
	// Boolean designating a FIFO queue. If not set, it defaults to `false` making it standard.
	FifoQueue interface{}
	// The length of time, in seconds, for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. An integer representing seconds, between 60 seconds (1 minute) and 86,400 seconds (24 hours). The default is 300 (5 minutes).
	KmsDataKeyReusePeriodSeconds interface{}
	// The ID of an AWS-managed customer master key (CMK) for Amazon SQS or a custom CMK. For more information, see [Key Terms](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-sse-key-terms).
	KmsMasterKeyId interface{}
	// The limit of how many bytes a message can contain before Amazon SQS rejects it. An integer from 1024 bytes (1 KiB) up to 262144 bytes (256 KiB). The default for this attribute is 262144 (256 KiB).
	MaxMessageSize interface{}
	// The number of seconds Amazon SQS retains a message. Integer representing seconds, from 60 (1 minute) to 1209600 (14 days). The default for this attribute is 345600 (4 days).
	MessageRetentionSeconds interface{}
	// This is the human-readable name of the queue. If omitted, this provider will assign a random name.
	Name interface{}
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix interface{}
	Policy interface{}
	// The time for which a ReceiveMessage call will wait for a message to arrive (long polling) before returning. An integer from 0 to 20 (seconds). The default for this attribute is 0, meaning that the call will return immediately.
	ReceiveWaitTimeSeconds interface{}
	// The JSON policy to set up the Dead Letter Queue, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/SQSDeadLetterQueue.html). **Note:** when specifying `maxReceiveCount`, you must specify it as an integer (`5`), and not a string (`"5"`).
	RedrivePolicy interface{}
	// A mapping of tags to assign to the queue.
	Tags interface{}
	// The visibility timeout for the queue. An integer from 0 to 43200 (12 hours). The default for this attribute is 30. For more information about visibility timeout, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/AboutVT.html).
	VisibilityTimeoutSeconds interface{}
}

// The set of arguments for constructing a Queue resource.
type QueueArgs struct {
	// Enables content-based deduplication for FIFO queues. For more information, see the [related documentation](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-exactly-once-processing)
	ContentBasedDeduplication interface{}
	// The time in seconds that the delivery of all messages in the queue will be delayed. An integer from 0 to 900 (15 minutes). The default for this attribute is 0 seconds.
	DelaySeconds interface{}
	// Boolean designating a FIFO queue. If not set, it defaults to `false` making it standard.
	FifoQueue interface{}
	// The length of time, in seconds, for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. An integer representing seconds, between 60 seconds (1 minute) and 86,400 seconds (24 hours). The default is 300 (5 minutes).
	KmsDataKeyReusePeriodSeconds interface{}
	// The ID of an AWS-managed customer master key (CMK) for Amazon SQS or a custom CMK. For more information, see [Key Terms](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-sse-key-terms).
	KmsMasterKeyId interface{}
	// The limit of how many bytes a message can contain before Amazon SQS rejects it. An integer from 1024 bytes (1 KiB) up to 262144 bytes (256 KiB). The default for this attribute is 262144 (256 KiB).
	MaxMessageSize interface{}
	// The number of seconds Amazon SQS retains a message. Integer representing seconds, from 60 (1 minute) to 1209600 (14 days). The default for this attribute is 345600 (4 days).
	MessageRetentionSeconds interface{}
	// This is the human-readable name of the queue. If omitted, this provider will assign a random name.
	Name interface{}
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix interface{}
	Policy interface{}
	// The time for which a ReceiveMessage call will wait for a message to arrive (long polling) before returning. An integer from 0 to 20 (seconds). The default for this attribute is 0, meaning that the call will return immediately.
	ReceiveWaitTimeSeconds interface{}
	// The JSON policy to set up the Dead Letter Queue, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/SQSDeadLetterQueue.html). **Note:** when specifying `maxReceiveCount`, you must specify it as an integer (`5`), and not a string (`"5"`).
	RedrivePolicy interface{}
	// A mapping of tags to assign to the queue.
	Tags interface{}
	// The visibility timeout for the queue. An integer from 0 to 43200 (12 hours). The default for this attribute is 30. For more information about visibility timeout, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/AboutVT.html).
	VisibilityTimeoutSeconds interface{}
}
