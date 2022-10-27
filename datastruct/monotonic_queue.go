package datastruct

import (
	"golang.org/x/exp/constraints"
	"sync"
)

// Monotonic queue contains a double end queue of any type
// All the elements stored in the dequeue is in descending order
type MonotonicQueue[V constraints.Ordered] struct {
	deque *Deque[V]
	lock  sync.RWMutex
}

// NewMonotonicQueue creates an empty double end queue of any type
func NewMonotonicQueue[V constraints.Ordered]() *MonotonicQueue[V] {
	return &MonotonicQueue[V]{deque: NewDeque[V]([]V{}), }
}

// Push pops up all elements smaller than the current element from the end of the dequeue
// and inserts the current element at the end of the underlying dequeue
func (mq *MonotonicQueue[V]) Push(elem V) {
	mq.lock.Lock()
	defer mq.lock.Unlock()

	for !mq.IsEmpty() && mq.deque.PeekLast() < elem {
		mq.deque.PopLast()
	}

	mq.deque.PushLast(elem)
}

// Pop tries to remove the given element if this element is the max of the queue
// if not, this action is ignored  
func (mq *MonotonicQueue[V]) Pop(elem V) {
	mq.lock.Lock()
	defer mq.lock.Unlock()

	if !mq.IsEmpty() {
		if elem == mq.deque.PeekFirst() {
			mq.deque.PopFirst()
		}
	}
}

// Max returns the max element of the queue, which is the first element of the underlying deque
func (mq *MonotonicQueue[V]) Max() V {
	return mq.deque.PeekFirst()
}

// IsEmpty checks if the monotonic queue is empty
func (mq *MonotonicQueue[V]) IsEmpty() bool {
	return mq.deque.IsEmpty()
}

// Size returns the length of monotonic queue
func (mq *MonotonicQueue[V]) Size() int {
	return mq.deque.Size()
}
