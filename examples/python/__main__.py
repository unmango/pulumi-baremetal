import pulumi
import unmango_baremetal as baremetal

tee = baremetal.coreutils.Tee("tee", args={
    "stdin": "whoops",
    "files": ["/tmp/tee-test.txt"],
})
pulumi.export("stdout", tee.stdout)
