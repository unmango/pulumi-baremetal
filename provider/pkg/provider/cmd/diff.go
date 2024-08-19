package cmd

import (
	"context"
	"slices"

	provider "github.com/pulumi/pulumi-go-provider"
)

func (s *State[T]) Diff(ctx context.Context, inputs CommandArgs[T]) (map[string]provider.PropertyDiff, error) {
	diff := map[string]provider.PropertyDiff{}
	if !slices.Equal(s.Triggers, inputs.Triggers) {
		diff["triggers"] = provider.PropertyDiff{Kind: provider.UpdateReplace}
	}

	if !slices.Equal(s.CustomUpdate, inputs.CustomUpdate) {
		diff["customUpdate"] = provider.PropertyDiff{Kind: provider.Update}
	}

	if len(inputs.CustomDelete) > len(s.CustomDelete) {
		diff["customDelete"] = provider.PropertyDiff{Kind: provider.Add}
	}

	return diff, nil
}

func Changed[T comparable](a *T, b *T) bool {
	if a == b {
		return false
	}

	if a == nil || b == nil {
		return true
	}

	return *a != *b
}
