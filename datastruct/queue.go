package datastruct

import (
	"sync"
)

// Queue contains a slice of any type
type Queue[V any] struct {
	Elems []V
	lock  sync.RWMutex
}

// NewQueue creates a queue from a slice of any type
func NewQueue[V any](slice []V) *Queue[V] {
	q := &Queue[V]{}

	for i := 0; i < len(slice); i++ {
		q.Push(slice[i])
	}

	return q
}

// Push appends a new element of any type to the end
func (q *Queue[V]) Push(elem V) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.Elems = append(q.Elems, elem)
}

// Pop returns the queue's head and remove it from the queue
func (q *Queue[V]) Pop() V {
	var zeroVal V

	if q.IsEmpty() {
		return zeroVal
	}

	q.lock.Lock()
	defer q.lock.Unlock()

	defer func(){
		q.Elems[0] = zeroVal
		q.Elems = q.Elems[1:]
	}()

	return q.Elems[0]
}

// Peek returns the queue's head
func (q *Queue[V]) Peek() V {
	var zeroVal V

	if q.IsEmpty() {
		return zeroVal
	}

	q.lock.RLock()
	defer q.lock.RUnlock()

	return q.Elems[0]
}

// IsEmpty checks if the queue is empty
func (q *Queue[V]) IsEmpty() bool {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return len(q.Elems) == 0
}

// Size returns the length of queue
func (q *Queue[V]) Size() int {
	q.lock.RLock()
	defer q.lock.RUnlock()

	return len(q.Elems)
}
