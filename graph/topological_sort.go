package graph

import (
	"github.com/ZhengjunHUO/goutil/datastruct"
)

type Dag struct {
	// graph's size
	size int
	// adjacent table for each vertex
	adjTable map[int][]int
	// stock the sort's result, "source" vertex on the top, the vertex "pointed to" at the bottom
	// which means the least dependent vertex will be poped out first 
	Sorted *datastruct.Stack[int]
	// to indicate if the vertex has been visited
	visited []bool
}

// NewDag initializes the graph
func NewDag(size int) *Dag {
	return &Dag{
		size: size,
		adjTable: make(map[int][]int),
		Sorted: datastruct.NewStack([]int{}),
		visited: make([]bool, size),
	}
}

// AddEdge builds the adjacent table
func (d *Dag) AddEdge(fromV, toV int) {
	if _, ok := d.adjTable[fromV]; !ok {
		d.adjTable[fromV] = []int{toV}
	}else{
		d.adjTable[fromV] = append(d.adjTable[fromV], toV)
	}
}

// findPath recursively goes through the current vertex' neighbours, and save the result to stack
func (d *Dag) findPath(vertex int) {
	d.visited[vertex] = true

	if v, ok := d.adjTable[vertex]; ok {
		for i := range v {
			if !d.visited[v[i]] {
				d.findPath(v[i])
			}
		}
	}

	d.Sorted.Push(vertex)
}

// TopologicalSort sorts the graph after the construction of adjacent table
func (d *Dag) TopologicalSort() {
	for i:=0; i<d.size; i++ {
		if !d.visited[i] {
			d.findPath(i)
		}
	}

}
