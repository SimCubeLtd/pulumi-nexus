// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus.Inputs
{

    public sealed class NexusBlobstoreS3BucketConfigurationBucketSecurityArgs : Pulumi.ResourceArgs
    {
        [Input("accessKeyId")]
        public Input<string>? AccessKeyId { get; set; }

        [Input("role")]
        public Input<string>? Role { get; set; }

        [Input("secretAccessKey")]
        public Input<string>? SecretAccessKey { get; set; }

        [Input("sessionToken")]
        public Input<string>? SessionToken { get; set; }

        public NexusBlobstoreS3BucketConfigurationBucketSecurityArgs()
        {
        }
    }
}
