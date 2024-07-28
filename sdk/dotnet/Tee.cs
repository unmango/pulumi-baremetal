// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.Baremetal
{
    /// <summary>
    /// TEE(1)                                                       User Commands                                                      TEE(1)
    /// 
    /// NAME
    ///        tee - read from standard input and write to standard output and files
    /// 
    /// SYNOPSIS
    ///        tee [OPTION]... [FILE]...
    /// 
    /// DESCRIPTION
    ///        Copy standard input to each FILE, and also to standard output.
    /// 
    ///        -a, --append
    ///               append to the given FILEs, do not overwrite
    ///     
    ///        -i, --ignore-interrupts
    ///               ignore interrupt signals
    ///     
    ///        -p     operate in a more appropriate MODE with pipes.
    ///     
    ///        --output-error[=MODE]
    ///               set behavior on write error.  See MODE below
    ///     
    ///        --help display this help and exit
    ///     
    ///        --version
    ///               output version information and exit
    /// 
    ///    MODE determines behavior with write errors on the outputs:
    ///        warn   diagnose errors writing to any output
    /// 
    ///        warn-nopipe
    ///               diagnose errors writing to any output not a pipe
    ///     
    ///        exit   exit on error writing to any output
    ///     
    ///        exit-nopipe
    ///               exit on error writing to any output not a pipe
    ///     
    ///        The default MODE for the -p option is 'warn-nopipe'.  With "nopipe" MODEs, exit immediately if all outputs become broken pipes.
    ///        The default operation when --output-error is not specified, is to exit immediately on error writing to a pipe, and diagnose er‐
    ///        rors writing to non pipe outputs.
    /// 
    /// AUTHOR
    ///        Written by Mike Parker, Richard M. Stallman, and David MacKenzie.
    /// 
    /// REPORTING BUGS
    ///        GNU coreutils online help: &lt;https://www.gnu.org/software/coreutils/&gt;
    ///        Report any translation bugs to &lt;https://translationproject.org/team/&gt;
    /// 
    /// COPYRIGHT
    ///        Copyright  ©  2024  Free  Software  Foundation,  Inc.   License  GPLv3+:  GNU  GPL  version  3  or  later  &lt;https://gnu.org/li‐
    ///        censes/gpl.html&gt;.
    ///        This is free software: you are free to change and redistribute it.  There is NO WARRANTY, to the extent permitted by law.
    /// 
    /// SEE ALSO
    ///        Full documentation &lt;https://www.gnu.org/software/coreutils/tee&gt;
    ///        or available locally via: info '(coreutils) tee invocation'
    /// 
    /// GNU coreutils 9.5                                             March 2024                                                        TEE(1)
    /// </summary>
    [BaremetalResourceType("baremetal:index:Tee")]
    public partial class Tee : global::Pulumi.CustomResource
    {
        [Output("create")]
        public Output<Outputs.TeeOpts?> Create { get; private set; } = null!;

        [Output("stderr")]
        public Output<string> Stderr { get; private set; } = null!;

        [Output("stdin")]
        public Output<string> Stdin { get; private set; } = null!;

        [Output("stdout")]
        public Output<string> Stdout { get; private set; } = null!;

        [Output("test")]
        public Output<UnMango.Baremetal.V1alpha1.Outputs.CommandRequest> Test { get; private set; } = null!;


        /// <summary>
        /// Create a Tee resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Tee(string name, TeeArgs args, CustomResourceOptions? options = null)
            : base("baremetal:index:Tee", name, args ?? new TeeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Tee(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("baremetal:index:Tee", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Tee resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Tee Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Tee(name, id, options);
        }
    }

    public sealed class TeeArgs : global::Pulumi.ResourceArgs
    {
        [Input("create")]
        public Input<Inputs.TeeOptsArgs>? Create { get; set; }

        [Input("stdin", required: true)]
        public Input<string> Stdin { get; set; } = null!;

        [Input("test", required: true)]
        public Input<UnMango.Baremetal.V1alpha1.Inputs.CommandRequestArgs> Test { get; set; } = null!;

        public TeeArgs()
        {
        }
        public static new TeeArgs Empty => new TeeArgs();
    }
}
