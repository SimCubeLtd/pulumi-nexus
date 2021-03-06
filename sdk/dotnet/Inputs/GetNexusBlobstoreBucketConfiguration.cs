// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Nexus.Inputs
{

    public sealed class GetNexusBlobstoreBucketConfigurationArgs : Pulumi.InvokeArgs
    {
        [Input("advancedBucketConnection")]
        public Inputs.GetNexusBlobstoreBucketConfigurationAdvancedBucketConnectionArgs? AdvancedBucketConnection { get; set; }

        [Input("bucket", required: true)]
        public Inputs.GetNexusBlobstoreBucketConfigurationBucketArgs Bucket { get; set; } = null!;

        [Input("bucketSecurity", required: true)]
        public Inputs.GetNexusBlobstoreBucketConfigurationBucketSecurityArgs BucketSecurity { get; set; } = null!;

        [Input("encryption")]
        public Inputs.GetNexusBlobstoreBucketConfigurationEncryptionArgs? Encryption { get; set; }

        public GetNexusBlobstoreBucketConfigurationArgs()
        {
        }
    }
}
