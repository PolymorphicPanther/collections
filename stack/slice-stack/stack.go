package slice_stack

import (
	"collections/stack"
)

var _ stack.Stack[int] = (*Stack[int])(nil)

func New[T any]() *Stack[T] {
	return &Stack[T]{
		slice: make([]T, 0),
	}
}

type Stack[T any] struct {
	slice []T
}

func (s *Stack[T]) Pop() (T, error) {
	if s.Length() == 0 {
		var empty T
		return empty, stack.EmptyError
	}

	e := s.slice[s.Length()-1]
	s.slice = s.slice[:s.Length()-1]

	return e, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.Length() == 0 {
		var empty T
		return empty, stack.EmptyError
	}

	e := s.slice[s.Length()-1]
	return e, nil
}

func (s *Stack[T]) Push(t T) {
	s.slice = append(s.slice, t)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Length() == 0
}

func (s *Stack[T]) Length() int {
	return len(s.slice)
}
