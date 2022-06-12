// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus.Inputs
{

    public sealed class GetNexusRepositoryHttpClientAuthenticationArgs : Pulumi.InvokeArgs
    {
        [Input("ntlmDomain")]
        public string? NtlmDomain { get; set; }

        [Input("ntlmHost")]
        public string? NtlmHost { get; set; }

        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        [Input("username")]
        public string? Username { get; set; }

        public GetNexusRepositoryHttpClientAuthenticationArgs()
        {
        }
    }
}
