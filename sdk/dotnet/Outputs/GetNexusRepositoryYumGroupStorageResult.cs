// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus.Outputs
{

    [OutputType]
    public sealed class GetNexusRepositoryYumGroupStorageResult
    {
        public readonly string BlobStoreName;
        public readonly bool StrictContentTypeValidation;

        [OutputConstructor]
        private GetNexusRepositoryYumGroupStorageResult(
            string blobStoreName,

            bool strictContentTypeValidation)
        {
            BlobStoreName = blobStoreName;
            StrictContentTypeValidation = strictContentTypeValidation;
        }
    }
}
