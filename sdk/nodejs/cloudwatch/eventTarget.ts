// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a CloudWatch Event Target resource.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const console = new aws.cloudwatch.EventRule("console", {
 *     description: "Capture all EC2 scaling events",
 *     eventPattern: `{
 *   "source": [
 *     "aws.autoscaling"
 *   ],
 *   "detail-type": [
 *     "EC2 Instance Launch Successful",
 *     "EC2 Instance Terminate Successful",
 *     "EC2 Instance Launch Unsuccessful",
 *     "EC2 Instance Terminate Unsuccessful"
 *   ]
 * }
 * `,
 * });
 * const testStream = new aws.kinesis.Stream("testStream", {
 *     shardCount: 1,
 * });
 * const yada = new aws.cloudwatch.EventTarget("yada", {
 *     arn: testStream.arn,
 *     rule: console.name,
 *     runCommandTargets: [
 *         {
 *             key: "tag:Name",
 *             values: ["FooBar"],
 *         },
 *         {
 *             key: "InstanceIds",
 *             values: ["i-162058cd308bffec2"],
 *         },
 *     ],
 * });
 * ```
 * 
 * ## Example SSM Document Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const ssmLifecycleTrust = aws.iam.getPolicyDocument({
 *     statements: [{
 *         actions: ["sts:AssumeRole"],
 *         principals: [{
 *             identifiers: ["events.amazonaws.com"],
 *             type: "Service",
 *         }],
 *     }],
 * });
 * const stopInstance = new aws.ssm.Document("stopInstance", {
 *     content: `  {
 *     "schemaVersion": "1.2",
 *     "description": "Stop an instance",
 *     "parameters": {
 * 
 *     },
 *     "runtimeConfig": {
 *       "aws:runShellScript": {
 *         "properties": [
 *           {
 *             "id": "0.aws:runShellScript",
 *             "runCommand": ["halt"]
 *           }
 *         ]
 *       }
 *     }
 *   }
 * `,
 *     documentType: "Command",
 * });
 * const ssmLifecyclePolicyDocument = stopInstance.arn.apply(arn => aws.iam.getPolicyDocument({
 *     statements: [
 *         {
 *             actions: ["ssm:SendCommand"],
 *             conditions: [{
 *                 test: "StringEquals",
 *                 values: ["*"],
 *                 variable: "ec2:ResourceTag/Terminate",
 *             }],
 *             effect: "Allow",
 *             resources: ["arn:aws:ec2:eu-west-1:1234567890:instance/*"],
 *         },
 *         {
 *             actions: ["ssm:SendCommand"],
 *             effect: "Allow",
 *             resources: [arn],
 *         },
 *     ],
 * }));
 * const ssmLifecycleRole = new aws.iam.Role("ssmLifecycle", {
 *     assumeRolePolicy: ssmLifecycleTrust.json,
 * });
 * const ssmLifecyclePolicy = new aws.iam.Policy("ssmLifecycle", {
 *     policy: ssmLifecyclePolicyDocument.json,
 * });
 * const stopInstancesEventRule = new aws.cloudwatch.EventRule("stopInstances", {
 *     description: "Stop instances nightly",
 *     scheduleExpression: "cron(0 0 * * ? *)",
 * });
 * const stopInstancesEventTarget = new aws.cloudwatch.EventTarget("stopInstances", {
 *     arn: stopInstance.arn,
 *     roleArn: ssmLifecycleRole.arn,
 *     rule: stopInstancesEventRule.name,
 *     runCommandTargets: [{
 *         key: "tag:Terminate",
 *         values: ["midnight"],
 *     }],
 * });
 * ```
 * 
 * ## Example RunCommand Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const stopInstancesEventRule = new aws.cloudwatch.EventRule("stopInstances", {
 *     description: "Stop instances nightly",
 *     scheduleExpression: "cron(0 0 * * ? *)",
 * });
 * const stopInstancesEventTarget = new aws.cloudwatch.EventTarget("stopInstances", {
 *     arn: `arn:aws:ssm:${var_aws_region}::document/AWS-RunShellScript`,
 *     input: "{\"commands\":[\"halt\"]}",
 *     roleArn: aws_iam_role_ssm_lifecycle.arn,
 *     rule: stopInstancesEventRule.name,
 *     runCommandTargets: [{
 *         key: "tag:Terminate",
 *         values: ["midnight"],
 *     }],
 * });
 * ```
 * 
 * ## Example ECS Run Task with Role and Task Override Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const ecsEvents = new aws.iam.Role("ecsEvents", {
 *     assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Sid": "",
 *       "Effect": "Allow",
 *       "Principal": {
 *         "Service": "events.amazonaws.com"
 *       },
 *       "Action": "sts:AssumeRole"
 *     }
 *   ]
 * }
 * `,
 * });
 * const ecsEventsRunTaskWithAnyRole = new aws.iam.RolePolicy("ecsEventsRunTaskWithAnyRole", {
 *     policy: aws_ecs_task_definition_task_name.arn.apply(arn => `{
 *     "Version": "2012-10-17",
 *     "Statement": [
 *         {
 *             "Effect": "Allow",
 *             "Action": "iam:PassRole",
 *             "Resource": "*"
 *         },
 *         {
 *             "Effect": "Allow",
 *             "Action": "ecs:RunTask",
 *             "Resource": "${arn.replace(/:\d+$/, ":*")}"
 *         }
 *     ]
 * }
 * `),
 *     role: ecsEvents.id,
 * });
 * const ecsScheduledTask = new aws.cloudwatch.EventTarget("ecsScheduledTask", {
 *     arn: aws_ecs_cluster_cluster_name.arn,
 *     ecsTarget: {
 *         taskCount: 1,
 *         taskDefinitionArn: aws_ecs_task_definition_task_name.arn,
 *     },
 *     input: `{
 *   "containerOverrides": [
 *     {
 *       "name": "name-of-container-to-override",
 *       "command": ["bin/console", "scheduled-task"]
 *     }
 *   ]
 * }
 * `,
 *     roleArn: ecsEvents.arn,
 *     rule: aws_cloudwatch_event_rule_every_hour.name,
 * });
 * ```
 * 
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_event_target.html.markdown.
 */
export class EventTarget extends pulumi.CustomResource {
    /**
     * Get an existing EventTarget resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventTargetState, opts?: pulumi.CustomResourceOptions): EventTarget {
        return new EventTarget(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cloudwatch/eventTarget:EventTarget';

    /**
     * Returns true if the given object is an instance of EventTarget.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventTarget {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventTarget.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) associated of the target.
     */
    public readonly arn!: pulumi.Output<string>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
     */
    public readonly batchTarget!: pulumi.Output<outputs.cloudwatch.EventTargetBatchTarget | undefined>;
    /**
     * Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
     */
    public readonly ecsTarget!: pulumi.Output<outputs.cloudwatch.EventTargetEcsTarget | undefined>;
    /**
     * Valid JSON text passed to the target.
     */
    public readonly input!: pulumi.Output<string | undefined>;
    /**
     * The value of the [JSONPath](http://goessner.net/articles/JsonPath/)
     * that is used for extracting part of the matched event when passing it to the target.
     */
    public readonly inputPath!: pulumi.Output<string | undefined>;
    /**
     * Parameters used when you are providing a custom input to a target based on certain event data.
     */
    public readonly inputTransformer!: pulumi.Output<outputs.cloudwatch.EventTargetInputTransformer | undefined>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
     */
    public readonly kinesisTarget!: pulumi.Output<outputs.cloudwatch.EventTargetKinesisTarget | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecsTarget` is used.
     */
    public readonly roleArn!: pulumi.Output<string | undefined>;
    /**
     * The name of the rule you want to add targets to.
     */
    public readonly rule!: pulumi.Output<string>;
    /**
     * Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
     */
    public readonly runCommandTargets!: pulumi.Output<outputs.cloudwatch.EventTargetRunCommandTarget[] | undefined>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
     */
    public readonly sqsTarget!: pulumi.Output<outputs.cloudwatch.EventTargetSqsTarget | undefined>;
    /**
     * The unique target assignment ID.  If missing, will generate a random, unique id.
     */
    public readonly targetId!: pulumi.Output<string>;

    /**
     * Create a EventTarget resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventTargetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventTargetArgs | EventTargetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EventTargetState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["batchTarget"] = state ? state.batchTarget : undefined;
            inputs["ecsTarget"] = state ? state.ecsTarget : undefined;
            inputs["input"] = state ? state.input : undefined;
            inputs["inputPath"] = state ? state.inputPath : undefined;
            inputs["inputTransformer"] = state ? state.inputTransformer : undefined;
            inputs["kinesisTarget"] = state ? state.kinesisTarget : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
            inputs["rule"] = state ? state.rule : undefined;
            inputs["runCommandTargets"] = state ? state.runCommandTargets : undefined;
            inputs["sqsTarget"] = state ? state.sqsTarget : undefined;
            inputs["targetId"] = state ? state.targetId : undefined;
        } else {
            const args = argsOrState as EventTargetArgs | undefined;
            if (!args || args.arn === undefined) {
                throw new Error("Missing required property 'arn'");
            }
            if (!args || args.rule === undefined) {
                throw new Error("Missing required property 'rule'");
            }
            inputs["arn"] = args ? args.arn : undefined;
            inputs["batchTarget"] = args ? args.batchTarget : undefined;
            inputs["ecsTarget"] = args ? args.ecsTarget : undefined;
            inputs["input"] = args ? args.input : undefined;
            inputs["inputPath"] = args ? args.inputPath : undefined;
            inputs["inputTransformer"] = args ? args.inputTransformer : undefined;
            inputs["kinesisTarget"] = args ? args.kinesisTarget : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["rule"] = args ? args.rule : undefined;
            inputs["runCommandTargets"] = args ? args.runCommandTargets : undefined;
            inputs["sqsTarget"] = args ? args.sqsTarget : undefined;
            inputs["targetId"] = args ? args.targetId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(EventTarget.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventTarget resources.
 */
