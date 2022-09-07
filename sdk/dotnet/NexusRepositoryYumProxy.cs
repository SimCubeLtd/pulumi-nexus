// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus
{
    [NexusResourceType("nexus:index/nexusRepositoryYumProxy:NexusRepositoryYumProxy")]
    public partial class NexusRepositoryYumProxy : Pulumi.CustomResource
    {
        /// <summary>
        /// Cleanup policies
        /// </summary>
        [Output("cleanups")]
        public Output<ImmutableArray<Outputs.NexusRepositoryYumProxyCleanup>> Cleanups { get; private set; } = null!;

        /// <summary>
        /// HTTP Client configuration for proxy repositories. Required for docker proxy repositories
        /// </summary>
        [Output("httpClient")]
        public Output<Outputs.NexusRepositoryYumProxyHttpClient?> HttpClient { get; private set; } = null!;

        /// <summary>
        /// A unique identifier for this repository
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Configuration of the negative cache handling
        /// </summary>
        [Output("negativeCache")]
        public Output<Outputs.NexusRepositoryYumProxyNegativeCache?> NegativeCache { get; private set; } = null!;

        /// <summary>
        /// Whether this repository accepts incoming requests
        /// </summary>
        [Output("online")]
        public Output<bool?> Online { get; private set; } = null!;

        /// <summary>
        /// Configuration for the proxy repository
        /// </summary>
        [Output("proxy")]
        public Output<Outputs.NexusRepositoryYumProxyProxy> Proxy { get; private set; } = null!;

        /// <summary>
        /// The name of the routing rule assigned to this repository
        /// </summary>
        [Output("routingRule")]
        public Output<string?> RoutingRule { get; private set; } = null!;

        /// <summary>
        /// The storage configuration of the repository
        /// </summary>
        [Output("storage")]
        public Output<Outputs.NexusRepositoryYumProxyStorage> Storage { get; private set; } = null!;

        /// <summary>
        /// Contains signing data of repositores
        /// </summary>
        [Output("yumSigning")]
        public Output<Outputs.NexusRepositoryYumProxyYumSigning?> YumSigning { get; private set; } = null!;


        /// <summary>
        /// Create a NexusRepositoryYumProxy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NexusRepositoryYumProxy(string name, NexusRepositoryYumProxyArgs args, CustomResourceOptions? options = null)
            : base("nexus:index/nexusRepositoryYumProxy:NexusRepositoryYumProxy", name, args ?? new NexusRepositoryYumProxyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NexusRepositoryYumProxy(string name, Input<string> id, NexusRepositoryYumProxyState? state = null, CustomResourceOptions? options = null)
            : base("nexus:index/nexusRepositoryYumProxy:NexusRepositoryYumProxy", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/SimCubeLtd/pulumi-nexus/releases/download/v${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NexusRepositoryYumProxy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NexusRepositoryYumProxy Get(string name, Input<string> id, NexusRepositoryYumProxyState? state = null, CustomResourceOptions? options = null)
        {
            return new NexusRepositoryYumProxy(name, id, state, options);
        }
    }

    public sealed class NexusRepositoryYumProxyArgs : Pulumi.ResourceArgs
    {
        [Input("cleanups")]
        private InputList<Inputs.NexusRepositoryYumProxyCleanupArgs>? _cleanups;

        /// <summary>
        /// Cleanup policies
        /// </summary>
        public InputList<Inputs.NexusRepositoryYumProxyCleanupArgs> Cleanups
        {
            get => _cleanups ?? (_cleanups = new InputList<Inputs.NexusRepositoryYumProxyCleanupArgs>());
            set => _cleanups = value;
        }

        /// <summary>
        /// HTTP Client configuration for proxy repositories. Required for docker proxy repositories
        /// </summary>
        [Input("httpClient")]
        public Input<Inputs.NexusRepositoryYumProxyHttpClientArgs>? HttpClient { get; set; }

        /// <summary>
        /// A unique identifier for this repository
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configuration of the negative cache handling
        /// </summary>
        [Input("negativeCache")]
        public Input<Inputs.NexusRepositoryYumProxyNegativeCacheArgs>? NegativeCache { get; set; }

        /// <summary>
        /// Whether this repository accepts incoming requests
        /// </summary>
        [Input("online")]
        public Input<bool>? Online { get; set; }

        /// <summary>
        /// Configuration for the proxy repository
        /// </summary>
        [Input("proxy", required: true)]
        public Input<Inputs.NexusRepositoryYumProxyProxyArgs> Proxy { get; set; } = null!;

        /// <summary>
        /// The name of the routing rule assigned to this repository
        /// </summary>
        [Input("routingRule")]
        public Input<string>? RoutingRule { get; set; }

        /// <summary>
        /// The storage configuration of the repository
        /// </summary>
        [Input("storage", required: true)]
        public Input<Inputs.NexusRepositoryYumProxyStorageArgs> Storage { get; set; } = null!;

        /// <summary>
        /// Contains signing data of repositores
        /// </summary>
        [Input("yumSigning")]
        public Input<Inputs.NexusRepositoryYumProxyYumSigningArgs>? YumSigning { get; set; }

        public NexusRepositoryYumProxyArgs()
        {
        }
    }

    public sealed class NexusRepositoryYumProxyState : Pulumi.ResourceArgs
    {
        [Input("cleanups")]
        private InputList<Inputs.NexusRepositoryYumProxyCleanupGetArgs>? _cleanups;

        /// <summary>
        /// Cleanup policies
        /// </summary>
        public InputList<Inputs.NexusRepositoryYumProxyCleanupGetArgs> Cleanups
        {
            get => _cleanups ?? (_cleanups = new InputList<Inputs.NexusRepositoryYumProxyCleanupGetArgs>());
            set => _cleanups = value;
        }

        /// <summary>
        /// HTTP Client configuration for proxy repositories. Required for docker proxy repositories
        /// </summary>
        [Input("httpClient")]
        public Input<Inputs.NexusRepositoryYumProxyHttpClientGetArgs>? HttpClient { get; set; }

        /// <summary>
        /// A unique identifier for this repository
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configuration of the negative cache handling
        /// </summary>
        [Input("negativeCache")]
        public Input<Inputs.NexusRepositoryYumProxyNegativeCacheGetArgs>? NegativeCache { get; set; }

        /// <summary>
        /// Whether this repository accepts incoming requests
        /// </summary>
        [Input("online")]
        public Input<bool>? Online { get; set; }

        /// <summary>
        /// Configuration for the proxy repository
        /// </summary>
        [Input("proxy")]
        public Input<Inputs.NexusRepositoryYumProxyProxyGetArgs>? Proxy { get; set; }

        /// <summary>
        /// The name of the routing rule assigned to this repository
        /// </summary>
        [Input("routingRule")]
        public Input<string>? RoutingRule { get; set; }

        /// <summary>
        /// The storage configuration of the repository
        /// </summary>
        [Input("storage")]
        public Input<Inputs.NexusRepositoryYumProxyStorageGetArgs>? Storage { get; set; }

        /// <summary>
        /// Contains signing data of repositores
        /// </summary>
        [Input("yumSigning")]
        public Input<Inputs.NexusRepositoryYumProxyYumSigningGetArgs>? YumSigning { get; set; }

        public NexusRepositoryYumProxyState()
        {
        }
    }
}