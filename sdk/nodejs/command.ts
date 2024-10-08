// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class Command extends pulumi.CustomResource {
    /**
     * Get an existing Command resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Command {
        return new Command(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'baremetal:index:Command';

    /**
     * Returns true if the given object is an instance of Command.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Command {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Command.__pulumiType;
    }

    public readonly connection!: pulumi.Output<outputs.ProvisionerConnection | undefined>;
    public readonly create!: pulumi.Output<string[]>;
    public readonly delete!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly exitCode!: pulumi.Output<number>;
    public /*out*/ readonly stderr!: pulumi.Output<string>;
    public /*out*/ readonly stdout!: pulumi.Output<string>;
    public readonly triggers!: pulumi.Output<any[] | undefined>;
    public readonly update!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Command resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CommandArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.create === undefined) && !opts.urn) {
                throw new Error("Missing required property 'create'");
            }
            resourceInputs["connection"] = args ? args.connection : undefined;
            resourceInputs["create"] = args ? args.create : undefined;
            resourceInputs["delete"] = args ? args.delete : undefined;
            resourceInputs["triggers"] = args ? args.triggers : undefined;
            resourceInputs["update"] = args ? args.update : undefined;
            resourceInputs["exitCode"] = undefined /*out*/;
            resourceInputs["stderr"] = undefined /*out*/;
            resourceInputs["stdout"] = undefined /*out*/;
        } else {
            resourceInputs["connection"] = undefined /*out*/;
            resourceInputs["create"] = undefined /*out*/;
            resourceInputs["delete"] = undefined /*out*/;
            resourceInputs["exitCode"] = undefined /*out*/;
            resourceInputs["stderr"] = undefined /*out*/;
            resourceInputs["stdout"] = undefined /*out*/;
            resourceInputs["triggers"] = undefined /*out*/;
            resourceInputs["update"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Command.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Command resource.
 */
export interface CommandArgs {
    connection?: pulumi.Input<inputs.ProvisionerConnectionArgs>;
    create: pulumi.Input<pulumi.Input<string>[]>;
    delete?: pulumi.Input<pulumi.Input<string>[]>;
    triggers?: pulumi.Input<any[]>;
    update?: pulumi.Input<pulumi.Input<string>[]>;
}
