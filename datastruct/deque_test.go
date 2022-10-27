package datastruct

import (
	"testing"
)

func TestDequeue(t *testing.T) {
	// Test a string slice with New method
	intStr := []string{"foo", "baz", "bar"}
	d := NewDeque(intStr)

	str := d.PopLast()
	if str != "bar" {
		t.Errorf("PopLast returns %v, expecting %v", str, "bar")
	}

	str = d.PeekLast()
	if str != "baz" {
		t.Errorf("PeekLast returns %v, expecting %v", str, "baz")
	}

	str = d.PopFirst()
	if str != "foo" {
		t.Errorf("PopFirst returns %v, expecting %v", str, "foo")
	}

	if d.IsEmpty() {
		t.Errorf("IsEmpty returns %v, expecting %v", d.IsEmpty(), "false")
	}

	// Test an empty int slice
	intSlice := []int{}
	dq := NewDeque(intSlice)

	if dq.PopFirst() != 0 {
		t.Errorf("Empty dequeue's PopFirst returns non zero-value element")
	}

	if dq.PopLast() != 0 {
		t.Errorf("Empty dequeue's PopLast returns non zero-value element")
	}

	i := dq.PeekFirst()
	if i != 0 {
		t.Errorf("PeekFirst returns %v, expecting %v", i, 0)
	}

	i = dq.PeekLast()
	if i != 0 {
		t.Errorf("PeekLast returns %v, expecting %v", i, 0)
	}

	dq.PushFirst(1)
	dq.PushFirst(2)

	i = dq.PeekFirst()
	if i != 2 {
		t.Errorf("PeekFirst returns %v, expecting %v", i, 2)
	}

	if dq.Size() != 2 {
		t.Errorf("Size returns %v, expecting %v", dq.Size(), 2)
	}
}
