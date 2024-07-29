import * as pulumi from "@pulumi/pulumi";
import * as baremetal from "@unmango/baremetal";

const tee = new baremetal.cmd.Tee("tee", {
    stdin: "whoops",
    create: {
        files: ["/tmp/tee/test.txt"],
    },
});
export const stdout = tee.stdout;
