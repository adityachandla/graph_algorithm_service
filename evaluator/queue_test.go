package evaluator_test

import (
	"testing"

	"github.com/adityachandla/graph_algorithm_service/evaluator"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := evaluator.NewQueue[int]()
	q.AddToFront(22)
	q.AddToFront(33)
	assert.False(t, q.Empty())
	val := q.PopBack()
	assert.Equal(t, 22, val)
	assert.False(t, q.Empty())
	val = q.PopBack()
	assert.Equal(t, 33, val)
	assert.True(t, q.Empty())
}

func TestMultipleAdds(t *testing.T) {
	q := evaluator.NewQueue[int]()
	q.AddToFront(22)
	assert.False(t, q.Empty())
	val := q.PopBack()
	assert.Equal(t, 22, val)

	q.AddToFront(33)
	assert.False(t, q.Empty())
	val = q.PopBack()
	assert.Equal(t, 33, val)
	assert.True(t, q.Empty())
}
