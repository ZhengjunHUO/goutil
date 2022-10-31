package datastruct

import (
	"testing"
)

func TestMaxPriorityQueue(t *testing.T) {
	values := []int{500, 300, 1000}
	prios := []float64{3.0, 2.5, 6.1}
	pq := NewPQ(values, prios, false)

	val, prio := 600, 9.8
	pq.Push(val, prio)
	v, p := pq.PopWithPriority()
	if v != val || p != prio {
		t.Errorf("PopWithPriority returns (%v, %v), expect (%v, %v)\n", v, p, val, prio)
	}

	if pq.Peek() != 1000 {
		t.Errorf("Peek returns %v, expect 1000\n", pq.Peek())
	}

	if pq.Size() != len(values) {
		t.Errorf("Size returns %v, expect %v", pq.Size(), len(values))
	}

	for pq.Size() > 0 {
		pq.Pop()
	}

	if v := pq.Peek(); v != 0 {
		t.Errorf("Peek empty pq returns %v, expect 0", v)
	}

	if v := pq.Pop(); v != 0 {
		t.Errorf("Pop empty pq returns %v, expect 0", v)
	}

	if v, p := pq.PopWithPriority(); v != 0 || p != 0.0 {
		t.Errorf("PopWithPriority empty pq returns (%v, %v), expect (0, 0.0)", v, p)
	}
}

func TestMinPriorityQueue(t *testing.T) {
	values := []int{500, 300, 1000}
	prios := []float64{3.0, 2.5}
	if pq := NewPQ(values, prios, true); pq != nil {
		t.Errorf("NewPQ with different length of values & prios returns %v, expect nil\n", pq)
	}

	prios = []float64{3.0, 2.5, 6.1}
	pq := NewPQ(values, prios, true)

	val, prio := 600, 0.2
	pq.Push(val, prio)
	v, p := pq.PopWithPriority()
	if v != val || p != prio {
		t.Errorf("PopWithPriority returns (%v, %v), expect (%v, %v)\n", v, p, val, prio)
	}

	if v := pq.Remove(400); v != nil {
		t.Errorf("Remove inexist value returns %v, expect nil", v)
	}

	if v := pq.Remove(300); v == nil {
		t.Errorf("Remove exist value returns %v, expect not nil", v)
	}

	if pq.Peek() != 500 {
		t.Errorf("Peek returns %v, expect 500\n", pq.Peek())
	}
}
