package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unmango/pulumi-baremetal/sdk/go/baremetal/cmd"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tee, err := cmd.NewTee(ctx, "tee", &cmd.TeeArgs{
			Stdin: pulumi.String("whoops"),
			Create: &cmd.TeeOptsArgs{
				Files: pulumi.StringArray{
					pulumi.String("/tmp/tee/test.txt"),
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
