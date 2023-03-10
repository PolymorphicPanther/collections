package stack

import "errors"

var EmptyError = errors.New("stack empty")

type Stack[T any] interface {
	Pop() (T, error)
	Peek() (T, error)
	Push(t T)
	IsEmpty() bool
	Length() int
}
