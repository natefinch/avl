package avl

import (
	"bytes"
	"testing"
)

func TestLeftRotate(t *testing.T) {
	// make a treee that looks like
	// 1
	//   2
	//     3
	actual := &Node{
		Val:    []byte{1},
		height: 3,
		Right: &Node{
			Val:    []byte{2},
			height: 2,
			Right: &Node{
				Val:    []byte{3},
				height: 1,
			},
		},
	}

	actual = actual.rotateLeft()

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
	treequal(t, expected, actual, "root")
}

func TestRightRotate(t *testing.T) {
	// make a tree that looks like
	//     3
	//   2
	// 1
	actual := &Node{
		Val:    []byte{3},
		height: 3,
		Left: &Node{
			Val:    []byte{2},
			height: 2,
			Left: &Node{
				Val:    []byte{1},
				height: 1,
			},
		},
	}

	actual = actual.rotateRight()

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
	treequal(t, expected, actual, "root")
}

func TestLeftRightRotate(t *testing.T) {
	// make a tree that looks like this:
	//    1
	//      3
	//    2
	actual := &Node{
		Val:    []byte{1},
		height: 3,
		Right: &Node{
			Val:    []byte{3},
			height: 2,
			Left: &Node{
				Val:    []byte{2},
				height: 1,
			},
		},
	}
	actual = actual.rebalance()

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
	treequal(t, expected, actual, "root")
}

func TestRightLeftRotate(t *testing.T) {
	// make a tree that looks like this:
	//    3
	//  1
	//    2
	actual := &Node{
		Val:    []byte{3},
		height: 3,
		Left: &Node{
			Val:    []byte{1},
			height: 2,
			Right: &Node{
				Val:    []byte{2},
				height: 1,
			},
		},
	}
	actual = actual.rebalance()

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
	treequal(t, expected, actual, "root")
}

func TestBalFactor(t *testing.T) {
	var n *Node
	if actual := n.balFactor(); actual != 0 {
		t.Errorf("Expected zero balance factor for nil node, but got %d", actual)
	}
	n = &Node{
		Val:    []byte{2},
		height: 2,
		Left: &Node{
			Val:    []byte{1},
			height: 1,
		},
	}
	if actual := n.balFactor(); actual != -1 {
		t.Errorf("Expected balance factor -1, but got %d", actual)
	}
}

func TestInsertSimple(t *testing.T) {
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
	treequal(t, expected, tree.Head, "root")
}

func TestInsertComplex(t *testing.T) {
	tree := New()
	tree.Insert([]byte{1})
	tree.Insert([]byte{2})
	tree.Insert([]byte{3})
	tree.Insert([]byte{4})
	tree.Insert([]byte{5})

	expected := &Node{
		Val:    []byte{2},
		height: 3,
		Left: &Node{
			Val:    []byte{1},
			height: 1,
		},
		Right: &Node{
			Val:    []byte{4},
			height: 2,
			Left: &Node{
				Val:    []byte{3},
				height: 1,
			},
			Right: &Node{
				Val:    []byte{5},
				height: 1,
			},
		},
	}
	treequal(t, expected, tree.Head, "root")
}

func TestDeleteLeaf(t *testing.T) {
	tree := New()
	tree.Insert([]byte{1})
	tree.Insert([]byte{2})
	tree.Insert([]byte{4})
	tree.Insert([]byte{3})

	tree.Delete([]byte{3})

	expected := &Node{
		Val:    []byte{2},
		height: 2,
		Left: &Node{
			Val:    []byte{1},
			height: 1,
		},
		Right: &Node{
			Val:    []byte{4},
			height: 1,
		},
	}
	treequal(t, expected, tree.Head, "root")
}

func TestDeleteRoot(t *testing.T) {
	tree := New()

	// starts as this:
	//       2
	//   1        6
	// 0        4   7
	//        3  5
	tree.Head = &Node{
		Val:    []byte{2},
		height: 4,
		Left: &Node{
			Val:    []byte{1},
			height: 2,
			Left: &Node{
				Val:    []byte{0},
				height: 1,
			},
		},
		Right: &Node{
			Val:    []byte{6},
			height: 3,
			Left: &Node{
				Val:    []byte{4},
				height: 2,
				Left: &Node{
					Val:    []byte{3},
					height: 1,
				},
				Right: &Node{
					Val:    []byte{5},
					height: 1,
				},
			},
			Right: &Node{
				Val:    []byte{7},
				height: 1,
			},
		},
	}

	// now delete 2, and we should just replace 2 with 3
	tree.Delete([]byte{2})

	expected := &Node{
		Val:    []byte{3},
		height: 4,
		Left: &Node{
			Val:    []byte{1},
			height: 2,
			Left: &Node{
				Val:    []byte{0},
				height: 1,
			},
		},
		Right: &Node{
			Val:    []byte{6},
			height: 3,
			Left: &Node{
				Val:    []byte{4},
				height: 2,
				Right: &Node{
					Val:    []byte{5},
					height: 1,
				},
			},
			Right: &Node{
				Val:    []byte{7},
				height: 1,
			},
		},
	}
	treequal(t, expected, tree.Head, "root")
}

func TestDeleteNil(t *testing.T) {
	tree := New()
	tree.Delete([]byte{3})
}

func TestSearch(t *testing.T) {
	tree := New()

	// starts as this:
	//       2
	//   1        6
	// 0        4   7
	//        3  5
	tree.Head = &Node{
		Val:    []byte{2},
		height: 4,
		Left: &Node{
			Val:    []byte{1},
			height: 2,
			Left: &Node{
				Val:    []byte{0},
				height: 1,
			},
		},
		Right: &Node{
			Val:    []byte{6},
			height: 3,
			Left: &Node{
				Val:    []byte{4},
				height: 2,
				Left: &Node{
					Val:    []byte{3},
					height: 1,
				},
				Right: &Node{
					Val:    []byte{5},
					height: 1,
				},
			},
			Right: &Node{
				Val:    []byte{7},
				height: 1,
			},
		},
	}

	actual := tree.Search([]byte{10})
	if actual != nil {
		t.Fatalf("Searched for value that shouldn't exist, but it does.")
	}

	actual = tree.Search([]byte{4})

	expected := &Node{
		Val:    []byte{4},
		height: 2,
		Left: &Node{
			Val:    []byte{3},
			height: 1,
		},
		Right: &Node{
			Val:    []byte{5},
			height: 1,
		},
	}
	treequal(t, expected, actual, "root")
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
