// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus.Inputs
{

    public sealed class NexusBlobstoreBucketConfigurationBucketArgs : Pulumi.ResourceArgs
    {
        [Input("expiration")]
        public Input<int>? Expiration { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public NexusBlobstoreBucketConfigurationBucketArgs()
        {
        }
    }
}
