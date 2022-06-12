// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class NexusRepositoryMavenHosted extends pulumi.CustomResource {
    /**
     * Get an existing NexusRepositoryMavenHosted resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NexusRepositoryMavenHostedState, opts?: pulumi.CustomResourceOptions): NexusRepositoryMavenHosted {
        return new NexusRepositoryMavenHosted(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nexus:index/nexusRepositoryMavenHosted:NexusRepositoryMavenHosted';

    /**
     * Returns true if the given object is an instance of NexusRepositoryMavenHosted.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NexusRepositoryMavenHosted {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NexusRepositoryMavenHosted.__pulumiType;
    }

    /**
     * Cleanup policies
     */
    public readonly cleanups!: pulumi.Output<outputs.NexusRepositoryMavenHostedCleanup[] | undefined>;
    /**
     * Component configuration for the hosted repository
     */
    public readonly component!: pulumi.Output<outputs.NexusRepositoryMavenHostedComponent>;
    /**
     * Maven contains additional data of maven repository
     */
    public readonly maven!: pulumi.Output<outputs.NexusRepositoryMavenHostedMaven>;
    /**
     * A unique identifier for this repository
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Whether this repository accepts incoming requests
     */
    public readonly online!: pulumi.Output<boolean | undefined>;
    /**
     * The storage configuration of the repository
     */
    public readonly storage!: pulumi.Output<outputs.NexusRepositoryMavenHostedStorage>;

    /**
     * Create a NexusRepositoryMavenHosted resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NexusRepositoryMavenHostedArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NexusRepositoryMavenHostedArgs | NexusRepositoryMavenHostedState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NexusRepositoryMavenHostedState | undefined;
            resourceInputs["cleanups"] = state ? state.cleanups : undefined;
            resourceInputs["component"] = state ? state.component : undefined;
            resourceInputs["maven"] = state ? state.maven : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["online"] = state ? state.online : undefined;
            resourceInputs["storage"] = state ? state.storage : undefined;
        } else {
            const args = argsOrState as NexusRepositoryMavenHostedArgs | undefined;
            if ((!args || args.maven === undefined) && !opts.urn) {
                throw new Error("Missing required property 'maven'");
            }
            if ((!args || args.storage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storage'");
            }
            resourceInputs["cleanups"] = args ? args.cleanups : undefined;
            resourceInputs["component"] = args ? args.component : undefined;
            resourceInputs["maven"] = args ? args.maven : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["online"] = args ? args.online : undefined;
            resourceInputs["storage"] = args ? args.storage : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NexusRepositoryMavenHosted.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NexusRepositoryMavenHosted resources.
 */
export interface NexusRepositoryMavenHostedState {
    /**
     * Cleanup policies
     */
    cleanups?: pulumi.Input<pulumi.Input<inputs.NexusRepositoryMavenHostedCleanup>[]>;
    /**
     * Component configuration for the hosted repository
     */
    component?: pulumi.Input<inputs.NexusRepositoryMavenHostedComponent>;
    /**
     * Maven contains additional data of maven repository
     */
    maven?: pulumi.Input<inputs.NexusRepositoryMavenHostedMaven>;
    /**
     * A unique identifier for this repository
     */
    name?: pulumi.Input<string>;
    /**
     * Whether this repository accepts incoming requests
     */
    online?: pulumi.Input<boolean>;
    /**
     * The storage configuration of the repository
     */
    storage?: pulumi.Input<inputs.NexusRepositoryMavenHostedStorage>;
}

/**
 * The set of arguments for constructing a NexusRepositoryMavenHosted resource.
 */
export interface NexusRepositoryMavenHostedArgs {
    /**
     * Cleanup policies
     */
    cleanups?: pulumi.Input<pulumi.Input<inputs.NexusRepositoryMavenHostedCleanup>[]>;
    /**
     * Component configuration for the hosted repository
     */
    component?: pulumi.Input<inputs.NexusRepositoryMavenHostedComponent>;
    /**
     * Maven contains additional data of maven repository
     */
    maven: pulumi.Input<inputs.NexusRepositoryMavenHostedMaven>;
    /**
     * A unique identifier for this repository
     */
    name?: pulumi.Input<string>;
    /**
     * Whether this repository accepts incoming requests
     */
    online?: pulumi.Input<boolean>;
    /**
     * The storage configuration of the repository
     */
    storage: pulumi.Input<inputs.NexusRepositoryMavenHostedStorage>;
}
