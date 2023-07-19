package structures

import "fmt"

type BinaryTreeNode[T any] struct {
	value T
	left  *BinaryTreeNode[T]
	right *BinaryTreeNode[T]
}

type BinaryTree[T any] struct {
	root *BinaryTreeNode[T]
}

func NewBinaryTreeNode[T any](value T) *BinaryTreeNode[T] {
	return &BinaryTreeNode[T]{value: value, left: nil, right: nil}
}

func NewBinaryTree[T any](root *BinaryTreeNode[T]) *BinaryTree[T] {
	return &BinaryTree[T]{root: root}
}

func (bt *BinaryTree[T]) Insert(node *BinaryTreeNode[T]) {
	// TODO: Make a queue data structure
	// queue := NewQueue()
}

// func (bt *BinaryTree) Remove(node *BinaryTreeNode) {

// }

// func (bt *BinaryTree) Search(node *BinaryTreeNode) bool {

// }

// func (bt *BinaryTree) GetNodeLevel(node *BinaryTreeNode) int {

// }

// func (bt *BinaryTree) GetHeight() int {

// }

// func (bt *BinaryTree) GetSize() int {

// }

func (bt *BinaryTree[T]) ToString() string {
	return bt.toString("", bt.root)
}

func (bt *BinaryTree[T]) toString(treeString string, current *BinaryTreeNode[T]) string {
	if current == nil {
		return ""
	}

	treeString += bt.toString(treeString, current.left)
	treeString += fmt.Sprintf("%v, ", current.value)
	treeString += bt.toString(treeString, current.right)

	return treeString
}
