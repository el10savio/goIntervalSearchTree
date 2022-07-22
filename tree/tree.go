package tree

import (
	"errors"
	"fmt"
)

var (
	ErrDuplicateLeftKey = errors.New("duplicate lo provided") // TODO: Pass lo to err
	ErrNodeDoesNotExist = errors.New("node does not exist")   // TODO: Pass in lo & hi
)

// IntervalTree ...
type IntervalTree struct {
	root *Node
}

// Node ...
type Node struct {
	interval Interval
	max      int
	left     *Node
	right    *Node
}

// Interval
type Interval struct {
	left  int
	right int
}

// New ...
func New() *IntervalTree {
	return &IntervalTree{}
}

// Clear ...
func (it *IntervalTree) Clear() {
	it.root = nil
}

// Put ...
func (it *IntervalTree) Put(lo, hi int) error {
	// TODO: Maintain stack for maxEnd update
	it.root, _ = put(it.root, it.root, lo, hi)
	return nil
}

// Search ...
func (it *IntervalTree) Search(lo, hi int) *Node {
	return search(it.root, lo, hi)
}

// Delete ...
func (it *IntervalTree) Delete(lo, hi int) error {
	// Search if node exists
	node := search(it.root, lo, hi)
	if node == nil {
		return ErrNodeDoesNotExist
	}

	deleteNode(node, lo, hi)

	return nil
}

// Delete node at point
func deleteNode(root *Node, lo, hi int) {
	node := search(root, lo, hi)
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
	deleteNode(node, replacement.interval.left, replacement.interval.right)

	node = &Node{
		interval: Interval{replacement.interval.left, replacement.interval.right},
		max:      replacement.max,
		left:     replacement.left,
		right:    replacement.right,
	}
}

// toList ...
func (it *IntervalTree) toList() []int {
	return toListHandler(it.root)
}

// toListHandler ...
func toListHandler(node *Node) []int {
	if node == nil {
		return []int{}
	}
	return append(
		[]int{node.interval.left, node.interval.right},
		append(toListHandler(node.left), toListHandler(node.right)...)...,
	)
}

func maxNodeInLeftSubtree(root *Node) *Node {
	if root == nil {
		return nil
	}

	if root.left == nil {
		return nil
	}

	root = root.left
	for root.right != nil {
		root = root.right
	}

	return root
}

// put ...
func put(root, node *Node, lo, hi int) (*Node, error) {
	if node == nil {
		return addNode(lo, hi, hi, nil, nil), nil
	}

	if node.interval.left < lo {
		fmt.Println("right")
		return put(root, node.right, lo, hi)
	}

	if node.interval.left > lo {
		fmt.Println("left")
		return put(root, node.left, lo, hi)
	}

	return root, ErrDuplicateLeftKey
}

// search ...
func search(root *Node, lo, hi int) *Node {
	if root == nil {
		return nil
	}

	if root.interval.left == lo && root.interval.right == hi {
		return root
	}

	if root.interval.left < lo {
		return search(root.right, lo, hi)
	}

	return search(root.left, lo, hi)
}

// addNode ...
func addNode(lo, hi, max int, left, right *Node) *Node {
	return &Node{Interval{lo, hi}, max, left, right}
}
