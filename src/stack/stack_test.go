package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	var S Stack = newLHStack()
	assert.Equal(t, true, S.Empty())
	S.Push("hi")
	S.Push("hey")
	assert.Equal(t, false, S.Empty())
	val := S.Pop()
	assert.Equal(t, "hey", val)
	tv := S.Top()
	assert.Equal(t, "hi", tv)
}
