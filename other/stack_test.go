package other

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_IsEmpty(t *testing.T) {
	s := &Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Pop()
	s.Pop()
	s.Pop()

	assert.True(t, s.IsEmpty())
}

func TestStack_PushPop(t *testing.T) {
	s := &Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	assert.Equal(t, 3, s.Pop())
	assert.Equal(t, 2, s.Pop())
	assert.Equal(t, 1, s.Pop())
}

func TestStack_Top(t *testing.T) {
	s := &Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Equal(t, 3, s.Top())
	assert.Equal(t, 3, s.Pop())
	assert.Equal(t, 2, s.Top())
}

func TestStack_EdgeCases(t *testing.T) {
	s := &Stack[int]{}
	s.Push(1)
	s.Pop()
	s.Pop()

	assert.Equal(t, 0, s.Pop())
	assert.Equal(t, 0, s.Pop())
	assert.Equal(t, 0, s.Top())
}
