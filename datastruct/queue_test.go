package datastruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	v, err := q.Pop()
	assert.Nil(t, err)
	assert.Equal(t, 1, v)
	assert.True(t, q.Empty())
}
