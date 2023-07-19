package structures

import "testing"

func Test_GivenQueue_WhenPush_ThenValueIsAdded(t *testing.T) {
	queue := NewQueue[int]()
	queue.Push(1)

	if length := queue.GetLength(); length != 1 {
		t.Errorf("Length = %d, want 1", length)
	}
}

func Test_GivenQueue_WhenPushPop_ThenCorrectValueReturned(t *testing.T) {
	queue := NewQueue[int]()
	queue.Push(1)
	value := queue.Pop()

	if length := queue.GetLength(); value != 1 || length != 0 {
		t.Errorf("Value = %v, Length = %d, want Value = 1, Length = 0", value, length)
	}
}

func Test_GivenQueue_WhenManyPushPop_ThenCorrectValueReturned(t *testing.T) {
	queue := NewQueue[int]()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	value := queue.Pop()

	if length := queue.GetLength(); value != 1 || length != 2 {
		t.Errorf("Value = %v, Length = %d, want Value = 1, Length = 0", value, length)
	}
}
