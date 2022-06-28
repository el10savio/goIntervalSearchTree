package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Refactor To Test Suite

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

// TestPop ...
func TestPop(t *testing.T) {
	// PopTestCase ...
	type PopTestCase struct {
		name          string
		values        []int
		expectedValue int
		expectedStack []int
		expectedError error
	}
	testCases := []PopTestCase{
		PopTestCase{"no values", []int{}, 0, nil, errIsEmptyStack},
		PopTestCase{"single value", []int{1}, 1, []int{}, nil},
		PopTestCase{"two values", []int{1, 2}, 2, []int{1}, nil},
		PopTestCase{"multiple values", []int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4}, nil},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			stack := New()
			defer stack.Clear()

			// TODO: Refactor as test utility function
			for _, value := range testCase.values {
				stack.Push(value)
			}

			value, err := stack.Pop()

			assert.Equal(t, testCase.expectedValue, value)
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedStack, stack.List())
		})
	}
}

// TestFront ...
func TestFront(t *testing.T) {
	// FrontTestCase ...
	type FrontTestCase struct {
		name          string
		values        []int
		expectedValue int
		expectedStack []int
		expectedError error
	}
	testCases := []FrontTestCase{
		FrontTestCase{"no values", []int{}, 0, nil, errIsEmptyStack},
		FrontTestCase{"single value", []int{1}, 1, []int{1}, nil},
		FrontTestCase{"two values", []int{1, 2}, 2, []int{1, 2}, nil},
		FrontTestCase{"multiple values", []int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}, nil},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			stack := New()
			defer stack.Clear()

			// TODO: Refactor as test utility function
			for _, value := range testCase.values {
				stack.Push(value)
			}

			value, err := stack.Front()

			assert.Equal(t, testCase.expectedValue, value)
			assert.Equal(t, testCase.expectedError, err)
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
