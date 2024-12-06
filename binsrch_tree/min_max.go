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

// getMaxSubtreeNodeWithParent returns the pointer to the
// rightmost (maximum) node of subtree, and it's parent node.
// If the root node is the max node in subtree - the parent node returned is nil.
// Returns (*maxNode, *parentNode)
func getMaxSubtreeNodeWithParent[T any](root *Node[T]) (*Node[T], *Node[T]) {
	if root == nil {
		return nil, nil
	}

	maxNode := root
	var parentNode *Node[T] = nil

	for maxNode.Right != nil {
		parentNode = maxNode
		maxNode = maxNode.Right
	}

	return maxNode, parentNode
}
