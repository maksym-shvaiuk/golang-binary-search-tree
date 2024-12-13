package binsrch_tree_test

import (
	"testing"

	"github.com/maksym-shvaiuk/golang-binary-search-tree/binsrch_tree"
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
	tree := binsrch_tree.New(intCmp)
	tree.SetAlgo(binsrch_tree.AlgoRecursive)

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

	tree := binsrch_tree.New(intCmp)
	tree.SetAlgo(binsrch_tree.AlgoRecursive)
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

	tree := binsrch_tree.New(intCmp)
	tree.SetAlgo(binsrch_tree.AlgoRecursive)
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

	tree := binsrch_tree.New(intCmp)
	tree.SetAlgo(binsrch_tree.AlgoRecursive)
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

	tree := binsrch_tree.New(intCmp)
	tree.SetAlgo(binsrch_tree.AlgoRecursive)
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

	tree := binsrch_tree.New(intCmp)
	tree.SetAlgo(binsrch_tree.AlgoRecursive)
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

	tree := binsrch_tree.New(intCmp)
	tree.SetAlgo(binsrch_tree.AlgoRecursive)
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

// TestInsertRecursiveMultipleValues tests multiple deletions in sequence.
func TestInsertRecursiveMultipleValues(t *testing.T) {
	intCmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	}

	tree := binsrch_tree.New(intCmp)
	tree.SetAlgo(binsrch_tree.AlgoRecursive)

	// Insert nodes into the tree
	_ = tree.Insert(62)
	_ = tree.Insert(83)
	_ = tree.Insert(33)
	_ = tree.Insert(92)
	_ = tree.Insert(18)
	_ = tree.Insert(83)
	_ = tree.Insert(87)
	_ = tree.Insert(46)
	_ = tree.Insert(17)
	_ = tree.Insert(41)

	// Verify remaining structure
	root := tree.GetRoot()
	if root == nil || root.Val != 62 {
		t.Fatalf("Expected root value to be 62, got: %v", root)
	}

	// Left subtree
	if root.Left == nil || root.Left.Val != 33 {
		t.Fatalf("Expected left child of root to be 33, got: %v", root.Left)
	}
	if root.Left.Left == nil || root.Left.Left.Val != 18 {
		t.Fatalf("Expected left child of node 33 to be 18, got: %v", root.Left.Left)
	}
	if root.Left.Left.Left == nil || root.Left.Left.Left.Val != 17 {
		t.Fatalf("Expected left child of node 18 to be 17, got: %v", root.Left.Left.Left)
	}
	if root.Left.Left.Right != nil {
		t.Fatalf("Expected right child of node 18 to be [null], got: %v", root.Left.Left.Right)
	}
	if root.Left.Right == nil || root.Left.Right.Val != 46 {
		t.Fatalf("Expected right child of node 33 to be 46, got: %v", root.Left.Right)
	}
	if root.Left.Right.Left == nil || root.Left.Right.Left.Val != 41 {
		t.Fatalf("Expected left child of node 46 to be 41, got: %v", root.Left.Right.Left)
	}
	if root.Left.Right.Right != nil {
		t.Fatalf("Expected right child of node 46 to be [null], got: %v", root.Left.Right.Right)
	}

	// Right subtree
	if root.Right == nil || root.Right.Val != 83 {
		t.Fatalf("Expected right child of root to be 83, got: %v", root.Right)
	}
	if root.Right.Left == nil || root.Right.Left.Val != 83 {
		t.Fatalf("Expected left child of node 83 to be 83, got: %v", root.Right.Left)
	}
	if root.Right.Right == nil || root.Right.Right.Val != 92 {
		t.Fatalf("Expected right child of node 83 to be 92, got: %v", root.Right.Right)
	}
}
