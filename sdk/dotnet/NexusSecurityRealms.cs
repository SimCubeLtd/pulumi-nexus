// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus
{
    [NexusResourceType("nexus:index/nexusSecurityRealms:NexusSecurityRealms")]
    public partial class NexusSecurityRealms : Pulumi.CustomResource
    {
        /// <summary>
        /// The realm IDs
        /// </summary>
        [Output("actives")]
        public Output<ImmutableArray<string>> Actives { get; private set; } = null!;


        /// <summary>
        /// Create a NexusSecurityRealms resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NexusSecurityRealms(string name, NexusSecurityRealmsArgs args, CustomResourceOptions? options = null)
            : base("nexus:index/nexusSecurityRealms:NexusSecurityRealms", name, args ?? new NexusSecurityRealmsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NexusSecurityRealms(string name, Input<string> id, NexusSecurityRealmsState? state = null, CustomResourceOptions? options = null)
            : base("nexus:index/nexusSecurityRealms:NexusSecurityRealms", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NexusSecurityRealms resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NexusSecurityRealms Get(string name, Input<string> id, NexusSecurityRealmsState? state = null, CustomResourceOptions? options = null)
        {
            return new NexusSecurityRealms(name, id, state, options);
        }
    }

    public sealed class NexusSecurityRealmsArgs : Pulumi.ResourceArgs
    {
        [Input("actives", required: true)]
        private InputList<string>? _actives;

        /// <summary>
        /// The realm IDs
        /// </summary>
        public InputList<string> Actives
        {
            get => _actives ?? (_actives = new InputList<string>());
            set => _actives = value;
        }

        public NexusSecurityRealmsArgs()
        {
        }
    }

    public sealed class NexusSecurityRealmsState : Pulumi.ResourceArgs
    {
        [Input("actives")]
        private InputList<string>? _actives;

        /// <summary>
        /// The realm IDs
        /// </summary>
        public InputList<string> Actives
        {
            get => _actives ?? (_actives = new InputList<string>());
            set => _actives = value;
        }

        public NexusSecurityRealmsState()
        {
        }
    }
}
