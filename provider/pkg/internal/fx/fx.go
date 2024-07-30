package fx

import (
	"fmt"

	px "github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

func Sprintf(i string, a ...px.Input[any]) px.Output[string] {
	return px.Apply(px.All(a...), func(a []any) string {
		return fmt.Sprintf(i, a...)
	})
}
