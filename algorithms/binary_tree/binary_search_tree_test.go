package binary_tree

import "testing"

func TestTreeSize(t *testing.T) {
	var tree *Tree = nil

	size := TreeSize(tree)

	if size != 0 {
		t.Error("Expected empty tree, got", size)
	}

	tree = &Tree{2, nil, &Tree{1, nil, nil, nil}, &Tree{3, nil, nil, nil}}
	size = TreeSize(tree)

	if size != 3 {
		t.Error("Expected tree of size three, got", size)
	}
}

func TestInsertion(t *testing.T) {
	var tree *Tree = nil
	Insert(tree, 1)
	Insert(tree, 2)
	Insert(tree, 3)
}
