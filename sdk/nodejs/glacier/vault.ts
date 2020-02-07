// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Glacier Vault Resource. You can refer to the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/working-with-vaults.html) for a full explanation of the Glacier Vault functionality
 * 
 * > **NOTE:** When removing a Glacier Vault, the Vault must be empty.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const awsSnsTopic = new aws.sns.Topic("aws.sns.Topic", {});
 * const myArchive = new aws.glacier.Vault("myArchive", {
 *     accessPolicy: `{
 *     "Version":"2012-10-17",
 *     "Statement":[
 *        {
 *           "Sid": "add-read-only-perm",
 *           "Principal": "*",
 *           "Effect": "Allow",
 *           "Action": [
 *              "glacier:InitiateJob",
 *              "glacier:GetJobOutput"
 *           ],
 *           "Resource": "arn:aws:glacier:eu-west-1:432981146916:vaults/MyArchive"
 *        }
 *     ]
 * }
 * `,
 *     notifications: [{
 *         events: [
 *             "ArchiveRetrievalCompleted",
 *             "InventoryRetrievalCompleted",
 *         ],
 *         snsTopic: awsSnsTopic.arn,
 *     }],
 *     tags: {
 *         Test: "MyArchive",
 *     },
 * });
 * ```
 * 
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/glacier_vault.html.markdown.
 */
export class Vault extends pulumi.CustomResource {
    /**
     * Get an existing Vault resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VaultState, opts?: pulumi.CustomResourceOptions): Vault {
        return new Vault(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:glacier/vault:Vault';

    /**
     * Returns true if the given object is an instance of Vault.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Vault {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Vault.__pulumiType;
    }

    /**
     * The policy document. This is a JSON formatted string.
     * The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
     */
    public readonly accessPolicy!: pulumi.Output<string | undefined>;
    /**
     * The ARN of the vault.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The URI of the vault that was created.
     */
    public /*out*/ readonly location!: pulumi.Output<string>;
    /**
     * The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The notifications for the Vault. Fields documented below.
     */
    public readonly notifications!: pulumi.Output<outputs.glacier.VaultNotification[] | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Vault resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VaultArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VaultArgs | VaultState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VaultState | undefined;
            inputs["accessPolicy"] = state ? state.accessPolicy : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["notifications"] = state ? state.notifications : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as VaultArgs | undefined;
            inputs["accessPolicy"] = args ? args.accessPolicy : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notifications"] = args ? args.notifications : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Vault.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Vault resources.
 */
export interface VaultState {
    /**
     * The policy document. This is a JSON formatted string.
     * The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
     */
    readonly accessPolicy?: pulumi.Input<string>;
    /**
     * The ARN of the vault.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The URI of the vault that was created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The notifications for the Vault. Fields documented below.
     */
    readonly notifications?: pulumi.Input<pulumi.Input<inputs.glacier.VaultNotification>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Vault resource.
 */
export interface VaultArgs {
    /**
     * The policy document. This is a JSON formatted string.
     * The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
     */
    readonly accessPolicy?: pulumi.Input<string>;
    /**
     * The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The notifications for the Vault. Fields documented below.
     */
    readonly notifications?: pulumi.Input<pulumi.Input<inputs.glacier.VaultNotification>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
