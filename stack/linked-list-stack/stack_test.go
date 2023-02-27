package linked_list_stack

import (
	collection "collections/stack"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Length(t *testing.T) {
	stack := New[int]()

	for i := 1; i < 10; i++ {
		stack.Push(i)
		assert.Equal(t, i, stack.Length())
	}

	for i := 9; i > 0; i-- {
		assert.Equal(t, i, stack.Length())
		_, err := stack.Pop()
		assert.Nil(t, err, "unexpected error")
	}
}

func TestStack_IsEmpty(t *testing.T) {
	stack := New[int]()

	assert.True(t, stack.IsEmpty())
	stack.Push(12)
	assert.False(t, stack.IsEmpty())
}

func TestStack_Push(t *testing.T) {
	stack := New[int]()

	stack.Push(1)
	assert.Equal(t, 1, stack.Length(), "exactly 1 element should be present")
	assert.False(t, stack.IsEmpty(), "stack should not be empty")

	stack.Push(2)
	assert.Equal(t, 2, stack.Length(), "exactly 2 elements should be present")
	assert.False(t, stack.IsEmpty(), "stack should not be empty")

	stack.Push(3)
	assert.Equal(t, 3, stack.Length(), "exactly 3 elements should be present")
	assert.False(t, stack.IsEmpty(), "stack should not be empty")

	stack.Push(4)
	assert.Equal(t, 4, stack.Length(), "exactly 4 elements should be present")
	assert.False(t, stack.IsEmpty(), "stack should not be empty")
}

func TestStack_Peek(t *testing.T) {
	stack := New[int]()

	_, err := stack.Peek()
	assert.ErrorIs(t, err, collection.StackEmptyError)

	stack.Push(2)
	stack.Push(3)
	stack.Push(5)
	stack.Push(8)

	assert.Equal(t, 4, stack.Length())

	v, err := stack.Peek()
	assert.Equal(t, 8, v)
	assert.Equal(t, 4, stack.Length())

	v, err = stack.Peek()
	assert.Equal(t, 8, v)
	assert.Equal(t, 4, stack.Length())
}

func TestStack_Pop(t *testing.T) {
	stack := New[string]()

	_, err := stack.Pop()
	assert.ErrorIs(t, err, collection.StackEmptyError)

	stack.Push("a")
	stack.Push("b")
	stack.Push("c")

	assert.Equal(t, 3, stack.Length())

	v, err := stack.Pop()
	assert.Equal(t, "c", v)
	assert.Nil(t, err)
	assert.Equal(t, 2, stack.Length())

	v, err = stack.Peek()
	assert.Equal(t, "b", v)

	v, err = stack.Pop()
	assert.Equal(t, "b", v)
	assert.Nil(t, err)
	assert.Equal(t, 1, stack.Length())

	v, err = stack.Peek()
	assert.Equal(t, "a", v)

	v, err = stack.Pop()
	assert.Equal(t, "a", v)
	assert.Nil(t, err)
	assert.Equal(t, 0, stack.Length())

	v, err = stack.Peek()
	assert.ErrorIs(t, err, collection.StackEmptyError)
}
