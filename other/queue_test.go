package other

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueOnStack(t *testing.T) {
	qs := NewQueueOnStack[int]()

	qs.Enqueue(1)
	qs.Enqueue(2)
	qs.Enqueue(3)

	assert.Equal(t, 1, qs.Dequeue())
	assert.Equal(t, 2, qs.Dequeue())
	assert.Equal(t, 3, qs.Dequeue())

	qs.Enqueue(4)
	qs.Enqueue(5)
	qs.Enqueue(6)

	assert.Equal(t, 4, qs.Dequeue())
	assert.Equal(t, 5, qs.Dequeue())
	assert.Equal(t, 6, qs.Dequeue())

}
