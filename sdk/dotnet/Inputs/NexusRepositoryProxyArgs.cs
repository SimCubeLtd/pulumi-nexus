// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus.Inputs
{

    public sealed class NexusRepositoryProxyArgs : Pulumi.ResourceArgs
    {
        [Input("contentMaxAge")]
        public Input<int>? ContentMaxAge { get; set; }

        [Input("metadataMaxAge")]
        public Input<int>? MetadataMaxAge { get; set; }

        [Input("remoteUrl")]
        public Input<string>? RemoteUrl { get; set; }

        public NexusRepositoryProxyArgs()
        {
        }
    }
}
