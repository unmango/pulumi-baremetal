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

    public sealed class TeeArgsArgs : global::Pulumi.ResourceArgs
    {
        [Input("append")]
        public Input<bool>? Append { get; set; }

        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        [Input("files", required: true)]
        private InputList<string>? _files;
        public InputList<string> Files
        {
            get => _files ?? (_files = new InputList<string>());
            set => _files = value;
        }

        public TeeArgsArgs()
        {
        }
        public static new TeeArgsArgs Empty => new TeeArgsArgs();
    }
}