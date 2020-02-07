// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about an EBS volume for use in other
 * resources.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const ebsVolume = aws.ebs.getVolume({
 *     filters: [
 *         {
 *             name: "volume-type",
 *             values: ["gp2"],
 *         },
 *         {
 *             name: "tag:Name",
 *             values: ["Example"],
 *         },
 *     ],
 *     mostRecent: true,
 * });
 * ```
 * 
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ebs_volume.html.markdown.
 */
export function getVolume(args?: GetVolumeArgs, opts?: pulumi.InvokeOptions): Promise<GetVolumeResult> & GetVolumeResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetVolumeResult> = pulumi.runtime.invoke("aws:ebs/getVolume:getVolume", {
        "filters": args.filters,
        "mostRecent": args.mostRecent,
        "tags": args.tags,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getVolume.
 */
export interface GetVolumeArgs {
    /**
     * One or more name/value pairs to filter off of. There are
     * several valid keys, for a full reference, check out
     * [describe-volumes in the AWS CLI reference][1].
     */
    readonly filters?: inputs.ebs.GetVolumeFilter[];
    /**
     * If more than one result is returned, use the most
     * recent Volume.
     */
    readonly mostRecent?: boolean;
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getVolume.
 */
export interface GetVolumeResult {
    /**
     * The volume ARN (e.g. arn:aws:ec2:us-east-1:0123456789012:volume/vol-59fcb34e).
     */
    readonly arn: string;
    /**
     * The AZ where the EBS volume exists.
     */
    readonly availabilityZone: string;
    /**
     * Whether the disk is encrypted.
     */
    readonly encrypted: boolean;
    readonly filters?: outputs.ebs.GetVolumeFilter[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The amount of IOPS for the disk.
     */
    readonly iops: number;
    /**
     * The ARN for the KMS encryption key.
     */
    readonly kmsKeyId: string;
    readonly mostRecent?: boolean;
    /**
     * The size of the drive in GiBs.
     */
    readonly size: number;
    /**
     * The snapshotId the EBS volume is based off.
     */
    readonly snapshotId: string;
    /**
     * A mapping of tags for the resource.
     */
    readonly tags: {[key: string]: any};
    /**
     * The volume ID (e.g. vol-59fcb34e).
     */
    readonly volumeId: string;
    /**
     * The type of EBS volume.
     */
    readonly volumeType: string;
}
