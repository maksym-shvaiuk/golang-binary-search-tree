package binsrch_tree

// findNode searches for a node with requested value in tree iteratively.
// No recursion to avoid stack overflow on huge trees.
// Returns pointer to node if found, nil otherwise.
// Restrictions: value should not be nil.
func (tree *Tree[T]) findNode(value T) *Node[T] {
	current := tree.root

	for {
		// If reached leaf and didn't find the value - it not exists in tree
		if current == nil {
			return nil
		}

		cmpRes := tree.cmp(value, current.Val)

		// NOTE: value == current
		if cmpRes == 0 {
			return current
		}

		// NOTE: value > current
		if cmpRes > 0 {
			current = current.Right
			continue
		}

		// NOTE: value < current
		if cmpRes < 0 {
			current = current.Left
		}
	}
}
