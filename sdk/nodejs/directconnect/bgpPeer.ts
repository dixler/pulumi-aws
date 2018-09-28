// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Direct Connect BGP peer resource.
 */
export class BgpPeer extends pulumi.CustomResource {
    /**
     * Get an existing BgpPeer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BgpPeerState): BgpPeer {
        return new BgpPeer(name, <any>state, { id });
    }

    /**
     * The address family for the BGP peer. `ipv4 ` or `ipv6`.
     */
    public readonly addressFamily: pulumi.Output<string>;
    /**
     * The IPv4 CIDR address to use to send traffic to Amazon.
     * Required for IPv4 BGP peers on public virtual interfaces.
     */
    public readonly amazonAddress: pulumi.Output<string>;
    /**
     * The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
     */
    public readonly bgpAsn: pulumi.Output<number>;
    /**
     * The authentication key for BGP configuration.
     */
    public readonly bgpAuthKey: pulumi.Output<string>;
    /**
     * The Up/Down state of the BGP peer.
     */
    public /*out*/ readonly bgpStatus: pulumi.Output<string>;
    /**
     * The IPv4 CIDR destination address to which Amazon should send traffic.
     * Required for IPv4 BGP peers on public virtual interfaces.
     */
    public readonly customerAddress: pulumi.Output<string>;
    /**
     * The ID of the Direct Connect virtual interface on which to create the BGP peer.
     */
    public readonly virtualInterfaceId: pulumi.Output<string>;

    /**
     * Create a BgpPeer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BgpPeerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BgpPeerArgs | BgpPeerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: BgpPeerState = argsOrState as BgpPeerState | undefined;
            inputs["addressFamily"] = state ? state.addressFamily : undefined;
            inputs["amazonAddress"] = state ? state.amazonAddress : undefined;
            inputs["bgpAsn"] = state ? state.bgpAsn : undefined;
            inputs["bgpAuthKey"] = state ? state.bgpAuthKey : undefined;
            inputs["bgpStatus"] = state ? state.bgpStatus : undefined;
            inputs["customerAddress"] = state ? state.customerAddress : undefined;
            inputs["virtualInterfaceId"] = state ? state.virtualInterfaceId : undefined;
        } else {
            const args = argsOrState as BgpPeerArgs | undefined;
            if (!args || args.addressFamily === undefined) {
                throw new Error("Missing required property 'addressFamily'");
            }
            if (!args || args.bgpAsn === undefined) {
                throw new Error("Missing required property 'bgpAsn'");
            }
            if (!args || args.virtualInterfaceId === undefined) {
                throw new Error("Missing required property 'virtualInterfaceId'");
            }
            inputs["addressFamily"] = args ? args.addressFamily : undefined;
            inputs["amazonAddress"] = args ? args.amazonAddress : undefined;
            inputs["bgpAsn"] = args ? args.bgpAsn : undefined;
            inputs["bgpAuthKey"] = args ? args.bgpAuthKey : undefined;
            inputs["customerAddress"] = args ? args.customerAddress : undefined;
            inputs["virtualInterfaceId"] = args ? args.virtualInterfaceId : undefined;
            inputs["bgpStatus"] = undefined /*out*/;
        }
        super("aws:directconnect/bgpPeer:BgpPeer", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BgpPeer resources.
 */
export interface BgpPeerState {
    /**
     * The address family for the BGP peer. `ipv4 ` or `ipv6`.
     */
    readonly addressFamily?: pulumi.Input<string>;
    /**
     * The IPv4 CIDR address to use to send traffic to Amazon.
     * Required for IPv4 BGP peers on public virtual interfaces.
     */
    readonly amazonAddress?: pulumi.Input<string>;
    /**
     * The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
     */
    readonly bgpAsn?: pulumi.Input<number>;
    /**
     * The authentication key for BGP configuration.
     */
    readonly bgpAuthKey?: pulumi.Input<string>;
    /**
     * The Up/Down state of the BGP peer.
     */
    readonly bgpStatus?: pulumi.Input<string>;
    /**
     * The IPv4 CIDR destination address to which Amazon should send traffic.
     * Required for IPv4 BGP peers on public virtual interfaces.
     */
    readonly customerAddress?: pulumi.Input<string>;
    /**
     * The ID of the Direct Connect virtual interface on which to create the BGP peer.
     */
    readonly virtualInterfaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BgpPeer resource.
 */
export interface BgpPeerArgs {
    /**
     * The address family for the BGP peer. `ipv4 ` or `ipv6`.
     */
    readonly addressFamily: pulumi.Input<string>;
    /**
     * The IPv4 CIDR address to use to send traffic to Amazon.
     * Required for IPv4 BGP peers on public virtual interfaces.
     */
    readonly amazonAddress?: pulumi.Input<string>;
    /**
     * The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
     */
    readonly bgpAsn: pulumi.Input<number>;
    /**
     * The authentication key for BGP configuration.
     */
    readonly bgpAuthKey?: pulumi.Input<string>;
    /**
     * The IPv4 CIDR destination address to which Amazon should send traffic.
     * Required for IPv4 BGP peers on public virtual interfaces.
     */
    readonly customerAddress?: pulumi.Input<string>;
    /**
     * The ID of the Direct Connect virtual interface on which to create the BGP peer.
     */
    readonly virtualInterfaceId: pulumi.Input<string>;
}
