// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.Baremetal.Cmd.Inputs
{

    public sealed class MktempArgsArgs : global::Pulumi.ResourceArgs
    {
        [Input("directory")]
        public Input<bool>? Directory { get; set; }

        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        [Input("help")]
        public Input<bool>? Help { get; set; }

        [Input("p")]
        public Input<string>? P { get; set; }

        [Input("quiet")]
        public Input<bool>? Quiet { get; set; }

        [Input("suffix")]
        public Input<string>? Suffix { get; set; }

        [Input("t")]
        public Input<bool>? T { get; set; }

        [Input("template")]
        public Input<string>? Template { get; set; }

        [Input("tmpdir")]
        public Input<bool>? Tmpdir { get; set; }

        [Input("version")]
        public Input<bool>? Version { get; set; }

        public MktempArgsArgs()
        {
        }
        public static new MktempArgsArgs Empty => new MktempArgsArgs();
    }
}
