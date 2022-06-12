// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus
{
    public static class GetNexusRepositoryAptProxy
    {
        public static Task<GetNexusRepositoryAptProxyResult> InvokeAsync(GetNexusRepositoryAptProxyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNexusRepositoryAptProxyResult>("nexus:index/getNexusRepositoryAptProxy:GetNexusRepositoryAptProxy", args ?? new GetNexusRepositoryAptProxyArgs(), options.WithDefaults());

        public static Output<GetNexusRepositoryAptProxyResult> Invoke(GetNexusRepositoryAptProxyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNexusRepositoryAptProxyResult>("nexus:index/getNexusRepositoryAptProxy:GetNexusRepositoryAptProxy", args ?? new GetNexusRepositoryAptProxyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNexusRepositoryAptProxyArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetNexusRepositoryAptProxyArgs()
        {
        }
    }

    public sealed class GetNexusRepositoryAptProxyInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetNexusRepositoryAptProxyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNexusRepositoryAptProxyResult
    {
        public readonly ImmutableArray<Outputs.GetNexusRepositoryAptProxyCleanupResult> Cleanups;
        public readonly string Distribution;
        public readonly bool Flat;
        public readonly ImmutableArray<Outputs.GetNexusRepositoryAptProxyHttpClientResult> HttpClients;
        public readonly string Id;
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetNexusRepositoryAptProxyNegativeCachResult> NegativeCaches;
        public readonly bool Online;
        public readonly ImmutableArray<Outputs.GetNexusRepositoryAptProxyProxyResult> Proxies;
        public readonly string RoutingRule;
        public readonly ImmutableArray<Outputs.GetNexusRepositoryAptProxyStorageResult> Storages;

        [OutputConstructor]
        private GetNexusRepositoryAptProxyResult(
            ImmutableArray<Outputs.GetNexusRepositoryAptProxyCleanupResult> cleanups,

            string distribution,

            bool flat,

            ImmutableArray<Outputs.GetNexusRepositoryAptProxyHttpClientResult> httpClients,

            string id,

            string name,

            ImmutableArray<Outputs.GetNexusRepositoryAptProxyNegativeCachResult> negativeCaches,

            bool online,

            ImmutableArray<Outputs.GetNexusRepositoryAptProxyProxyResult> proxies,

            string routingRule,

            ImmutableArray<Outputs.GetNexusRepositoryAptProxyStorageResult> storages)
        {
            Cleanups = cleanups;
            Distribution = distribution;
            Flat = flat;
            HttpClients = httpClients;
            Id = id;
            Name = name;
            NegativeCaches = negativeCaches;
            Online = online;
            Proxies = proxies;
            RoutingRule = routingRule;
            Storages = storages;
        }
    }
}
