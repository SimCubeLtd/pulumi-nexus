// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus.Inputs
{

    public sealed class GetNexusBlobstoreBucketConfigurationAdvancedBucketConnectionInputArgs : Pulumi.ResourceArgs
    {
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        [Input("forcePathStyle")]
        public Input<bool>? ForcePathStyle { get; set; }

        [Input("signerType")]
        public Input<string>? SignerType { get; set; }

        public GetNexusBlobstoreBucketConfigurationAdvancedBucketConnectionInputArgs()
        {
        }
    }
}
