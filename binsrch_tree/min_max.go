package binsrch_tree

// getMinSubtreeNode returns the pointer to the
// leftmost (minimal) node of subtree.
func getMinSubtreeNode[T any](root *Node[T]) *Node[T] {
	if root == nil {
		return nil
	}

	for root.Left != nil {
		root = root.Left
	}

	return root
}

// getMaxSubtreeNode returns the pointer to the
// rightmost (maximum) node of subtree.
func getMaxSubtreeNode[T any](root *Node[T]) *Node[T] {
	if root == nil {
		return nil
	}

	for root.Right != nil {
		root = root.Right
	}

	return root
}
