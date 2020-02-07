// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an WAF Regional Rule Resource for use with Application Load Balancer.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const ipset = new aws.wafregional.IpSet("ipset", {
 *     ipSetDescriptors: [{
 *         type: "IPV4",
 *         value: "192.0.7.0/24",
 *     }],
 * });
 * const wafrule = new aws.wafregional.Rule("wafrule", {
 *     metricName: "tfWAFRule",
 *     predicates: [{
 *         dataId: ipset.id,
 *         negated: false,
 *         type: "IPMatch",
 *     }],
 * });
 * ```
 * 
 * ## Nested Fields
 * 
 * ### `predicate`
 * 
 * See the [WAF Documentation](https://docs.aws.amazon.com/waf/latest/APIReference/API_Predicate.html) for more information.
 * 
 * #### Arguments
 * 
 * * `type` - (Required) The type of predicate in a rule. Valid values: `ByteMatch`, `GeoMatch`, `IPMatch`, `RegexMatch`, `SizeConstraint`, `SqlInjectionMatch`, or `XssMatch`
 * * `dataId` - (Required) The unique identifier of a predicate, such as the ID of a `ByteMatchSet` or `IPSet`.
 * * `negated` - (Required) Whether to use the settings or the negated settings that you specified in the objects.
 * 
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/wafregional_rule.html.markdown.
 */
export class Rule extends pulumi.CustomResource {
    /**
     * Get an existing Rule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleState, opts?: pulumi.CustomResourceOptions): Rule {
        return new Rule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:wafregional/rule:Rule';

    /**
     * Returns true if the given object is an instance of Rule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Rule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Rule.__pulumiType;
    }

    /**
     * The ARN of the WAF Regional Rule.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The name or description for the Amazon CloudWatch metric of this rule.
     */
    public readonly metricName!: pulumi.Output<string>;
    /**
     * The name or description of the rule.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The objects to include in a rule (documented below).
     */
    public readonly predicates!: pulumi.Output<outputs.wafregional.RulePredicate[] | undefined>;
    /**
     * Key-value mapping of resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Rule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RuleArgs | RuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RuleState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["metricName"] = state ? state.metricName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["predicates"] = state ? state.predicates : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as RuleArgs | undefined;
            if (!args || args.metricName === undefined) {
                throw new Error("Missing required property 'metricName'");
            }
            inputs["metricName"] = args ? args.metricName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["predicates"] = args ? args.predicates : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Rule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Rule resources.
 */
export interface RuleState {
    /**
     * The ARN of the WAF Regional Rule.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The name or description for the Amazon CloudWatch metric of this rule.
     */
    readonly metricName?: pulumi.Input<string>;
    /**
     * The name or description of the rule.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The objects to include in a rule (documented below).
     */
    readonly predicates?: pulumi.Input<pulumi.Input<inputs.wafregional.RulePredicate>[]>;
    /**
     * Key-value mapping of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Rule resource.
 */
export interface RuleArgs {
    /**
     * The name or description for the Amazon CloudWatch metric of this rule.
     */
    readonly metricName: pulumi.Input<string>;
    /**
     * The name or description of the rule.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The objects to include in a rule (documented below).
     */
    readonly predicates?: pulumi.Input<pulumi.Input<inputs.wafregional.RulePredicate>[]>;
    /**
     * Key-value mapping of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
