// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus.Outputs
{

    [OutputType]
    public sealed class NexusBlobstoreBucketConfiguration
    {
        public readonly Outputs.NexusBlobstoreBucketConfigurationAdvancedBucketConnection? AdvancedBucketConnection;
        public readonly Outputs.NexusBlobstoreBucketConfigurationBucket Bucket;
        public readonly Outputs.NexusBlobstoreBucketConfigurationBucketSecurity? BucketSecurity;
        public readonly Outputs.NexusBlobstoreBucketConfigurationEncryption? Encryption;

        [OutputConstructor]
        private NexusBlobstoreBucketConfiguration(
            Outputs.NexusBlobstoreBucketConfigurationAdvancedBucketConnection? advancedBucketConnection,

            Outputs.NexusBlobstoreBucketConfigurationBucket bucket,

            Outputs.NexusBlobstoreBucketConfigurationBucketSecurity? bucketSecurity,

            Outputs.NexusBlobstoreBucketConfigurationEncryption? encryption)
        {
            AdvancedBucketConnection = advancedBucketConnection;
            Bucket = bucket;
            BucketSecurity = bucketSecurity;
            Encryption = encryption;
        }
    }
}
