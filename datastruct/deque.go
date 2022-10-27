package datastruct

import (
	"sync"
)

// Deque is a double end queue, contains a slice of any type
type Deque[V any] struct {
	Elems []V
	Lock  sync.RWMutex
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
	d.Lock.Lock()
	defer d.Lock.Unlock()

	d.Elems = append([]V{elem}, d.Elems...)
}

// PushLast appends a new element of any type at the end
func (d *Deque[V]) PushLast(elem V) {
	d.Lock.Lock()
	defer d.Lock.Unlock()

	d.Elems = append(d.Elems, elem)
}

// PopFirst returns the queue's first element and remove it from the queue
func (d *Deque[V]) PopFirst() V {
	var zeroValue V

	if d.IsEmpty() {
		return zeroValue
	}

	d.Lock.Lock()
	defer d.Lock.Unlock()

	defer func(){
		d.Elems[0] = zeroValue
		d.Elems = d.Elems[1:]
	}()

	return d.Elems[0]
}

// PopLast returns the queue's last element and remove it from the queue
func (d *Deque[V]) PopLast() V {
	var zeroValue V

	if d.IsEmpty() {
		return zeroValue
	}

	d.Lock.Lock()
	defer d.Lock.Unlock()

	n := len(d.Elems)

	defer func(){
		d.Elems[n-1] = zeroValue
		d.Elems = d.Elems[:n-1]
	}()

	return d.Elems[n-1]
}

// PeekFirst returns the queue's first element
func (d *Deque[V]) PeekFirst() V {
	var zeroValue V

	if d.IsEmpty() {
		return zeroValue
	}

	d.Lock.RLock()
	defer d.Lock.RUnlock()

	return d.Elems[0]
}

// PeekLast returns the queue's last element
func (d *Deque[V]) PeekLast() V {
	var zeroValue V

	if d.IsEmpty() {
		return zeroValue
	}

	d.Lock.RLock()
	defer d.Lock.RUnlock()

	return d.Elems[len(d.Elems)-1]
}

// IsEmpty checks if the double end queue is empty
func (d *Deque[V]) IsEmpty() bool {
	d.Lock.RLock()
	defer d.Lock.RUnlock()

	return len(d.Elems) == 0
}

// Size returns the length of double end queue
func (d *Deque[V]) Size() int {
	d.Lock.RLock()
	defer d.Lock.RUnlock()

	return len(d.Elems)
}
