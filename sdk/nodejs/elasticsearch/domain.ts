// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {Tags} from "../index";

export class Domain extends pulumi.CustomResource {
    /**
     * Get an existing Domain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainState): Domain {
        return new Domain(name, <any>state, { id });
    }

    /**
     * IAM policy document specifying the access policies for the domain
     */
    public readonly accessPolicies: pulumi.Output<string>;
    /**
     * Key-value string pairs to specify advanced configuration options.
     * Note that the values for these configuration options must be strings (wrapped in quotes) or they
     * may be wrong and cause a perpetual diff, causing Terraform to want to recreate your Elasticsearch
     * domain on every apply.
     */
    public readonly advancedOptions: pulumi.Output<{[key: string]: any}>;
    /**
     * Amazon Resource Name (ARN) of the domain.
     */
    public /*out*/ readonly arn: pulumi.Output<string>;
    /**
     * Cluster configuration of the domain, see below.
     */
    public readonly clusterConfig: pulumi.Output<{ dedicatedMasterCount?: number, dedicatedMasterEnabled?: boolean, dedicatedMasterType?: string, instanceCount?: number, instanceType?: string, zoneAwarenessEnabled?: boolean }>;
    public readonly cognitoOptions: pulumi.Output<{ enabled?: boolean, identityPoolId: string, roleArn: string, userPoolId: string } | undefined>;
    /**
     * Unique identifier for the domain.
     */
    public /*out*/ readonly domainId: pulumi.Output<string>;
    /**
     * Name of the domain.
     */
    public readonly domainName: pulumi.Output<string>;
    /**
     * EBS related options, may be required based on chosen [instance size](https://aws.amazon.com/elasticsearch-service/pricing/). See below.
     */
    public readonly ebsOptions: pulumi.Output<{ ebsEnabled: boolean, iops?: number, volumeSize?: number, volumeType: string }>;
    /**
     * The version of Elasticsearch to deploy. Defaults to `1.5`
     */
    public readonly elasticsearchVersion: pulumi.Output<string | undefined>;
    /**
     * Encrypt at rest options. Only available for [certain instance types](http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-supported-instance-types.html). See below.
     */
    public readonly encryptAtRest: pulumi.Output<{ enabled: boolean, kmsKeyId: string }>;
    /**
     * Domain-specific endpoint used to submit index, search, and data upload requests.
     */
    public /*out*/ readonly endpoint: pulumi.Output<string>;
    /**
     * Domain-specific endpoint for kibana without https scheme.
     * * `vpc_options.0.availability_zones` - If the domain was created inside a VPC, the names of the availability zones the configured `subnet_ids` were created inside.
     * * `vpc_options.0.vpc_id` - If the domain was created inside a VPC, the ID of the VPC.
     */
    public /*out*/ readonly kibanaEndpoint: pulumi.Output<string>;
    /**
     * Options for publishing slow logs to CloudWatch Logs.
     */
    public readonly logPublishingOptions: pulumi.Output<{ cloudwatchLogGroupArn: string, enabled?: boolean, logType: string }[] | undefined>;
    /**
     * Snapshot related options, see below.
     */
    public readonly snapshotOptions: pulumi.Output<{ automatedSnapshotStartHour: number } | undefined>;
    /**
     * A mapping of tags to assign to the resource
     */
    public readonly tags: pulumi.Output<Tags | undefined>;
    /**
     * VPC related options, see below. Adding or removing this configuration forces a new resource ([documentation](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-vpc-limitations)).
     */
    public readonly vpcOptions: pulumi.Output<{ availabilityZones: string[], securityGroupIds?: string[], subnetIds?: string[], vpcId: string } | undefined>;

