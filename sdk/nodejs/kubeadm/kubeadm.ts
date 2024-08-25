// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class Kubeadm extends pulumi.CustomResource {
    /**
     * Get an existing Kubeadm resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Kubeadm {
        return new Kubeadm(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'baremetal:kubeadm:Kubeadm';

    /**
     * Returns true if the given object is an instance of Kubeadm.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Kubeadm {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Kubeadm.__pulumiType;
    }

    public readonly args!: pulumi.Output<outputs.kubeadm.KubeadmArgs>;
    public readonly connection!: pulumi.Output<outputs.config.ProvisionerConnection | undefined>;
    public /*out*/ readonly createdFiles!: pulumi.Output<string[]>;
    public readonly customDelete!: pulumi.Output<string[] | undefined>;
    public readonly customUpdate!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly exitCode!: pulumi.Output<number>;
    public /*out*/ readonly movedFiles!: pulumi.Output<{[key: string]: string}>;
    public /*out*/ readonly stderr!: pulumi.Output<string>;
    public /*out*/ readonly stdout!: pulumi.Output<string>;
    public readonly triggers!: pulumi.Output<any[] | undefined>;

    /**
     * Create a Kubeadm resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KubeadmArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.args === undefined) && !opts.urn) {
                throw new Error("Missing required property 'args'");
            }
            resourceInputs["args"] = args ? args.args : undefined;
            resourceInputs["connection"] = args ? args.connection : undefined;
            resourceInputs["customDelete"] = args ? args.customDelete : undefined;
            resourceInputs["customUpdate"] = args ? args.customUpdate : undefined;
            resourceInputs["triggers"] = args ? args.triggers : undefined;
            resourceInputs["createdFiles"] = undefined /*out*/;
            resourceInputs["exitCode"] = undefined /*out*/;
            resourceInputs["movedFiles"] = undefined /*out*/;
            resourceInputs["stderr"] = undefined /*out*/;
            resourceInputs["stdout"] = undefined /*out*/;
        } else {
            resourceInputs["args"] = undefined /*out*/;
            resourceInputs["connection"] = undefined /*out*/;
            resourceInputs["createdFiles"] = undefined /*out*/;
            resourceInputs["customDelete"] = undefined /*out*/;
            resourceInputs["customUpdate"] = undefined /*out*/;
            resourceInputs["exitCode"] = undefined /*out*/;
            resourceInputs["movedFiles"] = undefined /*out*/;
            resourceInputs["stderr"] = undefined /*out*/;
            resourceInputs["stdout"] = undefined /*out*/;
            resourceInputs["triggers"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Kubeadm.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Kubeadm resource.
 */
export interface KubeadmArgs {
    args: pulumi.Input<inputs.kubeadm.KubeadmArgsArgs>;
    connection?: pulumi.Input<inputs.config.ProvisionerConnectionArgs>;
    customDelete?: pulumi.Input<pulumi.Input<string>[]>;
    customUpdate?: pulumi.Input<pulumi.Input<string>[]>;
    triggers?: pulumi.Input<any[]>;
}
