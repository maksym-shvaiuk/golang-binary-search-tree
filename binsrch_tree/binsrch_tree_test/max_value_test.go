package binsrch_tree_test

import (
	"errors"
	"testing"

	"github.com/maksym-shvaiuk/golang-binary-search-tree/binsrch_tree"
)

func TestMaxValueBasicScenario(t *testing.T) {
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
	tree := binsrch_tree.New(intCmp)

	// Insert values
	values := []int{10, 5, 15, 3, 7, 12, 18}
	for _, val := range values {
		err := tree.Insert(val)
		if err != nil {
			t.Fatalf("Insert failed for value %d with error: %v", val, err)
		}
	}

	// Test MaxValue
	max, err := tree.MaxValue()
	if err != nil {
		t.Fatalf("MaxValue returned an unexpected error: %v", err)
	}

	expectedMax := 18
	if max != expectedMax {
		t.Fatalf("Expected max value %d, got %d", expectedMax, max)
	}
}

func TestMaxValueEmptyTree(t *testing.T) {
	// Comparison function for integers
	intCmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	// Create an empty tree
	tree := binsrch_tree.New(intCmp)

	// Test MaxValue
	_, err := tree.MaxValue()
	if !errors.Is(err, binsrch_tree.ErrorTreeRootIsNil) {
		t.Fatalf("Expected error ErrorTreeRootIsNil, got %v", err)
	}
}

func TestMaxValueSingleNodeTree(t *testing.T) {
	// Comparison function for integers
	intCmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	// Create a new tree with a single node
	tree := binsrch_tree.New(intCmp)
	err := tree.Insert(42)
	if err != nil {
		t.Fatalf("Insert failed for value 42 with error: %v", err)
	}

	// Test MaxValue
	max, err := tree.MaxValue()
	if err != nil {
		t.Fatalf("MaxValue returned an unexpected error: %v", err)
	}

	expectedMax := 42
	if max != expectedMax {
		t.Fatalf("Expected max value %d, got %d", expectedMax, max)
	}
}

func TestMaxValueSkewedTree(t *testing.T) {
	// Comparison function for integers
	intCmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	// Create a right-skewed tree
	tree := binsrch_tree.New(intCmp)
	values := []int{10, 20, 30, 40, 50}
	for _, val := range values {
		err := tree.Insert(val)
		if err != nil {
			t.Fatalf("Insert failed for value %d with error: %v", val, err)
		}
	}

	// Test MaxValue
	max, err := tree.MaxValue()
	if err != nil {
		t.Fatalf("MaxValue returned an unexpected error: %v", err)
	}

	expectedMax := 50
	if max != expectedMax {
		t.Fatalf("Expected max value %d, got %d", expectedMax, max)
	}
}

func TestMaxValueDuplicateValues1(t *testing.T) {
	// Comparison function for integers
	intCmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	// Create a tree with duplicate values
	tree := binsrch_tree.New(intCmp)
	values := []int{10, 20, 20, 20, 30}
	for _, val := range values {
		err := tree.Insert(val)
		if err != nil {
			t.Fatalf("Insert failed for value %d with error: %v", val, err)
		}
	}

	// Test MaxValue
	max, err := tree.MaxValue()
	if err != nil {
		t.Fatalf("MaxValue returned an unexpected error: %v", err)
	}

	expectedMax := 30
	if max != expectedMax {
		t.Fatalf("Expected max value %d, got %d", expectedMax, max)
	}
}

func TestMaxValueDuplicateValues2(t *testing.T) {
	// Comparison function for integers
	intCmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	// Create a tree with duplicate values
	tree := binsrch_tree.New(intCmp)
	values := []int{10, 20, 80, 20, 20, 30}
	for _, val := range values {
		err := tree.Insert(val)
		if err != nil {
			t.Fatalf("Insert failed for value %d with error: %v", val, err)
		}
	}

	// Test MaxValue
	max, err := tree.MaxValue()
	if err != nil {
		t.Fatalf("MaxValue returned an unexpected error: %v", err)
	}

	expectedMax := 80
	if max != expectedMax {
		t.Fatalf("Expected max value %d, got %d", expectedMax, max)
	}
}

func TestMaxValueDuplicateValues3(t *testing.T) {
	// Comparison function for integers
	intCmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	// Create a tree with duplicate values
	tree := binsrch_tree.New(intCmp)
	values := []int{10, 20, 20, 20, 30, 30, 80}
	for _, val := range values {
		err := tree.Insert(val)
		if err != nil {
			t.Fatalf("Insert failed for value %d with error: %v", val, err)
		}
	}

	// Test MaxValue
	max, err := tree.MaxValue()
	if err != nil {
		t.Fatalf("MaxValue returned an unexpected error: %v", err)
	}

	expectedMax := 80
	if max != expectedMax {
		t.Fatalf("Expected max value %d, got %d", expectedMax, max)
	}
}
