// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class NexusSecuritySaml extends pulumi.CustomResource {
    /**
     * Get an existing NexusSecuritySaml resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NexusSecuritySamlState, opts?: pulumi.CustomResourceOptions): NexusSecuritySaml {
        return new NexusSecuritySaml(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nexus:index/nexusSecuritySaml:NexusSecuritySaml';

    /**
     * Returns true if the given object is an instance of NexusSecuritySaml.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NexusSecuritySaml {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NexusSecuritySaml.__pulumiType;
    }

    /**
     * IdP field mappings for user's email address
     */
    public readonly emailAttribute!: pulumi.Output<string | undefined>;
    /**
     * Entity ID URI
     */
    public readonly entityId!: pulumi.Output<string | undefined>;
    /**
     * IdP field mappings for user's given name
     */
    public readonly firstNameAttribute!: pulumi.Output<string | undefined>;
    /**
     * IdP field mappings for user's groups
     */
    public readonly groupsAttribute!: pulumi.Output<string | undefined>;
    /**
     * SAML Identity Provider Metadata XML
     */
    public readonly idpMetadata!: pulumi.Output<string>;
    /**
     * IdP field mappings for user's family name
     */
    public readonly lastNameAttribute!: pulumi.Output<string | undefined>;
    /**
     * IdP field mappings for username
     */
    public readonly usernameAttribute!: pulumi.Output<string>;
    /**
     * By default, if a signing key is found in the IdP metadata, then NXRM will attempt to validate signatures on the
     * assertions.
     */
    public readonly validateAssertionSignature!: pulumi.Output<boolean | undefined>;
    /**
     * By default, if a signing key is found in the IdP metadata, then NXRM will attempt to validate signatures on the
     * response.
     */
    public readonly validateResponseSignature!: pulumi.Output<boolean | undefined>;

    /**
     * Create a NexusSecuritySaml resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NexusSecuritySamlArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NexusSecuritySamlArgs | NexusSecuritySamlState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NexusSecuritySamlState | undefined;
            resourceInputs["emailAttribute"] = state ? state.emailAttribute : undefined;
            resourceInputs["entityId"] = state ? state.entityId : undefined;
            resourceInputs["firstNameAttribute"] = state ? state.firstNameAttribute : undefined;
            resourceInputs["groupsAttribute"] = state ? state.groupsAttribute : undefined;
            resourceInputs["idpMetadata"] = state ? state.idpMetadata : undefined;
            resourceInputs["lastNameAttribute"] = state ? state.lastNameAttribute : undefined;
            resourceInputs["usernameAttribute"] = state ? state.usernameAttribute : undefined;
            resourceInputs["validateAssertionSignature"] = state ? state.validateAssertionSignature : undefined;
            resourceInputs["validateResponseSignature"] = state ? state.validateResponseSignature : undefined;
        } else {
            const args = argsOrState as NexusSecuritySamlArgs | undefined;
            if ((!args || args.idpMetadata === undefined) && !opts.urn) {
                throw new Error("Missing required property 'idpMetadata'");
            }
            if ((!args || args.usernameAttribute === undefined) && !opts.urn) {
                throw new Error("Missing required property 'usernameAttribute'");
            }
            resourceInputs["emailAttribute"] = args ? args.emailAttribute : undefined;
            resourceInputs["entityId"] = args ? args.entityId : undefined;
            resourceInputs["firstNameAttribute"] = args ? args.firstNameAttribute : undefined;
            resourceInputs["groupsAttribute"] = args ? args.groupsAttribute : undefined;
            resourceInputs["idpMetadata"] = args ? args.idpMetadata : undefined;
            resourceInputs["lastNameAttribute"] = args ? args.lastNameAttribute : undefined;
            resourceInputs["usernameAttribute"] = args ? args.usernameAttribute : undefined;
            resourceInputs["validateAssertionSignature"] = args ? args.validateAssertionSignature : undefined;
            resourceInputs["validateResponseSignature"] = args ? args.validateResponseSignature : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NexusSecuritySaml.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NexusSecuritySaml resources.
 */
export interface NexusSecuritySamlState {
    /**
     * IdP field mappings for user's email address
     */
    emailAttribute?: pulumi.Input<string>;
    /**
     * Entity ID URI
     */
    entityId?: pulumi.Input<string>;
    /**
     * IdP field mappings for user's given name
     */
    firstNameAttribute?: pulumi.Input<string>;
    /**
     * IdP field mappings for user's groups
     */
    groupsAttribute?: pulumi.Input<string>;
    /**
     * SAML Identity Provider Metadata XML
     */
    idpMetadata?: pulumi.Input<string>;
    /**
     * IdP field mappings for user's family name
     */
    lastNameAttribute?: pulumi.Input<string>;
    /**
     * IdP field mappings for username
     */
    usernameAttribute?: pulumi.Input<string>;
    /**
     * By default, if a signing key is found in the IdP metadata, then NXRM will attempt to validate signatures on the
     * assertions.
     */
    validateAssertionSignature?: pulumi.Input<boolean>;
    /**
     * By default, if a signing key is found in the IdP metadata, then NXRM will attempt to validate signatures on the
     * response.
     */
    validateResponseSignature?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a NexusSecuritySaml resource.
 */
export interface NexusSecuritySamlArgs {
    /**
     * IdP field mappings for user's email address
     */
    emailAttribute?: pulumi.Input<string>;
    /**
     * Entity ID URI
     */
    entityId?: pulumi.Input<string>;
    /**
     * IdP field mappings for user's given name
     */
    firstNameAttribute?: pulumi.Input<string>;
    /**
     * IdP field mappings for user's groups
     */
    groupsAttribute?: pulumi.Input<string>;
    /**
     * SAML Identity Provider Metadata XML
     */
    idpMetadata: pulumi.Input<string>;
    /**
     * IdP field mappings for user's family name
     */
    lastNameAttribute?: pulumi.Input<string>;
    /**
     * IdP field mappings for username
     */
    usernameAttribute: pulumi.Input<string>;
    /**
     * By default, if a signing key is found in the IdP metadata, then NXRM will attempt to validate signatures on the
     * assertions.
     */
    validateAssertionSignature?: pulumi.Input<boolean>;
    /**
     * By default, if a signing key is found in the IdP metadata, then NXRM will attempt to validate signatures on the
     * response.
     */
    validateResponseSignature?: pulumi.Input<boolean>;
}