package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPush ...
func TestPush(t *testing.T) {
	for _, testCase := range pushTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			stack := New()
			defer stack.Clear()

			stack = pushValues(stack, testCase.values)

			assert.Equal(t, testCase.expectedStack, stack.List())
		})
	}
}

// TestPop ...
func TestPop(t *testing.T) {
	for _, testCase := range popTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			stack := New()
			defer stack.Clear()

			stack = pushValues(stack, testCase.values)

			value, err := stack.Pop()

			assert.Equal(t, testCase.expectedValue, value)
			assert.ErrorIs(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedStack, stack.List())
		})
	}
}

// TestFront ...
func TestFront(t *testing.T) {
	for _, testCase := range frontTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			stack := New()
			defer stack.Clear()

			stack = pushValues(stack, testCase.values)

			value, err := stack.Front()

			assert.Equal(t, testCase.expectedValue, value)
			assert.ErrorIs(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedStack, stack.List())
		})
	}
}

// TestSize ...
func TestSize(t *testing.T) {
	for _, testCase := range sizeTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			stack := New()
			defer stack.Clear()

			stack = pushValues(stack, testCase.values)

			assert.Equal(t, testCase.expectedSize, stack.Size())
		})
	}
}

// pushValues ...
func pushValues(stack *Stack, values []int) *Stack {
	for _, value := range values {
		stack.Push(value)
	}
	return stack
}
