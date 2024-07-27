import pulumi
import pulumi_baremetal as baremetal

my_random_resource = baremetal.Random("myRandomResource", length=24)
pulumi.export("output", {
    "value": my_random_resource.result,
})
