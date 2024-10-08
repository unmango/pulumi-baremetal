// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.Baremetal.Coreutils.Inputs
{

    public sealed class ChmodArgsArgs : global::Pulumi.ResourceArgs
    {
        [Input("changes")]
        public Input<bool>? Changes { get; set; }

        [Input("files", required: true)]
        private InputList<string>? _files;
        public InputList<string> Files
        {
            get => _files ?? (_files = new InputList<string>());
            set => _files = value;
        }

        [Input("help")]
        public Input<bool>? Help { get; set; }

        [Input("mode")]
        private InputList<string>? _mode;
        public InputList<string> Mode
        {
            get => _mode ?? (_mode = new InputList<string>());
            set => _mode = value;
        }

        [Input("noPreserveRoot")]
        public Input<bool>? NoPreserveRoot { get; set; }

        [Input("octalMode")]
        public Input<string>? OctalMode { get; set; }

        [Input("preserveRoot")]
        public Input<bool>? PreserveRoot { get; set; }

        [Input("quiet")]
        public Input<bool>? Quiet { get; set; }

        [Input("recursive")]
        public Input<bool>? Recursive { get; set; }

        [Input("reference")]
        public Input<string>? Reference { get; set; }

        [Input("verbose")]
        public Input<bool>? Verbose { get; set; }

        [Input("version")]
        public Input<bool>? Version { get; set; }

        public ChmodArgsArgs()
        {
        }
        public static new ChmodArgsArgs Empty => new ChmodArgsArgs();
    }
}
