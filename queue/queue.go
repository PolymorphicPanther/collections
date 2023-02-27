package queue

import "errors"

var EmptyError = errors.New("queue empty")

type Queue[T any] interface {
	Add(t T)
	Remove() (T, error)
	Peek() (T, error)
	IsEmpty() bool
	Length() int
}
