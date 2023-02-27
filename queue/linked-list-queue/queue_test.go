package linked_list_queue

import (
	"collections/queue"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_FIFO(t *testing.T) {
	q := New[byte]()

	q.Add('A')
	q.Add('B')
	q.Add('C')

	v, err := q.Remove()
	assert.Nil(t, err)
	assert.Equal(t, byte('A'), v)

	v, err = q.Remove()
	assert.Nil(t, err)
	assert.Equal(t, byte('B'), v)

	v, err = q.Remove()
	assert.Nil(t, err)
	assert.Equal(t, byte('C'), v)
}

func TestQueue_Length(t *testing.T) {
	q := New[int]()
	assert.Equal(t, 0, q.Length())

	for i := 1; i < 10; i++ {
		q.Add(i)
		assert.Equal(t, i, q.Length())
	}

	for i := 1; i < 10; i++ {
		_, err := q.Remove()
		assert.Nil(t, err, "unexpected error")
		assert.Equal(t, 9-i, q.Length())
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	q := New[string]()
	assert.True(t, q.IsEmpty())

	q.Add("Do'h")
	assert.False(t, q.IsEmpty())
}

func TestQueue_Add(t *testing.T) {
	q := New[int]()

	q.Add(1)
	assert.Equal(t, 1, q.Length(), "exactly 1 element should be present")
	assert.False(t, q.IsEmpty(), "queue should not be empty")

	q.Add(2)
	assert.Equal(t, 2, q.Length(), "exactly 2 elements should be present")
	assert.False(t, q.IsEmpty(), "queue should not be empty")

	q.Add(3)
	assert.Equal(t, 3, q.Length(), "exactly 3 elements should be present")
	assert.False(t, q.IsEmpty(), "queue should not be empty")

	q.Add(4)
	assert.Equal(t, 4, q.Length(), "exactly 4 elements should be present")
	assert.False(t, q.IsEmpty(), "queue should not be empty")
}

func TestQueue_Peek(t *testing.T) {
	q := New[int]()

	_, err := q.Peek()
	assert.ErrorIs(t, err, queue.EmptyError)

	q.Add(2)
	q.Add(3)
	q.Add(5)
	q.Add(8)

	assert.Equal(t, 4, q.Length())

	v, err := q.Peek()
	assert.Equal(t, 2, v)
	assert.Equal(t, 4, q.Length())

	v, err = q.Peek()
	assert.Equal(t, 2, v)
	assert.Equal(t, 4, q.Length())
}

func TestQueue_Remove(t *testing.T) {
	q := New[string]()

	_, err := q.Remove()
	assert.ErrorIs(t, err, queue.EmptyError)

	q.Add("a")
	q.Add("b")
	q.Add("c")

	assert.Equal(t, 3, q.Length())

	v, err := q.Remove()
	assert.Equal(t, "a", v)
	assert.Nil(t, err)
	assert.Equal(t, 2, q.Length())

	v, err = q.Peek()
	assert.Equal(t, "b", v)

	v, err = q.Remove()
	assert.Equal(t, "b", v)
	assert.Nil(t, err)
	assert.Equal(t, 1, q.Length())

	v, err = q.Peek()
	assert.Equal(t, "c", v)

	v, err = q.Remove()
	assert.Equal(t, "c", v)
	assert.Nil(t, err)
	assert.Equal(t, 0, q.Length())

	v, err = q.Peek()
	assert.ErrorIs(t, err, queue.EmptyError)
}
