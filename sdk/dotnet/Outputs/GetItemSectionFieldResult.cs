// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.nexus.Outputs
{

    [OutputType]
    public sealed class GetItemSectionFieldResult
    {
        public readonly string Id;
        public readonly string Label;
        public readonly string Purpose;
        /// <summary>
        /// (Only applies to the database category) The type of database. One of ["db2" "filemaker" "msaccess" "mssql" "mysql" "oracle" "postgresql" "sqlite" "other"]
        /// </summary>
        public readonly string Type;
        public readonly string Value;

        [OutputConstructor]
        private GetItemSectionFieldResult(
            string id,

            string label,

            string purpose,

            string type,

            string value)
        {
            Id = id;
            Label = label;
            Purpose = purpose;
            Type = type;
            Value = value;
        }
    }
}