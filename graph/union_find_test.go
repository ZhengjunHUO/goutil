package graph

import (
	"testing"
	"reflect"
)

func TestUnionFind(t *testing.T) {
	uf := NewUF(10)

	uf.Union(0,1)
	uf.Union(6,0)
	uf.Union(2,3)
	uf.Union(2,5)
	uf.Union(1,3)
	uf.FindRoot(3)

	// do a snapshot
	c, p, s := uf.Count(), uf.Parent(), uf.Size()
	expectedC := 5
	expectedP := []int{0,0,0,0,4,2,0,7,8,9}
	expectedS := []int{6,1,2,1,1,1,1,1,1,1}

	if c != expectedC || !reflect.DeepEqual(p, expectedP) || !reflect.DeepEqual(s, expectedS) {
		t.Errorf("Union find returns count: %v, parent: %v, size: %v\nexpect count: %v, parent: %v, size: %v\n",
			c, p, s, expectedC, expectedP, expectedS)
	}

	uf.Union(6,5)
	if uf.IsLinked(3,5) != true {
		t.Errorf("Node 3 and 5 should be linked!\n")
	}

	// restore the snapshot
	uf.SetParent(p)
	uf.SetSize(s)
	uf.SetCount(c)
	if uf.Count() != c || !reflect.DeepEqual(uf.Parent(), p) || !reflect.DeepEqual(uf.Size(), s) {
		t.Errorf("Union find returns count: %v, parent: %v, size: %v\nexpect count: %v, parent: %v, size: %v\n",
			uf.Count(), uf.Parent(), uf.Size(), c, p, s)
	}
}
