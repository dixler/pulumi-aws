// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an Elastic File System (EFS) resource.
 * 
 * ## Example Usage
 * 
 * ### EFS File System w/ tags
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const foo = new aws.efs.FileSystem("foo", {
 *     tags: {
 *         Name: "MyProduct",
 *     },
 * });
 * ```
 * 
 * ### Using lifecycle policy
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const fooWithLifecylePolicy = new aws.efs.FileSystem("fooWithLifecylePolicy", {
 *     lifecyclePolicy: {
 *         transitionToIa: "AFTER_30_DAYS",
 *     },
 * });
 * ```
 * 
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/efs_file_system.html.markdown.
 */
export class FileSystem extends pulumi.CustomResource {
    /**
     * Get an existing FileSystem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FileSystemState, opts?: pulumi.CustomResourceOptions): FileSystem {
        return new FileSystem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:efs/fileSystem:FileSystem';

    /**
     * Returns true if the given object is an instance of FileSystem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FileSystem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FileSystem.__pulumiType;
    }

    /**
     * Amazon Resource Name of the file system.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A unique name (a maximum of 64 characters are allowed)
     * used as reference when creating the Elastic File System to ensure idempotent file
     * system creation. By default generated by this provider. See [Elastic File System]
     * (http://docs.aws.amazon.com/efs/latest/ug/) user guide for more information.
     */
    public readonly creationToken!: pulumi.Output<string>;
    /**
     * The DNS name for the filesystem per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
     */
    public /*out*/ readonly dnsName!: pulumi.Output<string>;
    /**
     * If true, the disk will be encrypted.
     */
    public readonly encrypted!: pulumi.Output<boolean>;
    /**
     * The ARN for the KMS encryption key. When specifying kms_key_id, encrypted needs to be set to true.
     */
    public readonly kmsKeyId!: pulumi.Output<string>;
    /**
     * A file system [lifecycle policy](https://docs.aws.amazon.com/efs/latest/ug/API_LifecyclePolicy.html) object (documented below).
     */
    public readonly lifecyclePolicy!: pulumi.Output<outputs.efs.FileSystemLifecyclePolicy | undefined>;
    /**
     * The file system performance mode. Can be either `"generalPurpose"` or `"maxIO"` (Default: `"generalPurpose"`).
     */
    public readonly performanceMode!: pulumi.Output<string>;
    /**
     * The throughput, measured in MiB/s, that you want to provision for the file system. Only applicable with `throughputMode` set to `provisioned`.
     */
    public readonly provisionedThroughputInMibps!: pulumi.Output<number | undefined>;
    /**
     * A mapping of tags to assign to the file system.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Throughput mode for the file system. Defaults to `bursting`. Valid values: `bursting`, `provisioned`. When using `provisioned`, also set `provisionedThroughputInMibps`.
     */
    public readonly throughputMode!: pulumi.Output<string | undefined>;

    /**
     * Create a FileSystem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FileSystemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FileSystemArgs | FileSystemState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FileSystemState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["creationToken"] = state ? state.creationToken : undefined;
            inputs["dnsName"] = state ? state.dnsName : undefined;
            inputs["encrypted"] = state ? state.encrypted : undefined;
            inputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            inputs["lifecyclePolicy"] = state ? state.lifecyclePolicy : undefined;
            inputs["performanceMode"] = state ? state.performanceMode : undefined;
            inputs["provisionedThroughputInMibps"] = state ? state.provisionedThroughputInMibps : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["throughputMode"] = state ? state.throughputMode : undefined;
        } else {
            const args = argsOrState as FileSystemArgs | undefined;
            inputs["creationToken"] = args ? args.creationToken : undefined;
            inputs["encrypted"] = args ? args.encrypted : undefined;
            inputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            inputs["lifecyclePolicy"] = args ? args.lifecyclePolicy : undefined;
            inputs["performanceMode"] = args ? args.performanceMode : undefined;
            inputs["provisionedThroughputInMibps"] = args ? args.provisionedThroughputInMibps : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["throughputMode"] = args ? args.throughputMode : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["dnsName"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(FileSystem.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FileSystem resources.
 */
export interface FileSystemState {
    /**
     * Amazon Resource Name of the file system.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * A unique name (a maximum of 64 characters are allowed)
     * used as reference when creating the Elastic File System to ensure idempotent file
     * system creation. By default generated by this provider. See [Elastic File System]
     * (http://docs.aws.amazon.com/efs/latest/ug/) user guide for more information.
     */
    readonly creationToken?: pulumi.Input<string>;
    /**
     * The DNS name for the filesystem per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
     */
    readonly dnsName?: pulumi.Input<string>;
    /**
     * If true, the disk will be encrypted.
     */
    readonly encrypted?: pulumi.Input<boolean>;
    /**
     * The ARN for the KMS encryption key. When specifying kms_key_id, encrypted needs to be set to true.
     */
    readonly kmsKeyId?: pulumi.Input<string>;
    /**
     * A file system [lifecycle policy](https://docs.aws.amazon.com/efs/latest/ug/API_LifecyclePolicy.html) object (documented below).
     */
    readonly lifecyclePolicy?: pulumi.Input<inputs.efs.FileSystemLifecyclePolicy>;
    /**
     * The file system performance mode. Can be either `"generalPurpose"` or `"maxIO"` (Default: `"generalPurpose"`).
     */
    readonly performanceMode?: pulumi.Input<string>;
    /**
     * The throughput, measured in MiB/s, that you want to provision for the file system. Only applicable with `throughputMode` set to `provisioned`.
     */
    readonly provisionedThroughputInMibps?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the file system.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Throughput mode for the file system. Defaults to `bursting`. Valid values: `bursting`, `provisioned`. When using `provisioned`, also set `provisionedThroughputInMibps`.
     */
    readonly throughputMode?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FileSystem resource.
 */
export interface FileSystemArgs {
    /**
     * A unique name (a maximum of 64 characters are allowed)
     * used as reference when creating the Elastic File System to ensure idempotent file
     * system creation. By default generated by this provider. See [Elastic File System]
     * (http://docs.aws.amazon.com/efs/latest/ug/) user guide for more information.
     */
    readonly creationToken?: pulumi.Input<string>;
    /**
     * If true, the disk will be encrypted.
     */
    readonly encrypted?: pulumi.Input<boolean>;
    /**
     * The ARN for the KMS encryption key. When specifying kms_key_id, encrypted needs to be set to true.
     */
    readonly kmsKeyId?: pulumi.Input<string>;
    /**
     * A file system [lifecycle policy](https://docs.aws.amazon.com/efs/latest/ug/API_LifecyclePolicy.html) object (documented below).
     */
    readonly lifecyclePolicy?: pulumi.Input<inputs.efs.FileSystemLifecyclePolicy>;
    /**
     * The file system performance mode. Can be either `"generalPurpose"` or `"maxIO"` (Default: `"generalPurpose"`).
     */
    readonly performanceMode?: pulumi.Input<string>;
    /**
     * The throughput, measured in MiB/s, that you want to provision for the file system. Only applicable with `throughputMode` set to `provisioned`.
     */
    readonly provisionedThroughputInMibps?: pulumi.Input<number>;
    /**
     * A mapping of tags to assign to the file system.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Throughput mode for the file system. Defaults to `bursting`. Valid values: `bursting`, `provisioned`. When using `provisioned`, also set `provisionedThroughputInMibps`.
     */
    readonly throughputMode?: pulumi.Input<string>;
}
