package stack

import (
	"errors"
	"sync"
)

var (
	errEmptyStack = errors.New("Stack is empty")
)

// Stack ...
type Stack struct {
	stack []int
	lock  sync.RWMutex
}

// Push ...
func (s *Stack) Push(name int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.stack = append(s.stack, name)
}

// Pop ...
func (s *Stack) Pop() error {
	if s.Empty() {
		return errEmptyStack
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	len := len(s.stack)
	s.stack = s.stack[:len-1]

	return nil
}

// Front ...
func (s *Stack) Front() (int, error) {
	if s.Empty() {
		return 0, errEmptyStack
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	len := len(s.stack)
	return s.stack[len-1], nil
}

// Size ...
func (s *Stack) Size() int {
	return len(s.stack)
}

// Empty ...
func (s *Stack) Empty() bool {
	return len(s.stack) == 0
}
