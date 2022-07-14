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
	node, _ := search(it.root, nil, lo, hi)
	return node
}

// Delete ...
func (it *IntervalTree) Delete(lo, hi int) error {
	// Search if node exists
	//
	// TODO: Remove parent from signature
	// if not required
	node, _ := search(it.root, nil, lo, hi)

	// Error out if it does not
	if node == nil {
		return ErrNodeDoesNotExist
	}

	return nil
}

// Delete node at point
func delete(lo, hi int) {
	node, parent := search(root, nil, lo, hi)
	if node == nil {
		return
	}

	if node.left == nil && node.right == nil {
		node = nil
		return
	}

	if node.left == nil && node.right != nil {
		node = node.right
		return
	}

	if node.right == nil && node.left != nil {
		node = node.left
		return
	}

	replacement := maxNodeInLeftSubtree(node)
	delete(replacement.Left, replacement.Right)

	node.Left, node.Right = replacement.Left, replacement.Right
}

func maxNodeInLeftSubtree(root *Node) *Node {
	if root == nil {
		return nil
	}

	if root.Left == nil {
		return nil
	}

	root = root.Left
	for root.Right != nil {
		root = root.Right
	}

	return root
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
func search(root, parent *Node, lo, hi int) (*Node, *Node) {
	if root == nil {
		return nil, nil
	}

	if root.left == lo && root.right == hi {
		return root, parent
	}

	if root.left < lo {
		return search(root.right, root, lo, hi)
	}

	return search(root.left, root, lo, hi)
}

// addNode ...
func addNode(lo, hi, max int, left, right *Node) *Node {
	return &Node{lo, hi, max, left, right}
}
