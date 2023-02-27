package slice_stack

import (
	collection "collections/stack"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_LIFO(t *testing.T) {
	s := New[string]()

	s.Push("A")
	s.Push("B")
	s.Push("C")

	v, err := s.Pop()
	assert.Nil(t, err)
	assert.Equal(t, "C", v)

	v, err = s.Pop()
	assert.Nil(t, err)
	assert.Equal(t, "B", v)

	v, err = s.Pop()
	assert.Nil(t, err)
	assert.Equal(t, "A", v)
}

func TestStack_Length(t *testing.T) {
	s := New[int]()

	for i := 1; i < 10; i++ {
		s.Push(i)
		assert.Equal(t, i, s.Length())
	}

	for i := 9; i > 0; i-- {
		assert.Equal(t, i, s.Length())
		_, err := s.Pop()
		assert.Nil(t, err, "unexpected error")
	}
}

func TestStack_IsEmpty(t *testing.T) {
	s := New[int]()

	assert.True(t, s.IsEmpty())
	s.Push(12)
	assert.False(t, s.IsEmpty())
}

func TestStack_Push(t *testing.T) {
	s := New[int]()

	s.Push(1)
	assert.Equal(t, 1, s.Length(), "exactly 1 element should be present")
	assert.False(t, s.IsEmpty(), "stack should not be empty")

	s.Push(2)
	assert.Equal(t, 2, s.Length(), "exactly 2 elements should be present")
	assert.False(t, s.IsEmpty(), "stack should not be empty")

	s.Push(3)
	assert.Equal(t, 3, s.Length(), "exactly 3 elements should be present")
	assert.False(t, s.IsEmpty(), "stack should not be empty")

	s.Push(4)
	assert.Equal(t, 4, s.Length(), "exactly 4 elements should be present")
	assert.False(t, s.IsEmpty(), "stack should not be empty")
}

func TestStack_Peek(t *testing.T) {
	s := New[int]()

	_, err := s.Peek()
	assert.ErrorIs(t, err, collection.EmptyError)

	s.Push(2)
	s.Push(3)
	s.Push(5)
	s.Push(8)

	assert.Equal(t, 4, s.Length())

	v, err := s.Peek()
	assert.Equal(t, 8, v)
	assert.Equal(t, 4, s.Length())

	v, err = s.Peek()
	assert.Equal(t, 8, v)
	assert.Equal(t, 4, s.Length())
}

func TestStack_Pop(t *testing.T) {
	s := New[string]()

	_, err := s.Pop()
	assert.ErrorIs(t, err, collection.EmptyError)

	s.Push("a")
	s.Push("b")
	s.Push("c")

	assert.Equal(t, 3, s.Length())

	v, err := s.Pop()
	assert.Equal(t, "c", v)
	assert.Nil(t, err)
	assert.Equal(t, 2, s.Length())

	v, err = s.Peek()
	assert.Equal(t, "b", v)

	v, err = s.Pop()
	assert.Equal(t, "b", v)
	assert.Nil(t, err)
	assert.Equal(t, 1, s.Length())

	v, err = s.Peek()
	assert.Equal(t, "a", v)

	v, err = s.Pop()
	assert.Equal(t, "a", v)
	assert.Nil(t, err)
	assert.Equal(t, 0, s.Length())

	v, err = s.Peek()
	assert.ErrorIs(t, err, collection.EmptyError)
}
