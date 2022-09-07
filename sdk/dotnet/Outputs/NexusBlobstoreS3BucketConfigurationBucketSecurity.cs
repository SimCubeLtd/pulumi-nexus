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
    public sealed class NexusBlobstoreS3BucketConfigurationBucketSecurity
    {
        public readonly string? AccessKeyId;
        public readonly string? Role;
        public readonly string? SecretAccessKey;
        public readonly string? SessionToken;

        [OutputConstructor]
        private NexusBlobstoreS3BucketConfigurationBucketSecurity(
            string? accessKeyId,

            string? role,

            string? secretAccessKey,

            string? sessionToken)
        {
            AccessKeyId = accessKeyId;
            Role = role;
            SecretAccessKey = secretAccessKey;
            SessionToken = sessionToken;
        }
    }
}