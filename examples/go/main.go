package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unmango/pulumi-baremetal/sdk/go/baremetal/coreutils"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tee, err := coreutils.NewTee(ctx, "tee", &coreutils.TeeArgs{
			Args: &coreutils.TeeArgsTypeArgs{
				Stdin: pulumi.String("whoops"),
				Files: pulumi.StringArray{
					pulumi.String("/tmp/tee-test.txt"),
				},
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("stdout", tee.Stdout)
		return nil
	})
}
