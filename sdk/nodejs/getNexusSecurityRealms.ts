// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getNexusSecurityRealms(opts?: pulumi.InvokeOptions): Promise<GetNexusSecurityRealmsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("nexus:index/getNexusSecurityRealms:GetNexusSecurityRealms", {
    }, opts);
}

/**
 * A collection of values returned by GetNexusSecurityRealms.
 */
export interface GetNexusSecurityRealmsResult {
    readonly actives: outputs.GetNexusSecurityRealmsActive[];
    readonly availables: outputs.GetNexusSecurityRealmsAvailable[];
    readonly id: string;
}
