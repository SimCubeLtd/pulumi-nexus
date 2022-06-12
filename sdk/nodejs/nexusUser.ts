// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class NexusUser extends pulumi.CustomResource {
    /**
     * Get an existing NexusUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NexusUserState, opts?: pulumi.CustomResourceOptions): NexusUser {
        return new NexusUser(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nexus:index/nexusUser:NexusUser';

    /**
     * Returns true if the given object is an instance of NexusUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NexusUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NexusUser.__pulumiType;
    }

    /**
     * The email address associated with the user.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * The first name of the user.
     */
    public readonly firstname!: pulumi.Output<string>;
    /**
     * The last name of the user.
     */
    public readonly lastname!: pulumi.Output<string>;
    /**
     * The password for the user.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * The roles which the user has been assigned within Nexus.
     */
    public readonly roles!: pulumi.Output<string[] | undefined>;
    /**
     * The user's status, e.g. active or disabled.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * The userid which is required for login. This value cannot be changed.
     */
    public readonly userid!: pulumi.Output<string>;

    /**
     * Create a NexusUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NexusUserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NexusUserArgs | NexusUserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NexusUserState | undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["firstname"] = state ? state.firstname : undefined;
            resourceInputs["lastname"] = state ? state.lastname : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["roles"] = state ? state.roles : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["userid"] = state ? state.userid : undefined;
        } else {
            const args = argsOrState as NexusUserArgs | undefined;
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            if ((!args || args.firstname === undefined) && !opts.urn) {
                throw new Error("Missing required property 'firstname'");
            }
            if ((!args || args.lastname === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lastname'");
            }
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            if ((!args || args.userid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userid'");
            }
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["firstname"] = args ? args.firstname : undefined;
            resourceInputs["lastname"] = args ? args.lastname : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["roles"] = args ? args.roles : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["userid"] = args ? args.userid : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NexusUser.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NexusUser resources.
 */
export interface NexusUserState {
    /**
     * The email address associated with the user.
     */
    email?: pulumi.Input<string>;
    /**
     * The first name of the user.
     */
    firstname?: pulumi.Input<string>;
    /**
     * The last name of the user.
     */
    lastname?: pulumi.Input<string>;
    /**
     * The password for the user.
     */
    password?: pulumi.Input<string>;
    /**
     * The roles which the user has been assigned within Nexus.
     */
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The user's status, e.g. active or disabled.
     */
    status?: pulumi.Input<string>;
    /**
     * The userid which is required for login. This value cannot be changed.
     */
    userid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NexusUser resource.
 */
export interface NexusUserArgs {
    /**
     * The email address associated with the user.
     */
    email: pulumi.Input<string>;
    /**
     * The first name of the user.
     */
    firstname: pulumi.Input<string>;
    /**
     * The last name of the user.
     */
    lastname: pulumi.Input<string>;
    /**
     * The password for the user.
     */
    password: pulumi.Input<string>;
    /**
     * The roles which the user has been assigned within Nexus.
     */
    roles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The user's status, e.g. active or disabled.
     */
    status?: pulumi.Input<string>;
    /**
     * The userid which is required for login. This value cannot be changed.
     */
    userid: pulumi.Input<string>;
}
