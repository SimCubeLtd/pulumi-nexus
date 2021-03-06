// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus
{
    [NexusResourceType("nexus:index/nexusSecurityUserToken:NexusSecurityUserToken")]
    public partial class NexusSecurityUserToken : Pulumi.CustomResource
    {
        /// <summary>
        /// Activate the feature of user tokens.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// Require user tokens for repository authentication. This does not effect UI access.
        /// </summary>
        [Output("protectContent")]
        public Output<bool?> ProtectContent { get; private set; } = null!;


        /// <summary>
        /// Create a NexusSecurityUserToken resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NexusSecurityUserToken(string name, NexusSecurityUserTokenArgs args, CustomResourceOptions? options = null)
            : base("nexus:index/nexusSecurityUserToken:NexusSecurityUserToken", name, args ?? new NexusSecurityUserTokenArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NexusSecurityUserToken(string name, Input<string> id, NexusSecurityUserTokenState? state = null, CustomResourceOptions? options = null)
            : base("nexus:index/nexusSecurityUserToken:NexusSecurityUserToken", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NexusSecurityUserToken resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NexusSecurityUserToken Get(string name, Input<string> id, NexusSecurityUserTokenState? state = null, CustomResourceOptions? options = null)
        {
            return new NexusSecurityUserToken(name, id, state, options);
        }
    }

    public sealed class NexusSecurityUserTokenArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Activate the feature of user tokens.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// Require user tokens for repository authentication. This does not effect UI access.
        /// </summary>
        [Input("protectContent")]
        public Input<bool>? ProtectContent { get; set; }

        public NexusSecurityUserTokenArgs()
        {
        }
    }

    public sealed class NexusSecurityUserTokenState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Activate the feature of user tokens.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Require user tokens for repository authentication. This does not effect UI access.
        /// </summary>
        [Input("protectContent")]
        public Input<bool>? ProtectContent { get; set; }

        public NexusSecurityUserTokenState()
        {
        }
    }
}
