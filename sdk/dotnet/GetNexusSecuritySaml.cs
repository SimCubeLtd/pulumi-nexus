// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus
{
    public static class GetNexusSecuritySaml
    {
        public static Task<GetNexusSecuritySamlResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNexusSecuritySamlResult>("nexus:index/getNexusSecuritySaml:GetNexusSecuritySaml", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetNexusSecuritySamlResult
    {
        public readonly string EmailAttribute;
        public readonly string EntityId;
        public readonly string FirstNameAttribute;
        public readonly string GroupsAttribute;
        public readonly string Id;
        public readonly string IdpMetadata;
        public readonly string LastNameAttribute;
        public readonly string UsernameAttribute;
        public readonly bool ValidateAssertionSignature;
        public readonly bool ValidateResponseSignature;

        [OutputConstructor]
        private GetNexusSecuritySamlResult(
            string emailAttribute,

            string entityId,

            string firstNameAttribute,

            string groupsAttribute,

            string id,

            string idpMetadata,

            string lastNameAttribute,

            string usernameAttribute,

            bool validateAssertionSignature,

            bool validateResponseSignature)
        {
            EmailAttribute = emailAttribute;
            EntityId = entityId;
            FirstNameAttribute = firstNameAttribute;
            GroupsAttribute = groupsAttribute;
            Id = id;
            IdpMetadata = idpMetadata;
            LastNameAttribute = lastNameAttribute;
            UsernameAttribute = usernameAttribute;
            ValidateAssertionSignature = validateAssertionSignature;
            ValidateResponseSignature = validateResponseSignature;
        }
    }
}
