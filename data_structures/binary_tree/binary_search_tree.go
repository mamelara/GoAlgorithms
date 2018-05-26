package binary_tree

type Tree struct {
	value  int
	parent *Tree
	left   *Tree
	right  *Tree
}

var root *Tree = nil

func IsLeaf(tree *Tree) bool {
	return tree.left == nil && tree.right == nil
}

func SearchTree(tree *Tree, item int) *Tree {

	if IsLeaf(tree) {
		return tree
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

func treeInsertion(tree *Tree, item int) {
	node := tree

	if node != nil {

		node = SearchTree(node, item)

		newNode := &Tree{value: item}
		if item > node.value {
			node.right = newNode
		} else if item < node.value {
			node.left = newNode
		}

		newNode.parent = node
	} else {
		tree = &Tree{value: item}
	}
}

func Insert(tree *Tree, item int) {
	treeInsertion(tree, item)
}

func TreeSize(tree *Tree) int {
	if tree == nil {
		return 0
	} else {
		return TreeSize(tree.left) + TreeSize(tree.right) + 1
	}
}
