// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package coreutils

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unmango/pulumi-baremetal/sdk/go/baremetal/internal"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "baremetal:coreutils:Cat":
		r = &Cat{}
	case "baremetal:coreutils:Chmod":
		r = &Chmod{}
	case "baremetal:coreutils:Mkdir":
		r = &Mkdir{}
	case "baremetal:coreutils:Mktemp":
		r = &Mktemp{}
	case "baremetal:coreutils:Mv":
		r = &Mv{}
	case "baremetal:coreutils:Rm":
		r = &Rm{}
	case "baremetal:coreutils:Tar":
		r = &Tar{}
	case "baremetal:coreutils:Tee":
		r = &Tee{}
	case "baremetal:coreutils:Wget":
		r = &Wget{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"baremetal",
		"coreutils",
		&module{version},
	)
}
