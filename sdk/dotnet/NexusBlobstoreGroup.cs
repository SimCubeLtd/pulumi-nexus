// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus
{
    [NexusResourceType("nexus:index/nexusBlobstoreGroup:NexusBlobstoreGroup")]
    public partial class NexusBlobstoreGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// Available space in Bytes
        /// </summary>
        [Output("availableSpaceInBytes")]
        public Output<int> AvailableSpaceInBytes { get; private set; } = null!;

        /// <summary>
        /// Count of blobs
        /// </summary>
        [Output("blobCount")]
        public Output<int> BlobCount { get; private set; } = null!;

        /// <summary>
        /// The policy how to fill the members. Possible values: `roundRobin` or `writeToFirst`
        /// </summary>
        [Output("fillPolicy")]
        public Output<string> FillPolicy { get; private set; } = null!;

        /// <summary>
        /// List of the names of blob stores that are members of this group
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// Blobstore name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Soft quota of the blobstore
        /// </summary>
        [Output("softQuota")]
        public Output<Outputs.NexusBlobstoreGroupSoftQuota?> SoftQuota { get; private set; } = null!;

        /// <summary>
        /// The total size of the blobstore in Bytes
        /// </summary>
        [Output("totalSizeInBytes")]
        public Output<int> TotalSizeInBytes { get; private set; } = null!;


        /// <summary>
        /// Create a NexusBlobstoreGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NexusBlobstoreGroup(string name, NexusBlobstoreGroupArgs args, CustomResourceOptions? options = null)
            : base("nexus:index/nexusBlobstoreGroup:NexusBlobstoreGroup", name, args ?? new NexusBlobstoreGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NexusBlobstoreGroup(string name, Input<string> id, NexusBlobstoreGroupState? state = null, CustomResourceOptions? options = null)
            : base("nexus:index/nexusBlobstoreGroup:NexusBlobstoreGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NexusBlobstoreGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NexusBlobstoreGroup Get(string name, Input<string> id, NexusBlobstoreGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new NexusBlobstoreGroup(name, id, state, options);
        }
    }

    public sealed class NexusBlobstoreGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The policy how to fill the members. Possible values: `roundRobin` or `writeToFirst`
        /// </summary>
        [Input("fillPolicy", required: true)]
        public Input<string> FillPolicy { get; set; } = null!;

        [Input("members", required: true)]
        private InputList<string>? _members;

        /// <summary>
        /// List of the names of blob stores that are members of this group
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// Blobstore name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Soft quota of the blobstore
        /// </summary>
        [Input("softQuota")]
        public Input<Inputs.NexusBlobstoreGroupSoftQuotaArgs>? SoftQuota { get; set; }

        public NexusBlobstoreGroupArgs()
        {
        }
    }

    public sealed class NexusBlobstoreGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Available space in Bytes
        /// </summary>
        [Input("availableSpaceInBytes")]
        public Input<int>? AvailableSpaceInBytes { get; set; }

        /// <summary>
        /// Count of blobs
        /// </summary>
        [Input("blobCount")]
        public Input<int>? BlobCount { get; set; }

        /// <summary>
        /// The policy how to fill the members. Possible values: `roundRobin` or `writeToFirst`
        /// </summary>
        [Input("fillPolicy")]
        public Input<string>? FillPolicy { get; set; }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// List of the names of blob stores that are members of this group
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// Blobstore name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Soft quota of the blobstore
        /// </summary>
        [Input("softQuota")]
        public Input<Inputs.NexusBlobstoreGroupSoftQuotaGetArgs>? SoftQuota { get; set; }

        /// <summary>
        /// The total size of the blobstore in Bytes
        /// </summary>
        [Input("totalSizeInBytes")]
        public Input<int>? TotalSizeInBytes { get; set; }

        public NexusBlobstoreGroupState()
        {
        }
    }
}