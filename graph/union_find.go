package graph

type UF struct {
	// number of trees
	count	int
	// parent of each nodes
	parent	[]int
	// size of nodes' descendents
	size	[]int
}

// NewUF initializes a UF struct
func NewUF(n int) *UF {
	parent, size := make([]int, n), make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}

	return &UF{n, parent, size}
}

// FindRoot finds current node's root, in the same time reduce tree's height
// After FindRoot the current tree got at most 2 levels
func (uf *UF) FindRoot(a int) int {
	for a != uf.parent[a] {
		if uf.parent[a] != uf.parent[uf.parent[a]] {
			// if the current node got a grand parent
			// attach the current node directly to the grand parent
			// uf.size[a] here should be 1
			// because a here could only be a size 1 leaf
			uf.size[uf.parent[a]] -= uf.size[a]
			uf.parent[a] = uf.parent[uf.parent[a]]
		}
		a = uf.parent[a]
	}

	return a
}

// Union connects two node by joining their roots
// Do nothing if they are already connected
// The tree generated by Union got at most 3 levels
func (uf *UF) Union(a, b int) {
	// ar, br are roots of trees at most 2 levels
	ar, br := uf.FindRoot(a), uf.FindRoot(b)
	if ar == br {
		return
	}

	// if both trees have 2 levels, the Union returns a 3 level tree
	if uf.size[ar] < uf.size[br] {
		uf.parent[ar] = br
		uf.size[br] += uf.size[ar]
	}else{
		uf.parent[br] = ar
		uf.size[ar] += uf.size[br]
	}

	uf.count--
}

// IsLinked test if two nodes are already connected
func (uf *UF) IsLinked(a, b int) bool {
	return uf.FindRoot(a) == uf.FindRoot(b)
}

// Count returns number of trees
func (uf *UF) Count() int {
	return uf.count
}

// Parent returns parent of each nodes
func (uf *UF) Parent() []int {
	p := make([]int, len(uf.parent))
	copy(p, uf.parent)
	return p
}

// Size returns size of nodes' descendents (tree/subtree)
func (uf *UF) Size() []int {
	s := make([]int, len(uf.size))
	copy(s, uf.size)
	return s
}

// SetCount sets UF's number of trees
func (uf *UF) SetCount(c int) {
	uf.count = c
}

// SetParent sets parent of each nodes
func (uf *UF) SetParent(p []int) {
	copy(uf.parent, p)
}

// SetSize sets size of nodes' descendents
func (uf *UF) SetSize(s []int) {
	copy(uf.size, s)
}
