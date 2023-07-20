package structures

import "fmt"

type linkedListNode[T any] struct {
	value T
	next  *linkedListNode[T]
}

func NewLinkedListNode[T any](value T) *linkedListNode[T] {
	return &linkedListNode[T]{value: value, next: nil}
}

// Get the T value of the node
func (n *linkedListNode[T]) GetValue() T {
	return n.value
}

// Get the string representation of the node
func (n *linkedListNode[T]) ToString() string {
	return fmt.Sprintf("Value %v, Next: %v", n.value, n.next)
}

type linkedList[T any] struct {
	head *linkedListNode[T]
}

func NewLinkedList[T any](node *linkedListNode[T]) *linkedList[T] {
	return &linkedList[T]{node}
}

// Add the node to the end of the linked list
func (ll *linkedList[T]) Append(node *linkedListNode[T]) {
	tail := ll.head

	if tail == nil {
		ll.head = node
	} else {
		tail := ll.GetTail()
		tail.next = node
	}
}

// Fetch the last node in the linked list
func (ll *linkedList[T]) GetTail() *linkedListNode[T] {
	current := ll.head

	for current.next != nil {
		current = current.next
	}

	return current
}

// Get the length of the linked list
func (ll *linkedList[T]) GetLength() int {
	current := ll.head
	length := 0

	if current != nil {
		length = 1
	}

	for current.next != nil {
		length++
		current = current.next
	}

	return length
}
