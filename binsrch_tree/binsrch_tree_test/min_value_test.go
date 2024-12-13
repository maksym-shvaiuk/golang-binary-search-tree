package binsrch_tree_test

import (
	"errors"
	"testing"

	"github.com/maksym-shvaiuk/golang-binary-search-tree/binsrch_tree"
)

func TestMinValueBasicScenario(t *testing.T) {
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

	// Test MinValue
	min, err := tree.MinValue()
	if err != nil {
		t.Fatalf("MinValue returned an unexpected error: %v", err)
	}

	expectedMin := 3
	if min != expectedMin {
		t.Fatalf("Expected min value %d, got %d", expectedMin, min)
	}
}

func TestMinValueEmptyTree(t *testing.T) {
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

	// Test MinValue
	_, err := tree.MinValue()
	if !errors.Is(err, binsrch_tree.ErrorTreeRootIsNil) {
		t.Fatalf("Expected error ErrorTreeRootIsNil, got %v", err)
	}
}

func TestMinValueSingleNodeTree(t *testing.T) {
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

	// Test MinValue
	min, err := tree.MinValue()
	if err != nil {
		t.Fatalf("MinValue returned an unexpected error: %v", err)
	}

	expectedMin := 42
	if min != expectedMin {
		t.Fatalf("Expected min value %d, got %d", expectedMin, min)
	}
}

func TestMinValueSkewedTree(t *testing.T) {
	// Comparison function for integers
	intCmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	// Create a left-skewed tree
	tree := binsrch_tree.New(intCmp)
	values := []int{50, 40, 30, 20, 10}
	for _, val := range values {
		err := tree.Insert(val)
		if err != nil {
			t.Fatalf("Insert failed for value %d with error: %v", val, err)
		}
	}

	// Test MinValue
	min, err := tree.MinValue()
	if err != nil {
		t.Fatalf("MinValue returned an unexpected error: %v", err)
	}

	expectedMin := 10
	if min != expectedMin {
		t.Fatalf("Expected min value %d, got %d", expectedMin, min)
	}
}

func TestMinValueDuplicateValues1(t *testing.T) {
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
	values := []int{30, 20, 20, 10, 10}
	for _, val := range values {
		err := tree.Insert(val)
		if err != nil {
			t.Fatalf("Insert failed for value %d with error: %v", val, err)
		}
	}

	// Test MinValue
	min, err := tree.MinValue()
	if err != nil {
		t.Fatalf("MinValue returned an unexpected error: %v", err)
	}

	expectedMin := 10
	if min != expectedMin {
		t.Fatalf("Expected min value %d, got %d", expectedMin, min)
	}
}

func TestMinValueDuplicateValues2(t *testing.T) {
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
	values := []int{30, 20, 8, 20, 10, 10}
	for _, val := range values {
		err := tree.Insert(val)
		if err != nil {
			t.Fatalf("Insert failed for value %d with error: %v", val, err)
		}
	}

	// Test MinValue
	min, err := tree.MinValue()
	if err != nil {
		t.Fatalf("MinValue returned an unexpected error: %v", err)
	}

	expectedMin := 8
	if min != expectedMin {
		t.Fatalf("Expected min value %d, got %d", expectedMin, min)
	}
}

func TestMinValueDuplicateValues3(t *testing.T) {
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
	values := []int{30, 20, 20, 20, 10, 8}
	for _, val := range values {
		err := tree.Insert(val)
		if err != nil {
			t.Fatalf("Insert failed for value %d with error: %v", val, err)
		}
	}

	// Test MinValue
	min, err := tree.MinValue()
	if err != nil {
		t.Fatalf("MinValue returned an unexpected error: %v", err)
	}

	expectedMin := 8
	if min != expectedMin {
		t.Fatalf("Expected min value %d, got %d", expectedMin, min)
	}
}
