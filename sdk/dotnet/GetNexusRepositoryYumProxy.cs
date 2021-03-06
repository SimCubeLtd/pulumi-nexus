// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus
{
    public static class GetNexusRepositoryYumProxy
    {
        public static Task<GetNexusRepositoryYumProxyResult> InvokeAsync(GetNexusRepositoryYumProxyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNexusRepositoryYumProxyResult>("nexus:index/getNexusRepositoryYumProxy:GetNexusRepositoryYumProxy", args ?? new GetNexusRepositoryYumProxyArgs(), options.WithDefaults());

        public static Output<GetNexusRepositoryYumProxyResult> Invoke(GetNexusRepositoryYumProxyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNexusRepositoryYumProxyResult>("nexus:index/getNexusRepositoryYumProxy:GetNexusRepositoryYumProxy", args ?? new GetNexusRepositoryYumProxyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNexusRepositoryYumProxyArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetNexusRepositoryYumProxyArgs()
        {
        }
    }

    public sealed class GetNexusRepositoryYumProxyInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetNexusRepositoryYumProxyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNexusRepositoryYumProxyResult
    {
        public readonly ImmutableArray<Outputs.GetNexusRepositoryYumProxyCleanupResult> Cleanups;
        public readonly ImmutableArray<Outputs.GetNexusRepositoryYumProxyHttpClientResult> HttpClients;
        public readonly string Id;
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetNexusRepositoryYumProxyNegativeCachResult> NegativeCaches;
        public readonly bool Online;
        public readonly ImmutableArray<Outputs.GetNexusRepositoryYumProxyProxyResult> Proxies;
        public readonly string RoutingRule;
        public readonly ImmutableArray<Outputs.GetNexusRepositoryYumProxyStorageResult> Storages;
        public readonly ImmutableArray<Outputs.GetNexusRepositoryYumProxyYumSigningResult> YumSignings;

        [OutputConstructor]
        private GetNexusRepositoryYumProxyResult(
            ImmutableArray<Outputs.GetNexusRepositoryYumProxyCleanupResult> cleanups,

            ImmutableArray<Outputs.GetNexusRepositoryYumProxyHttpClientResult> httpClients,

            string id,

            string name,

            ImmutableArray<Outputs.GetNexusRepositoryYumProxyNegativeCachResult> negativeCaches,

            bool online,

            ImmutableArray<Outputs.GetNexusRepositoryYumProxyProxyResult> proxies,

            string routingRule,

            ImmutableArray<Outputs.GetNexusRepositoryYumProxyStorageResult> storages,

            ImmutableArray<Outputs.GetNexusRepositoryYumProxyYumSigningResult> yumSignings)
        {
            Cleanups = cleanups;
            HttpClients = httpClients;
            Id = id;
            Name = name;
            NegativeCaches = negativeCaches;
            Online = online;
            Proxies = proxies;
            RoutingRule = routingRule;
            Storages = storages;
            YumSignings = yumSignings;
        }
    }
}
