// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace UnMango.Baremetal.Inputs
{

    public sealed class ProvisionerConnectionArgs : global::Pulumi.ResourceArgs
    {
        [Input("address", required: true)]
        public Input<string> Address { get; set; } = null!;

        [Input("caPem")]
        public Input<string>? CaPem { get; set; }

        [Input("certPem")]
        public Input<string>? CertPem { get; set; }

        [Input("keyPem")]
        private Input<string>? _keyPem;
        public Input<string>? KeyPem
        {
            get => _keyPem;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _keyPem = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("port")]
        public Input<string>? Port { get; set; }

        public ProvisionerConnectionArgs()
        {
        }
        public static new ProvisionerConnectionArgs Empty => new ProvisionerConnectionArgs();
    }
}