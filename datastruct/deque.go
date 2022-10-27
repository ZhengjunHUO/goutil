package datastruct

import (
	"sync"
)

// Deque is a double end queue, contains a slice of any type
type Deque[V any] struct {
	elems []V
	lock  sync.RWMutex
}

// NewDeque creates a double end queue from a slice of any type
func NewDeque[V any](slice []V) *Deque[V] {
	d := &Deque[V]{}

	for i := 0; i < len(slice); i++ {
		d.PushLast(slice[i])
	}

	return d
}

// PushFirst inserts a new element of any type at the begining
func (d *Deque[V]) PushFirst(elem V) {
	d.lock.Lock()
	defer d.lock.Unlock()

	d.elems = append([]V{elem}, d.elems...)
}

// PushLast appends a new element of any type at the end
func (d *Deque[V]) PushLast(elem V) {
	d.lock.Lock()
	defer d.lock.Unlock()

	d.elems = append(d.elems, elem)
}

// PopFirst returns the queue's first element and remove it from the queue
func (d *Deque[V]) PopFirst() V {
	var zeroValue V

	if d.IsEmpty() {
		return zeroValue
	}

	d.lock.Lock()
	defer d.lock.Unlock()

	defer func(){
		d.elems[0] = zeroValue
		d.elems = d.elems[1:]
	}()

	return d.elems[0]
}

// PopLast returns the queue's last element and remove it from the queue
func (d *Deque[V]) PopLast() V {
	var zeroValue V

	if d.IsEmpty() {
		return zeroValue
	}

	d.lock.Lock()
	defer d.lock.Unlock()

	n := len(d.elems)

	defer func(){
		d.elems[n-1] = zeroValue
		d.elems = d.elems[:n-1]
	}()

	return d.elems[n-1]
}

// PeekFirst returns the queue's first element
func (d *Deque[V]) PeekFirst() V {
	var zeroValue V

	if d.IsEmpty() {
		return zeroValue
	}

	d.lock.RLock()
	defer d.lock.RUnlock()

	return d.elems[0]
}

// PeekLast returns the queue's last element
func (d *Deque[V]) PeekLast() V {
	var zeroValue V

	if d.IsEmpty() {
		return zeroValue
	}

	d.lock.RLock()
	defer d.lock.RUnlock()

	return d.elems[len(d.elems)-1]
}

// IsEmpty checks if the double end queue is empty
func (d *Deque[V]) IsEmpty() bool {
	d.lock.RLock()
	defer d.lock.RUnlock()

	return len(d.elems) == 0
}

// Size returns the length of double end queue
func (d *Deque[V]) Size() int {
	d.lock.RLock()
	defer d.lock.RUnlock()

	return len(d.elems)
}
