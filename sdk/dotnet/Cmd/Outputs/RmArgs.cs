// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.Baremetal.Cmd.Outputs
{

    [OutputType]
    public sealed class RmArgs
    {
        public readonly bool? Dir;
        public readonly ImmutableArray<string> Files;
        public readonly bool? Force;
        public readonly bool? Help;
        public readonly bool? OneFileSystem;
        public readonly bool? Recursive;
        public readonly bool? Verbose;

        [OutputConstructor]
        private RmArgs(
            bool? dir,

            ImmutableArray<string> files,

            bool? force,

            bool? help,

            bool? oneFileSystem,

            bool? recursive,

            bool? verbose)
        {
            Dir = dir;
            Files = files;
            Force = force;
            Help = help;
            OneFileSystem = oneFileSystem;
            Recursive = recursive;
            Verbose = verbose;
        }
    }
}
