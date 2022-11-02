package datastruct

import (
	"sync"
)

type Stack[V any] struct {
	Elems	[]V
	lock	sync.RWMutex
}

func NewStack[V any](slice []V) *Stack[V] {
	s := &Stack[V]{}

	for i := 0; i < len(slice); i++ {
		s.Push(slice[i])
	}

	return s
}

func (s *Stack[V]) Push(elem V) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.Elems = append(s.Elems, elem)
}

func (s *Stack[V]) Pop() V {
	var zeroVal V

	if s.IsEmpty() {
		return zeroVal
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	n := len(s.Elems)

	defer func(){
		s.Elems[n-1] = zeroVal
		s.Elems = s.Elems[:n-1]
	}()

	return s.Elems[n-1]
}

func (s *Stack[V]) Peek() V {
	var zeroVal V

	if s.IsEmpty() {
		return zeroVal
	}

	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.Elems[len(s.Elems)-1]
}

func (s *Stack[V]) IsEmpty() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return len(s.Elems) == 0
}

func (s *Stack[V]) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return len(s.Elems)
}
