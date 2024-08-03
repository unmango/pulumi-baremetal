// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'baremetal';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === "pulumi:providers:" + Provider.__pulumiType;
    }

    public readonly address!: pulumi.Output<string>;
    public readonly caPem!: pulumi.Output<string | undefined>;
    public readonly certPem!: pulumi.Output<string | undefined>;
    public readonly keyPem!: pulumi.Output<string | undefined>;
    public readonly port!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            if ((!args || args.address === undefined) && !opts.urn) {
                throw new Error("Missing required property 'address'");
            }
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["caPem"] = args ? args.caPem : undefined;
            resourceInputs["certPem"] = args ? args.certPem : undefined;
            resourceInputs["keyPem"] = args ? args.keyPem : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    address: pulumi.Input<string>;
    caPem?: pulumi.Input<string>;
    certPem?: pulumi.Input<string>;
    keyPem?: pulumi.Input<string>;
    port?: pulumi.Input<string>;
}
