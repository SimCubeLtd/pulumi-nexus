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
    public sealed class GetNexusBlobstoreGroupSoftQuotaResult
    {
        public readonly int Limit;
        public readonly string Type;

        [OutputConstructor]
        private GetNexusBlobstoreGroupSoftQuotaResult(
            int limit,

            string type)
        {
            Limit = limit;
            Type = type;
        }
    }
}
