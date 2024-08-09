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

    public sealed class MkdirArgsArgs : global::Pulumi.ResourceArgs
    {
        [Input("directory", required: true)]
        private InputList<string>? _directory;
        public InputList<string> Directory
        {
            get => _directory ?? (_directory = new InputList<string>());
            set => _directory = value;
        }

        [Input("help")]
        public Input<bool>? Help { get; set; }

        [Input("mode")]
        public Input<string>? Mode { get; set; }

        [Input("parents")]
        public Input<bool>? Parents { get; set; }

        [Input("verbose")]
        public Input<bool>? Verbose { get; set; }

        [Input("version")]
        public Input<bool>? Version { get; set; }

        public MkdirArgsArgs()
        {
        }
        public static new MkdirArgsArgs Empty => new MkdirArgsArgs();
    }
}
