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
    public sealed class ChmodArgs
    {
        public readonly bool? Changes;
        public readonly ImmutableArray<string> Files;
        public readonly bool? Help;
        public readonly ImmutableArray<string> Mode;
        public readonly bool? NoPreserveRoot;
        public readonly string? OctalMode;
        public readonly bool? PreserveRoot;
        public readonly bool? Quiet;
        public readonly bool? Recursive;
        public readonly string? Reference;
        public readonly bool? Verbose;
        public readonly bool? Version;

        [OutputConstructor]
        private ChmodArgs(
            bool? changes,

            ImmutableArray<string> files,

            bool? help,

            ImmutableArray<string> mode,

            bool? noPreserveRoot,

            string? octalMode,

            bool? preserveRoot,

            bool? quiet,

            bool? recursive,

            string? reference,

            bool? verbose,

            bool? version)
        {
            Changes = changes;
            Files = files;
            Help = help;
            Mode = mode;
            NoPreserveRoot = noPreserveRoot;
            OctalMode = octalMode;
            PreserveRoot = preserveRoot;
            Quiet = quiet;
            Recursive = recursive;
            Reference = reference;
            Verbose = verbose;
            Version = version;
        }
    }
}
