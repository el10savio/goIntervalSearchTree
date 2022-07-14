package tree

import "errors"

var (
	ErrDuplicateLeftKey = errors.New("duplicate lo provided") // TODO: Pass lo to err
	ErrNodeDoesNotExist = errors.New("node does not exist")   // TODO: Pass in lo & hi
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

// Search ...
func (it *IntervalTree) Search(lo, hi int) *Node {
	return search(it.root, lo, hi)
}

// Delete ...
func (it *IntervalTree) Delete(lo, hi int) error {
	// Search if node exists
	node := search(it.root, lo, hi)

	// Error out if it does not
	if node == nil {
		return ErrNodeDoesNotExist
	}

	return nil
}

// Delete node at point
func delete(root *Node, value int) error {
	// Replace with left node
	// If left node does not exist
	// replace with right
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

// search ...
func search(root *Node, lo, hi int) *Node {
	if root == nil {
		return nil
	}

	if root.left == lo && root.right == hi {
		return root
	}

	if root.left < lo {
		return search(root.right, lo, hi)
	}

	return search(root.left, lo, hi)
}

// addNode ...
func addNode(lo, hi, max int, left, right *Node) *Node {
	return &Node{lo, hi, max, left, right}
}
