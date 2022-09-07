// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus
{
    public static class GetNexusBlobstoreS3
    {
        public static Task<GetNexusBlobstoreS3Result> InvokeAsync(GetNexusBlobstoreS3Args args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNexusBlobstoreS3Result>("nexus:index/getNexusBlobstoreS3:GetNexusBlobstoreS3", args ?? new GetNexusBlobstoreS3Args(), options.WithDefaults());

        public static Output<GetNexusBlobstoreS3Result> Invoke(GetNexusBlobstoreS3InvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNexusBlobstoreS3Result>("nexus:index/getNexusBlobstoreS3:GetNexusBlobstoreS3", args ?? new GetNexusBlobstoreS3InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNexusBlobstoreS3Args : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetNexusBlobstoreS3Args()
        {
        }
    }

    public sealed class GetNexusBlobstoreS3InvokeArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetNexusBlobstoreS3InvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNexusBlobstoreS3Result
    {
        public readonly int BlobCount;
        public readonly ImmutableArray<Outputs.GetNexusBlobstoreS3BucketConfigurationResult> BucketConfigurations;
        public readonly string Id;
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetNexusBlobstoreS3SoftQuotaResult> SoftQuotas;
        public readonly int TotalSizeInBytes;

        [OutputConstructor]
        private GetNexusBlobstoreS3Result(
            int blobCount,

            ImmutableArray<Outputs.GetNexusBlobstoreS3BucketConfigurationResult> bucketConfigurations,

            string id,

            string name,

            ImmutableArray<Outputs.GetNexusBlobstoreS3SoftQuotaResult> softQuotas,

            int totalSizeInBytes)
        {
            BlobCount = blobCount;
            BucketConfigurations = bucketConfigurations;
            Id = id;
            Name = name;
            SoftQuotas = softQuotas;
            TotalSizeInBytes = totalSizeInBytes;
        }
    }
}