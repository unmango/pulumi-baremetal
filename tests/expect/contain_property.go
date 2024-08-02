package expect

import (
	"fmt"
	"reflect"

	"github.com/onsi/gomega"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

type containProperty struct {
	Name  resource.PropertyKey
	Value *resource.PropertyValue
}

// Match implements types.GomegaMatcher.
func (e *containProperty) Match(actual interface{}) (success bool, err error) {
	props, ok := actual.(resource.PropertyMap)
	if !ok {
		return false, fmt.Errorf("unsupported match target: %s", reflect.TypeOf(actual))
	}

	out, ok := props[e.Name]
	if !ok {
		return false, fmt.Errorf("%s was not found in the property map", e.Name)
	}

	if e.Value == nil {
		return true, nil
	}

	if *e.Value != out {
		return false, fmt.Errorf("expected %s to match %s", out, e.Value)
	}

	return true, nil
}

// FailureMessage implements types.GomegaMatcher.
func (e *containProperty) FailureMessage(actual interface{}) (message string) {
	switch e.Value {
	case nil:
		return fmt.Sprintf("Expected %s to exist in the property map", e.Name)
	default:
		return fmt.Sprintf("Expected %s to have value %s", e.Name, e.Value)
	}
}

// NegatedFailureMessage implements types.GomegaMatcher.
func (e *containProperty) NegatedFailureMessage(actual interface{}) (message string) {
	switch e.Value {
	case nil:
		return fmt.Sprintf("Expected %s not to exist in the property map", e.Name)
	default:
		return fmt.Sprintf("Expected %s not to have value %s", e.Name, e.Value)
	}
}

func ContainProperty(name resource.PropertyKey) gomega.OmegaMatcher {
	return &containProperty{name, nil}
}

func ContainPropertyValue(name resource.PropertyKey, val resource.PropertyValue) gomega.OmegaMatcher {
	return &containProperty{name, &val}
}
