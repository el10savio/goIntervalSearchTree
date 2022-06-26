package stack

import (
	"errors"
	"sync"
)

var (
	errEmptyStack = errors.New("stack is empty")
)

// Stack ...
type Stack struct {
	stack []int
	lock  sync.RWMutex
}

// New ...
func New() *Stack {
	return &Stack{}
}

// Push ...
func (s *Stack) Push(value int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.stack = append(s.stack, value)
}

// Pop ...
func (s *Stack) Pop() (int, error) {
	if s.Empty() {
		return 0, errEmptyStack
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	len := len(s.stack)
	elem := s.stack[len-1]
	s.stack = s.stack[:len-1]

	return elem, nil
}

// Size ...
func (s *Stack) Size() int {
	return len(s.stack)
}

// Empty ...
func (s *Stack) Empty() bool {
	return len(s.stack) == 0
}
