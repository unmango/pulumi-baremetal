// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.Baremetal.Cmd
{
    [BaremetalResourceType("baremetal:cmd:Chmod")]
    public partial class Chmod : global::Pulumi.CustomResource
    {
        [Output("args")]
        public Output<Outputs.ChmodArgs> Args { get; private set; } = null!;

        [Output("createdFiles")]
        public Output<ImmutableArray<string>> CreatedFiles { get; private set; } = null!;

        [Output("exitCode")]
        public Output<int> ExitCode { get; private set; } = null!;

        [Output("movedFiles")]
        public Output<ImmutableDictionary<string, string>> MovedFiles { get; private set; } = null!;

        [Output("stderr")]
        public Output<string> Stderr { get; private set; } = null!;

        [Output("stdout")]
        public Output<string> Stdout { get; private set; } = null!;

        [Output("triggers")]
        public Output<ImmutableArray<object>> Triggers { get; private set; } = null!;


        /// <summary>
        /// Create a Chmod resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Chmod(string name, ChmodArgs args, CustomResourceOptions? options = null)
            : base("baremetal:cmd:Chmod", name, args ?? new ChmodArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Chmod(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("baremetal:cmd:Chmod", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/unmango",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Chmod resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Chmod Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Chmod(name, id, options);
        }
    }

    public sealed class ChmodArgs : global::Pulumi.ResourceArgs
    {
        [Input("args", required: true)]
        public Input<Inputs.ChmodArgsArgs> Args { get; set; } = null!;

        [Input("triggers")]
        private InputList<object>? _triggers;
        public InputList<object> Triggers
        {
            get => _triggers ?? (_triggers = new InputList<object>());
            set => _triggers = value;
        }

        public ChmodArgs()
        {
        }
        public static new ChmodArgs Empty => new ChmodArgs();
    }
}
