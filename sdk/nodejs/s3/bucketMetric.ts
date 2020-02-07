// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a S3 bucket [metrics configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html) resource.
 * 
 * ## Example Usage
 * 
 * ### Add metrics configuration for entire S3 bucket
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.s3.Bucket("example", {});
 * const exampleEntireBucket = new aws.s3.BucketMetric("example-entire-bucket", {
 *     bucket: example.bucket,
 * });
 * ```
 * 
 * ### Add metrics configuration with S3 bucket object filter
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.s3.Bucket("example", {});
 * const exampleFiltered = new aws.s3.BucketMetric("example-filtered", {
 *     bucket: example.bucket,
 *     filter: {
 *         prefix: "documents/",
 *         tags: {
 *             class: "blue",
 *             priority: "high",
 *         },
 *     },
 * });
 * ```
 * 
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/s3_bucket_metric.html.markdown.
 */
export class BucketMetric extends pulumi.CustomResource {
    /**
     * Get an existing BucketMetric resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BucketMetricState, opts?: pulumi.CustomResourceOptions): BucketMetric {
        return new BucketMetric(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:s3/bucketMetric:BucketMetric';

    /**
     * Returns true if the given object is an instance of BucketMetric.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BucketMetric {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BucketMetric.__pulumiType;
    }

    /**
     * The name of the bucket to put metric configuration.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * [Object filtering](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html#metrics-configurations-filter) that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
     */
    public readonly filter!: pulumi.Output<outputs.s3.BucketMetricFilter | undefined>;
    /**
     * Unique identifier of the metrics configuration for the bucket.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a BucketMetric resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BucketMetricArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BucketMetricArgs | BucketMetricState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BucketMetricState | undefined;
            inputs["bucket"] = state ? state.bucket : undefined;
            inputs["filter"] = state ? state.filter : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as BucketMetricArgs | undefined;
            if (!args || args.bucket === undefined) {
                throw new Error("Missing required property 'bucket'");
            }
            inputs["bucket"] = args ? args.bucket : undefined;
            inputs["filter"] = args ? args.filter : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(BucketMetric.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BucketMetric resources.
 */
export interface BucketMetricState {
    /**
     * The name of the bucket to put metric configuration.
     */
    readonly bucket?: pulumi.Input<string>;
    /**
     * [Object filtering](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html#metrics-configurations-filter) that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
     */
    readonly filter?: pulumi.Input<inputs.s3.BucketMetricFilter>;
    /**
     * Unique identifier of the metrics configuration for the bucket.
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BucketMetric resource.
 */
export interface BucketMetricArgs {
    /**
     * The name of the bucket to put metric configuration.
     */
    readonly bucket: pulumi.Input<string>;
    /**
     * [Object filtering](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html#metrics-configurations-filter) that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
     */
    readonly filter?: pulumi.Input<inputs.s3.BucketMetricFilter>;
    /**
     * Unique identifier of the metrics configuration for the bucket.
     */
    readonly name?: pulumi.Input<string>;
}
