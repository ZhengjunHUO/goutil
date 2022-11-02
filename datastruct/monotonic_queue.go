package datastruct

import (
	"golang.org/x/exp/constraints"
	"sync"
)

// Monotonic queue contains a double end queue of any type
// All the elements stored in the Dequeue is in descending order
type MonotonicQueue[V constraints.Ordered] struct {
	Deque *Deque[V]
	lock  sync.RWMutex
}

// NewMonotonicQueue creates an empty double end queue of any type
func NewMonotonicQueue[V constraints.Ordered]() *MonotonicQueue[V] {
	return &MonotonicQueue[V]{Deque: NewDeque[V]([]V{}), }
}

// Push pops up all elements smaller than the current element from the end of the Dequeue
// and inserts the current element at the end of the underlying Dequeue
func (mq *MonotonicQueue[V]) Push(elem V) {
	mq.lock.Lock()
	defer mq.lock.Unlock()

	for !mq.IsEmpty() && mq.Deque.PeekLast() < elem {
		mq.Deque.PopLast()
	}

	mq.Deque.PushLast(elem)
}

// Pop tries to remove the given element if this element is the max of the queue
// if not, this action is ignored
func (mq *MonotonicQueue[V]) Pop(elem V) {
	mq.lock.Lock()
	defer mq.lock.Unlock()

	if !mq.IsEmpty() {
		if elem == mq.Deque.PeekFirst() {
			mq.Deque.PopFirst()
		}
	}
}

// Max returns the max element of the queue, which is the first element of the underlying Deque
func (mq *MonotonicQueue[V]) Max() V {
	return mq.Deque.PeekFirst()
}

// IsEmpty checks if the monotonic queue is empty
func (mq *MonotonicQueue[V]) IsEmpty() bool {
	return mq.Deque.IsEmpty()
}

// Size returns the length of monotonic queue
func (mq *MonotonicQueue[V]) Size() int {
	return mq.Deque.Size()
}
