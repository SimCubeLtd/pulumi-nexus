// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus.Inputs
{

    public sealed class NexusRepositoryDockerHostedCleanupGetArgs : Pulumi.ResourceArgs
    {
        [Input("policyNames")]
        private InputList<string>? _policyNames;
        public InputList<string> PolicyNames
        {
            get => _policyNames ?? (_policyNames = new InputList<string>());
            set => _policyNames = value;
        }

        public NexusRepositoryDockerHostedCleanupGetArgs()
        {
        }
    }
}
