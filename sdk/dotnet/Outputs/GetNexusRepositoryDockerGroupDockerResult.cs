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
    public sealed class GetNexusRepositoryDockerGroupDockerResult
    {
        public readonly bool ForceBasicAuth;
        public readonly int HttpPort;
        public readonly int HttpsPort;
        public readonly bool V1Enabled;

        [OutputConstructor]
        private GetNexusRepositoryDockerGroupDockerResult(
            bool forceBasicAuth,

            int httpPort,

            int httpsPort,

            bool v1Enabled)
        {
            ForceBasicAuth = forceBasicAuth;
            HttpPort = httpPort;
            HttpsPort = httpsPort;
            V1Enabled = v1Enabled;
        }
    }
}
