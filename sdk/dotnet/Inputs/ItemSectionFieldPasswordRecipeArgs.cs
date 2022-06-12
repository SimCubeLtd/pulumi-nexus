// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.nexus.Inputs
{

    public sealed class ItemSectionFieldPasswordRecipeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Use digits [0-9] when generating the password.
        /// </summary>
        [Input("digits")]
        public Input<bool>? Digits { get; set; }

        /// <summary>
        /// The length of the password to be generated.
        /// </summary>
        [Input("length")]
        public Input<int>? Length { get; set; }

        /// <summary>
        /// Use letters [a-zA-Z] when generating the password.
        /// </summary>
        [Input("letters")]
        public Input<bool>? Letters { get; set; }

        /// <summary>
        /// Use symbols [!@.-_*] when generating the password.
        /// </summary>
        [Input("symbols")]
        public Input<bool>? Symbols { get; set; }

        public ItemSectionFieldPasswordRecipeArgs()
        {
        }
    }
}
