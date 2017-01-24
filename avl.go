// Package avl contains an implmentation of an AVL tree that contains []byte
// keys.
package avl

import (
	"bytes"
	"sync"
)

// New returns a new AVL Tree structure.
func New() *Tree {
	return &Tree{mu: &sync.Mutex{}}
}

// Tree is AVL tree designed to store slices of bytes.
type Tree struct {
	Head *Node
	mu   *sync.Mutex
}

// Insert adds the value to the tree.
func (t *Tree) Insert(b []byte) {
	defer t.mu.Unlock()
	t.mu.Lock()

	t.Head = t.Head.insert(b)
}

// Search returns the node of the tree with the value given, otherwise nil.
func (t *Tree) Search(b []byte) *Node {
	defer t.mu.Unlock()
	t.mu.Lock()
	n := t.Head
	for n != nil {
		switch bytes.Compare(b, n.Val) {
		case -1:
			n = n.Left
		case 0:
			return n
		case 1:
			n = n.Right
		}
	}
	return nil
}

// Delete removes the node of the tree with the value of b.
func (t *Tree) Delete(b []byte) {
	defer t.mu.Unlock()
	t.mu.Lock()

	t.Head = t.Head.delete(b)
}

// Node is a node in the AVL tree.
type Node struct {
	Val    []byte
	Left   *Node
	Right  *Node
	height int
}

// Height returns the height of this node's subtree, zero if n is nil.
func (n *Node) Height() int {
	if n == nil {
		return 0
	}
	return n.height
}

// insert inserts a node with the given value into the tree and returns the
// value that should be used as this subtree's root.
func (n *Node) insert(b []byte) *Node {
	if n == nil {
		return &Node{Val: b, height: 1}
	}

	switch bytes.Compare(b, n.Val) {
	case -1:
		n.Left = n.Left.insert(b)
	case 0, 1:
		n.Right = n.Right.insert(b)
	}
	return n.rebalance()
}

// delete finds the given node in the tree and deletes it, returning the Node
// that should replace it in the tree above.
func (n *Node) delete(b []byte) *Node {
	if n == nil {
		return nil
	}

	switch bytes.Compare(b, n.Val) {
	case -1:
		n.Left = n.Left.delete(b)
		n.rebalance()
		return n
	case 1:
		n.Right = n.Right.delete(b)
		n.rebalance()
		return n
	}
	// case 0
	if n.Left == nil {
		// don't need to rebalance here since n.Right isn't actually changing
		// and thus should still be balanced.
		return n.Right
	}

	// We're going to replace n with the minimal node from the right subtree
	// (and then n drops into garbage collection).
	min := findMin(n.Right)
	min.Right = removeMin(n.Right)
	min.Left = n.Left

	// return min or however the subtree under min rebalances.
	return min.rebalance()
}

// findmind returns the node with minimal key in n's subtree
func findMin(n *Node) *Node {
	for n.Left != nil {
		n = n.Left
	}
	return n
}

// removeMin removes the minimal (leftmost) node in this tree and then returns
// the tree that remains.
func removeMin(n *Node) *Node {
	if n.Left == nil {
		return n.Right
	}
	n.Left = removeMin(n.Left)
	return n.rebalance()
}

// rebalance rotates this tree if needed to regain balance, and returns the node
// that should replace this node in the tree above.
func (n *Node) rebalance() *Node {
	n.setHeight()
	// Since we constantly rebalance, we should only ever be at a max of +2 or
	// -2.
	switch n.balFactor() {
	case 2:
		if n.Right.balFactor() < 0 {
			// will need a right/left rotate
			n.Right = n.Right.rotateRight()
		}
		return n.rotateLeft()
	case -2:
		if n.Left.balFactor() > 0 {
			// will need a left/right rotate
			n.Left = n.Left.rotateLeft()
		}
		return n.rotateRight()
	default:
		// -1, 0, 1: already balanced
		return n
	}
}

func (n *Node) balFactor() int {
	if n == nil {
		return 0
	}
	return n.Right.Height() - n.Left.Height()
}

func (n *Node) setHeight() {
	n.height = max(n.Left.Height(), n.Right.Height()) + 1
}

// rotateRight rotates the tree under the given node right, returning the node
// that should replace this node's spot.
func (n *Node) rotateRight() *Node {
	// store our new root
	l := n.Left
	n.Left = l.Right
	l.Right = n
	n.setHeight()
	l.setHeight()
	return l
}

// rotateLeft rotates the tree under the given node left, returning the node
// that should replace this node in the tree above.
func (n *Node) rotateLeft() *Node {
	// store our new root
	r := n.Right
	n.Right = r.Left
	r.Left = n
	n.setHeight()
	r.setHeight()
	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
