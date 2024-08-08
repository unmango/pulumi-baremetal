package expect

import (
	"fmt"

	"github.com/onsi/gomega"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

type haveArrayValue[T AllowedPropertyValue] struct {
	Value []T
}

// Match implements types.GomegaMatcher.
func (e *haveArrayValue[T]) Match(actual interface{}) (success bool, err error) {
	if actual == nil {
		return false, fmt.Errorf("actual was nil")
	}

	value, ok := actual.(resource.PropertyValue)
	if !ok {
		return false, fmt.Errorf("unsupported match target: %T", actual)
	}

	return gomega.BeEquivalentTo(resource.NewPropertyValue(e.Value)).Match(value.ArrayValue())
}

// FailureMessage implements types.GomegaMatcher.
func (e *haveArrayValue[T]) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected %#v to have value %#v", actual, e.Value)
}

// NegatedFailureMessage implements types.GomegaMatcher.
func (e *haveArrayValue[T]) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected %#v not to have value %#v", actual, e.Value)
}

func HaveArrayValue[T AllowedPropertyValue](val []T) gomega.OmegaMatcher {
	return &haveArrayValue[T]{val}
}