    /**
     * Create a Domain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainArgs | DomainState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: DomainState = argsOrState as DomainState | undefined;
            inputs["accessPolicies"] = state ? state.accessPolicies : undefined;
            inputs["advancedOptions"] = state ? state.advancedOptions : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["clusterConfig"] = state ? state.clusterConfig : undefined;
            inputs["cognitoOptions"] = state ? state.cognitoOptions : undefined;
            inputs["domainId"] = state ? state.domainId : undefined;
            inputs["domainName"] = state ? state.domainName : undefined;
            inputs["ebsOptions"] = state ? state.ebsOptions : undefined;
            inputs["elasticsearchVersion"] = state ? state.elasticsearchVersion : undefined;
            inputs["encryptAtRest"] = state ? state.encryptAtRest : undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["kibanaEndpoint"] = state ? state.kibanaEndpoint : undefined;
            inputs["logPublishingOptions"] = state ? state.logPublishingOptions : undefined;
            inputs["snapshotOptions"] = state ? state.snapshotOptions : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vpcOptions"] = state ? state.vpcOptions : undefined;
        } else {
            const args = argsOrState as DomainArgs | undefined;
            inputs["accessPolicies"] = args ? args.accessPolicies : undefined;
            inputs["advancedOptions"] = args ? args.advancedOptions : undefined;
            inputs["clusterConfig"] = args ? args.clusterConfig : undefined;
            inputs["cognitoOptions"] = args ? args.cognitoOptions : undefined;
            inputs["domainName"] = args ? args.domainName : undefined;
            inputs["ebsOptions"] = args ? args.ebsOptions : undefined;
            inputs["elasticsearchVersion"] = args ? args.elasticsearchVersion : undefined;
            inputs["encryptAtRest"] = args ? args.encryptAtRest : undefined;
            inputs["logPublishingOptions"] = args ? args.logPublishingOptions : undefined;
            inputs["snapshotOptions"] = args ? args.snapshotOptions : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vpcOptions"] = args ? args.vpcOptions : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["domainId"] = undefined /*out*/;
            inputs["endpoint"] = undefined /*out*/;
            inputs["kibanaEndpoint"] = undefined /*out*/;
        }
        super("aws:elasticsearch/domain:Domain", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Domain resources.
 */
export interface DomainState {
    /**
     * IAM policy document specifying the access policies for the domain
     */
    readonly accessPolicies?: pulumi.Input<string>;
    /**
     * Key-value string pairs to specify advanced configuration options.
     * Note that the values for these configuration options must be strings (wrapped in quotes) or they
     * may be wrong and cause a perpetual diff, causing Terraform to want to recreate your Elasticsearch
     * domain on every apply.
     */
    readonly advancedOptions?: pulumi.Input<{[key: string]: any}>;
    /**
     * Amazon Resource Name (ARN) of the domain.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Cluster configuration of the domain, see below.
     */
    readonly clusterConfig?: pulumi.Input<{ dedicatedMasterCount?: pulumi.Input<number>, dedicatedMasterEnabled?: pulumi.Input<boolean>, dedicatedMasterType?: pulumi.Input<string>, instanceCount?: pulumi.Input<number>, instanceType?: pulumi.Input<string>, zoneAwarenessEnabled?: pulumi.Input<boolean> }>;
    readonly cognitoOptions?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, identityPoolId: pulumi.Input<string>, roleArn: pulumi.Input<string>, userPoolId: pulumi.Input<string> }>;
    /**
     * Unique identifier for the domain.
     */
    readonly domainId?: pulumi.Input<string>;
    /**
     * Name of the domain.
     */
    readonly domainName?: pulumi.Input<string>;
    /**
     * EBS related options, may be required based on chosen [instance size](https://aws.amazon.com/elasticsearch-service/pricing/). See below.
     */
    readonly ebsOptions?: pulumi.Input<{ ebsEnabled: pulumi.Input<boolean>, iops?: pulumi.Input<number>, volumeSize?: pulumi.Input<number>, volumeType?: pulumi.Input<string> }>;
    /**
     * The version of Elasticsearch to deploy. Defaults to `1.5`
     */
    readonly elasticsearchVersion?: pulumi.Input<string>;
    /**
     * Encrypt at rest options. Only available for [certain instance types](http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-supported-instance-types.html). See below.
     */
    readonly encryptAtRest?: pulumi.Input<{ enabled: pulumi.Input<boolean>, kmsKeyId?: pulumi.Input<string> }>;
    /**
     * Domain-specific endpoint used to submit index, search, and data upload requests.
     */
    readonly endpoint?: pulumi.Input<string>;
    /**
     * Domain-specific endpoint for kibana without https scheme.
     * * `vpc_options.0.availability_zones` - If the domain was created inside a VPC, the names of the availability zones the configured `subnet_ids` were created inside.
     * * `vpc_options.0.vpc_id` - If the domain was created inside a VPC, the ID of the VPC.
     */
    readonly kibanaEndpoint?: pulumi.Input<string>;
    /**
     * Options for publishing slow logs to CloudWatch Logs.
     */
    readonly logPublishingOptions?: pulumi.Input<pulumi.Input<{ cloudwatchLogGroupArn: pulumi.Input<string>, enabled?: pulumi.Input<boolean>, logType: pulumi.Input<string> }>[]>;
    /**
     * Snapshot related options, see below.
     */
    readonly snapshotOptions?: pulumi.Input<{ automatedSnapshotStartHour: pulumi.Input<number> }>;
    /**
     * A mapping of tags to assign to the resource
     */
    readonly tags?: pulumi.Input<Tags>;
    /**
     * VPC related options, see below. Adding or removing this configuration forces a new resource ([documentation](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-vpc-limitations)).
     */
    readonly vpcOptions?: pulumi.Input<{ availabilityZones?: pulumi.Input<pulumi.Input<string>[]>, securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>, subnetIds?: pulumi.Input<pulumi.Input<string>[]>, vpcId?: pulumi.Input<string> }>;
}

