// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getNexusRepository(args: GetNexusRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetNexusRepositoryResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("nexus:index/getNexusRepository:GetNexusRepository", {
        "aptSignings": args.aptSignings,
        "apts": args.apts,
        "cleanups": args.cleanups,
        "dockers": args.dockers,
        "format": args.format,
        "group": args.group,
        "httpClient": args.httpClient,
        "maven": args.maven,
        "name": args.name,
        "negativeCache": args.negativeCache,
        "online": args.online,
        "proxy": args.proxy,
        "storage": args.storage,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking GetNexusRepository.
 */
export interface GetNexusRepositoryArgs {
    aptSignings?: inputs.GetNexusRepositoryAptSigning[];
    apts?: inputs.GetNexusRepositoryApt[];
    cleanups?: inputs.GetNexusRepositoryCleanup[];
    dockers?: inputs.GetNexusRepositoryDocker[];
    format?: string;
    group?: inputs.GetNexusRepositoryGroup;
    httpClient?: inputs.GetNexusRepositoryHttpClient;
    maven?: inputs.GetNexusRepositoryMaven;
    name: string;
    negativeCache?: inputs.GetNexusRepositoryNegativeCache;
    online?: boolean;
    proxy?: inputs.GetNexusRepositoryProxy;
    storage?: inputs.GetNexusRepositoryStorage;
    type?: string;
}

/**
 * A collection of values returned by GetNexusRepository.
 */
export interface GetNexusRepositoryResult {
    readonly aptSignings?: outputs.GetNexusRepositoryAptSigning[];
    readonly apts?: outputs.GetNexusRepositoryApt[];
    readonly cleanups?: outputs.GetNexusRepositoryCleanup[];
    readonly dockers?: outputs.GetNexusRepositoryDocker[];
    readonly format?: string;
    readonly group?: outputs.GetNexusRepositoryGroup;
    readonly httpClient?: outputs.GetNexusRepositoryHttpClient;
    readonly id: string;
    readonly maven?: outputs.GetNexusRepositoryMaven;
    readonly name: string;
    readonly negativeCache?: outputs.GetNexusRepositoryNegativeCache;
    readonly online?: boolean;
    readonly proxy?: outputs.GetNexusRepositoryProxy;
    readonly storage?: outputs.GetNexusRepositoryStorage;
    readonly type?: string;
}

export function getNexusRepositoryOutput(args: GetNexusRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNexusRepositoryResult> {
    return pulumi.output(args).apply(a => getNexusRepository(a, opts))
}

/**
 * A collection of arguments for invoking GetNexusRepository.
 */
export interface GetNexusRepositoryOutputArgs {
    aptSignings?: pulumi.Input<pulumi.Input<inputs.GetNexusRepositoryAptSigningArgs>[]>;
    apts?: pulumi.Input<pulumi.Input<inputs.GetNexusRepositoryAptArgs>[]>;
    cleanups?: pulumi.Input<pulumi.Input<inputs.GetNexusRepositoryCleanupArgs>[]>;
    dockers?: pulumi.Input<pulumi.Input<inputs.GetNexusRepositoryDockerArgs>[]>;
    format?: pulumi.Input<string>;
    group?: pulumi.Input<inputs.GetNexusRepositoryGroupArgs>;
    httpClient?: pulumi.Input<inputs.GetNexusRepositoryHttpClientArgs>;
    maven?: pulumi.Input<inputs.GetNexusRepositoryMavenArgs>;
    name: pulumi.Input<string>;
    negativeCache?: pulumi.Input<inputs.GetNexusRepositoryNegativeCacheArgs>;
    online?: pulumi.Input<boolean>;
    proxy?: pulumi.Input<inputs.GetNexusRepositoryProxyArgs>;
    storage?: pulumi.Input<inputs.GetNexusRepositoryStorageArgs>;
    type?: pulumi.Input<string>;
}