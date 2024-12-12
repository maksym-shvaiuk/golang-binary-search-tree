package binsrch_tree

import "fmt"

func printTreeStructure[T any](node *Node[T], prefix string, isTail bool, logger Logger, isLeft bool) {
	if node == nil {
		return
	}

	// Use a marker for left or right child
	childLabel := "Root"
	if prefix != "" {
		if isLeft {
			childLabel = "L--"
		} else {
			childLabel = "R--"
		}
	}

	// Log the current node with its label
	logger.Print(fmt.Sprintf("%s%s %s %v", prefix, branchMarker(isTail), childLabel, node.Val))

	// Prepare prefix for children
	childPrefix := prefix + branchIndent(isTail)

	// Log children recursively
	if node.Left != nil || node.Right != nil {
		// Log the left child
		if node.Left != nil {
			printTreeStructure(node.Left, childPrefix, node.Right == nil, logger, true)
		} else {
			logger.Print(fmt.Sprintf("%s%s L-- [null]", childPrefix, branchMarker(true)))
		}

		// Log the right child
		if node.Right != nil {
			printTreeStructure(node.Right, childPrefix, true, logger, false)
		} else {
			logger.Print(fmt.Sprintf("%s%s R-- [null]", childPrefix, branchMarker(true)))
		}
	}
}

func branchMarker(isTail bool) string {
	if isTail {
		return "└──"
	}
	return "├──"
}

func branchIndent(isTail bool) string {
	if isTail {
		return "    "
	}
	return "│   "
}
