// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus
{
    [NexusResourceType("nexus:index/nexusBlobstoreS3:NexusBlobstoreS3")]
    public partial class NexusBlobstoreS3 : Pulumi.CustomResource
    {
        /// <summary>
        /// Count of blobs
        /// </summary>
        [Output("blobCount")]
        public Output<int> BlobCount { get; private set; } = null!;

        /// <summary>
        /// The S3 bucket configuration.
        /// </summary>
        [Output("bucketConfiguration")]
        public Output<Outputs.NexusBlobstoreS3BucketConfiguration> BucketConfiguration { get; private set; } = null!;

        /// <summary>
        /// Blobstore name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Soft quota of the blobstore
        /// </summary>
        [Output("softQuota")]
        public Output<Outputs.NexusBlobstoreS3SoftQuota?> SoftQuota { get; private set; } = null!;

        /// <summary>
        /// The total size of the blobstore in Bytes
        /// </summary>
        [Output("totalSizeInBytes")]
        public Output<int> TotalSizeInBytes { get; private set; } = null!;


        /// <summary>
        /// Create a NexusBlobstoreS3 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NexusBlobstoreS3(string name, NexusBlobstoreS3Args args, CustomResourceOptions? options = null)
            : base("nexus:index/nexusBlobstoreS3:NexusBlobstoreS3", name, args ?? new NexusBlobstoreS3Args(), MakeResourceOptions(options, ""))
        {
        }

        private NexusBlobstoreS3(string name, Input<string> id, NexusBlobstoreS3State? state = null, CustomResourceOptions? options = null)
            : base("nexus:index/nexusBlobstoreS3:NexusBlobstoreS3", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NexusBlobstoreS3 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NexusBlobstoreS3 Get(string name, Input<string> id, NexusBlobstoreS3State? state = null, CustomResourceOptions? options = null)
        {
            return new NexusBlobstoreS3(name, id, state, options);
        }
    }

    public sealed class NexusBlobstoreS3Args : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The S3 bucket configuration.
        /// </summary>
        [Input("bucketConfiguration", required: true)]
        public Input<Inputs.NexusBlobstoreS3BucketConfigurationArgs> BucketConfiguration { get; set; } = null!;

        /// <summary>
        /// Blobstore name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Soft quota of the blobstore
        /// </summary>
        [Input("softQuota")]
        public Input<Inputs.NexusBlobstoreS3SoftQuotaArgs>? SoftQuota { get; set; }

        public NexusBlobstoreS3Args()
        {
        }
    }

    public sealed class NexusBlobstoreS3State : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Count of blobs
        /// </summary>
        [Input("blobCount")]
        public Input<int>? BlobCount { get; set; }

        /// <summary>
        /// The S3 bucket configuration.
        /// </summary>
        [Input("bucketConfiguration")]
        public Input<Inputs.NexusBlobstoreS3BucketConfigurationGetArgs>? BucketConfiguration { get; set; }

        /// <summary>
        /// Blobstore name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Soft quota of the blobstore
        /// </summary>
        [Input("softQuota")]
        public Input<Inputs.NexusBlobstoreS3SoftQuotaGetArgs>? SoftQuota { get; set; }

        /// <summary>
        /// The total size of the blobstore in Bytes
        /// </summary>
        [Input("totalSizeInBytes")]
        public Input<int>? TotalSizeInBytes { get; set; }

        public NexusBlobstoreS3State()
        {
        }
    }
}
