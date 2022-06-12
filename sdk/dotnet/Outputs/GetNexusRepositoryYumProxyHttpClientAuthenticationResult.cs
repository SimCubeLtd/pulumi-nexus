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
    public sealed class GetNexusRepositoryYumProxyHttpClientAuthenticationResult
    {
        public readonly string NtlmDomain;
        public readonly string NtlmHost;
        public readonly string Password;
        public readonly string Type;
        public readonly string Username;

        [OutputConstructor]
        private GetNexusRepositoryYumProxyHttpClientAuthenticationResult(
            string ntlmDomain,

            string ntlmHost,

            string password,

            string type,

            string username)
        {
            NtlmDomain = ntlmDomain;
            NtlmHost = ntlmHost;
            Password = password;
            Type = type;
            Username = username;
        }
    }
}
