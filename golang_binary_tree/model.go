package golang_binary_tree

import "errors"

var ErrorCmpFunctionIsNil = errors.New("comparison function is nil")

type prevMoveDirection int

const (
	prevMoveLeft prevMoveDirection = iota
	prevMoveRight
)

type Algo int

const (
	AlgoIterative Algo = iota
	AlgoRecursive
)

type Node[T any] struct {
	Val   T
	Left  *Node[T]
	Right *Node[T]
}

type Tree[T any] struct {
	root *Node[T]
	// NOTE: cmp should return 1 if a > b, -1 if a < b, 0 if a == b
	cmp  func(a, b T) int
	algo Algo
}

// NOTE: Interface required in PrintTree.
type Logger interface {
	Print(string)
}
