// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus
{
    [NexusResourceType("nexus:index/nexusRole:NexusRole")]
    public partial class NexusRole : Pulumi.CustomResource
    {
        /// <summary>
        /// The description of this role.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the role.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The privileges of this role.
        /// </summary>
        [Output("privileges")]
        public Output<ImmutableArray<string>> Privileges { get; private set; } = null!;

        /// <summary>
        /// The id of the role.
        /// </summary>
        [Output("roleid")]
        public Output<string> Roleid { get; private set; } = null!;

        /// <summary>
        /// The roles of this role.
        /// </summary>
        [Output("roles")]
        public Output<ImmutableArray<string>> Roles { get; private set; } = null!;


        /// <summary>
        /// Create a NexusRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NexusRole(string name, NexusRoleArgs args, CustomResourceOptions? options = null)
            : base("nexus:index/nexusRole:NexusRole", name, args ?? new NexusRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NexusRole(string name, Input<string> id, NexusRoleState? state = null, CustomResourceOptions? options = null)
            : base("nexus:index/nexusRole:NexusRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NexusRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NexusRole Get(string name, Input<string> id, NexusRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new NexusRole(name, id, state, options);
        }
    }

    public sealed class NexusRoleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of this role.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the role.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("privileges")]
        private InputList<string>? _privileges;

        /// <summary>
        /// The privileges of this role.
        /// </summary>
        public InputList<string> Privileges
        {
            get => _privileges ?? (_privileges = new InputList<string>());
            set => _privileges = value;
        }

        /// <summary>
        /// The id of the role.
        /// </summary>
        [Input("roleid", required: true)]
        public Input<string> Roleid { get; set; } = null!;

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// The roles of this role.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        public NexusRoleArgs()
        {
        }
    }

    public sealed class NexusRoleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of this role.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the role.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("privileges")]
        private InputList<string>? _privileges;

        /// <summary>
        /// The privileges of this role.
        /// </summary>
        public InputList<string> Privileges
        {
            get => _privileges ?? (_privileges = new InputList<string>());
            set => _privileges = value;
        }

        /// <summary>
        /// The id of the role.
        /// </summary>
        [Input("roleid")]
        public Input<string>? Roleid { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;

        /// <summary>
        /// The roles of this role.
        /// </summary>
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        public NexusRoleState()
        {
        }
    }
}
