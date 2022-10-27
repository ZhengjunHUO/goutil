package datastruct

import (
	"sync"
)

type Deque[V any] struct {
	Elems []V
	Lock  sync.RWMutex
}

func NewDeque[V any](slice []V) *Deque[V] {
	d := &Deque[V]{}

	for i := 0; i < len(slice); i++ {
		d.PushLast(slice[i])
	}

	return d
}

func (d *Deque[V]) PushFirst(elem V) {
	d.Lock.Lock()
	defer d.Lock.Unlock()

	d.Elems = append([]V{elem}, d.Elems...)
}

func (d *Deque[V]) PushLast(elem V) {
	d.Lock.Lock()
	defer d.Lock.Unlock()

	d.Elems = append(d.Elems, elem)
}

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

func (d *Deque[V]) PeekFirst() V {
	var zeroValue V

	if d.IsEmpty() {
		return zeroValue
	}

	d.Lock.RLock()
	defer d.Lock.RUnlock()

	return d.Elems[0]
}

func (d *Deque[V]) PeekLast() V {
	var zeroValue V

	if d.IsEmpty() {
		return zeroValue
	}

	d.Lock.RLock()
	defer d.Lock.RUnlock()

	return d.Elems[len(d.Elems)-1]
}

func (d *Deque[V]) IsEmpty() bool {
	d.Lock.RLock()
	defer d.Lock.RUnlock()

	return len(d.Elems) == 0
}

func (d *Deque[V]) Size() int {
	d.Lock.RLock()
	defer d.Lock.RUnlock()

	return len(d.Elems)
}
