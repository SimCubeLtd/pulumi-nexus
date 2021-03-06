// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class NexusContentSelector extends pulumi.CustomResource {
    /**
     * Get an existing NexusContentSelector resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NexusContentSelectorState, opts?: pulumi.CustomResourceOptions): NexusContentSelector {
        return new NexusContentSelector(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nexus:index/nexusContentSelector:NexusContentSelector';

    /**
     * Returns true if the given object is an instance of NexusContentSelector.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NexusContentSelector {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NexusContentSelector.__pulumiType;
    }

    /**
     * A description of the content selector
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The content selector expression
     */
    public readonly expression!: pulumi.Output<string>;
    /**
     * Content selector name
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a NexusContentSelector resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NexusContentSelectorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NexusContentSelectorArgs | NexusContentSelectorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NexusContentSelectorState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["expression"] = state ? state.expression : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as NexusContentSelectorArgs | undefined;
            if ((!args || args.expression === undefined) && !opts.urn) {
                throw new Error("Missing required property 'expression'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["expression"] = args ? args.expression : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NexusContentSelector.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NexusContentSelector resources.
 */
export interface NexusContentSelectorState {
    /**
     * A description of the content selector
     */
    description?: pulumi.Input<string>;
    /**
     * The content selector expression
     */
    expression?: pulumi.Input<string>;
    /**
     * Content selector name
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NexusContentSelector resource.
 */
export interface NexusContentSelectorArgs {
    /**
     * A description of the content selector
     */
    description?: pulumi.Input<string>;
    /**
     * The content selector expression
     */
    expression: pulumi.Input<string>;
    /**
     * Content selector name
     */
    name?: pulumi.Input<string>;
}
