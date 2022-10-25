package datastruct

import (
	"sync"
)

type Stack[V any] struct {
	Elems	[]V
	Lock	sync.RWMutex
}

func NewStack[V any](slice []V) *Stack[V] {
	s := &Stack[V]{}

	for i := 0; i < len(slice); i++ {
		s.Push(slice[i])
	}

	return s
}

func (s *Stack[V]) Push(elem V) {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	s.Elems = append(s.Elems, elem)
}

func (s *Stack[V]) Pop() V {
	var zeroVal V

	if s.IsEmpty() {
		return zeroVal
	}

	s.Lock.Lock()
	defer s.Lock.Unlock()

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

	s.Lock.RLock()
	defer s.Lock.RUnlock()

	return s.Elems[len(s.Elems)-1]
}

func (s *Stack[V]) IsEmpty() bool {
	s.Lock.RLock()
	defer s.Lock.RUnlock()

	return len(s.Elems) == 0
}

func (s *Stack[V]) Size() int {
	s.Lock.RLock()
	defer s.Lock.RUnlock()

	return len(s.Elems)
}
