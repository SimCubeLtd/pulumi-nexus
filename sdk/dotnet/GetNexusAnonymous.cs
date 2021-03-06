// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus
{
    public static class GetNexusAnonymous
    {
        public static Task<GetNexusAnonymousResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNexusAnonymousResult>("nexus:index/getNexusAnonymous:GetNexusAnonymous", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetNexusAnonymousResult
    {
        public readonly bool Enabled;
        public readonly string Id;
        public readonly string RealmName;
        public readonly string UserId;

        [OutputConstructor]
        private GetNexusAnonymousResult(
            bool enabled,

            string id,

            string realmName,

            string userId)
        {
            Enabled = enabled;
            Id = id;
            RealmName = realmName;
            UserId = userId;
        }
    }
}
