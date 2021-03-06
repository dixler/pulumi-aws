# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class TopicRule(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the topic rule
    """
    cloudwatch_alarm: pulumi.Output[dict]
    cloudwatch_metric: pulumi.Output[dict]
    description: pulumi.Output[str]
    """
    The description of the rule.
    """
    dynamodb: pulumi.Output[dict]
    elasticsearch: pulumi.Output[dict]
    enabled: pulumi.Output[bool]
    """
    Specifies whether the rule is enabled.
    """
    firehose: pulumi.Output[dict]
    kinesis: pulumi.Output[dict]
    lambda_: pulumi.Output[dict]
    name: pulumi.Output[str]
    """
    The name of the rule.
    """
    republish: pulumi.Output[dict]
    s3: pulumi.Output[dict]
    sns: pulumi.Output[dict]
    sql: pulumi.Output[str]
    """
    The SQL statement used to query the topic. For more information, see AWS IoT SQL Reference (http://docs.aws.amazon.com/iot/latest/developerguide/iot-rules.html#aws-iot-sql-reference) in the AWS IoT Developer Guide.
    """
    sql_version: pulumi.Output[str]
    """
    The version of the SQL rules engine to use when evaluating the rule.
    """
    sqs: pulumi.Output[dict]
    def __init__(__self__, resource_name, opts=None, cloudwatch_alarm=None, cloudwatch_metric=None, description=None, dynamodb=None, elasticsearch=None, enabled=None, firehose=None, kinesis=None, lambda_=None, name=None, republish=None, s3=None, sns=None, sql=None, sql_version=None, sqs=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a TopicRule resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the rule.
        :param pulumi.Input[bool] enabled: Specifies whether the rule is enabled.
        :param pulumi.Input[str] name: The name of the rule.
        :param pulumi.Input[str] sql: The SQL statement used to query the topic. For more information, see AWS IoT SQL Reference (http://docs.aws.amazon.com/iot/latest/developerguide/iot-rules.html#aws-iot-sql-reference) in the AWS IoT Developer Guide.
        :param pulumi.Input[str] sql_version: The version of the SQL rules engine to use when evaluating the rule.
        
        The **cloudwatch_alarm** object supports the following:
        
          * `alarmName` (`pulumi.Input[str]`) - The CloudWatch alarm name.
          * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM role that grants access.
          * `stateReason` (`pulumi.Input[str]`) - The reason for the alarm change.
          * `stateValue` (`pulumi.Input[str]`) - The value of the alarm state. Acceptable values are: OK, ALARM, INSUFFICIENT_DATA.
        
        The **cloudwatch_metric** object supports the following:
        
          * `metric_name` (`pulumi.Input[str]`) - The CloudWatch metric name.
          * `metricNamespace` (`pulumi.Input[str]`) - The CloudWatch metric namespace name.
          * `metricTimestamp` (`pulumi.Input[str]`) - An optional Unix timestamp (http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#about_timestamp).
          * `metricUnit` (`pulumi.Input[str]`) - The metric unit (supported units can be found here: http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#Unit)
          * `metricValue` (`pulumi.Input[str]`) - The CloudWatch metric value.
          * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM role that grants access.
        
        The **dynamodb** object supports the following:
        
          * `hashKeyField` (`pulumi.Input[str]`) - The hash key name.
          * `hashKeyType` (`pulumi.Input[str]`) - The hash key type. Valid values are "STRING" or "NUMBER".
          * `hashKeyValue` (`pulumi.Input[str]`) - The hash key value.
          * `payloadField` (`pulumi.Input[str]`) - The action payload.
          * `rangeKeyField` (`pulumi.Input[str]`) - The range key name.
          * `rangeKeyType` (`pulumi.Input[str]`) - The range key type. Valid values are "STRING" or "NUMBER".
          * `rangeKeyValue` (`pulumi.Input[str]`) - The range key value.
          * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM role that grants access.
          * `table_name` (`pulumi.Input[str]`) - The name of the DynamoDB table.
        
        The **elasticsearch** object supports the following:
        
          * `endpoint` (`pulumi.Input[str]`) - The endpoint of your Elasticsearch domain.
          * `id` (`pulumi.Input[str]`) - The unique identifier for the document you are storing.
          * `index` (`pulumi.Input[str]`) - The Elasticsearch index where you want to store your data.
          * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM role that grants access.
          * `type` (`pulumi.Input[str]`) - The type of document you are storing.
        
        The **firehose** object supports the following:
        
          * `deliveryStreamName` (`pulumi.Input[str]`) - The delivery stream name.
          * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM role that grants access.
          * `separator` (`pulumi.Input[str]`) - A character separator that is used to separate records written to the Firehose stream. Valid values are: '\n' (newline), '\t' (tab), '\r\n' (Windows newline), ',' (comma).
        
        The **kinesis** object supports the following:
        
          * `partitionKey` (`pulumi.Input[str]`) - The partition key.
          * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM role that grants access.
          * `streamName` (`pulumi.Input[str]`) - The name of the Amazon Kinesis stream.
        
        The **lambda_** object supports the following:
        
          * `functionArn` (`pulumi.Input[str]`) - The ARN of the Lambda function.
        
        The **republish** object supports the following:
        
          * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM role that grants access.
          * `topic` (`pulumi.Input[str]`) - The name of the MQTT topic the message should be republished to.
        
        The **s3** object supports the following:
        
          * `bucketName` (`pulumi.Input[str]`) - The Amazon S3 bucket name.
          * `key` (`pulumi.Input[str]`) - The object key.
          * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM role that grants access.
        
        The **sns** object supports the following:
        
          * `messageFormat` (`pulumi.Input[str]`) - The message format of the message to publish. Accepted values are "JSON" and "RAW".
          * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM role that grants access.
          * `target_arn` (`pulumi.Input[str]`) - The ARN of the SNS topic.
        
        The **sqs** object supports the following:
        
          * `queueUrl` (`pulumi.Input[str]`) - The URL of the Amazon SQS queue.
          * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM role that grants access.
          * `useBase64` (`pulumi.Input[bool]`) - Specifies whether to use Base64 encoding.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iot_topic_rule.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['cloudwatch_alarm'] = cloudwatch_alarm
            __props__['cloudwatch_metric'] = cloudwatch_metric
            __props__['description'] = description
            __props__['dynamodb'] = dynamodb
            __props__['elasticsearch'] = elasticsearch
            if enabled is None:
                raise TypeError("Missing required property 'enabled'")
            __props__['enabled'] = enabled
            __props__['firehose'] = firehose
            __props__['kinesis'] = kinesis
            __props__['lambda_'] = lambda_
            __props__['name'] = name
            __props__['republish'] = republish
            __props__['s3'] = s3
            __props__['sns'] = sns
            if sql is None:
                raise TypeError("Missing required property 'sql'")
            __props__['sql'] = sql
            if sql_version is None:
                raise TypeError("Missing required property 'sql_version'")
            __props__['sql_version'] = sql_version
            __props__['sqs'] = sqs
            __props__['arn'] = None
        super(TopicRule, __self__).__init__(
            'aws:iot/topicRule:TopicRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, cloudwatch_alarm=None, cloudwatch_metric=None, description=None, dynamodb=None, elasticsearch=None, enabled=None, firehose=None, kinesis=None, lambda_=None, name=None, republish=None, s3=None, sns=None, sql=None, sql_version=None, sqs=None):
        """
        Get an existing TopicRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the topic rule
        :param pulumi.Input[str] description: The description of the rule.
        :param pulumi.Input[bool] enabled: Specifies whether the rule is enabled.
        :param pulumi.Input[str] name: The name of the rule.
        :param pulumi.Input[str] sql: The SQL statement used to query the topic. For more information, see AWS IoT SQL Reference (http://docs.aws.amazon.com/iot/latest/developerguide/iot-rules.html#aws-iot-sql-reference) in the AWS IoT Developer Guide.
        :param pulumi.Input[str] sql_version: The version of the SQL rules engine to use when evaluating the rule.
        
        The **cloudwatch_alarm** object supports the following:
        
          * `alarmName` (`pulumi.Input[str]`) - The CloudWatch alarm name.
          * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM role that grants access.
          * `stateReason` (`pulumi.Input[str]`) - The reason for the alarm change.
          * `stateValue` (`pulumi.Input[str]`) - The value of the alarm state. Acceptable values are: OK, ALARM, INSUFFICIENT_DATA.
        
        The **cloudwatch_metric** object supports the following:
        
          * `metric_name` (`pulumi.Input[str]`) - The CloudWatch metric name.
          * `metricNamespace` (`pulumi.Input[str]`) - The CloudWatch metric namespace name.
          * `metricTimestamp` (`pulumi.Input[str]`) - An optional Unix timestamp (http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#about_timestamp).
          * `metricUnit` (`pulumi.Input[str]`) - The metric unit (supported units can be found here: http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#Unit)
          * `metricValue` (`pulumi.Input[str]`) - The CloudWatch metric value.
          * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM role that grants access.
        
        The **dynamodb** object supports the following:
        
          * `hashKeyField` (`pulumi.Input[str]`) - The hash key name.
          * `hashKeyType` (`pulumi.Input[str]`) - The hash key type. Valid values are "STRING" or "NUMBER".
          * `hashKeyValue` (`pulumi.Input[str]`) - The hash key value.
          * `payloadField` (`pulumi.Input[str]`) - The action payload.
          * `rangeKeyField` (`pulumi.Input[str]`) - The range key name.
          * `rangeKeyType` (`pulumi.Input[str]`) - The range key type. Valid values are "STRING" or "NUMBER".
          * `rangeKeyValue` (`pulumi.Input[str]`) - The range key value.
          * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM role that grants access.
          * `table_name` (`pulumi.Input[str]`) - The name of the DynamoDB table.
        
        The **elasticsearch** object supports the following:
        
          * `endpoint` (`pulumi.Input[str]`) - The endpoint of your Elasticsearch domain.
          * `id` (`pulumi.Input[str]`) - The unique identifier for the document you are storing.
          * `index` (`pulumi.Input[str]`) - The Elasticsearch index where you want to store your data.
          * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM role that grants access.
          * `type` (`pulumi.Input[str]`) - The type of document you are storing.
        
        The **firehose** object supports the following:
        
          * `deliveryStreamName` (`pulumi.Input[str]`) - The delivery stream name.
          * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM role that grants access.
          * `separator` (`pulumi.Input[str]`) - A character separator that is used to separate records written to the Firehose stream. Valid values are: '\n' (newline), '\t' (tab), '\r\n' (Windows newline), ',' (comma).
        
        The **kinesis** object supports the following:
        
          * `partitionKey` (`pulumi.Input[str]`) - The partition key.
          * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM role that grants access.
          * `streamName` (`pulumi.Input[str]`) - The name of the Amazon Kinesis stream.
        
        The **lambda_** object supports the following:
        
          * `functionArn` (`pulumi.Input[str]`) - The ARN of the Lambda function.
        
        The **republish** object supports the following:
        
          * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM role that grants access.
          * `topic` (`pulumi.Input[str]`) - The name of the MQTT topic the message should be republished to.
        
        The **s3** object supports the following:
        
          * `bucketName` (`pulumi.Input[str]`) - The Amazon S3 bucket name.
          * `key` (`pulumi.Input[str]`) - The object key.
          * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM role that grants access.
        
        The **sns** object supports the following:
        
          * `messageFormat` (`pulumi.Input[str]`) - The message format of the message to publish. Accepted values are "JSON" and "RAW".
          * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM role that grants access.
          * `target_arn` (`pulumi.Input[str]`) - The ARN of the SNS topic.
        
        The **sqs** object supports the following:
        
          * `queueUrl` (`pulumi.Input[str]`) - The URL of the Amazon SQS queue.
          * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM role that grants access.
          * `useBase64` (`pulumi.Input[bool]`) - Specifies whether to use Base64 encoding.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iot_topic_rule.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["arn"] = arn
        __props__["cloudwatch_alarm"] = cloudwatch_alarm
        __props__["cloudwatch_metric"] = cloudwatch_metric
        __props__["description"] = description
        __props__["dynamodb"] = dynamodb
        __props__["elasticsearch"] = elasticsearch
        __props__["enabled"] = enabled
        __props__["firehose"] = firehose
        __props__["kinesis"] = kinesis
        __props__["lambda_"] = lambda_
        __props__["name"] = name
        __props__["republish"] = republish
        __props__["s3"] = s3
        __props__["sns"] = sns
        __props__["sql"] = sql
        __props__["sql_version"] = sql_version
        __props__["sqs"] = sqs
        return TopicRule(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

