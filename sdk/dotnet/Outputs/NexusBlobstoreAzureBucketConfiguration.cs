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
    public sealed class NexusBlobstoreAzureBucketConfiguration
    {
        public readonly string AccountName;
        public readonly Outputs.NexusBlobstoreAzureBucketConfigurationAuthentication Authentication;
        public readonly string ContainerName;

        [OutputConstructor]
        private NexusBlobstoreAzureBucketConfiguration(
            string accountName,

            Outputs.NexusBlobstoreAzureBucketConfigurationAuthentication authentication,

            string containerName)
        {
            AccountName = accountName;
            Authentication = authentication;
            ContainerName = containerName;
        }
    }
}