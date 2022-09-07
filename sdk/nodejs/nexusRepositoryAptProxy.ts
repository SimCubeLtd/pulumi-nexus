// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class NexusRepositoryAptProxy extends pulumi.CustomResource {
    /**
     * Get an existing NexusRepositoryAptProxy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NexusRepositoryAptProxyState, opts?: pulumi.CustomResourceOptions): NexusRepositoryAptProxy {
        return new NexusRepositoryAptProxy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nexus:index/nexusRepositoryAptProxy:NexusRepositoryAptProxy';

    /**
     * Returns true if the given object is an instance of NexusRepositoryAptProxy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NexusRepositoryAptProxy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NexusRepositoryAptProxy.__pulumiType;
    }

    /**
     * Cleanup policies
     */
    public readonly cleanups!: pulumi.Output<outputs.NexusRepositoryAptProxyCleanup[] | undefined>;
    /**
     * Distribution to fetch
     */
    public readonly distribution!: pulumi.Output<string>;
    /**
     * Distribution to fetch
     */
    public readonly flat!: pulumi.Output<boolean>;
    /**
     * HTTP Client configuration for proxy repositories. Required for docker proxy repositories
     */
    public readonly httpClient!: pulumi.Output<outputs.NexusRepositoryAptProxyHttpClient | undefined>;
    /**
     * A unique identifier for this repository
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Configuration of the negative cache handling
     */
    public readonly negativeCache!: pulumi.Output<outputs.NexusRepositoryAptProxyNegativeCache | undefined>;
    /**
     * Whether this repository accepts incoming requests
     */
    public readonly online!: pulumi.Output<boolean | undefined>;
    /**
     * Configuration for the proxy repository
     */
    public readonly proxy!: pulumi.Output<outputs.NexusRepositoryAptProxyProxy>;
    /**
     * The name of the routing rule assigned to this repository
     */
    public readonly routingRule!: pulumi.Output<string | undefined>;
    /**
     * The storage configuration of the repository
     */
    public readonly storage!: pulumi.Output<outputs.NexusRepositoryAptProxyStorage>;

    /**
     * Create a NexusRepositoryAptProxy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NexusRepositoryAptProxyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NexusRepositoryAptProxyArgs | NexusRepositoryAptProxyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NexusRepositoryAptProxyState | undefined;
            resourceInputs["cleanups"] = state ? state.cleanups : undefined;
            resourceInputs["distribution"] = state ? state.distribution : undefined;
            resourceInputs["flat"] = state ? state.flat : undefined;
            resourceInputs["httpClient"] = state ? state.httpClient : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["negativeCache"] = state ? state.negativeCache : undefined;
            resourceInputs["online"] = state ? state.online : undefined;
            resourceInputs["proxy"] = state ? state.proxy : undefined;
            resourceInputs["routingRule"] = state ? state.routingRule : undefined;
            resourceInputs["storage"] = state ? state.storage : undefined;
        } else {
            const args = argsOrState as NexusRepositoryAptProxyArgs | undefined;
            if ((!args || args.distribution === undefined) && !opts.urn) {
                throw new Error("Missing required property 'distribution'");
            }
            if ((!args || args.flat === undefined) && !opts.urn) {
                throw new Error("Missing required property 'flat'");
            }
            if ((!args || args.proxy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'proxy'");
            }
            if ((!args || args.storage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storage'");
            }
            resourceInputs["cleanups"] = args ? args.cleanups : undefined;
            resourceInputs["distribution"] = args ? args.distribution : undefined;
            resourceInputs["flat"] = args ? args.flat : undefined;
            resourceInputs["httpClient"] = args ? args.httpClient : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["negativeCache"] = args ? args.negativeCache : undefined;
            resourceInputs["online"] = args ? args.online : undefined;
            resourceInputs["proxy"] = args ? args.proxy : undefined;
            resourceInputs["routingRule"] = args ? args.routingRule : undefined;
            resourceInputs["storage"] = args ? args.storage : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NexusRepositoryAptProxy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NexusRepositoryAptProxy resources.
 */
export interface NexusRepositoryAptProxyState {
    /**
     * Cleanup policies
     */
    cleanups?: pulumi.Input<pulumi.Input<inputs.NexusRepositoryAptProxyCleanup>[]>;
    /**
     * Distribution to fetch
     */
    distribution?: pulumi.Input<string>;
    /**
     * Distribution to fetch
     */
    flat?: pulumi.Input<boolean>;
    /**
     * HTTP Client configuration for proxy repositories. Required for docker proxy repositories
     */
    httpClient?: pulumi.Input<inputs.NexusRepositoryAptProxyHttpClient>;
    /**
     * A unique identifier for this repository
     */
    name?: pulumi.Input<string>;
    /**
     * Configuration of the negative cache handling
     */
    negativeCache?: pulumi.Input<inputs.NexusRepositoryAptProxyNegativeCache>;
    /**
     * Whether this repository accepts incoming requests
     */
    online?: pulumi.Input<boolean>;
    /**
     * Configuration for the proxy repository
     */
    proxy?: pulumi.Input<inputs.NexusRepositoryAptProxyProxy>;
    /**
     * The name of the routing rule assigned to this repository
     */
    routingRule?: pulumi.Input<string>;
    /**
     * The storage configuration of the repository
     */
    storage?: pulumi.Input<inputs.NexusRepositoryAptProxyStorage>;
}

/**
 * The set of arguments for constructing a NexusRepositoryAptProxy resource.
 */
export interface NexusRepositoryAptProxyArgs {
    /**
     * Cleanup policies
     */
    cleanups?: pulumi.Input<pulumi.Input<inputs.NexusRepositoryAptProxyCleanup>[]>;
    /**
     * Distribution to fetch
     */
    distribution: pulumi.Input<string>;
    /**
     * Distribution to fetch
     */
    flat: pulumi.Input<boolean>;
    /**
     * HTTP Client configuration for proxy repositories. Required for docker proxy repositories
     */
    httpClient?: pulumi.Input<inputs.NexusRepositoryAptProxyHttpClient>;
    /**
     * A unique identifier for this repository
     */
    name?: pulumi.Input<string>;
    /**
     * Configuration of the negative cache handling
     */
    negativeCache?: pulumi.Input<inputs.NexusRepositoryAptProxyNegativeCache>;
    /**
     * Whether this repository accepts incoming requests
     */
    online?: pulumi.Input<boolean>;
    /**
     * Configuration for the proxy repository
     */
    proxy: pulumi.Input<inputs.NexusRepositoryAptProxyProxy>;
    /**
     * The name of the routing rule assigned to this repository
     */
    routingRule?: pulumi.Input<string>;
    /**
     * The storage configuration of the repository
     */
    storage: pulumi.Input<inputs.NexusRepositoryAptProxyStorage>;
}