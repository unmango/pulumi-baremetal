package sdk

import (
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/unmango/pulumi-baremetal/tests/util"
)

func ProgramTestOptions(p util.TestProvisioner) integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		Config: map[string]string{
			"baremetal:address": "localhost",
			"baremetal:port":    "6969",
		},
	}
}
