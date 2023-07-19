package structures

import "testing"

func Test_GivenLinkedList_WhenCreate_ThenHeadIsSet(t *testing.T) {
	list := NewLinkedList(NewLinkedListNode(1))

	if length := list.GetLength(); length != 1 {
		t.Errorf("Length = %d, want 1", length)
	}

	if list.head == nil || list.head.value != 1 || list.head.next != nil {
		t.Errorf("Head = %v, want Value: 1, Next: nil", list.head)
	}
}

func Test_GivenLinkedList_WhenGetTail_ThenTailIsHead(t *testing.T) {
	list := NewLinkedList(NewLinkedListNode(1))

	if length := list.GetLength(); length != 1 {
		t.Errorf("Length = %d, want 1", length)
	}

	if list.head == nil || list.head.value != 1 || list.head.next != nil {
		t.Errorf("Head = %v, want Value: 1, Next: nil", list.head.ToString())
	}

	if tail := list.GetTail(); tail != list.head {
		t.Errorf("Tail = %v, want %v", tail.ToString(), list.head.ToString())
	}
}

func Test_GivenLinkedList_WhenGetTail_ThenTailIsAppendedNode(t *testing.T) {
	list := NewLinkedList(NewLinkedListNode(1))
	expected := NewLinkedListNode(3)
	list.Append(expected)

	if length := list.GetLength(); length != 2 {
		t.Errorf("Length = %d, want 2", length)
	}

	if list.head == nil || list.head.value != 1 || list.head.next == nil {
		t.Errorf("Head = %v, want Value: 1, Next: nil", list.head.ToString())
	}

	if tail := list.GetTail(); tail != expected {
		t.Errorf("Tail = %v, want %v", tail.ToString(), expected.ToString())
	}
}

func Test_GivenLinkedList_WhenAppend_ThenNewNodeAdded(t *testing.T) {
	list := NewLinkedList(NewLinkedListNode(1))
	list.Append(NewLinkedListNode(2))

	if length := list.GetLength(); length != 2 {
		t.Errorf("Length = %d, want 1", length)
	}

	if tail := list.GetTail(); tail == nil || tail.value != 2 || tail.next != nil {
		t.Errorf("Tail = %v, want Value: 2, Next: nil", tail.ToString())
	}
}

func Test_GivenEmptyLinkedList_WhenAppend_ThenHeadIsSet(t *testing.T) {
	list := NewLinkedList[int](nil)
	list.Append(NewLinkedListNode(1))

	if list.head == nil || list.head.value != 1 || list.head.next != nil {
		t.Errorf("Head = %v, want Value = 1, Next = nil", list.head.ToString())
	}
}
