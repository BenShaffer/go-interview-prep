package structures

import "fmt"

type queue[T any] struct {
	values []T
}

func NewQueue[T any]() *queue[T] {
	return &queue[T]{values: []T{}}
}

// Push a T element to the back of the queue
func (q *queue[T]) Push(value T) {
	q.values = append(q.values, value)
}

// Pop a T element from the front of the queue
func (q *queue[T]) Pop() T {
	value := q.values[0]
	q.values = q.values[1:]
	return value
}

// Get the string representation of the queue
func (q *queue[T]) ToString() string {
	queueString := ""

	for i := 0; i < len(q.values); i++ {
		queueString += fmt.Sprintf("%v, ", q.values[i])
	}

	return queueString
}

// Get the length of the queue
func (q *queue[T]) GetLength() int {
	return len(q.values)
}
