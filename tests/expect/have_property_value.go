package expect

import (
	"fmt"

	"github.com/onsi/gomega"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

type AllowedPropertyValue interface {
	~string | ~int | ~bool
}

type havePropertyValue[T AllowedPropertyValue] struct {
	Value T
}

// Match implements types.GomegaMatcher.
func (e *havePropertyValue[T]) Match(actual interface{}) (success bool, err error) {
	if actual == nil {
		return false, fmt.Errorf("actual was nil")
	}

	value, ok := actual.(resource.PropertyValue)
	if !ok {
		return false, fmt.Errorf("unsupported match target: %T", actual)
	}

	if e.Value != value.V {
		return false, fmt.Errorf("expected %#v to match %#v", value.V, e.Value)
	}

	return true, nil
}

// FailureMessage implements types.GomegaMatcher.
func (e *havePropertyValue[T]) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected property to have value %#v", e.Value)
}

// NegatedFailureMessage implements types.GomegaMatcher.
func (e *havePropertyValue[T]) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected property not to have value %#v", e.Value)
}

func HavePropertyValue[T AllowedPropertyValue](val T) gomega.OmegaMatcher {
	return &havePropertyValue[T]{val}
}
