package datastruct

import (
	"testing"
	"reflect"
)

func TestLinkedList(t *testing.T) {
	intSlice := []int{3,2,0,-4}
	strSlice := []string{"foo", "bar", "fufu"}
	emptySlice := []any{}

	intList := NewLinkedList(intSlice)
	if !reflect.DeepEqual(intList.ToSlice(), intSlice) {
		t.Errorf("Expect %v, but get slice %v\n", intSlice, intList.ToSlice())
	}
	intList.PrintAll()

	strList := NewLinkedList(strSlice)
	if !reflect.DeepEqual(strList.ToSlice(), strSlice) {
		t.Errorf("Expect %v, but get slice %v\n", strSlice, strList.ToSlice())
	}

	emptyList := NewLinkedList(emptySlice)
	if len(emptyList.ToSlice()) != 0 {
		t.Errorf("Expect %v, but get slice %v\n", emptySlice, emptyList.ToSlice())
	}
}
