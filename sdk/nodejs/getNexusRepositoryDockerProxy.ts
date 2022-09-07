// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getNexusRepositoryDockerProxy(args: GetNexusRepositoryDockerProxyArgs, opts?: pulumi.InvokeOptions): Promise<GetNexusRepositoryDockerProxyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("nexus:index/getNexusRepositoryDockerProxy:GetNexusRepositoryDockerProxy", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking GetNexusRepositoryDockerProxy.
 */
export interface GetNexusRepositoryDockerProxyArgs {
    name: string;
}

/**
 * A collection of values returned by GetNexusRepositoryDockerProxy.
 */
export interface GetNexusRepositoryDockerProxyResult {
    readonly cleanups: outputs.GetNexusRepositoryDockerProxyCleanup[];
    readonly dockerProxies: outputs.GetNexusRepositoryDockerProxyDockerProxy[];
    readonly dockers: outputs.GetNexusRepositoryDockerProxyDocker[];
    readonly httpClients: outputs.GetNexusRepositoryDockerProxyHttpClient[];
    readonly id: string;
    readonly name: string;
    readonly negativeCaches: outputs.GetNexusRepositoryDockerProxyNegativeCach[];
    readonly online: boolean;
    readonly proxies: outputs.GetNexusRepositoryDockerProxyProxy[];
    readonly routingRule: string;
    readonly storages: outputs.GetNexusRepositoryDockerProxyStorage[];
}

export function getNexusRepositoryDockerProxyOutput(args: GetNexusRepositoryDockerProxyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNexusRepositoryDockerProxyResult> {
    return pulumi.output(args).apply(a => getNexusRepositoryDockerProxy(a, opts))
}

/**
 * A collection of arguments for invoking GetNexusRepositoryDockerProxy.
 */
export interface GetNexusRepositoryDockerProxyOutputArgs {
    name: pulumi.Input<string>;
}