// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class NexusRepositoryYumHosted extends pulumi.CustomResource {
    /**
     * Get an existing NexusRepositoryYumHosted resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NexusRepositoryYumHostedState, opts?: pulumi.CustomResourceOptions): NexusRepositoryYumHosted {
        return new NexusRepositoryYumHosted(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nexus:index/nexusRepositoryYumHosted:NexusRepositoryYumHosted';

    /**
     * Returns true if the given object is an instance of NexusRepositoryYumHosted.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NexusRepositoryYumHosted {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NexusRepositoryYumHosted.__pulumiType;
    }

    /**
     * Cleanup policies
     */
    public readonly cleanups!: pulumi.Output<outputs.NexusRepositoryYumHostedCleanup[] | undefined>;
    /**
     * Component configuration for the hosted repository
     */
    public readonly component!: pulumi.Output<outputs.NexusRepositoryYumHostedComponent>;
    /**
     * Validate that all paths are RPMs or yum metadata. Possible values: `STRICT` or `PERMISSIVE`
     */
    public readonly deployPolicy!: pulumi.Output<string | undefined>;
    /**
     * A unique identifier for this repository
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Whether this repository accepts incoming requests
     */
    public readonly online!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the repository depth where repodata folder(s) are created. Possible values: 0-5
     */
    public readonly repodataDepth!: pulumi.Output<number | undefined>;
    /**
     * The storage configuration of the repository
     */
    public readonly storage!: pulumi.Output<outputs.NexusRepositoryYumHostedStorage>;

    /**
     * Create a NexusRepositoryYumHosted resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NexusRepositoryYumHostedArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NexusRepositoryYumHostedArgs | NexusRepositoryYumHostedState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NexusRepositoryYumHostedState | undefined;
            resourceInputs["cleanups"] = state ? state.cleanups : undefined;
            resourceInputs["component"] = state ? state.component : undefined;
            resourceInputs["deployPolicy"] = state ? state.deployPolicy : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["online"] = state ? state.online : undefined;
            resourceInputs["repodataDepth"] = state ? state.repodataDepth : undefined;
            resourceInputs["storage"] = state ? state.storage : undefined;
        } else {
            const args = argsOrState as NexusRepositoryYumHostedArgs | undefined;
            if ((!args || args.storage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storage'");
            }
            resourceInputs["cleanups"] = args ? args.cleanups : undefined;
            resourceInputs["component"] = args ? args.component : undefined;
            resourceInputs["deployPolicy"] = args ? args.deployPolicy : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["online"] = args ? args.online : undefined;
            resourceInputs["repodataDepth"] = args ? args.repodataDepth : undefined;
            resourceInputs["storage"] = args ? args.storage : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NexusRepositoryYumHosted.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NexusRepositoryYumHosted resources.
 */
export interface NexusRepositoryYumHostedState {
    /**
     * Cleanup policies
     */
    cleanups?: pulumi.Input<pulumi.Input<inputs.NexusRepositoryYumHostedCleanup>[]>;
    /**
     * Component configuration for the hosted repository
     */
    component?: pulumi.Input<inputs.NexusRepositoryYumHostedComponent>;
    /**
     * Validate that all paths are RPMs or yum metadata. Possible values: `STRICT` or `PERMISSIVE`
     */
    deployPolicy?: pulumi.Input<string>;
    /**
     * A unique identifier for this repository
     */
    name?: pulumi.Input<string>;
    /**
     * Whether this repository accepts incoming requests
     */
    online?: pulumi.Input<boolean>;
    /**
     * Specifies the repository depth where repodata folder(s) are created. Possible values: 0-5
     */
    repodataDepth?: pulumi.Input<number>;
    /**
     * The storage configuration of the repository
     */
    storage?: pulumi.Input<inputs.NexusRepositoryYumHostedStorage>;
}

/**
 * The set of arguments for constructing a NexusRepositoryYumHosted resource.
 */
export interface NexusRepositoryYumHostedArgs {
    /**
     * Cleanup policies
     */
    cleanups?: pulumi.Input<pulumi.Input<inputs.NexusRepositoryYumHostedCleanup>[]>;
    /**
     * Component configuration for the hosted repository
     */
    component?: pulumi.Input<inputs.NexusRepositoryYumHostedComponent>;
    /**
     * Validate that all paths are RPMs or yum metadata. Possible values: `STRICT` or `PERMISSIVE`
     */
    deployPolicy?: pulumi.Input<string>;
    /**
     * A unique identifier for this repository
     */
    name?: pulumi.Input<string>;
    /**
     * Whether this repository accepts incoming requests
     */
    online?: pulumi.Input<boolean>;
    /**
     * Specifies the repository depth where repodata folder(s) are created. Possible values: 0-5
     */
    repodataDepth?: pulumi.Input<number>;
    /**
     * The storage configuration of the repository
     */
    storage: pulumi.Input<inputs.NexusRepositoryYumHostedStorage>;
}
