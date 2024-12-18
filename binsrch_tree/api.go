package binsrch_tree

// New creates a tree.
// By-default, the iterative algo is selected.
// Use the SetAlgo() to change it to recursive.
func New[T any](cmp func(a, b T) int) Tree[T] {
	return Tree[T]{
		root: nil,
		cmp:  cmp,
		algo: AlgoIterative,
	}
}

// Clone creates a deep copy of a tree.
func (tree *Tree[T]) Clone() Tree[T] {
	return tree.clone()
}

func (tree *Tree[T]) SetAlgo(algo Algo) {
	tree.algo = algo
}

func (tree *Tree[T]) GetRoot() *Node[T] {
	return tree.root
}

// Insert inserts a value.
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

// Delete deletes a node with the value provided from the tree.
// The deletion algo (iterative vs recursive) is defined by tree.algo.
// Restrictions: value should not be nil.
func (tree *Tree[T]) Delete(value T) (bool, error) {
	var res bool

	if tree.cmp == nil {
		return res, ErrorCmpFunctionIsNil
	}

	switch tree.algo {
	case AlgoIterative:
		res = tree.deleteIterative(value)
	case AlgoRecursive:
		res = tree.deleteRecursively(value)
	default:
		res = tree.deleteIterative(value)
	}

	return res, nil
}

// Exists checks if a node with requested value exists in tree.
// Returns true if node found, false otherwise.
// Restrictions: value should not be nil.
func (tree *Tree[T]) Exists(value T) bool {
	return tree.findNode(value) != nil
}

// MaxValue returns max value stored in tree.
func (tree *Tree[T]) MaxValue() (T, error) {
	if tree.root == nil {
		// return zero value of type T
		var res T
		return res, ErrorTreeRootIsNil
	}

	return getMaxSubtreeNode[T](tree.root).Val, nil
}

// MinValue returns min value stored in tree.
func (tree *Tree[T]) MinValue() (T, error) {
	if tree.root == nil {
		// return zero value of type T
		var res T
		return res, ErrorTreeRootIsNil
	}

	return getMinSubtreeNode[T](tree.root).Val, nil
}

// PrintTree prints the tree in a structured, human-readable format.
// The user should implement the custom logger.
func (tree *Tree[T]) PrintTree(logger Logger) {
	if tree.root == nil {
		logger.Print("The tree is empty.")
		return
	}
	printTreeStructure(tree.root, "", true, logger, false)
}