/**
 * The set of arguments for constructing a Domain resource.
 */
export interface DomainArgs {
    /**
     * IAM policy document specifying the access policies for the domain
     */
    readonly accessPolicies?: pulumi.Input<string>;
    /**
     * Key-value string pairs to specify advanced configuration options.
     * Note that the values for these configuration options must be strings (wrapped in quotes) or they
     * may be wrong and cause a perpetual diff, causing Terraform to want to recreate your Elasticsearch
     * domain on every apply.
     */
    readonly advancedOptions?: pulumi.Input<{[key: string]: any}>;
    /**
     * Cluster configuration of the domain, see below.
     */
    readonly clusterConfig?: pulumi.Input<{ dedicatedMasterCount?: pulumi.Input<number>, dedicatedMasterEnabled?: pulumi.Input<boolean>, dedicatedMasterType?: pulumi.Input<string>, instanceCount?: pulumi.Input<number>, instanceType?: pulumi.Input<string>, zoneAwarenessEnabled?: pulumi.Input<boolean> }>;
    readonly cognitoOptions?: pulumi.Input<{ enabled?: pulumi.Input<boolean>, identityPoolId: pulumi.Input<string>, roleArn: pulumi.Input<string>, userPoolId: pulumi.Input<string> }>;
    /**
     * Name of the domain.
     */
    readonly domainName?: pulumi.Input<string>;
    /**
     * EBS related options, may be required based on chosen [instance size](https://aws.amazon.com/elasticsearch-service/pricing/). See below.
     */
    readonly ebsOptions?: pulumi.Input<{ ebsEnabled: pulumi.Input<boolean>, iops?: pulumi.Input<number>, volumeSize?: pulumi.Input<number>, volumeType?: pulumi.Input<string> }>;
    /**
     * The version of Elasticsearch to deploy. Defaults to `1.5`
     */
    readonly elasticsearchVersion?: pulumi.Input<string>;
    /**
     * Encrypt at rest options. Only available for [certain instance types](http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-supported-instance-types.html). See below.
     */
    readonly encryptAtRest?: pulumi.Input<{ enabled: pulumi.Input<boolean>, kmsKeyId?: pulumi.Input<string> }>;
    /**
     * Options for publishing slow logs to CloudWatch Logs.
     */
    readonly logPublishingOptions?: pulumi.Input<pulumi.Input<{ cloudwatchLogGroupArn: pulumi.Input<string>, enabled?: pulumi.Input<boolean>, logType: pulumi.Input<string> }>[]>;
    /**
     * Snapshot related options, see below.
     */
    readonly snapshotOptions?: pulumi.Input<{ automatedSnapshotStartHour: pulumi.Input<number> }>;
    /**
     * A mapping of tags to assign to the resource
     */
    readonly tags?: pulumi.Input<Tags>;
    /**
     * VPC related options, see below. Adding or removing this configuration forces a new resource ([documentation](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-vpc-limitations)).
     */
    readonly vpcOptions?: pulumi.Input<{ availabilityZones?: pulumi.Input<pulumi.Input<string>[]>, securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>, subnetIds?: pulumi.Input<pulumi.Input<string>[]>, vpcId?: pulumi.Input<string> }>;
}
