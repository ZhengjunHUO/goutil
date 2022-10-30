package graph

import (
	"testing"
	"reflect"
)

func TestTopologicalSort(t *testing.T) {
	g := NewDag(6)
	g.AddEdge(5, 2)
        g.AddEdge(5, 0)
        g.AddEdge(4, 0)
        g.AddEdge(4, 1)
        g.AddEdge(2, 3)
        g.AddEdge(3, 1)
	g.TopologicalSort()

	expect := []int{5,4,2,3,1,0}

	result := make([]int, 0, g.Sorted.Size())
	for g.Sorted.Size() > 0 {
		result = append(result, g.Sorted.Pop())
	}

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("TopologicalSort returns %v, expect: %v", result, expect)
	}
}
