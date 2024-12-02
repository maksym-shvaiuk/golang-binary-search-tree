package golang_binary_tree_test

import (
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
