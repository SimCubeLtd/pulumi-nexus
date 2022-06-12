// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus
{
    public static class GetNexusRoutingRule
    {
        public static Task<GetNexusRoutingRuleResult> InvokeAsync(GetNexusRoutingRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNexusRoutingRuleResult>("nexus:index/getNexusRoutingRule:GetNexusRoutingRule", args ?? new GetNexusRoutingRuleArgs(), options.WithDefaults());

        public static Output<GetNexusRoutingRuleResult> Invoke(GetNexusRoutingRuleInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNexusRoutingRuleResult>("nexus:index/getNexusRoutingRule:GetNexusRoutingRule", args ?? new GetNexusRoutingRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNexusRoutingRuleArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetNexusRoutingRuleArgs()
        {
        }
    }

    public sealed class GetNexusRoutingRuleInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetNexusRoutingRuleInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNexusRoutingRuleResult
    {
        public readonly string Description;
        public readonly string Id;
        public readonly ImmutableArray<string> Matchers;
        public readonly string Mode;
        public readonly string Name;

        [OutputConstructor]
        private GetNexusRoutingRuleResult(
            string description,

            string id,

            ImmutableArray<string> matchers,

            string mode,

            string name)
        {
            Description = description;
            Id = id;
            Matchers = matchers;
            Mode = mode;
            Name = name;
        }
    }
}
