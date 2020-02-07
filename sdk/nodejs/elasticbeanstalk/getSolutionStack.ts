// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get the name of a elastic beanstalk solution stack.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const multiDocker = aws.elasticbeanstalk.getSolutionStack({
 *     mostRecent: true,
 *     nameRegex: "^64bit Amazon Linux (.*) Multi-container Docker (.*)$",
 * });
 * ```
 * 
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/elastic_beanstalk_solution_stack.html.markdown.
 */
export function getSolutionStack(args: GetSolutionStackArgs, opts?: pulumi.InvokeOptions): Promise<GetSolutionStackResult> & GetSolutionStackResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetSolutionStackResult> = pulumi.runtime.invoke("aws:elasticbeanstalk/getSolutionStack:getSolutionStack", {
        "mostRecent": args.mostRecent,
        "nameRegex": args.nameRegex,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getSolutionStack.
 */
export interface GetSolutionStackArgs {
    /**
     * If more than one result is returned, use the most
     * recent solution stack.
     */
    readonly mostRecent?: boolean;
    /**
     * A regex string to apply to the solution stack list returned
     * by AWS. See [Elastic Beanstalk Supported Platforms][beanstalk-platforms] from
     * AWS documentation for reference solution stack names.
     */
    readonly nameRegex: string;
}

/**
 * A collection of values returned by getSolutionStack.
 */
export interface GetSolutionStackResult {
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly mostRecent?: boolean;
    /**
     * The name of the solution stack.
     */
    readonly name: string;
    readonly nameRegex: string;
}
