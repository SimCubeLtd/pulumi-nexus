// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getNexusRepositoryAptHosted(args: GetNexusRepositoryAptHostedArgs, opts?: pulumi.InvokeOptions): Promise<GetNexusRepositoryAptHostedResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("nexus:index/getNexusRepositoryAptHosted:GetNexusRepositoryAptHosted", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking GetNexusRepositoryAptHosted.
 */
export interface GetNexusRepositoryAptHostedArgs {
    name: string;
}

/**
 * A collection of values returned by GetNexusRepositoryAptHosted.
 */
export interface GetNexusRepositoryAptHostedResult {
    readonly cleanups: outputs.GetNexusRepositoryAptHostedCleanup[];
    readonly components: outputs.GetNexusRepositoryAptHostedComponent[];
    readonly distribution: string;
    readonly id: string;
    readonly name: string;
    readonly online: boolean;
    readonly storages: outputs.GetNexusRepositoryAptHostedStorage[];
}

export function getNexusRepositoryAptHostedOutput(args: GetNexusRepositoryAptHostedOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNexusRepositoryAptHostedResult> {
    return pulumi.output(args).apply(a => getNexusRepositoryAptHosted(a, opts))
}

/**
 * A collection of arguments for invoking GetNexusRepositoryAptHosted.
 */
export interface GetNexusRepositoryAptHostedOutputArgs {
    name: pulumi.Input<string>;
}
