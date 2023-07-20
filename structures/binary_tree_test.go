package structures

import "testing"

func Test_GivenBinaryTree_WhenInsert_ThenNodeIsSet(t *testing.T) {
	tree := NewBinaryTree[int](NewBinaryTreeNode(1))
	expected := "2, 1, "

	tree.Insert(NewBinaryTreeNode(2))

	if tree.ToString() != expected {
		t.Errorf("Tree = %s, want %s", tree.ToString(), expected)
	}
}
