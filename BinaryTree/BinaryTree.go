package BinaryTree

import "cmp"

type Node[T cmp.Ordered] struct {
	data   T
	parent *Node[T]
	left   *Node[T]
	right  *Node[T]
}

type BinaryTree[T cmp.Ordered] struct {
	root *Node[T]
}

func (tree *BinaryTree[T]) InsertItem(i T) {
	if tree.root == nil {
		tree.root = &Node[T]{data: i}
		return
	}
	currentNode := tree.root
	for {
		if i > currentNode.data {
			if currentNode.right == nil {
				currentNode.right = &Node[T]{data: i, parent: currentNode}
				return
			}
			currentNode = currentNode.right
		} else {
			if currentNode.left == nil {
				currentNode.left = &Node[T]{data: i, parent: currentNode}
				return
			}
			currentNode = currentNode.left
		}
	}
}

func (tree *BinaryTree[T]) SearchItem(i T) (*Node[T], bool) {
	if tree.root == nil {
		return nil, false
	}
	currentNode := tree.root
	for currentNode != nil {

		if cmp.Compare(i, currentNode.data) == 0 {
			return currentNode, true
		} else if cmp.Compare(i, currentNode.data) == 1 {
			currentNode = currentNode.right
		} else if cmp.Compare(i, currentNode.data) == -1 {
			currentNode = currentNode.left
		}
	}
	return nil, false
}

func (tree *BinaryTree[T]) InorderTraversal(subtree *Node[T], callback func(T)) {
	if subtree.left != nil {
		tree.InorderTraversal(subtree.left, callback)
	}
	callback(subtree.data)
	if subtree.right != nil {
		tree.InorderTraversal(subtree.right, callback)
	}
}

func (tree *BinaryTree[T]) PreorderTraversal(subtree *Node[T], callback func(T)) {
	callback(subtree.data)
	if subtree.left != nil {
		tree.PreorderTraversal(subtree.left, callback)
	}
	if subtree.right != nil {
		tree.PreorderTraversal(subtree.right, callback)
	}
}

func (tree *BinaryTree[T]) PostorderTraversal(subtree *Node[T], callback func(T)) {
	if subtree.left != nil {
		tree.PostorderTraversal(subtree.left, callback)
	}
	if subtree.right != nil {
		tree.PostorderTraversal(subtree.right, callback)
	}
	callback(subtree.data)
}
