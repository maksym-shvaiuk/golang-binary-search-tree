package golang_binary_tree_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/maksym-shvaiuk/golang-binary-search-tree/golang_binary_tree"
)

const printTreesInUnitTests = true

type TestLogger struct {
}

func (l *TestLogger) Print(str string) {
	if !printTreesInUnitTests {
		return
	}
	fmt.Println(str)
}

func TestInsertBasicScenario(t *testing.T) {
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
	tree.SetAlgo(golang_binary_tree.AlgoIterative)

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

// TestInsertIntoEmptyTree checks insertion into an empty tree.
func TestInsertIntoEmptyTree(t *testing.T) {
	intCmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	tree := golang_binary_tree.New(intCmp)
	tree.SetAlgo(golang_binary_tree.AlgoIterative)
	err := tree.Insert(10)
	if err != nil {
		t.Fatalf("Insert failed with error: %v", err)
	}
	if tree.GetRoot() == nil || tree.GetRoot().Val != 10 {
		t.Fatalf("Expected root value to be 10, got %v", tree.GetRoot())
	}
}

// TestInsertDuplicateValues checks insertion of duplicate values.
func TestInsertDuplicateValues(t *testing.T) {
	intCmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	tree := golang_binary_tree.New(intCmp)
	tree.SetAlgo(golang_binary_tree.AlgoIterative)
	_ = tree.Insert(10)
	err := tree.Insert(10)
	if err != nil {
		t.Fatalf("Insert failed with error: %v", err)
	}
	if tree.GetRoot().Left == nil || tree.GetRoot().Left.Val != 10 {
		t.Fatalf("Expected left child value to be duplicate 10, got %v", tree.GetRoot().Left)
	}
}

// TestInsertLargerValues checks insertion of values larger than the root.
func TestInsertLargerValues(t *testing.T) {
	intCmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	tree := golang_binary_tree.New(intCmp)
	tree.SetAlgo(golang_binary_tree.AlgoIterative)
	_ = tree.Insert(10)
	err := tree.Insert(20)
	if err != nil {
		t.Fatalf("Insert failed with error: %v", err)
	}
	if tree.GetRoot().Right == nil || tree.GetRoot().Right.Val != 20 {
		t.Fatalf("Expected right child value to be 20, got %v", tree.GetRoot().Right)
	}
}

// TestInsertSmallerValues checks insertion of values smaller than the root.
func TestInsertSmallerValues(t *testing.T) {
	intCmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	tree := golang_binary_tree.New(intCmp)
	tree.SetAlgo(golang_binary_tree.AlgoIterative)
	_ = tree.Insert(10)
	err := tree.Insert(5)
	if err != nil {
		t.Fatalf("Insert failed with error: %v", err)
	}
	if tree.GetRoot().Left == nil || tree.GetRoot().Left.Val != 5 {
		t.Fatalf("Expected left child value to be 5, got %v", tree.GetRoot().Left)
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

// TestInsertMultipleValues checks insertion of multiple values to form a balanced tree.
func TestInsertMultipleValues(t *testing.T) {
	intCmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	tree := golang_binary_tree.New(intCmp)
	tree.SetAlgo(golang_binary_tree.AlgoIterative)
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

// TestUnbalancedTree checks the behavior with a deeply unbalanced tree.
func TestUnbalancedTree(t *testing.T) {
	intCmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	tree := golang_binary_tree.New(intCmp)
	tree.SetAlgo(golang_binary_tree.AlgoIterative)
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
