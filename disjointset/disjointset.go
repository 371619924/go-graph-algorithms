package disjointset

// DisjointSet is the interface for the disjoint-set (or union-find) data
// structure.
// Do not change the definition of this interface.
type DisjointSet interface {
	// UnionSet(s, t) merges (unions) the sets containing s and t,
	// and returns the representative of the resulting merged set.
	UnionSet(int, int) int
	// FindSet(s) returns representative of the class that s belongs to.
	FindSet(int) int
}

type disjointSet struct {
	parent map[int]int
	size   map[int]int
}

func (d *disjointSet) makeSet(x int) {
	if _, exists := d.parent[x]; !exists {
		d.parent[x] = x
		d.size[x] = 1
	}
}

func (d *disjointSet) FindSet(x int) int {
	d.makeSet(x)

	if d.parent[x] != x {
		d.parent[x] = d.FindSet(d.parent[x])
	}
	return d.parent[x]
}

func (d *disjointSet) UnionSet(a, b int) int {
	rootA := d.FindSet(a)
	rootB := d.FindSet(b)

	if rootA == rootB {
		return rootA
	}

	if d.size[rootA] < d.size[rootB] {
		rootA, rootB = rootB, rootA
	}

	d.parent[rootB] = rootA
	d.size[rootA] += d.size[rootB]

	return rootA
}

// TODO: implement a type that satisfies the DisjointSet interface.

// NewDisjointSet creates a struct of a type that satisfies the DisjointSet interface.
func NewDisjointSet() DisjointSet {
	return &disjointSet{
		parent: make(map[int]int),
		size:   make(map[int]int),
	}
}
