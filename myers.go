package godiff

import (
	"fmt"
	"slices"
)

type myersNode struct {
	k int
	x int
}

// Run the O((M+N)D) Myers Diff algorithm
// and return a list of patches between `a` and `b`.
// This uses O(D^2) space where D is the edit distance.
// https://publications.mpi-cbg.de/Myers_1986_6330.pdf
func MyersDiff(a, b []string) []Patch {
	// Compute the forward search in the edit graph
	// returning the trace of (d, k) starting at depth=-1
	// to depth=d - 1 where d is the shortest edit length
	shortestPath := myersShortestPath(a, b)
	fmt.Println("shortestPath: ", shortestPath)

	return []Patch{}
}

// Search the edit graph for the shortest path
// from `a` to `b`. Returns the list of myers frontiers.
func myersShortestPath(a, b []string) []myersNode {
	m, n := len(a), len(b)

	// the maximum edit length is composed of m deletions and n insertions
	maxDepth := m + n

	// frontier to store (k, x), the furthest
	// point on the edit graph for a given diagonal k
	// where k = x - y
	f := map[int]int{}
	f[1] = 0 // store (k=1, x=0, y=-1)

	// parents map a node to its parent in the edit path
	parents := map[myersNode]myersNode{}

	for d := 0; d <= maxDepth; d++ {
		// trace = append(trace, maps.Clone(f))
		for k := -d; k <= d; k += 2 {
			var parentMyersNode myersNode
			var x int
			if k == -d || (k != d && f[k-1] < f[k+1]) {
				// insertion advances along y and decrement k
				// so k + 1 becomes k and x stays the same
				parentMyersNode = myersNode{k: k + 1, x: f[k+1]}
				x = f[k+1]
			} else {
				// deletion advances along x and increments k
				// so k - 1 becomes k and x becomes x + 1
				parentMyersNode = myersNode{k: k - 1, x: f[k-1]}
				x = f[k-1] + 1
			}

			// rearranging k = x - y to get y = x - k
			y := x - k

			// walk along snake which is 0 cost
			// snake is a sequence of equal strings
			// between a and b in the edit graph
			for x < m && y < n && a[x] == b[y] {
				x++
				y++
			}

			// save (k, x)
			f[k] = x
			node := myersNode{k: k, x: x}
			if parentMyersNode.x-parentMyersNode.k >= 0 {
				parents[node] = parentMyersNode
			}

			// we have found the shortest edit path
			if x >= m && y >= n {
				path := myersBacktrack(parents, m, n)
				return path
			}
		}
	}

	// this should be impossible
	return []myersNode{}
}

// Return the edit path (k0, x0), (k1, x1) ...
// m - length of `a` src string sequence
// n - length of `n` dest string sequence
func myersBacktrack(parents map[myersNode]myersNode, m int, n int) []myersNode {
	node := myersNode{k: m - n, x: m}
	path := []myersNode{node}

	// get the parent, then iteratively get the parent
	// of the parent until we reach the start node (k=0, x=0)
	for {
		parent, exists := parents[node]
		if !exists {
			break
		}
		path = append(path, parent)
		node = parent
	}

	slices.Reverse(path)
	return path
}

// Construct the patches from the edit path
// returned by myersBacktrack
func patchesFromEditPath() {

}
