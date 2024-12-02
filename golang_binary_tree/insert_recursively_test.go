package golang_binary_tree_test

import (
	"testing"

	"github.com/maksym-shvaiuk/golang-binary-search-tree/golang_binary_tree"
)

func TestInsertRecursivelyBasicScenario(t *testing.T) {
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
	tree.SetAlgo(golang_binary_tree.AlgoRecursive)

	// Insert a single value
	err := tree.Insert(5)
	if err != nil {
		t.Fatalf("Insert failed with error: %v", err)
	}

	// Verify that the root is set correctly
	if tree.GetRoot() == nil || tree.GetRoot().Val != 5 {
		t.Fatalf("Expected root value to be 5, got %v", tree.GetRoot())
	}

	// Insert another value
	err = tree.Insert(3)
	if err != nil {
		t.Fatalf("Insert failed with error: %v", err)
	}

	// Verify that the left child is set correctly
	if tree.GetRoot().Left == nil || tree.GetRoot().Left.Val != 3 {
		t.Fatalf("Expected left child value to be 3, got %v", tree.GetRoot().Left)
	}

	// Insert a value greater than root
	err = tree.Insert(7)
	if err != nil {
		t.Fatalf("Insert failed with error: %v", err)
	}

	// Verify that the right child is set correctly
	if tree.GetRoot().Right == nil || tree.GetRoot().Right.Val != 7 {
		t.Fatalf("Expected right child value to be 7, got %v", tree.GetRoot().Right)
	}

	tree.PrintTree(&TestLogger{})
}

// TestInsertRecursivelyIntoEmptyTree checks insertion into an empty tree.
func TestInsertRecursivelyIntoEmptyTree(t *testing.T) {
	intCmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	tree := golang_binary_tree.New(intCmp)
	tree.SetAlgo(golang_binary_tree.AlgoRecursive)
	err := tree.Insert(10)
	if err != nil {
		t.Fatalf("Insert failed with error: %v", err)
	}
	if tree.GetRoot() == nil || tree.GetRoot().Val != 10 {
		t.Fatalf("Expected root value to be 10, got %v", tree.GetRoot())
	}
}

// TestInsertRecursivelyDuplicateValues checks insertion of duplicate values.
func TestInsertRecursivelyDuplicateValues(t *testing.T) {
	intCmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	tree := golang_binary_tree.New(intCmp)
	tree.SetAlgo(golang_binary_tree.AlgoRecursive)
	_ = tree.Insert(10)
	err := tree.Insert(10)
	if err != nil {
		t.Fatalf("Insert failed with error: %v", err)
	}
	if tree.GetRoot().Left == nil || tree.GetRoot().Left.Val != 10 {
		t.Fatalf("Expected left child value to be duplicate 10, got %v", tree.GetRoot().Left)
	}
}

// TestInsertRecursivelyLargerValues checks insertion of values larger than the root.
func TestInsertRecursivelyLargerValues(t *testing.T) {
	intCmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	tree := golang_binary_tree.New(intCmp)
	tree.SetAlgo(golang_binary_tree.AlgoRecursive)
	_ = tree.Insert(10)
	err := tree.Insert(20)
	if err != nil {
		t.Fatalf("Insert failed with error: %v", err)
	}
	if tree.GetRoot().Right == nil || tree.GetRoot().Right.Val != 20 {
		t.Fatalf("Expected right child value to be 20, got %v", tree.GetRoot().Right)
	}
}

// TestInsertRecursivelySmallerValues checks insertion of values smaller than the root.
func TestInsertRecursivelySmallerValues(t *testing.T) {
	intCmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	tree := golang_binary_tree.New(intCmp)
	tree.SetAlgo(golang_binary_tree.AlgoRecursive)
	_ = tree.Insert(10)
	err := tree.Insert(5)
	if err != nil {
		t.Fatalf("Insert failed with error: %v", err)
	}
	if tree.GetRoot().Left == nil || tree.GetRoot().Left.Val != 5 {
		t.Fatalf("Expected left child value to be 5, got %v", tree.GetRoot().Left)
	}
}

// TestInsertRecursivelyMultipleValues checks insertion of multiple values to form a balanced tree.
func TestInsertRecursivelyMultipleValues(t *testing.T) {
	intCmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	tree := golang_binary_tree.New(intCmp)
	tree.SetAlgo(golang_binary_tree.AlgoRecursive)
	values := []int{10, 5, 15, 3, 7, 12, 18}
	for _, val := range values {
		if err := tree.Insert(val); err != nil {
			t.Fatalf("Insert failed for value %d with error: %v", val, err)
		}
	}

	// Verify structure
	if tree.GetRoot().Val != 10 {
		t.Fatalf("Expected root value to be 10, got %v", tree.GetRoot().Val)
	}
	if tree.GetRoot().Left.Val != 5 || tree.GetRoot().Right.Val != 15 {
		t.Fatalf("Expected left child 5 and right child 15, got %v and %v", tree.GetRoot().Left, tree.GetRoot().Right)
	}
}

// TestInsertRecursivelyUnbalancedTree checks the behavior with a deeply unbalanced tree.
func TestInsertRecursivelyUnbalancedTree(t *testing.T) {
	intCmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	tree := golang_binary_tree.New(intCmp)
	tree.SetAlgo(golang_binary_tree.AlgoRecursive)
	values := []int{1, 2, 3, 4, 5}
	for _, val := range values {
		if err := tree.Insert(val); err != nil {
			t.Fatalf("Insert failed for value %d with error: %v", val, err)
		}
	}

	// Verify the tree is right-skewed
	current := tree.GetRoot()
	expected := 1
	for current != nil {
		if current.Val != expected {
			t.Fatalf("Expected node value to be %d, got %v", expected, current.Val)
		}
		current = current.Right
		expected++
	}
}

// TestInsertRecursivelyDataFromVideoTutorial checks the behavior with a data from the video tutorial (54).
func TestInsertRecursivelyDataFromVideoTutorial(t *testing.T) {
	intCmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	tree := golang_binary_tree.New(intCmp)
	tree.SetAlgo(golang_binary_tree.AlgoRecursive)
	values := []int{60, 41, 16, 53, 46, 55, 42, 54}
	for _, val := range values {
		if err := tree.Insert(val); err != nil {
			t.Fatalf("Insert failed for value %d with error: %v", val, err)
		}
	}

	// Verify the tree is right-skewed
	current := tree.GetRoot()
	expected := 60
	for current != nil {
		if current.Val != expected {
			t.Fatalf("Expected node value to be %d, got %v", expected, current.Val)
		}
		current = current.Right
		expected++
	}

	tree.PrintTree(&TestLogger{})
}
