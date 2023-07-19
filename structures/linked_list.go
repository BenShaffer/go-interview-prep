package structures

import "fmt"

type LinkedListNode[T any] struct {
	value T
	next  *LinkedListNode[T]
}

type LinkedList[T any] struct {
	head *LinkedListNode[T]
}

func NewLinkedListNode[T any](value T) *LinkedListNode[T] {
	return &LinkedListNode[T]{value: value, next: nil}
}

func (n *LinkedListNode[T]) GetValue() T {
	return n.value
}

func (n *LinkedListNode[T]) ToString() string {
	return fmt.Sprintf("Value %v, Next: %v", n.value, n.next)
}

func NewLinkedList[T any](node *LinkedListNode[T]) *LinkedList[T] {
	return &LinkedList[T]{node}
}

func (ll *LinkedList[T]) Append(node *LinkedListNode[T]) {
	tail := ll.head

	if tail == nil {
		ll.head = node
	} else {
		tail := ll.GetTail()
		tail.next = node
	}
}

func (ll *LinkedList[T]) GetTail() *LinkedListNode[T] {
	current := ll.head

	for current.next != nil {
		current = current.next
	}

	return current
}

func (ll *LinkedList[T]) GetLength() int {
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
