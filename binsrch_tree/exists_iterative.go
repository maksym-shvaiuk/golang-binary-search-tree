package binsrch_tree

// NOTE: Search a node with requested value in tree iteratively.
// No recursion to avoid stack overflow on huge trees.
// Returns true if node found, false otherwise.
// Restrictions: value should not be nil.
func (tree *Tree[T]) existsIterative(value T) bool {
	current := tree.root

	for {
		// If reached leaf and didn't find the value - it not exists in tree
		if current == nil {
			return false
		}

		cmpRes := tree.cmp(value, current.Val)

		// NOTE: value == current
		if cmpRes == 0 {
			return true
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
