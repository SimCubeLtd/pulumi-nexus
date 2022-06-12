// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus.Inputs
{

    public sealed class NexusRepositoryYumProxyHttpClientGetArgs : Pulumi.ResourceArgs
    {
        [Input("authentication")]
        public Input<Inputs.NexusRepositoryYumProxyHttpClientAuthenticationGetArgs>? Authentication { get; set; }

        [Input("autoBlock")]
        public Input<bool>? AutoBlock { get; set; }

        [Input("blocked")]
        public Input<bool>? Blocked { get; set; }

        [Input("connection")]
        public Input<Inputs.NexusRepositoryYumProxyHttpClientConnectionGetArgs>? Connection { get; set; }

        public NexusRepositoryYumProxyHttpClientGetArgs()
        {
        }
    }
}
