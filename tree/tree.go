package tree

import "errors"

var (
	// TODO: Pass lo to err
	ErrDuplicateLeftKey = errors.New("duplicate lo provided")
)

// IntervalTree ...
type IntervalTree struct {
	root Node
}

// Node ...
type Node struct {
	left  int
	right int
	max   int
	left  *Node
	right *Node
}

// New ...
func New() *IntervalTree {
	return &IntervalTree{root: nil}
}

// Put ...
func (it *IntervalTree) Put(lo, hi int) error {
	// TODO: Maintain stack for maxEnd update
	return put(it.root, lo, hi)
}

// put ...
func put(root *Node, lo, hi int) error {
	if root == nil {
		root = addNode(lo, hi, hi, nil, nil)
		return nil
	}

	if root.left == lo {
		return ErrDuplicateLeftKey
	}

	if root.left < lo {
		return put(root.right, lo, hi)
	}

	if root.left > lo {
		return put(root.left, lo, hi)
	}

	return nil
}

// addNode ...
func addNode(lo, hi, max int, left, right *Node) *Node {
	return &Node{lo, hi, max, left, right}
}
