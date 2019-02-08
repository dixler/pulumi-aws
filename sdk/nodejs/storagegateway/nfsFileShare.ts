// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an AWS Storage Gateway NFS File Share.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.storagegateway.NfsFileShare("example", {
 *     clientLists: ["0.0.0.0/0"],
 *     gatewayArn: aws_storagegateway_gateway_example.arn,
 *     locationArn: aws_s3_bucket_example.arn,
 *     roleArn: aws_iam_role_example.arn,
 * });
 * ```
 */
export class NfsFileShare extends pulumi.CustomResource {
    /**
     * Get an existing NfsFileShare resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NfsFileShareState, opts?: pulumi.CustomResourceOptions): NfsFileShare {
        return new NfsFileShare(name, <any>state, { ...opts, id: id });
    }

    /**
     * Amazon Resource Name (ARN) of the NFS File Share.
     */
    public /*out*/ readonly arn: pulumi.Output<string>;
    /**
     * The list of clients that are allowed to access the file gateway. The list must contain either valid IP addresses or valid CIDR blocks. Set to `["0.0.0.0/0"]` to not limit access. Minimum 1 item. Maximum 100 items.
     */
    public readonly clientLists: pulumi.Output<string[]>;
    /**
     * The default storage class for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`. Valid values: `S3_STANDARD`, `S3_STANDARD_IA`, `S3_ONEZONE_IA`.
     */
    public readonly defaultStorageClass: pulumi.Output<string | undefined>;
    /**
     * ID of the NFS File Share.
     */
    public /*out*/ readonly fileshareId: pulumi.Output<string>;
    /**
     * Amazon Resource Name (ARN) of the file gateway.
     */
    public readonly gatewayArn: pulumi.Output<string>;
    /**
     * Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
     */
    public readonly guessMimeTypeEnabled: pulumi.Output<boolean | undefined>;
    /**
     * Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
     */
    public readonly kmsEncrypted: pulumi.Output<boolean | undefined>;
    /**
     * Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kms_encrypted` is true.
     */
    public readonly kmsKeyArn: pulumi.Output<string | undefined>;
    /**
     * The ARN of the backed storage used for storing file data.
     */
    public readonly locationArn: pulumi.Output<string>;
    /**
     * Nested argument with file share default values. More information below.
     */
    public readonly nfsFileShareDefaults: pulumi.Output<{ directoryMode?: string, fileMode?: string, groupId?: number, ownerId?: number } | undefined>;
    /**
     * Access Control List permission for S3 bucket objects. Defaults to `private`.
     */
    public readonly objectAcl: pulumi.Output<string | undefined>;
    /**
     * Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
     */
    public readonly readOnly: pulumi.Output<boolean | undefined>;
    /**
     * Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
     */
    public readonly requesterPays: pulumi.Output<boolean | undefined>;
    /**
     * The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
     */
    public readonly roleArn: pulumi.Output<string>;
    /**
     * Maps a user to anonymous user. Defaults to `RootSquash`. Valid values: `RootSquash` (only root is mapped to anonymous user), `NoSquash` (no one is mapped to anonymous user), `AllSquash` (everyone is mapped to anonymous user)
     */
    public readonly squash: pulumi.Output<string | undefined>;

    /**
     * Create a NfsFileShare resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NfsFileShareArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NfsFileShareArgs | NfsFileShareState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: NfsFileShareState = argsOrState as NfsFileShareState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["clientLists"] = state ? state.clientLists : undefined;
            inputs["defaultStorageClass"] = state ? state.defaultStorageClass : undefined;
            inputs["fileshareId"] = state ? state.fileshareId : undefined;
            inputs["gatewayArn"] = state ? state.gatewayArn : undefined;
            inputs["guessMimeTypeEnabled"] = state ? state.guessMimeTypeEnabled : undefined;
            inputs["kmsEncrypted"] = state ? state.kmsEncrypted : undefined;
            inputs["kmsKeyArn"] = state ? state.kmsKeyArn : undefined;
            inputs["locationArn"] = state ? state.locationArn : undefined;
            inputs["nfsFileShareDefaults"] = state ? state.nfsFileShareDefaults : undefined;
            inputs["objectAcl"] = state ? state.objectAcl : undefined;
            inputs["readOnly"] = state ? state.readOnly : undefined;
            inputs["requesterPays"] = state ? state.requesterPays : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
            inputs["squash"] = state ? state.squash : undefined;
        } else {
            const args = argsOrState as NfsFileShareArgs | undefined;
            if (!args || args.clientLists === undefined) {
                throw new Error("Missing required property 'clientLists'");
            }
            if (!args || args.gatewayArn === undefined) {
                throw new Error("Missing required property 'gatewayArn'");
            }
            if (!args || args.locationArn === undefined) {
                throw new Error("Missing required property 'locationArn'");
            }
            if (!args || args.roleArn === undefined) {
                throw new Error("Missing required property 'roleArn'");
            }
            inputs["clientLists"] = args ? args.clientLists : undefined;
            inputs["defaultStorageClass"] = args ? args.defaultStorageClass : undefined;
            inputs["gatewayArn"] = args ? args.gatewayArn : undefined;
            inputs["guessMimeTypeEnabled"] = args ? args.guessMimeTypeEnabled : undefined;
            inputs["kmsEncrypted"] = args ? args.kmsEncrypted : undefined;
            inputs["kmsKeyArn"] = args ? args.kmsKeyArn : undefined;
            inputs["locationArn"] = args ? args.locationArn : undefined;
            inputs["nfsFileShareDefaults"] = args ? args.nfsFileShareDefaults : undefined;
            inputs["objectAcl"] = args ? args.objectAcl : undefined;
            inputs["readOnly"] = args ? args.readOnly : undefined;
            inputs["requesterPays"] = args ? args.requesterPays : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["squash"] = args ? args.squash : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["fileshareId"] = undefined /*out*/;
        }
        super("aws:storagegateway/nfsFileShare:NfsFileShare", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NfsFileShare resources.
 */
export interface NfsFileShareState {
    /**
     * Amazon Resource Name (ARN) of the NFS File Share.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The list of clients that are allowed to access the file gateway. The list must contain either valid IP addresses or valid CIDR blocks. Set to `["0.0.0.0/0"]` to not limit access. Minimum 1 item. Maximum 100 items.
     */
    readonly clientLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The default storage class for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`. Valid values: `S3_STANDARD`, `S3_STANDARD_IA`, `S3_ONEZONE_IA`.
     */
    readonly defaultStorageClass?: pulumi.Input<string>;
    /**
     * ID of the NFS File Share.
     */
    readonly fileshareId?: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) of the file gateway.
     */
    readonly gatewayArn?: pulumi.Input<string>;
    /**
     * Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
     */
    readonly guessMimeTypeEnabled?: pulumi.Input<boolean>;
    /**
     * Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
     */
    readonly kmsEncrypted?: pulumi.Input<boolean>;
    /**
     * Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kms_encrypted` is true.
     */
    readonly kmsKeyArn?: pulumi.Input<string>;
    /**
     * The ARN of the backed storage used for storing file data.
     */
    readonly locationArn?: pulumi.Input<string>;
    /**
     * Nested argument with file share default values. More information below.
     */
    readonly nfsFileShareDefaults?: pulumi.Input<{ directoryMode?: pulumi.Input<string>, fileMode?: pulumi.Input<string>, groupId?: pulumi.Input<number>, ownerId?: pulumi.Input<number> }>;
    /**
     * Access Control List permission for S3 bucket objects. Defaults to `private`.
     */
    readonly objectAcl?: pulumi.Input<string>;
    /**
     * Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
     */
    readonly readOnly?: pulumi.Input<boolean>;
    /**
     * Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
     */
    readonly requesterPays?: pulumi.Input<boolean>;
    /**
     * The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
     */
    readonly roleArn?: pulumi.Input<string>;
    /**
     * Maps a user to anonymous user. Defaults to `RootSquash`. Valid values: `RootSquash` (only root is mapped to anonymous user), `NoSquash` (no one is mapped to anonymous user), `AllSquash` (everyone is mapped to anonymous user)
     */
    readonly squash?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NfsFileShare resource.
 */
export interface NfsFileShareArgs {
    /**
     * The list of clients that are allowed to access the file gateway. The list must contain either valid IP addresses or valid CIDR blocks. Set to `["0.0.0.0/0"]` to not limit access. Minimum 1 item. Maximum 100 items.
     */
    readonly clientLists: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The default storage class for objects put into an Amazon S3 bucket by the file gateway. Defaults to `S3_STANDARD`. Valid values: `S3_STANDARD`, `S3_STANDARD_IA`, `S3_ONEZONE_IA`.
     */
    readonly defaultStorageClass?: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) of the file gateway.
     */
    readonly gatewayArn: pulumi.Input<string>;
    /**
     * Boolean value that enables guessing of the MIME type for uploaded objects based on file extensions. Defaults to `true`.
     */
    readonly guessMimeTypeEnabled?: pulumi.Input<boolean>;
    /**
     * Boolean value if `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Defaults to `false`.
     */
    readonly kmsEncrypted?: pulumi.Input<boolean>;
    /**
     * Amazon Resource Name (ARN) for KMS key used for Amazon S3 server side encryption. This value can only be set when `kms_encrypted` is true.
     */
    readonly kmsKeyArn?: pulumi.Input<string>;
    /**
     * The ARN of the backed storage used for storing file data.
     */
    readonly locationArn: pulumi.Input<string>;
    /**
     * Nested argument with file share default values. More information below.
     */
    readonly nfsFileShareDefaults?: pulumi.Input<{ directoryMode?: pulumi.Input<string>, fileMode?: pulumi.Input<string>, groupId?: pulumi.Input<number>, ownerId?: pulumi.Input<number> }>;
    /**
     * Access Control List permission for S3 bucket objects. Defaults to `private`.
     */
    readonly objectAcl?: pulumi.Input<string>;
    /**
     * Boolean to indicate write status of file share. File share does not accept writes if `true`. Defaults to `false`.
     */
    readonly readOnly?: pulumi.Input<boolean>;
    /**
     * Boolean who pays the cost of the request and the data download from the Amazon S3 bucket. Set this value to `true` if you want the requester to pay instead of the bucket owner. Defaults to `false`.
     */
    readonly requesterPays?: pulumi.Input<boolean>;
    /**
     * The ARN of the AWS Identity and Access Management (IAM) role that a file gateway assumes when it accesses the underlying storage.
     */
    readonly roleArn: pulumi.Input<string>;
    /**
     * Maps a user to anonymous user. Defaults to `RootSquash`. Valid values: `RootSquash` (only root is mapped to anonymous user), `NoSquash` (no one is mapped to anonymous user), `AllSquash` (everyone is mapped to anonymous user)
     */
    readonly squash?: pulumi.Input<string>;
}
