package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPush ...
func TestPush(t *testing.T) {
	// PushTestCase ...
	type PushTestCase struct {
		name          string
		values        []int
		expectedStack []int
	}
	testCases := []PushTestCase{
		PushTestCase{"no values", []int{}, nil},
		PushTestCase{"single value", []int{1}, []int{1}},
		PushTestCase{"two values", []int{1, 2}, []int{1, 2}},
		PushTestCase{"multiple values", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			stack := New()
			defer stack.Clear()

			// TODO: Refactor as test utility function
			for _, value := range testCase.values {
				stack.Push(value)
			}

			assert.Equal(t, testCase.expectedStack, stack.List())
		})
	}
}

// TestSize ...
func TestSize(t *testing.T) {
	// SizeTestCase ...
	type SizeTestCase struct {
		name         string
		values       []int
		expectedSize int
	}
	testCases := []SizeTestCase{
		SizeTestCase{"no values", []int{}, 0},
		SizeTestCase{"single value", []int{1}, 1},
		SizeTestCase{"two values", []int{1, 2}, 2},
		SizeTestCase{"multiple values", []int{1, 2, 3, 4, 5}, 5},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			stack := New()
			defer stack.Clear()

			for _, value := range testCase.values {
				stack.Push(value)
			}

			assert.Equal(t, testCase.expectedSize, stack.Size())
		})
	}
}
