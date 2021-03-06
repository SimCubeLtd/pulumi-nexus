// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus.Inputs
{

    public sealed class NexusBlobstoreS3BucketConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("advancedBucketConnection")]
        public Input<Inputs.NexusBlobstoreS3BucketConfigurationAdvancedBucketConnectionArgs>? AdvancedBucketConnection { get; set; }

        [Input("bucket", required: true)]
        public Input<Inputs.NexusBlobstoreS3BucketConfigurationBucketArgs> Bucket { get; set; } = null!;

        [Input("bucketSecurity")]
        public Input<Inputs.NexusBlobstoreS3BucketConfigurationBucketSecurityArgs>? BucketSecurity { get; set; }

        [Input("encryption")]
        public Input<Inputs.NexusBlobstoreS3BucketConfigurationEncryptionArgs>? Encryption { get; set; }

        public NexusBlobstoreS3BucketConfigurationArgs()
        {
        }
    }
}
