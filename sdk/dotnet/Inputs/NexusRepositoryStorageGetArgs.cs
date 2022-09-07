// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus.Inputs
{

    public sealed class NexusRepositoryStorageGetArgs : Pulumi.ResourceArgs
    {
        [Input("blobStoreName")]
        public Input<string>? BlobStoreName { get; set; }

        [Input("strictContentTypeValidation")]
        public Input<bool>? StrictContentTypeValidation { get; set; }

        [Input("writePolicy")]
        public Input<string>? WritePolicy { get; set; }

        public NexusRepositoryStorageGetArgs()
        {
        }
    }
}