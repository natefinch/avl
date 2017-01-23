package avl

import (
	"bytes"
	"testing"
)

func TestInsert(t *testing.T) {
	tree := New()
	// basic no balancing needed
	//   5
	// 4   6
	tree.Insert([]byte{5})
	tree.Insert([]byte{6})
	tree.Insert([]byte{4})

	expected := &Node{
		Val:    []byte{5},
		height: 2,
		Left: &Node{
			Val:    []byte{4},
			height: 1,
		},
		Right: &Node{
			Val:    []byte{6},
			height: 1,
		},
	}
	treequal(t, expected, tree.Head, "root|")
}

func TestLeftRotate(t *testing.T) {
	tree := New()
	// force a rebalance by making a treee that looks like
	// 1
	//   2
	//     3
	tree.Insert([]byte{1})
	tree.Insert([]byte{2})
	tree.Insert([]byte{3})

	expected := &Node{
		Val:    []byte{2},
		height: 2,
		Left: &Node{
			Val:    []byte{1},
			height: 1,
		},
		Right: &Node{
			Val:    []byte{3},
			height: 1,
		},
	}
	treequal(t, expected, tree.Head, "root|")
}

func TestRightRotate(t *testing.T) {
	tree := New()
	// force a rebalance by making a treee that looks like
	//     3
	//   2
	// 1
	tree.Insert([]byte{3})
	tree.Insert([]byte{2})
	tree.Insert([]byte{1})

	expected := &Node{
		Val:    []byte{2},
		height: 2,
		Left: &Node{
			Val:    []byte{1},
			height: 1,
		},
		Right: &Node{
			Val:    []byte{3},
			height: 1,
		},
	}
	treequal(t, expected, tree.Head, "root|")
}

func TestLeftRightRotate(t *testing.T) {
	tree := New()
	// force a rebalance by making a tree that looks like this:
	//    1
	//      3
	//    2
	tree.Insert([]byte{1})
	tree.Insert([]byte{3})
	tree.Insert([]byte{2})

	expected := &Node{
		Val:    []byte{2},
		height: 2,
		Left: &Node{
			Val:    []byte{1},
			height: 1,
		},
		Right: &Node{
			Val:    []byte{3},
			height: 1,
		},
	}
	treequal(t, expected, tree.Head, "root|")
}

func TestRightLeftRotate(t *testing.T) {
	tree := New()
	// force a rebalance by making a tree that looks like this:
	//    3
	//  1
	//    2

	tree.Insert([]byte{3})
	tree.Insert([]byte{1})
	tree.Insert([]byte{2})

	expected := &Node{
		Val:    []byte{2},
		height: 2,
		Left: &Node{
			Val:    []byte{1},
			height: 1,
		},
		Right: &Node{
			Val:    []byte{3},
			height: 1,
		},
	}
	treequal(t, expected, tree.Head, "root|")
}

func treequal(t *testing.T, expected, actual *Node, path string) {
	if expected == nil {
		if actual == nil {
			return
		}
		t.Fatalf("At node %v, expected nil, but got non-nil with val %v, height %v", path, actual.Val, actual.height)
	}
	if !bytes.Equal(expected.Val, actual.Val) {
		t.Fatalf("At node %v, expected val %x, but got %x", path, expected.Val, actual.Val)
	}
	if expected.height != actual.height {
		t.Fatalf("At node %v, expected height %v, but got %v", path, expected.height, actual.height)
	}
	treequal(t, expected.Left, actual.Left, path+"L")
	treequal(t, expected.Right, actual.Right, path+"R")
}
