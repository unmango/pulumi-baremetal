package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unmango/pulumi-baremetal/sdk/go/baremetal"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myRandomResource, err := baremetal.NewRandom(ctx, "myRandomResource", &baremetal.RandomArgs{
			Length: pulumi.Int(24),
		})
		if err != nil {
			return err
		}
		ctx.Export("output", map[string]interface{}{
			"value": myRandomResource.Result,
		})
		return nil
	})
}
