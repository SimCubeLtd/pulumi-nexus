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
    public sealed class GetNexusSecurityLdapLdapResult
    {
        public readonly string AuthPassword;
        public readonly string AuthRealm;
        public readonly string AuthSchema;
        public readonly string AuthUsername;
        public readonly int ConnectionRetryDelaySeconds;
        public readonly int ConnectionTimeoutSeconds;
        public readonly string GroupBaseDn;
        public readonly string GroupIdAttribute;
        public readonly string GroupMemberAttribute;
        public readonly string GroupMemberFormat;
        public readonly string GroupObjectClass;
        public readonly string GroupSubtree;
        public readonly string GroupType;
        public readonly string Host;
        public readonly string Id;
        public readonly bool LdapGroupsAsRoles;
        public readonly int MaxIncidentCount;
        public readonly string Name;
        public readonly int Port;
        public readonly string Protocol;
        public readonly string SearchBase;
        public readonly bool UseTrustStore;
        public readonly string UserBaseDn;
        public readonly string UserEmailAddressAttribute;
        public readonly string UserIdAttribute;
        public readonly string UserLdapFilter;
        public readonly string UserMemberOfAttribute;
        public readonly string UserObjectClass;
        public readonly string UserPasswordAttribute;
        public readonly string UserRealNameAttribute;
        public readonly bool UserSubtree;

        [OutputConstructor]
        private GetNexusSecurityLdapLdapResult(
            string authPassword,

            string authRealm,

            string authSchema,

            string authUsername,

            int connectionRetryDelaySeconds,

            int connectionTimeoutSeconds,

            string groupBaseDn,

            string groupIdAttribute,

            string groupMemberAttribute,

            string groupMemberFormat,

            string groupObjectClass,

            string groupSubtree,

            string groupType,

            string host,

            string id,

            bool ldapGroupsAsRoles,

            int maxIncidentCount,

            string name,

            int port,

            string protocol,

            string searchBase,

            bool useTrustStore,

            string userBaseDn,

            string userEmailAddressAttribute,

            string userIdAttribute,

            string userLdapFilter,

            string userMemberOfAttribute,

            string userObjectClass,

            string userPasswordAttribute,

            string userRealNameAttribute,

            bool userSubtree)
        {
            AuthPassword = authPassword;
            AuthRealm = authRealm;
            AuthSchema = authSchema;
            AuthUsername = authUsername;
            ConnectionRetryDelaySeconds = connectionRetryDelaySeconds;
            ConnectionTimeoutSeconds = connectionTimeoutSeconds;
            GroupBaseDn = groupBaseDn;
            GroupIdAttribute = groupIdAttribute;
            GroupMemberAttribute = groupMemberAttribute;
            GroupMemberFormat = groupMemberFormat;
            GroupObjectClass = groupObjectClass;
            GroupSubtree = groupSubtree;
            GroupType = groupType;
            Host = host;
            Id = id;
            LdapGroupsAsRoles = ldapGroupsAsRoles;
            MaxIncidentCount = maxIncidentCount;
            Name = name;
            Port = port;
            Protocol = protocol;
            SearchBase = searchBase;
            UseTrustStore = useTrustStore;
            UserBaseDn = userBaseDn;
            UserEmailAddressAttribute = userEmailAddressAttribute;
            UserIdAttribute = userIdAttribute;
            UserLdapFilter = userLdapFilter;
            UserMemberOfAttribute = userMemberOfAttribute;
            UserObjectClass = userObjectClass;
            UserPasswordAttribute = userPasswordAttribute;
            UserRealNameAttribute = userRealNameAttribute;
            UserSubtree = userSubtree;
        }
    }
}