// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus
{
    [NexusResourceType("nexus:index/nexusScript:NexusScript")]
    public partial class NexusScript : Pulumi.CustomResource
    {
        /// <summary>
        /// The content of this script.
        /// </summary>
        [Output("content")]
        public Output<string> Content { get; private set; } = null!;

        /// <summary>
        /// The name of the script.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The type of the script. Default: `groovy`
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a NexusScript resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NexusScript(string name, NexusScriptArgs args, CustomResourceOptions? options = null)
            : base("nexus:index/nexusScript:NexusScript", name, args ?? new NexusScriptArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NexusScript(string name, Input<string> id, NexusScriptState? state = null, CustomResourceOptions? options = null)
            : base("nexus:index/nexusScript:NexusScript", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NexusScript resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NexusScript Get(string name, Input<string> id, NexusScriptState? state = null, CustomResourceOptions? options = null)
        {
            return new NexusScript(name, id, state, options);
        }
    }

    public sealed class NexusScriptArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The content of this script.
        /// </summary>
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        /// <summary>
        /// The name of the script.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The type of the script. Default: `groovy`
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public NexusScriptArgs()
        {
        }
    }

    public sealed class NexusScriptState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The content of this script.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// The name of the script.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The type of the script. Default: `groovy`
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public NexusScriptState()
        {
        }
    }
}