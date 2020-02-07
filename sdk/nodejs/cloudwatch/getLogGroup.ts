// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about an AWS Cloudwatch Log Group
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = aws.cloudwatch.getLogGroup({
 *     name: "MyImportantLogs",
 * });
 * ```
 * 
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/cloudwatch_log_group.html.markdown.
 */
export function getLogGroup(args: GetLogGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetLogGroupResult> & GetLogGroupResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetLogGroupResult> = pulumi.runtime.invoke("aws:cloudwatch/getLogGroup:getLogGroup", {
        "name": args.name,
        "tags": args.tags,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getLogGroup.
 */
export interface GetLogGroupArgs {
    /**
     * The name of the Cloudwatch log group
     */
    readonly name: string;
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getLogGroup.
 */
export interface GetLogGroupResult {
    /**
     * The ARN of the Cloudwatch log group
     */
    readonly arn: string;
    /**
     * The creation time of the log group, expressed as the number of milliseconds after Jan 1, 1970 00:00:00 UTC.
     */
    readonly creationTime: number;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The ARN of the KMS Key to use when encrypting log data.
     */
    readonly kmsKeyId: string;
    readonly name: string;
    /**
     * The number of days log events retained in the specified log group.
     */
    readonly retentionInDays: number;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags: {[key: string]: any};
}
