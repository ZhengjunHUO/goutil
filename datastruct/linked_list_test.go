package datastruct

import (
	"testing"
	"reflect"
)

func TestLinkedList(t *testing.T) {
	// Test an int slice with New method
	intSlice := []int{3,2,0,-4}

	intList := NewLinkedList(intSlice)
	if !reflect.DeepEqual(intList.ToSlice(), intSlice) {
		t.Errorf("Expect %v, but get slice %v\n", intSlice, intList.ToSlice())
	}
	intList.PrintAll()

	// Test a string slice, built from scratch
	strSlice := []string{"foo", "bar", "fufu"}

	strList := &LinkedList[string]{}
	strList.Push("foo")
	strList.Push("bar")
	strList.Push("fufu")

	if !reflect.DeepEqual(strList.ToSlice(), strSlice) {
		t.Errorf("Expect %v, but get slice %v\n", strSlice, strList.ToSlice())
	}

	// Test corner case, with an empty slice of any type
	emptySlice := []any{}
	emptyList := NewLinkedList(emptySlice)
	if len(emptyList.ToSlice()) != 0 {
		t.Errorf("Expect %v, but get slice %v\n", emptySlice, emptyList.ToSlice())
	}
}
