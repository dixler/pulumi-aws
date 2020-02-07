// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an SNS topic policy resource
 * 
 * > **NOTE:** If a Principal is specified as just an AWS account ID rather than an ARN, AWS silently converts it to the ARN for the root user, causing future deployments to differ. To avoid this problem, just specify the full ARN, e.g. `arn:aws:iam::123456789012:root`
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const test = new aws.sns.Topic("test", {});
 * const snsTopicPolicy = test.arn.apply(arn => aws.iam.getPolicyDocument({
 *     policyId: "__default_policy_ID",
 *     statements: [{
 *         actions: [
 *             "SNS:Subscribe",
 *             "SNS:SetTopicAttributes",
 *             "SNS:RemovePermission",
 *             "SNS:Receive",
 *             "SNS:Publish",
 *             "SNS:ListSubscriptionsByTopic",
 *             "SNS:GetTopicAttributes",
 *             "SNS:DeleteTopic",
 *             "SNS:AddPermission",
 *         ],
 *         conditions: [{
 *             test: "StringEquals",
 *             values: [varAccountId],
 *             variable: "AWS:SourceOwner",
 *         }],
 *         effect: "Allow",
 *         principals: [{
 *             identifiers: ["*"],
 *             type: "AWS",
 *         }],
 *         resources: [arn],
 *         sid: "__default_statement_ID",
 *     }],
 * }));
 * const defaultTopicPolicy = new aws.sns.TopicPolicy("default", {
 *     arn: test.arn,
 *     policy: sns_topic_policy.json,
 * });
 * ```
 * 
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/sns_topic_policy.html.markdown.
 */
export class TopicPolicy extends pulumi.CustomResource {
    /**
     * Get an existing TopicPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TopicPolicyState, opts?: pulumi.CustomResourceOptions): TopicPolicy {
        return new TopicPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:sns/topicPolicy:TopicPolicy';

    /**
     * Returns true if the given object is an instance of TopicPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TopicPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TopicPolicy.__pulumiType;
    }

    /**
     * The ARN of the SNS topic
     */
    public readonly arn!: pulumi.Output<string>;
    public readonly policy!: pulumi.Output<string>;

    /**
     * Create a TopicPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TopicPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TopicPolicyArgs | TopicPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TopicPolicyState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["policy"] = state ? state.policy : undefined;
        } else {
            const args = argsOrState as TopicPolicyArgs | undefined;
            if (!args || args.arn === undefined) {
                throw new Error("Missing required property 'arn'");
            }
            if (!args || args.policy === undefined) {
                throw new Error("Missing required property 'policy'");
            }
            inputs["arn"] = args ? args.arn : undefined;
            inputs["policy"] = args ? args.policy : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TopicPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TopicPolicy resources.
 */
export interface TopicPolicyState {
    /**
     * The ARN of the SNS topic
     */
    readonly arn?: pulumi.Input<string>;
    readonly policy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TopicPolicy resource.
 */
export interface TopicPolicyArgs {
    /**
     * The ARN of the SNS topic
     */
    readonly arn: pulumi.Input<string>;
    readonly policy: pulumi.Input<string>;
}
