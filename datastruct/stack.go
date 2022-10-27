package datastruct

import (
	"sync"
)

type Stack[V any] struct {
	elems	[]V
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

	s.elems = append(s.elems, elem)
}

func (s *Stack[V]) Pop() V {
	var zeroVal V

	if s.IsEmpty() {
		return zeroVal
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	n := len(s.elems)

	defer func(){
		s.elems[n-1] = zeroVal
		s.elems = s.elems[:n-1]
	}()

	return s.elems[n-1]
}

func (s *Stack[V]) Peek() V {
	var zeroVal V

	if s.IsEmpty() {
		return zeroVal
	}

	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.elems[len(s.elems)-1]
}

func (s *Stack[V]) IsEmpty() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return len(s.elems) == 0
}

func (s *Stack[V]) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return len(s.elems)
}
