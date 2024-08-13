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
    public sealed class TarArgs
    {
        public readonly bool? Anchored;
        public readonly bool? Append;
        public readonly ImmutableArray<string> Args;
        public readonly bool? Bzip2;
        public readonly bool? Create;
        public readonly bool? Delete;
        public readonly bool? Diff;
        public readonly string? Directory;
        public readonly string? Exclude;
        public readonly string? ExcludeFrom;
        public readonly bool? ExcludeVcs;
        public readonly bool? ExcludeVcsIgnores;
        public readonly bool? Extract;
        public readonly string? File;
        public readonly bool? Gzip;
        public readonly bool? IgnoreCommandError;
        public readonly bool? KeepDirectorySymlink;
        public readonly bool? KeepNewerFiles;
        public readonly bool? KeepOldfiles;
        public readonly bool? List;
        public readonly bool? Lzip;
        public readonly bool? Lzma;
        public readonly bool? Lzop;
        public readonly bool? NoAnchored;
        public readonly bool? NoOverwriteDir;
        public readonly bool? NoSeek;
        public readonly bool? Overwrite;
        public readonly bool? OverwriteDir;
        public readonly bool? RemoveFiles;
        public readonly bool? SkipOldFiles;
        public readonly bool? Sparse;
        public readonly int? StripComponents;
        public readonly string? Suffix;
        public readonly bool? ToStdout;
        public readonly string? Transform;
        public readonly bool? UnlinkFirst;
        public readonly bool? Update;
        public readonly bool? Verbose;
        public readonly bool? Verify;
        public readonly bool? Version;
        public readonly bool? Xz;
        public readonly bool? Zstd;

        [OutputConstructor]
        private TarArgs(
            bool? anchored,

            bool? append,

            ImmutableArray<string> args,

            bool? bzip2,

            bool? create,

            bool? delete,

            bool? diff,

            string? directory,

            string? exclude,

            string? excludeFrom,

            bool? excludeVcs,

            bool? excludeVcsIgnores,

            bool? extract,

            string? file,

            bool? gzip,

            bool? ignoreCommandError,

            bool? keepDirectorySymlink,

            bool? keepNewerFiles,

            bool? keepOldfiles,

            bool? list,

            bool? lzip,

            bool? lzma,

            bool? lzop,

            bool? noAnchored,

            bool? noOverwriteDir,

            bool? noSeek,

            bool? overwrite,

            bool? overwriteDir,

            bool? removeFiles,

            bool? skipOldFiles,

            bool? sparse,

            int? stripComponents,

            string? suffix,

            bool? toStdout,

            string? transform,

            bool? unlinkFirst,

            bool? update,

            bool? verbose,

            bool? verify,

            bool? version,

            bool? xz,

            bool? zstd)
        {
            Anchored = anchored;
            Append = append;
            Args = args;
            Bzip2 = bzip2;
            Create = create;
            Delete = delete;
            Diff = diff;
            Directory = directory;
            Exclude = exclude;
            ExcludeFrom = excludeFrom;
            ExcludeVcs = excludeVcs;
            ExcludeVcsIgnores = excludeVcsIgnores;
            Extract = extract;
            File = file;
            Gzip = gzip;
            IgnoreCommandError = ignoreCommandError;
            KeepDirectorySymlink = keepDirectorySymlink;
            KeepNewerFiles = keepNewerFiles;
            KeepOldfiles = keepOldfiles;
            List = list;
            Lzip = lzip;
            Lzma = lzma;
            Lzop = lzop;
            NoAnchored = noAnchored;
            NoOverwriteDir = noOverwriteDir;
            NoSeek = noSeek;
            Overwrite = overwrite;
            OverwriteDir = overwriteDir;
            RemoveFiles = removeFiles;
            SkipOldFiles = skipOldFiles;
            Sparse = sparse;
            StripComponents = stripComponents;
            Suffix = suffix;
            ToStdout = toStdout;
            Transform = transform;
            UnlinkFirst = unlinkFirst;
            Update = update;
            Verbose = verbose;
            Verify = verify;
            Version = version;
            Xz = xz;
            Zstd = zstd;
        }
    }
}
