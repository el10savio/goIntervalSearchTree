package stack

import (
	"errors"
	"sync"
)

var (
	errIsEmptyStack = errors.New("stack is empty")
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

// Clear ...
func (s *Stack) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.stack = nil
}

// Push ...
func (s *Stack) Push(value int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.stack = append(s.stack, value)
}

// Pop ...
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errIsEmptyStack
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	len := len(s.stack)
	elem := s.stack[len-1]
	s.stack = s.stack[:len-1]

	return elem, nil
}

// Front ...
func (s *Stack) Front() (int, error) {
	if s.IsEmpty() {
		return 0, errIsEmptyStack
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	len := len(s.stack)
	elem := s.stack[len-1]

	return elem, nil
}

// Size ...
func (s *Stack) Size() int {
	return len(s.stack)
}

// List ...
func (s *Stack) List() []int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.stack
}

// IsEmpty ...
func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}
