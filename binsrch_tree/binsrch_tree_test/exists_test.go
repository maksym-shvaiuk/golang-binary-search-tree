package binsrch_tree_test

import (
	"testing"

	"github.com/maksym-shvaiuk/golang-binary-search-tree/binsrch_tree"
)

func TestExistsBasicScenario(t *testing.T) {
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
	tree.SetAlgo(binsrch_tree.AlgoIterative)

	// Insert values into the tree
	values := []int{10, 5, 15, 3, 7, 12, 18}
	for _, val := range values {
		err := tree.Insert(val)
		if err != nil {
			t.Fatalf("Insert failed for value %d with error: %v", val, err)
		}
	}

	// Test for existing values
	for _, val := range values {
		if !tree.Exists(val) {
			t.Fatalf("Expected value %d to exist in the tree, but it was not found", val)
		}
	}

	// Test for non-existing values
	nonExistingValues := []int{1, 6, 20}
	for _, val := range nonExistingValues {
		if tree.Exists(val) {
			t.Fatalf("Expected value %d to not exist in the tree, but it was found", val)
		}
	}
}

func TestExistsEmptyTree(t *testing.T) {
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
	tree.SetAlgo(binsrch_tree.AlgoIterative)

	// Test for any value in an empty tree
	if tree.Exists(10) {
		t.Fatalf("Expected no value to exist in an empty tree, but found a value")
	}
}

func TestExistsSingleNode(t *testing.T) {
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
	tree.SetAlgo(binsrch_tree.AlgoIterative)

	err := tree.Insert(10)
	if err != nil {
		t.Fatalf("Insert failed for value %d with error: %v", 10, err)
	}

	// Test for the existing value
	if !tree.Exists(10) {
		t.Fatalf("Expected value 10 to exist in the tree, but it was not found")
	}

	// Test for a non-existing value
	if tree.Exists(5) {
		t.Fatalf("Expected value 5 to not exist in the tree, but it was found")
	}
}

func TestExistsSkewedTree(t *testing.T) {
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
	tree.SetAlgo(binsrch_tree.AlgoIterative)

	// Create a right-skewed tree
	values := []int{10, 20, 30, 40, 50}
	for _, val := range values {
		err := tree.Insert(val)
		if err != nil {
			t.Fatalf("Insert failed for value %d with error: %v", val, err)
		}
	}

	// Test for existing values
	for _, val := range values {
		if !tree.Exists(val) {
			t.Fatalf("Expected value %d to exist in the tree, but it was not found", val)
		}
	}

	// Test for a non-existing value
	if tree.Exists(5) {
		t.Fatalf("Expected value 5 to not exist in the tree, but it was found")
	}
}

func TestExistsDuplicates(t *testing.T) {
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
	tree.SetAlgo(binsrch_tree.AlgoIterative)

	// Insert values into the tree
	values := []int{10, 5, 15, 3, 7, 15, 12, 5, 18}
	for _, val := range values {
		err := tree.Insert(val)
		if err != nil {
			t.Fatalf("Insert failed for value %d with error: %v", val, err)
		}
	}

	// Test for existing values
	for _, val := range values {
		if !tree.Exists(val) {
			t.Fatalf("Expected value %d to exist in the tree, but it was not found", val)
		}
	}

	// Test for existing values
	existingValues := []int{15, 3, 5}
	for _, val := range existingValues {
		if !tree.Exists(val) {
			t.Fatalf("Expected value %d to exist in the tree, but it was found", val)
		}
	}
}
