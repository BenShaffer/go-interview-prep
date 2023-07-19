package structures

import "fmt"

type Queue[T any] struct {
	values []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{values: []T{}}
}

func (q *Queue[T]) Push(value T) {
	q.values = append(q.values, value)
}

func (q *Queue[T]) Pop() T {
	value := q.values[0]
	q.values = q.values[1:]
	return value
}

func (q *Queue[T]) ToString() string {
	queueString := ""

	for i := 0; i < len(q.values); i++ {
		queueString += fmt.Sprintf("%v, ", q.values[i])
	}

	return queueString
}

func (q *Queue[T]) GetLength() int {
	return len(q.values)
}
