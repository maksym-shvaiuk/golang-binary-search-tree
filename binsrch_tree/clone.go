package binsrch_tree

// clone creates a deep copy of a tree recursively.
func (tree *Tree[T]) clone() Tree[T] {
	clonedTree := New(tree.cmp)
	clonedTree.SetAlgo(tree.algo)

	var cloneNode func(node *Node[T]) *Node[T]
	cloneNode = func(node *Node[T]) *Node[T] {
		if node == nil {
			return nil
		}

		newNode := Node[T]{
			Val:   node.Val,
			Left:  cloneNode(node.Left),
			Right: cloneNode(node.Right),
		}

		return &newNode
	}

	clonedTree.root = cloneNode(tree.GetRoot())
	return clonedTree
}
