package binsrch_tree_test

import (
	"testing"

	"github.com/maksym-shvaiuk/golang-binary-search-tree/binsrch_tree"
)

// TestDeleteRecursivelyBasicScenario tests deletion in a tree with a single value.
func TestDeleteRecursivelyBasicScenario(t *testing.T) {
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

	// Delete the single value
	deleted, err := tree.Delete(10)
	if err != nil {
		t.Fatalf("Delete failed with error: %v", err)
	}
	if !deleted {
		t.Fatalf("Expected value 10 to be deleted, but it was not")
	}
	if tree.GetRoot() != nil {
		t.Fatalf("Expected tree to be empty after deletion, but got root: %v", tree.GetRoot())
	}
}

// TestDeleteRecursivelyFromEmptyTree tests deletion from an empty tree.
func TestDeleteRecursivelyFromEmptyTree(t *testing.T) {
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

	// Try to delete from an empty tree
	deleted, err := tree.Delete(10)
	if err != nil {
		t.Fatalf("Delete failed with error: %v", err)
	}
	if deleted {
		t.Fatalf("Expected deletion to return false for an empty tree, but it returned true")
	}
}

// TestDeleteRecursivelyLeafNode tests deleting a leaf node.
func TestDeleteRecursivelyLeafNode(t *testing.T) {
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
	_ = tree.Insert(5)
	_ = tree.Insert(15)

	// Delete a leaf node
	deleted, err := tree.Delete(5)
	if err != nil {
		t.Fatalf("Delete failed with error: %v", err)
	}
	if !deleted {
		t.Fatalf("Expected value 5 to be deleted, but it was not")
	}
	if tree.GetRoot().Left != nil {
		t.Fatalf("Expected left child to be nil after deletion, but got: %v", tree.GetRoot().Left)
	}
}

// TestDeleteRecursivelyNodeWithOneChild tests deleting a node with one child.
func TestDeleteRecursivelyNodeWithOneChild(t *testing.T) {
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
	_ = tree.Insert(5)
	_ = tree.Insert(15)
	_ = tree.Insert(20)

	// Delete a node with one child
	deleted, err := tree.Delete(15)
	if err != nil {
		t.Fatalf("Delete failed with error: %v", err)
	}
	if !deleted {
		t.Fatalf("Expected value 15 to be deleted, but it was not")
	}
	if tree.GetRoot().Right == nil || tree.GetRoot().Right.Val != 20 {
		t.Fatalf("Expected right child of root to be 20 after deletion, but got: %v", tree.GetRoot().Right)
	}
}

// TestDeleteRecursivelyNodeWithTwoChildren tests deleting a node with two children.
func TestDeleteRecursivelyNodeWithTwoChildren(t *testing.T) {
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
	_ = tree.Insert(5)
	_ = tree.Insert(15)
	_ = tree.Insert(3)
	_ = tree.Insert(7)
	_ = tree.Insert(12)
	_ = tree.Insert(18)

	// Delete a node with two children
	deleted, err := tree.Delete(10)
	if err != nil {
		t.Fatalf("Delete failed with error: %v", err)
	}
	if !deleted {
		t.Fatalf("Expected value 10 to be deleted, but it was not")
	}
	if tree.GetRoot() == nil || tree.GetRoot().Val != 7 {
		t.Fatalf("Expected root value to be replaced by 7, got: %v", tree.GetRoot())
	}
}

// TestDeleteRecursivelyNonExistentNode tests deleting a value not in the tree.
func TestDeleteRecursivelyNonExistentNode(t *testing.T) {
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
	_ = tree.Insert(5)
	_ = tree.Insert(15)

	// Try to delete a value not in the tree
	deleted, err := tree.Delete(20)
	if err != nil {
		t.Fatalf("Delete failed with error: %v", err)
	}
	if deleted {
		t.Fatalf("Expected deletion to return false for non-existent value, but it returned true")
	}
}

// TestDeleteRecursiveMultipleValues tests multiple deletions in sequence.
func TestDeleteRecursiveMultipleValues(t *testing.T) {
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
	values := []int{62, 83, 33, 92, 18, 83, 87, 46, 17, 41}
	for _, val := range values {
		_ = tree.Insert(val)
	}

	// Use constant seed to get the same result each test run
	valuesToDelete := shuffleDataset(values, 123)

	// Delete all values
	for _, val := range valuesToDelete {
		deleted, err := tree.Delete(val)
		if err != nil {
			t.Fatalf("Delete failed for value %d with error: %v", val, err)
		}
		if !deleted {
			t.Fatalf("Expected value %d to be deleted, but it was not", val)
		}
	}

	// Verify remaining structure
	root := tree.GetRoot()
	if root != nil {
		t.Fatalf("Expected root value to be nil")
	}
}
