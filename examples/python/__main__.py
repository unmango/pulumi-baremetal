import pulumi
import unmango_baremetal as baremetal

tee = baremetal.cmd.Tee("tee",
    stdin="whoops",
    create={
        "files": ["/tmp/tee/test.txt"],
    })
pulumi.export("output", {
    "value": tee.stdout,
})
