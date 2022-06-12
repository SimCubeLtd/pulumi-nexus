// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getNexusRepositoryAptProxy(args: GetNexusRepositoryAptProxyArgs, opts?: pulumi.InvokeOptions): Promise<GetNexusRepositoryAptProxyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("nexus:index/getNexusRepositoryAptProxy:GetNexusRepositoryAptProxy", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking GetNexusRepositoryAptProxy.
 */
export interface GetNexusRepositoryAptProxyArgs {
    name: string;
}

/**
 * A collection of values returned by GetNexusRepositoryAptProxy.
 */
export interface GetNexusRepositoryAptProxyResult {
    readonly cleanups: outputs.GetNexusRepositoryAptProxyCleanup[];
    readonly distribution: string;
    readonly flat: boolean;
    readonly httpClients: outputs.GetNexusRepositoryAptProxyHttpClient[];
    readonly id: string;
    readonly name: string;
    readonly negativeCaches: outputs.GetNexusRepositoryAptProxyNegativeCach[];
    readonly online: boolean;
    readonly proxies: outputs.GetNexusRepositoryAptProxyProxy[];
    readonly routingRule: string;
    readonly storages: outputs.GetNexusRepositoryAptProxyStorage[];
}

export function getNexusRepositoryAptProxyOutput(args: GetNexusRepositoryAptProxyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNexusRepositoryAptProxyResult> {
    return pulumi.output(args).apply(a => getNexusRepositoryAptProxy(a, opts))
}

/**
 * A collection of arguments for invoking GetNexusRepositoryAptProxy.
 */
export interface GetNexusRepositoryAptProxyOutputArgs {
    name: pulumi.Input<string>;
}
