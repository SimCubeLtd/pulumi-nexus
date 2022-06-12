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
    public sealed class GetNexusRepositoryDockerProxyProxyResult
    {
        public readonly int ContentMaxAge;
        public readonly int MetadataMaxAge;
        public readonly string RemoteUrl;

        [OutputConstructor]
        private GetNexusRepositoryDockerProxyProxyResult(
            int contentMaxAge,

            int metadataMaxAge,

            string remoteUrl)
        {
            ContentMaxAge = contentMaxAge;
            MetadataMaxAge = metadataMaxAge;
            RemoteUrl = remoteUrl;
        }
    }
}
