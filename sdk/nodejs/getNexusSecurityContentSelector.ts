// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getNexusSecurityContentSelector(args: GetNexusSecurityContentSelectorArgs, opts?: pulumi.InvokeOptions): Promise<GetNexusSecurityContentSelectorResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("nexus:index/getNexusSecurityContentSelector:GetNexusSecurityContentSelector", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking GetNexusSecurityContentSelector.
 */
export interface GetNexusSecurityContentSelectorArgs {
    name: string;
}

/**
 * A collection of values returned by GetNexusSecurityContentSelector.
 */
export interface GetNexusSecurityContentSelectorResult {
    readonly description: string;
    readonly expression: string;
    readonly id: string;
    readonly name: string;
}

export function getNexusSecurityContentSelectorOutput(args: GetNexusSecurityContentSelectorOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNexusSecurityContentSelectorResult> {
    return pulumi.output(args).apply(a => getNexusSecurityContentSelector(a, opts))
}

/**
 * A collection of arguments for invoking GetNexusSecurityContentSelector.
 */
export interface GetNexusSecurityContentSelectorOutputArgs {
    name: pulumi.Input<string>;
}