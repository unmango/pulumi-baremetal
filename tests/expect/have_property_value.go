package expect

import (
	"fmt"

	"github.com/onsi/gomega"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

type havePropertyValue struct {
	Value interface{}
}

// Match implements types.GomegaMatcher.
func (e *havePropertyValue) Match(actual interface{}) (success bool, err error) {
	expected, ok := e.Value.(string)
	if !ok {
		panic("this matcher only supports strings right now")
	}

	if actual == nil {
		return false, fmt.Errorf("actual was nil")
	}

	value, ok := actual.(resource.PropertyValue)
	if !ok {
		return false, fmt.Errorf("unsupported match target: %T", actual)
	}

	if expected != value.V {
		return false, fmt.Errorf("expected %#v to match %#v", value.V, expected)
	}

	return true, nil
}

// FailureMessage implements types.GomegaMatcher.
func (e *havePropertyValue) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected property to have value %#v", e.Value)
}

// NegatedFailureMessage implements types.GomegaMatcher.
func (e *havePropertyValue) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected property not to have value %#v", e.Value)
}

func HavePropertyValue(val interface{}) gomega.OmegaMatcher {
	if _, ok := val.(string); !ok {
		panic("this matcher only supports strings right now")
	}

	return &havePropertyValue{val}
}
