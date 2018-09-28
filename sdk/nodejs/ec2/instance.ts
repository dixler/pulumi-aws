// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {InstanceProfile} from "../iam";
import {Tags} from "../index";
import {InstanceType} from "./instanceType";

/**
 * Provides an EC2 instance resource. This allows instances to be created, updated,
 * and deleted. Instances also support [provisioning](https://www.terraform.io/docs/provisioners/index.html).
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState): Instance {
        return new Instance(name, <any>state, { id });
    }

    /**
     * The AMI to use for the instance.
     */
    public readonly ami: pulumi.Output<string>;
    /**
     * The ARN of the instance.
     */
    public /*out*/ readonly arn: pulumi.Output<string>;
    /**
     * Associate a public ip address with an instance in a VPC.  Boolean value.
     */
    public readonly associatePublicIpAddress: pulumi.Output<boolean>;
    /**
     * The AZ to start the instance in.
     */
    public readonly availabilityZone: pulumi.Output<string>;
    /**
     * Sets the number of CPU cores for an instance. This option is 
     * only supported on creation of instance type that support CPU Options
     * [CPU Cores and Threads Per CPU Core Per Instance Type](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html#cpu-options-supported-instances-values) - specifying this option for unsupported instance types will return an error from the EC2 API.
     */
    public readonly cpuCoreCount: pulumi.Output<number>;
    /**
     * If set to to 1, hyperthreading is disabled on the launcehd instance. Defaults to 2 if not set. See [Optimizing CPU Options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html) for more information.
     */
    public readonly cpuThreadsPerCore: pulumi.Output<number>;
    /**
     * Customize the credit specification of the instance. See Credit Specification below for more details.
     */
    public readonly creditSpecification: pulumi.Output<{ cpuCredits?: string } | undefined>;
    /**
     * If true, enables [EC2 Instance
     * Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
     */
    public readonly disableApiTermination: pulumi.Output<boolean | undefined>;
    /**
     * Additional EBS block devices to attach to the
     * instance.  See Block Devices below for details.
     */
    public readonly ebsBlockDevices: pulumi.Output<{ deleteOnTermination?: boolean, deviceName: string, encrypted: boolean, iops: number, snapshotId: string, volumeId: string, volumeSize: number, volumeType: string }[]>;
    /**
     * If true, the launched EC2 instance will be EBS-optimized.
     * Note that if this is not set on an instance type that is optimized by default then
     * this will show as disabled but if the instance type is optimized by default then
     * there is no need to set this and there is no effect to disabling it.
     * See the [EBS Optimized section](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSOptimized.html) of the AWS User Guide for more information.
     */
    public readonly ebsOptimized: pulumi.Output<boolean | undefined>;
    /**
     * Customize Ephemeral (also known as
     * "Instance Store") volumes on the instance. See Block Devices below for details.
     */
    public readonly ephemeralBlockDevices: pulumi.Output<{ deviceName: string, noDevice?: boolean, virtualName?: string }[]>;
    /**
     * If true, wait for password data to become available and retrieve it. Useful for getting the administrator password for instances running Microsoft Windows. The password data is exported to the `password_data` attribute. See [GetPasswordData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetPasswordData.html) for more information.
     */
    public readonly getPasswordData: pulumi.Output<boolean | undefined>;
    /**
     * The IAM Instance Profile to
     * launch the instance with. Specified as the name of the Instance Profile. Ensure your credentials have the correct permission to assign the instance profile according to the [EC2 documentation](http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2.html#roles-usingrole-ec2instance-permissions), notably `iam:PassRole`.
     * * `ipv6_address_count`- (Optional) A number of IPv6 addresses to associate with the primary network interface. Amazon EC2 chooses the IPv6 addresses from the range of your subnet.
     */
    public readonly iamInstanceProfile: pulumi.Output<string | undefined>;
    /**
     * Shutdown behavior for the
     * instance. Amazon defaults this to `stop` for EBS-backed instances and
     * `terminate` for instance-store instances. Cannot be set on instance-store
     * instances. See [Shutdown Behavior](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingInstanceInitiatedShutdownBehavior) for more information.
     */
    public readonly instanceInitiatedShutdownBehavior: pulumi.Output<string | undefined>;
    public /*out*/ readonly instanceState: pulumi.Output<string>;
    /**
     * The type of instance to start. Updates to this field will trigger a stop/start of the EC2 instance.
     */
    public readonly instanceType: pulumi.Output<InstanceType>;
    public readonly ipv6AddressCount: pulumi.Output<number>;
    /**
     * Specify one or more IPv6 addresses from the range of the subnet to associate with the primary network interface
     */
    public readonly ipv6Addresses: pulumi.Output<string[]>;
    /**
     * The key name of the Key Pair to use for the instance; which can be managed using the `aws_key_pair` resource.
     */
    public readonly keyName: pulumi.Output<string>;
    /**
     * If true, the launched EC2 instance will have detailed monitoring enabled. (Available since v0.6.0)
     */
    public readonly monitoring: pulumi.Output<boolean | undefined>;
    /**
     * Customize network interfaces to be attached at instance boot time. See Network Interfaces below for more details.
     */
    public readonly networkInterfaces: pulumi.Output<{ deleteOnTermination?: boolean, deviceIndex: number, networkInterfaceId: string }[]>;
    /**
     * The ID of the network interface that was created with the instance.
     */
    public /*out*/ readonly networkInterfaceId: pulumi.Output<string>;
    /**
     * Base-64 encoded encrypted password data for the instance.
     * Useful for getting the administrator password for instances running Microsoft Windows.
     * This attribute is only exported if `get_password_data` is true.
     * Note that this encrypted value will be stored in the state file, as with all exported attributes.
     * See [GetPasswordData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetPasswordData.html) for more information.
     */
    public /*out*/ readonly passwordData: pulumi.Output<string>;
    /**
     * The Placement Group to start the instance in.
     */
    public readonly placementGroup: pulumi.Output<string>;
    /**
     * The ID of the instance's primary network interface.
     */
    public /*out*/ readonly primaryNetworkInterfaceId: pulumi.Output<string>;
    /**
     * The private DNS name assigned to the instance. Can only be
     * used inside the Amazon EC2, and only available if you've enabled DNS hostnames
     * for your VPC
     */
    public /*out*/ readonly privateDns: pulumi.Output<string>;
    /**
     * Private IP address to associate with the
     * instance in a VPC.
     */
    public readonly privateIp: pulumi.Output<string>;
    /**
     * The public DNS name assigned to the instance. For EC2-VPC, this
     * is only available if you've enabled DNS hostnames for your VPC
     */
    public /*out*/ readonly publicDns: pulumi.Output<string>;
    /**
     * The public IP address assigned to the instance, if applicable. **NOTE**: If you are using an [`aws_eip`](https://www.terraform.io/docs/providers/aws/r/eip.html) with your instance, you should refer to the EIP's address directly and not use `public_ip`, as this field will change after the EIP is attached.
     */
    public /*out*/ readonly publicIp: pulumi.Output<string>;
    /**
     * Customize details about the root block
     * device of the instance. See Block Devices below for details.
     */
    public readonly rootBlockDevice: pulumi.Output<{ deleteOnTermination?: boolean, iops: number, volumeId: string, volumeSize: number, volumeType: string }>;
    /**
     * A list of security group names (EC2-Classic) or IDs (default VPC) to associate with.
     */
    public readonly securityGroups: pulumi.Output<string[]>;
    /**
     * Controls if traffic is routed to the instance when
     * the destination address does not match the instance. Used for NAT or VPNs. Defaults true.
     */
    public readonly sourceDestCheck: pulumi.Output<boolean | undefined>;
    /**
     * The VPC Subnet ID to launch in.
     */
    public readonly subnetId: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags: pulumi.Output<Tags | undefined>;
    /**
     * The tenancy of the instance (if the instance is running in a VPC). An instance with a tenancy of dedicated runs on single-tenant hardware. The host tenancy is not supported for the import-instance command.
     */
    public readonly tenancy: pulumi.Output<string>;
    /**
     * The user data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see `user_data_base64` instead.
     */
    public readonly userData: pulumi.Output<string | undefined>;
    /**
     * Can be used instead of `user_data` to pass base64-encoded binary data directly. Use this instead of `user_data` whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption.
     */
    public readonly userDataBase64: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the devices created by the instance at launch time.
     */
    public readonly volumeTags: pulumi.Output<{[key: string]: any}>;
    /**
     * A list of security group IDs to associate with.
     */
    public readonly vpcSecurityGroupIds: pulumi.Output<string[]>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: InstanceState = argsOrState as InstanceState | undefined;
            inputs["ami"] = state ? state.ami : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["associatePublicIpAddress"] = state ? state.associatePublicIpAddress : undefined;
            inputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            inputs["cpuCoreCount"] = state ? state.cpuCoreCount : undefined;
            inputs["cpuThreadsPerCore"] = state ? state.cpuThreadsPerCore : undefined;
            inputs["creditSpecification"] = state ? state.creditSpecification : undefined;
            inputs["disableApiTermination"] = state ? state.disableApiTermination : undefined;
            inputs["ebsBlockDevices"] = state ? state.ebsBlockDevices : undefined;
            inputs["ebsOptimized"] = state ? state.ebsOptimized : undefined;
            inputs["ephemeralBlockDevices"] = state ? state.ephemeralBlockDevices : undefined;
            inputs["getPasswordData"] = state ? state.getPasswordData : undefined;
            inputs["iamInstanceProfile"] = state ? state.iamInstanceProfile : undefined;
            inputs["instanceInitiatedShutdownBehavior"] = state ? state.instanceInitiatedShutdownBehavior : undefined;
            inputs["instanceState"] = state ? state.instanceState : undefined;
            inputs["instanceType"] = state ? state.instanceType : undefined;
            inputs["ipv6AddressCount"] = state ? state.ipv6AddressCount : undefined;
            inputs["ipv6Addresses"] = state ? state.ipv6Addresses : undefined;
            inputs["keyName"] = state ? state.keyName : undefined;
            inputs["monitoring"] = state ? state.monitoring : undefined;
            inputs["networkInterfaces"] = state ? state.networkInterfaces : undefined;
            inputs["networkInterfaceId"] = state ? state.networkInterfaceId : undefined;
            inputs["passwordData"] = state ? state.passwordData : undefined;
            inputs["placementGroup"] = state ? state.placementGroup : undefined;
            inputs["primaryNetworkInterfaceId"] = state ? state.primaryNetworkInterfaceId : undefined;
            inputs["privateDns"] = state ? state.privateDns : undefined;
            inputs["privateIp"] = state ? state.privateIp : undefined;
            inputs["publicDns"] = state ? state.publicDns : undefined;
            inputs["publicIp"] = state ? state.publicIp : undefined;
            inputs["rootBlockDevice"] = state ? state.rootBlockDevice : undefined;
            inputs["securityGroups"] = state ? state.securityGroups : undefined;
            inputs["sourceDestCheck"] = state ? state.sourceDestCheck : undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tenancy"] = state ? state.tenancy : undefined;
            inputs["userData"] = state ? state.userData : undefined;
            inputs["userDataBase64"] = state ? state.userDataBase64 : undefined;
            inputs["volumeTags"] = state ? state.volumeTags : undefined;
            inputs["vpcSecurityGroupIds"] = state ? state.vpcSecurityGroupIds : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if (!args || args.ami === undefined) {
                throw new Error("Missing required property 'ami'");
            }
            if (!args || args.instanceType === undefined) {
                throw new Error("Missing required property 'instanceType'");
            }
            inputs["ami"] = args ? args.ami : undefined;
            inputs["associatePublicIpAddress"] = args ? args.associatePublicIpAddress : undefined;
            inputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            inputs["cpuCoreCount"] = args ? args.cpuCoreCount : undefined;
            inputs["cpuThreadsPerCore"] = args ? args.cpuThreadsPerCore : undefined;
            inputs["creditSpecification"] = args ? args.creditSpecification : undefined;
            inputs["disableApiTermination"] = args ? args.disableApiTermination : undefined;
            inputs["ebsBlockDevices"] = args ? args.ebsBlockDevices : undefined;
            inputs["ebsOptimized"] = args ? args.ebsOptimized : undefined;
            inputs["ephemeralBlockDevices"] = args ? args.ephemeralBlockDevices : undefined;
            inputs["getPasswordData"] = args ? args.getPasswordData : undefined;
            inputs["iamInstanceProfile"] = args ? args.iamInstanceProfile : undefined;
            inputs["instanceInitiatedShutdownBehavior"] = args ? args.instanceInitiatedShutdownBehavior : undefined;
            inputs["instanceType"] = args ? args.instanceType : undefined;
            inputs["ipv6AddressCount"] = args ? args.ipv6AddressCount : undefined;
            inputs["ipv6Addresses"] = args ? args.ipv6Addresses : undefined;
            inputs["keyName"] = args ? args.keyName : undefined;
            inputs["monitoring"] = args ? args.monitoring : undefined;
            inputs["networkInterfaces"] = args ? args.networkInterfaces : undefined;
            inputs["placementGroup"] = args ? args.placementGroup : undefined;
            inputs["privateIp"] = args ? args.privateIp : undefined;
            inputs["rootBlockDevice"] = args ? args.rootBlockDevice : undefined;
            inputs["securityGroups"] = args ? args.securityGroups : undefined;
            inputs["sourceDestCheck"] = args ? args.sourceDestCheck : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tenancy"] = args ? args.tenancy : undefined;
            inputs["userData"] = args ? args.userData : undefined;
            inputs["userDataBase64"] = args ? args.userDataBase64 : undefined;
            inputs["volumeTags"] = args ? args.volumeTags : undefined;
            inputs["vpcSecurityGroupIds"] = args ? args.vpcSecurityGroupIds : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["instanceState"] = undefined /*out*/;
            inputs["networkInterfaceId"] = undefined /*out*/;
            inputs["passwordData"] = undefined /*out*/;
            inputs["primaryNetworkInterfaceId"] = undefined /*out*/;
            inputs["privateDns"] = undefined /*out*/;
            inputs["publicDns"] = undefined /*out*/;
            inputs["publicIp"] = undefined /*out*/;
        }
        super("aws:ec2/instance:Instance", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * The AMI to use for the instance.
     */
    readonly ami?: pulumi.Input<string>;
    /**
     * The ARN of the instance.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Associate a public ip address with an instance in a VPC.  Boolean value.
     */
    readonly associatePublicIpAddress?: pulumi.Input<boolean>;
    /**
     * The AZ to start the instance in.
     */
    readonly availabilityZone?: pulumi.Input<string>;
    /**
     * Sets the number of CPU cores for an instance. This option is 
     * only supported on creation of instance type that support CPU Options
     * [CPU Cores and Threads Per CPU Core Per Instance Type](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html#cpu-options-supported-instances-values) - specifying this option for unsupported instance types will return an error from the EC2 API.
     */
    readonly cpuCoreCount?: pulumi.Input<number>;
    /**
     * If set to to 1, hyperthreading is disabled on the launcehd instance. Defaults to 2 if not set. See [Optimizing CPU Options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html) for more information.
     */
    readonly cpuThreadsPerCore?: pulumi.Input<number>;
    /**
     * Customize the credit specification of the instance. See Credit Specification below for more details.
     */
    readonly creditSpecification?: pulumi.Input<{ cpuCredits?: pulumi.Input<string> }>;
    /**
     * If true, enables [EC2 Instance
     * Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
     */
    readonly disableApiTermination?: pulumi.Input<boolean>;
    /**
     * Additional EBS block devices to attach to the
     * instance.  See Block Devices below for details.
     */
    readonly ebsBlockDevices?: pulumi.Input<pulumi.Input<{ deleteOnTermination?: pulumi.Input<boolean>, deviceName: pulumi.Input<string>, encrypted?: pulumi.Input<boolean>, iops?: pulumi.Input<number>, snapshotId?: pulumi.Input<string>, volumeId?: pulumi.Input<string>, volumeSize?: pulumi.Input<number>, volumeType?: pulumi.Input<string> }>[]>;
    /**
     * If true, the launched EC2 instance will be EBS-optimized.
     * Note that if this is not set on an instance type that is optimized by default then
     * this will show as disabled but if the instance type is optimized by default then
     * there is no need to set this and there is no effect to disabling it.
     * See the [EBS Optimized section](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSOptimized.html) of the AWS User Guide for more information.
     */
    readonly ebsOptimized?: pulumi.Input<boolean>;
    /**
     * Customize Ephemeral (also known as
     * "Instance Store") volumes on the instance. See Block Devices below for details.
     */
    readonly ephemeralBlockDevices?: pulumi.Input<pulumi.Input<{ deviceName: pulumi.Input<string>, noDevice?: pulumi.Input<boolean>, virtualName?: pulumi.Input<string> }>[]>;
    /**
     * If true, wait for password data to become available and retrieve it. Useful for getting the administrator password for instances running Microsoft Windows. The password data is exported to the `password_data` attribute. See [GetPasswordData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetPasswordData.html) for more information.
     */
    readonly getPasswordData?: pulumi.Input<boolean>;
    /**
     * The IAM Instance Profile to
     * launch the instance with. Specified as the name of the Instance Profile. Ensure your credentials have the correct permission to assign the instance profile according to the [EC2 documentation](http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2.html#roles-usingrole-ec2instance-permissions), notably `iam:PassRole`.
     * * `ipv6_address_count`- (Optional) A number of IPv6 addresses to associate with the primary network interface. Amazon EC2 chooses the IPv6 addresses from the range of your subnet.
     */
    readonly iamInstanceProfile?: pulumi.Input<string | InstanceProfile>;
    /**
     * Shutdown behavior for the
     * instance. Amazon defaults this to `stop` for EBS-backed instances and
     * `terminate` for instance-store instances. Cannot be set on instance-store
     * instances. See [Shutdown Behavior](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingInstanceInitiatedShutdownBehavior) for more information.
     */
    readonly instanceInitiatedShutdownBehavior?: pulumi.Input<string>;
    readonly instanceState?: pulumi.Input<string>;
    /**
     * The type of instance to start. Updates to this field will trigger a stop/start of the EC2 instance.
     */
    readonly instanceType?: pulumi.Input<InstanceType>;
    readonly ipv6AddressCount?: pulumi.Input<number>;
    /**
     * Specify one or more IPv6 addresses from the range of the subnet to associate with the primary network interface
     */
    readonly ipv6Addresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The key name of the Key Pair to use for the instance; which can be managed using the `aws_key_pair` resource.
     */
    readonly keyName?: pulumi.Input<string>;
    /**
     * If true, the launched EC2 instance will have detailed monitoring enabled. (Available since v0.6.0)
     */
    readonly monitoring?: pulumi.Input<boolean>;
    /**
     * Customize network interfaces to be attached at instance boot time. See Network Interfaces below for more details.
     */
    readonly networkInterfaces?: pulumi.Input<pulumi.Input<{ deleteOnTermination?: pulumi.Input<boolean>, deviceIndex: pulumi.Input<number>, networkInterfaceId: pulumi.Input<string> }>[]>;
    /**
     * The ID of the network interface that was created with the instance.
     */
    readonly networkInterfaceId?: pulumi.Input<string>;
    /**
     * Base-64 encoded encrypted password data for the instance.
     * Useful for getting the administrator password for instances running Microsoft Windows.
     * This attribute is only exported if `get_password_data` is true.
     * Note that this encrypted value will be stored in the state file, as with all exported attributes.
     * See [GetPasswordData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetPasswordData.html) for more information.
     */
    readonly passwordData?: pulumi.Input<string>;
    /**
     * The Placement Group to start the instance in.
     */
    readonly placementGroup?: pulumi.Input<string>;
    /**
     * The ID of the instance's primary network interface.
     */
    readonly primaryNetworkInterfaceId?: pulumi.Input<string>;
    /**
     * The private DNS name assigned to the instance. Can only be
     * used inside the Amazon EC2, and only available if you've enabled DNS hostnames
     * for your VPC
     */
    readonly privateDns?: pulumi.Input<string>;
    /**
     * Private IP address to associate with the
     * instance in a VPC.
     */
    readonly privateIp?: pulumi.Input<string>;
    /**
     * The public DNS name assigned to the instance. For EC2-VPC, this
     * is only available if you've enabled DNS hostnames for your VPC
     */
    readonly publicDns?: pulumi.Input<string>;
    /**
     * The public IP address assigned to the instance, if applicable. **NOTE**: If you are using an [`aws_eip`](https://www.terraform.io/docs/providers/aws/r/eip.html) with your instance, you should refer to the EIP's address directly and not use `public_ip`, as this field will change after the EIP is attached.
     */
    readonly publicIp?: pulumi.Input<string>;
    /**
     * Customize details about the root block
     * device of the instance. See Block Devices below for details.
     */
    readonly rootBlockDevice?: pulumi.Input<{ deleteOnTermination?: pulumi.Input<boolean>, iops?: pulumi.Input<number>, volumeId?: pulumi.Input<string>, volumeSize?: pulumi.Input<number>, volumeType?: pulumi.Input<string> }>;
    /**
     * A list of security group names (EC2-Classic) or IDs (default VPC) to associate with.
     */
    readonly securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Controls if traffic is routed to the instance when
     * the destination address does not match the instance. Used for NAT or VPNs. Defaults true.
     */
    readonly sourceDestCheck?: pulumi.Input<boolean>;
    /**
     * The VPC Subnet ID to launch in.
     */
    readonly subnetId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<Tags>;
    /**
     * The tenancy of the instance (if the instance is running in a VPC). An instance with a tenancy of dedicated runs on single-tenant hardware. The host tenancy is not supported for the import-instance command.
     */
    readonly tenancy?: pulumi.Input<string>;
    /**
     * The user data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see `user_data_base64` instead.
     */
    readonly userData?: pulumi.Input<string>;
    /**
     * Can be used instead of `user_data` to pass base64-encoded binary data directly. Use this instead of `user_data` whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption.
     */
    readonly userDataBase64?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the devices created by the instance at launch time.
     */
    readonly volumeTags?: pulumi.Input<{[key: string]: any}>;
    /**
     * A list of security group IDs to associate with.
     */
    readonly vpcSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * The AMI to use for the instance.
     */
    readonly ami: pulumi.Input<string>;
    /**
     * Associate a public ip address with an instance in a VPC.  Boolean value.
     */
    readonly associatePublicIpAddress?: pulumi.Input<boolean>;
    /**
     * The AZ to start the instance in.
     */
    readonly availabilityZone?: pulumi.Input<string>;
    /**
     * Sets the number of CPU cores for an instance. This option is 
     * only supported on creation of instance type that support CPU Options
     * [CPU Cores and Threads Per CPU Core Per Instance Type](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html#cpu-options-supported-instances-values) - specifying this option for unsupported instance types will return an error from the EC2 API.
     */
    readonly cpuCoreCount?: pulumi.Input<number>;
    /**
     * If set to to 1, hyperthreading is disabled on the launcehd instance. Defaults to 2 if not set. See [Optimizing CPU Options](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html) for more information.
     */
    readonly cpuThreadsPerCore?: pulumi.Input<number>;
    /**
     * Customize the credit specification of the instance. See Credit Specification below for more details.
     */
    readonly creditSpecification?: pulumi.Input<{ cpuCredits?: pulumi.Input<string> }>;
    /**
     * If true, enables [EC2 Instance
     * Termination Protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingDisableAPITermination)
     */
    readonly disableApiTermination?: pulumi.Input<boolean>;
    /**
     * Additional EBS block devices to attach to the
     * instance.  See Block Devices below for details.
     */
    readonly ebsBlockDevices?: pulumi.Input<pulumi.Input<{ deleteOnTermination?: pulumi.Input<boolean>, deviceName: pulumi.Input<string>, encrypted?: pulumi.Input<boolean>, iops?: pulumi.Input<number>, snapshotId?: pulumi.Input<string>, volumeId?: pulumi.Input<string>, volumeSize?: pulumi.Input<number>, volumeType?: pulumi.Input<string> }>[]>;
    /**
     * If true, the launched EC2 instance will be EBS-optimized.
     * Note that if this is not set on an instance type that is optimized by default then
     * this will show as disabled but if the instance type is optimized by default then
     * there is no need to set this and there is no effect to disabling it.
     * See the [EBS Optimized section](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSOptimized.html) of the AWS User Guide for more information.
     */
    readonly ebsOptimized?: pulumi.Input<boolean>;
    /**
     * Customize Ephemeral (also known as
     * "Instance Store") volumes on the instance. See Block Devices below for details.
     */
    readonly ephemeralBlockDevices?: pulumi.Input<pulumi.Input<{ deviceName: pulumi.Input<string>, noDevice?: pulumi.Input<boolean>, virtualName?: pulumi.Input<string> }>[]>;
    /**
     * If true, wait for password data to become available and retrieve it. Useful for getting the administrator password for instances running Microsoft Windows. The password data is exported to the `password_data` attribute. See [GetPasswordData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetPasswordData.html) for more information.
     */
    readonly getPasswordData?: pulumi.Input<boolean>;
    /**
     * The IAM Instance Profile to
     * launch the instance with. Specified as the name of the Instance Profile. Ensure your credentials have the correct permission to assign the instance profile according to the [EC2 documentation](http://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-ec2.html#roles-usingrole-ec2instance-permissions), notably `iam:PassRole`.
     * * `ipv6_address_count`- (Optional) A number of IPv6 addresses to associate with the primary network interface. Amazon EC2 chooses the IPv6 addresses from the range of your subnet.
     */
    readonly iamInstanceProfile?: pulumi.Input<string | InstanceProfile>;
    /**
     * Shutdown behavior for the
     * instance. Amazon defaults this to `stop` for EBS-backed instances and
     * `terminate` for instance-store instances. Cannot be set on instance-store
     * instances. See [Shutdown Behavior](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingInstanceInitiatedShutdownBehavior) for more information.
     */
    readonly instanceInitiatedShutdownBehavior?: pulumi.Input<string>;
    /**
     * The type of instance to start. Updates to this field will trigger a stop/start of the EC2 instance.
     */
    readonly instanceType: pulumi.Input<InstanceType>;
    readonly ipv6AddressCount?: pulumi.Input<number>;
    /**
     * Specify one or more IPv6 addresses from the range of the subnet to associate with the primary network interface
     */
    readonly ipv6Addresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The key name of the Key Pair to use for the instance; which can be managed using the `aws_key_pair` resource.
     */
    readonly keyName?: pulumi.Input<string>;
    /**
     * If true, the launched EC2 instance will have detailed monitoring enabled. (Available since v0.6.0)
     */
    readonly monitoring?: pulumi.Input<boolean>;
    /**
     * Customize network interfaces to be attached at instance boot time. See Network Interfaces below for more details.
     */
    readonly networkInterfaces?: pulumi.Input<pulumi.Input<{ deleteOnTermination?: pulumi.Input<boolean>, deviceIndex: pulumi.Input<number>, networkInterfaceId: pulumi.Input<string> }>[]>;
    /**
     * The Placement Group to start the instance in.
     */
    readonly placementGroup?: pulumi.Input<string>;
    /**
     * Private IP address to associate with the
     * instance in a VPC.
     */
    readonly privateIp?: pulumi.Input<string>;
    /**
     * Customize details about the root block
     * device of the instance. See Block Devices below for details.
     */
    readonly rootBlockDevice?: pulumi.Input<{ deleteOnTermination?: pulumi.Input<boolean>, iops?: pulumi.Input<number>, volumeId?: pulumi.Input<string>, volumeSize?: pulumi.Input<number>, volumeType?: pulumi.Input<string> }>;
    /**
     * A list of security group names (EC2-Classic) or IDs (default VPC) to associate with.
     */
    readonly securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Controls if traffic is routed to the instance when
     * the destination address does not match the instance. Used for NAT or VPNs. Defaults true.
     */
    readonly sourceDestCheck?: pulumi.Input<boolean>;
    /**
     * The VPC Subnet ID to launch in.
     */
    readonly subnetId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<Tags>;
    /**
     * The tenancy of the instance (if the instance is running in a VPC). An instance with a tenancy of dedicated runs on single-tenant hardware. The host tenancy is not supported for the import-instance command.
     */
    readonly tenancy?: pulumi.Input<string>;
    /**
     * The user data to provide when launching the instance. Do not pass gzip-compressed data via this argument; see `user_data_base64` instead.
     */
    readonly userData?: pulumi.Input<string>;
    /**
     * Can be used instead of `user_data` to pass base64-encoded binary data directly. Use this instead of `user_data` whenever the value is not a valid UTF-8 string. For example, gzip-encoded user data must be base64-encoded and passed via this argument to avoid corruption.
     */
    readonly userDataBase64?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the devices created by the instance at launch time.
     */
    readonly volumeTags?: pulumi.Input<{[key: string]: any}>;
    /**
     * A list of security group IDs to associate with.
     */
    readonly vpcSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
}
