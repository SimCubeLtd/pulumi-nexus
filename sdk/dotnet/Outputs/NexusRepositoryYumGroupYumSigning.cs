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
    public sealed class NexusRepositoryYumGroupYumSigning
    {
        public readonly string Keypair;
        public readonly string? Passphrase;

        [OutputConstructor]
        private NexusRepositoryYumGroupYumSigning(
            string keypair,

            string? passphrase)
        {
            Keypair = keypair;
            Passphrase = passphrase;
        }
    }
}
