// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus.Inputs
{

    public sealed class NexusRepositoryDockerGroupGroupArgs : Pulumi.ResourceArgs
    {
        [Input("memberNames", required: true)]
        private InputList<string>? _memberNames;
        public InputList<string> MemberNames
        {
            get => _memberNames ?? (_memberNames = new InputList<string>());
            set => _memberNames = value;
        }

        [Input("writableMember")]
        public Input<string>? WritableMember { get; set; }

        public NexusRepositoryDockerGroupGroupArgs()
        {
        }
    }
}
