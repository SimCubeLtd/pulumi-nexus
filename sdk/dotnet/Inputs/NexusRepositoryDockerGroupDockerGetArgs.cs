// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus.Inputs
{

    public sealed class NexusRepositoryDockerGroupDockerGetArgs : Pulumi.ResourceArgs
    {
        [Input("forceBasicAuth", required: true)]
        public Input<bool> ForceBasicAuth { get; set; } = null!;

        [Input("httpPort")]
        public Input<int>? HttpPort { get; set; }

        [Input("httpsPort")]
        public Input<int>? HttpsPort { get; set; }

        [Input("v1Enabled", required: true)]
        public Input<bool> V1Enabled { get; set; } = null!;

        public NexusRepositoryDockerGroupDockerGetArgs()
        {
        }
    }
}
