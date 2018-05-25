package binary_tree

type Tree struct {
	value  int
	parent *Tree
	left   *Tree
	right  *Tree
}

var root *Tree = Tree{}

func SearchTree(tree *Tree, item int) *Tree {

	if tree == nil {
		return nil
	}

	if tree.value == item {
		return tree
	}

	if item > tree.value {
		return SearchTree(tree.right, item)
	} else {
		return SearchTree(tree.left, item)
	}
}
