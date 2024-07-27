import * as pulumi from "@pulumi/pulumi";
import * as baremetal from "@unmango/baremetal";

const myRandomResource = new baremetal.Random("myRandomResource", {length: 24});
export const output = {
    value: myRandomResource.result,
};
