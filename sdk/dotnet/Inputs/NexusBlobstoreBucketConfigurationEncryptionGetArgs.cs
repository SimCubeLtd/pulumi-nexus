// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus.Inputs
{

    public sealed class NexusBlobstoreBucketConfigurationEncryptionGetArgs : Pulumi.ResourceArgs
    {
        [Input("encryptionKey")]
        public Input<string>? EncryptionKey { get; set; }

        [Input("encryptionType")]
        public Input<string>? EncryptionType { get; set; }

        public NexusBlobstoreBucketConfigurationEncryptionGetArgs()
        {
        }
    }
}
