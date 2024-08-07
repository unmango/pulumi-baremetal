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

    public sealed class RmArgsArgs : global::Pulumi.ResourceArgs
    {
        [Input("dir")]
        public Input<bool>? Dir { get; set; }

        [Input("files", required: true)]
        private InputList<string>? _files;
        public InputList<string> Files
        {
            get => _files ?? (_files = new InputList<string>());
            set => _files = value;
        }

        [Input("force")]
        public Input<bool>? Force { get; set; }

        [Input("help")]
        public Input<bool>? Help { get; set; }

        [Input("oneFileSystem")]
        public Input<bool>? OneFileSystem { get; set; }

        [Input("recursive")]
        public Input<bool>? Recursive { get; set; }

        [Input("verbose")]
        public Input<bool>? Verbose { get; set; }

        public RmArgsArgs()
        {
        }
        public static new RmArgsArgs Empty => new RmArgsArgs();
    }
}
