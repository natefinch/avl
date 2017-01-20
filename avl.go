// Package avl contains an implmentation of an AVL tree that contains []byte
// values.
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

// Node is a node in the tree.
type Node struct {
	Val   []byte
	Left  *Node
	Right *Node
}

// Insert adds the value to the tree.
func (t *Tree) Insert(b []byte) *Node {
	// TODO: rebalance
	defer t.mu.Unlock()
	t.mu.Lock()
	if t.Head == nil {
		t.Head = &Node{Val: b}
		return t.Head, nil
	}

	n := t.Head
	for {
		switch bytes.Compare(b, n.Val) {
		case -1:
			if n.Left == nil {
				n.Left = &Node{Val: b}
				return n.Left
			}
			n = n.Left
		case 0, 1:
			if n.Right == nil {
				n.Right = &Node{Val: b}
				return n.Right
			}
			n = n.Right
		}
	}
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

// Delete removes the node of the tree with the value of b.  It reports if the
// value existed in the tree.
func (t *Tree) Delete(b []byte) bool {
	// parent holds the parent of the node we're deleting
	var parent *Node

	n := t.Head
	for n != nil {
		switch bytes.Compare(b, n.Val) {
		case -1:
			parent = n
			n = n.Left
		case 1:
			parent = n
			n = n.Right
		case 0:
			// ooh, shiny.
			t.delete(n, parent)
			return true
		}
	}
	return false
}

func (t *Tree) delete(n, parent *Node) {
	// TODO: the hard stuff ;)
}

// Walk implements an in-order walk of a Tree using recursion.
//
// function f should return false if it wants the walk to stop
// Walk returns false if f ever returns false, otherwise true
func Walk(n *Node, f func(*Node) bool) bool {
	if n == nil {
		// Found a leaf.
		return true
	}
	if !Walk(n.Left, f) {
		return false
	}
	if !f(n) {
		return false
	}
	return Walk(n.Right, f)
}
