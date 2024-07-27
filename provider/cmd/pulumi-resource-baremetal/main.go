package main

import (
	p "github.com/pulumi/pulumi-go-provider"

	baremetal "github.com/unmango/pulumi-baremetal/provider"
)

func main() {
	err := p.RunProvider(
		baremetal.Name,
		baremetal.Version,
		baremetal.Provider(),
	)

	if err != nil {
		panic(err)
	}
}
