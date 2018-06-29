package stack_test

import (
	"testing"

	"github.com/puxin71/DesignPatternInGo/stack"
	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	myStack := stack.NewStack(0)
	assert.Equal(t, 0, myStack.Size(), "null stack")
	myStack = stack.NewStack(10)
	assert.Equal(t, 10, myStack.Size(), "stack with 10 bytes")
}

func TestPush(t *testing.T) {
	myStack := stack.NewStack(10)
	myString := []byte("hello world")
	for _, code := range myString {
		myStack.Push(code)
	}
	reversedString := make([]byte, 10)
	for i := 0; i < 10; i++ {
		reversedString[i] = myStack.Pop()
	}
	assert.Equal(t, "lrow olleh", string(reversedString), "done reversing")
}
