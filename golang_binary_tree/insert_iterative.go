package golang_binary_tree

// NOTE: Inserts a value iteratively.
// No recursion to avoid stack overflow on huge trees.
// Duplicate values are handled by inserting a new node to the left subtree.
// Restrictions: value should not be nil.
func (tree *Tree[T]) insertIterative(value T) {
	newNode := &Node[T]{
		Val:   value,
		Left:  nil,
		Right: nil,
	}

	if tree.root == nil {
		tree.root = newNode
		return
	}

	current := tree.root
	prev := current
	prevMove := prevMoveLeft

	// NOTE: handle value == root node value case separately
	if 0 == tree.cmp(value, tree.root.Val) {
		newNode.Left = tree.root.Left
		tree.root.Left = newNode
		return
	}

	for {
		if current == nil {
			insertSingleNode(prev, newNode, prevMove)
			return
		}

		cmpRes := tree.cmp(value, current.Val)

		// NOTE: value == current
		if cmpRes == 0 {
			newNode.Left = current
			insertSingleNode(prev, newNode, prevMove)
			return
		}

		prev = current

		// NOTE: value > current
		if cmpRes > 0 {
			current = current.Right
			prevMove = prevMoveRight
			continue
		}

		// NOTE: value < current
		if cmpRes < 0 {
			current = current.Left
			prevMove = prevMoveLeft
		}
	}
}

func insertSingleNode[T any](prevNode, newNode *Node[T], prevMove prevMoveDirection) {
	switch prevMove {
	case prevMoveLeft:
		prevNode.Left = newNode
	case prevMoveRight:
		prevNode.Right = newNode
	}
}
