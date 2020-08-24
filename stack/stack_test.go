package stack

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestStack(t *testing.T) {
	var s Stack
	s = NewArrayStack(10)

	t.Run("new stack have length 0 and empty", func(t *testing.T) {
		// a whole new stack
		assert.Equal(t, s.Len(), 0)
		assert.Equal(t, s.Empty(), true)
		assert.Equal(t, s.Peek(), nil)
		assert.Equal(t, s.Pop(), nil)
	})

	t.Run("push integer", func(t *testing.T) {
		intData := 123
		s.Push(intData)
		assert.Equal(t, s.Peek(), intData)
		assert.Equal(t, s.Len(), 1)
		assert.Equal(t, s.Empty(), false)
	})

	t.Run("push string and pop it", func(t *testing.T) {
		stringData := "another string"
		s.Push(stringData)
		assert.Equal(t, s.Peek(), stringData)
		assert.Equal(t, s.Len(), 2, "length not equal", "another message")
		assert.Equal(t, s.Empty(), false)

		//pop
		assert.Equal(t, s.Pop(), stringData)
		assert.Equal(t, s.Len(), 1)
	})
}
