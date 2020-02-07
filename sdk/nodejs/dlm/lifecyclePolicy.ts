// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a [Data Lifecycle Manager (DLM) lifecycle policy](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/snapshot-lifecycle.html) for managing snapshots.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const dlmLifecycleRole = new aws.iam.Role("dlmLifecycleRole", {
 *     assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": "sts:AssumeRole",
 *       "Principal": {
 *         "Service": "dlm.amazonaws.com"
 *       },
 *       "Effect": "Allow",
 *       "Sid": ""
 *     }
 *   ]
 * }
 * `,
 * });
 * const dlmLifecycle = new aws.iam.RolePolicy("dlmLifecycle", {
 *     policy: `{
 *    "Version": "2012-10-17",
 *    "Statement": [
 *       {
 *          "Effect": "Allow",
 *          "Action": [
 *             "ec2:CreateSnapshot",
 *             "ec2:DeleteSnapshot",
 *             "ec2:DescribeVolumes",
 *             "ec2:DescribeSnapshots"
 *          ],
 *          "Resource": "*"
 *       },
 *       {
 *          "Effect": "Allow",
 *          "Action": [
 *             "ec2:CreateTags"
 *          ],
 *          "Resource": "arn:aws:ec2:*::snapshot/*"
 *       }
 *    ]
 * }
 * `,
 *     role: dlmLifecycleRole.id,
 * });
 * const example = new aws.dlm.LifecyclePolicy("example", {
 *     description: "example DLM lifecycle policy",
 *     executionRoleArn: dlmLifecycleRole.arn,
 *     policyDetails: {
 *         resourceTypes: ["VOLUME"],
 *         schedules: [{
 *             copyTags: false,
 *             createRule: {
 *                 interval: 24,
 *                 intervalUnit: "HOURS",
 *                 times: "23:45",
 *             },
 *             name: "2 weeks of daily snapshots",
 *             retainRule: {
 *                 count: 14,
 *             },
 *             tagsToAdd: {
 *                 SnapshotCreator: "DLM",
 *             },
 *         }],
 *         targetTags: {
 *             Snapshot: "true",
 *         },
 *     },
 *     state: "ENABLED",
 * });
 * ```
 * 
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dlm_lifecycle_policy.html.markdown.
 */
export class LifecyclePolicy extends pulumi.CustomResource {
    /**
     * Get an existing LifecyclePolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LifecyclePolicyState, opts?: pulumi.CustomResourceOptions): LifecyclePolicy {
        return new LifecyclePolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:dlm/lifecyclePolicy:LifecyclePolicy';

    /**
     * Returns true if the given object is an instance of LifecyclePolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LifecyclePolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LifecyclePolicy.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of the DLM Lifecycle Policy.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A description for the DLM lifecycle policy.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The ARN of an IAM role that is able to be assumed by the DLM service.
     */
    public readonly executionRoleArn!: pulumi.Output<string>;
    /**
     * See the `policyDetails` configuration block. Max of 1.
     */
    public readonly policyDetails!: pulumi.Output<outputs.dlm.LifecyclePolicyPolicyDetails>;
    /**
     * Whether the lifecycle policy should be enabled or disabled. `ENABLED` or `DISABLED` are valid values. Defaults to `ENABLED`.
     */
    public readonly state!: pulumi.Output<string | undefined>;
    /**
     * Key-value mapping of resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a LifecyclePolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LifecyclePolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LifecyclePolicyArgs | LifecyclePolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LifecyclePolicyState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["executionRoleArn"] = state ? state.executionRoleArn : undefined;
            inputs["policyDetails"] = state ? state.policyDetails : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as LifecyclePolicyArgs | undefined;
            if (!args || args.description === undefined) {
                throw new Error("Missing required property 'description'");
            }
            if (!args || args.executionRoleArn === undefined) {
                throw new Error("Missing required property 'executionRoleArn'");
            }
            if (!args || args.policyDetails === undefined) {
                throw new Error("Missing required property 'policyDetails'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["executionRoleArn"] = args ? args.executionRoleArn : undefined;
            inputs["policyDetails"] = args ? args.policyDetails : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(LifecyclePolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LifecyclePolicy resources.
 */
export interface LifecyclePolicyState {
    /**
     * Amazon Resource Name (ARN) of the DLM Lifecycle Policy.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * A description for the DLM lifecycle policy.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The ARN of an IAM role that is able to be assumed by the DLM service.
     */
    readonly executionRoleArn?: pulumi.Input<string>;
    /**
     * See the `policyDetails` configuration block. Max of 1.
     */
    readonly policyDetails?: pulumi.Input<inputs.dlm.LifecyclePolicyPolicyDetails>;
    /**
     * Whether the lifecycle policy should be enabled or disabled. `ENABLED` or `DISABLED` are valid values. Defaults to `ENABLED`.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * Key-value mapping of resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a LifecyclePolicy resource.
 */
export interface LifecyclePolicyArgs {
    /**
     * A description for the DLM lifecycle policy.
     */
    readonly description: pulumi.Input<string>;
    /**
     * The ARN of an IAM role that is able to be assumed by the DLM service.
     */
    readonly executionRoleArn: pulumi.Input<string>;
    /**
     * See the `policyDetails` configuration block. Max of 1.
     */
    readonly policyDetails: pulumi.Input<inputs.dlm.LifecyclePolicyPolicyDetails>;
    /**
     * Whether the lifecycle policy should be enabled or disabled. `ENABLED` or `DISABLED` are valid values. Defaults to `ENABLED`.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * Key-value mapping of resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
