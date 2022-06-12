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
    public sealed class ItemPasswordRecipe
    {
        /// <summary>
        /// Use digits [0-9] when generating the password.
        /// </summary>
        public readonly bool? Digits;
        /// <summary>
        /// The length of the password to be generated.
        /// </summary>
        public readonly int? Length;
        /// <summary>
        /// Use letters [a-zA-Z] when generating the password.
        /// </summary>
        public readonly bool? Letters;
        /// <summary>
        /// Use symbols [!@.-_*] when generating the password.
        /// </summary>
        public readonly bool? Symbols;

        [OutputConstructor]
        private ItemPasswordRecipe(
            bool? digits,

            int? length,

            bool? letters,

            bool? symbols)
        {
            Digits = digits;
            Length = length;
            Letters = letters;
            Symbols = symbols;
        }
    }
}
