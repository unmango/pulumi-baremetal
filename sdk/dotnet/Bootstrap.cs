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
    [BaremetalResourceType("baremetal:index:Bootstrap")]
    public partial class Bootstrap : global::Pulumi.ComponentResource
    {
        [Output("download")]
        public Output<Pulumi.Command.Remote.Command> Download { get; private set; } = null!;

        [Output("mktemp")]
        public Output<Pulumi.Command.Remote.Command> Mktemp { get; private set; } = null!;

        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a Bootstrap resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Bootstrap(string name, BootstrapArgs args, ComponentResourceOptions? options = null)
            : base("baremetal:index:Bootstrap", name, args ?? new BootstrapArgs(), MakeResourceOptions(options, ""), remote: true)
        {
        }

        private static ComponentResourceOptions MakeResourceOptions(ComponentResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new ComponentResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/unmango",
            };
            var merged = ComponentResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class BootstrapArgs : global::Pulumi.ResourceArgs
    {
        [Input("connection", required: true)]
        public Input<Pulumi.Command.Remote.Inputs.ConnectionArgs> Connection { get; set; } = null!;

        [Input("version", required: true)]
        public string Version { get; set; } = null!;

        public BootstrapArgs()
        {
        }
        public static new BootstrapArgs Empty => new BootstrapArgs();
    }
}
