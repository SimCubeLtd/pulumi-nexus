// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus
{
    public static class GetNexusBlobstoreGroup
    {
        public static Task<GetNexusBlobstoreGroupResult> InvokeAsync(GetNexusBlobstoreGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNexusBlobstoreGroupResult>("nexus:index/getNexusBlobstoreGroup:GetNexusBlobstoreGroup", args ?? new GetNexusBlobstoreGroupArgs(), options.WithDefaults());

        public static Output<GetNexusBlobstoreGroupResult> Invoke(GetNexusBlobstoreGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNexusBlobstoreGroupResult>("nexus:index/getNexusBlobstoreGroup:GetNexusBlobstoreGroup", args ?? new GetNexusBlobstoreGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNexusBlobstoreGroupArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetNexusBlobstoreGroupArgs()
        {
        }
    }

    public sealed class GetNexusBlobstoreGroupInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetNexusBlobstoreGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNexusBlobstoreGroupResult
    {
        public readonly int AvailableSpaceInBytes;
        public readonly int BlobCount;
        public readonly string FillPolicy;
        public readonly string Id;
        public readonly ImmutableArray<string> Members;
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetNexusBlobstoreGroupSoftQuotaResult> SoftQuotas;
        public readonly int TotalSizeInBytes;

        [OutputConstructor]
        private GetNexusBlobstoreGroupResult(
            int availableSpaceInBytes,

            int blobCount,

            string fillPolicy,

            string id,

            ImmutableArray<string> members,

            string name,

            ImmutableArray<Outputs.GetNexusBlobstoreGroupSoftQuotaResult> softQuotas,

            int totalSizeInBytes)
        {
            AvailableSpaceInBytes = availableSpaceInBytes;
            BlobCount = blobCount;
            FillPolicy = fillPolicy;
            Id = id;
            Members = members;
            Name = name;
            SoftQuotas = softQuotas;
            TotalSizeInBytes = totalSizeInBytes;
        }
    }
}
