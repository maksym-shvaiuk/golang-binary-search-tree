package binsrch_tree

// deleteRecursively deletes a node with the requested value from the tree recursively.
// Returns true if the node is found and deleted, false otherwise.
// Restrictions: value should not be nil.
func (tree *Tree[T]) deleteRecursively(value T) bool {
	if tree.root == nil {
		return false
	}

	// Use a helper function to delete the node and update the root if necessary
	var deleted bool
	tree.root, deleted = deleteNodeRecursively(tree.root, value, tree.cmp)
	return deleted
}

// deleteNodeRecursively deletes a node from the tree and returns the updated subtree root.
// It also returns a boolean indicating whether the deletion occurred.
func deleteNodeRecursively[T any](node *Node[T], value T, cmp func(a, b T) int) (*Node[T], bool) {
	if node == nil {
		// Base case: Node not found
		return nil, false
	}

	cmpResult := cmp(value, node.Val)
	if cmpResult < 0 {
		// Value is in the left subtree
		var deleted bool
		node.Left, deleted = deleteNodeRecursively(node.Left, value, cmp)
		return node, deleted
	} else if cmpResult > 0 {
		// Value is in the right subtree
		var deleted bool
		node.Right, deleted = deleteNodeRecursively(node.Right, value, cmp)
		return node, deleted
	}

	var numberOfChildren int
	if node.Left != nil {
		numberOfChildren++
	}
	if node.Right != nil {
		numberOfChildren++
	}

	// Case 1: Node has no children (leaf node)
	if numberOfChildren == 0 {
		return nil, true
	}

	// Case 2: Node has one child
	if numberOfChildren == 1 {
		if node.Left != nil {
			return node.Left, true
		}
		return node.Right, true
	}

	// Case 3: Node has two children
	// Find the maximum node in the left subtree
	maxNode, maxNodeParent := getMaxSubtreeNodeWithParent(node.Left)

	// Replace the current node's value with the maximum node's value
	node.Val = maxNode.Val

	// Recursively delete the maximum node from the left subtree
	if maxNodeParent == nil {
		// maxNode is the direct left child of the current node
		node.Left = maxNode.Left
	} else {
		// Update the parent's right pointer to bypass the maxNode
		maxNodeParent.Right = maxNode.Left
	}

	return node, true
}
