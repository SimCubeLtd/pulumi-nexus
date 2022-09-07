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
    public sealed class GetNexusBlobstoreS3BucketConfigurationBucketResult
    {
        public readonly int Expiration;
        public readonly string Name;
        public readonly string Prefix;
        public readonly string Region;

        [OutputConstructor]
        private GetNexusBlobstoreS3BucketConfigurationBucketResult(
            int expiration,

            string name,

            string prefix,

            string region)
        {
            Expiration = expiration;
            Name = name;
            Prefix = prefix;
            Region = region;
        }
    }
}