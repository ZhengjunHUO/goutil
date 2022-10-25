package datastruct

import (
	"testing"
)

func TestQueue(t *testing.T) {
	// Test an int slice with New method
	intSlice := []int{3, 6, 9}
	intSliceLen := len(intSlice)
	intQ := NewQueue(intSlice)

	p := intQ.Pop()
	if p != intSlice[0] {
		t.Errorf("Pop returns %v, expect %v !\n", p, intSlice[intSliceLen-1])
	}

	p = intQ.Peek()
	if p != intSlice[1] {
		t.Errorf("Peek returns %v, expect %v !\n", p, intSlice[intSliceLen-2])
	}

	s := intQ.Size()
	if s != intSliceLen - 1 {
		t.Errorf("Size returns %v, expect %v !\n", s, intSliceLen-1)
	}

	// Test an empty string slice
	emptySlice := []string{}
	strQ := NewQueue(emptySlice)

	str := strQ.Pop()
	if str != "" {
		t.Errorf("Pop returns [%v], expect empty string !\n", str)
	}

	str = strQ.Peek()
	if str != "" {
		t.Errorf("Peek returns [%v], expect empty string !\n", str)
	}

	isEmpty := strQ.IsEmpty()
	if !isEmpty {
		t.Errorf("IsEmpty returns %v, expect true !\n", isEmpty)
	}
}
