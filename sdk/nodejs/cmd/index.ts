// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { MvArgs } from "./mv";
export type Mv = import("./mv").Mv;
export const Mv: typeof import("./mv").Mv = null as any;
utilities.lazyLoad(exports, ["Mv"], () => require("./mv"));

export { RmArgs } from "./rm";
export type Rm = import("./rm").Rm;
export const Rm: typeof import("./rm").Rm = null as any;
utilities.lazyLoad(exports, ["Rm"], () => require("./rm"));

export { TarArgs } from "./tar";
export type Tar = import("./tar").Tar;
export const Tar: typeof import("./tar").Tar = null as any;
utilities.lazyLoad(exports, ["Tar"], () => require("./tar"));

export { TeeArgs } from "./tee";
export type Tee = import("./tee").Tee;
export const Tee: typeof import("./tee").Tee = null as any;
utilities.lazyLoad(exports, ["Tee"], () => require("./tee"));

export { WgetArgs } from "./wget";
export type Wget = import("./wget").Wget;
export const Wget: typeof import("./wget").Wget = null as any;
utilities.lazyLoad(exports, ["Wget"], () => require("./wget"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "baremetal:cmd:Mv":
                return new Mv(name, <any>undefined, { urn })
            case "baremetal:cmd:Rm":
                return new Rm(name, <any>undefined, { urn })
            case "baremetal:cmd:Tar":
                return new Tar(name, <any>undefined, { urn })
            case "baremetal:cmd:Tee":
                return new Tee(name, <any>undefined, { urn })
            case "baremetal:cmd:Wget":
                return new Wget(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("baremetal", "cmd", _module)
