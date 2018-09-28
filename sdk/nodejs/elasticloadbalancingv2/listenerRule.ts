// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Load Balancer Listener Rule resource.
 * 
 * ~> **Note:** `aws_alb_listener_rule` is known as `aws_lb_listener_rule`. The functionality is identical.
 */
export class ListenerRule extends pulumi.CustomResource {
    /**
     * Get an existing ListenerRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ListenerRuleState): ListenerRule {
        return new ListenerRule(name, <any>state, { id });
    }

    /**
     * An Action block. Action blocks are documented below.
     */
    public readonly actions: pulumi.Output<{ fixedResponse?: { contentType: string, messageBody?: string, statusCode: string }, redirect?: { host?: string, path?: string, port?: string, protocol?: string, query?: string, statusCode: string }, targetGroupArn?: string, type: string }[]>;
    /**
     * The ARN of the rule (matches `id`)
     */
    public /*out*/ readonly arn: pulumi.Output<string>;
    /**
     * A Condition block. Condition blocks are documented below.
     */
    public readonly conditions: pulumi.Output<{ field?: string, values?: string }[]>;
    /**
     * The ARN of the listener to which to attach the rule.
     */
    public readonly listenerArn: pulumi.Output<string>;
    /**
     * The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
     */
    public readonly priority: pulumi.Output<number>;

    /**
     * Create a ListenerRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ListenerRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ListenerRuleArgs | ListenerRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ListenerRuleState = argsOrState as ListenerRuleState | undefined;
            inputs["actions"] = state ? state.actions : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["conditions"] = state ? state.conditions : undefined;
            inputs["listenerArn"] = state ? state.listenerArn : undefined;
            inputs["priority"] = state ? state.priority : undefined;
        } else {
            const args = argsOrState as ListenerRuleArgs | undefined;
            if (!args || args.actions === undefined) {
                throw new Error("Missing required property 'actions'");
            }
            if (!args || args.conditions === undefined) {
                throw new Error("Missing required property 'conditions'");
            }
            if (!args || args.listenerArn === undefined) {
                throw new Error("Missing required property 'listenerArn'");
            }
            inputs["actions"] = args ? args.actions : undefined;
            inputs["conditions"] = args ? args.conditions : undefined;
            inputs["listenerArn"] = args ? args.listenerArn : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        super("aws:elasticloadbalancingv2/listenerRule:ListenerRule", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ListenerRule resources.
 */
export interface ListenerRuleState {
    /**
     * An Action block. Action blocks are documented below.
     */
    readonly actions?: pulumi.Input<pulumi.Input<{ fixedResponse?: pulumi.Input<{ contentType: pulumi.Input<string>, messageBody?: pulumi.Input<string>, statusCode?: pulumi.Input<string> }>, redirect?: pulumi.Input<{ host?: pulumi.Input<string>, path?: pulumi.Input<string>, port?: pulumi.Input<string>, protocol?: pulumi.Input<string>, query?: pulumi.Input<string>, statusCode: pulumi.Input<string> }>, targetGroupArn?: pulumi.Input<string>, type: pulumi.Input<string> }>[]>;
    /**
     * The ARN of the rule (matches `id`)
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * A Condition block. Condition blocks are documented below.
     */
    readonly conditions?: pulumi.Input<pulumi.Input<{ field?: pulumi.Input<string>, values?: pulumi.Input<string> }>[]>;
    /**
     * The ARN of the listener to which to attach the rule.
     */
    readonly listenerArn?: pulumi.Input<string>;
    /**
     * The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
     */
    readonly priority?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ListenerRule resource.
 */
export interface ListenerRuleArgs {
    /**
     * An Action block. Action blocks are documented below.
     */
    readonly actions: pulumi.Input<pulumi.Input<{ fixedResponse?: pulumi.Input<{ contentType: pulumi.Input<string>, messageBody?: pulumi.Input<string>, statusCode?: pulumi.Input<string> }>, redirect?: pulumi.Input<{ host?: pulumi.Input<string>, path?: pulumi.Input<string>, port?: pulumi.Input<string>, protocol?: pulumi.Input<string>, query?: pulumi.Input<string>, statusCode: pulumi.Input<string> }>, targetGroupArn?: pulumi.Input<string>, type: pulumi.Input<string> }>[]>;
    /**
     * A Condition block. Condition blocks are documented below.
     */
    readonly conditions: pulumi.Input<pulumi.Input<{ field?: pulumi.Input<string>, values?: pulumi.Input<string> }>[]>;
    /**
     * The ARN of the listener to which to attach the rule.
     */
    readonly listenerArn: pulumi.Input<string>;
    /**
     * The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
     */
    readonly priority?: pulumi.Input<number>;
}
