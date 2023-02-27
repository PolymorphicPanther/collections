package linked_list_queue

import (
	"collections/queue"
	"container/list"
)

var _ queue.Queue[any] = (*Queue[any])(nil)

func New[T any]() *Queue[T] {

	l := list.New()
	l.Init()
	return &Queue[T]{
		list: l,
	}
}

type Queue[T any] struct {
	list *list.List
}

func (q *Queue[T]) Add(t T) {
	q.list.PushFront(t)
}

func (q *Queue[T]) Remove() (T, error) {
	if q.IsEmpty() {
		var empty T
		return empty, queue.EmptyError
	}

	e := q.list.Remove(q.list.Back())
	return e.(T), nil
}

func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var empty T
		return empty, queue.EmptyError
	}

	e := q.list.Back()
	return (e.Value).(T), nil
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Length() == 0
}

func (q *Queue[T]) Length() int {
	return q.list.Len()
}
