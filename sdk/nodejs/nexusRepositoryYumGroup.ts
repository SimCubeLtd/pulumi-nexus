// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class NexusRepositoryYumGroup extends pulumi.CustomResource {
    /**
     * Get an existing NexusRepositoryYumGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NexusRepositoryYumGroupState, opts?: pulumi.CustomResourceOptions): NexusRepositoryYumGroup {
        return new NexusRepositoryYumGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nexus:index/nexusRepositoryYumGroup:NexusRepositoryYumGroup';

    /**
     * Returns true if the given object is an instance of NexusRepositoryYumGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NexusRepositoryYumGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NexusRepositoryYumGroup.__pulumiType;
    }

    /**
     * Configuration for repository group
     */
    public readonly group!: pulumi.Output<outputs.NexusRepositoryYumGroupGroup>;
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
    public readonly storage!: pulumi.Output<outputs.NexusRepositoryYumGroupStorage>;
    /**
     * Contains signing data of repositores
     */
    public readonly yumSigning!: pulumi.Output<outputs.NexusRepositoryYumGroupYumSigning | undefined>;

    /**
     * Create a NexusRepositoryYumGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NexusRepositoryYumGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NexusRepositoryYumGroupArgs | NexusRepositoryYumGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NexusRepositoryYumGroupState | undefined;
            resourceInputs["group"] = state ? state.group : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["online"] = state ? state.online : undefined;
            resourceInputs["storage"] = state ? state.storage : undefined;
            resourceInputs["yumSigning"] = state ? state.yumSigning : undefined;
        } else {
            const args = argsOrState as NexusRepositoryYumGroupArgs | undefined;
            if ((!args || args.group === undefined) && !opts.urn) {
                throw new Error("Missing required property 'group'");
            }
            if ((!args || args.storage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storage'");
            }
            resourceInputs["group"] = args ? args.group : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["online"] = args ? args.online : undefined;
            resourceInputs["storage"] = args ? args.storage : undefined;
            resourceInputs["yumSigning"] = args ? args.yumSigning : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NexusRepositoryYumGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NexusRepositoryYumGroup resources.
 */
export interface NexusRepositoryYumGroupState {
    /**
     * Configuration for repository group
     */
    group?: pulumi.Input<inputs.NexusRepositoryYumGroupGroup>;
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
    storage?: pulumi.Input<inputs.NexusRepositoryYumGroupStorage>;
    /**
     * Contains signing data of repositores
     */
    yumSigning?: pulumi.Input<inputs.NexusRepositoryYumGroupYumSigning>;
}

/**
 * The set of arguments for constructing a NexusRepositoryYumGroup resource.
 */
export interface NexusRepositoryYumGroupArgs {
    /**
     * Configuration for repository group
     */
    group: pulumi.Input<inputs.NexusRepositoryYumGroupGroup>;
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
    storage: pulumi.Input<inputs.NexusRepositoryYumGroupStorage>;
    /**
     * Contains signing data of repositores
     */
    yumSigning?: pulumi.Input<inputs.NexusRepositoryYumGroupYumSigning>;
}