package datastruct

import (
	"container/heap"
	"sync"

	"golang.org/x/exp/constraints"
)

// Elem contains a value, its priority
// and an index for heap.Interface methods
type Elem[T any, O constraints.Ordered] struct {
	value		T
	priority	O
	index		int
}

// Heap implements heap.Interface's methods except Less(i, j int) bool
type Heap[T any, O constraints.Ordered] []*Elem[T, O]

// MaxHeap embeds Heap and provides a Less method, thus implements heap.Interface
// it also provides a GetElems method, to satisfy the GenericHeap interface
type MaxHeap[T any, O constraints.Ordered] struct {
	*Heap[T, O]
}

// MinHeap embeds Heap and provides a Less method, thus implements heap.Interface
// it also provides a GetElems method, to satisfy the GenericHeap interface
type MinHeap[T any, O constraints.Ordered] struct {
	*Heap[T, O]
}

// GenericHeap contains a heap.Interface and a GetElems method
// MaxHeap and MinHeap implement this interface
type GenericHeap[T any, O constraints.Ordered] interface {
	heap.Interface
	GetElems() []*Elem[T, O]
}

// PriorityQueue adds the thread safe to the GenericHeap
type PriorityQueue[T any, O constraints.Ordered] struct {
	data		GenericHeap[T, O]
	lock		sync.RWMutex
}

// NewPQ builds a priority queue with a slice of value and priority
// popLowest decides its a min heap or max heap
func NewPQ[T any, O constraints.Ordered](values []T, prios []O, popLowest bool) *PriorityQueue[T, O] {
	nv := len(values)
	np := len(prios)

	if nv != np {
                return nil
	}

	var data GenericHeap[T, O]
	h := make(Heap[T, O], nv)

	if popLowest {
		data = MinHeap[T, O]{&h,}
	}else{
		data = MaxHeap[T, O]{&h,}
	}

	elems := data.GetElems()
	for i:=0; i<nv; i++ {
		elems[i] = &Elem[T, O]{
			value:		values[i],
			priority:	prios[i],
			index:		i,
		}
	}

	heap.Init(data)

	return &PriorityQueue[T, O]{
		data: data,
	}
}

// Len implements sort/heap interface
func (h Heap[T, O]) Len() int {
	return len(h)
}

// Swap implements sort/heap interface
func (h Heap[T, O]) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

// Push implements heap interface
func (h *Heap[T, O]) Push(item any) {
	elem := item.(*Elem[T, O])
	elem.index = h.Len()
	*h = append(*h, elem)
}

// Pop implements heap interface
func (h *Heap[T, O]) Pop() any {
	n := h.Len()

	defer func() {
		(*h)[n-1] = nil
		*h = (*h)[:n-1]
	}()

	return (*h)[n-1]
}

// Less Implement sort/heap interface for MinHeap
func (h MinHeap[T, O]) Less(i, j int) bool {
	return (*h.Heap)[i].priority < (*h.Heap)[j].priority
}

// GetElems returns elements store in the heap
func (h MinHeap[T, O]) GetElems() []*Elem[T, O] {
	return *(h.Heap)
}

// Less Implement sort/heap interface for MaxHeap
func (h MaxHeap[T, O]) Less(i, j int) bool {
	return (*h.Heap)[i].priority > (*h.Heap)[j].priority
}

// GetElems returns elements store in the heap
func (h MaxHeap[T, O]) GetElems() []*Elem[T, O] {
	return *(h.Heap)
}

// Push adds a element to the queue
func (pq *PriorityQueue[T, O]) Push(value T, prio O) {
	pq.lock.Lock()
	defer pq.lock.Unlock()

	elem := &Elem[T, O]{
		value:		value,
		priority:	prio,
	}
	heap.Push(pq.data, elem)
}

// Pop returns the min/max value of the priority queue
func (pq *PriorityQueue[T, O]) Pop() T {
	var zeroVal T
	if pq.data.Len() < 1 {
		return zeroVal
	}

	pq.lock.Lock()
	defer pq.lock.Unlock()

	return heap.Pop(pq.data).(*Elem[T, O]).value
}

// PopWithPriority returns the min/max value of the priority queue along with its priority
func (pq *PriorityQueue[T, O]) PopWithPriority() (T, O){
	var zeroVal T
	var zeroOrdVal O
        if pq.data.Len() < 1 {
                return zeroVal, zeroOrdVal
        }

        pq.lock.Lock()
        defer pq.lock.Unlock()

	elem := heap.Pop(pq.data).(*Elem[T, O])

        return elem.value, elem.priority
}

// Peek checks the min/max value of the priority queue
func (pq *PriorityQueue[T, O]) Peek() T {
	var zeroVal T
	if pq.data.Len() < 1 {
		return zeroVal
	}

	pq.lock.RLock()
	defer pq.lock.RUnlock()

	return pq.data.GetElems()[0].value
}

// Size returns the number of elements in the priority queue
func (pq *PriorityQueue[T, O]) Size() int {
	return pq.data.Len()
}
