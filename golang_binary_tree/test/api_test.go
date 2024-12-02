package golang_binary_tree_test

import (
	"errors"
	"testing"

	"github.com/maksym-shvaiuk/golang-binary-search-tree/golang_binary_tree"
)

func TestNewTreeInitialization(t *testing.T) {
	// Comparison function for integers
	intCmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	// Create a new tree
	tree := golang_binary_tree.New(intCmp)

	// Test if tree.GetRoot() is initialized to nil
	if tree.GetRoot() != nil {
		t.Fatalf("Expected tree.root to be nil, got %v", tree.GetRoot())
	}
}

// TestNilComparisonFunction checks handling of a nil comparison function.
func TestNilComparisonFunction(t *testing.T) {
	tree := golang_binary_tree.New[int](nil)
	err := tree.Insert(5)
	if !errors.Is(err, golang_binary_tree.ErrorCmpFunctionIsNil) {
		t.Fatalf("Expected error for nil comparison function, got %v", err)
	}
}

// TestUnalteredTreeAfterError ensures the tree remains unaltered after a failed insert.
func TestUnalteredTreeAfterError(t *testing.T) {
	tree := golang_binary_tree.New[int](nil)
	_ = tree.Insert(10) // Should fail silently
	if tree.GetRoot() != nil {
		t.Fatalf("Expected tree.root to remain nil after failed insert, got %v", tree.GetRoot())
	}
}
