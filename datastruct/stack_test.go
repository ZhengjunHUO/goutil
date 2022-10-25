package datastruct

import (
	"testing"
)

func TestStack(t *testing.T) {
	// Test an int slice with New method
	intSlice := []int{3, 6, 9}
	intSliceLen := len(intSlice)
	intS := NewStack(intSlice)

	p := intS.Pop()
	if p != intSlice[intSliceLen-1] {
		t.Errorf("Pop returns %v, expect %v !\n", p, intSlice[intSliceLen-1])
	}

	p = intS.Peek()
	if p != intSlice[intSliceLen-2] {
		t.Errorf("Peek returns %v, expect %v !\n", p, intSlice[intSliceLen-2])
	}

	s := intS.Size()
	if s != intSliceLen - 1 {
		t.Errorf("Size returns %v, expect %v !\n", s, intSliceLen-1)
	}

	// Test an empty string slice
	emptySlice := []string{}
	strS := NewStack(emptySlice)

	str := strS.Pop()
	if str != "" {
		t.Errorf("Pop returns [%v], expect empty string !\n", str)
	}

	str = strS.Peek()
	if str != "" {
		t.Errorf("Peek returns [%v], expect empty string !\n", str)
	}

	isEmpty := strS.IsEmpty()
	if !isEmpty {
		t.Errorf("IsEmpty returns %v, expect true !\n", isEmpty)
	}
}
