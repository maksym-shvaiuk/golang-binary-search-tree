package binsrch_tree

// Creates a tree.
// By-default, the iterative algo is selected.
// Use the SetAlgo() to change it to recursive.
func New[T any](cmp func(a, b T) int) Tree[T] {
	return Tree[T]{
		root: nil,
		cmp:  cmp,
		algo: AlgoIterative,
	}
}

func (tree *Tree[T]) SetAlgo(algo Algo) {
	tree.algo = algo
}

func (tree *Tree[T]) GetRoot() *Node[T] {
	return tree.root
}

// NOTE: Inserts a value.
// The insertion algo (iterative vs recursive) is defined by tree.algo.
// Restrictions: value should not be nil.
func (tree *Tree[T]) Insert(value T) error {
	if tree.cmp == nil {
		return ErrorCmpFunctionIsNil
	}

	switch tree.algo {
	case AlgoIterative:
		tree.insertIterative(value)
	case AlgoRecursive:
		tree.insertRecursively(value)
	default:
		tree.insertIterative(value)
	}

	return nil
}

// Exists checks if a node with requested value exists in tree.
// Returns true if node found, false otherwise.
// Restrictions: value should not be nil.
func (tree *Tree[T]) Exists(value T) bool {
	return tree.existsIterative(value)
}

// PrintTree prints the tree in a structured, human-readable format.
// The user should implement the custom logger.
func (tree *Tree[T]) PrintTree(logger Logger) {
	if tree.root == nil {
		logger.Print("The tree is empty.")
		return
	}
	printTreeStructure(tree.root, "", true, logger)
}
