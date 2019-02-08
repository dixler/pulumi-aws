// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an ECR repository policy.
 * 
 * Note that currently only one policy may be applied to a repository.
 * 
 * > **NOTE on ECR Availability**: The EC2 Container Registry is not yet rolled out
 * in all regions - available regions are listed
 * [the AWS Docs](https://docs.aws.amazon.com/general/latest/gr/rande.html#ecr_region).
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const foo = new aws.ecr.Repository("foo", {});
 * const foopolicy = new aws.ecr.RepositoryPolicy("foopolicy", {
 *     policy: `{
 *     "Version": "2008-10-17",
 *     "Statement": [
 *         {
 *             "Sid": "new policy",
 *             "Effect": "Allow",
 *             "Principal": "*",
 *             "Action": [
 *                 "ecr:GetDownloadUrlForLayer",
 *                 "ecr:BatchGetImage",
 *                 "ecr:BatchCheckLayerAvailability",
 *                 "ecr:PutImage",
 *                 "ecr:InitiateLayerUpload",
 *                 "ecr:UploadLayerPart",
 *                 "ecr:CompleteLayerUpload",
 *                 "ecr:DescribeRepositories",
 *                 "ecr:GetRepositoryPolicy",
 *                 "ecr:ListImages",
 *                 "ecr:DeleteRepository",
 *                 "ecr:BatchDeleteImage",
 *                 "ecr:SetRepositoryPolicy",
 *                 "ecr:DeleteRepositoryPolicy"
 *             ]
 *         }
 *     ]
 * }
 * `,
 *     repository: foo.name,
 * });
 * ```
 */
export class RepositoryPolicy extends pulumi.CustomResource {
    /**
     * Get an existing RepositoryPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryPolicyState, opts?: pulumi.CustomResourceOptions): RepositoryPolicy {
        return new RepositoryPolicy(name, <any>state, { ...opts, id: id });
    }

    /**
     * The policy document. This is a JSON formatted string. For more information about building IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html)
     */
    public readonly policy: pulumi.Output<string>;
    /**
     * The registry ID where the repository was created.
     */
    public /*out*/ readonly registryId: pulumi.Output<string>;
    /**
     * Name of the repository to apply the policy.
     */
    public readonly repository: pulumi.Output<string>;

    /**
     * Create a RepositoryPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RepositoryPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepositoryPolicyArgs | RepositoryPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: RepositoryPolicyState = argsOrState as RepositoryPolicyState | undefined;
            inputs["policy"] = state ? state.policy : undefined;
            inputs["registryId"] = state ? state.registryId : undefined;
            inputs["repository"] = state ? state.repository : undefined;
        } else {
            const args = argsOrState as RepositoryPolicyArgs | undefined;
            if (!args || args.policy === undefined) {
                throw new Error("Missing required property 'policy'");
            }
            if (!args || args.repository === undefined) {
                throw new Error("Missing required property 'repository'");
            }
            inputs["policy"] = args ? args.policy : undefined;
            inputs["repository"] = args ? args.repository : undefined;
            inputs["registryId"] = undefined /*out*/;
        }
        super("aws:ecr/repositoryPolicy:RepositoryPolicy", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RepositoryPolicy resources.
 */
export interface RepositoryPolicyState {
    /**
     * The policy document. This is a JSON formatted string. For more information about building IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html)
     */
    readonly policy?: pulumi.Input<string>;
    /**
     * The registry ID where the repository was created.
     */
    readonly registryId?: pulumi.Input<string>;
    /**
     * Name of the repository to apply the policy.
     */
    readonly repository?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RepositoryPolicy resource.
 */
export interface RepositoryPolicyArgs {
    /**
     * The policy document. This is a JSON formatted string. For more information about building IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html)
     */
    readonly policy: pulumi.Input<string>;
    /**
     * Name of the repository to apply the policy.
     */
    readonly repository: pulumi.Input<string>;
}
