// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a RDS Global Cluster, which is an Aurora global database spread across multiple regions. The global database contains a single primary cluster with read-write capability, and a read-only secondary cluster that receives data from the primary cluster through high-speed replication performed by the Aurora storage subsystem.
 * 
 * More information about Aurora global databases can be found in the [Aurora User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html#aurora-global-database-creating).
 * 
 * > **NOTE:** RDS only supports the `aurora` engine (MySQL 5.6 compatible) for Global Clusters at this time.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const primary = new aws.Provider("primary", {
 *     region: "us-east-2",
 * });
 * const secondary = new aws.Provider("secondary", {
 *     region: "us-west-2",
 * });
 * const example = new aws.rds.GlobalCluster("example", {
 *     globalClusterIdentifier: "example",
 * }, {provider: primary});
 * const primaryCluster = new aws.rds.Cluster("primary", {
 *     // ... other configuration ...
 *     engineMode: "global",
 *     globalClusterIdentifier: example.id,
 * }, {provider: primary});
 * const primaryClusterInstance = new aws.rds.ClusterInstance("primary", {
 *     // ... other configuration ...
 *     clusterIdentifier: primaryCluster.id,
 * }, {provider: primary});
 * const secondaryCluster = new aws.rds.Cluster("secondary", {
 *     // ... other configuration ...
 *     engineMode: "global",
 *     globalClusterIdentifier: example.id,
 * }, {provider: secondary,dependsOn: [primaryClusterInstance]});
 * const secondaryClusterInstance = new aws.rds.ClusterInstance("secondary", {
 *     // ... other configuration ...
 *     clusterIdentifier: secondaryCluster.id,
 * }, {provider: secondary});
 * ```
 * 
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/rds_global_cluster.html.markdown.
 */
export class GlobalCluster extends pulumi.CustomResource {
    /**
     * Get an existing GlobalCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GlobalClusterState, opts?: pulumi.CustomResourceOptions): GlobalCluster {
        return new GlobalCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:rds/globalCluster:GlobalCluster';

    /**
     * Returns true if the given object is an instance of GlobalCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GlobalCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GlobalCluster.__pulumiType;
    }

    /**
     * RDS Global Cluster Amazon Resource Name (ARN)
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Name for an automatically created database on cluster creation.
     */
    public readonly databaseName!: pulumi.Output<string | undefined>;
    /**
     * If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
     */
    public readonly deletionProtection!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the database engine to be used for this DB cluster. Valid values: `aurora`. Defaults to `aurora`.
     */
    public readonly engine!: pulumi.Output<string | undefined>;
    /**
     * Engine version of the Aurora global database.
     */
    public readonly engineVersion!: pulumi.Output<string>;
    /**
     * The global cluster identifier.
     */
    public readonly globalClusterIdentifier!: pulumi.Output<string>;
    /**
     * AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed
     */
    public /*out*/ readonly globalClusterResourceId!: pulumi.Output<string>;
    /**
     * Specifies whether the DB cluster is encrypted. The default is `false`.
     */
    public readonly storageEncrypted!: pulumi.Output<boolean | undefined>;

    /**
     * Create a GlobalCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GlobalClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GlobalClusterArgs | GlobalClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GlobalClusterState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["databaseName"] = state ? state.databaseName : undefined;
            inputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            inputs["engine"] = state ? state.engine : undefined;
            inputs["engineVersion"] = state ? state.engineVersion : undefined;
            inputs["globalClusterIdentifier"] = state ? state.globalClusterIdentifier : undefined;
            inputs["globalClusterResourceId"] = state ? state.globalClusterResourceId : undefined;
            inputs["storageEncrypted"] = state ? state.storageEncrypted : undefined;
        } else {
            const args = argsOrState as GlobalClusterArgs | undefined;
            if (!args || args.globalClusterIdentifier === undefined) {
                throw new Error("Missing required property 'globalClusterIdentifier'");
            }
            inputs["databaseName"] = args ? args.databaseName : undefined;
            inputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            inputs["engine"] = args ? args.engine : undefined;
            inputs["engineVersion"] = args ? args.engineVersion : undefined;
            inputs["globalClusterIdentifier"] = args ? args.globalClusterIdentifier : undefined;
            inputs["storageEncrypted"] = args ? args.storageEncrypted : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["globalClusterResourceId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GlobalCluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GlobalCluster resources.
 */
export interface GlobalClusterState {
    /**
     * RDS Global Cluster Amazon Resource Name (ARN)
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Name for an automatically created database on cluster creation.
     */
    readonly databaseName?: pulumi.Input<string>;
    /**
     * If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
     */
    readonly deletionProtection?: pulumi.Input<boolean>;
    /**
     * Name of the database engine to be used for this DB cluster. Valid values: `aurora`. Defaults to `aurora`.
     */
    readonly engine?: pulumi.Input<string>;
    /**
     * Engine version of the Aurora global database.
     */
    readonly engineVersion?: pulumi.Input<string>;
    /**
     * The global cluster identifier.
     */
    readonly globalClusterIdentifier?: pulumi.Input<string>;
    /**
     * AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed
     */
    readonly globalClusterResourceId?: pulumi.Input<string>;
    /**
     * Specifies whether the DB cluster is encrypted. The default is `false`.
     */
    readonly storageEncrypted?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a GlobalCluster resource.
 */
export interface GlobalClusterArgs {
    /**
     * Name for an automatically created database on cluster creation.
     */
    readonly databaseName?: pulumi.Input<string>;
    /**
     * If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
     */
    readonly deletionProtection?: pulumi.Input<boolean>;
    /**
     * Name of the database engine to be used for this DB cluster. Valid values: `aurora`. Defaults to `aurora`.
     */
    readonly engine?: pulumi.Input<string>;
    /**
     * Engine version of the Aurora global database.
     */
    readonly engineVersion?: pulumi.Input<string>;
    /**
     * The global cluster identifier.
     */
    readonly globalClusterIdentifier: pulumi.Input<string>;
    /**
     * Specifies whether the DB cluster is encrypted. The default is `false`.
     */
    readonly storageEncrypted?: pulumi.Input<boolean>;
}
