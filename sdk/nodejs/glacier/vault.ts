// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Provides a Glacier Vault Resource. You can refer to the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/working-with-vaults.html) for a full explanation of the Glacier Vault functionality
 * 
 * ~> **NOTE:** When removing a Glacier Vault, the Vault must be empty.
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
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VaultState): Vault {
        return new Vault(name, <any>state, { id });
    }

    /**
     * The policy document. This is a JSON formatted string.
     * The heredoc syntax or `file` function is helpful here. Use the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html) for more information on Glacier Vault Policy
     */
    public readonly accessPolicy: pulumi.Output<string | undefined>;
    /**
     * The ARN of the vault.
     */
    public /*out*/ readonly arn: pulumi.Output<string>;
    /**
     * The URI of the vault that was created.
     */
    public /*out*/ readonly location: pulumi.Output<string>;
    /**
     * The name of the Vault. Names can be between 1 and 255 characters long and the valid characters are a-z, A-Z, 0-9, '_' (underscore), '-' (hyphen), and '.' (period).
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The notifications for the Vault. Fields documented below.
     */
    public readonly notifications: pulumi.Output<{ events: string[], snsTopic: string }[] | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Vault resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VaultArgs, opts?: pulumi.ResourceOptions)
    constructor(name: string, argsOrState?: VaultArgs | VaultState, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: VaultState = argsOrState as VaultState | undefined;
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
        super("aws:glacier/vault:Vault", name, inputs, opts);
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
    readonly notifications?: pulumi.Input<{ events: pulumi.Input<pulumi.Input<string>[]>, snsTopic: pulumi.Input<string> }[]>;
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
    readonly notifications?: pulumi.Input<{ events: pulumi.Input<pulumi.Input<string>[]>, snsTopic: pulumi.Input<string> }[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}