package structures

import "fmt"

type binaryTreeNode[T any] struct {
	value T
	left  *binaryTreeNode[T]
	right *binaryTreeNode[T]
}

func NewBinaryTreeNode[T any](value T) *binaryTreeNode[T] {
	return &binaryTreeNode[T]{value: value, left: nil, right: nil}
}

type binaryTree[T any] struct {
	root *binaryTreeNode[T]
}

func NewBinaryTree[T any](root *binaryTreeNode[T]) *binaryTree[T] {
	return &binaryTree[T]{root: root}
}

// Insert a node to the binary tree via level-order
func (bt *binaryTree[T]) Insert(node *binaryTreeNode[T]) {
	// TODO: Add more tests for this method (nil root case, etc.)
	queue := NewQueue[*binaryTreeNode[T]]()
	queue.Push(bt.root)

	for queue.GetLength() != 0 {
		current := queue.Pop()

		if current.left == nil {
			current.left = node
			return
		} else {
			queue.Push(current.left)
		}

		if current.right == nil {
			current.right = node
			return
		} else {
			queue.Push(current.right)
		}
	}
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

// Get the string representation of the binary tree
func (bt *binaryTree[T]) ToString() string {
	return bt.toString("", bt.root)
}

func (bt *binaryTree[T]) toString(treeString string, current *binaryTreeNode[T]) string {
	// TODO: Update the formatting here to better handle the end of the tree
	if current == nil {
		return ""
	}

	treeString += bt.toString(treeString, current.left)
	treeString += fmt.Sprintf("%v, ", current.value)
	treeString += bt.toString(treeString, current.right)

	return treeString
}
