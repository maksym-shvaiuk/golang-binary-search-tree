package binsrch_tree

import "fmt"

func printTreeStructure[T any](node *Node[T], prefix string, isTail bool, logger Logger) {
	if node == nil {
		return
	}

	// Log the current node
	logger.Print(fmt.Sprintf("%s%s── %v", prefix, branchMarker(isTail), node.Val))

	// Adjust prefix for children
	childPrefix := prefix + branchIndent(isTail)

	// Recursively log children
	if node.Left != nil || node.Right != nil {
		printTreeStructure(node.Left, childPrefix, false, logger)
		printTreeStructure(node.Right, childPrefix, true, logger)
	}
}

func branchMarker(isTail bool) string {
	if isTail {
		return "└"
	}
	return "├"
}

func branchIndent(isTail bool) string {
	if isTail {
		return "    "
	}
	return "│   "
}