export interface EventTargetState {
    /**
     * The Amazon Resource Name (ARN) associated of the target.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
     */
    readonly batchTarget?: pulumi.Input<inputs.cloudwatch.EventTargetBatchTarget>;
    /**
     * Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
     */
    readonly ecsTarget?: pulumi.Input<inputs.cloudwatch.EventTargetEcsTarget>;
    /**
     * Valid JSON text passed to the target.
     */
    readonly input?: pulumi.Input<string>;
    /**
     * The value of the [JSONPath](http://goessner.net/articles/JsonPath/)
     * that is used for extracting part of the matched event when passing it to the target.
     */
    readonly inputPath?: pulumi.Input<string>;
    /**
     * Parameters used when you are providing a custom input to a target based on certain event data.
     */
    readonly inputTransformer?: pulumi.Input<inputs.cloudwatch.EventTargetInputTransformer>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
     */
    readonly kinesisTarget?: pulumi.Input<inputs.cloudwatch.EventTargetKinesisTarget>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecsTarget` is used.
     */
    readonly roleArn?: pulumi.Input<string>;
    /**
     * The name of the rule you want to add targets to.
     */
    readonly rule?: pulumi.Input<string>;
    /**
     * Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
     */
    readonly runCommandTargets?: pulumi.Input<pulumi.Input<inputs.cloudwatch.EventTargetRunCommandTarget>[]>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
     */
    readonly sqsTarget?: pulumi.Input<inputs.cloudwatch.EventTargetSqsTarget>;
    /**
     * The unique target assignment ID.  If missing, will generate a random, unique id.
     */
    readonly targetId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EventTarget resource.
 */
export interface EventTargetArgs {
    /**
     * The Amazon Resource Name (ARN) associated of the target.
     */
    readonly arn: pulumi.Input<string>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
     */
    readonly batchTarget?: pulumi.Input<inputs.cloudwatch.EventTargetBatchTarget>;
    /**
     * Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
     */
    readonly ecsTarget?: pulumi.Input<inputs.cloudwatch.EventTargetEcsTarget>;
    /**
     * Valid JSON text passed to the target.
     */
    readonly input?: pulumi.Input<string>;
    /**
     * The value of the [JSONPath](http://goessner.net/articles/JsonPath/)
     * that is used for extracting part of the matched event when passing it to the target.
     */
    readonly inputPath?: pulumi.Input<string>;
    /**
     * Parameters used when you are providing a custom input to a target based on certain event data.
     */
    readonly inputTransformer?: pulumi.Input<inputs.cloudwatch.EventTargetInputTransformer>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
     */
    readonly kinesisTarget?: pulumi.Input<inputs.cloudwatch.EventTargetKinesisTarget>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecsTarget` is used.
     */
    readonly roleArn?: pulumi.Input<string>;
    /**
     * The name of the rule you want to add targets to.
     */
    readonly rule: pulumi.Input<string>;
    /**
     * Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
     */
    readonly runCommandTargets?: pulumi.Input<pulumi.Input<inputs.cloudwatch.EventTargetRunCommandTarget>[]>;
    /**
     * Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
     */
    readonly sqsTarget?: pulumi.Input<inputs.cloudwatch.EventTargetSqsTarget>;
    /**
     * The unique target assignment ID.  If missing, will generate a random, unique id.
     */
    readonly targetId?: pulumi.Input<string>;
}
