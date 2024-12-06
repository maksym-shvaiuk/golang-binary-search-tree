package binsrch_tree

// deleteIterative deletes a node with requested value from tree iteratively.
// Returns true if node found and deleted, false otherwise.
// Restrictions: value should not be nil.
func (tree *Tree[T]) deleteIterative(value T) bool {
	if tree.root == nil {
		return false
	}

	if tree.cmp(value, tree.root.Val) == 0 {
		tree.deleteRootNodeIterative()
		return true
	}

	current := tree.root
	prev := tree.root
	prevMove := prevMoveLeft

	for {
		if current == nil {
			return false
		}

		cmpRes := tree.cmp(value, current.Val)

		if cmpRes == 0 {
			deleteNodeIterative(prev, current, prevMove)
			return true
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

func (tree *Tree[T]) deleteRootNodeIterative() {
	var numberOfChildren int
	if tree.root.Left != nil {
		numberOfChildren++
	}
	if tree.root.Right != nil {
		numberOfChildren++
	}

	if numberOfChildren == 0 {
		tree.root = nil
		return
	} else if numberOfChildren == 1 {
		if tree.root.Left != nil {
			tree.root = tree.root.Left
		} else {
			tree.root = tree.root.Right
		}
		return
	}

	deleteNodeTwoChildrenIterative(tree.root)
	return
}

func deleteNodeIterative[T any](prevNode, nodeToDelete *Node[T], prevMove prevMoveDirection) {
	var numberOfChildren int
	if nodeToDelete.Left != nil {
		numberOfChildren++
	}
	if nodeToDelete.Right != nil {
		numberOfChildren++
	}

	if numberOfChildren == 0 {
		deleteNodeNoChildrenIterative[T](prevNode, prevMove)
	} else if numberOfChildren == 1 {
		deleteNodeSingleChildIterative[T](prevNode, nodeToDelete, prevMove)
	} else {
		deleteNodeTwoChildrenIterative(nodeToDelete)
	}
}

func deleteNodeNoChildrenIterative[T any](prevNode *Node[T], prevMove prevMoveDirection) {
	switch prevMove {
	case prevMoveLeft:
		prevNode.Left = nil
	case prevMoveRight:
		prevNode.Right = nil
	}
}

func deleteNodeSingleChildIterative[T any](prevNode, nodeToDelete *Node[T], prevMove prevMoveDirection) {
	switch prevMove {
	case prevMoveLeft:
		if nodeToDelete.Left != nil {
			prevNode.Left = nodeToDelete.Left
		} else if nodeToDelete.Right != nil {
			prevNode.Left = nodeToDelete.Right
		}
	case prevMoveRight:
		if nodeToDelete.Left != nil {
			prevNode.Right = nodeToDelete.Left
		} else if nodeToDelete.Right != nil {
			prevNode.Right = nodeToDelete.Right
		}
	}
}

func deleteNodeTwoChildrenIterative[T any](nodeToDelete *Node[T]) {
	// step 1: find node with max value in left subtree and it's parent node
	maxNodeLeftSubTree, maxNodeLeftSubTreeParent := getMaxSubtreeNodeWithParent[T](nodeToDelete.Left)
	
	
	// step 2: copy the value from max node of the left subtree to the nodeToDelete
	nodeToDelete.Val = maxNodeLeftSubTree.Val

	// Step 3: Adjust pointers to remove maxNodeLeftSubTree from its position
	// If maxNodeLeftSubTreeParent is nil, maxNodeLeftSubTree is the direct left child of nodeToDelete
	if maxNodeLeftSubTreeParent == nil {
		// If the maxNodeLeftSubTreeParent is nil - it means that
		// maxNodeLeftSubTree == nodeToDelete.Left
		// so, the nodeToDelete is the parent of the maxNodeLeftSubTree.
		nodeToDelete.Left = maxNodeLeftSubTree.Left
	} else {
		// Otherwise, bypass maxNodeLeftSubTree
		maxNodeLeftSubTreeParent.Right = maxNodeLeftSubTree.Left
	}
}
