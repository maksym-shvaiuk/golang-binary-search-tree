package golang_binary_tree

// NOTE: Inserts a value recursively.
func (tree *Tree[T]) insertRecursively(value T) {
	tree.root = insertRecursivelyInternal(tree.root, value, tree.cmp)
}

func insertRecursivelyInternal[T any](current *Node[T], value T, cmp func(T, T) int) *Node[T] {
	if current == nil {
		return &Node[T]{
			Val:   value,
			Left:  nil,
			Right: nil,
		}
	}

	cmpRes := cmp(value, current.Val)

	// NOTE: value == current
	if cmpRes == 0 {
		return &Node[T]{
			Val:   value,
			Left:  current,
			Right: nil,
		}
	}

	// NOTE: value > current
	if cmpRes > 0 {
		current.Right = insertRecursivelyInternal(current.Right, value, cmp)
	}

	// NOTE: value < current
	if cmpRes < 0 {
		current.Left = insertRecursivelyInternal(current.Left, value, cmp)
	}

	return current
}
