// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.Baremetal.Coreutils.Outputs
{

    [OutputType]
    public sealed class TeeArgs
    {
        public readonly bool? Append;
        public readonly ImmutableArray<string> Files;
        public readonly string? Stdin;

        [OutputConstructor]
        private TeeArgs(
            bool? append,

            ImmutableArray<string> files,

            string? stdin)
        {
            Append = append;
            Files = files;
            Stdin = stdin;
        }
    }
}
