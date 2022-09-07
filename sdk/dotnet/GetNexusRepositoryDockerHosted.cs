// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus
{
    public static class GetNexusRepositoryDockerHosted
    {
        public static Task<GetNexusRepositoryDockerHostedResult> InvokeAsync(GetNexusRepositoryDockerHostedArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNexusRepositoryDockerHostedResult>("nexus:index/getNexusRepositoryDockerHosted:GetNexusRepositoryDockerHosted", args ?? new GetNexusRepositoryDockerHostedArgs(), options.WithDefaults());

        public static Output<GetNexusRepositoryDockerHostedResult> Invoke(GetNexusRepositoryDockerHostedInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNexusRepositoryDockerHostedResult>("nexus:index/getNexusRepositoryDockerHosted:GetNexusRepositoryDockerHosted", args ?? new GetNexusRepositoryDockerHostedInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNexusRepositoryDockerHostedArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetNexusRepositoryDockerHostedArgs()
        {
        }
    }

    public sealed class GetNexusRepositoryDockerHostedInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetNexusRepositoryDockerHostedInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNexusRepositoryDockerHostedResult
    {
        public readonly ImmutableArray<Outputs.GetNexusRepositoryDockerHostedCleanupResult> Cleanups;
        public readonly ImmutableArray<Outputs.GetNexusRepositoryDockerHostedComponentResult> Components;
        public readonly ImmutableArray<Outputs.GetNexusRepositoryDockerHostedDockerResult> Dockers;
        public readonly string Id;
        public readonly string Name;
        public readonly bool Online;
        public readonly ImmutableArray<Outputs.GetNexusRepositoryDockerHostedStorageResult> Storages;

        [OutputConstructor]
        private GetNexusRepositoryDockerHostedResult(
            ImmutableArray<Outputs.GetNexusRepositoryDockerHostedCleanupResult> cleanups,

            ImmutableArray<Outputs.GetNexusRepositoryDockerHostedComponentResult> components,

            ImmutableArray<Outputs.GetNexusRepositoryDockerHostedDockerResult> dockers,

            string id,

            string name,

            bool online,

            ImmutableArray<Outputs.GetNexusRepositoryDockerHostedStorageResult> storages)
        {
            Cleanups = cleanups;
            Components = components;
            Dockers = dockers;
            Id = id;
            Name = name;
            Online = online;
            Storages = storages;
        }
    }
}