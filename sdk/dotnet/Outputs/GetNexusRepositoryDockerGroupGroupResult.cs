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
    public sealed class GetNexusRepositoryDockerGroupGroupResult
    {
        public readonly ImmutableArray<string> MemberNames;
        public readonly string WritableMember;

        [OutputConstructor]
        private GetNexusRepositoryDockerGroupGroupResult(
            ImmutableArray<string> memberNames,

            string writableMember)
        {
            MemberNames = memberNames;
            WritableMember = writableMember;
        }
    }
}