package tree

import "errors"

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
	if it.root == nil {
		it.root = addNode(lo, hi, hi, nil, nil)
		return nil
	}

	if it.root.left == lo {
		// TODO: Move err to var
		// TODO: Pass lo to err
		return errors.New("duplicate lo provided")
	}
}

// addNode ...
func addNode(lo, hi, max int, left, right *Node) *Node {
	return &Node{lo, hi, max, left, right}
}
