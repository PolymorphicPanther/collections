package linked_list_stack

import (
	"collections/stack"
	"container/list"
)

var _ stack.Stack[int] = (*Stack[int])(nil)

func New[T any]() *Stack[T] {

	l := list.New()
	l.Init()
	return &Stack[T]{
		list: l,
	}
}

type Stack[T any] struct {
	list *list.List
}

func (s *Stack[T]) Pop() (T, error) {
	if s.list.Back() == nil {
		var empty T
		return empty, stack.EmptyError
	}

	e := s.list.Back()
	s.list.Remove(e)

	return e.Value.(T), nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.list.Back() == nil {
		var empty T
		return empty, stack.EmptyError
	}

	e := s.list.Back()
	return e.Value.(T), nil
}

func (s *Stack[T]) Push(t T) {
	s.list.PushBack(t)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Length() == 0
}

func (s *Stack[T]) Length() int {
	return s.list.Len()
}
