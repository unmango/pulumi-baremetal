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
    public sealed class MvArgs
    {
        public readonly string? Backup;
        public readonly string? Destination;
        public readonly string? Directory;
        public readonly bool? Force;
        public readonly bool? Help;
        public readonly bool? NoClobber;
        public readonly bool? NoTargetDirectory;
        public readonly ImmutableArray<string> Source;
        public readonly bool? StripTrailingSlashes;
        public readonly string? Suffix;
        public readonly string? TargetDirectory;
        public readonly bool? Update;
        public readonly bool? Verbose;
        public readonly bool? Version;

        [OutputConstructor]
        private MvArgs(
            string? backup,

            string? destination,

            string? directory,

            bool? force,

            bool? help,

            bool? noClobber,

            bool? noTargetDirectory,

            ImmutableArray<string> source,

            bool? stripTrailingSlashes,

            string? suffix,

            string? targetDirectory,

            bool? update,

            bool? verbose,

            bool? version)
        {
            Backup = backup;
            Destination = destination;
            Directory = directory;
            Force = force;
            Help = help;
            NoClobber = noClobber;
            NoTargetDirectory = noTargetDirectory;
            Source = source;
            StripTrailingSlashes = stripTrailingSlashes;
            Suffix = suffix;
            TargetDirectory = targetDirectory;
            Update = update;
            Verbose = verbose;
            Version = version;
        }
    }
}
