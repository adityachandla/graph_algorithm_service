package evaluator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	st := NewStack[int]()
	st.Push(22)
	st.Push(11)
	val := st.Pop()
	assert.Equal(t, val, 11)
	assert.False(t, st.Empty())
	val = st.Pop()
	assert.Equal(t, 22, val)
	assert.True(t, st.Empty())
}
