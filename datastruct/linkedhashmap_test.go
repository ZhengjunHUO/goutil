package datastruct

import (
	"fmt"
	"testing"
)

func TestLinkedHashmap(t *testing.T) {
	emptyLhm := NewLinkedHashmap()

	if emptyLhm.List.Pop() != nil {
		t.Errorf("Expect nil when poping from empty list.")
	}

	if emptyLhm.PopEldest() != nil {
		t.Errorf("Expect nil when poping from empty list.")
	}

	/*
	// Add
	emptyLhm.Put(2, "two")
	emptyLhm.Put(3, "three")
	emptyLhm.Put(5, "five")
	emptyLhm.Put(7, "seven")
	*/

	lhm := NewLinkedHashmapFromKV([][2]interface{}{
		[2]interface{}{2, "two"},
		[2]interface{}{3, "three"},
		[2]interface{}{5, "five"},
		[2]interface{}{7, "seven"},
	})

	// Read through the ordered hashmap
	list := ""
	expectedList := "two three five seven "
	for elem := range lhm.IntoIter() {
		list += fmt.Sprintf("%v ", elem.Val)
	}

	if list != expectedList {
		t.Errorf("Expect %v, but got %v", expectedList, list)
	}

	// Update existing value
	cinq := "cinq"
	lhm.Put(5, cinq)
	if lhm.Get(5) != cinq {
		t.Errorf("Expect %v, but got %v", cinq, lhm.Get(5))
	}

	// Delete non existing value
	size := lhm.Size()
	lhm.Delete(11)
	if lhm.Size() != size {
		t.Errorf("Expect %v, but got %v", size, lhm.Size())
	}

	// Delete existing value
	if !lhm.Contains(3) {
		t.Errorf("The linked hashmap should contain 3")
	}

	lhm.Delete(3)
	if lhm.Get(3) != nil {
		t.Errorf("Expect nil, but got %v", lhm.Get(11))
	}

	if lhm.Contains(3) {
		t.Errorf("The linked hashmap should not contain 3 after delete")
	}

	// Get eldest
	two := "two"
	eldest := lhm.PopEldest()
	if eldest.Val != two {
		t.Errorf("Expect %v, but got %v", two, eldest.Val)
	}

	// Move node to the end
	lhm.BecomeNewest(4)
	lhm.BecomeNewest(5)
	if lhm.List.Tail.Prev.Val != cinq {
		t.Errorf("Expect %v, but got %v", cinq, lhm.List.Tail.Val)
	}
}
