// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getNexusUser(args: GetNexusUserArgs, opts?: pulumi.InvokeOptions): Promise<GetNexusUserResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("nexus:index/getNexusUser:GetNexusUser", {
        "userid": args.userid,
    }, opts);
}

/**
 * A collection of arguments for invoking GetNexusUser.
 */
export interface GetNexusUserArgs {
    userid: string;
}

/**
 * A collection of values returned by GetNexusUser.
 */
export interface GetNexusUserResult {
    readonly email: string;
    readonly firstname: string;
    readonly id: string;
    readonly lastname: string;
    readonly roles: string[];
    readonly status: string;
    readonly userid: string;
}

export function getNexusUserOutput(args: GetNexusUserOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNexusUserResult> {
    return pulumi.output(args).apply(a => getNexusUser(a, opts))
}

/**
 * A collection of arguments for invoking GetNexusUser.
 */
export interface GetNexusUserOutputArgs {
    userid: pulumi.Input<string>;
}
