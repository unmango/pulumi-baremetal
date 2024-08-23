import * as pulumi from "@pulumi/pulumi";
import * as baremetal from "@unmango/baremetal";

const tee = new baremetal.coreutils.Tee("tee", {args: {
    stdin: "whoops",
    files: ["/tmp/tee-test.txt"],
}});
export const stdout = tee.stdout;
